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

type RunConfigApi struct{}

// CreateRunConfig 创建运行配置
// @Tags RunConfig
// @Summary 创建运行配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.RunConfig true "创建运行配置"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /rc/createRunConfig [post]
func (rcApi *RunConfigApi) CreateRunConfig(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var rc platform.RunConfig
	err := c.ShouldBindJSON(&rc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = rcService.CreateRunConfig(ctx, &rc)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteRunConfig 删除运行配置
// @Tags RunConfig
// @Summary 删除运行配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.RunConfig true "删除运行配置"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /rc/deleteRunConfig [delete]
func (rcApi *RunConfigApi) DeleteRunConfig(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	projectId := utils.GetProjectIDInt64(c)
	err := rcService.DeleteRunConfig(ctx, ID, projectId)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteRunConfigByIds 批量删除运行配置
// @Tags RunConfig
// @Summary 批量删除运行配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /rc/deleteRunConfigByIds [delete]
func (rcApi *RunConfigApi) DeleteRunConfigByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := rcService.DeleteRunConfigByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateRunConfig 更新运行配置
// @Tags RunConfig
// @Summary 更新运行配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.RunConfig true "更新运行配置"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /rc/updateRunConfig [put]
func (rcApi *RunConfigApi) UpdateRunConfig(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var rc platform.RunConfig
	err := c.ShouldBindJSON(&rc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	projectId := utils.GetProjectIDInt64(c)
	err = rcService.UpdateRunConfig(ctx, rc, projectId)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindRunConfig 用id查询运行配置
// @Tags RunConfig
// @Summary 用id查询运行配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询运行配置"
// @Success 200 {object} response.Response{data=platform.RunConfig,msg=string} "查询成功"
// @Router /rc/findRunConfig [get]
func (rcApi *RunConfigApi) FindRunConfig(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	rerc, err := rcService.GetRunConfig(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rerc, c)
}

// GetRunConfigList 分页获取运行配置列表
// @Tags RunConfig
// @Summary 分页获取运行配置列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query platformReq.RunConfigSearch true "分页获取运行配置列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /rc/getRunConfigList [get]
func (rcApi *RunConfigApi) GetRunConfigList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo platformReq.RunConfigSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := rcService.GetRunConfigInfoList(ctx, pageInfo)
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

// GetRunConfigPublic 不需要鉴权的运行配置接口
// @Tags RunConfig
// @Summary 不需要鉴权的运行配置接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /rc/getRunConfigPublic [get]
func (rcApi *RunConfigApi) GetRunConfigPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	rcService.GetRunConfigPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的运行配置接口信息",
	}, "获取成功", c)
}
