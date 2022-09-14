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
	var apiCaseApi = v1.ApiGroupApp.InterfaceCaseApiGroup.InterfaceTemplateApi
	{
		apicaseRouter.POST("createInterfaceTemplate", apiCaseApi.CreateInterfaceTemplate)             // 新建InterfaceTemplate
		apicaseRouter.DELETE("deleteInterfaceTemplate", apiCaseApi.DeleteInterfaceTemplate)           // 删除InterfaceTemplate
		apicaseRouter.DELETE("deleteInterfaceTemplateByIds", apiCaseApi.DeleteInterfaceTemplateByIds) // 批量删除InterfaceTemplate
		apicaseRouter.PUT("updateInterfaceTemplate", apiCaseApi.UpdateInterfaceTemplate)              // 更新InterfaceTemplate
		apicaseRouter.PUT("updateDebugTalk", apiCaseApi.UpdateDebugTalk)
		apicaseRouter.POST("getDebugTalk", apiCaseApi.GetDebugTalk)
	}
	{
		apicaseRouterWithoutRecord.GET("findInterfaceTemplate", apiCaseApi.FindInterfaceTemplate)       // 根据ID获取InterfaceTemplate
		apicaseRouterWithoutRecord.GET("getInterfaceTemplateList", apiCaseApi.GetInterfaceTemplateList) // 获取InterfaceTemplate列表
		apicaseRouterWithoutRecord.GET("getDebugTalk", apiCaseApi.GetDebugTalk)
	}
}
