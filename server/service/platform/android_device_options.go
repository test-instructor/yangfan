package platform

import (
	"context"
	"errors"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	platformReq "github.com/test-instructor/yangfan/server/v2/model/platform/request"
)

type AndroidDeviceOptionsService struct{}

// CreateAndroidDeviceOptions 创建安卓设备记录
// Author [yourname](https://github.com/yourname)
func (adoService *AndroidDeviceOptionsService) CreateAndroidDeviceOptions(ctx context.Context, ado *platform.AndroidDeviceOptions) (err error) {
	err = global.GVA_DB.Create(ado).Error
	return err
}

// DeleteAndroidDeviceOptions 删除安卓设备记录
// Author [yourname](https://github.com/yourname)
func (adoService *AndroidDeviceOptionsService) DeleteAndroidDeviceOptions(ctx context.Context, ID string, projectId int64) (err error) {
	var ado platform.AndroidDeviceOptions
	err = global.GVA_DB.Where("id = ?", ID).First(&ado).Error
	if err != nil {
		return err
	}
	if ado.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}
	err = global.GVA_DB.Delete(&platform.AndroidDeviceOptions{}, "id = ?", ID).Error
	return err
}

// DeleteAndroidDeviceOptionsByIds 批量删除安卓设备记录
// Author [yourname](https://github.com/yourname)
func (adoService *AndroidDeviceOptionsService) DeleteAndroidDeviceOptionsByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]platform.AndroidDeviceOptions{}, "id in ?", IDs).Error
	return err
}

// UpdateAndroidDeviceOptions 更新安卓设备记录
// Author [yourname](https://github.com/yourname)
func (adoService *AndroidDeviceOptionsService) UpdateAndroidDeviceOptions(ctx context.Context, ado platform.AndroidDeviceOptions, projectId int64) (err error) {
	var oldAndroidDeviceOptions platform.AndroidDeviceOptions
	err = global.GVA_DB.Model(&oldAndroidDeviceOptions).Where("id = ?", ado.ID).First(&oldAndroidDeviceOptions).Error
	if err != nil {
		return err
	}
	if oldAndroidDeviceOptions.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}

	err = global.GVA_DB.Model(&platform.AndroidDeviceOptions{}).Where("id = ?", ado.ID).Updates(&ado).Error
	return err
}

// GetAndroidDeviceOptions 根据ID获取安卓设备记录
// Author [yourname](https://github.com/yourname)
func (adoService *AndroidDeviceOptionsService) GetAndroidDeviceOptions(ctx context.Context, ID string) (ado platform.AndroidDeviceOptions, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&ado).Error
	return
}

// GetAndroidDeviceOptionsInfoList 分页获取安卓设备记录
// Author [yourname](https://github.com/yourname)
func (adoService *AndroidDeviceOptionsService) GetAndroidDeviceOptionsInfoList(ctx context.Context, info platformReq.AndroidDeviceOptionsSearch) (list []platform.AndroidDeviceOptions, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&platform.AndroidDeviceOptions{})
	var ados []platform.AndroidDeviceOptions
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
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["created_at"] = true
	orderMap["serial_number"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&ados).Error
	return ados, total, err
}
func (adoService *AndroidDeviceOptionsService) GetAndroidDeviceOptionsPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
