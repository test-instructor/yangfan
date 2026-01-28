package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/middleware"
)

type BrowserDeviceOptionsRouter struct{}

// InitBrowserDeviceOptionsRouter 初始化 浏览器设备选项 路由信息
func (s *BrowserDeviceOptionsRouter) InitBrowserDeviceOptionsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	bdoRouter := Router.Group("bdo").Use(middleware.OperationRecord())
	bdoRouterWithoutRecord := Router.Group("bdo")
	bdoRouterWithoutAuth := PublicRouter.Group("bdo")
	{
		bdoRouter.POST("createBrowserDeviceOptions", bdoApi.CreateBrowserDeviceOptions)             // 新建浏览器设备选项
		bdoRouter.DELETE("deleteBrowserDeviceOptions", bdoApi.DeleteBrowserDeviceOptions)           // 删除浏览器设备选项
		bdoRouter.DELETE("deleteBrowserDeviceOptionsByIds", bdoApi.DeleteBrowserDeviceOptionsByIds) // 批量删除浏览器设备选项
		bdoRouter.PUT("updateBrowserDeviceOptions", bdoApi.UpdateBrowserDeviceOptions)              // 更新浏览器设备选项
	}
	{
		bdoRouterWithoutRecord.GET("findBrowserDeviceOptions", bdoApi.FindBrowserDeviceOptions)       // 根据ID获取浏览器设备选项
		bdoRouterWithoutRecord.GET("getBrowserDeviceOptionsList", bdoApi.GetBrowserDeviceOptionsList) // 获取浏览器设备选项列表
	}
	{
		bdoRouterWithoutAuth.GET("getBrowserDeviceOptionsPublic", bdoApi.GetBrowserDeviceOptionsPublic) // 浏览器设备选项开放接口
	}
}
