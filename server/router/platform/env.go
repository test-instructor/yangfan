package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/middleware"
)

type EnvRouter struct{}

// InitEnvRouter 初始化 环境配置 路由信息
func (s *EnvRouter) InitEnvRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	envRouter := Router.Group("env").Use(middleware.OperationRecord())
	envRouterWithoutRecord := Router.Group("env")
	envRouterWithoutAuth := PublicRouter.Group("env")
	{
		envRouter.POST("createEnv", envApi.CreateEnv)             // 新建环境配置
		envRouter.DELETE("deleteEnv", envApi.DeleteEnv)           // 删除环境配置
		envRouter.DELETE("deleteEnvByIds", envApi.DeleteEnvByIds) // 批量删除环境配置
		envRouter.PUT("updateEnv", envApi.UpdateEnv)              // 更新环境配置
	}
	{
		envRouterWithoutRecord.GET("findEnv", envApi.FindEnv)       // 根据ID获取环境配置
		envRouterWithoutRecord.GET("getEnvList", envApi.GetEnvList) // 获取环境配置列表
	}
	{
		envRouterWithoutAuth.GET("getEnvPublic", envApi.GetEnvPublic) // 环境配置开放接口
	}
}
