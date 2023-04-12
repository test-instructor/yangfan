package interfacecase

import (
	"github.com/test-instructor/yangfan/server/global"
	"gorm.io/datatypes"
)

type ApiRequest struct {
	global.GVA_MODEL
	Agreement      string            `json:"agreement" form:"agreement" gorm:"column:agreement;comment:协议"`
	Method         string            `json:"method" form:"method" gorm:"column:method;comment:请求方法"`
	HTTP2          bool              `json:"http2,omitempty" form:"http2"  gorm:"column:http2;comment:是否为http2;"`
	Url            string            `json:"url" form:"url" gorm:"column:url;comment:请求url;"`
	Params         datatypes.JSONMap `json:"params" form:"params" gorm:"column:params;comment:url参数;type:text"`
	Headers        datatypes.JSONMap `json:"headers" form:"headers" gorm:"column:headers;comment:请求头;type:text"`
	Data           datatypes.JSONMap `json:"data" form:"data" gorm:"column:data;comment:request body data;type:text"`
	ParamsJson     datatypes.JSON    `json:"params_json" form:"params_json" gorm:"column:params_json;comment:url参数json数据格式;"`
	HeadersJson    datatypes.JSON    `json:"headers_json" form:"headers_json" gorm:"column:headers_json;comment:请求头json数据格式;"`
	DataJson       datatypes.JSON    `json:"data_json" form:"data_json" gorm:"column:data_json;comment:request body data json数据格式;"`
	Json           datatypes.JSON    `json:"json" form:"json"`
	Timeout        float32           `json:"timeout,omitempty" form:"timeout"  gorm:"comment:超时时间"`
	AllowRedirects bool              `json:"allow_redirects,omitempty" form:"allow_redirects" gorm:"column:allow_redirects;comment:;"`
	Verify         bool              `json:"verify" form:"verify" gorm:"column:verify;comment:;"`
}
