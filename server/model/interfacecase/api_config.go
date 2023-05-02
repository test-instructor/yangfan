// 自动生成模板ApiConfig
package interfacecase

import (
	"gorm.io/datatypes"

	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/interfacecase/customType"
)

// ApiConfig 结构体
// 如果含有time.Time 请自行import time包
type ApiConfig struct {
	global.GVA_MODEL
	Operator
	Name           string              `json:"name" form:"name" gorm:"column:name;comment:配置名称;size:255;"`
	BaseUrl        string              `json:"base_url" form:"base_url" gorm:"column:base_url;comment:默认域名;size:255;"`
	Variables      datatypes.JSONMap   `json:"variables" form:"variables" gorm:"column:variables;comment:变量;type:text"`
	Headers        datatypes.JSONMap   `json:"headers" form:"headers" gorm:"column:headers;comment:请求头;type:text"`
	Parameters     datatypes.JSONMap   `json:"parameters" form:"parameters" gorm:"column:parameters;comment:参数;type:text"`
	VariablesJson  datatypes.JSON      `json:"variables_json" form:"variables_json"`
	HeadersJson    datatypes.JSON      `json:"headers_json" form:"headers_json" gorm:"column:headers_json;comment:;type:text"`
	Weight         int                 `json:"weight" form:"weight" gorm:"column:weight;weight:;"`
	Default        bool                `json:"default" form:"default" gorm:"column:default;comment:默认配置;"`
	Timeout        float32             `json:"timeout,omitempty" form:"timeout" gorm:"comment:超时时间" `
	AllowRedirects bool                `json:"allow_redirects,omitempty" form:"allow_redirects" gorm:"column:allow_redirects;comment:;"`
	Verify         bool                `json:"verify" form:"verify" gorm:"column:verify;comment:;"`
	Export         customType.TypeArgs `json:"export,omitempty" gorm:"column:export;comment:导出参数;"`
	SetupCase      *ApiCaseStep        `json:"setup_case" `
	SetupCaseID    *uint               `json:"setup_case_id" form:"setup_case_id" gorm:"comment:前置用例"`
	Environs       map[string]string   `json:"environs,omitempty" yaml:"environs,omitempty" gorm:"-"` // environment variables
	CaseID         uint                `json:"case_id" gorm:"-"`
	ReportID       uint                `json:"report_id" gorm:"-"`
}
