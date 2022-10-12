package runTestCase

import (
	"github.com/test-instructor/cheetah/server/model/common/request"
	"github.com/test-instructor/cheetah/server/model/interfacecase"
)

type ToTestCase struct {
	Config    interfacecase.ApiConfig
	TestSteps []interfacecase.HrpCaseStep
}

func RunApi(runCaseReq request.RunCaseReq, runType interfacecase.RunType) (reports *interfacecase.ApiReport, err error) {
	api := NewRunApi(runCaseReq, runType)
	report, err := RunTestCase(api)
	return report, nil
}

func RunStep(runCaseReq request.RunCaseReq, runType interfacecase.RunType) (reports *interfacecase.ApiReport, err error) {
	api := NewRunStep(runCaseReq, runType)
	report, err := RunTestCase(api)
	return report, nil
}

func RunCase(runCaseReq request.RunCaseReq, runType interfacecase.RunType) (reports *interfacecase.ApiReport, err error) {
	api := NewRunCase(runCaseReq, runType)
	report, err := RunTestCase(api)
	return report, nil
}

func RunTimerTask(runCaseReq request.RunCaseReq, runType interfacecase.RunType) (reports *interfacecase.ApiReport, err error) {
	api := NewRunTask(runCaseReq, runType)
	report, err := RunTestCase(api)
	return report, nil
}

func RunTimerTaskBack(taskID uint) func() {
	return func() {
		var runCaseReq request.RunCaseReq
		runCaseReq.TaskID = taskID
		RunTimerTask(runCaseReq, interfacecase.RunTypeRunTimer)
	}
}

func RunApiCaseBack(apiCaseID uint) func() {
	return func() {
		//RunCase(apiCaseID, interfacecase.CaseTypeCases, interfacecase.RunTypeRunTimer)
	}
}
