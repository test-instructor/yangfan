package projectmgr

import "github.com/test-instructor/yangfan/server/v2/service"

type ApiGroup struct {
	UserProjectAccessApi
	ProjectApi
	ReportNotifyApi
}

var (
	upaService = service.ServiceGroupApp.ProjectmgrServiceGroup.UserProjectAccessService
	pjService  = service.ServiceGroupApp.ProjectmgrServiceGroup.ProjectService
	ntService  = service.ServiceGroupApp.ProjectmgrServiceGroup.ReportNotifyService
)
