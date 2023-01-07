package interfacecase

import (
	"github.com/test-instructor/cheetah/server/model/common/request"
	"github.com/test-instructor/cheetah/server/model/interfacecase"
	"github.com/test-instructor/cheetah/server/service/interfacecase/runTestCase"
)

type RunCaseService struct {
}

// RunTestCase TestCase排序

func (apicaseService *RunCaseService) RunTestCaseStep(runCase request.RunCaseReq, runType interfacecase.RunType) (reports *interfacecase.ApiReport, err error) {
	reports, err = runTestCase.RunStep(runCase, runType)
	return
}

func (apicaseService *RunCaseService) RunApiCase(runCase request.RunCaseReq, runType interfacecase.RunType) (report *interfacecase.ApiReport, err error) {
	report, err = runTestCase.RunCase(runCase, runType)
	if err != nil {
		return
	}
	return
}

func (apicaseService *RunCaseService) RunBoomerDebug(runCase request.RunCaseReq, runType interfacecase.RunType) (report *interfacecase.ApiReport, err error) {
	report, err = runTestCase.RunBoomerDebug(runCase, runType)
	return
}

func (apicaseService *RunCaseService) RunBoomer(runCase request.RunCaseReq, runType interfacecase.RunType) (report *interfacecase.ApiReport, err error) {
	report, err = runTestCase.RunBoomer(runCase, runType)
	return
}

func (apicaseService *RunCaseService) RunTimerTask(runCase request.RunCaseReq, runType interfacecase.RunType) {
	runTestCase.RunTimerTask(runCase, runType)
	return
}

func (apicaseService *RunCaseService) RunApi(runCase request.RunCaseReq) (reports *interfacecase.ApiReport, err error) {
	report, err := runTestCase.RunApi(runCase, interfacecase.RunType(runCase.RunType))
	if err != nil {
		return nil, err
	}
	return report, nil
}
