package platform

import (
	"context"
	"errors"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	platformReq "github.com/test-instructor/yangfan/server/v2/model/platform/request"
)

type HarmonyDeviceOptionsService struct{}

// CreateHarmonyDeviceOptions 创建设备选项记录
// Author [yourname](https://github.com/yourname)
func (hdoService *HarmonyDeviceOptionsService) CreateHarmonyDeviceOptions(ctx context.Context, hdo *platform.HarmonyDeviceOptions) (err error) {
	err = global.GVA_DB.Create(hdo).Error
	return err
}

// DeleteHarmonyDeviceOptions 删除设备选项记录
// Author [yourname](https://github.com/yourname)
func (hdoService *HarmonyDeviceOptionsService) DeleteHarmonyDeviceOptions(ctx context.Context, ID string, projectId int64) (err error) {
	var hdo platform.HarmonyDeviceOptions
	err = global.GVA_DB.Where("id = ?", ID).First(&hdo).Error
	if err != nil {
		return err
	}
	if hdo.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}
	err = global.GVA_DB.Delete(&platform.HarmonyDeviceOptions{}, "id = ?", ID).Error
	return err
}

// DeleteHarmonyDeviceOptionsByIds 批量删除设备选项记录
// Author [yourname](https://github.com/yourname)
func (hdoService *HarmonyDeviceOptionsService) DeleteHarmonyDeviceOptionsByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]platform.HarmonyDeviceOptions{}, "id in ?", IDs).Error
	return err
}

// UpdateHarmonyDeviceOptions 更新设备选项记录
// Author [yourname](https://github.com/yourname)
func (hdoService *HarmonyDeviceOptionsService) UpdateHarmonyDeviceOptions(ctx context.Context, hdo platform.HarmonyDeviceOptions, projectId int64) (err error) {
	var oldHarmonyDeviceOptions platform.HarmonyDeviceOptions
	err = global.GVA_DB.Model(&oldHarmonyDeviceOptions).Where("id = ?", hdo.ID).First(&oldHarmonyDeviceOptions).Error
	if err != nil {
		return err
	}
	if oldHarmonyDeviceOptions.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}

	err = global.GVA_DB.Model(&platform.HarmonyDeviceOptions{}).Where("id = ?", hdo.ID).Updates(&hdo).Error
	return err
}

// GetHarmonyDeviceOptions 根据ID获取设备选项记录
// Author [yourname](https://github.com/yourname)
func (hdoService *HarmonyDeviceOptionsService) GetHarmonyDeviceOptions(ctx context.Context, ID string) (hdo platform.HarmonyDeviceOptions, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&hdo).Error
	return
}

// GetHarmonyDeviceOptionsInfoList 分页获取设备选项记录
// Author [yourname](https://github.com/yourname)
func (hdoService *HarmonyDeviceOptionsService) GetHarmonyDeviceOptionsInfoList(ctx context.Context, info platformReq.HarmonyDeviceOptionsSearch) (list []platform.HarmonyDeviceOptions, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&platform.HarmonyDeviceOptions{})
	var hdos []platform.HarmonyDeviceOptions
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

	err = db.Find(&hdos).Error
	return hdos, total, err
}
func (hdoService *HarmonyDeviceOptionsService) GetHarmonyDeviceOptionsPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
