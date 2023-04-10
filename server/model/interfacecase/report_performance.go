package interfacecase

import (
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/system"
	"gorm.io/datatypes"
)

type PerformanceReport struct {
	global.GVA_MODEL
	Name                    string                    `json:"name"`
	State                   State                     `json:"state"`
	PerformanceReportDetail []PerformanceReportDetail `json:"performance_report_detail"`
	PerformanceID           uint                      `json:"performance_id"`
	ProjectID               uint                      `json:"-"`
	Project                 system.Project            `json:"-"`
	TotalStats              datatypes.JSONMap         `json:"stats_total" gorm:"column:stats_total;comment:;type:text"`
	TotalAvgResponseTime    float64                   `json:"total_avg_response_time"`
	TotalMinResponseTime    float64                   `json:"total_min_response_time"`
	TotalMaxResponseTime    float64                   `json:"total_max_response_time"`
	TotalRPS                float64                   `json:"total_rps"`
	TotalFailRatio          float64                   `json:"total_fail_ratio"`
	TotalFailPerSec         float64                   `json:"total_fail_per_sec"`
	UserCount               int64                     `json:"user_count"`
}

type State int

const (
	StateInit     State = iota + 1 // initializing
	StateSpawning                  // spawning
	StateRunning                   // running
	StateStopping                  // stopping
	StateStopped                   // stopped
	StateQuitting                  // quitting
	StateMissing                   // missing
	StateError                     // 运行报错
)

type PerformanceReportDetail struct {
	global.GVA_MODEL
	UserCount            int64             `json:"user_count"`
	State                State             `json:"state"`
	TotalStats           datatypes.JSONMap `json:"stats_total" gorm:"column:stats_total;comment:;type:text"`
	TransactionsPassed   int64             `json:"transactions_passed"`
	TransactionsFailed   int64             `json:"transactions_failed"`
	TotalAvgResponseTime float64           `json:"total_avg_response_time"`
	TotalMinResponseTime float64           `json:"total_min_response_time"`
	TotalMaxResponseTime float64           `json:"total_max_response_time"`
	TotalRPS             float64           `json:"total_rps"`
	TotalFailRatio       float64           `json:"total_fail_ratio"`
	TotalFailPerSec      float64           `json:"total_fail_per_sec"`
	Duration             float64           `json:"duration"`
	Errors               datatypes.JSON    `json:"errors" gorm:"column:errors;comment:;type:text"`
	PerformanceReportID  uint

	PerformanceReportTotalStats []PerformanceReportTotalStats `json:"stats"`
}

type PerformanceReportTotalStats struct {
	global.GVA_MODEL
	Name                      string  `json:"name"`
	Method                    string  `json:"method"`
	StartTime                 int64   `json:"start_time"`
	NumFailures               int     `json:"num_failures"`
	NumRequests               int     `json:"num_requests"`
	MaxResponseTime           int     `json:"max_response_time"`
	MinResponseTime           int     `json:"min_response_time"`
	NumNoneRequests           int     `json:"num_none_requests"`
	TotalResponseTime         int     `json:"total_response_time"`
	TotalContentLength        int     `json:"total_content_length"`
	LastRequestTimestamp      int64   `json:"last_request_timestamp"`
	CurrentRps                float64 `json:"current_rps"`
	CurrentFailPerSec         float64 `json:"current_fail_per_sec"`
	PerformanceReportDetailID uint

	ResponseTimer datatypes.JSONMap `json:"response_timer"`
}
