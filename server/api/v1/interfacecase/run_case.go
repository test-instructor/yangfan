package interfacecase

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/model/common/request"
	"github.com/test-instructor/cheetah/server/model/common/response"
	"github.com/test-instructor/cheetah/server/model/interfacecase"
	"github.com/test-instructor/cheetah/server/service"
	"go.uber.org/zap"
)

type RunCaseApi struct {
}

var runCaseService = service.ServiceGroupApp.InterfacecaseServiceGroup.RunCaseService

func (runCaseApi *RunCaseApi) RunTestCase(c *gin.Context) {
	var runCase request.RunCaseReq
	_ = c.ShouldBindJSON(&runCase)
	reports, err := runCaseService.RunTestCase(runCase)
	if err != nil {
		global.GVA_LOG.Error("运行失败!", zap.Error(err))
		response.FailWithMessage("运行失败", c)
	} else {
		response.OkWithData(gin.H{"reports": reports}, c)
	}
}

func (runCaseApi *RunCaseApi) RunTimerTask(c *gin.Context) {
	var runTimerTask interfacecase.TimerTask
	_ = c.ShouldBindJSON(&runTimerTask)
	runCaseService.RunTimerTask(runTimerTask)
	response.OkWithData("运行成功", c)

	//reports, err := runCaseService.RunTimerTask(runTimerTask)
	//if err != nil {
	//	global.GVA_LOG.Error("运行失败!", zap.Error(err))
	//	response.FailWithMessage("运行失败", c)
	//} else {
	//	response.OkWithData(gin.H{"reports": reports}, c)
	//}
}
