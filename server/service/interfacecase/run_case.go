package interfacecase

import (
	"github.com/test-instructor/cheetah/server/model/common/request"
	"github.com/test-instructor/cheetah/server/model/interfacecase"
	"github.com/test-instructor/cheetah/server/service/interfacecase/runTestCase"
)

type RunCaseService struct {
}

// RunTestCase TestCase排序

func (apicaseService *RunCaseService) RunTestCase(apiCaseID request.RunCaseReq) (reports interfacecase.ApiReport, err error) {

	reports, err = runTestCase.RunCase(apiCaseID)
	return
}

func (apicaseService *RunCaseService) RunTimerTask(timerTaskStr interfacecase.TimerTask) {
	runTestCase.RunTask(timerTaskStr.ID, 3, 2)
	return
}
