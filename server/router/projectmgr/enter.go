package projectmgr

import api "github.com/test-instructor/yangfan/server/v2/api/v1"

type RouterGroup struct {
	UserProjectAccessRouter
	ProjectRouter
	ReportNotifyRouter
}

var (
	upaApi = api.ApiGroupApp.ProjectmgrApiGroup.UserProjectAccessApi
	pjApi  = api.ApiGroupApp.ProjectmgrApiGroup.ProjectApi
	ntApi  = api.ApiGroupApp.ProjectmgrApiGroup.ReportNotifyApi
)
