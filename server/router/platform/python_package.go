package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/middleware"
)

type PythonPackageRouter struct{}

// InitPythonPackageRouter 初始化 py 第三方库 路由信息
func (s *PythonPackageRouter) InitPythonPackageRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	ppRouter := Router.Group("pp").Use(middleware.OperationRecord())
	ppRouterWithoutRecord := Router.Group("pp")
	ppRouterWithoutAuth := PublicRouter.Group("pp")
	{
		ppRouter.POST("createPythonPackage", ppApi.CreatePythonPackage)             // 新建py 第三方库
		ppRouter.DELETE("deletePythonPackage", ppApi.DeletePythonPackage)           // 删除py 第三方库
		ppRouter.DELETE("deletePythonPackageByIds", ppApi.DeletePythonPackageByIds) // 批量删除py 第三方库
		ppRouter.PUT("updatePythonPackage", ppApi.UpdatePythonPackage)              // 更新py 第三方库
	}
	{
		ppRouterWithoutRecord.GET("findPythonPackage", ppApi.FindPythonPackage)               // 根据ID获取py 第三方库
		ppRouterWithoutRecord.GET("findPythonPackageVersion", ppApi.FindPythonPackageVersion) // 根据ID获取py 第三方库
		ppRouterWithoutRecord.GET("getPythonPackageList", ppApi.GetPythonPackageList)         // 获取py 第三方库列表
	}
	{
		ppRouterWithoutAuth.GET("getPythonPackagePublic", ppApi.GetPythonPackagePublic) // py 第三方库开放接口
	}
}
