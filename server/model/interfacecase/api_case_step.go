// 自动生成模板TestCase
package interfacecase

import (
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/model/system"
)

// ApiCaseStep 结构体
// 如果含有time.Time 请自行import time包
type ApiCaseStep struct {
	global.GVA_MODEL
	Name      string    `json:"name" form:"name" gorm:"column:name;comment:;"`
	FrontCase *bool     `json:"front_case" orm:"front_case"`
	TStep     []ApiStep `json:"TStep" form:"TStep" gorm:"many2many:ApiCaseStepRelationship;"`
	ApiCase   []ApiCase `json:"case" form:"case" gorm:"many2many:ApiCaseRelationship;"`
	//RunConfig   *ApiConfig     `json:"runConfig" form:"runConfig"`
	RunConfigID uint           `json:"RunConfigID" form:"RunConfigID"`
	ProjectID   uint           `json:"-"`
	Project     system.Project `json:"-"`
	ApiMenu     ApiMenu        `json:"-"`
	ApiMenuID   uint           `json:"-"`
	CreatedBy   system.SysUser `json:"-"`
	CreatedByID uint           `json:"-"`
	UpdateBy    system.SysUser `json:"-"`
	UpdateByID  uint           `json:"-"`
	DeleteBy    system.SysUser `json:"-"`
	DeleteByID  uint           `json:"-"`
}

type HrpCaseStep struct {
	ID       uint
	Name     string
	TestCase HrpTestCase `json:"testcase,omitempty" yaml:"testcase,omitempty"`
}

type HrpTestCase struct {
	ID        uint
	Name      string
	Confing   ApiConfig `json:"config" form:"config"`
	TestSteps []ApiStep `json:"teststeps,omitempty" yaml:"teststeps,omitempty"`
}

type HrpCase struct {
	ID        uint
	Name      string
	Confing   ApiConfig     `json:"config" form:"config"`
	TestSteps []HrpCaseStep `json:"teststeps,omitempty" yaml:"teststeps,omitempty"`
}
