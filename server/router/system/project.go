package system

import (
	"github.com/gin-gonic/gin"

	"github.com/test-instructor/yangfan/server/api/v1"
	"github.com/test-instructor/yangfan/server/middleware"
)

type ProjectRouter struct {
}

// InitProjectRouter 初始化 Project 路由信息
func (s *ProjectRouter) InitProjectRouter(Router *gin.RouterGroup) {
	projectRouter := Router.Group("project").Use(middleware.OperationRecord())
	projectRouterWithoutRecord := Router.Group("project")
	projectApi := v1.ApiGroupApp.SystemApiGroup.ProjectApi
	{
		projectRouter.POST("createProject", projectApi.CreateProject)               // 新建Project
		projectRouter.DELETE("deleteProject", projectApi.DeleteProject)             // 删除Project
		projectRouter.DELETE("deleteProjectByIds", projectApi.DeleteProjectByIds)   // 批量删除Project
		projectRouter.PUT("updateProject", projectApi.UpdateProject)                // 更新Project
		projectRouter.POST("setUserProjectAuth", projectApi.SetUserProjectAuth)     // 设置用户项目权限
		projectRouter.DELETE("deleteUserProjectAuth", projectApi.DeleteProjectAuth) // 删除用户项目权限
	}
	{
		projectRouterWithoutRecord.GET("findProject", projectApi.FindProject)               // 根据ID获取Project
		projectRouterWithoutRecord.GET("getProjectList", projectApi.GetProjectList)         // 获取Project列表
		projectRouterWithoutRecord.GET("getProjectUserList", projectApi.GetProjectUserList) // 获取项目用户列表
	}
}
