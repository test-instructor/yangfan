package runTestCase

import (
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/automation"
)

// Replaces interfacecase.ApiCaseRelationship
// Yangfan: AutoCaseStepList
func caseSort(caseId uint) []automation.AutoCaseStepList {
	var apiCaseCase []automation.AutoCaseStepList
	caseDB := global.GVA_DB.Model(automation.AutoCaseStepList{}).
		Preload("AutoCase").
		Preload("AutoCaseStep").
		Where("case_id = ?", caseId).
		Order("sort")
	caseDB.Find(&apiCaseCase)

	return apiCaseCase
}

// Replaces interfacecase.ApiTimerTaskRelationship
// Yangfan: TimerTaskCaseList
func taskSort(taskId uint) []automation.TimerTaskCaseList {
	var apiCaseCase []automation.TimerTaskCaseList
	caseDB := global.GVA_DB.Model(automation.TimerTaskCaseList{}).
		Preload("AutoCase").
		Where("task_id = ?", taskId).
		Order("sort")
	caseDB.Find(&apiCaseCase)
	return apiCaseCase
}
