// 自动生成模板ApiConfig
package interfacecase

import (
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/model/interfacecase/customType"
	"github.com/test-instructor/cheetah/server/model/system"
	"gorm.io/datatypes"
)

// ApiConfig 结构体
// 如果含有time.Time 请自行import time包
type ApiConfig struct {
	global.GVA_MODEL
	Name           string              `json:"name" form:"name" gorm:"column:name;comment:;size:255;"`
	BaseUrl        string              `json:"base_url" form:"base_url" gorm:"column:base_url;comment:;size:255;"`
	Variables      datatypes.JSONMap   `json:"variables" form:"variables" gorm:"column:variables;comment:;type:text"`
	Headers        datatypes.JSONMap   `json:"headers" form:"headers" gorm:"column:headers;comment:;type:text"`
	Parameters     datatypes.JSONMap   `json:"parameters" form:"parameters" gorm:"column:parameters;comment:;type:text"`
	VariablesJson  datatypes.JSON      `json:"variables_json" form:"variables_json"`
	HeadersJson    datatypes.JSON      `json:"headers_json" form:"headers_json" gorm:"column:headers_json;comment:;"`
	ParametersJson datatypes.JSON      `json:"parameters_json" form:"parameters_json" gorm:"column:parameters_json;comment:;"`
	Weight         int                 `json:"weight" form:"weight" gorm:"column:weight;weight:;"`
	Default        *bool               `json:"default" form:"default" gorm:"column:default;comment:;"`
	Timeout        float32             `json:"timeout,omitempty" form:"timeout" `
	AllowRedirects bool                `json:"allow_redirects,omitempty" form:"allow_redirects" gorm:"column:allow_redirects;comment:;"`
	Verify         bool                `json:"verify" form:"verify" gorm:"column:verify;comment:;"`
	Export         customType.TypeArgs `json:"export,omitempty" gorm:"column:export;comment:;"`
	ProjectID      uint                `json:"-"`
	Project        system.Project      `json:"-"`
}
