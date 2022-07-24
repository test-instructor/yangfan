package interfacecase

import (
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/model/common/request"
	"github.com/test-instructor/cheetah/server/model/interfacecase"
	interfacecaseReq "github.com/test-instructor/cheetah/server/model/interfacecase/request"
)

type ApiConfigService struct {
}

// CreateApiConfig 创建ApiConfig记录

func (acService *ApiConfigService) CreateApiConfig(ac interfacecase.ApiConfig) (err error) {
	err = global.GVA_DB.Create(&ac).Error
	return err
}

// DeleteApiConfig 删除ApiConfig记录

func (acService *ApiConfigService) DeleteApiConfig(ac interfacecase.ApiConfig) (err error) {
	err = global.GVA_DB.Delete(&ac).Error
	return err
}

// DeleteApiConfigByIds 批量删除ApiConfig记录

func (acService *ApiConfigService) DeleteApiConfigByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]interfacecase.ApiConfig{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateApiConfig 更新ApiConfig记录

func (acService *ApiConfigService) UpdateApiConfig(ac interfacecase.ApiConfig) (err error) {
	err = global.GVA_DB.Where(&interfacecase.ApiConfig{GVA_MODEL: global.GVA_MODEL{ID: ac.ID}}).
		Save(&ac).Error
	return err
}

// GetApiConfig 根据id获取ApiConfig记录

func (acService *ApiConfigService) GetApiConfig(id uint) (err error, ac interfacecase.ApiConfig) {
	err = global.GVA_DB.
		Where("id = ?", id).First(&ac).Error
	return
}

type configStruct struct {
	ID   uint
	Name string
}

// GetApiConfigInfoList 分页获取ApiConfig记录

func (acService *ApiConfigService) GetApiConfigInfoList(info interfacecaseReq.ApiConfigSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&interfacecase.ApiConfig{})
	var acs []interfacecase.ApiConfig
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	//configStruct2 := configStruct{}
	err = db.Preload("Project").
		Model(&interfacecase.ApiConfig{}).
		Limit(limit).Offset(offset).Find(&acs, projectDB(db, info.Project.ID)).
		Select("api_configs.id", "api_configs.name").
		Error
	return err, acs, total
}
