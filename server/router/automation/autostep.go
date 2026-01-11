package automation

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/middleware"
)

type AutoStepRouter struct{}

// InitAutoStepRouter 初始化 自动化步骤 路由信息
func (s *AutoStepRouter) InitAutoStepRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	asRouter := Router.Group("as").Use(middleware.OperationRecord())
	asRouterWithoutRecord := Router.Group("as")
	asRouterWithoutAuth := PublicRouter.Group("as")
	{
		asRouter.POST("createAutoStep", asApi.CreateAutoStep)             // 新建自动化步骤
		asRouter.DELETE("deleteAutoStep", asApi.DeleteAutoStep)           // 删除自动化步骤
		asRouter.DELETE("deleteAutoStepByIds", asApi.DeleteAutoStepByIds) // 批量删除自动化步骤
		asRouter.PUT("updateAutoStep", asApi.UpdateAutoStep)              // 更新自动化步骤
	}
	{
		asRouterWithoutRecord.GET("findAutoStep", asApi.FindAutoStep)       // 根据ID获取自动化步骤
		asRouterWithoutRecord.GET("getAutoStepList", asApi.GetAutoStepList) // 获取自动化步骤列表
	}
	{
		asRouterWithoutAuth.GET("getAutoStepPublic", asApi.GetAutoStepPublic) // 自动化步骤开放接口
	}
}
