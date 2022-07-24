package interfacecase

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/cheetah/server/api/v1"
	"github.com/test-instructor/cheetah/server/middleware"
)

type InterfaceTemplateRouter struct {
}

// InitInterfaceTemplateRouter 初始化 InterfaceTemplate 路由信息
func (s *InterfaceTemplateRouter) InitInterfaceTemplateRouter(Router *gin.RouterGroup) {
	apicaseRouter := Router.Group("").Use(middleware.OperationRecord())
	apicaseRouterWithoutRecord := Router.Group("")
	var apicaseApi = v1.ApiGroupApp.InterfaceCaseApiGroup.InterfaceTemplateApi
	{
		apicaseRouter.POST("createInterfaceTemplate", apicaseApi.CreateInterfaceTemplate)             // 新建InterfaceTemplate
		apicaseRouter.DELETE("deleteInterfaceTemplate", apicaseApi.DeleteInterfaceTemplate)           // 删除InterfaceTemplate
		apicaseRouter.DELETE("deleteInterfaceTemplateByIds", apicaseApi.DeleteInterfaceTemplateByIds) // 批量删除InterfaceTemplate
		apicaseRouter.PUT("updateInterfaceTemplate", apicaseApi.UpdateInterfaceTemplate)              // 更新InterfaceTemplate
		apicaseRouter.PUT("updateDebugTalk", apicaseApi.UpdateDebugTalk)
		apicaseRouter.POST("getDebugTalk", apicaseApi.GetDebugTalk)
	}
	{
		apicaseRouterWithoutRecord.GET("findInterfaceTemplate", apicaseApi.FindInterfaceTemplate)       // 根据ID获取InterfaceTemplate
		apicaseRouterWithoutRecord.GET("getInterfaceTemplateList", apicaseApi.GetInterfaceTemplateList) // 获取InterfaceTemplate列表
		apicaseRouterWithoutRecord.GET("getDebugTalk", apicaseApi.GetDebugTalk)
	}
}
