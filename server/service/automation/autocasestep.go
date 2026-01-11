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

type AutoCaseStepService struct{}

// CreateAutoCaseStep 创建测试步骤记录
// Author [yourname](https://github.com/yourname)
func (acsService *AutoCaseStepService) CreateAutoCaseStep(ctx context.Context, acs *automation.AutoCaseStep) (err error) {
	err = global.GVA_DB.Create(acs).Error
	return err
}

func (acsService *AutoCaseStepService) AddAutoCaseStepApi(ctx context.Context, acs *automationReq.AutoCaseStepSearchApi) (data map[string]interface{}, err error) {
	var as automation.AutoStep
	err = global.GVA_DB.Model(&automation.AutoStep{}).Preload("Request").Where("id = ?", acs.ApiID).First(&as).Error
	if err != nil {
		return
	}
	as.ID = 0
	as.Request.ID = 0
	as.StepType = 11
	err = global.GVA_DB.Create(&as).Error
	if err != nil {
		return
	}
	var acsRequest automation.AutoCaseStepRelation
	acsRequest.AutoCaseStepID = acs.ID
	acsRequest.AutoStepID = as.ID
	acsRequest.Sort = 999
	err = global.GVA_DB.Create(&acsRequest).Error
	if err != nil {
		return
	}
	data = map[string]interface{}{
		"id": as.ID,
	}
	return
}

func (acsService *AutoCaseStepService) SortAutoCaseStepApi(ctx context.Context, acs *automationReq.AutoCaseStepSearchApi) (err error) {
	// 1. 校验输入数据是否为空
	if len(acs.Data) == 0 {
		return fmt.Errorf("更新数据不能为空")
	}

	// 2. 提取ID和构建CASE语句
	caseStmt := "CASE auto_step_id "
	ids := make([]uint, 0, len(acs.Data))
	var args []interface{}
	idMap := make(map[uint]uint) // 用于记录ID与目标sort的映射，方便后续校验

	for _, step := range acs.Data {
		caseStmt += "WHEN ? THEN ? "
		args = append(args, step.ID, step.Sort)
		ids = append(ids, step.ID)
		idMap[step.ID] = step.Sort // 暂存ID对应的目标sort
	}
	caseStmt += "END"

	// 3. 先查询当前ID的sort值，确认是否需要更新（可选，用于排查问题）
	var currentSteps []struct {
		ID   uint
		Sort uint
	}
	err = global.GVA_DB.WithContext(ctx).
		Model(&automation.AutoCaseStepRelation{}).
		Where("auto_step_id IN (?)", ids).
		Find(&currentSteps).Error
	if err != nil {
		return fmt.Errorf("查询当前数据失败: %v", err)
	}
	// 检查是否有ID不存在
	if len(currentSteps) != len(ids) {
		// 找出不存在的ID
		existIDs := make(map[uint]bool)
		for _, s := range currentSteps {
			existIDs[s.ID] = true
		}
		var notExistIDs []uint
		for _, id := range ids {
			if !existIDs[id] {
				notExistIDs = append(notExistIDs, id)
			}
		}
		return fmt.Errorf("以下ID不存在或已被删除: %v", notExistIDs)
	}

	// 4. 执行更新
	result := global.GVA_DB.WithContext(ctx).Debug().
		Model(&automation.AutoCaseStepRelation{}).
		Where("auto_step_id IN (?)", ids).
		Update("sort", gorm.Expr(caseStmt, args...))
	if result.Error != nil {
		return fmt.Errorf("更新失败: %v", result.Error)
	}

	// 5. 输出更新结果日志（方便排查）
	if result.RowsAffected == 0 {
		// 检查是否因为值未变化导致
		allSame := true
		for _, s := range currentSteps {
			if idMap[s.ID] != s.Sort {
				allSame = false
				break
			}
		}
		if allSame {
			fmt.Printf("提示: 所有ID的sort值与目标值一致，无需更新\n")
		} else {
			return fmt.Errorf("更新成功但未影响任何行，可能存在数据不一致")
		}
	} else {
		fmt.Printf("更新成功，影响行数: %d\n", result.RowsAffected)
	}

	return nil
}

