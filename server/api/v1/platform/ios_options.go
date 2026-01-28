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

type IOSDeviceOptionsApi struct{}

// CreateIOSDeviceOptions 创建iOS设备
// @Tags IOSDeviceOptions
// @Summary 创建iOS设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.IOSDeviceOptions true "创建iOS设备"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /ido/createIOSDeviceOptions [post]
func (idoApi *IOSDeviceOptionsApi) CreateIOSDeviceOptions(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var ido platform.IOSDeviceOptions
	err := c.ShouldBindJSON(&ido)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = idoService.CreateIOSDeviceOptions(ctx, &ido)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteIOSDeviceOptions 删除iOS设备
// @Tags IOSDeviceOptions
// @Summary 删除iOS设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.IOSDeviceOptions true "删除iOS设备"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /ido/deleteIOSDeviceOptions [delete]
func (idoApi *IOSDeviceOptionsApi) DeleteIOSDeviceOptions(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	projectId := utils.GetProjectIDInt64(c)
	err := idoService.DeleteIOSDeviceOptions(ctx, ID, projectId)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteIOSDeviceOptionsByIds 批量删除iOS设备
// @Tags IOSDeviceOptions
// @Summary 批量删除iOS设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /ido/deleteIOSDeviceOptionsByIds [delete]
func (idoApi *IOSDeviceOptionsApi) DeleteIOSDeviceOptionsByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := idoService.DeleteIOSDeviceOptionsByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateIOSDeviceOptions 更新iOS设备
// @Tags IOSDeviceOptions
// @Summary 更新iOS设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.IOSDeviceOptions true "更新iOS设备"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /ido/updateIOSDeviceOptions [put]
func (idoApi *IOSDeviceOptionsApi) UpdateIOSDeviceOptions(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var ido platform.IOSDeviceOptions
	err := c.ShouldBindJSON(&ido)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	projectId := utils.GetProjectIDInt64(c)
	err = idoService.UpdateIOSDeviceOptions(ctx, ido, projectId)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindIOSDeviceOptions 用id查询iOS设备
// @Tags IOSDeviceOptions
// @Summary 用id查询iOS设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询iOS设备"
// @Success 200 {object} response.Response{data=platform.IOSDeviceOptions,msg=string} "查询成功"
// @Router /ido/findIOSDeviceOptions [get]
func (idoApi *IOSDeviceOptionsApi) FindIOSDeviceOptions(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reido, err := idoService.GetIOSDeviceOptions(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reido, c)
}

// GetIOSDeviceOptionsList 分页获取iOS设备列表
// @Tags IOSDeviceOptions
// @Summary 分页获取iOS设备列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query platformReq.IOSDeviceOptionsSearch true "分页获取iOS设备列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /ido/getIOSDeviceOptionsList [get]
func (idoApi *IOSDeviceOptionsApi) GetIOSDeviceOptionsList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo platformReq.IOSDeviceOptionsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := idoService.GetIOSDeviceOptionsInfoList(ctx, pageInfo)
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

// GetIOSDeviceOptionsPublic 不需要鉴权的iOS设备接口
// @Tags IOSDeviceOptions
// @Summary 不需要鉴权的iOS设备接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ido/getIOSDeviceOptionsPublic [get]
func (idoApi *IOSDeviceOptionsApi) GetIOSDeviceOptionsPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	idoService.GetIOSDeviceOptionsPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的iOS设备接口信息",
	}, "获取成功", c)
}
