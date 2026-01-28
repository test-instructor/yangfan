// Package platform 自动生成模板RunConfig
package platform

import (
	"github.com/test-instructor/yangfan/server/v2/global"
	"gorm.io/datatypes"
)

// 运行配置 结构体  RunConfig
type RunConfig struct {
	global.GVA_MODEL
	Name               string            `json:"name" form:"name" mapstructure:"name" gorm:"comment:配置名称;column:name;" binding:"required"`                              //名称
	Type               string            `json:"type" form:"type" gorm:"-"`                                                                                             //类型
	BaseUrl            string            `json:"base_url" form:"base_url" mapstructure:"base_url" gorm:"comment:默认域名;column:base_url;"`                                 //域名
	Variables          datatypes.JSONMap `json:"variables" form:"variables" mapstructure:"variables" gorm:"comment:变量;column:variables;" swaggertype:"object"`          //变量
	Headers            datatypes.JSONMap `json:"headers" form:"headers" mapstructure:"headers" gorm:"comment:请求头;column:headers;" swaggertype:"object"`                 //请求头
	Parameters         datatypes.JSONMap `json:"parameters" form:"parameters" mapstructure:"parameters" gorm:"comment:参数;column:parameters;" swaggertype:"object"`      //参数
	ParametersTemp     datatypes.JSONMap `json:"parameters_temp" form:"parameters_temp" gorm:"comment:参数;column:parameters_temp;" swaggertype:"object"`                 //参数
	DataWarehouse      datatypes.JSONMap `json:"data_warehouse" form:"data_warehouse" gorm:"comment:数据仓库;column:data_warehouse;" swaggertype:"object"`                  //数据仓库
	DataWarehouseTemp  datatypes.JSONMap `json:"data_warehouse_temp" form:"data_warehouse_temp" gorm:"comment:数据仓库临时;column:data_warehouse_temp;" swaggertype:"object"` //数据仓库临时
	VariableTemp       datatypes.JSON    `json:"variable_temp" form:"variable_temp" gorm:"comment:变量;column:variable_temp;" swaggertype:"object"`
	HeaderTemp         datatypes.JSON    `json:"header_temp" form:"header_temp" gorm:"column:header_temp;"`                                            //请求头JSON
	Weight             int               `json:"weight" form:"weight" gorm:"column:weight;"`                                                           //权重
	Timeout            float32           `json:"timeout" form:"timeout" mapstructure:"request_timeout" gorm:"comment:超时时间;column:timeout;"`            //超时
	AllowRedirects     bool              `json:"allow_redirects" form:"allow_redirects" mapstructure:"allow_redirects" gorm:"column:allow_redirects;"` //允许重定向
	Verify             bool              `json:"verify" form:"verify" mapstructure:"verify" gorm:"column:verify;"`                                     //验证
	PreparatorySteps   int64             `json:"preparatorySteps" form:"preparatorySteps" gorm:"column:preparatory_steps;"`                            //前置步骤
	PreparatoryStepsID uint              `json:"setup_case_id" form:"setup_case_id" gorm:"comment:前置用例;column:setup_case_id;"`                         //前置步骤ID
	Environs           map[string]string `json:"environs" form:"environs" gorm:"-" swaggertype:"object"`                                               //环境变量
	ReportID           int64             `json:"report_id" form:"report_id" gorm:"-"`                                                                  //报告ID
	Retry              int               `json:"retry" form:"retry" gorm:"comment:重试次数;column:retry;"`                                                 //重试次数
	ProjectId          int64             `json:"projectId" form:"projectId" gorm:"column:project_id;"`

	AndroidDeviceOptionsID *uint                 `json:"androidDeviceOptionsId" form:"androidDeviceOptionsId" gorm:"column:android_device_options_id;"`
	AndroidDeviceOptions   *AndroidDeviceOptions `json:"androidDeviceOptions" form:"androidDeviceOptions" gorm:"foreignKey:AndroidDeviceOptionsID;"`

	IOSDeviceOptionsID *uint             `json:"iosDeviceOptionsId" form:"iosDeviceOptionsId" gorm:"column:ios_device_options_id;"`
	IOSDeviceOptions   *IOSDeviceOptions `json:"iosDeviceOptions" form:"iosDeviceOptions" gorm:"foreignKey:IOSDeviceOptionsID;"`

	HarmonyDeviceOptionsID *uint                 `json:"harmonyDeviceOptionsId" form:"harmonyDeviceOptionsId" gorm:"column:harmony_device_options_id;"`
	HarmonyDeviceOptions   *HarmonyDeviceOptions `json:"harmonyDeviceOptions" form:"harmonyDeviceOptions" gorm:"foreignKey:HarmonyDeviceOptionsID;"`

	BrowserDeviceOptionsID *uint                 `json:"browserDeviceOptionsId" form:"browserDeviceOptionsId" gorm:"column:browser_device_options_id;"`
	BrowserDeviceOptions   *BrowserDeviceOptions `json:"browserDeviceOptions" form:"browserDeviceOptions" gorm:"foreignKey:BrowserDeviceOptionsID;"`
}

// TableName 运行配置 RunConfig自定义表名 run_configs
func (RunConfig) TableName() string {
	return "yf_run_configs"
}
