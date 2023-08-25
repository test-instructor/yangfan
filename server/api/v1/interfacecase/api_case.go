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

type ApiCaseApi struct {
}

var apiCaseService = service.ServiceGroupApp.InterfacecaseServiceGroup.ApiCaseService

// CreateApiCase 创建ApiCase
//
//	@Tags		ApiCase
//	@Summary	创建ApiCase
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		interfacecase.ApiCase	true	"创建ApiCase"
//	@Success	200		{string}	string					"{"success":true,"data":{},"msg":"获取成功"}"
//	@Router		/testCase/createApiCase [post]
func (apiCase *ApiCaseApi) CreateApiCase(c *gin.Context) {
	var testCase interfacecase.ApiCase
	_ = c.ShouldBindJSON(&testCase)
	testCase.ProjectID = utils.GetUserProject(c)
	menuStr := c.Query("menu")
	menuInt, _ := strconv.Atoi(menuStr)
	menu := interfacecase.ApiMenu{GVA_MODEL: global.GVA_MODEL{ID: uint(menuInt)}}
	testCase.ApiMenu = &menu
	if err := apiCaseService.CreateApiCase(testCase); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteApiCase 删除ApiCase
//
//	@Tags		ApiCase
//	@Summary	删除ApiCase
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		interfacecase.ApiCase	true	"删除ApiCase"
//	@Success	200		{string}	string					"{"success":true,"data":{},"msg":"删除成功"}"
//	@Router		/testCase/deleteApiCase [delete]
func (apiCase *ApiCaseApi) DeleteApiCase(c *gin.Context) {
	var testCase interfacecase.ApiCase
	_ = c.ShouldBindJSON(&testCase)
	testCase.ProjectID = utils.GetUserProject(c)
	testCase.DeleteBy = utils.GetUserIDAddress(c)
	if err := apiCaseService.DeleteApiCase(testCase); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteApiCaseByIds 批量删除ApiCase
//
//	@Tags		ApiCase
//	@Summary	批量删除ApiCase
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		request.IdsReq	true	"批量删除ApiCase"
//	@Success	200		{string}	string			"{"success":true,"data":{},"msg":"批量删除成功"}"
//	@Router		/testCase/deleteApiCaseByIds [delete]
func (apiCase *ApiCaseApi) DeleteApiCaseByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := apiCaseService.DeleteApiCaseByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateApiCase 更新ApiCase
//
//	@Tags		ApiCase
//	@Summary	更新ApiCase
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		interfacecase.ApiCase	true	"更新ApiCase"
//	@Success	200		{string}	string					"{"success":true,"data":{},"msg":"更新成功"}"
//	@Router		/testCase/updateApiCase [put]
func (apiCase *ApiCaseApi) UpdateApiCase(c *gin.Context) {
	var testCase interfacecase.ApiCase
	_ = c.ShouldBindJSON(&testCase)
	testCase.ProjectID = utils.GetUserProject(c)
	testCase.UpdateBy = utils.GetUserIDAddress(c)
	menuStr := c.Query("menu")
	menuInt, _ := strconv.Atoi(menuStr)
	menu := interfacecase.ApiMenu{GVA_MODEL: global.GVA_MODEL{ID: uint(menuInt)}}
	testCase.ApiMenu = &menu
	if err := apiCaseService.UpdateApiCase(testCase); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// SortApisCase 测试用例排序
//
//	@Tags		ApiCase
//	@Summary	测试用例排序
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		[]interfacecase.ApiCaseRelationshipf	true	"测试用例排序"
//	@Success	200		{string}	string									"{"success":true,"data":{},"msg":"更新成功"}"
//	@Router		/testCase/addApiTestCase [delete]
func (apiCase *ApiCaseApi) SortApisCase(c *gin.Context) {
	var testCase []interfacecase.ApiCaseRelationship
	_ = c.ShouldBindJSON(&testCase)
	if err := apiCaseService.SortApisCase(testCase); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

type addApisCaseReq struct {
	TestCaseID uint   `json:"testCase_id"` // 测试用例id
	CaseID     []uint `json:"case_id"`     // 测试步骤id
}

// AddApisCase 添加测试用例
//
//	@Tags		ApiCase
//	@Summary	删除测试步骤
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		addApisCaseReq	true	"删除测试用例"
//	@Success	200		{string}	string			"{"success":true,"data":{},"msg":"更新成功"}"
//	@Router		/testCase/addApiTestCase [delete]
func (apiCase *ApiCaseApi) AddApisCase(c *gin.Context) {
	var testCase addApisCaseReq
	_ = c.ShouldBindJSON(&testCase)
	if err := apiCaseService.AddApisCase(testCase.TestCaseID, testCase.CaseID); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// DelApisCase 删除测试用例
//
//	@Tags		ApiCase
//	@Summary	删除测试步骤
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		interfacecase.ApiCaseRelationship	true	"删除测试用例"
//	@Success	200		{string}	string								"{"success":true,"data":{},"msg":"删除成功"}"
//	@Router		/testCase/addApiTestCase [delete]
func (apiCase *ApiCaseApi) DelApisCase(c *gin.Context) {
	var testCase interfacecase.ApiCaseRelationship
	_ = c.ShouldBindJSON(&testCase)
	if err := apiCaseService.DelApisCase(testCase); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

type apiCaseResp struct {
	Name     string        `json:"name"`
	TestCase []apiTestCase `json:"test_case"`
}

type apiTestCase struct {
	ID   uint                      `json:"id"`
	Case interfacecase.ApiCaseStep `json:"case"`
}

// FindApiTestCase 测试步骤增加用例
//
//	@Tags		ApiCase
//	@Summary	测试步骤增加用例
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		interfacecase.ApiCase	true	"测试步骤增加用例"
//	@Success	200		{string}	string					"{"success":true,"data":{},"msg":"添加用例成功"}"
//	@Router		/testCase/addApiTestCase [post]
func (apiCase *ApiCaseApi) FindApiTestCase(c *gin.Context) {
	var testCase interfacecase.ApiCase
	_ = c.ShouldBindQuery(&testCase)
	testCase.ProjectID = utils.GetUserProject(c)
	var reapicase apiCaseResp
	err, resp, name := apiCaseService.FindApiTestCase(testCase.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		reapicase.Name = name
		for _, v := range resp {
			var testcase apiTestCase
			testcase.ID = v.ID
			v.ApiCaseStep.CreatedBy = nil
			v.ApiCaseStep.UpdateBy = nil
			v.ApiCaseStep.DeleteBy = nil
			v.ApiCaseStep.FrontCase = nil
			v.ApiCaseStep.ProjectID = 0
			v.ApiCaseStep.RunConfigID = 0
			v.ApiCaseStep.ApiEnvID = 0
			v.ApiCaseStep.RunConfigName = nil
			v.ApiCaseStep.ApiEnvName = nil
			testcase.Case = v.ApiCaseStep
			reapicase.TestCase = append(reapicase.TestCase, testcase)
		}
		response.OkWithData(gin.H{"reapicase": reapicase}, c)
	}
}

// AddApiTestCase 测试用例添加步骤
//
//	@Tags		ApiCase
//	@Summary	测试用例添加步骤
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		request.ApiCaseIdReq	true	"测试步骤增加用例"
//	@Success	200		{string}	string					"{"success":true,"data":{},"msg":"添加用例成功"}"
//	@Router		/testCase/addApiTestCase [post]
func (apiCase *ApiCaseApi) AddApiTestCase(c *gin.Context) {
	var apiCaseID request.ApiCaseIdReq
	_ = c.ShouldBindJSON(&apiCaseID)
	caseApiDetail, err := apiCaseService.AddApiTestCase(apiCaseID)
	if err != nil {
		global.GVA_LOG.Error("添加用例失败!", zap.Error(err))
		response.FailWithMessage("添加用例失败", c)
	} else {
		response.OkWithDetailed(caseApiDetail, "添加用例成功", c)
	}
}

// SetApisCase 设置测试步骤内容
//
//	@Tags		ApiCase
//	@Summary	设置测试步骤内容
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		interfacecaseReq.SetTimerCares	true	"设置测试步骤内容"
//	@Success	200		{string}	string							"{"success":true,"data":{},"msg":"修改成功"}"
//	@Router		/testCase/setApisCase [post]
func (apiCase *ApiCaseApi) SetApisCase(c *gin.Context) {
	var sua interfacecaseReq.SetTimerCares
	_ = c.ShouldBindJSON(&sua)
	if err := apiCaseService.SetApisCase(sua.ID, sua.CaseIds); err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// FindApiCase 用id查询ApiCase
//
//	@Tags		ApiCase
//	@Summary	用id查询ApiCase
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	query		request.ApiCaseIdReq	true	"用id查询ApiCase"
//	@Success	200		{string}	string					"{"success":true,"data":{},"msg":"查询成功"}"
//	@Router		/testCase/findApiCase [get]
func (apiCase *ApiCaseApi) FindApiCase(c *gin.Context) {
	var testCase request.ApiCaseIdReq
	_ = c.ShouldBindQuery(&testCase)
	if err, retestCase := apiCaseService.GetApiCase(testCase.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retestCase": retestCase}, c)
	}
}

// GetApiCaseList 分页获取ApiCase列表
//
//	@Tags		ApiCase
//	@Summary	分页获取ApiCase列表
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	query		interfacecaseReq.ApiCaseSearch	true	"分页获取ApiCase列表"
//	@Success	200		{string}	string							"{"success":true,"data":{},"msg":"获取成功"}"
//	@Router		/testCase/getApiCaseList [get]
func (apiCase *ApiCaseApi) GetApiCaseList(c *gin.Context) {
	var pageInfo interfacecaseReq.ApiCaseSearch
	_ = c.ShouldBindQuery(&pageInfo)
	pageInfo.ProjectID = utils.GetUserProject(c)
	menuId := c.Query("menu")
	menuidInt, _ := strconv.Atoi(menuId)
	pageInfo.ApiMenuID = uint(menuidInt)
	if err, list, total := apiCaseService.GetApiCaseInfoList(pageInfo); err != nil {
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
