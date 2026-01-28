package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/middleware"
)

type IOSDeviceOptionsRouter struct{}

// InitIOSDeviceOptionsRouter 初始化 iOS设备选项 路由信息
func (s *IOSDeviceOptionsRouter) InitIOSDeviceOptionsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	idoRouter := Router.Group("ido").Use(middleware.OperationRecord())
	idoRouterWithoutRecord := Router.Group("ido")
	idoRouterWithoutAuth := PublicRouter.Group("ido")
	{
		idoRouter.POST("createIOSDeviceOptions", idoApi.CreateIOSDeviceOptions)             // 新建iOS设备选项
		idoRouter.DELETE("deleteIOSDeviceOptions", idoApi.DeleteIOSDeviceOptions)           // 删除iOS设备选项
		idoRouter.DELETE("deleteIOSDeviceOptionsByIds", idoApi.DeleteIOSDeviceOptionsByIds) // 批量删除iOS设备选项
		idoRouter.PUT("updateIOSDeviceOptions", idoApi.UpdateIOSDeviceOptions)              // 更新iOS设备选项
	}
	{
		idoRouterWithoutRecord.GET("findIOSDeviceOptions", idoApi.FindIOSDeviceOptions)       // 根据ID获取iOS设备选项
		idoRouterWithoutRecord.GET("getIOSDeviceOptionsList", idoApi.GetIOSDeviceOptionsList) // 获取iOS设备选项列表
	}
	{
		idoRouterWithoutAuth.GET("getIOSDeviceOptionsPublic", idoApi.GetIOSDeviceOptionsPublic) // iOS设备选项开放接口
	}
}
