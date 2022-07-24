package interfacecase

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/cheetah/server/api/v1"
	"github.com/test-instructor/cheetah/server/middleware"
)

type ApiMenuRouter struct {
}

// ApiMenuRouter 初始化 ApiMenu 路由信息
func (s *ApiMenuRouter) InitApiMenuRouter(Router *gin.RouterGroup) {
	apiMenuRouter := Router.Group("").Use(middleware.OperationRecord())
	apiMenuRouterWithoutRecord := Router.Group("")
	var apicaseApi = v1.ApiGroupApp.InterfaceCaseApiGroup.ApiMenuApi
	{
		apiMenuRouter.POST("createApiMenu", apicaseApi.CreateApiMenu)             // 新建ApiMenu
		apiMenuRouter.DELETE("deleteApiMenu", apicaseApi.DeleteApiMenu)           // 删除ApiMenu
		apiMenuRouter.DELETE("deleteApiMenuByIds", apicaseApi.DeleteApiMenuByIds) // 批量删除ApiMenu
		apiMenuRouter.PUT("updateApiMenu", apicaseApi.UpdateApiMenu)              // 更新ApiMenu
	}
	{
		apiMenuRouterWithoutRecord.GET("findApiMenu", apicaseApi.FindApiMenu)       // 根据ID获取ApiMenu
		apiMenuRouterWithoutRecord.GET("getApiMenuList", apicaseApi.GetApiMenuList) // 获取ApiMenu列表
	}
}
