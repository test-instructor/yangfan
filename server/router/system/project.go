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
		projectRouter.POST("createProject", projectApi.CreateProject)             // 新建Project
		projectRouter.DELETE("deleteProject", projectApi.DeleteProject)           // 删除Project
		projectRouter.DELETE("deleteProjectByIds", projectApi.DeleteProjectByIds) // 批量删除Project
		projectRouter.PUT("updateProject", projectApi.UpdateProject)              // 更新Project
	}
	{
		projectRouterWithoutRecord.GET("findProject", projectApi.FindProject)       // 根据ID获取Project
		projectRouterWithoutRecord.GET("getProjectList", projectApi.GetProjectList) // 获取Project列表
	}
}
