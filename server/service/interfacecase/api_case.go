package interfacecase

import (
	"strconv"

	"gorm.io/gorm"

	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/common/request"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	interfacecaseReq "github.com/test-instructor/yangfan/server/model/interfacecase/request"
)

type ApiCaseService struct {
}

// CreateApiCase 创建ApiCase记录

func (testCaseService *ApiCaseService) CreateApiCase(testCase interfacecase.ApiCase) (err error) {
	err = global.GVA_DB.Create(&testCase).Error
	return err
}

// DeleteApiCase 删除ApiCase记录

func (testCaseService *ApiCaseService) DeleteApiCase(testCase interfacecase.ApiCase) (err error) {
	err = global.GVA_DB.Delete(&testCase).Error
	return err
}

// DeleteApiCaseByIds 批量删除ApiCase记录

func (testCaseService *ApiCaseService) DeleteApiCaseByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]interfacecase.ApiCase{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateApiCase 更新ApiCase记录

func (testCaseService *ApiCaseService) UpdateApiCase(testCase interfacecase.ApiCase) (err error) {
	var oId interfacecase.Operator
	global.GVA_DB.Model(interfacecase.ApiCase{}).Where("id = ?", testCase.ID).First(&oId)
	testCase.CreatedBy = oId.CreatedBy
	testCase.ApiCaseStep = []interfacecase.ApiCaseStep{}
	err = global.GVA_DB.Where("id = ?", testCase.ID).Save(&testCase).Error
	if err != nil {
		return
	}
	global.GVA_Timer.Remove(strconv.Itoa(int(testCase.ID)), testCase.EntryID)
	//if *testCase.Status {
	//	id, err := global.GVA_Timer.AddTaskByFunc(strconv.Itoa(int(testCase.ID)), testCase.RunTime, runTestCase.RunApiCaseBack(testCase.ID), cron.WithSeconds())
	//	if err != nil {
	//		return err
	//	}
	//	testCase.EntryID = int(id)
	//	err = global.GVA_DB.Save(&testCase).Error
	//	if err != nil {
	//		return err
	//	}
	//}
	return err
}

func (testCaseService *ApiCaseService) AddApisCase(taskID uint, caseIDs []uint) (err error) {

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		for _, v := range caseIDs {
			var testCase interfacecase.ApiCaseRelationship
			testCase.ApiCase.ID = taskID
			testCase.Sort = 9999
			testCase.ApiCaseStep.ID = v
			err := tx.Model(interfacecase.ApiCaseRelationship{}).Save(&testCase).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func (testCaseService *ApiCaseService) DelApisCase(testCase interfacecase.ApiCaseRelationship) (err error) {
	err = global.GVA_DB.Delete(&testCase).Error
	return err
}

func (testCaseService *ApiCaseService) SortApisCase(testCase []interfacecase.ApiCaseRelationship) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		for _, v := range testCase {
			err := tx.Find(&interfacecase.ApiCaseRelationship{
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

func (testCaseService *ApiCaseService) FindApiTestCase(id uint) (err error, apiCaseRelationship []interfacecase.ApiCaseRelationship, name string) {
	apiCase := interfacecase.ApiCase{GVA_MODEL: global.GVA_MODEL{ID: id}}
	global.GVA_DB.First(&apiCase)
	name = apiCase.Name
	global.GVA_DB.Model(interfacecase.ApiCaseRelationship{}).
		Where("api_case_id = ?", id).
		Preload("ApiCaseStep").
		Preload("ApiCase").
		Order("Sort").
		Find(&apiCaseRelationship)

	return
}

func (testCaseService *ApiCaseService) AddApiTestCase(apiCaseID request.ApiCaseIdReq) (caseApiDetail interfacecase.ApiStep, err error) {
	caseApiDetail = interfacecase.ApiStep{GVA_MODEL: global.GVA_MODEL{ID: apiCaseID.ApiID}}
	err = global.GVA_DB.Preload("Request").First(&caseApiDetail).Error
	if err != nil {
		return interfacecase.ApiStep{}, err
	}
	caseApiDetail.Parent = caseApiDetail.ID
	caseApiDetail.ID = 0
	caseApiDetail.Request.ID = 0
	caseApiDetail.ApiType = 2
	caseDetail := interfacecase.ApiCaseStep{GVA_MODEL: global.GVA_MODEL{ID: apiCaseID.CaseID}}
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var err error
		err = tx.Create(&caseApiDetail).Error
		if err != nil {
			return err
		}
		err = tx.Model(&caseDetail).Association("TStep").Append(&caseApiDetail)
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		return interfacecase.ApiStep{}, err
	}
	return
}

func (testCaseService *ApiCaseService) SetApisCase(id uint, caseIds []uint) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		TxErr := tx.Delete(&[]interfacecase.ApiCaseRelationship{}, "timer_task_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}
		var timerCase []interfacecase.ApiCaseRelationship
		for _, caseID := range caseIds {
			timerCase = append(timerCase, interfacecase.ApiCaseRelationship{
				ApiCaseId:     id,
				ApiCaseStepId: caseID,
			})
		}
		TxErr = tx.Create(&timerCase).Error
		if TxErr != nil {
			return TxErr
		}
		return nil
	})
}

// GetApiCase 根据id获取ApiCase记录

func (testCaseService *ApiCaseService) GetApiCase(id uint, detail bool) (err error, testCase interfacecase.ApiCase) {
	err = global.GVA_DB.Preload("Project").Where("id = ?", id).First(&testCase).Error
	return
}

// GetApiCaseInfoList 分页获取ApiCase记录

func (testCaseService *ApiCaseService) GetApiCaseInfoList(info interfacecaseReq.ApiCaseSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&interfacecase.ApiCase{}).
		Preload("RunConfig").
		Preload("ApiCaseStep")
	//Preload("Project").Joins("Project").Where("Project.ID = ?", info.ProjectID)
	if info.ApiMenuID > 0 {
		db.Preload("ApiMenu").Joins("ApiMenu").Where("ApiMenu.ID = ?", info.ApiMenuID)
	}
	var testCases []interfacecase.ApiCase
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.FrontCase {
		db.Where("front_case = ?", 1)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Preload("Project").Limit(limit).Offset(offset).Find(&testCases, projectDB(db, info.ProjectID)).Error
	return err, testCases, total
}
