package datacategory

// CallType 调用类型常量
const (
	CallTypeTestStep = 1 // 测试步骤
	CallTypePython   = 2 // Python
	CallTypeDelete   = 3 // 直接删除 (仅用于 CleanCallType)
)

// DataStatus 数据状态常量
const (
	DataStatusAvailable = 0 // 可用
	DataStatusUsed      = 1 // 已占用
	DataStatusCleaned   = 2 // 已清洗
)

// PythonCodeType Python代码类型
const (
	PythonCodeTypeAutomation   = 1 // 自动化代码
	PythonCodeTypeDataCategory = 2 // 数据分类代码
)

// ProcessResult 处理结果
type ProcessResult struct {
	CategoryID   uint                   `json:"categoryId"`
	CategoryName string                 `json:"categoryName"`
	EnvID        uint                   `json:"envId"`
	Action       string                 `json:"action"` // "clean" or "create"
	Success      bool                   `json:"success"`
	Message      string                 `json:"message"`
	DataCount    int                    `json:"dataCount"`
	LastData     map[string]interface{} `json:"-"` // 用于回传更新后的 LastData，不序列化到前端
}

// ExecuteContext 执行上下文
type ExecuteContext struct {
	ProjectID      int64
	CategoryID     uint
	CategoryName   string
	Type           *string // 数据类型
	EnvID          uint
	EnvName        string
	CreateCallType *int64
	CleanCallType  *int64
}
