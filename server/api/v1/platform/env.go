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

type EnvApi struct{}

// CreateEnv 创建环境配置
// @Tags Env
// @Summary 创建环境配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.Env true "创建环境配置"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /env/createEnv [post]
func (envApi *EnvApi) CreateEnv(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var env platform.Env
	err := c.ShouldBindJSON(&env)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = envService.CreateEnv(ctx, &env)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteEnv 删除环境配置
// @Tags Env
// @Summary 删除环境配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.Env true "删除环境配置"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /env/deleteEnv [delete]
func (envApi *EnvApi) DeleteEnv(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	projectId := utils.GetProjectIDInt64(c)
	err := envService.DeleteEnv(ctx, ID, projectId)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteEnvByIds 批量删除环境配置
// @Tags Env
// @Summary 批量删除环境配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /env/deleteEnvByIds [delete]
func (envApi *EnvApi) DeleteEnvByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := envService.DeleteEnvByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateEnv 更新环境配置
// @Tags Env
// @Summary 更新环境配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.Env true "更新环境配置"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /env/updateEnv [put]
func (envApi *EnvApi) UpdateEnv(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var env platform.Env
	err := c.ShouldBindJSON(&env)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	projectId := utils.GetProjectIDInt64(c)
	err = envService.UpdateEnv(ctx, env, projectId)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindEnv 用id查询环境配置
// @Tags Env
// @Summary 用id查询环境配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询环境配置"
// @Success 200 {object} response.Response{data=platform.Env,msg=string} "查询成功"
// @Router /env/findEnv [get]
func (envApi *EnvApi) FindEnv(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reenv, err := envService.GetEnv(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reenv, c)
}

// GetEnvList 分页获取环境配置列表
// @Tags Env
// @Summary 分页获取环境配置列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query platformReq.EnvSearch true "分页获取环境配置列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /env/getEnvList [get]
func (envApi *EnvApi) GetEnvList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo platformReq.EnvSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := envService.GetEnvInfoList(ctx, pageInfo)
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

// GetEnvPublic 不需要鉴权的环境配置接口
// @Tags Env
// @Summary 不需要鉴权的环境配置接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /env/getEnvPublic [get]
func (envApi *EnvApi) GetEnvPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	envService.GetEnvPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的环境配置接口信息",
	}, "获取成功", c)
}
