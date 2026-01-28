package platform

import (
	"github.com/test-instructor/yangfan/server/v2/service"
)

type ApiGroup struct {
	EnvApi
	EnvDetailApi
	PythonCodeApi
	PythonCodeDebugApi
	PythonPackageApi
	PythonCodeFuncApi
	RunConfigApi
	CategoryMenuApi
	RunnerNodeApi
	RunnerApi
	LLMModelConfigApi
}

var (
	envService       = service.ServiceGroupApp.PlatformServiceGroup.EnvService
	edService        = service.ServiceGroupApp.PlatformServiceGroup.EnvDetailService
	pcService        = service.ServiceGroupApp.PlatformServiceGroup.PythonCodeService
	pcdService       = service.ServiceGroupApp.PlatformServiceGroup.PythonCodeDebugService
	ppService        = service.ServiceGroupApp.PlatformServiceGroup.PythonPackageService
	pcfService       = service.ServiceGroupApp.PlatformServiceGroup.PythonCodeFuncService
	rcService        = service.ServiceGroupApp.PlatformServiceGroup.RunConfigService
	cmService        = service.ServiceGroupApp.PlatformServiceGroup.CategoryMenuService
	rnService        = service.ServiceGroupApp.PlatformServiceGroup.RunnerNodeService
	runnerService    = service.ServiceGroupApp.PlatformServiceGroup.RunnerService
	llmconfigService = service.ServiceGroupApp.PlatformServiceGroup.LLMModelConfigService
)
