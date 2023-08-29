// 自动生成模板InterfaceTemplate
package interfacecase

import (
	"gorm.io/datatypes"

	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/interfacecase/customType"
)

type ApiType int

const (
	ApiTypeTemplate        ApiType = 1 // 接口模板
	ApiTypeCase            ApiType = 2 // 接口测试用例
	ApiTypeCasePerformance ApiType = 3 // 性能测试用例
)

var _ = []ApiType{ApiTypeTemplate, ApiTypeCase}

type TransactionType string

const (
	TransactionStart TransactionType = "start"
	TransactionEnd   TransactionType = "end"
)

// ApiStep 结构体
// 如果含有time.Time 请自行import time包
type ApiStep struct {
	global.GVA_MODEL
	Operator
	Name    string      `json:"name,omitempty" form:"name" gorm:"column:name;comment:接口名称"`
	ApiType ApiType     `json:"type,omitempty" form:"type" gorm:"column:api_type;comment:接口类型"`
	Request *ApiRequest `json:"request,omitempty" form:"request;comment:http接口"`
	Grpc    *ApiGrpc    `json:"gRPC,omitempty" yaml:"gRPC,omitempty"`

	Transaction *ApiStepTransaction `json:"transaction,omitempty" yaml:"transaction,omitempty;comment:事务"`
	Rendezvous  *ApiStepRendezvous  `json:"rendezvous,omitempty" yaml:"rendezvous,omitempty;comment:集合点"`
	ThinkTime   *ApiStepThinkTime   `json:"think_time,omitempty" yaml:"think_time,omitempty;comment:思考时间"`

	ThinkTimeID   uint `json:"think_time_id,omitempty" gorm:"comment:思考时间"`
	TransactionID uint `json:"transaction_id,omitempty" gorm:"comment:事务"`
	RendezvousID  uint `json:"rendezvous_id,omitempty" gorm:"comment:集合点"`
	RequestID     uint `json:"request_id,omitempty" gorm:"comment:http请求"`
	GrpcID        uint `json:"grpc_id,omitempty" gorm:"comment:grpc请求"`

	Variables       datatypes.JSONMap      `json:"variables,omitempty" form:"variables" gorm:"column:variables;comment:变量;type:text"`
	Extract         datatypes.JSONMap      `json:"extract,omitempty" form:"extract" gorm:"column:extract;comment:导出参数;type:text"`
	Validate        customType.TypeArgsMap `json:"validate,omitempty" form:"validate" gorm:"column:validate;comment:断言;type:text"`
	ValidateNumber  uint                   `json:"validate_number,omitempty" form:"validate_number;comment:断言数量"`
	ValidateJson    datatypes.JSON         `json:"validate_json,omitempty" form:"validate_json;comment:变量json格式" `
	ExtractJson     datatypes.JSON         `json:"extract_json,omitempty" form:"extract_json;comment:导出参数json格式"`
	VariablesJson   datatypes.JSON         `json:"variables_json,omitempty" form:"variables_json;comment:断言json类型"`
	Hooks           string                 `json:"hooks,omitempty" form:"hooks" gorm:"column:hooks;"`
	SetupHooks      customType.TypeArgs    `json:"setup_hooks,omitempty" form:"setup_hooks,omitempty" gorm:"column:setup_hooks;type:text"`
	TeardownHooks   customType.TypeArgs    `json:"teardown_hooks,omitempty" form:"teardown_hooks,omitempty" gorm:"column:teardown_hooks;type:text"`
	TTestCase       []ApiCaseStep          `json:"testCase,omitempty" form:"testCase" gorm:"many2many:ApiCaseStepRelationship;"`
	Sort            uint                   `json:"sort,omitempty" form:"sort" gorm:"column:sort;"`
	ExportHeader    datatypes.JSON         `json:"export_header,omitempty" gorm:"column:export_header;comment:导出请求头到全局config;type:text"`
	ExportParameter datatypes.JSON         `json:"export_parameter,omitempty" gorm:"column:export_parameter;comment:导出参数到全局config;type:text"`
	Retry           uint                   `json:"retry" gorm:"comment:重试次数"`

	Parent    uint    `json:"-"`
	ApiMenuID uint    `json:"-"`
	ApiMenu   ApiMenu `json:"-"`
}

type ApiStepTransaction struct {
	global.GVA_MODEL
	Name string          `json:"name,omitempty" yaml:"name"`
	Type TransactionType `json:"type,omitempty" yaml:"type"`
}

type ApiStepRendezvous struct {
	global.GVA_MODEL
	Name    string  `json:"name,omitempty" yaml:"name"`                 // required
	Percent float32 `json:"percent,omitempty" yaml:"percent,omitempty"` // default to 1(100%)
	Number  int64   `json:"number,omitempty" yaml:"number,omitempty"`
	Timeout int64   `json:"timeout,omitempty" yaml:"timeout,omitempty"` // milliseconds
}

type ApiStepThinkTime struct {
	global.GVA_MODEL
	Time float64 `json:"time,omitempty" yaml:"time"`
}
