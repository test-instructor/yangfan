package request

type RunnerRequest struct {
	CaseType  string `json:"case_type"` // 接口/步骤/用例/任务/标签
	CaseID    uint   `json:"case_id"`
	ReportID  uint   `json:"report_id"` // 测试报告ID
	RunMode   string `json:"run_mode"`  // 调试模式/保存并调试/立即运行/后台运行/定时执行/CI
	NodeName  string `json:"node_name"` // 可选，指定节点；不填则随机
	Timeout   int    `json:"timeout"`   // 可选，任务超时时间（秒）
	EnvID     int    `json:"env_id"`    // 环境ID
	ConfigID  int    `json:"config_id"` // 配置ID
	ProjectId uint   `json:"projectId"`

	NotifyEnabled    *bool  `json:"notify_enabled"`
	NotifyRule       string `json:"notify_rule"`
	NotifyChannelIDs []uint `json:"notify_channel_ids"`
	Failfast         *bool  `json:"failfast"`
}

type RunnerResponse struct {
	TaskID   string `json:"task_id"`
	ReportID uint   `json:"report_id"` // 测试报告ID
	NodeName string `json:"node_name"`
	SendTime int64  `json:"send_time"`
}
