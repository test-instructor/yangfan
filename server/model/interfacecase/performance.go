package interfacecase

import (
	"github.com/test-instructor/yangfan/server/global"
)

type Performance struct {
	global.GVA_MODEL
	Operator
	Name        string     `json:"name" form:"name" gorm:"column:name;comment:性能任务名称;"`
	State       State      `json:"status" form:"status" gorm:"column:status;comment:状态;"`
	Describe    string     `json:"describe" form:"describe" gorm:"column:describe;comment:备注;"`
	RunConfig   *ApiConfig `json:"config" form:"config"`
	RunConfigID uint       `json:"RunConfigID" form:"RunConfigID" gorm:"column:RunConfigID;comment:运行配置;"`

	FrontCase           *bool              `json:"front_case" form:"front_case"`
	EntryID             int                `json:"-"`
	ApiMenuID           uint               `json:"-" gorm:"comment:所属菜单"`
	PerformanceReport   *PerformanceReport `json:"Report"`
	PerformanceReportId uint               `json:"report_id"`
	ApiEnvName          *string            `json:"api_env_name" gorm:"comment:所属环境名称;"`
	ApiEnvID            uint               `json:"api_env_id" gorm:"comment:所属环境;"`
}
