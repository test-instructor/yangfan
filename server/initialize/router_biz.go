package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/middleware"
	"github.com/test-instructor/yangfan/server/v2/router"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}
func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]
	var projectGroup *gin.RouterGroup
	if len(routers) > 2 {
		projectGroup = routers[2]
	}
	var openGroup *gin.RouterGroup
	if len(routers) > 3 {
		openGroup = routers[3]
	}
	holder(publicGroup, privateGroup, projectGroup, openGroup)
	{
		projectmgrRouter := router.RouterGroupApp.Projectmgr
		projectmgrRouter.InitProjectRouter(privateGroup, publicGroup)
		projectmgrRouter.InitUserProjectAccessRouter(privateGroup, publicGroup)
		projectmgrRouter.InitReportNotifyRouter(privateGroup, publicGroup)
	}
	privateGroup.Use(middleware.ProjectAuth())
	{
		platformRouter := router.RouterGroupApp.Platform
		platformRouter.InitEnvRouter(privateGroup, publicGroup)
		platformRouter.InitEnvDetailRouter(privateGroup, publicGroup)
		platformRouter.InitPythonCodeRouter(privateGroup, publicGroup)
		platformRouter.InitPythonPackageRouter(privateGroup, publicGroup)
		platformRouter.InitPythonCodeDebugRouter(privateGroup, publicGroup)
		platformRouter.InitPythonCodeFuncRouter(privateGroup, publicGroup)
		platformRouter.InitRunConfigRouter(privateGroup, publicGroup)
		platformRouter.InitCategoryMenuRouter(privateGroup, publicGroup)
		platformRouter.InitRunnerNodeRouter(privateGroup, publicGroup)
		platformRouter.InitRunnerRouter(privateGroup, publicGroup)
		platformRouter.InitOpenRunnerRouter(openGroup)
		platformRouter.InitLLMModelConfigRouter(privateGroup, publicGroup)
		platformRouter.InitAndroidDeviceOptionsRouter(privateGroup, publicGroup)
		platformRouter.InitIOSDeviceOptionsRouter(privateGroup, publicGroup)
		platformRouter.InitHarmonyDeviceOptionsRouter(privateGroup, publicGroup) // 占位方法，保证文件可以正确加载，避免go空变量检测报错，请勿删除。
		platformRouter.InitBrowserDeviceOptionsRouter(privateGroup, publicGroup)
	}
	{
		automationRouter := router.RouterGroupApp.Automation
		automationRouter.InitAutoStepRouter(privateGroup, publicGroup)
		automationRouter.InitRequestRouter(privateGroup, publicGroup)
		automationRouter.InitAutoCaseStepRouter(privateGroup, publicGroup)
		automationRouter.InitAutoCaseRouter(privateGroup, publicGroup)
		automationRouter.InitTimerTaskRouter(privateGroup, publicGroup)
		automationRouter.InitAutoReportRouter(privateGroup, publicGroup)
	}
	{
		datawarehouseRouter := router.RouterGroupApp.Datawarehouse
		datawarehouseRouter.InitDataCategoryManagementRouter(privateGroup, publicGroup)
	}
}
