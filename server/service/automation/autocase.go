package automation

import (
	"context"
	"errors"
	"fmt"

	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/automation"
	automationReq "github.com/test-instructor/yangfan/server/v2/model/automation/request"
	"gorm.io/gorm"
)

type AutoCaseService struct{}

// CreateAutoCase 创建测试用例记录
// Author [yourname](https://github.com/yourname)
func (acService *AutoCaseService) CreateAutoCase(ctx context.Context, ac *automation.AutoCase) (err error) {
	err = global.GVA_DB.Create(ac).Error
	return err
}

// DeleteAutoCase 删除测试用例记录
// Author [yourname](https://github.com/yourname)
func (acService *AutoCaseService) DeleteAutoCase(ctx context.Context, ID string, projectId int64) (err error) {
	var ac automation.AutoCase
	err = global.GVA_DB.Where("id = ?", ID).First(&ac).Error
	if err != nil {
		return err
	}
	if ac.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}
	err = global.GVA_DB.Delete(&automation.AutoCase{}, "id = ?", ID).Error
	return err
}

// DeleteAutoCaseByIds 批量删除测试用例记录
// Author [yourname](https://github.com/yourname)
func (acService *AutoCaseService) DeleteAutoCaseByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]automation.AutoCase{}, "id in ?", IDs).Error
	return err
}

// UpdateAutoCase 更新测试用例记录
// Author [yourname](https://github.com/yourname)
func (acService *AutoCaseService) UpdateAutoCase(ctx context.Context, ac automation.AutoCase, projectId int64) (err error) {
	var oldAutoCase automation.AutoCase
	err = global.GVA_DB.Model(&oldAutoCase).Where("id = ?", ac.ID).First(&oldAutoCase).Error
	if err != nil {
		return err
	}
	if oldAutoCase.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}

	err = global.GVA_DB.Model(&automation.AutoCase{}).Where("id = ?", ac.ID).Updates(&ac).Error
	return err
}

// GetAutoCase 根据ID获取测试用例记录
// Author [yourname](https://github.com/yourname)
func (acService *AutoCaseService) GetAutoCase(ctx context.Context, ID string) (ac automation.AutoCase, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&ac).Error
	return
}

// GetAutoCaseInfoList 分页获取测试用例记录
// Author [yourname](https://github.com/yourname)
func (acService *AutoCaseService) GetAutoCaseInfoList(ctx context.Context, info automationReq.AutoCaseSearch) (list []automation.AutoCase, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&automation.AutoCase{})
	var acs []automation.AutoCase
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}
	db.Order("id desc")
	db.Where("project_id = ? ", info.ProjectId)

	if info.CaseName != nil && *info.CaseName != "" {
		db = db.Where("case_name LIKE ?", "%"+*info.CaseName+"%")
	}
	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}
	if info.EnvName != nil && *info.EnvName != "" {
		db = db.Where("env_name LIKE ?", "%"+*info.EnvName+"%")
	}
	if info.ConfigName != nil && *info.ConfigName != "" {
		db = db.Where("config_name LIKE ?", "%"+*info.ConfigName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["created_at"] = true
	orderMap["case_name"] = true
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

	err = db.Find(&acs).Error
	return acs, total, err
}
func (acService *AutoCaseService) GetAutoCasePublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

func (acService *AutoCaseService) AddAutoCaseStep(ctx context.Context, req automationReq.AutoCaseStepReq) (err error) {
	var list automation.AutoCaseStepList
	list.AutoCaseID = req.CaseID
	list.AutoCaseStepID = req.StepID
	list.Sort = 999
	err = global.GVA_DB.Create(&list).Error
	return
}

func (acService *AutoCaseService) SortAutoCaseStep(ctx context.Context, req automationReq.AutoCaseStepSort) (err error) {
	if len(req.Data) == 0 {
		return fmt.Errorf("更新数据不能为空")
	}
	caseStmt := "CASE id "
	ids := make([]uint, 0, len(req.Data))
	var args []interface{}
	for _, step := range req.Data {
		caseStmt += "WHEN ? THEN ? "
		args = append(args, step.ID, step.Sort)
		ids = append(ids, step.ID)
	}
	caseStmt += "END"
	result := global.GVA_DB.WithContext(ctx).Debug().
		Model(&automation.AutoCaseStepList{}).
		Where("id IN (?)", ids).
		Update("sort", gorm.Expr(caseStmt, args...))
	if result.Error != nil {
		return fmt.Errorf("更新失败: %v", result.Error)
	}
	return nil
}

func (acService *AutoCaseService) DelAutoCaseStep(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&automation.AutoCaseStepList{}, "id = ?", id).Error
	return
}

type AutoCaseStep struct {
	IsConfig     bool   `json:"isConfig" form:"isConfig"`
	IsStepConfig bool   `json:"isStepConfig" form:"isStepConfig"`
	ParentId     uint   `json:"parentId" form:"parentId" `
	ID           uint   `json:"id" form:"id" gorm:"column:id;"`
	ConfigName   string `json:"configName" form:"configName"`
	EnvName      string `json:"envName" form:"envName" `
	Name         string `json:"name" form:"name" `
}

func (acService *AutoCaseService) GetAutoCaseSteps(ctx context.Context, caseID string) (list []AutoCaseStep, err error) {
	var acsList []automation.AutoCaseStepList
	err = global.GVA_DB.Preload("AutoCaseStep").Where("case_id = ?", caseID).Order("sort asc").Find(&acsList).Error
	if err != nil {
		return
	}
	for _, acs := range acsList {
		autoCaseStep := AutoCaseStep{}
		autoCaseStep.ID = acs.AutoCaseStepID
		autoCaseStep.ConfigName = acs.AutoCaseStep.ConfigName
		autoCaseStep.EnvName = acs.AutoCaseStep.EnvName
		autoCaseStep.IsConfig = acs.IsConfig
		autoCaseStep.IsStepConfig = acs.IsStepConfig
		autoCaseStep.ParentId = acs.ID
		autoCaseStep.Name = acs.AutoCaseStep.StepName
		list = append(list, autoCaseStep)
	}
	return
}

func (acService *AutoCaseService) SetStepConfig(ctx context.Context, req automationReq.SetStepConfigReq) (err error) {
	err = global.GVA_DB.Model(&automation.AutoCaseStepList{}).Where("id = ?", req.ID).Updates(map[string]interface{}{"is_config": req.IsConfig, "is_step_config": req.IsStepConfig}).Error
	return
}
