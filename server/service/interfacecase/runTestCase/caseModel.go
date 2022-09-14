package runTestCase

import (
	"github.com/test-instructor/cheetah/server/hrp"
	"github.com/test-instructor/cheetah/server/model/interfacecase"
)

type ApisCaseModel struct {
	Case      []hrp.ITestCase
	Config    interfacecase.ApiConfig
	SetupCase bool
}

type CaseList struct {
	Case      []hrp.ITestCase
	SetupCase bool
}
