package platform

import (
	"context"
	"errors"

	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	platformReq "github.com/test-instructor/yangfan/server/v2/model/platform/request"
)

type EnvDetailService struct{}

// CreateEnvDetail 创建环境详情记录
// Author [yourname](https://github.com/yourname)
func (edService *EnvDetailService) CreateEnvDetail(ctx context.Context, ed *platform.EnvDetail) (err error) {
	err = global.GVA_DB.Create(ed).Error
	return err
}

// DeleteEnvDetail 删除环境详情记录
// Author [yourname](https://github.com/yourname)
func (edService *EnvDetailService) DeleteEnvDetail(ctx context.Context, ID string, projectId int64) (err error) {
	var ed platform.EnvDetail
	err = global.GVA_DB.Where("id = ?", ID).First(&ed).Error
	if err != nil {
		return err
	}
	if ed.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}
	err = global.GVA_DB.Delete(&platform.EnvDetail{}, "id = ?", ID).Error
	return err
}

// DeleteEnvDetailByIds 批量删除环境详情记录
// Author [yourname](https://github.com/yourname)
func (edService *EnvDetailService) DeleteEnvDetailByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]platform.EnvDetail{}, "id in ?", IDs).Error
	return err
}

// UpdateEnvDetail 更新环境详情记录
// Author [yourname](https://github.com/yourname)
func (edService *EnvDetailService) UpdateEnvDetail(ctx context.Context, ed platform.EnvDetail, projectId int64) (err error) {
	var oldEnvDetail platform.EnvDetail
	err = global.GVA_DB.Model(&oldEnvDetail).Where("id = ?", ed.ID).First(&oldEnvDetail).Error
	if err != nil {
		return err
	}
	if oldEnvDetail.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}

	err = global.GVA_DB.Model(&platform.EnvDetail{}).Where("id = ?", ed.ID).Updates(&ed).Error
	return err
}

// GetEnvDetail 根据ID获取环境详情记录
// Author [yourname](https://github.com/yourname)
func (edService *EnvDetailService) GetEnvDetail(ctx context.Context, ID string) (ed platform.EnvDetail, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&ed).Error
	return
}

// GetEnvDetailInfoList 分页获取环境详情记录
// Author [yourname](https://github.com/yourname)
func (edService *EnvDetailService) GetEnvDetailInfoList(ctx context.Context, info platformReq.EnvDetailSearch) (list []platform.EnvDetail, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&platform.EnvDetail{}).Order("id DESC")
	var eds []platform.EnvDetail
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&eds).Error
	return eds, total, err
}
func (edService *EnvDetailService) GetEnvDetailPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
