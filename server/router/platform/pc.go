package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/middleware"
)

type PythonCodeRouter struct{}

// InitPythonCodeRouter 初始化 python 函数 路由信息
func (s *PythonCodeRouter) InitPythonCodeRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	pcRouter := Router.Group("pc").Use(middleware.OperationRecord())
	pcRouterWithoutRecord := Router.Group("pc")
	pcRouterWithoutAuth := PublicRouter.Group("pc")
	{
		//pcRouter.POST("createPythonCode", pcApi.CreatePythonCode)   // 新建python 函数
		//pcRouter.DELETE("deletePythonCode", pcApi.DeletePythonCode) // 删除python 函数
		//pcRouter.DELETE("deletePythonCodeByIds", pcApi.DeletePythonCodeByIds) // 批量删除python 函数
		pcRouter.PUT("updatePythonCode", pcApi.UpdatePythonCode) // 更新python 函数
	}
	{
		pcRouterWithoutRecord.GET("findPythonCode", pcApi.FindPythonCode)       // 根据ID获取python 函数
		pcRouterWithoutRecord.GET("getPythonCodeList", pcApi.GetPythonCodeList) // 获取python 函数列表
	}
	{
		pcRouterWithoutAuth.GET("getPythonCodePublic", pcApi.GetPythonCodePublic) // python 函数开放接口
	}
}
