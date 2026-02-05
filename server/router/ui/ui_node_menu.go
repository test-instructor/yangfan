package ui

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/middleware"
)

type UINodeMenuRouter struct{}

func (r *UINodeMenuRouter) InitUINodeMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	menuRouter := Router.Group("ui/node/menu").Use(middleware.OperationRecord())
	menuRouterWithoutRecord := Router.Group("ui/node/menu")
	{
		menuRouter.POST("create", uiNodeMenuApi.CreateMenu)
		menuRouter.PUT("update", uiNodeMenuApi.UpdateMenu)
		menuRouter.DELETE("delete", uiNodeMenuApi.DeleteMenu)
	}
	{
		menuRouterWithoutRecord.POST("getMenuTree", uiNodeMenuApi.GetMenuTree)
		menuRouterWithoutRecord.GET("list", uiNodeMenuApi.ListMenus)
	}
	return menuRouter
}
