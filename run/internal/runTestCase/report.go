package runTestCase

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/test-instructor/yangfan/httprunner/hrp"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/automation"
	"go.uber.org/zap"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ReportOperation struct {
	report *automation.AutoReport
}

// UpdateFromSummary 根据 hrp.Summary 计算并更新报告的用例 / 步骤 / 接口运行概要
// 参考 yangfan/run/runTestCase/report.go 的 resetReport 思路，按 lingce 的 AutoReport 结构落库。
func (r *ReportOperation) UpdateFromSummary(s *hrp.Summary) {
	if s == nil {
		return
	}

	// 顶层成功状态
	success := s.Success
	summaryReport := &automation.AutoReport{
		Success: &success,
	}

	// 1. 用例 / 步骤统计 (Summary.Stat)
	if s.Stat != nil {
		// 用例统计
		caseStat := &automation.AutoReportStatTestcases{
			Total:   s.Stat.TestCases.Total,
			Success: s.Stat.TestCases.Success,
			Fail:    s.Stat.TestCases.Fail,
		}

		// 步骤统计 + 接口统计
		stepStat := &automation.AutoReportStatTeststeps{
			Actions: datatypes.JSONMap{},
		}

		// 接口统计
		apiStat := &automation.AutoReportStatTeststepapi{}

		// 统计步骤和接口
		stepTotal := 0
		stepSuccess := 0
		stepFail := 0
		apiTotal := 0
		apiSuccess := 0
		apiFail := 0
		for _, cs := range s.Details {
			if cs == nil {
				continue
			}
			for _, step := range cs.Records {
				if step == nil {
					continue
				}
				// 重新统计步骤
				stepTotal++
				if step.Success {
					stepSuccess++
				} else {
					stepFail++
				}

				// 统计接口 (包含 "request" 或 "api" 类型的步骤)
				stepTypeStr := string(step.StepType)
				if strings.Contains(stepTypeStr, "request") || strings.Contains(stepTypeStr, "api") {
					apiTotal++
					if step.Success {
						apiSuccess++
					} else {
						apiFail++
					}
				}
			}
		}

		stepStat.Total = stepTotal
		stepStat.Successes = stepSuccess
		stepStat.Failures = stepFail

		apiStat.Total = apiTotal
		apiStat.Success = apiSuccess
		apiStat.Fail = apiFail

		summaryReport.Stat = &automation.AutoReportStat{
			Testcases:   caseStat,
			Teststeps:   stepStat,
			Teststepapi: apiStat,
		}
	}

	// 2. 整体时间信息：从 StartAt 到当前时间重新计算 Duration
	if s.Time != nil {
		duration := time.Since(s.Time.StartAt).Seconds()
		summaryReport.Time = &automation.AutoReportTime{
			StartAt:  s.Time.StartAt,
			Duration: duration,
		}
	}

	// 3. 平台信息
	if s.Platform != nil {
		platform := datatypes.JSONMap{
			"httprunner_version": s.Platform.HttprunnerVersion,
			"go_version":         s.Platform.GoVersion,
			"platform":           s.Platform.Platform,
		}
		summaryReport.Platform = platform
	}

	// 4. 详情 & 步骤记录
	for _, cs := range s.Details {
		if cs == nil {
			continue
		}

		// 用例维度时间 &统计
		caseStat := datatypes.JSONMap{}
		if cs.Stat != nil {
			caseStat["total"] = cs.Stat.Total
			caseStat["successes"] = cs.Stat.Successes
			caseStat["failures"] = cs.Stat.Failures
		}

		caseTime := datatypes.JSONMap{}
		if cs.Time != nil {
			caseTime["start_at"] = cs.Time.StartAt
			caseTime["duration"] = cs.Time.Duration
		}

		inOut := datatypes.JSONMap{}
		if cs.InOut != nil {
			inOut["config_vars"] = cs.InOut.ConfigVars
			inOut["export_vars"] = cs.InOut.ExportVars
		}

		// 构建 AutoReportDetail
		name := cs.Name
		if cs.InOut != nil && cs.InOut.ConfigVars != nil {
			if v, ok := cs.InOut.ConfigVars["__case_name__"]; ok {
				if s, ok := v.(string); ok && s != "" {
					name = s
				}
			}
		}

		detail := automation.AutoReportDetail{
			Name:    name,
			Success: cs.Success,
			Stat:    caseStat,
			Time:    caseTime,
			InOut:   inOut,
		}

		// 步骤记录
		for _, step := range cs.Records {
			if step == nil {
				continue
			}

			httpStat := datatypes.JSONMap{}
			for k, v := range step.HttpStat {
				httpStat[k] = v
			}

			// Data 统一转成 map[string]interface{} 存到 JSONMap
			var dataMap datatypes.JSONMap
			if step.Data != nil {
				// 先尝试直接断言
				if m, ok := step.Data.(map[string]interface{}); ok {
					dataMap = datatypes.JSONMap(m)
				} else {
					// 回退到 JSON 编解码
					b, err := json.Marshal(step.Data)
					if err == nil {
						var tmp map[string]interface{}
						if err := json.Unmarshal(b, &tmp); err == nil {
							dataMap = datatypes.JSONMap(tmp)
						}
					}
				}
			}

			var exportVars datatypes.JSONMap
			if len(step.ExportVars) > 0 {
				exportVars = datatypes.JSONMap{}
				for k, v := range step.ExportVars {
					exportVars[k] = v
				}
			}

			record := automation.AutoReportRecord{
				Name:        step.Name,
				StartTime:   step.StartTime,
				StepType:    string(step.StepType),
				Success:     step.Success,
				ElapsedMs:   step.Elapsed,
				HttpStat:    httpStat,
				Data:        dataMap,
				ContentSize: step.ContentSize,
				ExportVars:  exportVars,
			}
			detail.Records = append(detail.Records, record)
		}

		summaryReport.Details = append(summaryReport.Details, detail)
	}

	// 复用已有的 UpdateReport 逻辑，将 summaryReport 合并到 r.report 并持久化
	r.UpdateReport(summaryReport)
}

