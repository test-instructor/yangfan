package projectmgr

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/automation"
	"github.com/test-instructor/yangfan/server/v2/model/projectmgr"
	projectmgrReq "github.com/test-instructor/yangfan/server/v2/model/projectmgr/request"
	"github.com/test-instructor/yangfan/server/v2/utils/notify"
	"go.uber.org/zap"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ReportNotifyService struct{}

type AutoReportNotifyChannelStatus struct {
	ChannelId uint   `json:"channelId"`
	Provider  string `json:"provider"`
	Name      string `json:"name"`
	SendRule  string `json:"send_rule"`
	State     string `json:"state"`
	Ok        *bool  `json:"ok"`
	Error     string `json:"error"`
	Display   string `json:"display"`
}

type AutoReportNotifyStatusResponse struct {
	ReportId     uint                            `json:"reportId"`
	ProjectId    int64                           `json:"projectId"`
	ReportStatus int64                           `json:"reportStatus"`
	Result       string                          `json:"result"`
	Items        []AutoReportNotifyChannelStatus `json:"items"`
}

func (s *ReportNotifyService) CreateReportNotifyChannel(ctx context.Context, channel *projectmgr.ProjectReportNotifyChannel, projectId int64) error {
	if channel == nil {
		return errors.New("invalid channel")
	}
	channel.ProjectId = projectId
	channel.Provider = projectmgr.ReportNotifyProvider(strings.ToLower(string(channel.Provider)))
	channel.SendRule = projectmgr.ReportNotifySendRule(strings.ToLower(string(channel.SendRule)))
	if channel.Provider == "" {
		return errors.New("provider required")
	}
	if channel.Name == "" {
		return errors.New("name required")
	}
	if channel.SendRule == "" {
		channel.SendRule = projectmgr.ReportNotifySendRuleAlways
	}
	return global.GVA_DB.Create(channel).Error
}

func (s *ReportNotifyService) DeleteReportNotifyChannel(ctx context.Context, id uint, projectId int64) error {
	return global.GVA_DB.Where("id = ? AND project_id = ?", id, projectId).Delete(&projectmgr.ProjectReportNotifyChannel{}).Error
}

func (s *ReportNotifyService) DeleteReportNotifyChannelByIds(ctx context.Context, ids []uint, projectId int64) error {
	if len(ids) == 0 {
		return nil
	}
	return global.GVA_DB.Where("project_id = ?", projectId).Delete(&[]projectmgr.ProjectReportNotifyChannel{}, "id in ?", ids).Error
}

func (s *ReportNotifyService) UpdateReportNotifyChannel(ctx context.Context, channel projectmgr.ProjectReportNotifyChannel, projectId int64) error {
	channel.Provider = projectmgr.ReportNotifyProvider(strings.ToLower(string(channel.Provider)))
	channel.SendRule = projectmgr.ReportNotifySendRule(strings.ToLower(string(channel.SendRule)))
	if channel.ID == 0 {
		return errors.New("ID required")
	}
	var existing projectmgr.ProjectReportNotifyChannel
	if err := global.GVA_DB.Where("id = ? AND project_id = ?", channel.ID, projectId).First(&existing).Error; err != nil {
		return err
	}

	updates := map[string]any{
		"provider":         channel.Provider,
		"name":             channel.Name,
		"enabled":          channel.Enabled,
		"send_rule":        channel.SendRule,
		"webhook_url":      channel.WebhookURL,
		"webhook_secret":   channel.WebhookSecret,
		"template_success": channel.TemplateSuccess,
		"template_fail":    channel.TemplateFail,
		"web_base_url":     channel.WebBaseURL,
		"extra":            channel.Extra,
	}
	return global.GVA_DB.Model(&projectmgr.ProjectReportNotifyChannel{}).Where("id = ? AND project_id = ?", channel.ID, projectId).Updates(updates).Error
}

func (s *ReportNotifyService) GetReportNotifyChannel(ctx context.Context, id uint, projectId int64) (projectmgr.ProjectReportNotifyChannel, error) {
	var channel projectmgr.ProjectReportNotifyChannel
	err := global.GVA_DB.Where("id = ? AND project_id = ?", id, projectId).First(&channel).Error
	return channel, err
}

