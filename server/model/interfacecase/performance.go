package interfacecase

import (
	"github.com/test-instructor/yangfan/server/global"
)

type Performance struct {
	global.GVA_MODEL
	Operator
	Name        string     `json:"name,omitempty" form:"name" gorm:"column:name;comment:性能任务名称;"`
	State       State      `json:"status,omitempty" form:"status" gorm:"column:status;comment:状态;"`
	Describe    string     `json:"describe,omitempty" form:"describe" gorm:"column:describe;comment:备注;"`
	RunConfig   *ApiConfig `json:"config,omitempty" form:"config"`
	RunConfigID uint       `json:"RunConfigID,omitempty" form:"RunConfigID" gorm:"column:RunConfigID;comment:运行配置;"`

	FrontCase           *bool              `json:"front_case,omitempty" form:"front_case"`
	EntryID             int                `json:"-"`
	PerformanceReport   *PerformanceReport `json:"Report"`
	PerformanceReportId uint               `json:"report_id"`
	ApiEnvName          *string            `json:"api_env_name" gorm:"comment:所属环境名称;"`
	ApiEnvID            uint               `json:"api_env_id" gorm:"comment:所属环境;"`
	TestCase            []*PerformanceCase `json:"-" yaml:"-" gorm:"many2many:PerformanceCaseRelationship;"`
}

type PerformanceCase struct {
	global.GVA_MODEL
	Operator
	Name     string `json:"name,omitempty" form:"name" gorm:"column:name;comment:性能任务名称;"`
	State    State  `json:"status,omitempty" form:"status" gorm:"column:status;comment:状态;"`
	Describe string `json:"describe,omitempty" form:"describe" gorm:"column:describe;comment:备注;"`

	PerformanceReport   *PerformanceReport `json:"Report"`
	PerformanceReportId uint               `json:"report_id"`
	ApiEnvName          *string            `json:"api_env_name" gorm:"comment:所属环境名称;"`
	ApiEnvID            uint               `json:"api_env_id" gorm:"comment:所属环境;"`
	TestCase            []*Performance     `json:"test_case" yaml:"test_case" gorm:"many2many:PerformanceCaseRelationship;"`
}
