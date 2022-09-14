// 自动生成模板InterfaceTemplate
package interfacecase

import (
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/model/interfacecase/customType"
	"github.com/test-instructor/cheetah/server/model/system"
	"gorm.io/datatypes"
)

type ApiType int

const (
	ApiTypeTemplate ApiType = 1 // 接口模板
	ApiTypeCase     ApiType = 2 // 测试用例
)

// ApiStep 结构体
// 如果含有time.Time 请自行import time包
type ApiStep struct {
	global.GVA_MODEL
	Name            string                 `json:"name" form:"name" gorm:"column:name;comment:接口名称"`
	ApiType         ApiType                `json:"type" form:"type" gorm:"column:api_type;comment:接口名称"`
	RequestID       uint                   `json:"-"`
	Request         *ApiRequest            `json:"request" form:"request"`
	Variables       datatypes.JSONMap      `json:"variables" form:"variables" gorm:"column:variables;comment:;type:text"`
	Extract         datatypes.JSONMap      `json:"extract" form:"extract" gorm:"column:extract;comment:;type:text"`
	Validate        customType.TypeArgsMap `json:"validate" form:"validate" gorm:"column:validate;comment:;type:text"`
	ValidateNumber  uint                   `json:"validate_number" form:"validate_number"`
	ValidateJson    datatypes.JSON         `json:"validate_json" form:"validate_json" `
	ExtractJson     datatypes.JSON         `json:"extract_json" form:"extract_json"`
	VariablesJson   datatypes.JSON         `json:"variables_json" form:"variables_json"`
	Hooks           string                 `json:"hooks" form:"hooks" gorm:"column:hooks;"`
	SetupHooks      customType.TypeArgs    `json:"setup_hooks,omitempty" form:"setup_hooks,omitempty" gorm:"column:setup_hooks;type:text"`
	TeardownHooks   customType.TypeArgs    `json:"teardown_hooks,omitempty" form:"teardown_hooks,omitempty" gorm:"column:teardown_hooks;type:text"`
	ProjectID       uint                   `json:"-"`
	TTestCase       []ApiCaseStep          `json:"testCase" form:"testCase" gorm:"many2many:ApiCaseStepRelationship;"`
	Sort            uint                   `json:"sort" form:"sort" gorm:"column:sort;"`
	ExportHeader    datatypes.JSON         `json:"export_header" gorm:"column:export_header;comment:;type:text"`
	ExportParameter datatypes.JSON         `json:"export_parameter" gorm:"column:export_parameter;comment:;type:text"`
	Parent          uint                   `json:"-"`
	Project         system.Project         `json:"-"`
	ApiMenuID       uint                   `json:"-"`
	ApiMenu         ApiMenu                `json:"-"`
	CreatedBy       system.SysUser         `json:"-"`
	CreatedByID     uint                   `json:"-"`
	UpdateBy        system.SysUser         `json:"-"`
	UpdateByID      uint                   `json:"-"`
	DeleteBy        system.SysUser         `json:"-"`
	DeleteByID      uint                   `json:"-"`
	TestCase        interface{}            `gorm:"-"`
}
