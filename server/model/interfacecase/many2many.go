package interfacecase

import (
	"github.com/test-instructor/yangfan/server/global"
)

type ApiCaseStepRelationship struct {
	global.GVA_MODEL
	ApiStep       ApiStep
	ApiStepId     uint `gorm:"comment:测试步骤"`
	ApiCaseStep   ApiCaseStep
	ApiCaseStepId uint `gorm:"comment:测试步骤"`
	Sort          uint `gorm:"comment:排序"`
}

type ApiCaseRelationship struct {
	global.GVA_MODEL
	ApiCaseStep   ApiCaseStep
	ApiCaseStepId uint `gorm:"comment:测试步骤"` // 测试步骤
	ApiCase       ApiCase
	ApiCaseId     uint `gorm:"comment:测试用例"` // 测试用例
	Sort          uint `gorm:"comment:排序"`   // 排序
}

type PerformanceRelationship struct {
	global.GVA_MODEL
	Performance   *Performance
	PerformanceId uint `json:"PerformanceId" gorm:"comment:性能任务"`
	//ApiCase       *ApiCase
	//ApiCaseId     uint `gorm:"comment:测试用例"`
	ApiCaseStep   *ApiCaseStep
	ApiCaseStepId uint `json:"ApiCaseStepId" gorm:"comment:测试步骤"`
	Sort          uint `gorm:"comment:排序"`
}

type ApiTimerTaskRelationship struct {
	global.GVA_MODEL
	ApiTimerTask   ApiTimerTask
	ApiTimerTaskId uint `gorm:"comment:定时任务"`
	ApiCase        ApiCase
	ApiCaseId      uint `gorm:"comment:测试用例"`
	Sort           uint `gorm:"comment:排序"`
}

type PerformanceCaseRelationship struct {
	global.GVA_MODEL
	Performance       Performance
	PerformanceId     uint
	PerformanceCase   PerformanceCase
	PerformanceCaseId uint
}
