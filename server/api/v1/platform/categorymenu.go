package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/common/response"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	platformReq "github.com/test-instructor/yangfan/server/v2/model/platform/request"
	"github.com/test-instructor/yangfan/server/v2/utils"
	"go.uber.org/zap"
)

type CategoryMenuApi struct{}

// CreateCategoryMenu 创建自动化菜单
// @Tags CategoryMenu
// @Summary 创建自动化菜单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.CategoryMenu true "创建自动化菜单"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /cm/createCategoryMenu [post]
func (cmApi *CategoryMenuApi) CreateCategoryMenu(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var cm platform.CategoryMenu
	err := c.ShouldBindJSON(&cm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = cmService.CreateCategoryMenu(ctx, &cm)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteCategoryMenu 删除自动化菜单
// @Tags CategoryMenu
// @Summary 删除自动化菜单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.CategoryMenu true "删除自动化菜单"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /cm/deleteCategoryMenu [delete]
func (cmApi *CategoryMenuApi) DeleteCategoryMenu(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("id")
	projectId := utils.GetProjectIDInt64(c)
	err := cmService.DeleteCategoryMenu(ctx, ID, projectId)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteCategoryMenuByIds 批量删除自动化菜单
// @Tags CategoryMenu
// @Summary 批量删除自动化菜单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /cm/deleteCategoryMenuByIds [delete]
func (cmApi *CategoryMenuApi) DeleteCategoryMenuByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := cmService.DeleteCategoryMenuByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateCategoryMenu 更新自动化菜单
// @Tags CategoryMenu
// @Summary 更新自动化菜单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.CategoryMenu true "更新自动化菜单"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /cm/updateCategoryMenu [put]
func (cmApi *CategoryMenuApi) UpdateCategoryMenu(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var cm platform.CategoryMenu
	err := c.ShouldBindJSON(&cm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	projectId := utils.GetProjectIDInt64(c)
	err = cmService.UpdateCategoryMenu(ctx, cm, projectId)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindCategoryMenu 用id查询自动化菜单
// @Tags CategoryMenu
// @Summary 用id查询自动化菜单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询自动化菜单"
// @Success 200 {object} response.Response{data=platform.CategoryMenu,msg=string} "查询成功"
// @Router /cm/findCategoryMenu [get]
func (cmApi *CategoryMenuApi) FindCategoryMenu(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	recm, err := cmService.GetCategoryMenu(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(recm, c)
}

// GetCategoryMenuList 分页获取自动化菜单列表
// @Tags CategoryMenu
// @Summary 分页获取自动化菜单列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query platformReq.CategoryMenuSearch true "分页获取自动化菜单列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /cm/getCategoryMenuList [get]
func (cmApi *CategoryMenuApi) GetCategoryMenuList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo platformReq.CategoryMenuSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, err := cmService.GetCategoryMenuInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List: list,
	}, "获取成功", c)
}

// GetCategoryMenuPublic 不需要鉴权的自动化菜单接口
// @Tags CategoryMenu
// @Summary 不需要鉴权的自动化菜单接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cm/getCategoryMenuPublic [get]
func (cmApi *CategoryMenuApi) GetCategoryMenuPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	cmService.GetCategoryMenuPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的自动化菜单接口信息",
	}, "获取成功", c)
}
