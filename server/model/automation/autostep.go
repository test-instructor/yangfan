// Package automation 自动生成模板AutoStep
package automation

import (
	"github.com/test-instructor/yangfan/server/v2/global"
	"gorm.io/datatypes"
)

// AutoStep 自动化步骤 结构体  AutoStep
type AutoStep struct {
	global.GVA_MODEL
	StepConfig

	Request   *Request `json:"request,omitempty" form:"request;comment:http接口" mapstructure:"request"`
	RequestID uint     `json:"request_id,omitempty" gorm:"comment:http请求"`

	Android *MobileStep `json:"android,omitempty" gorm:"column:android;type:json;serializer:json"`
	IOS     *MobileStep `json:"ios,omitempty" gorm:"column:ios;type:json;serializer:json"`
	Harmony *MobileStep `json:"harmony,omitempty" gorm:"column:harmony;type:json;serializer:json"`
	Browser *MobileStep `json:"browser,omitempty" gorm:"column:browser;type:json;serializer:json"`

	ProjectId int64 `json:"projectId" form:"projectId" gorm:"column:project_id;"` //项目信息
	Menu      int64 `json:"menu" form:"menu" gorm:"column:menu;"`
	ParentId  uint  `json:"parentId" form:"parentId" gorm:"column:parent_id;"`
	StepType  int   `json:"type" form:"type" gorm:"column:step_type;"`
}

// TableName 自动化步骤 AutoStep自定义表名 AutoStep
func (AutoStep) TableName() string {
	return "yf_auto_step"
}

type StepConfig struct {
	StepName          string            `json:"name" form:"name" mapstructure:"name" gorm:"column:step_name;" binding:"required"` //步骤名称
	Variables         datatypes.JSONMap `json:"variables" form:"variables" mapstructure:"variables" gorm:"column:variables;"`     //变量
	VariablesTemp     datatypes.JSON    `json:"variables_temp" form:"variables_temp" gorm:"column:variables_temp;"`
	Parameters        datatypes.JSONMap `json:"parameters" form:"parameters" mapstructure:"parameters" gorm:"column:parameters;"`                                            //参数
	ParametersTemp    datatypes.JSONMap `json:"parameters_temp" form:"parameters_temp" gorm:"column:parameters_temp;"`                                                       //参数
	DataWarehouse     datatypes.JSONMap `json:"data_warehouse" form:"data_warehouse" mapstructure:"data_warehouse" gorm:"column:data_warehouse;"`                            //数据仓库
	DataWarehouseTemp datatypes.JSONMap `json:"data_warehouse_temp" form:"data_warehouse_temp" gorm:"column:data_warehouse_temp;"`                                           //数据仓库临时
	SetupHooks        datatypes.JSON    `json:"setup_hooks" form:"setup_hooks" mapstructure:"setup_hooks" gorm:"column:setup_hooks;" swaggertype:"array,object"`             //设置钩子
	TeardownHooks     datatypes.JSON    `json:"teardown_hooks" form:"teardown_hooks" mapstructure:"teardown_hooks" gorm:"column:teardown_hooks;" swaggertype:"array,object"` //清理钩子
	Extract           datatypes.JSONMap `json:"extract" form:"extract" mapstructure:"extract" gorm:"column:extract;"`                                                        //提取
	ExtractTemp       datatypes.JSON    `json:"extract_temp" form:"extract_temp" gorm:"column:extract_temp;"`
	Validators        datatypes.JSON    `json:"validate" form:"validate" mapstructure:"validate" gorm:"column:validators;" swaggertype:"array,object"` //验证器
	ValidatorsTemp    datatypes.JSON    `json:"validators_temp" form:"validators_temp" gorm:"column:validators_temp;"`
	StepExport        datatypes.JSON    `json:"export" form:"export" mapstructure:"export" gorm:"column:step_export;" swaggertype:"array,object"`
	Skip              datatypes.JSON    `json:"skip" form:"skip" gorm:"column:skip;"`                                                               //变量
	SkipTemp          datatypes.JSON    `json:"skip_temp" form:"skip_temp" gorm:"column:skip_temp;"`                                                //步骤导出
	Loops             int64             `json:"loops" form:"loops" mapstructure:"loops" gorm:"column:loops;"`                                       //循环次数
	IgnorePopup       bool              `json:"ignore_popup" form:"ignore_popup" gorm:"column:ignore_popup;"`                                       //忽略弹出窗口
	AutoPopupHandler  bool              `json:"auto_popup_handler,omitempty" yaml:"auto_popup_handler,omitempty" gorm:"column:auto_popup_handler;"` // enable auto popup handler for this step
	Sort              int               `json:"sort" form:"sort" gorm:"column:sort;"`
	Retry             int               `json:"retry" form:"retry" gorm:"column:retry;"`
}

type Test struct {
	NameTemp datatypes.JSON    `json:"name_temp" form:"name_temp" gorm:"column:name_temp;"`
	Name     datatypes.JSONMap `json:"name" form:"name" gorm:"column:-;"`
}
