// 自动生成模板AutoReport
package automation

import (
	"context"
	"fmt"
	"time"

	"github.com/test-instructor/yangfan/server/v2/global"
	"gorm.io/datatypes"
)

// 报告状态常量
const (
	ReportStatusPending = 0 // 待运行
	ReportStatusRunning = 1 // 运行中
	ReportStatusFailed  = 2 // 失败
	ReportStatusSuccess = 3 // 成功
)

// 自动报告 结构体  AutoReport
type AutoReport struct {
	global.GVA_MODEL
	Name       *string             `json:"name" form:"name" gorm:"column:name;"`          //名称
	Success    *bool               `json:"success" form:"success" gorm:"column:success;"` //成功
	StatID     *int64              `json:"stat_id" form:"stat_id" gorm:"column:stat_id;"` //统计ID
	Stat       *AutoReportStat     `json:"stat" gorm:"foreignKey:StatID;references:ID"`
	TimeID     *int64              `json:"time_id" form:"time_id" gorm:"column:time_id;"` //时间ID
	Time       *AutoReportTime     `json:"time" gorm:"foreignKey:TimeID;references:ID"`
	ProgressID *uint               `json:"progress_id" form:"progress_id" gorm:"column:progress_id;"` //进度ID
	Progress   *AutoReportProgress `json:"progress" gorm:"foreignKey:ProgressID;references:ID"`
	Platform   datatypes.JSONMap   `json:"platform" form:"platform" gorm:"column:platform;" swaggertype:"object"` //平台
	Status     *int64              `json:"status" form:"status" gorm:"column:status;"`                            //状态
	SetupCase  *bool               `json:"setup_case" form:"setup_case" gorm: "column:setup_case;"`               //设置案例
	Describe   *string             `json:"describe" form:"describe" gorm:"column:describe;"`                      //描述
	Hostname   *string             `json:"hostname" form:"hostname" gorm:"comment:主机名;;​​column:hostname;"`       //主机名
	Details    []AutoReportDetail  `json:"details" gorm:"foreignKey:AutoReportID"`

	ProjectId  int64  `json:"projectId" form:"projectId" gorm:"column:project_id;"` //项目信息
	CaseType   string `json:"case_type" form:"case_type" gorm:"column:case_type;"`  // 用例类型
	RunMode    string `json:"run_mode" form:"run_mode" gorm:"column:run_mode;"`     // 运行模式
	ConfigID   int    `json:"config_id" form:"config_id" gorm:"column:config_id;"`  // 配置ID
	EnvID      int    `json:"env_id" form:"env_id" gorm:"column:env_id;"`           // 环境ID
	EnvName    string `json:"env_name" form:"env_name" gorm:"column:env_name;"`
	ConfigName string `json:"config_name" form:"config_name" gorm:"column:config_name;"`
	CaseID     uint   `json:"case_id" form:"case_id" gorm:"column:case_id;"` // 用例ID
}

// TableName 自动报告 AutoReport自定义表名 auto_reports
func (AutoReport) TableName() string {
	return "lc_auto_reports"
}

// Redis key prefix for progress data
const redisProgressKeyPrefix = "test_report_progress"

// buildProgressKey formats a redis key like: test_report_progress:{report_id}:{field}
func buildProgressKey(reportID uint, field string) string {
	return fmt.Sprintf("%s:%d:%s", redisProgressKeyPrefix, reportID, field)
}

// LoadProgressFromRedis 如果 Progress 为空，尝试从 Redis 加载进度数据
// 这个方法应该在查询报告后调用，以确保进度数据完整
func (ar *AutoReport) LoadProgressFromRedis() {
	// 如果已有 Progress 数据，无需从 Redis 加载
	if ar.Progress != nil {
		return
	}

	// 如果 Redis 不可用，无法加载
	if global.GVA_REDIS == nil {
		return
	}

	// 从 Redis 读取进度数据
	ctx := context.Background()
	pipe := global.GVA_REDIS.TxPipeline()

	totalCasesCmd := pipe.Get(ctx, buildProgressKey(ar.ID, "total_cases"))
	totalStepsCmd := pipe.Get(ctx, buildProgressKey(ar.ID, "total_steps"))
	totalApisCmd := pipe.Get(ctx, buildProgressKey(ar.ID, "total_apis"))
	executedCasesCmd := pipe.Get(ctx, buildProgressKey(ar.ID, "executed_cases"))
	executedStepsCmd := pipe.Get(ctx, buildProgressKey(ar.ID, "executed_steps"))
	executedApisCmd := pipe.Get(ctx, buildProgressKey(ar.ID, "executed_apis"))

	if _, err := pipe.Exec(ctx); err != nil {
		// Redis 中没有数据是正常的（可能已过期或从未存入）
		return
	}

	// 解析数据，如果任何一个字段有数据，则创建 Progress 对象
	totalCases, err1 := totalCasesCmd.Int()
	totalSteps, err2 := totalStepsCmd.Int()
	totalApis, err3 := totalApisCmd.Int()
	executedCases, err4 := executedCasesCmd.Int()
	executedSteps, err5 := executedStepsCmd.Int()
	executedApis, err6 := executedApisCmd.Int()

	// 如果所有字段都读取失败，说明 Redis 中没有数据
	if err1 != nil && err2 != nil && err3 != nil && err4 != nil && err5 != nil && err6 != nil {
		return
	}

	// 创建 Progress 对象（不存储到数据库，仅用于当前查询返回）
	ar.Progress = &AutoReportProgress{
		TotalCases:    totalCases,
		TotalSteps:    totalSteps,
		TotalApis:     totalApis,
		ExecutedCases: executedCases,
		ExecutedSteps: executedSteps,
		ExecutedApis:  executedApis,
	}
}

