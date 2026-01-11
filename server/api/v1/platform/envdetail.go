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

type EnvDetailApi struct{}

// CreateEnvDetail 创建环境详情
// @Tags EnvDetail
// @Summary 创建环境详情
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.EnvDetail true "创建环境详情"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /ed/createEnvDetail [post]
func (edApi *EnvDetailApi) CreateEnvDetail(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var ed platform.EnvDetail
	err := c.ShouldBindJSON(&ed)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = edService.CreateEnvDetail(ctx, &ed)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteEnvDetail 删除环境详情
// @Tags EnvDetail
// @Summary 删除环境详情
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.EnvDetail true "删除环境详情"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /ed/deleteEnvDetail [delete]
func (edApi *EnvDetailApi) DeleteEnvDetail(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	projectId := utils.GetProjectIDInt64(c)
	err := edService.DeleteEnvDetail(ctx, ID, projectId)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteEnvDetailByIds 批量删除环境详情
// @Tags EnvDetail
// @Summary 批量删除环境详情
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /ed/deleteEnvDetailByIds [delete]
func (edApi *EnvDetailApi) DeleteEnvDetailByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := edService.DeleteEnvDetailByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateEnvDetail 更新环境详情
// @Tags EnvDetail
// @Summary 更新环境详情
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.EnvDetail true "更新环境详情"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /ed/updateEnvDetail [put]
func (edApi *EnvDetailApi) UpdateEnvDetail(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var ed platform.EnvDetail
	err := c.ShouldBindJSON(&ed)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	projectId := utils.GetProjectIDInt64(c)
	err = edService.UpdateEnvDetail(ctx, ed, projectId)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindEnvDetail 用id查询环境详情
// @Tags EnvDetail
// @Summary 用id查询环境详情
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询环境详情"
// @Success 200 {object} response.Response{data=platform.EnvDetail,msg=string} "查询成功"
// @Router /ed/findEnvDetail [get]
func (edApi *EnvDetailApi) FindEnvDetail(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reed, err := edService.GetEnvDetail(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reed, c)
}

// GetEnvDetailList 分页获取环境详情列表
// @Tags EnvDetail
// @Summary 分页获取环境详情列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query platformReq.EnvDetailSearch true "分页获取环境详情列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /ed/getEnvDetailList [get]
func (edApi *EnvDetailApi) GetEnvDetailList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo platformReq.EnvDetailSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := edService.GetEnvDetailInfoList(ctx, pageInfo)
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

// GetEnvDetailPublic 不需要鉴权的环境详情接口
// @Tags EnvDetail
// @Summary 不需要鉴权的环境详情接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ed/getEnvDetailPublic [get]
func (edApi *EnvDetailApi) GetEnvDetailPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	edService.GetEnvDetailPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的环境详情接口信息",
	}, "获取成功", c)
}
