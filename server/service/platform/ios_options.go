package platform

import (
	"context"
	"errors"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	platformReq "github.com/test-instructor/yangfan/server/v2/model/platform/request"
)

type IOSDeviceOptionsService struct{}

// CreateIOSDeviceOptions 创建iOS设备选项记录
// Author [yourname](https://github.com/yourname)
func (idoService *IOSDeviceOptionsService) CreateIOSDeviceOptions(ctx context.Context, ido *platform.IOSDeviceOptions) (err error) {
	err = global.GVA_DB.Create(ido).Error
	return err
}

// DeleteIOSDeviceOptions 删除iOS设备选项记录
// Author [yourname](https://github.com/yourname)
func (idoService *IOSDeviceOptionsService) DeleteIOSDeviceOptions(ctx context.Context, ID string, projectId int64) (err error) {
	var ido platform.IOSDeviceOptions
	err = global.GVA_DB.Where("id = ?", ID).First(&ido).Error
	if err != nil {
		return err
	}
	if ido.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}
	err = global.GVA_DB.Delete(&platform.IOSDeviceOptions{}, "id = ?", ID).Error
	return err
}

// DeleteIOSDeviceOptionsByIds 批量删除iOS设备选项记录
// Author [yourname](https://github.com/yourname)
func (idoService *IOSDeviceOptionsService) DeleteIOSDeviceOptionsByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]platform.IOSDeviceOptions{}, "id in ?", IDs).Error
	return err
}

// UpdateIOSDeviceOptions 更新iOS设备选项记录
// Author [yourname](https://github.com/yourname)
func (idoService *IOSDeviceOptionsService) UpdateIOSDeviceOptions(ctx context.Context, ido platform.IOSDeviceOptions, projectId int64) (err error) {
	var oldIOSDeviceOptions platform.IOSDeviceOptions
	err = global.GVA_DB.Model(&oldIOSDeviceOptions).Where("id = ?", ido.ID).First(&oldIOSDeviceOptions).Error
	if err != nil {
		return err
	}
	if oldIOSDeviceOptions.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}

	err = global.GVA_DB.Model(&platform.IOSDeviceOptions{}).Where("id = ?", ido.ID).Updates(&ido).Error
	return err
}

// GetIOSDeviceOptions 根据ID获取iOS设备选项记录
// Author [yourname](https://github.com/yourname)
func (idoService *IOSDeviceOptionsService) GetIOSDeviceOptions(ctx context.Context, ID string) (ido platform.IOSDeviceOptions, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&ido).Error
	return
}

// GetIOSDeviceOptionsInfoList 分页获取iOS设备选项记录
// Author [yourname](https://github.com/yourname)
func (idoService *IOSDeviceOptionsService) GetIOSDeviceOptionsInfoList(ctx context.Context, info platformReq.IOSDeviceOptionsSearch) (list []platform.IOSDeviceOptions, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&platform.IOSDeviceOptions{})
	var idos []platform.IOSDeviceOptions
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}
	db.Order("id desc")
	db.Where("project_id = ? ", info.ProjectId)

	if info.Name != nil && *info.Name != "" {
		db = db.Where("name LIKE ?", "%"+*info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&idos).Error
	return idos, total, err
}
func (idoService *IOSDeviceOptionsService) GetIOSDeviceOptionsPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
