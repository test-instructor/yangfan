package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/middleware"
)

type HarmonyDeviceOptionsRouter struct{}

// InitHarmonyDeviceOptionsRouter 初始化 设备选项 路由信息
func (s *HarmonyDeviceOptionsRouter) InitHarmonyDeviceOptionsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	hdoRouter := Router.Group("hdo").Use(middleware.OperationRecord())
	hdoRouterWithoutRecord := Router.Group("hdo")
	hdoRouterWithoutAuth := PublicRouter.Group("hdo")
	{
		hdoRouter.POST("createHarmonyDeviceOptions", hdoApi.CreateHarmonyDeviceOptions)             // 新建设备选项
		hdoRouter.DELETE("deleteHarmonyDeviceOptions", hdoApi.DeleteHarmonyDeviceOptions)           // 删除设备选项
		hdoRouter.DELETE("deleteHarmonyDeviceOptionsByIds", hdoApi.DeleteHarmonyDeviceOptionsByIds) // 批量删除设备选项
		hdoRouter.PUT("updateHarmonyDeviceOptions", hdoApi.UpdateHarmonyDeviceOptions)              // 更新设备选项
	}
	{
		hdoRouterWithoutRecord.GET("findHarmonyDeviceOptions", hdoApi.FindHarmonyDeviceOptions)       // 根据ID获取设备选项
		hdoRouterWithoutRecord.GET("getHarmonyDeviceOptionsList", hdoApi.GetHarmonyDeviceOptionsList) // 获取设备选项列表
	}
	{
		hdoRouterWithoutAuth.GET("getHarmonyDeviceOptionsPublic", hdoApi.GetHarmonyDeviceOptionsPublic) // 设备选项开放接口
	}
}
