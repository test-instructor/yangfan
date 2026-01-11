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

type TimerTaskService struct{}

// CreateTimerTask 创建定时任务记录
// Author [yourname](https://github.com/yourname)
func (tkService *TimerTaskService) CreateTimerTask(ctx context.Context, tk *automation.TimerTask) (err error) {
	err = global.GVA_DB.Create(tk).Error
	return err
}

// DeleteTimerTask 删除定时任务记录
// Author [yourname](https://github.com/yourname)
func (tkService *TimerTaskService) DeleteTimerTask(ctx context.Context, ID string, projectId int64) (err error) {
	var tk automation.TimerTask
	err = global.GVA_DB.Where("id = ?", ID).First(&tk).Error
	if err != nil {
		return err
	}
	if tk.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}
	err = global.GVA_DB.Delete(&automation.TimerTask{}, "id = ?", ID).Error
	return err
}

// DeleteTimerTaskByIds 批量删除定时任务记录
// Author [yourname](https://github.com/yourname)
func (tkService *TimerTaskService) DeleteTimerTaskByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]automation.TimerTask{}, "id in ?", IDs).Error
	return err
}

// UpdateTimerTask 更新定时任务记录
// Author [yourname](https://github.com/yourname)
func (tkService *TimerTaskService) UpdateTimerTask(ctx context.Context, tk automation.TimerTask, projectId int64) (err error) {
	var oldTimerTask automation.TimerTask
	err = global.GVA_DB.Model(&oldTimerTask).Where("id = ?", tk.ID).First(&oldTimerTask).Error
	if err != nil {
		return err
	}
	if oldTimerTask.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}

	err = global.GVA_DB.Model(&automation.TimerTask{}).Where("id = ?", tk.ID).Updates(&tk).Error
	return err
}

// GetTimerTask 根据ID获取定时任务记录
// Author [yourname](https://github.com/yourname)
func (tkService *TimerTaskService) GetTimerTask(ctx context.Context, ID string) (tk automation.TimerTask, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&tk).Error
	return
}

// GetTimerTaskInfoList 分页获取定时任务记录
// Author [yourname](https://github.com/yourname)
func (tkService *TimerTaskService) GetTimerTaskInfoList(ctx context.Context, info automationReq.TimerTaskSearch) (list []automation.TimerTask, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&automation.TimerTask{})
	var tks []automation.TimerTask
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}
	db.Order("id desc")
	db.Where("project_id = ? ", info.ProjectId)

	if info.ConfigName != nil && *info.ConfigName != "" {
		db = db.Where("config_name LIKE ?", "%"+*info.ConfigName+"%")
	}
	if info.EnvName != nil && *info.EnvName != "" {
		db = db.Where("env_name LIKE ?", "%"+*info.EnvName+"%")
	}
	if info.MessageName != nil && *info.MessageName != "" {
		db = db.Where("message_name LIKE ?", "%"+*info.MessageName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&tks).Error
	return tks, total, err
}
func (tkService *TimerTaskService) GetTimerTaskPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// 任务-用例管理
func (tkService *TimerTaskService) AddTimerTaskCase(ctx context.Context, req automationReq.TimerTaskCaseReq) (err error) {
	// 可选：去重，避免同一任务重复引用同一用例
	var cnt int64
	err = global.GVA_DB.WithContext(ctx).
		Model(&automation.TimerTaskCaseList{}).
		Where("task_id = ? AND case_id = ?", req.TaskID, req.CaseID).
		Count(&cnt).Error
	if err != nil {
		return err
	}
	if cnt > 0 {
		return fmt.Errorf("该任务已引用该用例")
	}

	list := automation.TimerTaskCaseList{TimerTaskID: req.TaskID, AutoCaseID: req.CaseID, Sort: 999}
	err = global.GVA_DB.WithContext(ctx).Create(&list).Error
	return
}

