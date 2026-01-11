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

type PythonCodeApi struct{}

// CreatePythonCode 创建python 函数
// @Tags PythonCode
// @Summary 创建python 函数
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.PythonCode true "创建python 函数"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /pc/createPythonCode [post]
func (pcApi *PythonCodeApi) CreatePythonCode(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pc platform.PythonCode
	err := c.ShouldBindJSON(&pc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = pcService.CreatePythonCode(ctx, &pc)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeletePythonCode 删除python 函数
// @Tags PythonCode
// @Summary 删除python 函数
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.PythonCode true "删除python 函数"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /pc/deletePythonCode [delete]
func (pcApi *PythonCodeApi) DeletePythonCode(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	projectId := utils.GetProjectIDInt64(c)
	err := pcService.DeletePythonCode(ctx, ID, projectId)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeletePythonCodeByIds 批量删除python 函数
// @Tags PythonCode
// @Summary 批量删除python 函数
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /pc/deletePythonCodeByIds [delete]
func (pcApi *PythonCodeApi) DeletePythonCodeByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := pcService.DeletePythonCodeByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdatePythonCode 更新python 函数
// @Tags PythonCode
// @Summary 更新python 函数
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.PythonCode true "更新python 函数"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /pc/updatePythonCode [put]
func (pcApi *PythonCodeApi) UpdatePythonCode(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var pc platform.PythonCode
	err := c.ShouldBindJSON(&pc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	projectId := utils.GetProjectIDInt64(c)
	pc.UpdateBy = utils.GetUserID(c)
	err = pcService.UpdatePythonCode(ctx, pc, projectId)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindPythonCode 用id查询python 函数
// @Tags PythonCode
// @Summary 用id查询python 函数
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询python 函数"
// @Success 200 {object} response.Response{data=platform.PythonCode,msg=string} "查询成功"
// @Router /pc/findPythonCode [get]
func (pcApi *PythonCodeApi) FindPythonCode(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()
	var pageInfo platformReq.PythonCodeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	projectId := utils.GetProjectID(c)
	repc, err := pcService.GetPythonCode(ctx, pageInfo, projectId)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(repc, c)
}

// GetPythonCodeList 分页获取python 函数列表
// @Tags PythonCode
// @Summary 分页获取python 函数列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query platformReq.PythonCodeSearch true "分页获取python 函数列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /pc/getPythonCodeList [get]
func (pcApi *PythonCodeApi) GetPythonCodeList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo platformReq.PythonCodeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := pcService.GetPythonCodeInfoList(ctx, pageInfo)
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

// GetPythonCodePublic 不需要鉴权的python 函数接口
// @Tags PythonCode
// @Summary 不需要鉴权的python 函数接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /pc/getPythonCodePublic [get]
func (pcApi *PythonCodeApi) GetPythonCodePublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	pcService.GetPythonCodePublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的python 函数接口信息",
	}, "获取成功", c)
}
