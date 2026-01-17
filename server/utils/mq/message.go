package mq

// TaskMessage 定义统一的任务消息结构
type TaskMessage struct {
	TaskID           string      `json:"task_id"`
	ReportID         uint        `json:"report_id"` // 测试报告ID
	CaseType         string      `json:"case_type"` // 接口/步骤/用例/任务/标签
	CaseID           uint        `json:"case_id"`
	RunMode          string      `json:"run_mode"` // 调试模式/保存并调试/立即运行/后台运行/定时执行/CI
	NodeName         string      `json:"node_name"`
	Data             interface{} `json:"data"`
	EnvID            int         `json:"env_id"`
	ConfigID         int         `json:"config_id"`
	ProjectID        uint        `json:"project_id"`
	NotifyEnabled    *bool       `json:"notify_enabled,omitempty"`
	NotifyRule       string      `json:"notify_rule,omitempty"`
	NotifyChannelIDs []uint      `json:"notify_channel_ids,omitempty"`
	CreateTime       int64       `json:"create_time"`
	ExpireTime       int64       `json:"expire_time"`
}
