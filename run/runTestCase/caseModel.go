package runTestCase

import (
	"github.com/test-instructor/yangfan/hrp"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
)

type ApisCaseModel struct {
	Case      []hrp.ITestCase
	Config    interfacecase.ApiConfig
	SetupCase bool
	Environs  map[string]string
}

type CaseList struct {
	Case      []hrp.ITestCase
	SetupCase bool
}
