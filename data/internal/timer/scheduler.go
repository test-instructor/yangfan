package timer

import (
	"fmt"

	"github.com/test-instructor/yangfan/data/internal/datacategory"
	"github.com/test-instructor/yangfan/server/v2/global"
	"go.uber.org/zap"
)

const (
	// CronName 定时任务组名称
	CronName = "DataWarehouse"
	// TaskName 任务名称
	TaskName = "DataCategoryProcess"
	// CronSpec 定时任务表达式 - 每小时执行一次
	CronSpec = "@every 1h"
	// CronSpecDebug 调试用定时任务表达式 - 每分钟执行一次
	CronSpecDebug = "@every 1m"
)

// Scheduler 定时任务调度器
type Scheduler struct {
	service   *datacategory.Service
	debugMode bool
}

// NewScheduler 创建调度器
func NewScheduler() *Scheduler {
	return &Scheduler{
		service:   datacategory.NewService(),
		debugMode: false,
	}
}

// NewDebugScheduler 创建调试模式调度器（每分钟执行）
func NewDebugScheduler() *Scheduler {
	return &Scheduler{
		service:   datacategory.NewService(),
		debugMode: true,
	}
}

// Start 启动定时任务
func (s *Scheduler) Start() error {
	spec := CronSpec
	if s.debugMode {
		spec = CronSpecDebug
	}

	global.GVA_LOG.Info("启动数据分类定时任务",
		zap.String("cronName", CronName),
		zap.String("taskName", TaskName),
		zap.String("spec", spec),
		zap.Bool("debugMode", s.debugMode),
	)

	_, err := global.GVA_Timer.AddTaskByFunc(CronName, spec, func() {
		global.GVA_LOG.Info("定时任务触发: 开始处理数据分类")
		s.service.ProcessAll()
	}, TaskName)

	if err != nil {
		global.GVA_LOG.Error("添加定时任务失败", zap.Error(err))
		return fmt.Errorf("添加定时任务失败: %w", err)
	}

	global.GVA_LOG.Info("数据分类定时任务已启动")
	return nil
}

// Stop 停止定时任务
func (s *Scheduler) Stop() {
	global.GVA_LOG.Info("停止数据分类定时任务")
	global.GVA_Timer.Clear(CronName)
}

// RunOnce 立即执行一次（用于测试或手动触发）
func (s *Scheduler) RunOnce() {
	global.GVA_LOG.Info("手动触发: 开始处理数据分类")
	s.service.ProcessAll()
}