func (tkService *TimerTaskService) SortTimerTaskCase(ctx context.Context, req automationReq.TimerTaskCaseSort) (err error) {
	if len(req.Data) == 0 {
		return fmt.Errorf("更新数据不能为空")
	}
	caseStmt := "CASE id "
	ids := make([]uint, 0, len(req.Data))
	var args []interface{}
	for _, item := range req.Data {
		caseStmt += "WHEN ? THEN ? "
		args = append(args, item.ID, item.Sort)
		ids = append(ids, item.ID)
	}
	caseStmt += "END"
	result := global.GVA_DB.WithContext(ctx).
		Model(&automation.TimerTaskCaseList{}).
		Where("id IN (?)", ids).
		Update("sort", gorm.Expr(caseStmt, args...))
	if result.Error != nil {
		return fmt.Errorf("更新失败: %v", result.Error)
	}
	return nil
}

func (tkService *TimerTaskService) DelTimerTaskCase(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.WithContext(ctx).Delete(&automation.TimerTaskCaseList{}, "id = ?", id).Error
	return
}

type TaskCaseItem struct {
	ParentId   uint   `json:"parentId" form:"parentId"`
	ID         uint   `json:"ID" form:"ID"`
	CaseName   string `json:"caseName" form:"caseName"`
	EnvName    string `json:"envName" form:"envName"`
	ConfigName string `json:"configName" form:"configName"`
}

func (tkService *TimerTaskService) GetTimerTaskCases(ctx context.Context, taskID string) (list []TaskCaseItem, err error) {
	var refs []automation.TimerTaskCaseList
	err = global.GVA_DB.Debug().
		Preload("AutoCase").
		Where("task_id = ?", taskID).
		Order("sort asc").
		Find(&refs).Error
	if err != nil {
		return
	}
	for _, r := range refs {
		// 载入 AutoCase 信息
		var ac automation.AutoCase
		if err = global.GVA_DB.Where("id = ?", r.AutoCaseID).First(&ac).Error; err != nil {
			return
		}
		item := TaskCaseItem{
			ParentId:   r.ID,
			ID:         r.AutoCaseID,
			CaseName:   ac.CaseName,
			EnvName:    ac.EnvName,
			ConfigName: ac.ConfigName,
		}
		list = append(list, item)
	}
	return
}

// CreateTag 创建标签
// Author [yourname](https://github.com/yourname)
func (tkService *TimerTaskService) CreateTag(ctx context.Context, tag *automation.TimerTaskTag) (err error) {
	err = global.GVA_DB.Create(tag).Error
	return err
}

// UpdateTag 更新标签
// Author [yourname](https://github.com/yourname)
func (tkService *TimerTaskService) UpdateTag(ctx context.Context, tag automation.TimerTaskTag, projectId int64) (err error) {
	var old automation.TimerTaskTag
	err = global.GVA_DB.Model(&old).Where("id = ?", tag.ID).First(&old).Error
	if err != nil {
		return err
	}
	if old.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}
	err = global.GVA_DB.Model(&automation.TimerTaskTag{}).Where("id = ?", tag.ID).Updates(&tag).Error
	return err
}

// DeleteTag 删除标签
// Author [yourname](https://github.com/yourname)
func (tkService *TimerTaskService) DeleteTag(ctx context.Context, id string, projectId int64) (err error) {
	var old automation.TimerTaskTag
	err = global.GVA_DB.Model(&old).Where("id = ?", id).First(&old).Error
	if err != nil {
		return err
	}
	if old.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}
	err = global.GVA_DB.Delete(&automation.TimerTaskTag{}, "id = ?", id).Error
	return err
}

// GetTagList 分页获取标签列表
// Author [yourname](https://github.com/yourname)
func (tkService *TimerTaskService) GetTagList(ctx context.Context, info automationReq.TagSearch) (list []automation.TimerTaskTag, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&automation.TimerTaskTag{})
	var tags []automation.TimerTaskTag
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}
	db = db.Where("project_id = ?", info.ProjectId)
	if info.Name != nil && *info.Name != "" {
		db = db.Where("name LIKE ?", "%"+*info.Name+"%")
	}
	db = db.Order("id desc")
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Find(&tags).Error
	return tags, total, err
}
