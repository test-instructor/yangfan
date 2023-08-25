package interfacecase

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/common/response"
	interfacecaseReq "github.com/test-instructor/yangfan/server/model/interfacecase/request"
	"github.com/test-instructor/yangfan/server/service"
	"go.uber.org/zap"
)

type ApiCIApi struct{}

var ciService = service.ServiceGroupApp.InterfacecaseServiceGroup.ApiCIService

// RunTag @Tags RunTag
// @Summary CI运行标签
// @accept application/json
// @Produce application/json
// @Param data body interfacecaseReq.CIRun true "tag, env, uuid, secret"
// @Success 200 {object} string "{"code":0,"data":{"key": "f2dc1396-1d3e-4f12-8b9a-f9d35e88cd7e","report": 12},"msg":"运行成功"}
// @Router /ci/runTag [post]
func (ci *ApiCIApi) RunTag(c *gin.Context) {
	var tagReq interfacecaseReq.CIRun
	_ = c.ShouldBindQuery(&tagReq)
	if tagReq.TagID == 0 || tagReq.ProjectID == 0 || tagReq.EnvID == 0 {
		_ = c.ShouldBindJSON(&tagReq)
	}
	if err, data := ciService.RunTag(tagReq); err != nil {
		global.GVA_LOG.Error("运行失败!", zap.Error(err))
		response.FailWithMessage("运行失败", c)
	} else {
		response.OkWithDetailed(data, "运行成功", c)
	}
}

func (ci *ApiCIApi) GetReport(c *gin.Context) {
	var resp interfacecaseReq.CIRun
	_ = c.ShouldBindQuery(&resp)
	if err, data := ciService.GetRepost(resp); err != nil {
		global.GVA_LOG.Error("运行中!", zap.Error(err))
		response.FailWithMessage("测试执行中，请稍后查询", c)
	} else {
		response.OkWithDetailed(data, "运行成功", c)
	}
}
