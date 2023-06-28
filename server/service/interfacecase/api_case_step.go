package interfacecase

import (
	"gorm.io/gorm"

	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/common/request"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	interfacecaseReq "github.com/test-instructor/yangfan/server/model/interfacecase/request"
)

type TestCaseService struct {
}

// CreateTestCase 创建TestCase记录

func (t *TestCaseService) CreateTestCaseStep(apicase interfacecase.ApiCaseStep) (err error) {
	err = global.GVA_DB.Create(&apicase).Error
	return err
}

// DeleteTestCase 删除TestCase记录

func (t *TestCaseService) DeleteTestCaseStep(apicase interfacecase.ApiCaseStep) (err error) {
	err = global.GVA_DB.Delete(&apicase).Error
	return err
}

// DeleteTestCaseByIds 批量删除TestCase记录

func (t *TestCaseService) DeleteTestCaseStepByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]interfacecase.ApiCaseStep{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTestCase 更新TestCase记录

func (t *TestCaseService) UpdateTestCaseStep(apicase interfacecase.ApiCaseStep) (err error) {
	var oId interfacecase.Operator
	global.GVA_DB.Model(interfacecase.ApiCaseStep{}).Where("id = ?", apicase.ID).First(&oId)
	apicase.CreatedBy = oId.CreatedBy
	apicase.TStep = []*interfacecase.ApiStep{}
	err = global.GVA_DB.Save(&apicase).Error
	return err
}

// UpdateTestCase TestCase排序

func (t *TestCaseService) SortTestCaseStep(apicase interfacecase.ApiCaseStep) (err error) {

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		for _, v := range apicase.TStep {
			err := tx.Find(&interfacecase.ApiStep{
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

// AddTestCase TestCase排序

func (t *TestCaseService) AddTestCaseStep(apiCaseID request.ApiCaseIdReq) (caseApiDetail *interfacecase.ApiStep, err error) {
	caseApiDetail = &interfacecase.ApiStep{GVA_MODEL: global.GVA_MODEL{ID: apiCaseID.ApiID}}
	err = global.GVA_DB.Preload("Request").Preload("Grpc").First(&caseApiDetail).Error
	if err != nil {
		return nil, err
	}
	caseApiDetail.Parent = caseApiDetail.ID
	caseApiDetail.ID = 0
	if caseApiDetail.Request != nil {
		caseApiDetail.Request.ID = 0
	}
	if apiCaseID.Type == "copy" {
		caseApiDetail.Name = caseApiDetail.Name + "_copy"
	}
	if caseApiDetail.Grpc != nil {
		caseApiDetail.Grpc.ID = 0
	}
	caseApiDetail.ApiType = 2
	caseDetail := interfacecase.ApiCaseStep{GVA_MODEL: global.GVA_MODEL{ID: apiCaseID.CaseID}}
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var err error
		err = tx.Create(caseApiDetail).Error
		if err != nil {
			return err
		}
		err = tx.Model(&caseDetail).Association("TStep").Append(caseApiDetail)
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return
}

// DelTestCase

func (t *TestCaseService) DelTestCaseStep(apiCaseID request.ApiCaseIdReq) (err error) {
	caseApiDetail := interfacecase.ApiStep{GVA_MODEL: global.GVA_MODEL{ID: apiCaseID.ApiID}}
	//err = global.GVA_DB.First(&caseApiDetail).Error
	//if err != nil {
	//	return err
	//}
	caseDetail := interfacecase.ApiCaseStep{GVA_MODEL: global.GVA_MODEL{ID: apiCaseID.CaseID}}
	err = global.GVA_DB.Model(&caseDetail).Association("TStep").Delete(&caseApiDetail)
	return
}

type ToTestCase struct {
	Config    interfacecase.ApiConfig
	TestSteps []interfacecase.ApiStep
}

// GetTestCase 根据id获取TestCase记录

func (t *TestCaseService) FindTestCaseStep(id uint) (err error, apicase interfacecase.ApiCaseStep) {
	err = global.GVA_DB.Model(interfacecase.ApiCaseStep{}).
		Preload("Project").
		Preload("TStep", func(db2 *gorm.DB) *gorm.DB {
			return db2.Joins("Request").Joins("Grpc").Order("Sort")
		}).
		Where("id = ?", id).Select("ID,name").First(&apicase).Error
	resetApiCaseStep(apicase.TStep...)
	return
}

// GetTestCaseInfoList 分页获取TestCase记录

func (t *TestCaseService) GetTestCaseStepInfoList(info interfacecaseReq.TestCaseSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&interfacecase.ApiCaseStep{}).
		Preload("Project").Joins("Project").Where("Project.ID = ?", info.ProjectID)
	if info.ApiMenuID > 0 {
		db.Preload("ApiMenu").Joins("ApiMenu").Where("ApiMenu.ID = ?", info.ApiMenuID)
	}
	db.Where("type = ?", info.ApiType)
	var apicases []interfacecase.ApiCaseStep
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
	err = db.Preload("Project").
		Select("run_config_name,api_env_name,front_case,api_case_steps.created_at,api_case_steps.name,api_case_steps.ID").
		Limit(limit).Offset(offset).Find(&apicases).Error
	return err, apicases, total
}