// DeleteAutoCaseStep 删除测试步骤记录
// Author [yourname](https://github.com/yourname)
func (acsService *AutoCaseStepService) DeleteAutoCaseStep(ctx context.Context, ID string, projectId int64) (err error) {
	var acs automation.AutoCaseStep
	err = global.GVA_DB.Where("id = ?", ID).First(&acs).Error
	if err != nil {
		return err
	}
	if acs.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}
	err = global.GVA_DB.Delete(&automation.AutoCaseStep{}, "id = ?", ID).Error
	return err
}

func (acsService *AutoCaseStepService) DeleteAutoCaseStepApi(ctx context.Context, ID string, projectId int64) (err error) {
	var acs automation.AutoCaseStepRelation
	err = global.GVA_DB.Where("auto_step_id = ?", ID).First(&acs).Error
	if err != nil {
		return err
	}

	err = global.GVA_DB.Delete(&automation.AutoCaseStepRelation{}, "auto_step_id = ?", ID).Error
	if err != nil {
		return err
	}
	err = global.GVA_DB.Delete(&automation.AutoStep{}, "id = ?", ID).Error
	return nil
}

// DeleteAutoCaseStepByIds 批量删除测试步骤记录
// Author [yourname](https://github.com/yourname)
func (acsService *AutoCaseStepService) DeleteAutoCaseStepByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]automation.AutoCaseStep{}, "id in ?", IDs).Error
	return err
}

// UpdateAutoCaseStep 更新测试步骤记录
// Author [yourname](https://github.com/yourname)
func (acsService *AutoCaseStepService) UpdateAutoCaseStep(ctx context.Context, acs automation.AutoCaseStep, projectId int64) (err error) {
	var oldAutoCaseStep automation.AutoCaseStep
	err = global.GVA_DB.Model(&oldAutoCaseStep).Where("id = ?", acs.ID).First(&oldAutoCaseStep).Error
	if err != nil {
		return err
	}
	if oldAutoCaseStep.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}

	err = global.GVA_DB.Model(&automation.AutoCaseStep{}).Where("id = ?", acs.ID).Updates(&acs).Error
	return err
}

// GetAutoCaseStep 根据ID获取测试步骤记录
// Author [yourname](https://github.com/yourname)
func (acsService *AutoCaseStepService) GetAutoCaseStep(ctx context.Context, ID string) (acs automation.AutoCaseStep, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&acs).Error
	return
}

func (acsService *AutoCaseStepService) GetAutoCaseStepApi(ctx context.Context, ID string) (setpApiList []automation.AutoStep, err error) {
	var acsps []automation.AutoCaseStepRelation
	err = global.GVA_DB.Model(&automation.AutoCaseStepRelation{}).Preload("AutoStep.Request").
		Where("auto_case_step_id = ?", ID).Order("sort ASC").
		Find(&acsps).Error
	if err != nil {
		return
	}
	for _, v := range acsps {
		setpApiList = append(setpApiList, v.AutoStep)
	}
	return
}

// GetAutoCaseStepInfoList 分页获取测试步骤记录
// Author [yourname](https://github.com/yourname)
func (acsService *AutoCaseStepService) GetAutoCaseStepInfoList(ctx context.Context, info automationReq.AutoCaseStepSearch) (list []automation.AutoCaseStep, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&automation.AutoCaseStep{})
	var acss []automation.AutoCaseStep
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.StepName != nil && *info.StepName != "" {
		db = db.Where("step_name LIKE ?", "%"+*info.StepName+"%")
	}
	if info.Menu != 0 {
		db = db.Where("menu = ?", info.Menu)
	}
	db.Order("id desc")
	db.Where("project_id = ? ", info.ProjectId)
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["created_at"] = true
	orderMap["step_name"] = true
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

	err = db.Find(&acss).Error
	return acss, total, err
}
func (acsService *AutoCaseStepService) GetAutoCaseStepPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
