package automation

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/automation"
	automationReq "github.com/test-instructor/yangfan/server/v2/model/automation/request"
	"gorm.io/gorm"
)

type AutoCaseStepService struct{}

// CreateAutoCaseStep 创建测试步骤记录
// Author [yourname](https://github.com/yourname)
func (acsService *AutoCaseStepService) CreateAutoCaseStep(ctx context.Context, acs *automation.AutoCaseStep) (err error) {
	if strings.TrimSpace(acs.Type) == "" {
		acs.Type = "api"
	}
	err = global.GVA_DB.Create(acs).Error
	return err
}

func (acsService *AutoCaseStepService) AddAutoCaseStepApi(ctx context.Context, acs *automationReq.AutoCaseStepSearchApi) (data map[string]interface{}, err error) {
	if acs == nil {
		return nil, fmt.Errorf("请求参数不能为空")
	}
	if acs.ID == 0 {
		return nil, fmt.Errorf("步骤集合ID不能为空")
	}
	if acs.ApiID == 0 {
		return nil, fmt.Errorf("api_id不能为空")
	}

	db := global.GVA_DB.WithContext(ctx)
	var createdStepID uint
	err = db.Transaction(func(tx *gorm.DB) error {
		var src automation.AutoStep
		if e := tx.Preload("Request").First(&src, "id = ?", acs.ApiID).Error; e != nil {
			return e
		}

		hasRequest := src.Request != nil
		hasAndroid := src.Android != nil && len(src.Android.Actions) > 0
		hasIOS := src.IOS != nil && len(src.IOS.Actions) > 0
		hasHarmony := src.Harmony != nil && len(src.Harmony.Actions) > 0
		hasBrowser := src.Browser != nil && len(src.Browser.Actions) > 0
		if !hasRequest && !hasAndroid && !hasIOS && !hasHarmony && !hasBrowser {
			if src.RequestID != 0 {
				return fmt.Errorf("该步骤的HTTP Request不存在或已被删除")
			}
			return fmt.Errorf("该步骤未配置 Request/Android/IOS/Harmony/Browser，无法添加为接口步骤")
		}

		var reqCopyID uint
		if src.Request != nil {
			reqCopy := *src.Request
			reqCopy.ID = 0
			reqCopy.CreatedAt = time.Time{}
			reqCopy.UpdatedAt = time.Time{}
			reqCopy.DeletedAt = gorm.DeletedAt{}
			reqCopy.Request = nil
			reqCopy.RequestID = 0
			if e := tx.Create(&reqCopy).Error; e != nil {
				return e
			}
			reqCopyID = reqCopy.ID
		}

		stepCopy := src
		stepCopy.ID = 0
		stepCopy.CreatedAt = time.Time{}
		stepCopy.UpdatedAt = time.Time{}
		stepCopy.DeletedAt = gorm.DeletedAt{}
		stepCopy.StepType = 11
		stepCopy.Request = nil
		stepCopy.RequestID = reqCopyID
		if e := tx.Create(&stepCopy).Error; e != nil {
			return e
		}

		sort := uint(999)
		if acs.Sort != 0 {
			sort = acs.Sort
		}
		rel := automation.AutoCaseStepRelation{
			AutoCaseStepID: acs.ID,
			AutoStepID:     stepCopy.ID,
			Sort:           sort,
			ProjectId:      stepCopy.ProjectId,
		}
		if e := tx.Create(&rel).Error; e != nil {
			return e
		}

		createdStepID = stepCopy.ID
		return nil
	})
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{"id": createdStepID}, nil
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
	err = global.GVA_DB.Model(&automation.AutoCaseStepRelation{}).
		Preload("AutoStep").
		Preload("AutoStep.Request").
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
	db = db.Where("project_id = ? ", info.ProjectId)
	t := ""
	if info.Type != nil {
		t = strings.TrimSpace(*info.Type)
	}
	if t == "" || strings.EqualFold(t, "api") {
		db = db.Where("(type = ? OR type = '' OR type IS NULL)", "api")
	} else if t != "" {
		db = db.Where("type = ?", t)
	}
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
