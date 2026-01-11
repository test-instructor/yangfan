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

type PythonCodeFuncApi struct{}

// CreatePythonCodeFunc 创建python函数详情
// @Tags PythonCodeFunc
// @Summary 创建python函数详情
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.PythonCodeFunc true "创建python函数详情"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /pcf/createPythonCodeFunc [post]
func (pcfApi *PythonCodeFuncApi) CreatePythonCodeFunc(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pcf platform.PythonCodeFunc
	err := c.ShouldBindJSON(&pcf)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = pcfService.CreatePythonCodeFunc(ctx, &pcf)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeletePythonCodeFunc 删除python函数详情
// @Tags PythonCodeFunc
// @Summary 删除python函数详情
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.PythonCodeFunc true "删除python函数详情"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /pcf/deletePythonCodeFunc [delete]
func (pcfApi *PythonCodeFuncApi) DeletePythonCodeFunc(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	projectId := utils.GetProjectIDInt64(c)
	err := pcfService.DeletePythonCodeFunc(ctx, ID, projectId)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeletePythonCodeFuncByIds 批量删除python函数详情
// @Tags PythonCodeFunc
// @Summary 批量删除python函数详情
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /pcf/deletePythonCodeFuncByIds [delete]
func (pcfApi *PythonCodeFuncApi) DeletePythonCodeFuncByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := pcfService.DeletePythonCodeFuncByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdatePythonCodeFunc 更新python函数详情
// @Tags PythonCodeFunc
// @Summary 更新python函数详情
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.PythonCodeFunc true "更新python函数详情"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /pcf/updatePythonCodeFunc [put]
func (pcfApi *PythonCodeFuncApi) UpdatePythonCodeFunc(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var pcf platformReq.PythonCodeFuncSearch
	err := c.ShouldBindJSON(&pcf)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	projectId := utils.GetProjectIDInt64(c)
	err = pcfService.UpdatePythonCodeFunc(ctx, pcf, projectId)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindPythonCodeFunc 用id查询python函数详情
// @Tags PythonCodeFunc
// @Summary 用id查询python函数详情
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询python函数详情"
// @Success 200 {object} response.Response{data=platform.PythonCodeFunc,msg=string} "查询成功"
// @Router /pcf/findPythonCodeFunc [get]
func (pcfApi *PythonCodeFuncApi) FindPythonCodeFunc(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	repcf, err := pcfService.GetPythonCodeFunc(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(repcf, c)
}

// GetPythonCodeFuncList 分页获取python函数详情列表
// @Tags PythonCodeFunc
// @Summary 分页获取python函数详情列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query platformReq.PythonCodeFuncSearch true "分页获取python函数详情列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /pcf/getPythonCodeFuncList [get]
func (pcfApi *PythonCodeFuncApi) GetPythonCodeFuncList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo platformReq.PythonCodeFuncSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := pcfService.GetPythonCodeFuncInfoList(ctx, pageInfo)
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

// GetPythonCodeFuncPublic 不需要鉴权的python函数详情接口
// @Tags PythonCodeFunc
// @Summary 不需要鉴权的python函数详情接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /pcf/getPythonCodeFuncPublic [get]
func (pcfApi *PythonCodeFuncApi) GetPythonCodeFuncPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	pcfService.GetPythonCodeFuncPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的python函数详情接口信息",
	}, "获取成功", c)
}