func (s *ReportNotifyService) GetReportNotifyChannelList(ctx context.Context, info projectmgrReq.ReportNotifyChannelSearch) (list []projectmgr.ProjectReportNotifyChannel, total int64, err error) {
	db := global.GVA_DB.Model(&projectmgr.ProjectReportNotifyChannel{})
	if info.ProjectId != 0 {
		db = db.Where("project_id = ?", info.ProjectId)
	}
	if info.Provider != "" {
		db = db.Where("provider = ?", strings.ToLower(info.Provider))
	}
	if info.Enabled != nil {
		db = db.Where("enabled = ?", *info.Enabled)
	}
	if info.Keyword != "" {
		kw := "%" + info.Keyword + "%"
		db = db.Where("name LIKE ?", kw)
	}
	db = db.Order("id desc")

	if err = db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if info.PageSize != 0 {
		db = db.Scopes(info.PageInfo.Paginate())
	}
	var items []projectmgr.ProjectReportNotifyChannel
	if err = db.Find(&items).Error; err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

func (s *ReportNotifyService) GetAutoReportNotifyStatus(ctx context.Context, reportId uint) (AutoReportNotifyStatusResponse, error) {
	var report automation.AutoReport
	if err := global.GVA_DB.Select("id", "project_id", "status", "notify_enabled", "notify_rule", "notify_channel_ids").Where("id = ?", reportId).First(&report).Error; err != nil {
		return AutoReportNotifyStatusResponse{}, err
	}
	status := int64(0)
	if report.Status != nil {
		status = *report.Status
	}

	notifyEnabled := true
	if report.NotifyEnabled != nil {
		notifyEnabled = *report.NotifyEnabled
	}
	selectedChannelIDs := parseNotifyChannelIDs(report.NotifyChannelIDs)

	result := ""
	switch status {
	case int64(automation.ReportStatusSuccess):
		result = string(projectmgr.ReportNotifyReportResultSuccess)
	case int64(automation.ReportStatusFailed):
		result = string(projectmgr.ReportNotifyReportResultFail)
	}

	var channels []projectmgr.ProjectReportNotifyChannel
	channelDB := global.GVA_DB.Where("project_id = ? AND enabled = ?", report.ProjectId, true)
	if len(selectedChannelIDs) > 0 {
		channelDB = channelDB.Where("id IN ?", selectedChannelIDs)
	} else if !notifyEnabled {
		channels = []projectmgr.ProjectReportNotifyChannel{}
		return AutoReportNotifyStatusResponse{
			ReportId:     reportId,
			ProjectId:    report.ProjectId,
			ReportStatus: status,
			Result:       result,
			Items:        []AutoReportNotifyChannelStatus{},
		}, nil
	}
	if err := channelDB.Order("id desc").Find(&channels).Error; err != nil {
		return AutoReportNotifyStatusResponse{}, err
	}

	logByChannel := map[uint]projectmgr.ProjectReportNotifyLog{}
	if result != "" {
		var logs []projectmgr.ProjectReportNotifyLog
		if err := global.GVA_DB.Where("report_id = ? AND report_result = ?", reportId, result).Find(&logs).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return AutoReportNotifyStatusResponse{}, err
		}
		for _, l := range logs {
			logByChannel[l.ChannelId] = l
		}
	}

	items := make([]AutoReportNotifyChannelStatus, 0, len(channels))
	for _, ch := range channels {
		rule := ch.SendRule
		if report.NotifyRule != "" {
			rule = projectmgr.ReportNotifySendRule(report.NotifyRule)
		}
		state, okPtr, errMsg := s.channelState(status, notifyEnabled, rule, result, logByChannel[ch.ID])
		items = append(items, AutoReportNotifyChannelStatus{
			ChannelId: ch.ID,
			Provider:  string(ch.Provider),
			Name:      ch.Name,
			SendRule:  string(rule),
			State:     state,
			Ok:        okPtr,
			Error:     errMsg,
			Display:   fmt.Sprintf("%s-%s（%s）", providerLabel(ch.Provider), ch.Name, state),
		})
	}

	return AutoReportNotifyStatusResponse{
		ReportId:     reportId,
		ProjectId:    report.ProjectId,
		ReportStatus: status,
		Result:       result,
		Items:        items,
	}, nil
}

