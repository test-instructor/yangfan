// 自动生成模板AutoCaseStep
package automation

import (
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	"gorm.io/datatypes"
)

// 测试步骤 结构体  AutoCaseStep
type AutoCaseStep struct {
	global.GVA_MODEL
	StepConfig

	TestCase   []AutoStep `json:"testcase,omitempty" form:"testcase;comment:http接口;" gorm:"-;"`
	ConfigName string     `json:"configName" form:"configName" gorm:"column:config_name;"`
	ConfigID   int64      `json:"configID" form:"configID" gorm:"column:config_id;"` //运行配置ID
	Config     *platform.RunConfig
	EnvName    string `json:"envName" form:"envName" gorm:"column:env_name;"`       //运行环境
	EnvID      int64  `json:"envID" form:"envID" gorm:"column:env_id;"`             //运行环境ID
	ProjectId  int64  `json:"projectId" form:"projectId" gorm:"column:project_id;"` //项目信息
	Menu       int64  `json:"menu" form:"menu" gorm:"column:menu;"`
	ParentId   uint   `json:"parentId" form:"parentId" gorm:"column:parent_id;"`
}

// TableName 测试步骤 AutoCaseStep自定义表名 AutoCaseStep
func (AutoCaseStep) TableName() string {
	return "lc_auto_case_step"
}

type AutoCaseStepRelation struct {
	global.GVA_MODEL
	AutoCaseStep   AutoCaseStep // 测试步骤集合
	AutoCaseStepID uint         `json:"autoCaseStepID" form:"autoCaseStepID" gorm:"column:auto_case_step_id;"`
	AutoStep       AutoStep     // 测试步骤
	AutoStepID     uint         `json:"autoStepID" form:"autoStepID" gorm:"column:auto_step_id;"`
	Sort           uint         `json:"sort" form:"sort" gorm:"column:sort;"`
	ProjectId      int64        `json:"projectId" form:"projectId" gorm:"column:project_id;"`
}

func (AutoCaseStepRelation) TableName() string {
	return "lc_auto_case_step_relation"
}

type CaseStepConfig struct {
	StepName          string            `json:"name" form:"name" gorm:"column:step_name;" binding:"required"` //步骤名称
	Variables         datatypes.JSONMap `json:"variables" form:"variables" gorm:"column:variables;"`          //变量
	VariablesTemp     datatypes.JSON    `json:"variables_temp" form:"variables_temp" gorm:"column:variables_temp;"`
	Parameters        datatypes.JSONMap `json:"parameters" form:"parameters" gorm:"column:parameters;"`                                        //参数
	ParametersTemp    datatypes.JSONMap `json:"parameters_temp" form:"parameters_temp" gorm:"column:parameters_temp;"`                         //参数
	DataWarehouse     datatypes.JSONMap `json:"data_warehouse" form:"data_warehouse" gorm:"column:data_warehouse;"`                            //数据仓库
	DataWarehouseTemp datatypes.JSONMap `json:"data_warehouse_temp" form:"data_warehouse_temp" gorm:"column:data_warehouse_temp;"`             //数据仓库临时
	SetupHooks        datatypes.JSON    `json:"setup_hooks" form:"setup_hooks" gorm:"column:setup_hooks;" swaggertype:"array,object"`          //设置钩子
	TeardownHooks     datatypes.JSON    `json:"teardown_hooks" form:"teardown_hooks" gorm:"column:teardown_hooks;" swaggertype:"array,object"` //清理钩子
	Extract           datatypes.JSONMap `json:"extract" form:"extract" gorm:"column:extract;"`                                                 //提取
	ExtractTemp       datatypes.JSON    `json:"extract_temp" form:"extract_temp" gorm:"column:extract_temp;"`
	Validators        datatypes.JSON    `json:"validate" form:"validate" gorm:"column:validators;" swaggertype:"array,object"` //验证器
	ValidatorsTemp    datatypes.JSON    `json:"validators_temp" form:"validators_temp" gorm:"column:validators_temp;"`
	StepExport        datatypes.JSON    `json:"export" form:"export" gorm:"column:step_export;" swaggertype:"array,object"` //步骤导出
	Skip              datatypes.JSON    `json:"skip" form:"skip" gorm:"column:skip;"`                                       //变量
	SkipTemp          datatypes.JSON    `json:"skip_temp" form:"skip_temp" gorm:"column:skip_temp;"`
	Loops             int64             `json:"loops" form:"loops" gorm:"column:loops;"`                                                            //循环次数
	IgnorePopup       bool              `json:"ignore_popup" form:"ignore_popup" gorm:"column:ignore_popup;"`                                       //忽略弹出窗口
	AutoPopupHandler  bool              `json:"auto_popup_handler,omitempty" yaml:"auto_popup_handler,omitempty" gorm:"column:auto_popup_handler;"` // enable auto popup handler for this step
	Sort              int               `json:"sort" form:"sort" gorm:"column:sort;"`
	Retry             int               `json:"retry" form:"retry" gorm:"column:retry;"`
}
