package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/middleware"
)

type AndroidDeviceOptionsRouter struct{}

// InitAndroidDeviceOptionsRouter 初始化 安卓设备 路由信息
func (s *AndroidDeviceOptionsRouter) InitAndroidDeviceOptionsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	adoRouter := Router.Group("ado").Use(middleware.OperationRecord())
	adoRouterWithoutRecord := Router.Group("ado")
	adoRouterWithoutAuth := PublicRouter.Group("ado")
	{
		adoRouter.POST("createAndroidDeviceOptions", adoApi.CreateAndroidDeviceOptions)             // 新建安卓设备
		adoRouter.DELETE("deleteAndroidDeviceOptions", adoApi.DeleteAndroidDeviceOptions)           // 删除安卓设备
		adoRouter.DELETE("deleteAndroidDeviceOptionsByIds", adoApi.DeleteAndroidDeviceOptionsByIds) // 批量删除安卓设备
		adoRouter.PUT("updateAndroidDeviceOptions", adoApi.UpdateAndroidDeviceOptions)              // 更新安卓设备
	}
	{
		adoRouterWithoutRecord.GET("findAndroidDeviceOptions", adoApi.FindAndroidDeviceOptions)       // 根据ID获取安卓设备
		adoRouterWithoutRecord.GET("getAndroidDeviceOptionsList", adoApi.GetAndroidDeviceOptionsList) // 获取安卓设备列表
	}
	{
		adoRouterWithoutAuth.GET("getAndroidDeviceOptionsPublic", adoApi.GetAndroidDeviceOptionsPublic) // 安卓设备开放接口
	}
}
