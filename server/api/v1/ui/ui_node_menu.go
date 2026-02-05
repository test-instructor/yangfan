package ui

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/model/common/response"
	sysModel "github.com/test-instructor/yangfan/server/v2/model/system"
)

type UINodeMenuApi struct{}

func (a *UINodeMenuApi) GetMenuTree(c *gin.Context) {
	menus, err := uiNodeMenuService.GetMenuTree()
	if err != nil {
		response.FailWithMessage("获取菜单失败:"+err.Error(), c)
		return
	}
	response.OkWithData(menus, c)
}

func (a *UINodeMenuApi) CreateMenu(c *gin.Context) {
	var menu sysModel.SysUINodeMenu
	if err := c.ShouldBindJSON(&menu); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := uiNodeMenuService.CreateMenu(menu); err != nil {
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (a *UINodeMenuApi) UpdateMenu(c *gin.Context) {
	var menu sysModel.SysUINodeMenu
	if err := c.ShouldBindJSON(&menu); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := uiNodeMenuService.UpdateMenu(menu); err != nil {
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (a *UINodeMenuApi) DeleteMenu(c *gin.Context) {
	idStr := c.Query("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	if err := uiNodeMenuService.DeleteMenu(uint(id)); err != nil {
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (a *UINodeMenuApi) ListMenus(c *gin.Context) {
	list, err := uiNodeMenuService.ListMenus()
	if err != nil {
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithData(list, c)
}
