// 自动生成模板TimerTask
package interfacecase

import (
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/model/system"
	"time"
)

// TimerTask 结构体
// 如果含有time.Time 请自行import time包
type ApiTimerTask struct {
	global.GVA_MODEL
	Name        string         `json:"name" form:"name" gorm:"column:name;comment:;"`
	RunTime     string         `json:"runTime" form:"runTime" gorm:"column:run_time;comment:;"`
	NextRunTime *time.Time     `json:"nextRunTime" form:"nextRunTime" gorm:"column:next_run_time;comment:;"`
	Status      *bool          `json:"status" form:"status" gorm:"column:status;comment:;"`
	Describe    string         `json:"describe" form:"describe" gorm:"column:describe;comment:;"`
	RunNumber   *int           `json:"runNumber" form:"runNumber" gorm:"column:run_number;comment:;"`
	ProjectID   uint           `json:"-"`
	Project     system.Project `json:"-"`
	RunConfig   ApiConfig      `json:"config" form:"config"`
	RunConfigID uint           `json:"RunConfigID" form:"RunConfigID"`
	TestCase    []*ApiCase     `json:"TestCase" form:"TestCase" gorm:"many2many:ApiTimerTaskRelationship;"`
	EntryID     int            `json:"-"`
	CreatedBy   system.SysUser `json:"-"`
	CreatedByID uint           `json:"-"`
	UpdateBy    system.SysUser `json:"-"`
	UpdateByID  uint           `json:"-"`
	DeleteBy    system.SysUser `json:"-"`
	DeleteByID  uint           `json:"-"`
}
