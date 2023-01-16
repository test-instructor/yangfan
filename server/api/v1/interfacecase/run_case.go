package interfacecase

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/model/common/request"
	"github.com/test-instructor/cheetah/server/model/common/response"
	"github.com/test-instructor/cheetah/server/model/interfacecase"
	"github.com/test-instructor/cheetah/server/service"
	"github.com/test-instructor/cheetah/server/utils"
	"go.uber.org/zap"
)

type RunCaseApi struct {
}

var runCaseService = service.ServiceGroupApp.InterfacecaseServiceGroup.RunCaseService

func (runCaseApi *RunCaseApi) RunTestCaseStep(c *gin.Context) {
	var runCase request.RunCaseReq
	_ = c.ShouldBindJSON(&runCase)
	if runCase.RunType == 1 {
		reports, err := runCaseService.RunTestCaseStep(runCase, interfacecase.RunTypeDebug)
		if err != nil {
			global.GVA_LOG.Error("运行失败!", zap.Error(err))
			response.FailWithMessage("运行失败", c)
		} else {
			response.OkWithData(gin.H{"id": reports.ID}, c)
		}
	} else {
		go runCaseService.RunTestCaseStep(runCase, interfacecase.RunTypeRunBack)
		response.OkWithData("运行成功", c)
	}

}

func (runCaseApi *RunCaseApi) RunApiCase(c *gin.Context) {
	var runApiCase request.RunCaseReq
	_ = c.ShouldBindJSON(&runApiCase)
	go runCaseService.RunApiCase(runApiCase, interfacecase.RunTypeRunBack)
	response.OkWithData("运行成功", c)
}

func (runCaseApi *RunCaseApi) RunBoomerDebug(c *gin.Context) {
	var runApiCase request.RunCaseReq
	_ = c.ShouldBindJSON(&runApiCase)

	reports, err := runCaseService.RunBoomerDebug(runApiCase, interfacecase.RunTypeDebug)
	if err != nil {
		global.GVA_LOG.Error("运行失败!", zap.Error(err))
		response.FailWithMessage("运行失败", c)
	} else {
		response.OkWithData(gin.H{"id": reports.ID}, c)
	}
}

func (runCaseApi *RunCaseApi) RunBoomer(c *gin.Context) {
	var runApiCase request.RunCaseReq
	_ = c.ShouldBindJSON(&runApiCase)

	go runCaseService.RunBoomer(runApiCase, interfacecase.RunTypeRuning)
	response.OkWithData("运行成功", c)
	//if err != nil {
	//	global.GVA_LOG.Error("运行失败!", zap.Error(err))
	//	response.FailWithMessage("运行失败", c)
	//} else {
	//	response.OkWithData(gin.H{"id": reports.ID}, c)
	//}
}

func (runCaseApi *RunCaseApi) RunTimerTask(c *gin.Context) {
	var runApiCase request.RunCaseReq
	_ = c.ShouldBindJSON(&runApiCase)
	runApiCase.ProjectID = utils.GetUserProject(c)
	go runCaseService.RunTimerTask(runApiCase, interfacecase.RunTypeRunBack)
	response.OkWithData("运行成功", c)
}

func (runCaseApi *RunCaseApi) RunApi(c *gin.Context) {
	var runCase request.RunCaseReq
	_ = c.ShouldBindJSON(&runCase)
	reports, err := runCaseService.RunApi(runCase)
	if err != nil {
		global.GVA_LOG.Error("运行失败!", zap.Error(err))
		response.FailWithMessage("运行失败", c)
	} else {
		response.OkWithData(gin.H{"id": reports.ID}, c)
	}
}
