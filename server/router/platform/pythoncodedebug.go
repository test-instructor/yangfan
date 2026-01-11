package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/middleware"
)

type PythonCodeDebugRouter struct{}

// InitPythonCodeDebugRouter 初始化 调试信息 路由信息
func (s *PythonCodeDebugRouter) InitPythonCodeDebugRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	pcdRouter := Router.Group("pcd").Use(middleware.OperationRecord())
	pcdRouterWithoutRecord := Router.Group("pcd")
	pcdRouterWithoutAuth := PublicRouter.Group("pcd")
	{
		pcdRouter.POST("createPythonCodeDebug", pcdApi.CreatePythonCodeDebug)             // 新建调试信息
		pcdRouter.DELETE("deletePythonCodeDebug", pcdApi.DeletePythonCodeDebug)           // 删除调试信息
		pcdRouter.DELETE("deletePythonCodeDebugByIds", pcdApi.DeletePythonCodeDebugByIds) // 批量删除调试信息
		pcdRouter.PUT("updatePythonCodeDebug", pcdApi.UpdatePythonCodeDebug)              // 更新调试信息
	}
	{
		pcdRouterWithoutRecord.GET("findPythonCodeDebug", pcdApi.FindPythonCodeDebug)       // 根据ID获取调试信息
		pcdRouterWithoutRecord.GET("getPythonCodeDebugList", pcdApi.GetPythonCodeDebugList) // 获取调试信息列表
	}
	{
		pcdRouterWithoutAuth.GET("getPythonCodeDebugPublic", pcdApi.GetPythonCodeDebugPublic) // 调试信息开放接口
	}
}
