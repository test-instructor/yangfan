// Package automation 自动生成模板Request
package automation

import (
	"github.com/test-instructor/yangfan/server/v2/global"
	"gorm.io/datatypes"
)

// Request 请求 结构体  Request
type Request struct {
	global.GVA_MODEL

	Method         string            `json:"method" form:"method" mapstructure:"method" gorm:"column:method;"` //方法
	URL            string            `json:"url" form:"url" mapstructure:"url" gorm:"column:url;"`             //URL
	HTTP2          bool              `json:"http2" form:"http2" mapstructure:"http2" gorm:"column:http2;"`     //HTTP2
	Params         datatypes.JSONMap `json:"params" form:"params" mapstructure:"params" gorm:"column:params;"` //参数
	ParamTemp      datatypes.JSON    `json:"param_temp" form:"param_temp" gorm:"column:param_temp;"`
	Headers        datatypes.JSONMap `json:"headers" form:"headers" mapstructure:"headers" gorm:"column:headers;"` //头
	HeaderTemp     datatypes.JSON    `json:"header_temp" form:"header_temp" gorm:"column:header_temp;"`
	Json           datatypes.JSONMap `json:"json" form:"json" mapstructure:"body,omitempty" gorm:"column:json;"` //Json
	DataWarehouse  datatypes.JSONMap `json:"data_warehouse" form:"data_warehouse" mapstructure:"data_warehouse" gorm:"column:data_warehouse;"`
	Data           datatypes.JSONMap `json:"data" form:"data" mapstructure:"data,omitempty" gorm:"column:data;"` //数据
	DataTemp       datatypes.JSON    `json:"data_temp" form:"data_temp" gorm:"column:data_temp;"`
	Timeout        float64           `json:"timeout" form:"timeout" mapstructure:"timeout" gorm:"column:timeout;"`                                 //超时
	AllowRedirects bool              `json:"allow_redirects" form:"allow_redirects" mapstructure:"allow_redirects" gorm:"column:allow_redirects;"` //允许重定向
	Verify         bool              `json:"verify" form:"verify" mapstructure:"verify" gorm:"column:verify;"`                                     //验证
	Upload         datatypes.JSONMap `json:"upload" form:"upload" mapstructure:"upload" gorm:"column:upload;"`                                     //上传
	ProjectId      int64             `json:"projectId" form:"projectId" gorm:"column:project_id;"`                                                 //项目信息

	Request   *Request `json:"request,omitempty" form:"request;comment:http接口"`
	RequestID uint     `json:"request_id,omitempty" gorm:"comment:http请求"`
	MenuID    uint     `json:"menuId" form:"menuId" gorm:"column:menu_id;"`

	//Cookies        map[string]string      `json:"cookies" form:"cookies" gorm:"column:cookies;"`                         //Cookies
	//Body           interface{}            `json:"body" form:"body" gorm:"column:body;"`                                  //Body
}

// TableName 请求 Request自定义表名 request
func (Request) TableName() string {
	return "yf_request"
}
