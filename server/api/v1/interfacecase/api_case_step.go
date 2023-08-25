package interfacecase

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/common/request"
	"github.com/test-instructor/yangfan/server/model/common/response"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	interfacecaseReq "github.com/test-instructor/yangfan/server/model/interfacecase/request"
	"github.com/test-instructor/yangfan/server/service"
	"github.com/test-instructor/yangfan/server/utils"
)

type ApiCase struct {
}

var testCaseService = service.ServiceGroupApp.InterfacecaseServiceGroup.TestCaseService

// CreateTestCaseStep 创建TestCase
//	@Tags		TestCase
//	@Summary	创建TestCase
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		interfacecase.ApiCaseStep	true	"创建TestCase"
//	@Success	200		{string}	string						"{"success":true,"data":{},"msg":"获取成功"}"
//	@Router		/apicase/createTestCase [post]
func (apiCase *ApiCase) CreateTestCaseStep(c *gin.Context) {
	var apicase interfacecase.ApiCaseStep
	var menu interfacecase.ApiMenu
	_ = c.ShouldBindJSON(&apicase)
	apicase.ProjectID = utils.GetUserProject(c)
	apicase.CreatedBy = utils.GetUserIDAddress(c)
	menuStr := c.Query("menu")
	menuInt, _ := strconv.Atoi(menuStr)
	menu = interfacecase.ApiMenu{GVA_MODEL: global.GVA_MODEL{ID: uint(menuInt)}}
	apicase.ApiMenu = menu
	if err := testCaseService.CreateTestCaseStep(apicase); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTestCaseStep 删除TestCase
//	@Tags		TestCase
//	@Summary	删除TestCase
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		interfacecase.ApiCaseStep	true	"删除TestCase"
//	@Success	200		{string}	string						"{"success":true,"data":{},"msg":"删除成功"}"
//	@Router		/apicase/deleteTestCase [delete]
func (apiCase *ApiCase) DeleteTestCaseStep(c *gin.Context) {
	var apicase interfacecase.ApiCaseStep
	_ = c.ShouldBindJSON(&apicase)
	apicase.ProjectID = utils.GetUserProject(c)
	apicase.DeleteBy = utils.GetUserIDAddress(c)
	if err := testCaseService.DeleteTestCaseStep(apicase); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTestCaseStepByIds 批量删除TestCase
//	@Tags		TestCase
//	@Summary	批量删除TestCase
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		request.IdsReq	true	"批量删除TestCase"
//	@Success	200		{string}	string			"{"success":true,"data":{},"msg":"批量删除成功"}"
//	@Router		/apicase/deleteTestCaseByIds [delete]
func (apiCase *ApiCase) DeleteTestCaseStepByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := testCaseService.DeleteTestCaseStepByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTestCaseStep 更新TestCase
//	@Tags		TestCase
//	@Summary	更新TestCase
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		interfacecase.ApiCaseStep	true	"更新TestCase"
//	@Success	200		{string}	string						"{"success":true,"data":{},"msg":"更新成功"}"
//	@Router		/apicase/updateTestCase [put]
func (apiCase *ApiCase) UpdateTestCaseStep(c *gin.Context) {
	var apicase interfacecase.ApiCaseStep
	_ = c.ShouldBindJSON(&apicase)
	apicase.ProjectID = utils.GetUserProject(c)
	apicase.UpdateBy = utils.GetUserIDAddress(c)
	menuStr := c.Query("menu")
	menuInt, _ := strconv.Atoi(menuStr)
	menu := interfacecase.ApiMenu{GVA_MODEL: global.GVA_MODEL{ID: uint(menuInt)}}
	apicase.ApiMenu = menu
	if err := testCaseService.UpdateTestCaseStep(apicase); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// SortTestCaseStep 用例列表排序
//	@Tags		ApiConfig
//	@Summary	更新用例列表排序
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		interfacecase.ApiConfig	true	"更新ApiConfig"
//	@Success	200		{string}	string					"{"success":true,"data":{},"msg":"更新成功"}"
//	@Router		/ac/updateApiConfig [put]
func (apiCase *ApiCase) SortTestCaseStep(c *gin.Context) {
	var apicase interfacecase.ApiCaseStep
	_ = c.ShouldBindJSON(&apicase)
	apicase.ProjectID = utils.GetUserProject(c)
	if err := testCaseService.SortTestCaseStep(apicase); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// AddTestCaseStep 用例列表排序
//	@Tags		ApiConfig
//	@Summary	更新用例列表排序
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		interfacecase.ApiConfig	true	"更新ApiConfig"
//	@Success	200		{string}	string					"{"success":true,"data":{},"msg":"更新成功"}"
//	@Router		/ac/updateApiConfig [put]
func (apiCase *ApiCase) AddTestCaseStep(c *gin.Context) {
	var apiCaseID request.ApiCaseIdReq
	_ = c.ShouldBindJSON(&apiCaseID)
	caseApiDetail, err := testCaseService.AddTestCaseStep(apiCaseID)
	if err != nil {
		global.GVA_LOG.Error("添加用例失败!", zap.Error(err))
		response.FailWithMessage("添加用例失败", c)
	} else {
		response.OkWithDetailed(caseApiDetail, "添加用例成功", c)
	}
}

// DelTestCaseStep 用例列表排序
//	@Tags		ApiConfig
//	@Summary	更新用例列表排序
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		interfacecase.ApiConfig	true	"更新ApiConfig"
//	@Success	200		{string}	string					"{"success":true,"data":{},"msg":"更新成功"}"
//	@Router		/ac/updateApiConfig [put]
func (apiCase *ApiCase) DelTestCaseStep(c *gin.Context) {
	var apiCaseID request.ApiCaseIdReq
	_ = c.ShouldBindJSON(&apiCaseID)
	err := testCaseService.DelTestCaseStep(apiCaseID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// FindTestCaseStep 用id查询TestCase
//	@Tags		TestCase
//	@Summary	用id查询TestCase
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	query		interfacecase.ApiCaseStep	true	"用id查询TestCase"
//	@Success	200		{string}	string						"{"success":true,"data":{},"msg":"查询成功"}"
//	@Router		/apicase/findTestCase [get]
func (apiCase *ApiCase) FindTestCaseStep(c *gin.Context) {
	var apicase request.ApiCaseIdReq
	_ = c.ShouldBindQuery(&apicase)
	if err, reapicase := testCaseService.FindTestCaseStep(apicase.ID, apicase.Detail); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reapicase": reapicase}, c)
	}
}

// GetTestCaseStepList 分页获取TestCase列表
//	@Tags		TestCase
//	@Summary	分页获取TestCase列表
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	query		interfacecaseReq.TestCaseSearch	true	"分页获取TestCase列表"
//	@Success	200		{string}	string							"{"success":true,"data":{},"msg":"获取成功"}"
//	@Router		/ApiCase/getTestCaseList [get]
func (apiCase *ApiCase) GetTestCaseStepList(c *gin.Context) {
	var pageInfo interfacecaseReq.TestCaseSearch
	_ = c.ShouldBindQuery(&pageInfo)
	pageInfo.ProjectID = utils.GetUserProject(c)
	menuId := c.Query("menu")
	menuidInt, _ := strconv.Atoi(menuId)
	pageInfo.ApiMenuID = uint(menuidInt)
	if err, list, total := testCaseService.GetTestCaseStepInfoList(pageInfo); err != nil {
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
