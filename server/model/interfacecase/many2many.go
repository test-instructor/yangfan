package interfacecase

type ApiCaseRelationship struct {
	ID            uint
	ApiStep       ApiStep
	ApiStepId     uint
	ApiTestCase   ApiTestCase
	ApiTestCaseId uint
}

type TimerTaskRelationship struct {
	ID            uint
	ApiTestCase   ApiTestCase
	ApiTestCaseId uint
	TimerTask     TimerTask
	TimerTaskId   uint
}
