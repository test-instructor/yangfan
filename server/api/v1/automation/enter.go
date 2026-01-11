package automation

import "github.com/test-instructor/yangfan/server/v2/service"

type ApiGroup struct {
	AutoStepApi
	RequestApi
	AutoCaseStepApi
	AutoCaseApi
	TimerTaskApi
	AutoReportApi
}

var (
	asService  = service.ServiceGroupApp.AutomationServiceGroup.AutoStepService
	reqService = service.ServiceGroupApp.AutomationServiceGroup.RequestService
	acsService = service.ServiceGroupApp.AutomationServiceGroup.AutoCaseStepService
	acService  = service.ServiceGroupApp.AutomationServiceGroup.AutoCaseService
	tkService  = service.ServiceGroupApp.AutomationServiceGroup.TimerTaskService
	arService  = service.ServiceGroupApp.AutomationServiceGroup.AutoReportService
)
