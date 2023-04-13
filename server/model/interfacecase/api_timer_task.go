// 自动生成模板TimerTask
package interfacecase

import (
	"time"

	"github.com/test-instructor/yangfan/server/global"
)

// ApiTimerTask  结构体
// 如果含有time.Time 请自行import time包
type ApiTimerTask struct {
	global.GVA_MODEL
	Operator
	Name            string            `json:"name" form:"name" gorm:"column:name;comment:任务名称;"`
	RunTime         string            `json:"runTime" form:"runTime" gorm:"column:run_time;comment:运行时间cron;"`
	NextRunTime     *time.Time        `json:"nextRunTime" form:"nextRunTime" gorm:"column:next_run_time;comment:下次运行时间;"`
	Status          *bool             `json:"status" form:"status" gorm:"column:status;comment:运行状态;"`
	Describe        string            `json:"describe" form:"describe" gorm:"column:describe;comment:备注;"`
	RunNumber       *int              `json:"runNumber" form:"runNumber" gorm:"column:run_number;comment:运行次数;"`
	RunConfig       ApiConfig         `json:"config" form:"config"`
	RunConfigID     uint              `json:"RunConfigID" form:"RunConfigID" gorm:"comment:运行配置"`
	TestCase        []*ApiCase        `json:"TestCase" form:"TestCase" gorm:"many2many:ApiTimerTaskRelationship;"`
	EntryID         int               `json:"-"`
	ApiTimerTaskTag []ApiTimerTaskTag `json:"apiTimerTaskTag" form:"apiTimerTaskTag" gorm:"many2many:ApiTimerTaskTagRelationship;"`
	TagIds          []uint            `json:"tagIds" gorm:"-"`
	ApiEnvName      *string           `json:"api_env_name" gorm:"comment:所属环境名称;"`
	ApiEnvID        uint              `json:"api_env_id" gorm:"comment:所属环境;"`
}

type ApiTimerTaskTag struct {
	global.GVA_MODEL
	Operator
	ApiTimerTask []ApiTimerTask `json:"apiTimerTask" form:"apiTimerTask" gorm:"many2many:ApiTimerTaskTagRelationship;"`
	Name         string         `json:"name"`
	Remarks      string         `json:"remarks"`
}

type ApiTimerTaskTagRelationship struct {
	ApiTimerTask      ApiTimerTask
	ApiTimerTaskId    uint `gorm:"comment:定时任务"`
	ApiTimerTaskTag   ApiTimerTaskTag
	ApiTimerTaskTagId uint `gorm:"comment:定时任务标签"`
}
