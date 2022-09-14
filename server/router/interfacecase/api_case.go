package interfacecase

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/cheetah/server/api/v1"
	"github.com/test-instructor/cheetah/server/middleware"
)

type ApiCaseRouter struct {
}

// InitApiCaseRouter 初始化 ApiCase 路由信息
func (s *ApiCaseRouter) InitApiCaseRouter(Router *gin.RouterGroup) {
	taskRouter := Router.Group("").Use(middleware.OperationRecord())
	taskRouterWithoutRecord := Router.Group("")
	var taskApi = v1.ApiGroupApp.InterfaceCaseApiGroup.ApiCaseApi
	{
		taskRouter.POST("createApiCase", taskApi.CreateApiCase)             // 新建ApiCase
		taskRouter.DELETE("deleteApiCase", taskApi.DeleteApiCase)           // 删除ApiCase
		taskRouter.DELETE("deleteApiCaseByIds", taskApi.DeleteApiCaseByIds) // 批量删除ApiCase
		taskRouter.PUT("updateApiCase", taskApi.UpdateApiCase)              // 更新ApiCase
		taskRouter.POST("sortApisCase", taskApi.SortApisCase)
		taskRouter.POST("addApisCase", taskApi.AddApisCase)
		taskRouter.DELETE("delApisCase", taskApi.DelApisCase)
		taskRouter.GET("findApiTestCase", taskApi.FindApiTestCase)
		taskRouter.POST("addApiTestCase", taskApi.AddApiTestCase)
		taskRouter.POST("setApisCase", taskApi.SetApisCase)
	}
	{
		taskRouterWithoutRecord.GET("findApiCase", taskApi.FindApiCase)       // 根据ID获取ApiCase
		taskRouterWithoutRecord.GET("getApiCaseList", taskApi.GetApiCaseList) // 获取ApiCase列表
	}
}
