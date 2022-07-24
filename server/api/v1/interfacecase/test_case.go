package interfacecase

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/model/common/request"
	"github.com/test-instructor/cheetah/server/model/common/response"
	"github.com/test-instructor/cheetah/server/model/interfacecase"
	interfacecaseReq "github.com/test-instructor/cheetah/server/model/interfacecase/request"
	"github.com/test-instructor/cheetah/server/model/system"
	"github.com/test-instructor/cheetah/server/service"
	"go.uber.org/zap"
	"strconv"
)

type TestCaseApi struct {
}

var testCaseService = service.ServiceGroupApp.InterfacecaseServiceGroup.TestCaseService

// CreateTestCase 创建TestCase
// @Tags TestCase
// @Summary 创建TestCase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body interfacecase.TestCase true "创建TestCase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apicase/createTestCase [post]
func (apicaseApi *TestCaseApi) CreateTestCase(c *gin.Context) {
	var apicase interfacecase.ApiTestCase
	var menu interfacecase.ApiMenu
	_ = c.ShouldBindJSON(&apicase)
	projectsss, _ := c.Get("project")
	apicase.Project = projectsss.(system.Project)
	menuStr := c.Query("menu")
	menuInt, _ := strconv.Atoi(menuStr)
	menu = interfacecase.ApiMenu{GVA_MODEL: global.GVA_MODEL{ID: uint(menuInt)}}
	apicase.ApiMenu = menu
	if err := testCaseService.CreateTestCase(apicase); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTestCase 删除TestCase
// @Tags TestCase
// @Summary 删除TestCase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body interfacecase.TestCase true "删除TestCase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /apicase/deleteTestCase [delete]
func (apicaseApi *TestCaseApi) DeleteTestCase(c *gin.Context) {
	var apicase interfacecase.ApiTestCase
	_ = c.ShouldBindJSON(&apicase)
	projectsss, _ := c.Get("project")
	apicase.Project = projectsss.(system.Project)
	if err := testCaseService.DeleteTestCase(apicase); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTestCaseByIds 批量删除TestCase
// @Tags TestCase
// @Summary 批量删除TestCase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TestCase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /apicase/deleteTestCaseByIds [delete]
func (apicaseApi *TestCaseApi) DeleteTestCaseByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := testCaseService.DeleteTestCaseByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTestCase 更新TestCase
// @Tags TestCase
// @Summary 更新TestCase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body interfacecase.TestCase true "更新TestCase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /apicase/updateTestCase [put]
func (apicaseApi *TestCaseApi) UpdateTestCase(c *gin.Context) {
	var apicase interfacecase.ApiTestCase
	_ = c.ShouldBindJSON(&apicase)
	projectsss, _ := c.Get("project")
	apicase.Project = projectsss.(system.Project)
	menuStr := c.Query("menu")
	menuInt, _ := strconv.Atoi(menuStr)
	menu := interfacecase.ApiMenu{GVA_MODEL: global.GVA_MODEL{ID: uint(menuInt)}}
	apicase.ApiMenu = menu
	if err := testCaseService.UpdateTestCase(apicase); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// SortTestCase 用例列表排序
// @Tags ApiConfig
// @Summary 更新用例列表排序
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body interfacecase.ApiConfig true "更新ApiConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ac/updateApiConfig [put]
func (acApi *TestCaseApi) SortTestCase(c *gin.Context) {
	var apicase interfacecase.ApiTestCase
	_ = c.ShouldBindJSON(&apicase)
	projectsss, _ := c.Get("project")
	apicase.Project = projectsss.(system.Project)
	if err := testCaseService.SortTestCase(apicase); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// AddTestCase 用例列表排序
// @Tags ApiConfig
// @Summary 更新用例列表排序
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body interfacecase.ApiConfig true "更新ApiConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ac/updateApiConfig [put]
func (acApi *TestCaseApi) AddTestCase(c *gin.Context) {
	var apiCaseID request.ApiCaseIdReq
	_ = c.ShouldBindJSON(&apiCaseID)
	caseApiDetail, err := testCaseService.AddTestCase(apiCaseID)
	if err != nil {
		global.GVA_LOG.Error("添加用例失败!", zap.Error(err))
		response.FailWithMessage("添加用例失败", c)
	} else {
		response.OkWithDetailed(caseApiDetail, "添加用例成功", c)
	}
}

// DelTestCase 用例列表排序
// @Tags ApiConfig
// @Summary 更新用例列表排序
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body interfacecase.ApiConfig true "更新ApiConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ac/updateApiConfig [put]
func (acApi *TestCaseApi) DelTestCase(c *gin.Context) {
	var apiCaseID request.ApiCaseIdReq
	_ = c.ShouldBindJSON(&apiCaseID)
	err := testCaseService.DelTestCase(apiCaseID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// FindTestCase 用id查询TestCase
// @Tags TestCase
// @Summary 用id查询TestCase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query interfacecase.TestCase true "用id查询TestCase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /apicase/findTestCase [get]
func (apicaseApi *TestCaseApi) FindTestCase(c *gin.Context) {
	var apicase interfacecase.ApiTestCase
	_ = c.ShouldBindQuery(&apicase)
	projectsss, _ := c.Get("project")
	apicase.Project = projectsss.(system.Project)
	if err, reapicase := testCaseService.FindTestCase(apicase.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reapicase": reapicase}, c)
	}
}

// GetTestCaseList 分页获取TestCase列表
// @Tags TestCase
// @Summary 分页获取TestCase列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query interfacecaseReq.TestCaseSearch true "分页获取TestCase列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apicase/getTestCaseList [get]
func (apicaseApi *TestCaseApi) GetTestCaseList(c *gin.Context) {
	var pageInfo interfacecaseReq.TestCaseSearch
	_ = c.ShouldBindQuery(&pageInfo)
	projectsss, _ := c.Get("project")
	pageInfo.Project = projectsss.(system.Project)

	menuId := c.Query("menu")
	menuidInt, _ := strconv.Atoi(menuId)
	pageInfo.ApiMenuID = uint(menuidInt)
	if err, list, total := testCaseService.GetTestCaseInfoList(pageInfo); err != nil {
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
