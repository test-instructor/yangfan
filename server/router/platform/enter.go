package platform

import api "github.com/test-instructor/yangfan/server/v2/api/v1"

type RouterGroup struct {
	EnvRouter
	EnvDetailRouter
	PythonCodeRouter
	PythonCodeDebugRouter
	PythonPackageRouter
	PythonCodeFuncRouter
	RunConfigRouter
	CategoryMenuRouter
	RunnerNodeRouter
	RunnerRouter
	OpenRunnerRouter
}

var (
	envApi    = api.ApiGroupApp.PlatformApiGroup.EnvApi
	edApi     = api.ApiGroupApp.PlatformApiGroup.EnvDetailApi
	pcApi     = api.ApiGroupApp.PlatformApiGroup.PythonCodeApi
	pcdApi    = api.ApiGroupApp.PlatformApiGroup.PythonCodeDebugApi
	ppApi     = api.ApiGroupApp.PlatformApiGroup.PythonPackageApi
	pcfApi    = api.ApiGroupApp.PlatformApiGroup.PythonCodeFuncApi
	rcApi     = api.ApiGroupApp.PlatformApiGroup.RunConfigApi
	cmApi     = api.ApiGroupApp.PlatformApiGroup.CategoryMenuApi
	rnApi     = api.ApiGroupApp.PlatformApiGroup.RunnerNodeApi
	runnerApi = api.ApiGroupApp.PlatformApiGroup.RunnerApi
)
