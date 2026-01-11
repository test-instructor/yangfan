package projectmgr

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/middleware"
)

type ProjectRouter struct{}

// InitProjectRouter 初始化 项目配置 路由信息
func (s *ProjectRouter) InitProjectRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	pjRouter := Router.Group("pj").Use(middleware.OperationRecord())
	pjRouterWithoutRecord := Router.Group("pj")
	pjRouterWithoutAuth := PublicRouter.Group("pj")
	{
		pjRouter.POST("createProject", pjApi.CreateProject)             // 新建项目配置
		pjRouter.DELETE("deleteProject", pjApi.DeleteProject)           // 删除项目配置
		pjRouter.DELETE("deleteProjectByIds", pjApi.DeleteProjectByIds) // 批量删除项目配置
		pjRouter.PUT("updateProject", pjApi.UpdateProject)              // 更新项目配置
	}
	{
		pjRouterWithoutRecord.GET("findProject", pjApi.FindProject)       // 根据ID获取项目配置
		pjRouterWithoutRecord.GET("getProjectList", pjApi.GetProjectList) // 获取项目配置列表
	}
	{
		pjRouterWithoutAuth.GET("getProjectPublic", pjApi.GetProjectPublic) // 项目配置开放接口
	}
}
