package platform

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/automation"
	"github.com/test-instructor/yangfan/server/v2/model/platform/request"
	"go.uber.org/zap"
)

type RunnerService struct{}

var (
	validCaseTypes = map[string]bool{
		"接口":   true,
		"步骤":   true,
		"用例":   true,
		"任务":   true,
		"标签":   true,
		"api":  true,
		"step": true,
		"case": true,
		"task": true,
		"tag":  true,
	}
	validRunModes = map[string]bool{
		"调试模式":                true,
		"保存并调试":               true,
		"立即运行":                true,
		"后台运行":                true,
		"定时执行":                true,
		"CI":                  true,
		"Debug":               true,
		"Save & Debug":        true,
		"Run Immediately":     true,
		"Background Run":      true,
		"Scheduled Execution": true,
	}
)

// RunTask 运行任务
func (s *RunnerService) RunTask(req request.RunnerRequest) (*request.RunnerResponse, error) {
	// 1. Validate Enums
	if !validCaseTypes[req.CaseType] {
		return nil, errors.New("invalid case_type")
	}
	if !validRunModes[req.RunMode] {
		return nil, errors.New("invalid run_mode")
	}

	// 2. Generate Task ID
	taskID := uuid.New().String()

	// 3. 创建测试报告（状态为待运行）
	report, err := s.CreatePendingReport(req)
	if err != nil {
		global.GVA_LOG.Error("创建测试报告失败", zap.Error(err))
		return nil, errors.New("创建测试报告失败: " + err.Error())
	}

	// 4. Send to MQ (携带报告ID)
	if global.GVA_MQ_PRODUCER == nil {
		global.GVA_LOG.Error("MQ未初始化，无法运行任务")
		return nil, errors.New("MQ未初始化，无法运行任务")
	}
	nodeName, err := global.GVA_MQ_PRODUCER.SendTask(
		taskID,
		report.ID,
		req.CaseType,
		req.CaseID,
		req.RunMode,
		req.NodeName,
		req.EnvID,
		req.ConfigID,
		req.ProjectId,
		req.NotifyEnabled,
		req.NotifyRule,
		req.NotifyChannelIDs,
		req.Failfast,
	)

	if err != nil {
		return nil, err
	}

	if uerr := global.GVA_DB.Model(&automation.AutoReport{}).Where("id = ?", report.ID).Update("node_name", nodeName).Error; uerr != nil {
		global.GVA_LOG.Warn("写入测试报告 node_name 失败", zap.Uint("report_id", report.ID), zap.String("node_name", nodeName), zap.Error(uerr))
	}

	return &request.RunnerResponse{
		TaskID:   taskID,
		ReportID: report.ID,
		NodeName: nodeName,
		SendTime: time.Now().Unix(),
	}, nil
}

// CreatePendingReport 根据请求参数创建待运行状态的测试报告
func (s *RunnerService) CreatePendingReport(req request.RunnerRequest) (*automation.AutoReport, error) {
	var name string
	var projectId int64

	// 根据不同的 CaseType 查询对应的名称和项目ID
	switch req.CaseType {
	case "接口", "api":
		var step automation.AutoStep
		if err := global.GVA_DB.Model(&automation.AutoStep{}).Select("step_name, project_id").First(&step, "id = ?", req.CaseID).Error; err != nil {
			return nil, errors.New("获取接口信息失败")
		}
		name = step.StepName
		projectId = step.ProjectId

	case "步骤", "step":
		var step automation.AutoCaseStep
		if err := global.GVA_DB.Model(&automation.AutoCaseStep{}).Select("step_name, project_id").First(&step, "id = ?", req.CaseID).Error; err != nil {
			return nil, errors.New("获取步骤信息失败")
		}
		name = step.StepName
		projectId = step.ProjectId

	case "用例", "case":
		var autoCase automation.AutoCase
		if err := global.GVA_DB.Model(&automation.AutoCase{}).Select("case_name, project_id").First(&autoCase, "id = ?", req.CaseID).Error; err != nil {
			return nil, errors.New("获取用例信息失败")
		}
		name = autoCase.CaseName
		projectId = autoCase.ProjectId

	case "任务", "task":
		var task automation.TimerTask
		if err := global.GVA_DB.Model(&automation.TimerTask{}).
			Select("name, project_id, message_id, notify_enabled, notify_rule").
			First(&task, "id = ?", req.CaseID).Error; err != nil {
			return nil, errors.New("获取任务信息失败")
		}
		if task.Name != nil {
			name = *task.Name
		}
		projectId = task.ProjectId
		if req.NotifyEnabled == nil && req.NotifyRule == "" && len(req.NotifyChannelIDs) == 0 {
			if task.NotifyEnabled != nil {
				req.NotifyEnabled = task.NotifyEnabled
			}
			if task.NotifyRule != nil {
				req.NotifyRule = *task.NotifyRule
			}
			if task.MessageID != nil {
				req.NotifyChannelIDs = []uint{*task.MessageID}
			}
		}
		if req.Failfast == nil && task.Failfast != nil {
			req.Failfast = task.Failfast
		}

	case "标签", "tag":
		var tag automation.TimerTaskTag
		if err := global.GVA_DB.Model(&automation.TimerTaskTag{}).Select("name, project_id").First(&tag, "id = ?", req.CaseID).Error; err != nil {
			return nil, errors.New("获取标签信息失败")
		}
		name = tag.Name
		projectId = tag.ProjectId

	default:
		return nil, errors.New("unsupported case_type")
	}

	// 如果请求中有 projectId，使用请求中的
	if req.ProjectId != 0 {
		projectId = int64(req.ProjectId)
	}

	// 创建报告，状态为待运行
	pendingStatus := int64(automation.ReportStatusPending)

	notifyEnabled := req.NotifyEnabled
	if notifyEnabled == nil && len(req.NotifyChannelIDs) > 0 {
		v := true
		notifyEnabled = &v
	}
	notifyChannelBytes := []byte("[]")
	if len(req.NotifyChannelIDs) > 0 {
		if b, err := json.Marshal(req.NotifyChannelIDs); err == nil {
			notifyChannelBytes = b
		}
	}
	var nodeName *string
	if req.NodeName != "" {
		v := req.NodeName
		nodeName = &v
	}
	report := &automation.AutoReport{
		Name:             &name,
		Status:           &pendingStatus,
		ProjectId:        projectId,
		CaseType:         req.CaseType,
		RunMode:          req.RunMode,
		ConfigID:         req.ConfigID,
		EnvID:            req.EnvID,
		CaseID:           req.CaseID,
		NodeName:         nodeName,
		NotifyEnabled:    notifyEnabled,
		NotifyRule:       req.NotifyRule,
		NotifyChannelIDs: notifyChannelBytes,
	}

	if err := global.GVA_DB.Create(report).Error; err != nil {
		return nil, err
	}

	return report, nil
}