func (s *ReportNotifyService) NotifyAutoReport(ctx context.Context, reportId uint) error {
	global.GVA_LOG.Debug("NotifyAutoReport called", zap.Uint("reportId", reportId))
	var report automation.AutoReport
	if err := global.GVA_DB.
		Preload("Stat").
		Preload("Stat.Testcases").
		Preload("Time").
		Preload("Details").
		Where("id = ?", reportId).
		First(&report).Error; err != nil {
		global.GVA_LOG.Error("NotifyAutoReport fetch report failed", zap.Uint("reportId", reportId), zap.Error(err))
		return err
	}
	status := int64(0)
	if report.Status != nil {
		status = *report.Status
	}
	var result projectmgr.ReportNotifyReportResult
	switch status {
	case int64(automation.ReportStatusSuccess):
		result = projectmgr.ReportNotifyReportResultSuccess
	case int64(automation.ReportStatusFailed):
		result = projectmgr.ReportNotifyReportResultFail
	default:
		return nil
	}

	notifyEnabled := true
	if report.NotifyEnabled != nil {
		notifyEnabled = *report.NotifyEnabled
	}
	global.GVA_LOG.Debug("NotifyAutoReport check enabled", zap.Uint("reportId", reportId), zap.Bool("enabled", notifyEnabled))
	if !notifyEnabled {
		return nil
	}

	var channels []projectmgr.ProjectReportNotifyChannel
	channelDB := global.GVA_DB.Where("project_id = ? AND enabled = ?", report.ProjectId, true)
	selectedChannelIDs := parseNotifyChannelIDs(report.NotifyChannelIDs)
	if len(selectedChannelIDs) > 0 {
		channelDB = channelDB.Where("id IN ?", selectedChannelIDs)
	}
	if err := channelDB.Find(&channels).Error; err != nil {
		global.GVA_LOG.Error("NotifyAutoReport fetch channels failed", zap.Uint("reportId", reportId), zap.Error(err))
		return err
	}
	global.GVA_LOG.Debug("NotifyAutoReport channels found", zap.Uint("reportId", reportId), zap.Int("count", len(channels)))
	for _, ch := range channels {
		global.GVA_LOG.Debug("NotifyAutoReport processing channel", zap.Uint("reportId", reportId), zap.Uint("channelId", ch.ID), zap.String("name", ch.Name))
		rule := ch.SendRule
		if report.NotifyRule != "" {
			rule = projectmgr.ReportNotifySendRule(report.NotifyRule)
		}
		if !matchRule(rule, result) {
			global.GVA_LOG.Debug("NotifyAutoReport rule mismatch", zap.Uint("reportId", reportId), zap.Uint("channelId", ch.ID), zap.String("rule", string(rule)), zap.String("result", string(result)))
			continue
		}
		if err := s.sendOnce(ctx, report, ch, result); err != nil {
			global.GVA_LOG.Error("NotifyAutoReport sendOnce failed", zap.Uint("reportId", reportId), zap.Uint("channelId", ch.ID), zap.Error(err))
			continue
		}
		global.GVA_LOG.Debug("NotifyAutoReport sendOnce success", zap.Uint("reportId", reportId), zap.Uint("channelId", ch.ID))
	}
	return nil
}

func (s *ReportNotifyService) sendOnce(ctx context.Context, report automation.AutoReport, ch projectmgr.ProjectReportNotifyChannel, result projectmgr.ReportNotifyReportResult) error {
	global.GVA_LOG.Debug("sendOnce called", zap.Uint("reportId", report.ID), zap.Uint("channelId", ch.ID))
	now := time.Now()
	logRow := projectmgr.ProjectReportNotifyLog{
		ProjectId:    report.ProjectId,
		ReportId:     report.ID,
		ChannelId:    ch.ID,
		Provider:     ch.Provider,
		ReportResult: result,
		Ok:           false,
		SentAt:       now,
	}

	res := global.GVA_DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&logRow)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return nil
	}

	payloadBytes, sendErr := s.sendToChannel(ctx, report, ch, result)
	updates := map[string]any{
		"sent_at":         time.Now(),
		"request_payload": datatypes.JSON(payloadBytes),
	}
	if sendErr != nil {
		updates["ok"] = false
		updates["error"] = sendErr.Error()
	} else {
		updates["ok"] = true
		updates["error"] = ""
	}

	return global.GVA_DB.Model(&projectmgr.ProjectReportNotifyLog{}).
		Where("report_id = ? AND channel_id = ? AND report_result = ?", report.ID, ch.ID, result).
		Updates(updates).Error
}

