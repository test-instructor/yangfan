package projectmgr

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/common/response"
	"github.com/test-instructor/yangfan/server/v2/model/projectmgr"
	projectmgrReq "github.com/test-instructor/yangfan/server/v2/model/projectmgr/request"
	"go.uber.org/zap"
)

type ProjectApi struct{}

// CreateProject 创建项目配置
// @Tags Project
// @Summary 创建项目配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body projectmgr.Project true "创建项目配置"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /pj/createProject [post]
func (pjApi *ProjectApi) CreateProject(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pj projectmgr.Project
	err := c.ShouldBindJSON(&pj)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = pjService.CreateProject(c, ctx, &pj)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteProject 删除项目配置
// @Tags Project
// @Summary 删除项目配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body projectmgr.Project true "删除项目配置"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /pj/deleteProject [delete]
func (pjApi *ProjectApi) DeleteProject(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := pjService.DeleteProject(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteProjectByIds 批量删除项目配置
// @Tags Project
// @Summary 批量删除项目配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /pj/deleteProjectByIds [delete]
func (pjApi *ProjectApi) DeleteProjectByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := pjService.DeleteProjectByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateProject 更新项目配置
// @Tags Project
// @Summary 更新项目配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body projectmgr.Project true "更新项目配置"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /pj/updateProject [put]
func (pjApi *ProjectApi) UpdateProject(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var pj projectmgr.Project
	err := c.ShouldBindJSON(&pj)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = pjService.UpdateProject(ctx, pj)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// ResetProjectAuth 重设项目CI鉴权信息
// @Tags Project
// @Summary 重设项目CI鉴权信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body projectmgrReq.ResetProjectAuthReq true "重设项目CI鉴权信息"
// @Success 200 {object} response.Response{data=object,msg=string} "重设成功"
// @Router /pj/resetProjectAuth [put]
func (pjApi *ProjectApi) ResetProjectAuth(c *gin.Context) {
	ctx := c.Request.Context()

	var req projectmgrReq.ResetProjectAuthReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	pj, err := pjService.ResetProjectAuth(ctx, req.ID)
	if err != nil {
		global.GVA_LOG.Error("重设失败!", zap.Error(err))
		response.FailWithMessage("重设失败:"+err.Error(), c)
		return
	}

	response.OkWithData(gin.H{
		"uuid":   pj.UUID,
		"secret": pj.Secret,
	}, c)
}

// FindProject 用id查询项目配置
// @Tags Project
// @Summary 用id查询项目配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询项目配置"
// @Success 200 {object} response.Response{data=projectmgr.Project,msg=string} "查询成功"
// @Router /pj/findProject [get]
func (pjApi *ProjectApi) FindProject(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	repj, err := pjService.GetProject(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(repj, c)
}

// GetProjectList 分页获取项目配置列表
// @Tags Project
// @Summary 分页获取项目配置列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query projectmgrReq.ProjectSearch true "分页获取项目配置列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /pj/getProjectList [get]
func (pjApi *ProjectApi) GetProjectList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo projectmgrReq.ProjectSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := pjService.GetProjectInfoList(ctx, pageInfo)
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

// GetProjectPublic 不需要鉴权的项目配置接口
// @Tags Project
// @Summary 不需要鉴权的项目配置接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /pj/getProjectPublic [get]
func (pjApi *ProjectApi) GetProjectPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	pjService.GetProjectPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的项目配置接口信息",
	}, "获取成功", c)
}
