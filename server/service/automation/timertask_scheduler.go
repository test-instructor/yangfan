package automation

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/automation"
	platformReq "github.com/test-instructor/yangfan/server/v2/model/platform/request"
	platformService "github.com/test-instructor/yangfan/server/v2/service/platform"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const timerTaskCronName = "AutomationTimerTask"

var timerTaskCronParser = cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)

func RefreshTimerTaskSchedules(ctx context.Context) {
	global.GVA_Timer.Clear(timerTaskCronName)

	var tasks []automation.TimerTask
	db := global.GVA_DB.WithContext(ctx).Model(&automation.TimerTask{}).
		Where("status = ?", true).
		Where("run_time IS NOT NULL AND run_time <> ''").
		Where("(runner_node_name IS NULL OR runner_node_name = '')")
	if err := db.Find(&tasks).Error; err != nil {
		return
	}

	now := time.Now()
	for _, t := range tasks {
		if t.RunTime == nil || strings.TrimSpace(*t.RunTime) == "" {
			continue
		}
		spec := normalizeCronSpec(*t.RunTime)
		taskID := t.ID
		taskName := fmt.Sprintf("timer_task_%d", taskID)

		_, err := global.GVA_Timer.AddTaskByFuncWithSecond(timerTaskCronName, spec, func() {
			runTimerTask(taskID, spec)
		}, taskName)
		if err != nil {
			continue
		}

		next := calcNext(spec, now)
		if next != nil {
			_ = global.GVA_DB.WithContext(ctx).Model(&automation.TimerTask{}).
				Where("id = ?", taskID).
				Update("next_run_time", *next).Error
		}
	}
}

func runTimerTask(taskID uint, spec string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var task automation.TimerTask
	err := global.GVA_DB.WithContext(ctx).
		Select("id", "status", "project_id", "env_id", "config_id", "run_time").
		Where("id = ?", taskID).
		First(&task).Error
	if err != nil {
		return
	}
	if task.Status == nil || !*task.Status {
		return
	}

	envID := 0
	if task.EnvID != nil {
		envID = int(*task.EnvID)
	}
	configID := 0
	if task.ConfigID != nil {
		configID = int(*task.ConfigID)
	}

	req := platformReq.RunnerRequest{
		CaseType:  "task",
		CaseID:    task.ID,
		RunMode:   "定时执行",
		EnvID:     envID,
		ConfigID:  configID,
		ProjectId: uint(task.ProjectId),
	}

	if _, err := (&platformService.RunnerService{}).RunTask(req); err != nil {
		global.GVA_LOG.Warn("timer task run failed", zap.Uint("task_id", taskID), zap.Error(err))
	}

	_ = global.GVA_DB.WithContext(ctx).Model(&automation.TimerTask{}).
		Where("id = ?", taskID).
		UpdateColumn("run_number", gorm.Expr("COALESCE(run_number,0)+1")).Error

	next := calcNext(spec, time.Now())
	if next != nil {
		_ = global.GVA_DB.WithContext(ctx).Model(&automation.TimerTask{}).
			Where("id = ?", taskID).
			Update("next_run_time", *next).Error
	}
}

func normalizeCronSpec(spec string) string {
	s := strings.TrimSpace(spec)
	if s == "" || strings.HasPrefix(s, "@") {
		return s
	}
	fields := strings.Fields(s)
	if len(fields) == 5 {
		return "0 " + s
	}
	return s
}

func calcNext(spec string, from time.Time) *time.Time {
	schedule, err := timerTaskCronParser.Parse(spec)
	if err != nil {
		return nil
	}
	next := schedule.Next(from)
	return &next
}
