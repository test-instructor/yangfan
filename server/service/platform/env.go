package platform

import (
	"context"
	"errors"

	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	platformReq "github.com/test-instructor/yangfan/server/v2/model/platform/request"
)

type EnvService struct{}

// CreateEnv 创建环境配置记录
// Author [yourname](https://github.com/yourname)
func (envService *EnvService) CreateEnv(ctx context.Context, env *platform.Env) (err error) {
	err = global.GVA_DB.Create(env).Error
	return err
}

// DeleteEnv 删除环境配置记录
// Author [yourname](https://github.com/yourname)
func (envService *EnvService) DeleteEnv(ctx context.Context, ID string, projectId int64) (err error) {
	var env platform.Env
	err = global.GVA_DB.Where("id = ?", ID).First(&env).Error
	if err != nil {
		return
	}
	if env.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}

	err = global.GVA_DB.Delete(&platform.Env{}, "id = ?", ID).Error
	return err
}

func Del(ctx context.Context, ID string, projectId int64) (err error) {
	var env platform.Env
	err = global.GVA_DB.Where("id = ?", ID).First(&env).Error
	if err != nil {
		return
	}
	if env.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}
	return
}

// DeleteEnvByIds 批量删除环境配置记录
// Author [yourname](https://github.com/yourname)
func (envService *EnvService) DeleteEnvByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]platform.Env{}, "id in ?", IDs).Error
	return err
}

// UpdateEnv 更新环境配置记录
// Author [yourname](https://github.com/yourname)
func (envService *EnvService) UpdateEnv(ctx context.Context, env platform.Env, projectId int64) (err error) {
	var oldEnv platform.Env
	err = global.GVA_DB.Model(&oldEnv).Where("id = ?", env.ID).First(&oldEnv).Error
	if err != nil {
		return err
	}
	if oldEnv.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}

	err = global.GVA_DB.Model(&platform.Env{}).Where("id = ?", env.ID).Updates(&env).Error
	return err
}

// GetEnv 根据ID获取环境配置记录
// Author [yourname](https://github.com/yourname)
func (envService *EnvService) GetEnv(ctx context.Context, ID string) (env platform.Env, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&env).Error
	return
}

// GetEnvInfoList 分页获取环境配置记录
// Author [yourname](https://github.com/yourname)
func (envService *EnvService) GetEnvInfoList(ctx context.Context, info platformReq.EnvSearch) (list []platform.Env, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&platform.Env{})
	var envs []platform.Env
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.ProjectId != nil {
		db = db.Where("project_id = ?", *info.ProjectId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&envs).Error
	return envs, total, err
}
func (envService *EnvService) GetEnvPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
