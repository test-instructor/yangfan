package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/middleware"
)

type CategoryMenuRouter struct{}

// InitCategoryMenuRouter 初始化 自动化菜单 路由信息
func (s *CategoryMenuRouter) InitCategoryMenuRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	cmRouter := Router.Group("cm").Use(middleware.OperationRecord())
	cmRouterWithoutRecord := Router.Group("cm")
	cmRouterWithoutAuth := PublicRouter.Group("cm")
	{
		cmRouter.POST("createCategoryMenu", cmApi.CreateCategoryMenu)             // 新建自动化菜单
		cmRouter.DELETE("deleteCategoryMenu", cmApi.DeleteCategoryMenu)           // 删除自动化菜单
		cmRouter.DELETE("deleteCategoryMenuByIds", cmApi.DeleteCategoryMenuByIds) // 批量删除自动化菜单
		cmRouter.PUT("updateCategoryMenu", cmApi.UpdateCategoryMenu)              // 更新自动化菜单
	}
	{
		cmRouterWithoutRecord.GET("findCategoryMenu", cmApi.FindCategoryMenu)       // 根据ID获取自动化菜单
		cmRouterWithoutRecord.GET("getCategoryMenuList", cmApi.GetCategoryMenuList) // 获取自动化菜单列表
	}
	{
		cmRouterWithoutAuth.GET("getCategoryMenuPublic", cmApi.GetCategoryMenuPublic) // 自动化菜单开放接口
	}
}
