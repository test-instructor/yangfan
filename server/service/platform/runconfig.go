package platform

import (
	"context"
	"errors"

	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	platformReq "github.com/test-instructor/yangfan/server/v2/model/platform/request"
)

type RunConfigService struct{}

// CreateRunConfig 创建运行配置记录
// Author [yourname](https://github.com/yourname)
func (rcService *RunConfigService) CreateRunConfig(ctx context.Context, rc *platform.RunConfig) (err error) {
	err = global.GVA_DB.Create(rc).Error
	return err
}

// DeleteRunConfig 删除运行配置记录
// Author [yourname](https://github.com/yourname)
func (rcService *RunConfigService) DeleteRunConfig(ctx context.Context, ID string, projectId int64) (err error) {
	var rc platform.RunConfig
	err = global.GVA_DB.Where("id = ?", ID).First(&rc).Error
	if err != nil {
		return err
	}
	if rc.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}
	err = global.GVA_DB.Delete(&platform.RunConfig{}, "id = ?", ID).Error
	return err
}

// DeleteRunConfigByIds 批量删除运行配置记录
// Author [yourname](https://github.com/yourname)
func (rcService *RunConfigService) DeleteRunConfigByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]platform.RunConfig{}, "id in ?", IDs).Error
	return err
}

// UpdateRunConfig 更新运行配置记录
// Author [yourname](https://github.com/yourname)
func (rcService *RunConfigService) UpdateRunConfig(ctx context.Context, rc platform.RunConfig, projectId int64) (err error) {
	var oldRunConfig platform.RunConfig
	err = global.GVA_DB.Model(&oldRunConfig).Where("id = ?", rc.ID).First(&oldRunConfig).Error
	if err != nil {
		return err
	}
	if oldRunConfig.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}

	err = global.GVA_DB.Model(&platform.RunConfig{}).Where("id = ?", rc.ID).Updates(&rc).Error
	return err
}

// GetRunConfig 根据ID获取运行配置记录
// Author [yourname](https://github.com/yourname)
func (rcService *RunConfigService) GetRunConfig(ctx context.Context, ID string) (rc platform.RunConfig, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&rc).Error
	return
}

// GetRunConfigInfoList 分页获取运行配置记录
// Author [yourname](https://github.com/yourname)
func (rcService *RunConfigService) GetRunConfigInfoList(ctx context.Context, info platformReq.RunConfigSearch) (list []platform.RunConfig, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&platform.RunConfig{})
	var rcs []platform.RunConfig
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}
	db.Where("project_id = ?", info.ProjectId)
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

	err = db.Find(&rcs).Error
	return rcs, total, err
}
func (rcService *RunConfigService) GetRunConfigPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
