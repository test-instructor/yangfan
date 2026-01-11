package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/common/response"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	platformReq "github.com/test-instructor/yangfan/server/v2/model/platform/request"
	"go.uber.org/zap"
)

type RunnerNodeApi struct{}

// CreateRunnerNode 创建节点
// @Tags RunnerNode
// @Summary 创建节点
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.RunnerNode true "创建节点"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /rn/createRunnerNode [post]
func (rnApi *RunnerNodeApi) CreateRunnerNode(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var rn platform.RunnerNode
	err := c.ShouldBindJSON(&rn)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = rnService.CreateRunnerNode(ctx, &rn)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteRunnerNode 删除节点
// @Tags RunnerNode
// @Summary 删除节点
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.RunnerNode true "删除节点"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /rn/deleteRunnerNode [delete]
func (rnApi *RunnerNodeApi) DeleteRunnerNode(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := rnService.DeleteRunnerNode(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteRunnerNodeByIds 批量删除节点
// @Tags RunnerNode
// @Summary 批量删除节点
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /rn/deleteRunnerNodeByIds [delete]
func (rnApi *RunnerNodeApi) DeleteRunnerNodeByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := rnService.DeleteRunnerNodeByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateRunnerNode 更新节点
// @Tags RunnerNode
// @Summary 更新节点
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.RunnerNode true "更新节点"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /rn/updateRunnerNode [put]
func (rnApi *RunnerNodeApi) UpdateRunnerNode(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var rn platform.RunnerNode
	err := c.ShouldBindJSON(&rn)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = rnService.UpdateRunnerNode(ctx, rn)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindRunnerNode 用id查询节点
// @Tags RunnerNode
// @Summary 用id查询节点
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询节点"
// @Success 200 {object} response.Response{data=platform.RunnerNode,msg=string} "查询成功"
// @Router /rn/findRunnerNode [get]
func (rnApi *RunnerNodeApi) FindRunnerNode(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	rern, err := rnService.GetRunnerNode(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rern, c)
}

// GetRunnerNodeList 分页获取节点列表
// @Tags RunnerNode
// @Summary 分页获取节点列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query platformReq.RunnerNodeSearch true "分页获取节点列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /rn/getRunnerNodeList [get]
func (rnApi *RunnerNodeApi) GetRunnerNodeList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo platformReq.RunnerNodeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := rnService.GetRunnerNodeInfoList(ctx, pageInfo)
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

// GetRunnerNodePublic 不需要鉴权的节点接口
// @Tags RunnerNode
// @Summary 不需要鉴权的节点接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /rn/getRunnerNodePublic [get]
func (rnApi *RunnerNodeApi) GetRunnerNodePublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	rnService.GetRunnerNodePublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的节点接口信息",
	}, "获取成功", c)
}
