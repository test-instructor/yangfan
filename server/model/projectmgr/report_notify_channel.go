package projectmgr

import (
	"github.com/test-instructor/yangfan/server/v2/global"
	"gorm.io/datatypes"
)

type ReportNotifySendRule string

const (
	ReportNotifySendRuleSuccess ReportNotifySendRule = "success"
	ReportNotifySendRuleFail    ReportNotifySendRule = "fail"
	ReportNotifySendRuleAlways  ReportNotifySendRule = "always"
)

type ReportNotifyProvider string

const (
	ReportNotifyProviderFeishu   ReportNotifyProvider = "feishu"
	ReportNotifyProviderDingTalk ReportNotifyProvider = "dingtalk"
	ReportNotifyProviderWeCom    ReportNotifyProvider = "wecom"
)

type ProjectReportNotifyChannel struct {
	global.GVA_MODEL
	ProjectId       int64                `json:"projectId" form:"projectId" gorm:"column:project_id;index:idx_project_provider,priority:1;uniqueIndex:uniq_project_provider_name,priority:1"`
	Provider        ReportNotifyProvider `json:"provider" form:"provider" gorm:"column:provider;type:varchar(16);index:idx_project_provider,priority:2;uniqueIndex:uniq_project_provider_name,priority:2"`
	Name            string               `json:"name" form:"name" gorm:"column:name;type:varchar(64);uniqueIndex:uniq_project_provider_name,priority:3"`
	Enabled         bool                 `json:"enabled" form:"enabled" gorm:"column:enabled"`
	SendRule        ReportNotifySendRule `json:"send_rule" form:"send_rule" gorm:"column:send_rule;type:varchar(16)"`
	WebhookURL      string               `json:"webhook_url" form:"webhook_url" gorm:"column:webhook_url;type:varchar(512)"`
	WebhookSecret   string               `json:"webhook_secret" form:"webhook_secret" gorm:"column:webhook_secret;type:varchar(255)"`
	TemplateSuccess string               `json:"template_success" form:"template_success" gorm:"column:template_success;type:text"`
	TemplateFail    string               `json:"template_fail" form:"template_fail" gorm:"column:template_fail;type:text"`
	WebBaseURL      string               `json:"web_base_url" form:"web_base_url" gorm:"column:web_base_url;type:varchar(255)"`
	Extra           datatypes.JSONMap    `json:"extra" form:"extra" gorm:"column:extra" swaggertype:"object"`
}

func (ProjectReportNotifyChannel) TableName() string {
	return "yf_project_report_notify_channels"
}
