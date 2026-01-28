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

type BrowserDeviceOptionsApi struct{}

// CreateBrowserDeviceOptions 创建浏览器设备选项
// @Tags BrowserDeviceOptions
// @Summary 创建浏览器设备选项
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.BrowserDeviceOptions true "创建浏览器设备选项"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /bdo/createBrowserDeviceOptions [post]
func (bdoApi *BrowserDeviceOptionsApi) CreateBrowserDeviceOptions(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var bdo platform.BrowserDeviceOptions
	err := c.ShouldBindJSON(&bdo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = bdoService.CreateBrowserDeviceOptions(ctx, &bdo)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteBrowserDeviceOptions 删除浏览器设备选项
// @Tags BrowserDeviceOptions
// @Summary 删除浏览器设备选项
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.BrowserDeviceOptions true "删除浏览器设备选项"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /bdo/deleteBrowserDeviceOptions [delete]
func (bdoApi *BrowserDeviceOptionsApi) DeleteBrowserDeviceOptions(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	projectId := utils.GetProjectIDInt64(c)
	err := bdoService.DeleteBrowserDeviceOptions(ctx, ID, projectId)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteBrowserDeviceOptionsByIds 批量删除浏览器设备选项
// @Tags BrowserDeviceOptions
// @Summary 批量删除浏览器设备选项
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /bdo/deleteBrowserDeviceOptionsByIds [delete]
func (bdoApi *BrowserDeviceOptionsApi) DeleteBrowserDeviceOptionsByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := bdoService.DeleteBrowserDeviceOptionsByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateBrowserDeviceOptions 更新浏览器设备选项
// @Tags BrowserDeviceOptions
// @Summary 更新浏览器设备选项
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.BrowserDeviceOptions true "更新浏览器设备选项"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /bdo/updateBrowserDeviceOptions [put]
func (bdoApi *BrowserDeviceOptionsApi) UpdateBrowserDeviceOptions(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var bdo platform.BrowserDeviceOptions
	err := c.ShouldBindJSON(&bdo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	projectId := utils.GetProjectIDInt64(c)
	err = bdoService.UpdateBrowserDeviceOptions(ctx, bdo, projectId)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindBrowserDeviceOptions 用id查询浏览器设备选项
// @Tags BrowserDeviceOptions
// @Summary 用id查询浏览器设备选项
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询浏览器设备选项"
// @Success 200 {object} response.Response{data=platform.BrowserDeviceOptions,msg=string} "查询成功"
// @Router /bdo/findBrowserDeviceOptions [get]
func (bdoApi *BrowserDeviceOptionsApi) FindBrowserDeviceOptions(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	rebdo, err := bdoService.GetBrowserDeviceOptions(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rebdo, c)
}

// GetBrowserDeviceOptionsList 分页获取浏览器设备选项列表
// @Tags BrowserDeviceOptions
// @Summary 分页获取浏览器设备选项列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query platformReq.BrowserDeviceOptionsSearch true "分页获取浏览器设备选项列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /bdo/getBrowserDeviceOptionsList [get]
func (bdoApi *BrowserDeviceOptionsApi) GetBrowserDeviceOptionsList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo platformReq.BrowserDeviceOptionsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := bdoService.GetBrowserDeviceOptionsInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetBrowserDeviceOptionsPublic 不需要鉴权的浏览器设备选项接口
// @Tags BrowserDeviceOptions
// @Summary 不需要鉴权的浏览器设备选项接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /bdo/getBrowserDeviceOptionsPublic [get]
func (bdoApi *BrowserDeviceOptionsApi) GetBrowserDeviceOptionsPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	bdoService.GetBrowserDeviceOptionsPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的浏览器设备选项接口信息",
	}, "获取成功", c)
}
