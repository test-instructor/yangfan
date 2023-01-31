package interfacecase

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/test-instructor/cheetah/server/api/v1"
	"github.com/test-instructor/cheetah/server/middleware"
)

func (p *PyPkgRouter) InitPyPkgRouter(Router *gin.RouterGroup) {
	pyPkgRouter := Router.Group("pyPkg").Use(middleware.OperationRecord())
	pyPkgRouterWithoutRecord := Router.Group("pyPkg")
	pkg := v1.ApiGroupApp.InterfaceCaseApiGroup.PyPkg
	{
		pyPkgRouter.GET("pyPkgList", pkg.GetPyPkgList) // 获取Python包列表
	}
	{
		pyPkgRouterWithoutRecord.POST("installPyPkg", pkg.InstallPyPkg) // 安装Python包
		pyPkgRouterWithoutRecord.POST("uninstallPyPkg", pkg.UninstallPyPkg)
		pyPkgRouterWithoutRecord.POST("updatePyPkg", pkg.UpdatePyPkg)
		pyPkgRouterWithoutRecord.POST("searchPyPkg", pkg.SearchPyPkg)
	}
}
