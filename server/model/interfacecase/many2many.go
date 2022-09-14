package interfacecase

import "github.com/test-instructor/cheetah/server/global"

type ApiCaseStepRelationship struct {
	global.GVA_MODEL
	ApiStep       ApiStep
	ApiStepId     uint
	ApiCaseStep   ApiCaseStep
	ApiCaseStepId uint
	Sort          uint
}

type ApiCaseRelationship struct {
	global.GVA_MODEL
	ApiCaseStep   ApiCaseStep
	ApiCaseStepId uint
	ApiCase       ApiCase
	ApiCaseId     uint
	Sort          uint
}

type ApiTimerTaskRelationship struct {
	global.GVA_MODEL
	ApiTimerTask   ApiTimerTask
	ApiTimerTaskId uint
	ApiCase        ApiCase
	ApiCaseId      uint
	Sort           uint
}
