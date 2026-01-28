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

type AndroidDeviceOptionsApi struct{}

// CreateAndroidDeviceOptions 创建安卓设备
// @Tags AndroidDeviceOptions
// @Summary 创建安卓设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.AndroidDeviceOptions true "创建安卓设备"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /ado/createAndroidDeviceOptions [post]
func (adoApi *AndroidDeviceOptionsApi) CreateAndroidDeviceOptions(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var ado platform.AndroidDeviceOptions
	err := c.ShouldBindJSON(&ado)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = adoService.CreateAndroidDeviceOptions(ctx, &ado)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteAndroidDeviceOptions 删除安卓设备
// @Tags AndroidDeviceOptions
// @Summary 删除安卓设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.AndroidDeviceOptions true "删除安卓设备"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /ado/deleteAndroidDeviceOptions [delete]
func (adoApi *AndroidDeviceOptionsApi) DeleteAndroidDeviceOptions(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	projectId := utils.GetProjectIDInt64(c)
	err := adoService.DeleteAndroidDeviceOptions(ctx, ID, projectId)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAndroidDeviceOptionsByIds 批量删除安卓设备
// @Tags AndroidDeviceOptions
// @Summary 批量删除安卓设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /ado/deleteAndroidDeviceOptionsByIds [delete]
func (adoApi *AndroidDeviceOptionsApi) DeleteAndroidDeviceOptionsByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := adoService.DeleteAndroidDeviceOptionsByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAndroidDeviceOptions 更新安卓设备
// @Tags AndroidDeviceOptions
// @Summary 更新安卓设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.AndroidDeviceOptions true "更新安卓设备"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /ado/updateAndroidDeviceOptions [put]
func (adoApi *AndroidDeviceOptionsApi) UpdateAndroidDeviceOptions(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var ado platform.AndroidDeviceOptions
	err := c.ShouldBindJSON(&ado)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	projectId := utils.GetProjectIDInt64(c)
	err = adoService.UpdateAndroidDeviceOptions(ctx, ado, projectId)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAndroidDeviceOptions 用id查询安卓设备
// @Tags AndroidDeviceOptions
// @Summary 用id查询安卓设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询安卓设备"
// @Success 200 {object} response.Response{data=platform.AndroidDeviceOptions,msg=string} "查询成功"
// @Router /ado/findAndroidDeviceOptions [get]
func (adoApi *AndroidDeviceOptionsApi) FindAndroidDeviceOptions(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reado, err := adoService.GetAndroidDeviceOptions(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reado, c)
}

// GetAndroidDeviceOptionsList 分页获取安卓设备列表
// @Tags AndroidDeviceOptions
// @Summary 分页获取安卓设备列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query platformReq.AndroidDeviceOptionsSearch true "分页获取安卓设备列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /ado/getAndroidDeviceOptionsList [get]
func (adoApi *AndroidDeviceOptionsApi) GetAndroidDeviceOptionsList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo platformReq.AndroidDeviceOptionsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := adoService.GetAndroidDeviceOptionsInfoList(ctx, pageInfo)
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

// GetAndroidDeviceOptionsPublic 不需要鉴权的安卓设备接口
// @Tags AndroidDeviceOptions
// @Summary 不需要鉴权的安卓设备接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ado/getAndroidDeviceOptionsPublic [get]
func (adoApi *AndroidDeviceOptionsApi) GetAndroidDeviceOptionsPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	adoService.GetAndroidDeviceOptionsPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的安卓设备接口信息",
	}, "获取成功", c)
}
