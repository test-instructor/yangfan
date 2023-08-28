package request

import "gorm.io/datatypes"

type CIRun struct {
	TagID       uint              `json:"tag" form:"tag"`                   // 标签ID
	TaskID      uint              `json:"task" form:"task"`                 // 任务ID
	EnvID       uint              `json:"env" form:"env"`                   // 环境ID
	UUID        string            `json:"uuid" form:"uuid"`                 // 项目UUID
	Secret      string            `json:"secret" form:"secret"`             // 项目Secret
	ProjectID   uint              `json:"project" form:"project"`           // 项目ID
	MessageID   uint              `json:"message" form:"message"`           // 消息ID
	ReportID    uint              `json:"report" form:"report"`             // 报告ID
	Key         string            `json:"key" form:"key"`                   // 运行Key
	CallbackUrl string            `json:"callback_url" form:"callback_url"` // 回调地址
	Other       datatypes.JSONMap `json:"other" form:"other"`               // 其他参数
}
