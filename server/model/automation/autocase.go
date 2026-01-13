// Package automation 自动生成模板AutoCase
package automation

import (
	"github.com/test-instructor/yangfan/server/v2/global"
)

// AutoCase 测试用例 结构体  AutoCase
type AutoCase struct {
	global.GVA_MODEL
	CaseName   string `json:"caseName" form:"caseName" gorm:"column:case_name;" binding:"required"`                                          //用例名称
	RunNumber  int64  `json:"runNumber" form:"runNumber" gorm:"column:run_number;"`                                                          //运行次数
	Status     string `json:"status" form:"status" gorm:"column:status;type:enum('测试中','待评审','评审不通过','已发布','禁用','已废弃');" binding:"required"` //状态
	ConfigID   int64  `json:"configID" form:"configID" gorm:"column:config_id;"`                                                             //运行配置
	EnvName    string `json:"envName" form:"envName" gorm:"column:env_name;"`                                                                //运行环境
	EnvID      int64  `json:"envID" form:"envID" gorm:"column:env_id;"`                                                                      //运行环境ID
	Desc       string `json:"desc" form:"desc" gorm:"column:desc;"`                                                                          //描述
	ConfigName string `json:"configName" form:"configName" gorm:"column:config_name;"`                                                       //配置名称
	Menu       int64  `json:"menu" form:"menu" gorm:"column:menu;"`
	ProjectId  int64  `json:"projectId" form:"projectId" gorm:"column:project_id;"` //项目信息
}

// TableName 测试用例 AutoCase自定义表名 auto_case
func (AutoCase) TableName() string {
	return "yf_auto_case"
}

type AutoCaseStepList struct {
	global.GVA_MODEL
	AutoCase       AutoCase     `json:"case" form:"case"`
	AutoCaseID     uint         `json:"case_id" form:"case_id" gorm:"column:case_id;"`
	AutoCaseStep   AutoCaseStep `json:"step" form:"step" `
	AutoCaseStepID uint         `json:"step_id" form:"step_id" gorm:"column:step_id;"`
	ConfigID       int64        `json:"config_id" form:"config_id" gorm:"column:config_id;"`
	Sort           int64        `json:"sort" form:"sort" gorm:"column:sort;"`
	IsConfig       bool         `json:"isConfig" form:"isConfig" gorm:"column:is_config;"`
	IsStepConfig   bool         `json:"isStepConfig" form:"isStepConfig" gorm:"column:is_step_config;default:false"`
}

func (AutoCaseStepList) TableName() string {
	return "yf_auto_case_step_list"
}
