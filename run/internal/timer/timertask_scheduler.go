package timer

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/test-instructor/yangfan/run/internal/runTestCase"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/automation"
	platformReq "github.com/test-instructor/yangfan/server/v2/model/platform/request"
	platformService "github.com/test-instructor/yangfan/server/v2/service/platform"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TimerTaskScheduler struct {
	nodeName string
	cronName string
}

func NewTimerTaskScheduler(nodeName string) *TimerTaskScheduler {
	return &TimerTaskScheduler{
		nodeName: strings.TrimSpace(nodeName),
		cronName: fmt.Sprintf("RunnerTimerTask_%s", strings.TrimSpace(nodeName)),
	}
}

func (s *TimerTaskScheduler) Start(ctx context.Context) {
	global.GVA_Timer.Clear(s.cronName)
	s.resync(ctx)
}

func (s *TimerTaskScheduler) Stop() {
	global.GVA_Timer.Clear(s.cronName)
}

func (s *TimerTaskScheduler) UpsertTask(ctx context.Context, taskID uint) {
	s.upsertFromDB(ctx, taskID)
}

func (s *TimerTaskScheduler) RemoveTask(taskID uint) {
	taskName := fmt.Sprintf("timer_task_%d", taskID)
	global.GVA_Timer.RemoveTaskByName(s.cronName, taskName)
}

func (s *TimerTaskScheduler) resync(ctx context.Context) {
	if s.nodeName == "" {
		return
	}
	var tasks []automation.TimerTask
	if err := global.GVA_DB.WithContext(ctx).
		Model(&automation.TimerTask{}).
		Where("runner_node_name = ?", s.nodeName).
		Where("status = ?", true).
		Where("run_time IS NOT NULL AND run_time <> ''").
		Find(&tasks).Error; err != nil {
		return
	}
	now := time.Now()
	for _, t := range tasks {
		if t.RunTime == nil || strings.TrimSpace(*t.RunTime) == "" {
			continue
		}
		spec := normalizeCronSpec(*t.RunTime)
		s.addOrReplaceTask(ctx, t.ID, spec, now)
	}
}

func (s *TimerTaskScheduler) upsertFromDB(ctx context.Context, taskID uint) {
	var task automation.TimerTask
	err := global.GVA_DB.WithContext(ctx).
		Select("id", "status", "run_time", "runner_node_name").
		First(&task, "id = ?", taskID).Error
	if err != nil {
		return
	}
	if task.RunnerNodeName == nil || strings.TrimSpace(*task.RunnerNodeName) != s.nodeName {
		s.RemoveTask(taskID)
		return
	}

	enabled := task.Status != nil && *task.Status
	if !enabled || task.RunTime == nil || strings.TrimSpace(*task.RunTime) == "" {
		s.RemoveTask(taskID)
		return
	}

	spec := normalizeCronSpec(*task.RunTime)
	s.addOrReplaceTask(ctx, taskID, spec, time.Now())
}

func (s *TimerTaskScheduler) addOrReplaceTask(ctx context.Context, taskID uint, spec string, now time.Time) {
	taskName := fmt.Sprintf("timer_task_%d", taskID)
	global.GVA_Timer.RemoveTaskByName(s.cronName, taskName)

	_, err := global.GVA_Timer.AddTaskByFuncWithSecond(s.cronName, spec, func() {
		s.runTimerTask(taskID, spec)
	}, taskName)
	if err != nil {
		return
	}

	next := calcNext(spec, now)
	if next != nil {
		_ = global.GVA_DB.WithContext(ctx).Model(&automation.TimerTask{}).
			Where("id = ?", taskID).
			Update("next_run_time", *next).Error
	}
}

func (s *TimerTaskScheduler) runTimerTask(taskID uint, spec string) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var task automation.TimerTask
	err := global.GVA_DB.WithContext(ctx).
		Select("id", "status", "project_id", "env_id", "config_id", "run_time", "runner_node_name").
		Where("id = ?", taskID).
		First(&task).Error
	if err != nil {
		return
	}
	if task.RunnerNodeName == nil || strings.TrimSpace(*task.RunnerNodeName) != s.nodeName {
		s.RemoveTask(taskID)
		return
	}
	if task.Status == nil || !*task.Status {
		s.RemoveTask(taskID)
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
		NodeName:  s.nodeName,
	}

	report, err := (&platformService.RunnerService{}).CreatePendingReport(req)
	if err != nil {
		global.GVA_LOG.Warn("timer task create report failed", zap.Uint("task_id", taskID), zap.Error(err))
		return
	}
	req.ReportID = report.ID

	msgBody := fmt.Sprintf(`{"source":"timer","task_id":%d,"node_name":"%s"}`, taskID, s.nodeName)
	if _, err := runTestCase.Entry(req, &msgBody); err != nil {
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

var timerTaskCronParser = cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)

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
