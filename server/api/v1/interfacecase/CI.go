package interfacecase

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/common/response"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	interfacecaseReq "github.com/test-instructor/yangfan/server/model/interfacecase/request"
	"github.com/test-instructor/yangfan/server/service"
	"go.uber.org/zap"
)

type ApiCIApi struct{}

var ciService = service.ServiceGroupApp.InterfacecaseServiceGroup.ApiCIService

// RunTag @Tags RunTag
//
//	@Summary	CI运行标签
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		interfacecaseReq.CIRun	true	"tag, env, uuid, secret"
//	@Success	200		{object}	string					"{"code":0,"data":{"key": "f2dc1396-1d3e-4f12-8b9a-f9d35e88cd7e","report": 12},"msg":"运行成功"}
//	@Router		/ci/runTag [post]
//	@Router		/ci/runTag [get]
func (ci *ApiCIApi) RunTag(c *gin.Context) {
	var tagReq interfacecaseReq.CIRun
	req, ok := c.Get("req")
	if !ok {
		global.GVA_LOG.Error("提取参数失败")
	}
	tagReq = req.(interfacecaseReq.CIRun)
	global.GVA_LOG.Debug("请求参数", zap.Any("", req))
	if err, data := ciService.RunTag(tagReq); err != nil {
		global.GVA_LOG.Error("运行失败!", zap.Error(err))
		response.FailWithMessage("运行失败", c)
	} else {
		response.OkWithDetailed(data, "运行成功", c)
	}
}

// GetReport @Tags GetReport
//
//	@Summary	获取CI运行的测试报告
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	query		interfacecaseReq.CIRun	true	"key, project, uuid, secret"
//	@Success	200		{object}	string					"{"code":0,"data":{"success": true},"msg":"查询成功"}
//	@Router		/ci/runTag [get]
func (ci *ApiCIApi) GetReport(c *gin.Context) {
	var resp interfacecaseReq.CIRun
	req, ok := c.Get("req")
	if !ok {
		global.GVA_LOG.Error("提取参数失败")
	}
	resp = req.(interfacecaseReq.CIRun)
	if err, data := ciService.GetRepost(resp); err != nil {
		global.GVA_LOG.Error("运行中!", zap.Error(err))
		response.FailWithMessage("测试执行中，请稍后查询", c)
	} else {
		response.OkWithDetailed(data, "查询成功", c)
	}
}

func (ci *ApiCIApi) GetReportDetail(c *gin.Context) {
	var resp interfacecaseReq.CIRun
	req, ok := c.Get("req")
	if !ok {
		global.GVA_LOG.Error("提取参数失败")
	}
	resp = req.(interfacecaseReq.CIRun)
	if err, data := ciService.GetRepost(resp); err != nil {
		global.GVA_LOG.Error("运行中!", zap.Error(err))
		response.FailWithMessage("测试执行中，请稍后查询", c)
	} else {
		response.OkWithDetailed(data, "查询成功", c)
	}
}

func (ci *ApiCIApi) FindReport(c *gin.Context) {
	var resp interfacecase.ApiReport
	req, ok := c.Get("req")
	if !ok {
		global.GVA_LOG.Error("提取参数失败")
	}
	resp = req.(interfacecase.ApiReport)
	if err, reapicase := reportService.FindReport(resp); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"reapicase": reapicase}, c)
	}
}
