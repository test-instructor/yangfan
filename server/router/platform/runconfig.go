package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/middleware"
)

type RunConfigRouter struct{}

// InitRunConfigRouter 初始化 运行配置 路由信息
func (s *RunConfigRouter) InitRunConfigRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	rcRouter := Router.Group("rc").Use(middleware.OperationRecord())
	rcRouterWithoutRecord := Router.Group("rc")
	rcRouterWithoutAuth := PublicRouter.Group("rc")
	{
		rcRouter.POST("createRunConfig", rcApi.CreateRunConfig)             // 新建运行配置
		rcRouter.DELETE("deleteRunConfig", rcApi.DeleteRunConfig)           // 删除运行配置
		rcRouter.DELETE("deleteRunConfigByIds", rcApi.DeleteRunConfigByIds) // 批量删除运行配置
		rcRouter.PUT("updateRunConfig", rcApi.UpdateRunConfig)              // 更新运行配置
	}
	{
		rcRouterWithoutRecord.GET("findRunConfig", rcApi.FindRunConfig)       // 根据ID获取运行配置
		rcRouterWithoutRecord.GET("getRunConfigList", rcApi.GetRunConfigList) // 获取运行配置列表
	}
	{
		rcRouterWithoutAuth.GET("getRunConfigPublic", rcApi.GetRunConfigPublic) // 运行配置开放接口
	}
}
