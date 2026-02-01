package automation

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/automation"
	automationReq "github.com/test-instructor/yangfan/server/v2/model/automation/request"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	"gorm.io/gorm"
)

type TimerTaskService struct{}

// CreateTimerTask 创建定时任务记录
// Author [yourname](https://github.com/yourname)
func (tkService *TimerTaskService) CreateTimerTask(ctx context.Context, tk *automation.TimerTask) (err error) {
	if tk.Type == nil || *tk.Type == "" {
		t := "api"
		tk.Type = &t
	}
	if err := tkService.validateTimerTaskRunnerNode(ctx, tk.ProjectId, tk.RunnerNodeName); err != nil {
		return err
	}
	err = global.GVA_DB.Create(tk).Error
	if err == nil {
		RefreshTimerTaskSchedules(ctx)
		tkService.publishTimerTaskControl(ctx, tk.ID, tk.RunnerNodeName)
	}
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
	tkService.publishTimerTaskControlDelete(ctx, tk.ID, tk.RunnerNodeName)
	err = global.GVA_DB.Delete(&automation.TimerTask{}, "id = ?", ID).Error
	if err == nil {
		RefreshTimerTaskSchedules(ctx)
	}
	return err
}

// DeleteTimerTaskByIds 批量删除定时任务记录
// Author [yourname](https://github.com/yourname)
func (tkService *TimerTaskService) DeleteTimerTaskByIds(ctx context.Context, IDs []string) (err error) {
	var tasks []automation.TimerTask
	_ = global.GVA_DB.WithContext(ctx).
		Select("id", "runner_node_name").
		Find(&tasks, "id in ?", IDs).Error
	for _, t := range tasks {
		tkService.publishTimerTaskControlDelete(ctx, t.ID, t.RunnerNodeName)
	}
	err = global.GVA_DB.Delete(&[]automation.TimerTask{}, "id in ?", IDs).Error
	if err == nil {
		RefreshTimerTaskSchedules(ctx)
	}
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

	validateProjectID := oldTimerTask.ProjectId
	if tk.ProjectId != 0 {
		validateProjectID = tk.ProjectId
	}
	if err := tkService.validateTimerTaskRunnerNode(ctx, validateProjectID, tk.RunnerNodeName); err != nil {
		return err
	}

	err = global.GVA_DB.Model(&automation.TimerTask{}).Where("id = ?", tk.ID).Updates(&tk).Error
	if err == nil {
		RefreshTimerTaskSchedules(ctx)
		var current automation.TimerTask
		_ = global.GVA_DB.WithContext(ctx).
			Select("id", "run_time", "status", "runner_node_name").
			First(&current, "id = ?", tk.ID).Error

		oldNode := ""
		if oldTimerTask.RunnerNodeName != nil {
			oldNode = strings.TrimSpace(*oldTimerTask.RunnerNodeName)
		}
		newNode := ""
		if current.RunnerNodeName != nil {
			newNode = strings.TrimSpace(*current.RunnerNodeName)
		}

		if oldNode != "" && oldNode != newNode {
			tkService.publishTimerTaskControlDelete(ctx, tk.ID, oldTimerTask.RunnerNodeName)
		}
		tkService.publishTimerTaskControl(ctx, tk.ID, current.RunnerNodeName)
	}
	return err
}

func (tkService *TimerTaskService) validateTimerTaskRunnerNode(ctx context.Context, projectID int64, runnerNodeName *string) error {
	if runnerNodeName == nil {
		return nil
	}
	node := strings.TrimSpace(*runnerNodeName)
	if node == "" {
		return nil
	}
	if projectID == 0 {
		return nil
	}

	var rn platform.RunnerNode
	err := global.GVA_DB.WithContext(ctx).
		Model(&platform.RunnerNode{}).
		Select("run_content").
		Where("node_name = ?", node).
		First(&rn).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("运行节点不存在: %s", node)
	}
	if err != nil {
		return err
	}
	if rn.RunContent == nil || strings.TrimSpace(*rn.RunContent) == "" {
		return nil
	}
	switch strings.ToLower(strings.TrimSpace(*rn.RunContent)) {
	case "timer", "all":
		return nil
	default:
		return fmt.Errorf("运行节点不支持定时任务，仅支持 timer/all: %s", node)
	}
}

func (tkService *TimerTaskService) publishTimerTaskControl(ctx context.Context, taskID uint, runnerNodeName *string) {
	if global.GVA_MQ_TIMER_PRODUCER == nil {
		return
	}
	if runnerNodeName == nil {
		return
	}
	node := strings.TrimSpace(*runnerNodeName)
	if node == "" {
		return
	}

	var t automation.TimerTask
	if err := global.GVA_DB.WithContext(ctx).
		Select("id", "run_time", "status", "runner_node_name").
		First(&t, "id = ?", taskID).Error; err != nil {
		return
	}

	enabled := t.Status != nil && *t.Status
	hasSpec := t.RunTime != nil && strings.TrimSpace(*t.RunTime) != ""
	action := "delete"
	if enabled && hasSpec {
		action = "upsert"
	}
	_ = global.GVA_MQ_TIMER_PRODUCER.Send(action, taskID, node)
}

func (tkService *TimerTaskService) publishTimerTaskControlDelete(ctx context.Context, taskID uint, runnerNodeName *string) {
	if global.GVA_MQ_TIMER_PRODUCER == nil {
		return
	}
	if runnerNodeName == nil {
		return
	}
	node := strings.TrimSpace(*runnerNodeName)
	if node == "" {
		return
	}
	_ = global.GVA_MQ_TIMER_PRODUCER.Send("delete", taskID, node)
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
	db = db.Where("project_id = ? ", info.ProjectId)

	if info.ConfigName != nil && *info.ConfigName != "" {
		db = db.Where("config_name LIKE ?", "%"+*info.ConfigName+"%")
	}
	if info.EnvName != nil && *info.EnvName != "" {
		db = db.Where("env_name LIKE ?", "%"+*info.EnvName+"%")
	}
	if info.MessageName != nil && *info.MessageName != "" {
		db = db.Where("message_name LIKE ?", "%"+*info.MessageName+"%")
	}
	if info.Type != nil {
		t := strings.TrimSpace(*info.Type)
		if t == "" || strings.EqualFold(t, "api") {
			db = db.Where("(type = ? OR type = '' OR type IS NULL)", "api")
		} else {
			db = db.Where("type = ?", t)
		}
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
