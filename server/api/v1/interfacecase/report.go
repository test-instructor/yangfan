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

type ReportApi struct {
}

var reportService = service.ServiceGroupApp.InterfacecaseServiceGroup.ReportService

func (acApi *ReportApi) GetReportList(c *gin.Context) {
	var pageInfo interfacecaseReq.ReportSearch
	_ = c.ShouldBindQuery(&pageInfo)
	pageInfo.ProjectID = utils.GetUserProject(c)
	if err, list, total := reportService.GetReportList(pageInfo); err != nil {
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

func (acApi *ReportApi) GetReportDetail(c *gin.Context) {
	var pageInfo interfacecaseReq.ReportSearch
	_ = c.ShouldBindQuery(&pageInfo)
	pageInfo.ProjectID = utils.GetUserProject(c)
	if err, data := reportService.GetReportDetail(pageInfo.ID); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"data": data}, c)
	}
}

func (acApi *ReportApi) FindReport(c *gin.Context) {
	var apiReport interfacecase.ApiReport
	_ = c.ShouldBindQuery(&apiReport)

	if err, reapicase := reportService.FindReport(apiReport); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"reapicase": reapicase}, c)
	}
}

func (acApi *ReportApi) DelReport(c *gin.Context) {
	var apiReport interfacecase.ApiReport
	_ = c.ShouldBindJSON(&apiReport)

	apiReport.ProjectID = utils.GetUserProject(c)
	apiReport.DeleteBy = utils.GetUserIDAddress(c)

	if err := reportService.DelReport(apiReport); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
