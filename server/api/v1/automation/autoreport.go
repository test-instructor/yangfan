package automation

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/automation"
	automationReq "github.com/test-instructor/yangfan/server/v2/model/automation/request"
	"github.com/test-instructor/yangfan/server/v2/model/common/response"
	"github.com/test-instructor/yangfan/server/v2/utils"
	"go.uber.org/zap"
)

type AutoReportApi struct{}

// CreateAutoReport 创建自动报告
// @Tags AutoReport
// @Summary 创建自动报告
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body automation.AutoReport true "创建自动报告"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /ar/createAutoReport [post]
func (arApi *AutoReportApi) CreateAutoReport(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var ar automation.AutoReport
	err := c.ShouldBindJSON(&ar)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = arService.CreateAutoReport(ctx, &ar)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteAutoReport 删除自动报告
// @Tags AutoReport
// @Summary 删除自动报告
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body automation.AutoReport true "删除自动报告"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /ar/deleteAutoReport [delete]
func (arApi *AutoReportApi) DeleteAutoReport(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	projectId := utils.GetProjectIDInt64(c)
	err := arService.DeleteAutoReport(ctx, ID, projectId)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAutoReportByIds 批量删除自动报告
// @Tags AutoReport
// @Summary 批量删除自动报告
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /ar/deleteAutoReportByIds [delete]
func (arApi *AutoReportApi) DeleteAutoReportByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := arService.DeleteAutoReportByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAutoReport 更新自动报告
// @Tags AutoReport
// @Summary 更新自动报告
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body automation.AutoReport true "更新自动报告"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /ar/updateAutoReport [put]
func (arApi *AutoReportApi) UpdateAutoReport(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var ar automation.AutoReport
	err := c.ShouldBindJSON(&ar)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	projectId := utils.GetProjectIDInt64(c)
	err = arService.UpdateAutoReport(ctx, ar, projectId)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAutoReport 用id查询自动报告
// @Tags AutoReport
// @Summary 用id查询自动报告
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询自动报告"
// @Success 200 {object} response.Response{data=automation.AutoReport,msg=string} "查询成功"
// @Router /ar/findAutoReport [get]
func (arApi *AutoReportApi) FindAutoReport(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	rear, err := arService.GetAutoReport(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rear, c)
}

// GetAutoReportList 分页获取自动报告列表
// @Tags AutoReport
// @Summary 分页获取自动报告列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query automationReq.AutoReportSearch true "分页获取自动报告列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /ar/getAutoReportList [get]
func (arApi *AutoReportApi) GetAutoReportList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo automationReq.AutoReportSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := arService.GetAutoReportInfoList(ctx, pageInfo)
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

// GetAutoReportPublic 不需要鉴权的自动报告接口
// @Tags AutoReport
// @Summary 不需要鉴权的自动报告接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ar/getAutoReportPublic [get]
func (arApi *AutoReportApi) GetAutoReportPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	arService.GetAutoReportPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的自动报告接口信息",
	}, "获取成功", c)
}
