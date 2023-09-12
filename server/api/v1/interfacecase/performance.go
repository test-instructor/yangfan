package interfacecase

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/common/response"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	"github.com/test-instructor/yangfan/server/model/interfacecase/request"
	interfacecaseReq "github.com/test-instructor/yangfan/server/model/interfacecase/request"
	"github.com/test-instructor/yangfan/server/service"
	"github.com/test-instructor/yangfan/server/utils"
)

type PerformanceApi struct {
}

var performanceService = service.ServiceGroupApp.InterfacecaseServiceGroup.PerformanceService

func (apiCase *PerformanceApi) CreatePerformance(c *gin.Context) {
	var testCase interfacecase.Performance
	_ = c.ShouldBindJSON(&testCase)
	testCase.ProjectID = utils.GetUserProject(c)
	testCase.CreatedBy = utils.GetUserIDAddress(c)
	if err := performanceService.CreatePerformance(testCase); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (apiCase *PerformanceApi) GetPerformanceList(c *gin.Context) {
	var pageInfo request.PerformancekSearch
	_ = c.ShouldBindQuery(&pageInfo)
	pageInfo.ProjectID = utils.GetUserProject(c)

	if err, list, total := performanceService.GetPerformanceList(pageInfo); err != nil {
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

func (apiCase *PerformanceApi) DeletePerformance(c *gin.Context) {
	var testCase interfacecase.Performance
	_ = c.ShouldBindJSON(&testCase)
	testCase.ProjectID = utils.GetUserProject(c)
	testCase.DeleteBy = utils.GetUserIDAddress(c)
	if err := performanceService.DeletePerformance(testCase); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (apiCase *PerformanceApi) UpdatePerformance(c *gin.Context) {
	var testCase interfacecase.Performance
	_ = c.ShouldBindJSON(&testCase)
	testCase.ProjectID = utils.GetUserProject(c)
	testCase.UpdateBy = utils.GetUserIDAddress(c)
	if err := performanceService.UpdatePerformance(testCase); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (apiCase *PerformanceApi) SortPerformanceCase(c *gin.Context) {
	var testCase []interfacecase.PerformanceRelationship
	_ = c.ShouldBindJSON(&testCase)
	if err := performanceService.SortPerformanceCase(testCase); err != nil {
		global.GVA_LOG.Error("排序失败!", zap.Error(err))
		response.FailWithMessage("排序失败", c)
	} else {
		response.OkWithMessage("排序成功", c)
	}
}

type addPerformanceReq struct {
	PerformanceID uint   `json:"task_id"`
	CaseID        []uint `json:"case_id"`
}

func (apiCase *PerformanceApi) AddPerformanceCase(c *gin.Context) {
	var testCase addPerformanceReq
	_ = c.ShouldBindJSON(&testCase)
	if err := performanceService.AddPerformanceCase(testCase.PerformanceID, testCase.CaseID); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

type Operation struct {
	ApiStep interfacecase.ApiStep `json:"api_step"`
	Pid     uint                  `json:"pid"`
}

func (apiCase *PerformanceApi) AddOperation(c *gin.Context) {
	var testCase Operation
	_ = c.ShouldBindJSON(&testCase)
	testCase.ApiStep.ProjectID = utils.GetUserProject(c)
	testCase.ApiStep.UpdateBy = utils.GetUserIDAddress(c)
	if err := performanceService.AddOperation(testCase.ApiStep, testCase.Pid); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (apiCase *PerformanceApi) DelPerformanceCase(c *gin.Context) {
	var testCase interfacecase.PerformanceRelationship
	_ = c.ShouldBindJSON(&testCase)
	if err := performanceService.DelPerformanceCase(testCase); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (apiCase *PerformanceApi) FindPerformance(c *gin.Context) {
	var testCase interfacecase.Performance
	_ = c.ShouldBindQuery(&testCase)
	testCase.ProjectID = utils.GetUserProject(c)
	testCase.CreatedBy = utils.GetUserIDAddress(c)
	if err, reapicase := performanceService.FindPerformance(testCase.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reapicase": reapicase}, c)
	}
}

func (apiCase *PerformanceApi) FindPerformanceCase(c *gin.Context) {
	var testCase interfacecase.Performance
	_ = c.ShouldBindQuery(&testCase)
	testCase.ProjectID = utils.GetUserProject(c)
	testCase.CreatedBy = utils.GetUserIDAddress(c)
	if err, reapicase, name := performanceService.FindPerformanceCase(testCase.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reapicase": reapicase, "name": name}, c)
	}
}

func (apiCase *PerformanceApi) FindPerformanceStep(c *gin.Context) {
	var testCase interfacecase.ApiCaseStep
	_ = c.ShouldBindQuery(&testCase)
	testCase.ProjectID = utils.GetUserProject(c)
	testCase.CreatedBy = utils.GetUserIDAddress(c)
	if err, reapicase := performanceService.FindPerformanceStep(testCase.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reapicase": reapicase}, c)
	}
}

func (apiCase *PerformanceApi) GetReportList(c *gin.Context) {
	var pageInfo interfacecaseReq.PReportSearch
	_ = c.ShouldBindQuery(&pageInfo)
	pageInfo.ProjectID = utils.GetUserProject(c)
	if err, list, total := performanceService.GetReportList(pageInfo); err != nil {
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

func (apiCase *PerformanceApi) DeleteReport(c *gin.Context) {
	var report interfacecase.PerformanceReport
	_ = c.ShouldBindJSON(&report)
	if err := performanceService.DeleteReport(report); err != nil {
		global.GVA_LOG.Error("删除测试报告失败!", zap.Error(err))
		response.FailWithMessage("删除测试报告失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{}, "删除测试报告成功", c)
	}
}

func (apiCase *PerformanceApi) FindReport(c *gin.Context) {
	var pReport interfacecaseReq.PReportDetail
	_ = c.ShouldBindQuery(&pReport)

	if err, reapicase := performanceService.FindReport(pReport); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		//grafana-host: http://localhost:3000/
		//grafana-dashboard: ERv3OaBPYe6A
		response.OkWithData(gin.H{
			"reapicase":                    reapicase,
			"grafana_host":                 global.GVA_CONFIG.YangFan.GrafanaHost,
			"grafana_dashboard":            global.GVA_CONFIG.YangFan.GrafanaDashboard,
			"grafana_dashboard_name":       global.GVA_CONFIG.YangFan.GrafanaDashboardName,
			"grafana_dashboard_stats":      global.GVA_CONFIG.YangFan.GrafanaDashboardStats,
			"grafana_dashboard_stats_name": global.GVA_CONFIG.YangFan.GrafanaDashboardStatsName,
		}, c)
	}
}

func (apiCase *PerformanceApi) Create(c *gin.Context) {
	var testCase interfacecase.PerformanceCase
	_ = c.ShouldBindJSON(&testCase)
	testCase.ProjectID = utils.GetUserProject(c)
	testCase.CreatedBy = utils.GetUserIDAddress(c)
	if err := performanceService.Create(testCase); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
