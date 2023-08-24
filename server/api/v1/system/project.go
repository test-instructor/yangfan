package system

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/common/request"
	"github.com/test-instructor/yangfan/server/model/common/response"
	"github.com/test-instructor/yangfan/server/model/system"
	interfacecaseReq "github.com/test-instructor/yangfan/server/model/system/request"
	"github.com/test-instructor/yangfan/server/service"
	"go.uber.org/zap"
)

type ProjectApi struct {
}

var projectService = service.ServiceGroupApp.SystemServiceGroup.ProjectService

// CreateProject 创建Project
// @Tags Project
// @Summary 创建Project
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Project true "创建Project"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /project/createProject [post]
func (projectApi *ProjectApi) CreateProject(c *gin.Context) {
	var project system.Project
	_ = c.ShouldBindJSON(&project)
	if err := projectService.CreateProject(project); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteProject 删除Project
// @Tags Project
// @Summary 删除Project
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Project true "删除Project"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /project/deleteProject [delete]
func (projectApi *ProjectApi) DeleteProject(c *gin.Context) {
	var project system.Project
	_ = c.ShouldBindJSON(&project)
	if err := projectService.DeleteProject(project); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteProjectByIds 批量删除Project
// @Tags Project
// @Summary 批量删除Project
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Project"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /project/deleteProjectByIds [delete]
func (projectApi *ProjectApi) DeleteProjectByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := projectService.DeleteProjectByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateProject 更新Project
// @Tags Project
// @Summary 更新Project
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Project true "更新Project"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /project/updateProject [put]
func (projectApi *ProjectApi) UpdateProject(c *gin.Context) {
	var project system.Project
	_ = c.ShouldBindJSON(&project)
	if err := projectService.UpdateProject(project); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (projectApi *ProjectApi) SetUserProjectAuth(c *gin.Context) {
	var sua system.SysUserProject
	_ = c.ShouldBindJSON(&sua)
	if err := projectService.SetUserProjectAuth(sua); err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

func (projectApi *ProjectApi) SetKey(c *gin.Context) {
	var sp system.Project
	_ = c.ShouldBindJSON(&sp)
	if err, data := projectService.SetKey(sp); err != nil {
		global.GVA_LOG.Error("设置密钥失败!", zap.Error(err))
		response.FailWithMessage("设置密钥失败", c)
	} else {
		response.OkWithData(data, c)
	}
}

func (projectApi *ProjectApi) FindKey(c *gin.Context) {
	var sp system.Project
	_ = c.ShouldBindQuery(&sp)
	if err, data := projectService.FindKey(sp); err != nil {
		global.GVA_LOG.Error("设置密钥失败!", zap.Error(err))
		response.FailWithMessage("设置密钥失败", c)
	} else {
		response.OkWithData(data, c)
	}
}

func (projectApi *ProjectApi) DeleteProjectAuth(c *gin.Context) {
	var sua system.SysUserProject
	_ = c.ShouldBindJSON(&sua)
	if err := projectService.DeleteProjectAuth(sua); err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// FindProject 用id查询Project
// @Tags Project
// @Summary 用id查询Project
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query system.Project true "用id查询Project"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /project/findProject [get]
func (projectApi *ProjectApi) FindProject(c *gin.Context) {
	var project system.Project
	_ = c.ShouldBindQuery(&project)
	if err, reproject := projectService.GetProject(project.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reproject": reproject}, c)
	}
}

// GetProjectList 分页获取Project列表
// @Tags Project
// @Summary 分页获取Project列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query interfacecaseReq.ProjectSearch true "分页获取Project列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /project/getProjectList [get]
func (projectApi *ProjectApi) GetProjectList(c *gin.Context) {
	var pageInfo interfacecaseReq.ProjectSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := projectService.GetProjectInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (projectApi *ProjectApi) GetProjectUserList(c *gin.Context) {
	var search interfacecaseReq.SysProjectUsers
	_ = c.ShouldBindQuery(&search)
	list, total, err := projectService.GetProjectUserList(search)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     search.Page,
		PageSize: search.PageSize,
	}, "获取成功", c)
}
