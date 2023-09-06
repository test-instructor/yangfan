package interfacecase

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/common/response"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	interfacecaseReq "github.com/test-instructor/yangfan/server/model/interfacecase/request"
	"github.com/test-instructor/yangfan/server/service"
	"github.com/test-instructor/yangfan/server/utils"
)

type EnvironmentAPi struct{}

var environmentService = service.ServiceGroupApp.InterfacecaseServiceGroup.EnvironmentService

func (env *EnvironmentAPi) CreateEnv(c *gin.Context) {
	var environment interfacecase.ApiEnv
	_ = c.ShouldBindJSON(&environment)
	environment.ProjectID = utils.GetUserProject(c)
	environment.CreatedBy = utils.GetUserIDAddress(c)
	if envs, err := environmentService.CreateEnv(environment); err != nil {
		global.GVA_LOG.Error("环境操作失败!", zap.Error(err))
		response.FailWithMessage("环境操作失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: envs,
		}, "获取成功", c)
	}

}
func (env *EnvironmentAPi) DeleteEnv(c *gin.Context) {
	var environment interfacecase.ApiEnv
	_ = c.ShouldBindJSON(&environment)
	environment.DeleteBy = utils.GetUserIDAddress(c)
	if err := environmentService.DeleteEnv(environment); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}

}
func (env *EnvironmentAPi) DeleteEnvByIds(c *gin.Context) {
}

func (env *EnvironmentAPi) UpdateEnv(c *gin.Context) {
	var environment interfacecase.ApiEnv
	_ = c.ShouldBindJSON(&environment)
	environment.ProjectID = utils.GetUserProject(c)
	environment.UpdateBy = utils.GetUserIDAddress(c)
	if id, err := environmentService.UpdateEnv(environment); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithDetailed(gin.H{"id": id}, "更新成功", c)
	}
}
func (env *EnvironmentAPi) FindEnv(c *gin.Context) {
	var environment interfacecase.ApiEnv
	_ = c.ShouldBindJSON(&environment)
	if err, env := environmentService.FindEnv(environment.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"env": env}, c)
	}
}

func (env *EnvironmentAPi) GetEnvList(c *gin.Context) {
	var pageInfo interfacecaseReq.EnvSearch
	_ = c.ShouldBindQuery(&pageInfo)
	pageInfo.ProjectID = utils.GetUserProject(c)
	if err, list := environmentService.GetEnvList(pageInfo); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: list,
		}, "获取成功", c)
	}
}

func (env *EnvironmentAPi) CreateEnvVariable(c *gin.Context) {
	var envVar interfacecase.ApiEnvDetail
	_ = c.ShouldBindJSON(&envVar)
	envVar.ProjectID = utils.GetUserProject(c)
	if err := environmentService.CreateEnvVariable(envVar); err != nil {
		global.GVA_LOG.Error("操作失败!", zap.Error(err))
		response.FailWithMessage("操作失败", c)
	} else {
		response.OkWithMessage("操作成功", c)
	}
}

func (env *EnvironmentAPi) DeleteEnvVariable(c *gin.Context) {
	var environment interfacecase.ApiEnvDetail
	_ = c.ShouldBindJSON(&environment)
	environment.DeleteBy = utils.GetUserIDAddress(c)
	if err := environmentService.DeleteEnvVariable(environment); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (env *EnvironmentAPi) FindEnvVariable(c *gin.Context) {
	var environment interfacecase.ApiEnvDetail
	_ = c.ShouldBindQuery(&environment)
	if err, env := environmentService.FindEnvVariable(environment.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"env": env}, c)
	}
}

func (env *EnvironmentAPi) GetEnvVariableList(c *gin.Context) {
	var pageInfo interfacecaseReq.EnvVariableSearch
	_ = c.ShouldBindQuery(&pageInfo)
	pageInfo.ProjectID = utils.GetUserProject(c)
	if err, list, total := environmentService.GetEnvVariableList(pageInfo); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (env *EnvironmentAPi) CreateEnvMock(c *gin.Context) {
	var envVar interfacecase.ApiEnvMock
	_ = c.ShouldBindJSON(&envVar)
	envVar.ProjectID = utils.GetUserProject(c)
	if err := environmentService.CreateEnvMock(envVar); err != nil {
		global.GVA_LOG.Error("操作失败!", zap.Error(err))
		response.FailWithMessage("操作失败", c)
	} else {
		response.OkWithMessage("操作成功", c)
	}
}

func (env *EnvironmentAPi) DeleteEnvMock(c *gin.Context) {
	var environment interfacecase.ApiEnvMock
	_ = c.ShouldBindJSON(&environment)
	environment.DeleteBy = utils.GetUserIDAddress(c)
	if err := environmentService.DeleteEnvMock(environment); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (env *EnvironmentAPi) FindEnvMock(c *gin.Context) {
	var environment interfacecase.ApiEnvMock
	_ = c.ShouldBindQuery(&environment)
	if err, env := environmentService.FindEnvMock(environment.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"env": env}, c)
	}
}

func (env *EnvironmentAPi) GetEnvMockList(c *gin.Context) {
	var pageInfo interfacecaseReq.EnvMockSearch
	_ = c.ShouldBindQuery(&pageInfo)
	pageInfo.ProjectID = utils.GetUserProject(c)
	if err, list, total := environmentService.GetEnvMockList(pageInfo); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
