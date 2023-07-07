package runTestCase

import (
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	"gorm.io/gorm"
)

func caseSort(caseId uint) []interfacecase.ApiCaseRelationship {
	var apiCaseCase []interfacecase.ApiCaseRelationship
	caseDB := global.GVA_DB.Model(interfacecase.ApiCaseRelationship{}).
		Preload("ApiCase").
		Preload("ApiCaseStep").
		Preload("ApiCaseStep.TStep", func(db2 *gorm.DB) *gorm.DB {
			return db2.Order("Sort")
		}).
		Preload("ApiCaseStep.TStep.Request").
		Where("api_case_id = ?", caseId).
		Order("Sort")
	caseDB.Find(&apiCaseCase)
	return apiCaseCase
}

func taskSort(taskId uint) []interfacecase.ApiTimerTaskRelationship {
	var apiCaseCase []interfacecase.ApiTimerTaskRelationship
	caseDB := global.GVA_DB.Model(interfacecase.ApiTimerTaskRelationship{}).
		Preload("ApiTimerTask").
		Preload("ApiCase").
		Preload("ApiCase.ApiCaseStep.TStep.Request").
		Where("api_timer_task_id = ?", taskId).
		Order("Sort")
	caseDB.Find(&apiCaseCase)
	return apiCaseCase
}
