package interfacecase

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/cheetah/server/api/v1"
	"github.com/test-instructor/cheetah/server/middleware"
)

type ApiConfigRouter struct {
}

// InitApiConfigRouter 初始化 ApiConfig 路由信息
func (s *ApiConfigRouter) InitApiConfigRouter(Router *gin.RouterGroup) {
	acRouter := Router.Group("").Use(middleware.OperationRecord())
	acRouterWithoutRecord := Router.Group("")
	var acApi = v1.ApiGroupApp.InterfaceCaseApiGroup.ApiConfigApi
	{
		acRouter.POST("createApiConfig", acApi.CreateApiConfig)             // 新建ApiConfig
		acRouter.DELETE("deleteApiConfig", acApi.DeleteApiConfig)           // 删除ApiConfig
		acRouter.DELETE("deleteApiConfigByIds", acApi.DeleteApiConfigByIds) // 批量删除ApiConfig
		acRouter.PUT("updateApiConfig", acApi.UpdateApiConfig)              // 更新ApiConfig
	}
	{
		acRouterWithoutRecord.GET("findApiConfig", acApi.FindApiConfig)       // 根据ID获取ApiConfig
		acRouterWithoutRecord.GET("getApiConfigList", acApi.GetApiConfigList) // 获取ApiConfig列表
	}
}
