// 自动生成模板TimerTask
package interfacecase

import (
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/model/system"
	"time"
)

// ApiTimerTask  结构体
// 如果含有time.Time 请自行import time包
type ApiTimerTask struct {
	global.GVA_MODEL
	Operator
	Name        string         `json:"name" form:"name" gorm:"column:name;comment:任务名称;"`
	RunTime     string         `json:"runTime" form:"runTime" gorm:"column:run_time;comment:运行时间cron;"`
	NextRunTime *time.Time     `json:"nextRunTime" form:"nextRunTime" gorm:"column:next_run_time;comment:下次运行时间;"`
	Status      *bool          `json:"status" form:"status" gorm:"column:status;comment:运行状态;"`
	Describe    string         `json:"describe" form:"describe" gorm:"column:describe;comment:备注;"`
	RunNumber   *int           `json:"runNumber" form:"runNumber" gorm:"column:run_number;comment:运行次数;"`
	ProjectID   uint           `json:"-" gorm:"comment:所属项目"`
	Project     system.Project `json:"-"`
	RunConfig   ApiConfig      `json:"config" form:"config"`
	RunConfigID uint           `json:"RunConfigID" form:"RunConfigID" gorm:"comment:运行配置"`
	TestCase    []*ApiCase     `json:"TestCase" form:"TestCase" gorm:"many2many:ApiTimerTaskRelationship;"`
	EntryID     int            `json:"-"`
}