// 统计信息
type AutoReportStat struct {
	global.GVA_MODEL
	TestcasesID   *int64                     `json:"testcases_id" gorm:"column:testcases_id"`
	Testcases     *AutoReportStatTestcases   `json:"testcases" gorm:"foreignKey:TestcasesID;references:ID"`
	TeststepsID   *int64                     `json:"teststeps_id" gorm:"column:teststeps_id"`
	Teststeps     *AutoReportStatTeststeps   `json:"teststeps" gorm:"foreignKey:TeststepsID;references:ID"`
	TeststepapiID *int64                     `json:"teststepapi_id" gorm:"column:teststepapi_id"`
	Teststepapi   *AutoReportStatTeststepapi `json:"teststepapi" gorm:"foreignKey:TeststepapiID;references:ID"`
}

func (AutoReportStat) TableName() string {
	return "lc_auto_report_stats"
}

type AutoReportStatTestcases struct {
	global.GVA_MODEL
	Total   int `json:"total" gorm:"column:total"`
	Success int `json:"success" gorm:"column:success"`
	Fail    int `json:"fail" gorm:"column:fail"`
}

func (AutoReportStatTestcases) TableName() string {
	return "lc_auto_report_stat_testcases"
}

type AutoReportStatTeststeps struct {
	global.GVA_MODEL
	Total     int               `json:"total" gorm:"column:total"`
	Successes int               `json:"successes" gorm:"column:successes"`
	Failures  int               `json:"failures" gorm:"column:failures"`
	Actions   datatypes.JSONMap `json:"actions" gorm:"column:actions"`
}

func (AutoReportStatTeststeps) TableName() string {
	return "lc_auto_report_stat_teststeps"
}

type AutoReportStatTeststepapi struct {
	global.GVA_MODEL
	Total   int `json:"total" gorm:"column:total"`
	Success int `json:"success" gorm:"column:success"`
	Fail    int `json:"fail" gorm:"column:fail"`
}

func (AutoReportStatTeststepapi) TableName() string {
	return "lc_auto_report_stat_teststepapi"
}

// 时间信息
type AutoReportTime struct {
	global.GVA_MODEL
	StartAt  time.Time `json:"start_at" gorm:"column:start_at"`
	Duration float64   `json:"duration" gorm:"column:duration"`
}

func (AutoReportTime) TableName() string {
	return "lc_auto_report_times"
}

// 详情
type AutoReportDetail struct {
	global.GVA_MODEL
	AutoReportID uint               `json:"auto_report_id" gorm:"column:auto_report_id"`
	Name         string             `json:"name" gorm:"column:name"`
	Success      bool               `json:"success" gorm:"column:success"`
	Stat         datatypes.JSONMap  `json:"stat" gorm:"column:stat"`
	Time         datatypes.JSONMap  `json:"time" gorm:"column:time"`
	InOut        datatypes.JSONMap  `json:"in_out" gorm:"column:in_out"`
	Records      []AutoReportRecord `json:"records" gorm:"foreignKey:AutoReportDetailID"`
}

func (AutoReportDetail) TableName() string {
	return "lc_auto_report_details"
}

// 记录
type AutoReportRecord struct {
	global.GVA_MODEL
	AutoReportDetailID uint              `json:"auto_report_detail_id" gorm:"column:auto_report_detail_id"`
	Name               string            `json:"name" gorm:"column:name"`
	StartTime          int64             `json:"start_time" gorm:"column:start_time"`
	StepType           string            `json:"step_type" gorm:"column:step_type"`
	Success            bool              `json:"success" gorm:"column:success"`
	ElapsedMs          int64             `json:"elapsed_ms" gorm:"column:elapsed_ms"`
	HttpStat           datatypes.JSONMap `json:"httpstat" gorm:"column:httpstat"`
	Data               datatypes.JSONMap `json:"data" gorm:"column:data"`
	ContentSize        int64             `json:"content_size" gorm:"column:content_size"`
	ExportVars         datatypes.JSONMap `json:"export_vars" gorm:"column:export_vars"`
}

func (AutoReportRecord) TableName() string {
	return "lc_auto_report_records"
}
