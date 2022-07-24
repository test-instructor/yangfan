package interfacecase

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/model/common/response"
	"github.com/test-instructor/cheetah/server/model/interfacecase"
	interfacecaseReq "github.com/test-instructor/cheetah/server/model/interfacecase/request"
	"github.com/test-instructor/cheetah/server/model/system"
	"github.com/test-instructor/cheetah/server/service"
	"go.uber.org/zap"
)

type ReportApi struct {
}

var reportService = service.ServiceGroupApp.InterfacecaseServiceGroup.ReportService

func (acApi *ReportApi) GetReportList(c *gin.Context) {
	var pageInfo interfacecaseReq.ReportSearch
	_ = c.ShouldBindQuery(&pageInfo)
	project, _ := c.Get("project")
	pageInfo.Project = project.(system.Project)

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
