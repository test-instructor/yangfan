package runTestCase

import (
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/common/request"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
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
	api := NewRunTask(runCaseReq, runType)
	report, err := RunTestCase(api)
	return report, nil
}

func RunTimerTag(runCaseReq request.RunCaseReq, runType interfacecase.RunType) (reports *interfacecase.ApiReport, err error) {
	api := NewRunTag(runCaseReq, runType)
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
