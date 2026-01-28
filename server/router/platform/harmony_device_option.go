package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/middleware"
)

type HarmonyDeviceOptionsRouter struct{}

// InitHarmonyDeviceOptionsRouter 初始化 鸿蒙设备 路由信息
func (s *HarmonyDeviceOptionsRouter) InitHarmonyDeviceOptionsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	hdoRouter := Router.Group("hdo").Use(middleware.OperationRecord())
	hdoRouterWithoutRecord := Router.Group("hdo")
	hdoRouterWithoutAuth := PublicRouter.Group("hdo")
	{
		hdoRouter.POST("createHarmonyDeviceOptions", hdoApi.CreateHarmonyDeviceOptions)             // 新建鸿蒙设备
		hdoRouter.DELETE("deleteHarmonyDeviceOptions", hdoApi.DeleteHarmonyDeviceOptions)           // 删除鸿蒙设备
		hdoRouter.DELETE("deleteHarmonyDeviceOptionsByIds", hdoApi.DeleteHarmonyDeviceOptionsByIds) // 批量删除鸿蒙设备
		hdoRouter.PUT("updateHarmonyDeviceOptions", hdoApi.UpdateHarmonyDeviceOptions)              // 更新鸿蒙设备
	}
	{
		hdoRouterWithoutRecord.GET("findHarmonyDeviceOptions", hdoApi.FindHarmonyDeviceOptions)       // 根据ID获取鸿蒙设备
		hdoRouterWithoutRecord.GET("getHarmonyDeviceOptionsList", hdoApi.GetHarmonyDeviceOptionsList) // 获取鸿蒙设备列表
	}
	{
		hdoRouterWithoutAuth.GET("getHarmonyDeviceOptionsPublic", hdoApi.GetHarmonyDeviceOptionsPublic) // 鸿蒙设备开放接口
	}
}
