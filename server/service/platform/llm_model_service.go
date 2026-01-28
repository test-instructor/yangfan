package platform

import (
	"context"
	"errors"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	platformReq "github.com/test-instructor/yangfan/server/v2/model/platform/request"
)

type LLMModelConfigService struct{}

// CreateLLMModelConfig 创建大语言模型配置记录
// Author [yourname](https://github.com/yourname)
func (llmconfigService *LLMModelConfigService) CreateLLMModelConfig(ctx context.Context, llmconfig *platform.LLMModelConfig) (err error) {
	err = global.GVA_DB.Create(llmconfig).Error
	return err
}

// DeleteLLMModelConfig 删除大语言模型配置记录
// Author [yourname](https://github.com/yourname)
func (llmconfigService *LLMModelConfigService) DeleteLLMModelConfig(ctx context.Context, ID string, projectId int64) (err error) {
	var llmconfig platform.LLMModelConfig
	err = global.GVA_DB.Where("id = ?", ID).First(&llmconfig).Error
	if err != nil {
		return err
	}
	if llmconfig.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}
	err = global.GVA_DB.Delete(&platform.LLMModelConfig{}, "id = ?", ID).Error
	return err
}

// DeleteLLMModelConfigByIds 批量删除大语言模型配置记录
// Author [yourname](https://github.com/yourname)
func (llmconfigService *LLMModelConfigService) DeleteLLMModelConfigByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]platform.LLMModelConfig{}, "id in ?", IDs).Error
	return err
}

// UpdateLLMModelConfig 更新大语言模型配置记录
// Author [yourname](https://github.com/yourname)
func (llmconfigService *LLMModelConfigService) UpdateLLMModelConfig(ctx context.Context, llmconfig platform.LLMModelConfig, projectId int64) (err error) {
	var oldLLMModelConfig platform.LLMModelConfig
	err = global.GVA_DB.Model(&oldLLMModelConfig).Where("id = ?", llmconfig.ID).First(&oldLLMModelConfig).Error
	if err != nil {
		return err
	}
	if oldLLMModelConfig.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}

	err = global.GVA_DB.Model(&platform.LLMModelConfig{}).Where("id = ?", llmconfig.ID).Updates(&llmconfig).Error
	return err
}

// GetLLMModelConfig 根据ID获取大语言模型配置记录
// Author [yourname](https://github.com/yourname)
func (llmconfigService *LLMModelConfigService) GetLLMModelConfig(ctx context.Context, ID string) (llmconfig platform.LLMModelConfig, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&llmconfig).Error
	return
}

// GetLLMModelConfigInfoList 分页获取大语言模型配置记录
// Author [yourname](https://github.com/yourname)
func (llmconfigService *LLMModelConfigService) GetLLMModelConfigInfoList(ctx context.Context, info platformReq.LLMModelConfigSearch) (list []platform.LLMModelConfig, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&platform.LLMModelConfig{})
	var llmconfigs []platform.LLMModelConfig
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}
	db.Order("id desc")
	db.Where("project_id = ? ", info.ProjectId)

	if info.Name != nil && *info.Name != "" {
		db = db.Where("name LIKE ?", "%"+*info.Name+"%")
	}
	if info.RequestSchema != nil && *info.RequestSchema != "" {
		db = db.Where("request_schema = ?", *info.RequestSchema)
	}
	if info.Model != nil && *info.Model != "" {
		db = db.Where("model = ?", *info.Model)
	}
	if info.ReasoningEffort != nil && *info.ReasoningEffort != "" {
		db = db.Where("reasoning_effort = ?", *info.ReasoningEffort)
	}
	if info.Enabled != nil {
		db = db.Where("enabled = ?", *info.Enabled)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&llmconfigs).Error
	return llmconfigs, total, err
}
func (llmconfigService *LLMModelConfigService) GetLLMModelConfigPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
