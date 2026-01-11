package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/common/response"
	"github.com/test-instructor/yangfan/server/v2/model/platform/request"
	"go.uber.org/zap"
)

type RunnerApi struct{}

// RunTask
// @Tags Runner
// @Summary 发送运行任务
// @accept application/json
// @Produce application/json
// @Param data body request.RunnerRequest true "运行任务参数"
// @Success 200 {object} response.Response{data=request.RunnerResponse,msg=string} "发送成功"
// @Router /runner/api [post]
func (r *RunnerApi) RunTask(c *gin.Context) {
	var req request.RunnerRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage("参数校验失败: "+err.Error(), c)
		return
	}

	// Call Service
	// Note: rnService is RunnerNodeService, I need RunnerService.
	// In api/v1/platform/enter.go, I will add runnerService variable.
	// But here I can access it via the global service group or just use the variable I will define in enter.go.
	// Since I am in the same package `platform`, I can use the variable `runnerService` that I will define in `enter.go`.

	res, err := runnerService.RunTask(req)
	if err != nil {
		global.GVA_LOG.Error("RunTask failed", zap.Error(err))
		response.FailWithMessage("任务发送失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(res, "success", c)
}
