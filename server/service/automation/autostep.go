package automation

import (
	"context"
	"errors"

	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/automation"
	automationReq "github.com/test-instructor/yangfan/server/v2/model/automation/request"
)

type AutoStepService struct{}

// CreateAutoStep 创建自动化步骤记录
// Author [yourname](https://github.com/yourname)
func (asService *AutoStepService) CreateAutoStep(ctx context.Context, as *automation.AutoStep) (err error) {
	as.StepType = 1
	err = global.GVA_DB.Create(as).Error
	return err
}

// DeleteAutoStep 删除自动化步骤记录
// Author [yourname](https://github.com/yourname)
func (asService *AutoStepService) DeleteAutoStep(ctx context.Context, ID string, projectId int64) (err error) {
	var as automation.AutoStep
	err = global.GVA_DB.Where("id = ?", ID).First(&as).Error
	if err != nil {
		return err
	}
	if as.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}
	err = global.GVA_DB.Delete(&automation.AutoStep{}, "id = ?", ID).Error
	return err
}

// DeleteAutoStepByIds 批量删除自动化步骤记录
// Author [yourname](https://github.com/yourname)
func (asService *AutoStepService) DeleteAutoStepByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]automation.AutoStep{}, "id in ?", IDs).Error
	return err
}

// UpdateAutoStep 更新自动化步骤记录
// Author [yourname](https://github.com/yourname)
func (asService *AutoStepService) UpdateAutoStep(ctx context.Context, as automation.AutoStep, projectId int64) (err error) {
	var oldAutoStep automation.AutoStep
	err = global.GVA_DB.Model(&oldAutoStep).Where("id = ?", as.ID).First(&oldAutoStep).Error
	if err != nil {
		return err
	}
	if oldAutoStep.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}

	err = global.GVA_DB.Model(&automation.Request{}).Where("id = ?", as.RequestID).Save(as.Request).Error
	if err != nil {
		return err
	}
	err = global.GVA_DB.Model(&automation.AutoStep{}).Where("id = ?", as.ID).Updates(&as).Error

	return err
}

// GetAutoStep 根据ID获取自动化步骤记录
// Author [yourname](https://github.com/yourname)
func (asService *AutoStepService) GetAutoStep(ctx context.Context, ID string) (as automation.AutoStep, err error) {
	err = global.GVA_DB.Model(&automation.AutoStep{}).Preload("Request").Where("id = ?", ID).First(&as).Error
	return
}

// GetAutoStepInfoList 分页获取自动化步骤记录
// Author [yourname](https://github.com/yourname)
func (asService *AutoStepService) GetAutoStepInfoList(ctx context.Context, info automationReq.AutoStepSearch) (list []automation.AutoStep, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&automation.AutoStep{})
	var ass []automation.AutoStep
	if info.StepName != nil && *info.StepName != "" {
		db = db.Where("step_name LIKE ?", "%"+*info.StepName+"%")
	}
	if info.Menu != nil && *info.Menu != 0 {
		db = db.Where("menu = ?", info.Menu)
	}
	var stepType = 1
	if info.StepType != nil && *info.StepType != 0 {
		stepType = *info.StepType
	}
	db = db.Where("step_type = ?", stepType)
	db.Order("id desc")
	db = db.Where("project_id = ?", info.ProjectId)
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	db.Preload("Request")
	err = db.Find(&ass).Error
	return ass, total, err
}
func (asService *AutoStepService) GetAutoStepPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
