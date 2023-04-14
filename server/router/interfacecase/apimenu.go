package interfacecase

import (
	"github.com/gin-gonic/gin"

	"github.com/test-instructor/yangfan/server/api/v1"
	"github.com/test-instructor/yangfan/server/middleware"
)

type ApiMenuRouter struct {
}

// InitApiMenuRouter  初始化 ApiMenu 路由信息
func (s *ApiMenuRouter) InitApiMenuRouter(Router *gin.RouterGroup) {
	apiMenuRouter := Router.Group("").Use(middleware.OperationRecord())
	apiMenuRouterWithoutRecord := Router.Group("")
	var apiCaseApi = v1.ApiGroupApp.InterfaceCaseApiGroup.ApiMenuApi
	{
		apiMenuRouter.POST("createApiMenu", apiCaseApi.CreateApiMenu)             // 新建ApiMenu
		apiMenuRouter.DELETE("deleteApiMenu", apiCaseApi.DeleteApiMenu)           // 删除ApiMenu
		apiMenuRouter.DELETE("deleteApiMenuByIds", apiCaseApi.DeleteApiMenuByIds) // 批量删除ApiMenu
		apiMenuRouter.PUT("updateApiMenu", apiCaseApi.UpdateApiMenu)              // 更新ApiMenu
	}
	{
		apiMenuRouterWithoutRecord.GET("findApiMenu", apiCaseApi.FindApiMenu)       // 根据ID获取ApiMenu
		apiMenuRouterWithoutRecord.GET("getApiMenuList", apiCaseApi.GetApiMenuList) // 获取ApiMenu列表
	}
}
