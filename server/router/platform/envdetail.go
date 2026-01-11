package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/middleware"
)

type EnvDetailRouter struct{}

// InitEnvDetailRouter 初始化 环境详情 路由信息
func (s *EnvDetailRouter) InitEnvDetailRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	edRouter := Router.Group("ed").Use(middleware.OperationRecord())
	edRouterWithoutRecord := Router.Group("ed")
	edRouterWithoutAuth := PublicRouter.Group("ed")
	{
		edRouter.POST("createEnvDetail", edApi.CreateEnvDetail)             // 新建环境详情
		edRouter.DELETE("deleteEnvDetail", edApi.DeleteEnvDetail)           // 删除环境详情
		edRouter.DELETE("deleteEnvDetailByIds", edApi.DeleteEnvDetailByIds) // 批量删除环境详情
		edRouter.PUT("updateEnvDetail", edApi.UpdateEnvDetail)              // 更新环境详情
	}
	{
		edRouterWithoutRecord.GET("findEnvDetail", edApi.FindEnvDetail)       // 根据ID获取环境详情
		edRouterWithoutRecord.GET("getEnvDetailList", edApi.GetEnvDetailList) // 获取环境详情列表
	}
	{
		edRouterWithoutAuth.GET("getEnvDetailPublic", edApi.GetEnvDetailPublic) // 环境详情开放接口
	}
}
