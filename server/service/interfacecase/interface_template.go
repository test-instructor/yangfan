package interfacecase

import (
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/model/common/request"
	"github.com/test-instructor/cheetah/server/model/interfacecase"
	interfacecaseReq "github.com/test-instructor/cheetah/server/model/interfacecase/request"
	"gorm.io/gorm"
)

type ApiParams struct {
	Requests string `json:"request" form:"request"`
	Validate string `json:"validate" form:"validate"`
}

type InterfaceTemplateService struct {
}

// CreateInterfaceTemplate 创建InterfaceTemplate记录

func (apicaseService *InterfaceTemplateService) CreateInterfaceTemplate(apicase interfacecase.ApiStep) (err error) {
	apicase.ValidateNumber = uint(len(apicase.Validate))
	err = global.GVA_DB.Create(&apicase).Error
	return err
}

// DeleteInterfaceTemplate 删除InterfaceTemplate记录

func (apicaseService *InterfaceTemplateService) DeleteInterfaceTemplate(apicase interfacecase.ApiStep) (err error) {
	err = global.GVA_DB.Delete(&apicase).Error
	return err
}

// DeleteInterfaceTemplateByIds 批量删除InterfaceTemplate记录

func (apicaseService *InterfaceTemplateService) DeleteInterfaceTemplateByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]interfacecase.ApiStep{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateInterfaceTemplate 更新InterfaceTemplate记录

func (apicaseService *InterfaceTemplateService) UpdateInterfaceTemplate(apicase interfacecase.ApiStep) (id uint, err error) {
	var oId getOperationId
	global.GVA_DB.Model(interfacecase.ApiStep{}).Where("id = ?", apicase.ID).First(&oId)
	apicase.CreatedByID = oId.CreatedByID
	apicase.ValidateNumber = uint(len(apicase.Validate))
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Where(&interfacecase.ApiStep{
			GVA_MODEL: global.GVA_MODEL{ID: apicase.ID},
		}).
			Save(&apicase).Error
		if err != nil {
			return err
		}
		err = tx.Where(&interfacecase.ApiRequest{GVA_MODEL: global.GVA_MODEL{ID: apicase.Request.ID}}).
			Save(&apicase.Request).Error
		return err
	})

	return apicase.ID, err
}

// GetInterfaceTemplate 根据id获取InterfaceTemplate记录

func (apicaseService *InterfaceTemplateService) GetInterfaceTemplate(id uint) (err error, apiCase interfacecase.ApiStep) {
	db := global.GVA_DB.
		Preload("Request").
		Preload("Validate").
		Model(&interfacecase.ApiStep{})
	db.Where("id = ?", id).Find(&apiCase)
	return
}

// GetInterfaceTemplateInfoList 分页获取InterfaceTemplate记录

func (apicaseService *InterfaceTemplateService) GetInterfaceTemplateInfoList(info interfacecaseReq.InterfaceTemplateSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.
		Model(&interfacecase.ApiStep{}).
		Preload("Request").Preload("ApiMenu").
		Preload("Project").Joins("Project").Where("Project.ID = ?", info.ProjectID)
	if info.ApiMenuID > 0 {
		db.Preload("ApiMenu").Joins("ApiMenu").Where("ApiMenu.ID = ?", info.ApiMenuID)
	}
	var apicases []interfacecase.ApiStep

	//查询对应的类型
	db.Where("api_type = ?", info.ApiType)

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}

	err = db.Limit(limit).Offset(offset).Find(&apicases).Error
	return err, apicases, total
}

func (apicaseService *InterfaceTemplateService) UpdateDebugTalk(debugTalk interfacecase.ApiDebugTalk) (err error) {
	err = global.GVA_DB.Save(&debugTalk).Error
	return err
}

func (apicaseService *InterfaceTemplateService) GetDebugTalk(debugTalk interfacecase.ApiDebugTalk) (err error, debugTalkFirst interfacecase.ApiDebugTalk) {

	db := global.GVA_DB.
		Model(&interfacecase.ApiDebugTalk{}).
		Preload("Project").Joins("Project").Where("Project.ID = ?", debugTalk.Project.ID)
	//查询对应的类型
	db.Where("file_type = ?", debugTalk.FileType).Order("id desc")
	err = db.First(&debugTalkFirst).Error
	if err != nil {
		defaultDB := global.GVA_DB.Model(&interfacecase.ApiDebugTalk{}).
			Preload("Project").Joins("Project").Where("Project.ID = ?", 2)
		defaultDB.Where("file_type = ?", debugTalk.FileType)
		err = defaultDB.First(&debugTalkFirst).Error
	}
	return err, debugTalkFirst
}
