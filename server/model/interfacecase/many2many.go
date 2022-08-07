package interfacecase

import "github.com/test-instructor/cheetah/server/global"

type ApiCaseRelationship struct {
	ID            uint
	ApiStep       ApiStep
	ApiStepId     uint
	ApiTestCase   ApiTestCase
	ApiTestCaseId uint
}

type TimerTaskRelationship struct {
	global.GVA_MODEL
	ApiTestCase   ApiTestCase
	ApiTestCaseId uint
	TimerTask     TimerTask
	TimerTaskId   uint
	Sort          uint
}
