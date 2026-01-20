package runTestCase

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/test-instructor/yangfan/httprunner/hrp"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/automation"
	projectmgrsvc "github.com/test-instructor/yangfan/server/v2/service/projectmgr"
	"go.uber.org/zap"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ReportOperation struct {
	report *automation.AutoReport
}

func stringifyAttachments(v interface{}) string {
	if v == nil {
		return ""
	}
	switch vv := v.(type) {
	case string:
		return vv
	case []byte:
		return string(vv)
	default:
		b, err := json.Marshal(vv)
		if err == nil {
			s := strings.TrimSpace(string(b))
			if s != "" && s != "null" && s != `""` {
				return s
			}
		}
		return fmt.Sprint(vv)
	}
}

func buildReportFromSummary(s *hrp.Summary) *automation.AutoReport {
	if s == nil {
		return nil
	}

	success := s.Success
	summaryReport := &automation.AutoReport{
		Success: &success,
	}

	if s.Stat != nil {
		caseStat := &automation.AutoReportStatTestcases{
			Total:   s.Stat.TestCases.Total,
			Success: s.Stat.TestCases.Success,
			Fail:    s.Stat.TestCases.Fail,
			Skip:    s.Stat.TestCases.Skipped,
		}

		stepStat := &automation.AutoReportStatTeststeps{
			Actions:   datatypes.JSONMap{},
			Total:     s.Stat.TestSteps.Total,
			Successes: s.Stat.TestSteps.Successes,
			Failures:  s.Stat.TestSteps.Failures,
			Skip:      s.Stat.TestSteps.Skipped,
		}

		apiStat := &automation.AutoReportStatTeststepapi{}

		apiTotal := 0
		apiSuccess := 0
		apiFail := 0
		apiSkip := 0
		for _, cs := range s.Details {
			if cs == nil {
				continue
			}
			for _, step := range cs.Records {
				if step == nil {
					continue
				}

				stepTypeStr := string(step.StepType)
				if strings.Contains(stepTypeStr, "request") || strings.Contains(stepTypeStr, "api") {
					apiTotal++
					if step.Skipped {
						apiSkip++
					} else if step.Success {
						apiSuccess++
					} else {
						apiFail++
					}
				}
			}
		}

		apiStat.Total = apiTotal
		apiStat.Success = apiSuccess
		apiStat.Fail = apiFail
		apiStat.Skip = apiSkip

		summaryReport.Stat = &automation.AutoReportStat{
			Testcases:   caseStat,
			Teststeps:   stepStat,
			Teststepapi: apiStat,
		}
	}

	if s.Time != nil {
		duration := s.Time.Duration
		if duration <= 0 {
			duration = time.Since(s.Time.StartAt).Seconds()
		}
		summaryReport.Time = &automation.AutoReportTime{
			StartAt:  s.Time.StartAt,
			Duration: duration,
		}
	}

	if s.Platform != nil {
		platform := datatypes.JSONMap{
			"httprunner_version": s.Platform.HttprunnerVersion,
			"go_version":         s.Platform.GoVersion,
			"platform":           s.Platform.Platform,
		}
		summaryReport.Platform = platform
	}

	for _, cs := range s.Details {
		if cs == nil {
			continue
		}

		caseStat := datatypes.JSONMap{}
		if cs.Stat != nil {
			caseStat["total"] = cs.Stat.Total
			caseStat["successes"] = cs.Stat.Successes
			caseStat["failures"] = cs.Stat.Failures
			caseStat["skip"] = cs.Stat.Skipped
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
			Skip:    cs.Skipped,
			Stat:    caseStat,
			Time:    caseTime,
			InOut:   inOut,
		}

		for _, step := range cs.Records {
			if step == nil {
				continue
			}

			httpStat := datatypes.JSONMap{}
			for k, v := range step.HttpStat {
				httpStat[k] = v
			}

			var dataMap datatypes.JSONMap
			if step.Data != nil {
				if m, ok := step.Data.(map[string]interface{}); ok {
					dataMap = datatypes.JSONMap(m)
				} else {
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
				Skip:        step.Skipped,
				ElapsedMs:   step.Elapsed,
				HttpStat:    httpStat,
				Data:        dataMap,
				Attachments: stringifyAttachments(step.Attachments),
				ContentSize: step.ContentSize,
				ExportVars:  exportVars,
			}
			detail.Records = append(detail.Records, record)
		}

		summaryReport.Details = append(summaryReport.Details, detail)
	}

	return summaryReport
}

// UpdateFromSummary 根据 hrp.Summary 计算并更新报告的用例 / 步骤 / 接口运行概要
// 参考 yangfan/run/runTestCase/report.go 的 resetReport 思路，按 yangfan 的 AutoReport 结构落库。
func (r *ReportOperation) UpdateFromSummary(s *hrp.Summary) {
	summaryReport := buildReportFromSummary(s)
	if summaryReport == nil {
		return
	}
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

	nodeName := r.report.NodeName
	if nodeName == nil {
		var dbReport automation.AutoReport
		if err := global.GVA_DB.Select("node_name").First(&dbReport, "id = ?", r.report.ID).Error; err == nil {
			nodeName = dbReport.NodeName
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
	r.report.NodeName = nodeName

	// 使用 Save + FullSaveAssociations：
	// 1）根据主键 ID 更新 yf_auto_reports 本身；
	// 2）自动创建 / 更新 Stat、Time、Details、Records 等所有关联表记录。
	err := global.GVA_DB.
		Session(&gorm.Session{FullSaveAssociations: true}).
		Save(r.report).Error
	if err != nil {
		global.GVA_LOG.Error("更新报告失败", zap.Error(err))
		return
	}

	reportID := r.report.ID
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := (&projectmgrsvc.ReportNotifyService{}).NotifyAutoReport(ctx, reportID); err != nil {
			global.GVA_LOG.Warn("发送报告通知失败", zap.Uint("report_id", reportID), zap.Error(err))
		}
	}()
}
