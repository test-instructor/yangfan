package automation

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/middleware"
)

type AutoReportRouter struct{}

// InitAutoReportRouter 初始化 自动报告 路由信息
func (s *AutoReportRouter) InitAutoReportRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	arRouter := Router.Group("ar").Use(middleware.OperationRecord())
	arRouterWithoutRecord := Router.Group("ar")
	arRouterWithoutAuth := PublicRouter.Group("ar")
	{
		arRouter.POST("createAutoReport", arApi.CreateAutoReport)             // 新建自动报告
		arRouter.DELETE("deleteAutoReport", arApi.DeleteAutoReport)           // 删除自动报告
		arRouter.DELETE("deleteAutoReportByIds", arApi.DeleteAutoReportByIds) // 批量删除自动报告
		arRouter.PUT("updateAutoReport", arApi.UpdateAutoReport)              // 更新自动报告
	}
	{
		arRouterWithoutRecord.GET("findAutoReport", arApi.FindAutoReport)       // 根据ID获取自动报告
		arRouterWithoutRecord.GET("getAutoReportList", arApi.GetAutoReportList) // 获取自动报告列表
	}
	{
		arRouterWithoutAuth.GET("getAutoReportPublic", arApi.GetAutoReportPublic) // 自动报告开放接口
	}
}