func (s *ReportNotifyService) sendToChannel(ctx context.Context, report automation.AutoReport, ch projectmgr.ProjectReportNotifyChannel, result projectmgr.ReportNotifyReportResult) ([]byte, error) {
	global.GVA_LOG.Debug("sendToChannel called", zap.Uint("reportId", report.ID), zap.Uint("channelId", ch.ID), zap.String("provider", string(ch.Provider)))
	vars := buildTemplateVars(report, ch.WebBaseURL)
	if ch.Provider != projectmgr.ReportNotifyProviderFeishu {
		if list, ok := vars["content"].([]Content); ok {
			vars["content"] = generateTableContent(list)
		}
	}

	ctx, cancel := context.WithTimeout(ctx, 6*time.Second)
	defer cancel()

	switch ch.Provider {
	case projectmgr.ReportNotifyProviderFeishu:
		templateID := feishuSuccessTemplateID
		if result == projectmgr.ReportNotifyReportResultFail {
			templateID = feishuFailTemplateID
		}
		body := map[string]any{
			"msg_type": "interactive",
			"card": map[string]any{
				"type": "template",
				"data": map[string]any{
					"template_id":       templateID,
					"template_variable": vars,
				},
			},
		}
		if ch.WebhookSecret != "" {
			timestamp := strconv.FormatInt(time.Now().Unix(), 10)
			body["timestamp"] = timestamp
			body["sign"] = feishuSign(timestamp, ch.WebhookSecret)
		}
		b, _ := json.Marshal(body)
		global.GVA_LOG.Debug("NotifyAutoReport request", zap.Uint("reportId", report.ID), zap.Uint("channelId", ch.ID), zap.String("body", string(b)))
		resp, err := notify.PostJSON(ctx, ch.WebhookURL, body)
		if len(resp) > 0 {
			global.GVA_LOG.Debug("NotifyAutoReport response", zap.Uint("reportId", report.ID), zap.Uint("channelId", ch.ID), zap.String("body", string(resp)))
		}
		return b, err
	case projectmgr.ReportNotifyProviderDingTalk:
		text := renderTemplate(selectTemplate(ch, result), vars)
		body := map[string]any{
			"msgtype": "markdown",
			"markdown": map[string]any{
				"title": vars["title"],
				"text":  text,
			},
		}
		b, _ := json.Marshal(body)
		targetURL := dingTalkSignedURL(ch.WebhookURL, ch.WebhookSecret)
		global.GVA_LOG.Debug("NotifyAutoReport request", zap.Uint("reportId", report.ID), zap.Uint("channelId", ch.ID), zap.String("body", string(b)))
		resp, err := notify.PostJSON(ctx, targetURL, body)
		if len(resp) > 0 {
			global.GVA_LOG.Debug("NotifyAutoReport response", zap.Uint("reportId", report.ID), zap.Uint("channelId", ch.ID), zap.String("body", string(resp)))
		}
		return b, err
	case projectmgr.ReportNotifyProviderWeCom:
		text := renderTemplate(selectTemplate(ch, result), vars)
		body := map[string]any{
			"msgtype": "markdown",
			"markdown": map[string]any{
				"content": text,
			},
		}
		b, _ := json.Marshal(body)
		global.GVA_LOG.Debug("NotifyAutoReport request", zap.Uint("reportId", report.ID), zap.Uint("channelId", ch.ID), zap.String("body", string(b)))
		resp, err := notify.PostJSON(ctx, ch.WebhookURL, body)
		if len(resp) > 0 {
			global.GVA_LOG.Debug("NotifyAutoReport response", zap.Uint("reportId", report.ID), zap.Uint("channelId", ch.ID), zap.String("body", string(resp)))
		}
		return b, err
	default:
		body := map[string]any{
			"msgtype": "text",
			"text": map[string]any{
				"content": renderTemplate(selectTemplate(ch, result), vars),
			},
		}
		b, _ := json.Marshal(body)
		global.GVA_LOG.Debug("NotifyAutoReport request", zap.Uint("reportId", report.ID), zap.Uint("channelId", ch.ID), zap.String("body", string(b)))
		resp, err := notify.PostJSON(ctx, ch.WebhookURL, body)
		if len(resp) > 0 {
			global.GVA_LOG.Debug("NotifyAutoReport response", zap.Uint("reportId", report.ID), zap.Uint("channelId", ch.ID), zap.String("body", string(resp)))
		}
		return b, err
	}
}

