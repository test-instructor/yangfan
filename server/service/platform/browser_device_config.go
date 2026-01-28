package platform

import (
	"context"
	"errors"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	platformReq "github.com/test-instructor/yangfan/server/v2/model/platform/request"
)

type BrowserDeviceOptionsService struct{}

// CreateBrowserDeviceOptions 创建浏览器设备选项记录
// Author [yourname](https://github.com/yourname)
func (bdoService *BrowserDeviceOptionsService) CreateBrowserDeviceOptions(ctx context.Context, bdo *platform.BrowserDeviceOptions) (err error) {
	err = global.GVA_DB.Create(bdo).Error
	return err
}

// DeleteBrowserDeviceOptions 删除浏览器设备选项记录
// Author [yourname](https://github.com/yourname)
func (bdoService *BrowserDeviceOptionsService) DeleteBrowserDeviceOptions(ctx context.Context, ID string, projectId int64) (err error) {
	var bdo platform.BrowserDeviceOptions
	err = global.GVA_DB.Where("id = ?", ID).First(&bdo).Error
	if err != nil {
		return err
	}
	if bdo.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}
	err = global.GVA_DB.Delete(&platform.BrowserDeviceOptions{}, "id = ?", ID).Error
	return err
}

// DeleteBrowserDeviceOptionsByIds 批量删除浏览器设备选项记录
// Author [yourname](https://github.com/yourname)
func (bdoService *BrowserDeviceOptionsService) DeleteBrowserDeviceOptionsByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]platform.BrowserDeviceOptions{}, "id in ?", IDs).Error
	return err
}

// UpdateBrowserDeviceOptions 更新浏览器设备选项记录
// Author [yourname](https://github.com/yourname)
func (bdoService *BrowserDeviceOptionsService) UpdateBrowserDeviceOptions(ctx context.Context, bdo platform.BrowserDeviceOptions, projectId int64) (err error) {
	var oldBrowserDeviceOptions platform.BrowserDeviceOptions
	err = global.GVA_DB.Model(&oldBrowserDeviceOptions).Where("id = ?", bdo.ID).First(&oldBrowserDeviceOptions).Error
	if err != nil {
		return err
	}
	if oldBrowserDeviceOptions.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}

	err = global.GVA_DB.Model(&platform.BrowserDeviceOptions{}).Where("id = ?", bdo.ID).Updates(&bdo).Error
	return err
}

// GetBrowserDeviceOptions 根据ID获取浏览器设备选项记录
// Author [yourname](https://github.com/yourname)
func (bdoService *BrowserDeviceOptionsService) GetBrowserDeviceOptions(ctx context.Context, ID string) (bdo platform.BrowserDeviceOptions, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&bdo).Error
	return
}

// GetBrowserDeviceOptionsInfoList 分页获取浏览器设备选项记录
// Author [yourname](https://github.com/yourname)
func (bdoService *BrowserDeviceOptionsService) GetBrowserDeviceOptionsInfoList(ctx context.Context, info platformReq.BrowserDeviceOptionsSearch) (list []platform.BrowserDeviceOptions, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&platform.BrowserDeviceOptions{})
	var bdos []platform.BrowserDeviceOptions
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}
	db.Order("id desc")
	db.Where("project_id = ? ", info.ProjectId)

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&bdos).Error
	return bdos, total, err
}
func (bdoService *BrowserDeviceOptionsService) GetBrowserDeviceOptionsPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
