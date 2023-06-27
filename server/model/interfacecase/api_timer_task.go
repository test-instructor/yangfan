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
	Name            string            `json:"name,omitempty" form:"name" gorm:"column:name;comment:任务名称;"`
	RunTime         string            `json:"runTime,omitempty" form:"runTime" gorm:"column:run_time;comment:运行时间cron;"`
	NextRunTime     *time.Time        `json:"nextRunTime,omitempty" form:"nextRunTime" gorm:"column:next_run_time;comment:下次运行时间;"`
	Status          *bool             `json:"status,omitempty" form:"status" gorm:"column:status;comment:运行状态;"`
	Describe        string            `json:"describe,omitempty" form:"describe" gorm:"column:describe;comment:备注;"`
	RunNumber       *int              `json:"runNumber,omitempty" form:"runNumber" gorm:"column:run_number;comment:运行次数;"`
	RunConfig       ApiConfig         `json:"config,omitempty" form:"config"`
	RunConfigID     uint              `json:"RunConfigID,omitempty" form:"RunConfigID" gorm:"comment:运行配置"`
	TestCase        []*ApiCase        `json:"TestCase,omitempty" form:"TestCase" gorm:"many2many:ApiTimerTaskRelationship;"`
	EntryID         int               `json:"-"`
	ApiTimerTaskTag []ApiTimerTaskTag `json:"apiTimerTaskTag,omitempty" form:"apiTimerTaskTag" gorm:"many2many:ApiTimerTaskTagRelationship;"`
	TagIds          []uint            `json:"tagIds" gorm:"-"`
	ApiEnvName      *string           `json:"api_env_name" gorm:"comment:所属环境名称;"`
	ApiEnvID        uint              `json:"api_env_id" gorm:"comment:所属环境;"`
}

type ApiTimerTaskTag struct {
	global.GVA_MODEL
	Operator
	ApiTimerTask []ApiTimerTask `json:"apiTimerTask,omitempty" form:"apiTimerTask" gorm:"many2many:ApiTimerTaskTagRelationship;"`
	Name         string         `json:"name"`
	Remarks      string         `json:"remarks"`
}

type ApiTimerTaskTagRelationship struct {
	ApiTimerTask      ApiTimerTask
	ApiTimerTaskId    uint `gorm:"comment:定时任务"`
	ApiTimerTaskTag   ApiTimerTaskTag
	ApiTimerTaskTagId uint `gorm:"comment:定时任务标签"`
}