func (r *ReportOperation) CreateReport() {
	if r.report == nil {
		return
	}
	err := global.GVA_DB.Create(r.report).Error
	if err != nil {
		global.GVA_LOG.Error("创建报告失败", zap.Error(err))
	}
}

// LoadReport 根据报告ID加载已存在的报告，并更新状态为运行中
func (r *ReportOperation) LoadReport(reportID uint) error {
	if reportID == 0 {
		return nil
	}

	var report automation.AutoReport
	if err := global.GVA_DB.First(&report, "id = ?", reportID).Error; err != nil {
		global.GVA_LOG.Error("加载报告失败", zap.Uint("report_id", reportID), zap.Error(err))
		return err
	}

	// 更新状态为运行中
	runningStatus := int64(automation.ReportStatusRunning)
	report.Status = &runningStatus
	if err := global.GVA_DB.Model(&report).Update("status", runningStatus).Error; err != nil {
		global.GVA_LOG.Error("更新报告状态失败", zap.Error(err))
		return err
	}

	r.report = &report
	return nil
}

func (r *ReportOperation) UpdateReport(report *automation.AutoReport) {
	if r.report == nil {
		return
	}

	// ID 必须存在，否则 GORM 会禁止无 WHERE 的批量更新
	if r.report.ID == 0 {
		global.GVA_LOG.Error("更新报告失败，ID 为空，已跳过以避免全表更新")
		return
	}

	// 获取 ProgressID：优先使用内存中的值，如果为空则从数据库重新读取
	// （Redis 模式下，finalizeReportProgress 已经更新了数据库中的 ProgressID）
	progressID := r.report.ProgressID
	if progressID == nil {
		var dbReport automation.AutoReport
		if err := global.GVA_DB.Select("progress_id").First(&dbReport, "id = ?", r.report.ID).Error; err == nil {
			progressID = dbReport.ProgressID
		}
	}

	// Update fields from latest execution result
	r.report.Success = report.Success
	var finishStatus int64 = 2
	if report.Success != nil && *report.Success {
		finishStatus = 3
	}
	r.report.Status = &finishStatus // 1: running, 2: fail, 3: success

	r.report.Time = report.Time
	r.report.Stat = report.Stat
	r.report.Platform = report.Platform
	r.report.Details = report.Details

	// 设置 ProgressID，确保不被覆盖
	r.report.ProgressID = progressID

	// 使用 Save + FullSaveAssociations：
	// 1）根据主键 ID 更新 lc_auto_reports 本身；
	// 2）自动创建 / 更新 Stat、Time、Details、Records 等所有关联表记录。
	err := global.GVA_DB.
		Session(&gorm.Session{FullSaveAssociations: true}).
		Save(r.report).Error
	if err != nil {
		global.GVA_LOG.Error("更新报告失败", zap.Error(err))
	}
}
