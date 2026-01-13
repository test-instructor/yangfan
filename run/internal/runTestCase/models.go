package runTestCase

import (
	"github.com/test-instructor/yangfan/httprunner/hrp"
	"github.com/test-instructor/yangfan/server/v2/model/automation"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
)

// Structures adapted from yangfan/server/model/interfacecase

type ApisCaseModel struct {
	Case      []hrp.ITestCase
	Config    *hrp.TConfig
	SetupCase bool
	Environs  map[string]string
}

type CaseList struct {
	Case      []hrp.ITestCase
	SetupCase bool
}

type HrpCaseStep struct {
	ID          uint
	Name        string
	TestCase    interface{}  `json:"testcase,omitempty" yaml:"testcase,omitempty"`
	Transaction *Transaction `json:"transaction,omitempty" yaml:"transaction,omitempty;comment:事务"`
	Rendezvous  *Rendezvous  `json:"rendezvous,omitempty" yaml:"rendezvous,omitempty;comment:集合点"`
	ThinkTime   *ThinkTime   `json:"think_time,omitempty" yaml:"think_time,omitempty;comment:思考时间"`
	Len         int
}

type HrpTestCase struct {
	ID        uint
	Name      string
	Confing   platform.RunConfig     `json:"config,omitempty" form:"config"`
	TestSteps []*automation.AutoStep `json:"teststeps,omitempty" yaml:"teststeps,omitempty"`
}

type HrpCase struct {
	ID        uint
	Name      string
	Confing   platform.RunConfig `json:"config,omitempty" form:"config"`
	TestSteps []interface{}      `json:"teststeps,omitempty" yaml:"teststeps,omitempty"`
}

// YangfanTestCase 实现 hrp.ITestCase 接口
// 用于将业务数据直接转换为 httprunner 可执行格式
type YangfanTestCase struct {
	ID        uint
	Name      string
	Config    *hrp.TConfig
	TestSteps []hrp.IStep
}

// GetTestCase 实现 hrp.ITestCase 接口
func (ltc *YangfanTestCase) GetTestCase() (*hrp.TestCase, error) {
	if ltc.Name != "" && ltc.Config != nil {
		ltc.Config.Name = ltc.Name
		if ltc.Config.Variables == nil {
			ltc.Config.Variables = make(map[string]interface{})
		}
		ltc.Config.Variables["__case_name__"] = ltc.Name
	}
	return &hrp.TestCase{
		Config:    ltc.Config,
		TestSteps: ltc.TestSteps,
	}, nil
}

type Transaction struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Rendezvous struct {
	Name    string  `json:"name"`
	Percent float64 `json:"percent"`
	Timeout int64   `json:"timeout"`
}

type ThinkTime struct {
	Time float64 `json:"time"`
}
