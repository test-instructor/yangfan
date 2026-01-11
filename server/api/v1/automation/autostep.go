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

type AutoStepApi struct{}

// CreateAutoStep 创建自动化步骤
// @Tags AutoStep
// @Summary 创建自动化步骤
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body automation.AutoStep true "创建自动化步骤"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /as/createAutoStep [post]
func (asApi *AutoStepApi) CreateAutoStep(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var as automation.AutoStep
	err := c.ShouldBindJSON(&as)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = asService.CreateAutoStep(ctx, &as)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteAutoStep 删除自动化步骤
// @Tags AutoStep
// @Summary 删除自动化步骤
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body automation.AutoStep true "删除自动化步骤"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /as/deleteAutoStep [delete]
func (asApi *AutoStepApi) DeleteAutoStep(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	projectId := utils.GetProjectIDInt64(c)
	err := asService.DeleteAutoStep(ctx, ID, projectId)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAutoStepByIds 批量删除自动化步骤
// @Tags AutoStep
// @Summary 批量删除自动化步骤
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /as/deleteAutoStepByIds [delete]
func (asApi *AutoStepApi) DeleteAutoStepByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := asService.DeleteAutoStepByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAutoStep 更新自动化步骤
// @Tags AutoStep
// @Summary 更新自动化步骤
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body automation.AutoStep true "更新自动化步骤"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /as/updateAutoStep [put]
func (asApi *AutoStepApi) UpdateAutoStep(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var as automation.AutoStep
	err := c.ShouldBindJSON(&as)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	projectId := utils.GetProjectIDInt64(c)
	err = asService.UpdateAutoStep(ctx, as, projectId)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAutoStep 用id查询自动化步骤
// @Tags AutoStep
// @Summary 用id查询自动化步骤
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询自动化步骤"
// @Success 200 {object} response.Response{data=automation.AutoStep,msg=string} "查询成功"
// @Router /as/findAutoStep [get]
func (asApi *AutoStepApi) FindAutoStep(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reas, err := asService.GetAutoStep(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reas, c)
}

// GetAutoStepList 分页获取自动化步骤列表
// @Tags AutoStep
// @Summary 分页获取自动化步骤列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query automationReq.AutoStepSearch true "分页获取自动化步骤列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /as/getAutoStepList [get]
func (asApi *AutoStepApi) GetAutoStepList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo automationReq.AutoStepSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := asService.GetAutoStepInfoList(ctx, pageInfo)
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

// GetAutoStepPublic 不需要鉴权的自动化步骤接口
// @Tags AutoStep
// @Summary 不需要鉴权的自动化步骤接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /as/getAutoStepPublic [get]
func (asApi *AutoStepApi) GetAutoStepPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	asService.GetAutoStepPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的自动化步骤接口信息",
	}, "获取成功", c)
}
