package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/middleware"
)

type PythonCodeFuncRouter struct{}

// InitPythonCodeFuncRouter 初始化 python函数详情 路由信息
func (s *PythonCodeFuncRouter) InitPythonCodeFuncRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	pcfRouter := Router.Group("pcf").Use(middleware.OperationRecord())
	pcfRouterWithoutRecord := Router.Group("pcf")
	pcfRouterWithoutAuth := PublicRouter.Group("pcf")
	{
		pcfRouter.POST("createPythonCodeFunc", pcfApi.CreatePythonCodeFunc)             // 新建python函数详情
		pcfRouter.DELETE("deletePythonCodeFunc", pcfApi.DeletePythonCodeFunc)           // 删除python函数详情
		pcfRouter.DELETE("deletePythonCodeFuncByIds", pcfApi.DeletePythonCodeFuncByIds) // 批量删除python函数详情
		pcfRouter.PUT("updatePythonCodeFunc", pcfApi.UpdatePythonCodeFunc)              // 更新python函数详情
	}
	{
		pcfRouterWithoutRecord.GET("findPythonCodeFunc", pcfApi.FindPythonCodeFunc)       // 根据ID获取python函数详情
		pcfRouterWithoutRecord.GET("getPythonCodeFuncList", pcfApi.GetPythonCodeFuncList) // 获取python函数详情列表
	}
	{
		pcfRouterWithoutAuth.GET("getPythonCodeFuncPublic", pcfApi.GetPythonCodeFuncPublic) // python函数详情开放接口
	}
}
