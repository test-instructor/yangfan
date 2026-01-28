package automation

import (
	"gorm.io/datatypes"
)

// MobileStep 包装了 UI 动作列表，存储在 AutoStep 的 json 字段中
type MobileStep struct {
	Actions []MobileAction `json:"actions,omitempty"`
}

// MobileAction 定义了单个 UI 动作
// 对应 httprunner/uixt/option/action.go 中的 MobileAction
type MobileAction struct {
	Method  string         `json:"method,omitempty"`
	Params  interface{}    `json:"params,omitempty"`
	Options *ActionOptions `json:"options,omitempty"`
}

// ActionOptions 定义了动作的可选参数
// 对应 httprunner/uixt/option/action.go 中的 ActionOptions
type ActionOptions struct {
	// Device targeting
	Platform string `json:"platform,omitempty"`
	Serial   string `json:"serial,omitempty"`

	// Common action parameters
	X     float64 `json:"x,omitempty"`
	Y     float64 `json:"y,omitempty"`
	FromX float64 `json:"from_x,omitempty"`
	FromY float64 `json:"from_y,omitempty"`
	ToX   float64 `json:"to_x,omitempty"`
	ToY   float64 `json:"to_y,omitempty"`
	Text  string  `json:"text,omitempty"`

	// App/Package related
	PackageName string `json:"packageName,omitempty"`
	AppName     string `json:"appName,omitempty"`
	AppUrl      string `json:"appUrl,omitempty"`

	// Wait/Timeout
	MaxRetryTimes       int     `json:"maxRetryTimes,omitempty"`
	Interval            float64 `json:"interval,omitempty"`
	Timeout             float64 `json:"timeout,omitempty"`
	IgnoreNotFoundError bool    `json:"ignoreNotFoundError,omitempty"`

	// Selector
	Selector string      `json:"selector,omitempty"`
	Index    int         `json:"index,omitempty"`
	Regex    bool        `json:"regex,omitempty"`
	Offset   interface{} `json:"offset,omitempty"` // map or struct

	// Other
	Identifier       string            `json:"identifier,omitempty"`
	PreMarkOperation bool              `json:"pre_mark_operation,omitempty"`
	OutputSchema     datatypes.JSONMap `json:"output_schema,omitempty"`
}
