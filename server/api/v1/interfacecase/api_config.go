package interfacecase

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/model/common/request"
	"github.com/test-instructor/cheetah/server/model/common/response"
	"github.com/test-instructor/cheetah/server/model/interfacecase"
	interfacecaseReq "github.com/test-instructor/cheetah/server/model/interfacecase/request"
	"github.com/test-instructor/cheetah/server/service"
	"github.com/test-instructor/cheetah/server/utils"
	"go.uber.org/zap"
)

type ApiConfigApi struct {
}

var acService = service.ServiceGroupApp.InterfacecaseServiceGroup.ApiConfigService

// CreateApiConfig 创建ApiConfig
// @Tags ApiConfig
// @Summary 创建ApiConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body interfacecase.ApiConfig true "创建ApiConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ac/createApiConfig [post]
func (acApi *ApiConfigApi) CreateApiConfig(c *gin.Context) {
	var ac interfacecase.ApiConfig
	err := c.ShouldBindJSON(&ac)
	if err != nil {
		global.GVA_LOG.Error("获取创建配置参数失败", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	}
	ac.CreatedByID = utils.GetUserIDAddress(c)
	ac.ProjectID = utils.GetUserProject(c)
	if err := acService.CreateApiConfig(ac); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}

}

// DeleteApiConfig 删除ApiConfig
// @Tags ApiConfig
// @Summary 删除ApiConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body interfacecase.ApiConfig true "删除ApiConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ac/deleteApiConfig [delete]
func (acApi *ApiConfigApi) DeleteApiConfig(c *gin.Context) {
	var ac interfacecase.ApiConfig
	_ = c.ShouldBindJSON(&ac)
	ac.ProjectID = utils.GetUserProject(c)
	ac.DeleteByID = utils.GetUserIDAddress(c)
	if err := acService.DeleteApiConfig(ac); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteApiConfigByIds 批量删除ApiConfig
// @Tags ApiConfig
// @Summary 批量删除ApiConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ApiConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ac/deleteApiConfigByIds [delete]
func (acApi *ApiConfigApi) DeleteApiConfigByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := acService.DeleteApiConfigByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateApiConfig 更新ApiConfig
// @Tags ApiConfig
// @Summary 更新ApiConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body interfacecase.ApiConfig true "更新ApiConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ac/updateApiConfig [put]
func (acApi *ApiConfigApi) UpdateApiConfig(c *gin.Context) {
	var ac interfacecase.ApiConfig
	_ = c.ShouldBindJSON(&ac)
	ac.ProjectID = utils.GetUserProject(c)
	ac.UpdateByID = utils.GetUserIDAddress(c)
	if ac.SetupCaseID == nil {
		ac.SetupCaseID = nil
	}
	if err := acService.UpdateApiConfig(ac); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindApiConfig 用id查询ApiConfig
// @Tags ApiConfig
// @Summary 用id查询ApiConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query interfacecase.ApiConfig true "用id查询ApiConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ac/findApiConfig [get]
func (acApi *ApiConfigApi) FindApiConfig(c *gin.Context) {
	var ac interfacecase.ApiConfig
	_ = c.ShouldBindQuery(&ac)
	ac.ProjectID = utils.GetUserProject(c)
	if err, reac := acService.GetApiConfig(ac.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reac": reac}, c)
	}
}

// GetApiConfigList 分页获取ApiConfig列表
// @Tags ApiConfig
// @Summary 分页获取ApiConfig列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query interfacecaseReq.ApiConfigSearch true "分页获取ApiConfig列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ac/getApiConfigList [get]
func (acApi *ApiConfigApi) GetApiConfigList(c *gin.Context) {
	var pageInfo interfacecaseReq.ApiConfigSearch
	_ = c.ShouldBindQuery(&pageInfo)
	pageInfo.ProjectID = utils.GetUserProject(c)
	if err, list, total := acService.GetApiConfigInfoList(pageInfo); err != nil {
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
