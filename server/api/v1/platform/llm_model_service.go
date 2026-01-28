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

type LLMModelConfigApi struct{}

// CreateLLMModelConfig 创建大语言模型配置
// @Tags LLMModelConfig
// @Summary 创建大语言模型配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.LLMModelConfig true "创建大语言模型配置"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /llmconfig/createLLMModelConfig [post]
func (llmconfigApi *LLMModelConfigApi) CreateLLMModelConfig(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var llmconfig platform.LLMModelConfig
	err := c.ShouldBindJSON(&llmconfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = llmconfigService.CreateLLMModelConfig(ctx, &llmconfig)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteLLMModelConfig 删除大语言模型配置
// @Tags LLMModelConfig
// @Summary 删除大语言模型配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.LLMModelConfig true "删除大语言模型配置"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /llmconfig/deleteLLMModelConfig [delete]
func (llmconfigApi *LLMModelConfigApi) DeleteLLMModelConfig(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	projectId := utils.GetProjectIDInt64(c)
	err := llmconfigService.DeleteLLMModelConfig(ctx, ID, projectId)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteLLMModelConfigByIds 批量删除大语言模型配置
// @Tags LLMModelConfig
// @Summary 批量删除大语言模型配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /llmconfig/deleteLLMModelConfigByIds [delete]
func (llmconfigApi *LLMModelConfigApi) DeleteLLMModelConfigByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := llmconfigService.DeleteLLMModelConfigByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateLLMModelConfig 更新大语言模型配置
// @Tags LLMModelConfig
// @Summary 更新大语言模型配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.LLMModelConfig true "更新大语言模型配置"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /llmconfig/updateLLMModelConfig [put]
func (llmconfigApi *LLMModelConfigApi) UpdateLLMModelConfig(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var llmconfig platform.LLMModelConfig
	err := c.ShouldBindJSON(&llmconfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	projectId := utils.GetProjectIDInt64(c)
	err = llmconfigService.UpdateLLMModelConfig(ctx, llmconfig, projectId)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindLLMModelConfig 用id查询大语言模型配置
// @Tags LLMModelConfig
// @Summary 用id查询大语言模型配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询大语言模型配置"
// @Success 200 {object} response.Response{data=platform.LLMModelConfig,msg=string} "查询成功"
// @Router /llmconfig/findLLMModelConfig [get]
func (llmconfigApi *LLMModelConfigApi) FindLLMModelConfig(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	rellmconfig, err := llmconfigService.GetLLMModelConfig(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rellmconfig, c)
}

// GetLLMModelConfigList 分页获取大语言模型配置列表
// @Tags LLMModelConfig
// @Summary 分页获取大语言模型配置列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query platformReq.LLMModelConfigSearch true "分页获取大语言模型配置列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /llmconfig/getLLMModelConfigList [get]
func (llmconfigApi *LLMModelConfigApi) GetLLMModelConfigList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo platformReq.LLMModelConfigSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := llmconfigService.GetLLMModelConfigInfoList(ctx, pageInfo)
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

// GetLLMModelConfigPublic 不需要鉴权的大语言模型配置接口
// @Tags LLMModelConfig
// @Summary 不需要鉴权的大语言模型配置接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /llmconfig/getLLMModelConfigPublic [get]
func (llmconfigApi *LLMModelConfigApi) GetLLMModelConfigPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	llmconfigService.GetLLMModelConfigPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的大语言模型配置接口信息",
	}, "获取成功", c)
}
