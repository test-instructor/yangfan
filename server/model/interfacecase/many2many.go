package interfacecase

import (
	"github.com/test-instructor/cheetah/server/global"
)

type ApiCaseStepRelationship struct {
	global.GVA_MODEL
	ApiStep       ApiStep
	ApiStepId     uint `gorm:"comment:测试步骤"`
	ApiCaseStep   ApiCaseStep
	ApiCaseStepId uint `gorm:"comment:测试套件"`
	Sort          uint `gorm:"comment:排序"`
}

type ApiCaseRelationship struct {
	global.GVA_MODEL
	ApiCaseStep   ApiCaseStep
	ApiCaseStepId uint `gorm:"comment:测试套件"`
	ApiCase       ApiCase
	ApiCaseId     uint `gorm:"comment:测试用例"`
	Sort          uint `gorm:"comment:排序"`
}

type PerformanceRelationship struct {
	global.GVA_MODEL
	Performance   *Performance
	PerformanceId uint `json:"PerformanceId" gorm:"comment:性能任务"`
	//ApiCase       *ApiCase
	//ApiCaseId     uint `gorm:"comment:测试用例"`
	ApiCaseStep   *ApiCaseStep
	ApiCaseStepId uint `json:"ApiCaseStepId" gorm:"comment:测试套件"`
	Sort          uint `gorm:"comment:排序"`
}

type ApiTimerTaskRelationship struct {
	global.GVA_MODEL
	ApiTimerTask   ApiTimerTask
	ApiTimerTaskId uint `gorm:"comment:性能任务"`
	ApiCase        ApiCase
	ApiCaseId      uint `gorm:"comment:测试用例"`
	Sort           uint `gorm:"comment:排序"`
}
