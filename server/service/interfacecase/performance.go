package interfacecase

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	"github.com/test-instructor/yangfan/server/model/interfacecase/request"
	interfacecaseReq "github.com/test-instructor/yangfan/server/model/interfacecase/request"
)

type PerformanceService struct {
}

func (testCaseService *PerformanceService) CreatePerformance(testCase interfacecase.Performance) (err error) {
	err = global.GVA_DB.Create(&testCase).Error
	return err
}

func (testCaseService *PerformanceService) GetPerformanceList(info request.PerformancekSearch) (err error, list interface{}, total int64) {

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&interfacecase.Performance{}).
		Preload("RunConfig")
	db.Preload("Project").Limit(limit).Offset(offset)
	var testCase []interfacecase.Performance
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("performances.name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Find(&testCase, projectDB(db, info.ProjectID)).Error
	return err, testCase, total
}

func (testCaseService *PerformanceService) DeletePerformance(testCase interfacecase.Performance) (err error) {
	err = global.GVA_DB.Delete(&testCase).Error
	return err
}

func (testCaseService *PerformanceService) UpdatePerformance(testCase interfacecase.Performance) (err error) {
	var getCase interfacecase.Operator
	global.GVA_DB.Model(&interfacecase.Performance{}).Where("id = ?", testCase.ID).First(&getCase)
	testCase.CreatedBy = getCase.CreatedBy
	err = global.GVA_DB.Save(&testCase).Error
	return err
}

func (testCaseService *PerformanceService) SortPerformanceCase(testCase []interfacecase.PerformanceRelationship) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		for _, v := range testCase {
			err := tx.Find(&interfacecase.PerformanceRelationship{
				GVA_MODEL: global.GVA_MODEL{ID: v.ID},
			}).Update("Sort", v.Sort).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func (testCaseService *PerformanceService) AddPerformanceCase(performanceID uint, caseIDs []uint) (err error) {

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		for _, v := range caseIDs {
			var testcase interfacecase.PerformanceRelationship
			testcase.PerformanceId = performanceID
			testcase.Sort = 9999
			testcase.ApiCaseStepId = v
			err := tx.Model(interfacecase.PerformanceRelationship{}).Create(&testcase).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func (testCaseService *PerformanceService) AddOperation(testcase interfacecase.ApiStep, pid uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&interfacecase.ApiStep{}).Save(&testcase).Error
		if err != nil {
			return err
		}
		var apiCaseStep interfacecase.ApiCaseStep
		apiCaseStep.Name = testcase.Name
		apiCaseStep.ProjectID = testcase.ProjectID
		apiCaseStep.CreatedBy = testcase.CreatedBy
		apiCaseStep.Type = interfacecase.ApiTypeCasePerformance
		if testcase.Transaction != nil {
			var transaction interfacecase.ApiStepType
			if testcase.Transaction.Type == interfacecase.TransactionStart {
				transaction = interfacecase.ApiStepTypeTransactionStart
			}
			if testcase.Transaction.Type == interfacecase.TransactionEnd {
				transaction = interfacecase.ApiStepTypeTransactionEnd
			}
			apiCaseStep.ApiStepType = transaction
		}
		if testcase.Rendezvous != nil {
			apiCaseStep.ApiStepType = interfacecase.ApiStepTypeRendezvous
		}
		err = tx.Model(&interfacecase.ApiCaseStep{}).Create(&apiCaseStep).Error
		if err != nil {
			return err
		}
		var apiCaseStepRelationship interfacecase.ApiCaseStepRelationship
		apiCaseStepRelationship.ApiStepId = testcase.ID
		apiCaseStepRelationship.ApiCaseStepId = apiCaseStep.ID
		err = tx.Model(&interfacecase.ApiCaseStepRelationship{}).Create(&apiCaseStepRelationship).Error
		if err != nil {
			return err
		}

		var performanceRelationship interfacecase.PerformanceRelationship
		performanceRelationship.PerformanceId = pid
		performanceRelationship.ApiCaseStepId = apiCaseStep.ID
		performanceRelationship.Sort = 999
		err = tx.Model(&interfacecase.PerformanceRelationship{}).Create(&performanceRelationship).Error

		return err
	})
	return err
}

func (testCaseService *PerformanceService) DelPerformanceCase(testCase interfacecase.PerformanceRelationship) (err error) {
	err = global.GVA_DB.Delete(&testCase).Error
	return err
}

func (testCaseService *PerformanceService) FindPerformance(id uint) (err error, testCase interfacecase.Performance) {
	testCase = interfacecase.Performance{GVA_MODEL: global.GVA_MODEL{ID: id}}
	err = global.GVA_DB.First(&testCase).Error
	return err, testCase
}

func (testCaseService *PerformanceService) FindPerformanceCase(id uint) (err error, performanceRelationship []interfacecase.PerformanceRelationship, name string) {
	testCase := interfacecase.Performance{GVA_MODEL: global.GVA_MODEL{ID: id}}
	err = global.GVA_DB.First(&testCase).Error
	name = testCase.Name
	global.GVA_DB.Model(&interfacecase.PerformanceRelationship{}).
		Where("performance_id = ?", id).
		Preload("ApiCaseStep").
		Order("Sort").
		Find(&performanceRelationship)
	return
}

func (testCaseService *PerformanceService) FindPerformanceStep(id uint) (err error, step interfacecase.ApiStep) {
	testCase := interfacecase.ApiCaseStepRelationship{ApiCaseStepId: id}
	err = global.GVA_DB.Model(&interfacecase.ApiCaseStepRelationship{}).
		Where("api_case_step_id=?", id).
		Preload("ApiCaseStep").
		Preload("ApiStep").
		First(&testCase).Error
	err = global.GVA_DB.Model(&interfacecase.ApiStep{}).
		Where("id=?", &testCase.ApiStepId).
		Preload("Transaction").
		Preload("Rendezvous").
		Find(&step).Error
	return
}

func (testCaseService *PerformanceService) GetReportList(info interfacecaseReq.PReportSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	var pReport []interfacecase.PerformanceReport
	db := global.GVA_DB.
		Model(&interfacecase.PerformanceReport{}).
		Preload("Project").Joins("Project").Where("Project.ID = ?", info.ProjectID)
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	fmt.Println(info.ProjectID)
	err = db.Limit(limit).Offset(offset).Order("ID desc").Find(&pReport).Error

	return err, pReport, total
}

func (testCaseService *PerformanceService) FindReport(pReportReq interfacecaseReq.PReportDetail) (err error, report interface{}) {
	var pReport interfacecase.PerformanceReport
	// 创建db
	db := global.GVA_DB.
		Model(&interfacecase.PerformanceReport{})
	db.Preload("PerformanceReportDetail.PerformanceReportTotalStats")

	if pReportReq.DetailID == 0 {
		db.Preload("PerformanceReportDetail")
	} else {
		db.Preload("PerformanceReportDetail", "id > ?", pReportReq.DetailID)
	}

	err = db.Find(&pReport, "id = ?", pReportReq.ID).Error

	return err, pReport
}
