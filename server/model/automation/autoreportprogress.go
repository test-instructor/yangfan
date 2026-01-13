// Package automation 测试报告进度表
package automation

import (
	"github.com/test-instructor/yangfan/server/v2/global"
)

// AutoReportProgress 测试报告进度 结构体
type AutoReportProgress struct {
	global.GVA_MODEL
	TotalCases    int `json:"total_cases" form:"total_cases" gorm:"column:total_cases;comment:用例总数"`
	TotalSteps    int `json:"total_steps" form:"total_steps" gorm:"column:total_steps;comment:步骤总数(考虑循环)"`
	TotalApis     int `json:"total_apis" form:"total_apis" gorm:"column:total_apis;comment:接口总数(考虑循环)"`
	ExecutedCases int `json:"executed_cases" form:"executed_cases" gorm:"column:executed_cases;default:0;comment:已执行用例数"`
	ExecutedSteps int `json:"executed_steps" form:"executed_steps" gorm:"column:executed_steps;default:0;comment:已执行步骤数"`
	ExecutedApis  int `json:"executed_apis" form:"executed_apis" gorm:"column:executed_apis;default:0;comment:已执行接口数"`
}

// TableName 测试报告进度 自定义表名
func (AutoReportProgress) TableName() string {
	return "yf_auto_report_progress"
}
