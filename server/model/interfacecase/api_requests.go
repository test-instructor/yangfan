package interfacecase

import (
	"github.com/test-instructor/cheetah/server/global"
	"gorm.io/datatypes"
)

type ApiRequest struct {
	global.GVA_MODEL
	Agreement      string            `json:"agreement" form:"agreement" gorm:"column:agreement;comment:协议"`
	Method         string            `json:"method" form:"method" gorm:"column:method;comment:请求方法"`
	HTTP2          bool              `json:"http2,omitempty" form:"http2"  gorm:"column:http2;comment:;"`
	Url            string            `json:"url" form:"url" gorm:"column:url;comment:;"`
	Params         datatypes.JSONMap `json:"params" form:"params" gorm:"column:params;comment:;type:text"`
	Headers        datatypes.JSONMap `json:"headers" form:"headers" gorm:"column:headers;comment:;type:text"`
	Data           datatypes.JSONMap `json:"data" form:"data" gorm:"column:data;comment:;type:text"`
	ParamsJson     datatypes.JSON    `json:"params_json" form:"params_json" gorm:"column:params_json;comment:;"`
	HeadersJson    datatypes.JSON    `json:"headers_json" form:"headers_json" gorm:"column:headers_json;comment:;"`
	DataJson       datatypes.JSON    `json:"data_json" form:"data_json" gorm:"column:data_json;comment:;"`
	Json           datatypes.JSON    `json:"json" form:"json"`
	Timeout        float32           `json:"timeout,omitempty" form:"timeout" `
	AllowRedirects bool              `json:"allow_redirects,omitempty" form:"allow_redirects" gorm:"column:allow_redirects;comment:;"`
	Verify         bool              `json:"verify" form:"verify" gorm:"column:verify;comment:;"`
}
