package interfacecase

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/common/request"
	"github.com/test-instructor/yangfan/server/model/common/response"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	"github.com/test-instructor/yangfan/server/service"
	"github.com/test-instructor/yangfan/server/utils"
	"go.uber.org/zap"
)

type RunCaseApi struct {
}

var runCaseService = service.ServiceGroupApp.InterfacecaseServiceGroup.RunCaseService

func (runCaseApi *RunCaseApi) RunTestCaseStep(c *gin.Context) {
	var runCase request.RunCaseReq
	_ = c.ShouldBindJSON(&runCase)
	if runCase.RunType == 1 {
		runCase.RunType = uint(interfacecase.RunTypeDebug)
		reports, err := runCaseService.RunTestCaseStep(runCase)
		if err != nil {
			global.GVA_LOG.Error("运行失败!", zap.Error(err))
			response.FailWithMessage(err.Error(), c)
		} else {
			response.OkWithData(gin.H{"id": reports.ID}, c)
		}
	} else {
		runCase.RunType = uint(interfacecase.RunTypeRunBack)
		go runCaseService.RunTestCaseStep(runCase)
		response.OkWithData("运行成功", c)
	}

}

func (runCaseApi *RunCaseApi) RunApiCase(c *gin.Context) {
	var runApiCase request.RunCaseReq
	_ = c.ShouldBindJSON(&runApiCase)
	runApiCase.RunType = uint(interfacecase.RunTypeRunBack)
	go runCaseService.RunApiCase(runApiCase)
	response.OkWithData("运行成功", c)
}

func (runCaseApi *RunCaseApi) RunBoomerDebug(c *gin.Context) {
	var runApiCase request.RunCaseReq
	_ = c.ShouldBindJSON(&runApiCase)
	runApiCase.RunType = uint(interfacecase.RunTypeDebug)
	reports, err := runCaseService.RunBoomerDebug(runApiCase)
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
	boomer, err := runCaseService.RunMasterBoomer(runApiCase, interfacecase.RunTypeRuning)
	if err != nil {
		global.GVA_LOG.Error("运行失败!", zap.Error(err))
		response.FailWithDetailed(err, "运行失败", c)
	} else {
		response.OkWithData(gin.H{"id": boomer.ID}, c)
	}
}

func (runCaseApi *RunCaseApi) Rebalance(c *gin.Context) {
	var runApiCase request.RunCaseReq
	_ = c.ShouldBindJSON(&runApiCase)
	boomer, err := runCaseService.Rebalance(runApiCase)
	if err != nil {
		return
	}
	if err != nil {
		global.GVA_LOG.Error("运行失败!", zap.Error(err))
		response.FailWithMessage("运行失败", c)
	} else {
		response.OkWithData(gin.H{"id": boomer.ID}, c)
	}
}

func (runCaseApi *RunCaseApi) Stop(c *gin.Context) {
	var runApiCase request.RunCaseReq
	_ = c.ShouldBindQuery(&runApiCase)
	err := runCaseService.Stop(runApiCase)
	if err != nil {
		return
	}
	if err != nil {
		global.GVA_LOG.Error("停止性能呢任务失败!", zap.Error(err))
		response.FailWithMessage("停止性能呢任务失败", c)
	} else {
		response.OkWithMessage("停止性能任务成功", c)
	}
}

func (runCaseApi *RunCaseApi) RunTimerTask(c *gin.Context) {
	var runApiCase request.RunCaseReq
	_ = c.ShouldBindJSON(&runApiCase)
	runApiCase.ProjectID = utils.GetUserProject(c)
	runApiCase.RunType = uint(interfacecase.RunTypeRunBack)
	go runCaseService.RunTimerTask(runApiCase)
	response.OkWithData("运行成功", c)
}

func (runCaseApi *RunCaseApi) RunApi(c *gin.Context) {
	var runCase request.RunCaseReq
	_ = c.ShouldBindJSON(&runCase)
	runCase.RunType = uint(interfacecase.RunTypeRunSave)
	reports, err := runCaseService.RunApi(runCase)
	if err != nil {
		global.GVA_LOG.Error("运行失败!", zap.Error(err))
		response.FailWithMessage("运行失败", c)
	} else {
		response.OkWithData(gin.H{"id": reports.ID}, c)
	}
}
