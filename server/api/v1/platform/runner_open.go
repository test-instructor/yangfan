package platform

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/common/response"
	"github.com/test-instructor/yangfan/server/v2/model/platform/request"
	"github.com/test-instructor/yangfan/server/v2/model/projectmgr"
	"github.com/test-instructor/yangfan/server/v2/utils/notify"
	"go.uber.org/zap"
)

type OpenRunnerRunRequest struct {
	ProjectId        uint   `json:"projectId" form:"projectId"`
	UUID             string `json:"uuid" form:"uuid"`
	Secret           string `json:"secret" form:"secret"`
	CaseType         string `json:"case_type" form:"case_type"`
	CaseID           uint   `json:"case_id" form:"case_id"`
	EnvID            int    `json:"env_id" form:"env_id"`
	ConfigID         int    `json:"config_id" form:"config_id"`
	NodeName         string `json:"node_name" form:"node_name"`
	RunMode          string `json:"run_mode" form:"run_mode"`
	NotifyEnabled    *bool  `json:"notify_enabled" form:"notify_enabled"`
	NotifyRule       string `json:"notify_rule" form:"notify_rule"`
	NotifyChannelIDs []uint `json:"notify_channel_ids" form:"notify_channel_ids"`
	Failfast         *bool  `json:"failfast" form:"failfast"`
	ResponseMode     string `json:"response_mode" form:"response_mode"`
	CallbackURL      string `json:"callback_url" form:"callback_url"`
	WebhookChannelID uint   `json:"webhook_channel_id" form:"webhook_channel_id"`
	WebhookType      string `json:"webhook_type" form:"webhook_type"`
	WebhookURL       string `json:"webhook_url" form:"webhook_url"`
	WebhookSecret    string `json:"webhook_secret" form:"webhook_secret"`
}

func (r *RunnerApi) OpenRun(c *gin.Context) {
	var queryReq OpenRunnerRunRequest
	_ = c.ShouldBindQuery(&queryReq)

	var bodyReq OpenRunnerRunRequest
	if c.Request.Method != http.MethodGet {
		if err := c.ShouldBindJSON(&bodyReq); err != nil && !errors.Is(err, io.EOF) {
			response.FailWithMessage("参数校验失败: "+err.Error(), c)
			return
		}
	}

	req := mergeOpenRunnerRunRequest(queryReq, bodyReq)

	if req.CaseType == "" {
		response.FailWithMessage("参数校验失败: case_type 不能为空", c)
		return
	}
	if req.CaseID == 0 {
		response.FailWithMessage("参数校验失败: case_id 不能为空", c)
		return
	}
	if req.EnvID == 0 {
		response.FailWithMessage("参数校验失败: env_id 不能为空", c)
		return
	}

	projectID := req.ProjectId
	if v, ok := c.Get("projectId"); ok {
		if vv, ok := v.(uint); ok {
			projectID = vv
		}
	}

	runMode := "CI"
	if req.RunMode != "" {
		runMode = req.RunMode
	}
	if runMode != "CI" {
		runMode = "CI"
	}

	responseMode := req.ResponseMode
	if responseMode == "" {
		responseMode = "sync"
	}
	switch responseMode {
	case "sync":
	case "callback":
		if req.CallbackURL == "" {
			response.FailWithMessage("参数校验失败: callback_url 不能为空", c)
			return
		}
	case "webhook":
		if req.WebhookChannelID != 0 {
			var ch projectmgr.ProjectReportNotifyChannel
			err := global.GVA_DB.Where("id = ? AND project_id = ?", req.WebhookChannelID, projectID).First(&ch).Error
			if err != nil {
				response.FailWithMessage("参数校验失败: webhook_channel_id 不存在", c)
				return
			}
			if !ch.Enabled {
				response.FailWithMessage("参数校验失败: webhook_channel_id 未启用", c)
				return
			}
			if strings.TrimSpace(ch.WebhookURL) == "" {
				response.FailWithMessage("参数校验失败: webhook_channel_id 未配置 webhook_url", c)
				return
			}
			req.WebhookType = strings.ToLower(string(ch.Provider))
			req.WebhookURL = ch.WebhookURL
			req.WebhookSecret = ch.WebhookSecret
			if req.WebhookType != "custom" && req.WebhookType != "feishu" && req.WebhookType != "wecom" && req.WebhookType != "dingtalk" {
				response.FailWithMessage("参数校验失败: webhook_channel_id provider 不支持", c)
				return
			}
		} else {
			if req.WebhookType == "" || req.WebhookURL == "" {
				response.FailWithMessage("参数校验失败: webhook_type/webhook_url 不能为空", c)
				return
			}
		}
	default:
		response.FailWithMessage("参数校验失败: response_mode 不合法", c)
		return
	}

	runnerReq := request.RunnerRequest{
		CaseType:         req.CaseType,
		CaseID:           req.CaseID,
		RunMode:          runMode,
		NodeName:         req.NodeName,
		EnvID:            req.EnvID,
		ConfigID:         req.ConfigID,
		ProjectId:        projectID,
		NotifyEnabled:    req.NotifyEnabled,
		NotifyRule:       req.NotifyRule,
		NotifyChannelIDs: req.NotifyChannelIDs,
	}

	res, err := runnerService.RunTask(runnerReq)
	if err != nil {
		global.GVA_LOG.Error("OpenRun failed", zap.Error(err))
		response.FailWithMessage("任务发送失败: "+err.Error(), c)
		return
	}

	if responseMode == "callback" || responseMode == "webhook" {
		payload := notify.EnqueuedPayload{
			Event:        "ci_run_enqueued",
			ProjectId:    projectID,
			TaskID:       res.TaskID,
			ReportID:     res.ReportID,
			NodeName:     res.NodeName,
			SendTimeUnix: res.SendTime,
			CaseType:     req.CaseType,
			CaseID:       req.CaseID,
			EnvID:        req.EnvID,
			ConfigID:     req.ConfigID,
			RunMode:      runMode,
		}

		go func() {
			ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
			defer cancel()

			var err error
			if responseMode == "callback" {
				err = notify.SendCallback(ctx, req.CallbackURL, payload)
			} else {
				err = notify.SendWebhook(ctx, req.WebhookType, req.WebhookURL, req.WebhookSecret, payload)
			}
			if err != nil {
				global.GVA_LOG.Warn("notify failed", zap.String("mode", responseMode), zap.Error(err))
			}
		}()
	}

	response.OkWithDetailed(res, "success", c)
}

