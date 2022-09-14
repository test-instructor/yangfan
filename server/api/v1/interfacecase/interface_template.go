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
	"strconv"
)

type InterfaceTemplateApi struct {
}

var apicaseServices = service.ServiceGroupApp.InterfacecaseServiceGroup.InterfaceTemplateService

// CreateInterfaceTemplate 创建InterfaceTemplate
// @Tags InterfaceTemplate
// @Summary 创建InterfaceTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body interfacecase.InterfaceTemplate true "创建InterfaceTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apicase/createInterfaceTemplate [post]
func (apiCaseApi *InterfaceTemplateApi) CreateInterfaceTemplate(c *gin.Context) {
	var apicase interfacecase.ApiStep
	_ = c.ShouldBindJSON(&apicase)
	apicase.ProjectID = utils.GetUserProject(c)
	apicase.CreatedByID = utils.GetUserID(c)
	menuStr := c.Query("menu")
	menuInt, _ := strconv.Atoi(menuStr)
	menu := interfacecase.ApiMenu{GVA_MODEL: global.GVA_MODEL{ID: uint(menuInt)}}
	apicase.ApiMenu = menu
	if err := apicaseServices.CreateInterfaceTemplate(apicase); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteInterfaceTemplate 删除InterfaceTemplate
// @Tags InterfaceTemplate
// @Summary 删除InterfaceTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body interfacecase.InterfaceTemplate true "删除InterfaceTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /apicase/deleteInterfaceTemplate [delete]
func (apiCaseApi *InterfaceTemplateApi) DeleteInterfaceTemplate(c *gin.Context) {
	var apicase interfacecase.ApiStep
	_ = c.ShouldBindJSON(&apicase)
	apicase.ProjectID = utils.GetUserProject(c)
	apicase.DeleteByID = utils.GetUserID(c)
	if err := apicaseServices.DeleteInterfaceTemplate(apicase); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteInterfaceTemplateByIds 批量删除InterfaceTemplate
// @Tags InterfaceTemplate
// @Summary 批量删除InterfaceTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除InterfaceTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /apicase/deleteInterfaceTemplateByIds [delete]
func (apiCaseApi *InterfaceTemplateApi) DeleteInterfaceTemplateByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := apicaseServices.DeleteInterfaceTemplateByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateInterfaceTemplate 更新InterfaceTemplate
// @Tags InterfaceTemplate
// @Summary 更新InterfaceTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body interfacecase.InterfaceTemplate true "更新InterfaceTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /apicase/updateInterfaceTemplate [put]
func (apiCaseApi *InterfaceTemplateApi) UpdateInterfaceTemplate(c *gin.Context) {
	var apicase interfacecase.ApiStep
	_ = c.ShouldBindJSON(&apicase)
	apicase.ProjectID = utils.GetUserProject(c)
	apicase.UpdateByID = utils.GetUserID(c)
	menuStr := c.Query("menu")
	menuInt, _ := strconv.Atoi(menuStr)
	menu := interfacecase.ApiMenu{GVA_MODEL: global.GVA_MODEL{ID: uint(menuInt)}}
	apicase.ApiMenu = menu
	if id, err := apicaseServices.UpdateInterfaceTemplate(apicase); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithDetailed(gin.H{"id": id}, "更新成功", c)
	}
}

// FindInterfaceTemplate 用id查询InterfaceTemplate
// @Tags InterfaceTemplate
// @Summary 用id查询InterfaceTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query interfacecase.InterfaceTemplate true "用id查询InterfaceTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /apicase/findInterfaceTemplate [get]
func (apiCaseApi *InterfaceTemplateApi) FindInterfaceTemplate(c *gin.Context) {
	var apicase interfacecase.ApiStep
	_ = c.ShouldBindQuery(&apicase)
	apicase.ProjectID = utils.GetUserProject(c)
	if err, reapicase := apicaseServices.GetInterfaceTemplate(apicase.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reapicase": reapicase}, c)
	}
}

// GetInterfaceTemplateList 分页获取InterfaceTemplate列表
// @Tags InterfaceTemplate
// @Summary 分页获取InterfaceTemplate列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query interfacecaseReq.InterfaceTemplateSearch true "分页获取InterfaceTemplate列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apicase/getInterfaceTemplateList [get]
func (apiCaseApi *InterfaceTemplateApi) GetInterfaceTemplateList(c *gin.Context) {
	var pageInfo interfacecaseReq.InterfaceTemplateSearch
	_ = c.ShouldBindQuery(&pageInfo)
	pageInfo.ProjectID = utils.GetUserProject(c)
	menuid := c.Query("menu")
	menuid_int, _ := strconv.Atoi(menuid)
	pageInfo.ApiMenuID = uint(menuid_int)

	if err, list, total := apicaseServices.GetInterfaceTemplateInfoList(pageInfo); err != nil {
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

func (apiCaseApi *InterfaceTemplateApi) UpdateDebugTalk(c *gin.Context) {
	var debugTalk interfacecase.ApiDebugTalk
	_ = c.ShouldBindJSON(&debugTalk)
	debugTalk.ProjectID = utils.GetUserProject(c)
	if err := apicaseServices.UpdateDebugTalk(debugTalk); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (apiCaseApi *InterfaceTemplateApi) GetDebugTalk(c *gin.Context) {
	var debugTalk interfacecase.ApiDebugTalk
	_ = c.ShouldBindJSON(&debugTalk)
	debugTalk.Project.ID = utils.GetUserProject(c)
	if err, reapicase := apicaseServices.GetDebugTalk(debugTalk); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reapicase": reapicase}, c)
	}
}
