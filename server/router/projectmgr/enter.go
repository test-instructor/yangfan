package projectmgr

import api "github.com/test-instructor/yangfan/server/v2/api/v1"

type RouterGroup struct {
	UserProjectAccessRouter
	ProjectRouter
}

var (
	upaApi = api.ApiGroupApp.ProjectmgrApiGroup.UserProjectAccessApi
	pjApi  = api.ApiGroupApp.ProjectmgrApiGroup.ProjectApi
)