func mergeOpenRunnerRunRequest(a OpenRunnerRunRequest, b OpenRunnerRunRequest) OpenRunnerRunRequest {
	out := a
	if b.ProjectId != 0 {
		out.ProjectId = b.ProjectId
	}
	if b.UUID != "" {
		out.UUID = b.UUID
	}
	if b.Secret != "" {
		out.Secret = b.Secret
	}
	if b.CaseType != "" {
		out.CaseType = b.CaseType
	}
	if b.CaseID != 0 {
		out.CaseID = b.CaseID
	}
	if b.EnvID != 0 {
		out.EnvID = b.EnvID
	}
	if b.ConfigID != 0 {
		out.ConfigID = b.ConfigID
	}
	if b.NodeName != "" {
		out.NodeName = b.NodeName
	}
	if b.RunMode != "" {
		out.RunMode = b.RunMode
	}
	if b.NotifyEnabled != nil {
		out.NotifyEnabled = b.NotifyEnabled
	}
	if b.NotifyRule != "" {
		out.NotifyRule = b.NotifyRule
	}
	if len(b.NotifyChannelIDs) > 0 {
		out.NotifyChannelIDs = b.NotifyChannelIDs
	}
	if b.Failfast != nil {
		out.Failfast = b.Failfast
	}
	if b.ResponseMode != "" {
		out.ResponseMode = b.ResponseMode
	}
	if b.CallbackURL != "" {
		out.CallbackURL = b.CallbackURL
	}
	if b.WebhookChannelID != 0 {
		out.WebhookChannelID = b.WebhookChannelID
	}
	if b.WebhookType != "" {
		out.WebhookType = b.WebhookType
	}
	if b.WebhookURL != "" {
		out.WebhookURL = b.WebhookURL
	}
	if b.WebhookSecret != "" {
		out.WebhookSecret = b.WebhookSecret
	}
	return out
}
