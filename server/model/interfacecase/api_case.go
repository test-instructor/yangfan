// 自动生成模板ApiCase
package interfacecase

import (
	"github.com/test-instructor/yangfan/server/global"
)

// ApiCase 结构体
// 如果含有time.Time 请自行import time包
type ApiCase struct {
	global.GVA_MODEL
	Operator
	Name        string        `json:"name" form:"name" gorm:"column:name;comment:用例名称;"`
	Status      *bool         `json:"status" form:"status" gorm:"column:status;comment:状态;"`
	Describe    string        `json:"describe" form:"describe" gorm:"column:describe;comment:备注;"`
	RunNumber   *int          `json:"runNumber" form:"runNumber" gorm:"column:run_number;comment:运行次数;"`
	RunConfig   *ApiConfig    `json:"runConfig" form:"runConfig"`
	RunConfigID uint          `json:"RunConfigID" form:"RunConfigID" gorm:"column:RunConfigID;comment:运行配置;"`
	ApiCaseStep []ApiCaseStep `json:"case" form:"case" gorm:"many2many:ApiCaseRelationship;"`
	//Performance []Performance   `json:"performance" form:"performance" gorm:"many2many:PerformanceRelationship;"`
	FrontCase  *bool    `json:"front_case" form:"front_case" gorm:"comment:是否为前置用例"`
	EntryID    int      `json:"-"`
	ApiMenuID  uint     `json:"-" gorm:"comment:所属菜单"`
	ApiMenu    *ApiMenu `json:"-"`
	ApiEnvName *string  `json:"api_env_name" gorm:"comment:所属环境名称;"`
	ApiEnvID   uint     `json:"api_env_id" gorm:"comment:所属环境;"`
}
