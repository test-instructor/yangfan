package automation

import api "github.com/test-instructor/yangfan/server/v2/api/v1"

type RouterGroup struct {
	AutoStepRouter
	RequestRouter
	AutoCaseStepRouter
	AutoCaseRouter
	TimerTaskRouter
	AutoReportRouter
}

var (
	asApi  = api.ApiGroupApp.AutomationApiGroup.AutoStepApi
	reqApi = api.ApiGroupApp.AutomationApiGroup.RequestApi
	acsApi = api.ApiGroupApp.AutomationApiGroup.AutoCaseStepApi
	acApi  = api.ApiGroupApp.AutomationApiGroup.AutoCaseApi
	tkApi  = api.ApiGroupApp.AutomationApiGroup.TimerTaskApi
	arApi  = api.ApiGroupApp.AutomationApiGroup.AutoReportApi
)
