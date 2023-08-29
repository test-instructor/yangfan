package runTestCase

import (
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/common/request"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	"go.uber.org/zap"
)

type ToTestCase struct {
	Config    interfacecase.ApiConfig
	TestSteps []interface{}
}

func RunApi(runCaseReq request.RunCaseReq, runType interfacecase.RunType) (reports *interfacecase.ApiReport, err error) {
	api := NewRunApi(runCaseReq, runType)
	report, err := RunTestCase(api)
	return report, nil
}

func RunStep(runCaseReq request.RunCaseReq, runType interfacecase.RunType) (reports *interfacecase.ApiReport, err error) {
	api := NewRunStep(runCaseReq, runType)
	reports, err = RunTestCase(api)
	return
}

func RunCase(runCaseReq request.RunCaseReq, runType interfacecase.RunType) (reports *interfacecase.ApiReport, err error) {
	api := NewRunCase(runCaseReq, runType)
	report, err := RunTestCase(api)
	return report, nil
}

func RunBoomerDebug(runCaseReq request.RunCaseReq, runType interfacecase.RunType) (reports *interfacecase.ApiReport, err error) {
	api := NewBoomerDebug(runCaseReq, runType)
	report, err := RunTestCase(api)
	return report, nil
}

func RunBoomer(runCaseReq request.RunCaseReq, runType interfacecase.RunType) (reports *interfacecase.ApiReport, err error) {
	if runCaseReq.Operation.Running == request.RunningTypeRun {
		api := NewBoomer(runCaseReq, runType)
		report, err := RunTestCase(api)
		return report, err
	}
	RunYangfanBoomer(nil, nil, nil, runCaseReq)
	return nil, nil
}

func RunTimerTask(runCaseReq request.RunCaseReq, runType interfacecase.RunType) (reports *interfacecase.ApiReport, err error) {
	defer func() {
		if msg := recover(); msg != nil {
			global.GVA_LOG.Error("测试用例运行时报错：", zap.Any("运行任务标签时出现异常", msg))
		}
	}()
	api := NewRunTask(runCaseReq, runType, nil)
	report, err := RunTestCase(api)
	return report, nil
}

func RunTimerTagCI(runCaseReq request.RunCaseReq, runType interfacecase.RunType) (reports *interfacecase.ApiReport, err error) {
	api := NewRunTag(runCaseReq, runType, nil)
	report, err := RunTestCase(api)
	return report, nil
}

func RunTimerTaskBack(taskID uint) func() {
	return func() {
		var runCaseReq request.RunCaseReq
		runCaseReq.TaskID = taskID
		global.GVA_LOG.Debug("执行定时任务")
		RunTimerTask(runCaseReq, interfacecase.RunTypeRunTimer)
	}
}