const (
	feishuSuccessTemplateID = "ctp_AAmjuIxqEQjK"
	feishuFailTemplateID    = "ctp_AAmIkXEaoPea"
)

func providerLabel(p projectmgr.ReportNotifyProvider) string {
	switch p {
	case projectmgr.ReportNotifyProviderFeishu:
		return "飞书"
	case projectmgr.ReportNotifyProviderDingTalk:
		return "钉钉"
	case projectmgr.ReportNotifyProviderWeCom:
		return "企业微信"
	default:
		return string(p)
	}
}

func matchRule(rule projectmgr.ReportNotifySendRule, result projectmgr.ReportNotifyReportResult) bool {
	switch rule {
	case projectmgr.ReportNotifySendRuleAlways:
		return true
	case projectmgr.ReportNotifySendRuleSuccess:
		return result == projectmgr.ReportNotifyReportResultSuccess
	case projectmgr.ReportNotifySendRuleFail:
		return result == projectmgr.ReportNotifyReportResultFail
	default:
		return false
	}
}

func (s *ReportNotifyService) channelState(reportStatus int64, notifyEnabled bool, rule projectmgr.ReportNotifySendRule, result string, log projectmgr.ProjectReportNotifyLog) (string, *bool, string) {
	if !notifyEnabled {
		return "未发送", nil, ""
	}
	if reportStatus != int64(automation.ReportStatusSuccess) && reportStatus != int64(automation.ReportStatusFailed) {
		return "待发送", nil, ""
	}
	if result == "" {
		return "待发送", nil, ""
	}
	if !matchRule(rule, projectmgr.ReportNotifyReportResult(result)) {
		return "未发送", nil, ""
	}
	if log.ID == 0 {
		return "待发送", nil, ""
	}
	ok := log.Ok
	if ok {
		return "成功发送", &ok, ""
	}
	return "发送失败", &ok, log.Error
}

func parseNotifyChannelIDs(raw datatypes.JSON) []uint {
	if len(raw) == 0 {
		return nil
	}
	var ids []uint
	if err := json.Unmarshal(raw, &ids); err != nil {
		return nil
	}
	return ids
}

func selectTemplate(ch projectmgr.ProjectReportNotifyChannel, result projectmgr.ReportNotifyReportResult) string {
	if result == projectmgr.ReportNotifyReportResultFail {
		if ch.TemplateFail != "" {
			return ch.TemplateFail
		}
	}
	if ch.TemplateSuccess != "" {
		return ch.TemplateSuccess
	}
	return defaultTextTemplate()
}

func defaultTextTemplate() string {
	return "**{{title}}**\n\n- 状态：{{status}}\n- 环境：{{env}}\n- 成功：{{success}}\n- 失败：{{fail}}\n- 耗时：{{time}}\n- 详情：{{detail}}\n"
}

type Content struct {
	Name    string `json:"name"`
	Success int    `json:"success"`
	Fail    int    `json:"fail"`
	Time    int    `json:"time"`
	Total   int    `json:"total"`
	Error   int    `json:"error"`
	Skip    int    `json:"skip"`
}

