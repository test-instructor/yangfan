package hrp

import (
	"github.com/test-instructor/yangfan/server/model/interfacecase"
)

type ApisCaseModel struct {
	Case      []ITestCase
	Config    interfacecase.ApiConfig
	SetupCase bool
	Environs  map[string]string
}

type CaseList struct {
	Case      []ITestCase
	SetupCase bool
}
