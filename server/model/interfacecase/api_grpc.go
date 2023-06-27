package interfacecase

import (
	"github.com/test-instructor/yangfan/server/global"
	"gorm.io/datatypes"
)

type GrpcType string

var (
	GrpcTypeSimple              GrpcType = "Simple"
	GrpcTypeServerSideStream    GrpcType = "ServerSideStream"
	GrpcTypeClientSideStream    GrpcType = "ClientSideStream"
	GrpcTypeBidirectionalStream GrpcType = "BidirectionalStream"
)

type ApiGrpc struct {
	global.GVA_MODEL
	URL         string            `json:"url" yaml:"url" gorm:"column:url;comment:请求地址"`
	Headers     datatypes.JSONMap `json:"headers,omitempty" form:"headers" gorm:"column:headers;comment:请求头;type:text"`
	HeadersJson datatypes.JSON    `json:"headers_json,omitempty" form:"headers_json" gorm:"column:headers_json;comment:请求头json数据格式;"`
	Body        datatypes.JSONMap `json:"body,omitempty" gorm:"column:body;comment:请求体"`
	Timeout     float32           `json:"timeout,omitempty" yaml:"timeout,omitempty" gorm:"column:timeout;comment:超时时间"`
	Type        GrpcType          `json:"type,omitempty" yaml:"type,omitempty" gorm:"column:type;comment:请求类型"`
	Detail      datatypes.JSONMap `json:"detail,omitempty" form:"detail" gorm:"column:detail;comment:请求服务器详情;type:text"`

	Parent uint `json:"-"`
}
