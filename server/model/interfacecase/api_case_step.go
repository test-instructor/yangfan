// 自动生成模板TestCase
package interfacecase

import (
	"github.com/test-instructor/yangfan/server/global"
)

type ApiStepType string

const (
	ApiStepTypeTransaction      ApiStepType = "Transaction"
	ApiStepTypeTransactionStart ApiStepType = "TransactionStart"
	ApiStepTypeTransactionEnd   ApiStepType = "TransactionEnd"
	ApiStepTypeRendezvous       ApiStepType = "Rendezvous"
)

var _ = ApiStepTypeTransaction

// ApiCaseStep 结构体
// 如果含有time.Time 请自行import time包
type ApiCaseStep struct {
	global.GVA_MODEL
	Operator
	Name      string    `json:"name" form:"name" gorm:"column:name;comment:步骤名称;"`
	FrontCase *bool     `json:"front_case" form:"front_case" gorm:"comment:允许设置为前置用例;"`
	TStep     []ApiStep `json:"TStep" form:"TStep" gorm:"many2many:ApiCaseStepRelationship;"`
	//Performance []Performance `json:"performance" form:"performance" gorm:"many2many:PerformanceRelationship;"`
	ApiCase []ApiCase `json:"case" form:"case" gorm:"many2many:ApiCaseRelationship;"`
	//RunConfig   *ApiConfig     `json:"runConfig" form:"runConfig"`
	RunConfigID   uint        `json:"RunConfigID" form:"RunConfigID" gorm:"comment:运行配置;"`
	RunConfigName *string     `json:"RunConfigName" form:"RunConfigName" gorm:"comment:运行配置名称;"`
	ApiMenu       ApiMenu     `json:"-"`
	ApiMenuID     uint        `json:"-" gorm:"comment:所属菜单;"`
	Type          ApiType     `json:"type" form:"type" gorm:"column:type;comment:接口类型"`
	ApiStepType   ApiStepType `json:"api_step_type" gorm:"column:api_step_type;comment:性能测试step类型"`
	ApiEnvName    *string     `json:"api_env_name" gorm:"comment:所属环境名称;"`
	ApiEnvID      uint        `json:"api_env_id" gorm:"comment:所属环境;"`
}

type HrpCaseStep struct {
	ID          uint
	Name        string
	TestCase    interface{}         `json:"testcase,omitempty" yaml:"testcase,omitempty"`
	Transaction *ApiStepTransaction `json:"transaction,omitempty" yaml:"transaction,omitempty;comment:事务"`
	Rendezvous  *ApiStepRendezvous  `json:"rendezvous,omitempty" yaml:"rendezvous,omitempty;comment:集合点"`
	ThinkTime   *ApiStepThinkTime   `json:"think_time,omitempty" yaml:"think_time,omitempty;comment:思考时间"`
	Len         int
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
	TestSteps []interface{} `json:"teststeps,omitempty" yaml:"teststeps,omitempty"`
}
