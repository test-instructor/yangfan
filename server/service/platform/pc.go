package platform

import (
	"context"
	"errors"

	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	platformReq "github.com/test-instructor/yangfan/server/v2/model/platform/request"
)

type PythonCodeService struct{}

// CreatePythonCode 创建python 函数记录
// Author [yourname](https://github.com/yourname)
func (pcService *PythonCodeService) CreatePythonCode(ctx context.Context, pc *platform.PythonCode) (err error) {
	err = global.GVA_DB.Create(pc).Error
	return err
}

// DeletePythonCode 删除python 函数记录
// Author [yourname](https://github.com/yourname)
func (pcService *PythonCodeService) DeletePythonCode(ctx context.Context, ID string, projectId int64) (err error) {
	var pc platform.PythonCode
	err = global.GVA_DB.Where("id = ?", ID).First(&pc).Error
	if err != nil {
		return err
	}
	if pc.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}
	err = global.GVA_DB.Delete(&platform.PythonCode{}, "id = ?", ID).Error
	return err
}

// DeletePythonCodeByIds 批量删除python 函数记录
// Author [yourname](https://github.com/yourname)
func (pcService *PythonCodeService) DeletePythonCodeByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]platform.PythonCode{}, "id in ?", IDs).Error
	return err
}

// UpdatePythonCode 更新python 函数记录
// Author [yourname](https://github.com/yourname)
func (pcService *PythonCodeService) UpdatePythonCode(ctx context.Context, pc platform.PythonCode, projectId int64) (err error) {
	err = global.GVA_DB.Model(&platform.PythonCode{}).Create(&pc).Error
	return err
}

// GetPythonCode 根据ID获取python 函数记录
// Author [yourname](https://github.com/yourname)
func (pcService *PythonCodeService) GetPythonCode(ctx context.Context, pageInfo platformReq.PythonCodeSearch, projectId uint) (pc platform.PythonCode, err error) {
	db := global.GVA_DB.Model(&platform.PythonCode{})
	if pageInfo.ID != nil && *pageInfo.ID != 0 {
		db.Where("id = ?", *pageInfo.ID)
	} else {

		if pageInfo.Type != 0 {
			db.Where("type = ?", pageInfo.Type)
		}
		if pageInfo.UniqueKey != "" {
			db.Where("unique_key = ?", pageInfo.UniqueKey)
		}
		db.Where("project_id = ?", projectId)
	}
	err = db.Order("id desc").First(&pc).Error

	return
}

// GetPythonCodeInfoList 分页获取python 函数记录
// Author [yourname](https://github.com/yourname)
func (pcService *PythonCodeService) GetPythonCodeInfoList(ctx context.Context, info platformReq.PythonCodeSearch) (list []platform.PythonCode, total int64, err error) {

	// 创建db
	db := global.GVA_DB.Model(&platform.PythonCode{})
	var pcs []platform.PythonCode
	db.Select("id, created_at, updated_at, deleted_at, type, project_id, update_by")
	db.Where("project_id = ?", info.ProjectId)
	if info.UniqueKey != "" {
		db.Where("unique_key = ?", info.UniqueKey)
	}
	if info.Type != 0 {
		db.Where("type = ?", info.Type)
	}

	db.Order("id DESC")
	err = db.Find(&pcs).Error
	return pcs, total, err
}

func (pcService *PythonCodeService) GetPythonCodePublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
