package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/middleware"
)

type AndroidDeviceOptionsRouter struct{}

// InitAndroidDeviceOptionsRouter 初始化 设备选项 路由信息
func (s *AndroidDeviceOptionsRouter) InitAndroidDeviceOptionsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	adoRouter := Router.Group("ado").Use(middleware.OperationRecord())
	adoRouterWithoutRecord := Router.Group("ado")
	adoRouterWithoutAuth := PublicRouter.Group("ado")
	{
		adoRouter.POST("createAndroidDeviceOptions", adoApi.CreateAndroidDeviceOptions)             // 新建设备选项
		adoRouter.DELETE("deleteAndroidDeviceOptions", adoApi.DeleteAndroidDeviceOptions)           // 删除设备选项
		adoRouter.DELETE("deleteAndroidDeviceOptionsByIds", adoApi.DeleteAndroidDeviceOptionsByIds) // 批量删除设备选项
		adoRouter.PUT("updateAndroidDeviceOptions", adoApi.UpdateAndroidDeviceOptions)              // 更新设备选项
	}
	{
		adoRouterWithoutRecord.GET("findAndroidDeviceOptions", adoApi.FindAndroidDeviceOptions)       // 根据ID获取设备选项
		adoRouterWithoutRecord.GET("getAndroidDeviceOptionsList", adoApi.GetAndroidDeviceOptionsList) // 获取设备选项列表
	}
	{
		adoRouterWithoutAuth.GET("getAndroidDeviceOptionsPublic", adoApi.GetAndroidDeviceOptionsPublic) // 设备选项开放接口
	}
}