func buildTemplateVars(report automation.AutoReport, webBaseURL string) map[string]any {
	title := ""
	if report.Name != nil {
		title = *report.Name
	}
	status := ""
	if report.Status != nil {
		if *report.Status == int64(automation.ReportStatusSuccess) {
			status = "成功"
		} else if *report.Status == int64(automation.ReportStatusFailed) {
			status = "失败"
		}
	}
	successCount := 0
	failCount := 0
	if report.Stat != nil && report.Stat.Testcases != nil {
		successCount = report.Stat.Testcases.Success
		failCount = report.Stat.Testcases.Fail
	}
	duration := 0.0
	if report.Time != nil {
		duration = report.Time.Duration
	}
	detail := ""
	if webBaseURL != "" {
		base := strings.TrimRight(webBaseURL, "/")
		detail = fmt.Sprintf("%s/#/layout/auto-report-detail/%d", base, report.ID)
	}

	var contents []Content
	for _, v := range report.Details {
		statMap := v.Stat
		timeMap := v.Time

		// Debug log to inspect map content
		global.GVA_LOG.Debug("inspecting detail stat",
			zap.String("name", v.Name),
			zap.Any("stat", statMap),
			zap.Any("time", timeMap))

		getFloat := func(m map[string]any, k string) float64 {
			if m == nil {
				return 0
			}
			if val, ok := m[k]; ok {
				switch v := val.(type) {
				case float64:
					return v
				case int:
					return float64(v)
				case int64:
					return float64(v)
				case float32:
					return float64(v)
				case string:
					f, _ := strconv.ParseFloat(v, 64)
					return f
				case json.Number:
					f, _ := v.Float64()
					return f
				}
			}
			return 0
		}

		durationValue := getFloat(timeMap, "duration")

		content := Content{
			Name:    v.Name,
			Total:   int(getFloat(statMap, "total")),
			Fail:    int(getFloat(statMap, "failures")),
			Success: int(getFloat(statMap, "successes")),
			Error:   int(getFloat(statMap, "error")),
			Skip:    int(getFloat(statMap, "skip")),
			Time:    int(durationValue),
		}
		contents = append(contents, content)
	}

	return map[string]any{
		"title":   title,
		"name":    title,
		"status":  status,
		"env":     report.EnvName,
		"success": fmt.Sprintf("%d", successCount),
		"fail":    fmt.Sprintf("%d", failCount),
		"time":    fmt.Sprintf("%.2f秒", duration),
		"detail":  detail,
		"content": contents,
	}
}

func renderTemplate(tpl string, vars map[string]any) string {
	out := tpl
	for k, v := range vars {
		s := fmt.Sprintf("%v", v)
		out = strings.ReplaceAll(out, "{{"+k+"}}", s)
		out = strings.ReplaceAll(out, "${"+k+"}", s)
	}
	return out
}

func generateTableContent(data []Content) (tableContent string) {
	for _, row := range data {
		tableContent += "<tr>"
		tableContent += fmt.Sprintf("<td>%s</td>", row.Name)
		tableContent += fmt.Sprintf("<td>%d</td>", row.Success)
		tableContent += fmt.Sprintf("<td>%d</td>", row.Fail)
		tableContent += fmt.Sprintf("<td>%d秒</td>", row.Time)
		tableContent += "</tr>"
	}
	tableContent = fmt.Sprintf("<table><tr><th style=\"min-width: 80px; max-width: 240px;\">用例名称</th>"+
		"<th style=\"min-width: 80px; max-width: 240px;\">成功数</th>"+
		"<th style=\"min-width: 80px; max-width: 240px;\">失败数</th>"+
		"<th style=\"min-width: 80px; max-width: 240px;\">耗时</th></tr>%s</table>", tableContent)

	return tableContent
}

func feishuSign(timestamp, secret string) string {
	toSign := timestamp + "\n" + secret
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(toSign))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func dingTalkSignedURL(webhookURL, secret string) string {
	if secret == "" {
		return webhookURL
	}
	u, err := url.Parse(webhookURL)
	if err != nil {
		return webhookURL
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return webhookURL
	}
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	sign := dingTalkSign(timestamp, secret)
	q := u.Query()
	q.Set("timestamp", timestamp)
	q.Set("sign", sign)
	u.RawQuery = q.Encode()
	return u.String()
}

func dingTalkSign(timestamp, secret string) string {
	toSign := timestamp + "\n" + secret
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(toSign))
	return url.QueryEscape(base64.StdEncoding.EncodeToString(mac.Sum(nil)))
}
