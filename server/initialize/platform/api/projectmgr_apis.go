package api

import (
	sysModel "github.com/test-instructor/yangfan/server/v2/model/system"
)

func projectmgrApis() []sysModel.SysApi {
	return []sysModel.SysApi{
		{Path: "/upa/createUserProjectAccess", Description: "新增项目成员与权限", ApiGroup: "项目成员与权限", Method: "POST"},
		{Path: "/upa/deleteUserProjectAccess", Description: "删除项目成员与权限", ApiGroup: "项目成员与权限", Method: "DELETE"},
		{Path: "/upa/deleteUserProjectAccessByIds", Description: "批量删除项目成员与权限", ApiGroup: "项目成员与权限", Method: "DELETE"},
		{Path: "/upa/updateUserProjectAccess", Description: "更新项目成员与权限", ApiGroup: "项目成员与权限", Method: "PUT"},
		{Path: "/upa/findUserProjectAccess", Description: "根据ID获取项目成员与权限", ApiGroup: "项目成员与权限", Method: "GET"},
		{Path: "/upa/getUserProjectAccessList", Description: "获取项目成员与权限列表", ApiGroup: "项目成员与权限", Method: "GET"},

		{Path: "/pj/createProject", Description: "新增项目配置", ApiGroup: "项目配置", Method: "POST"},
		{Path: "/pj/deleteProject", Description: "删除项目配置", ApiGroup: "项目配置", Method: "DELETE"},
		{Path: "/pj/deleteProjectByIds", Description: "批量删除项目配置", ApiGroup: "项目配置", Method: "DELETE"},
		{Path: "/pj/updateProject", Description: "更新项目配置", ApiGroup: "项目配置", Method: "PUT"},
		{Path: "/pj/resetProjectAuth", Description: "重设项目CI鉴权信息", ApiGroup: "项目配置", Method: "PUT"},
		{Path: "/pj/findProject", Description: "根据ID获取项目配置", ApiGroup: "项目配置", Method: "GET"},
		{Path: "/pj/getProjectList", Description: "获取项目配置列表", ApiGroup: "项目配置", Method: "GET"},
	}
}
