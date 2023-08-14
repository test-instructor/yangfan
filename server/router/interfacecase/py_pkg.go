package interfacecase

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/test-instructor/yangfan/server/api/v1"
	"github.com/test-instructor/yangfan/server/middleware"
)

type PyPkgRouter struct{}

func (p *PyPkgRouter) InitPyPkgRouter(Router *gin.RouterGroup) {
	pyPkgRouter := Router.Group("pyPkg").Use(middleware.OperationRecord())
	pyPkgRouterWithoutRecord := Router.Group("pyPkg")
	pkg := v1.ApiGroupApp.InterfaceCaseApiGroup.PyPkg
	{
		pyPkgRouter.POST("installPyPkg", pkg.InstallPyPkg) // 安装Python包
		pyPkgRouter.POST("uninstallPyPkg", pkg.UninstallPyPkg)
		pyPkgRouter.POST("updatePyPkg", pkg.UpdatePyPkg)
		pyPkgRouter.POST("searchPyPkg", pkg.SearchPyPkg)
	}
	{
		pyPkgRouterWithoutRecord.GET("pyPkgList", pkg.GetPyPkgList) // 获取Python包列表
		pyPkgRouter.GET("getPkgVersionList", pkg.GetPkgVersion)
	}
}
