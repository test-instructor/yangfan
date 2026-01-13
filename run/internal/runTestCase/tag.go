package runTestCase

import (
	"encoding/json"
	"errors"
	"os"
	"testing"

	"github.com/test-instructor/yangfan/httprunner/hrp"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/automation"
	"github.com/test-instructor/yangfan/server/v2/model/platform/request"
	"go.uber.org/zap"
)

func NewRunTag(runCaseReq request.RunnerRequest, msg *string) TestCase {
	return &runTag{
		TagID:      uint(runCaseReq.CaseID),
		runCaseReq: runCaseReq,
		msg:        msg,
	}
}

type runTag struct {
	reportOperation *ReportOperation
	TagID           uint
	runCaseReq      request.RunnerRequest
	tcm             ApisCaseModel
	d               debugTalkOperation
	envVars         map[string]string
	msg             *string

	// 用于统计标签运行所涉及的用例/步骤/接口总数
	totalCases int
	totalSteps int
	totalApis  int
}

func (r *runTag) LoadCase() (err error) {
	// 获取标签信息
	var tag automation.TimerTaskTag
	err = global.GVA_DB.Model(&automation.TimerTaskTag{}).Where("id = ?", r.TagID).First(&tag).Error
	if err != nil {
		return errors.New("获取标签失败")
	}

	// 获取环境变量
	var envName string
	r.envVars, envName, err = GetEnvVar(tag.ProjectId, int64(r.runCaseReq.EnvID))
	if err != nil {
		return errors.New("获取环境变量失败")
	}
	_ = envName

	// DebugTalk
	r.d.ProjectID = uint(tag.ProjectId)
	r.d.ID = r.TagID
	r.d.RunDebugTalkFile()

	// 查找包含此标签的定时任务
	var tasks []automation.TimerTask
	err = global.GVA_DB.Model(&automation.TimerTask{}).
		Where("project_id = ? AND JSON_CONTAINS(tag, ?)", tag.ProjectId, r.TagID).
		Find(&tasks).Error
	if err != nil {
		// 尝试其他查询方式
		err = global.GVA_DB.Model(&automation.TimerTask{}).
			Where("project_id = ?", tag.ProjectId).
			Find(&tasks).Error
		if err != nil {
			return errors.New("查询关联任务失败")
		}
	}

	// 筛选包含指定标签的任务
	var filteredTasks []automation.TimerTask
	for _, task := range tasks {
		if task.Tag != nil {
			var tagIDs []uint
			if err := json.Unmarshal(task.Tag, &tagIDs); err == nil {
				for _, tagID := range tagIDs {
					if tagID == r.TagID {
						filteredTasks = append(filteredTasks, task)
						break
					}
				}
			}
		}
	}

	// 遍历每个任务，加载其关联的用例
	for _, task := range filteredTasks {
		taskCaseList := taskSort(task.ID)

		for _, taskCase := range taskCaseList {
			// 获取用例配置: 优先使用 mq 消息中的 ConfigID，否则使用用例自身的配置
			configID := uint(r.runCaseReq.ConfigID)
			if configID == 0 {
				configID = uint(taskCase.AutoCase.ConfigID)
			}
			apiConfig, err := getRunConfig(configID)
			if err != nil {
				global.GVA_LOG.Warn("Failed to get config for case",
					zap.Uint("case_id", taskCase.AutoCaseID),
					zap.Error(err))
				continue
			}

			// 转换配置为 httprunner TConfig
			tConfig := convertConfigToTConfig(apiConfig, r.envVars, taskCase.AutoCase.ProjectId, r.runCaseReq.EnvID)
			tConfig.Path = r.d.FilePath

			// 构建步骤列表
			var steps []hrp.IStep

			// 前置步骤
			if apiConfig.PreparatoryStepsID != 0 {
				var setupCaseStep automation.AutoCaseStep
				err := global.GVA_DB.Model(&automation.AutoCaseStep{}).
					First(&setupCaseStep, "id = ?", apiConfig.PreparatoryStepsID).Error
				if err == nil {
					setupStep, err := convertAutoCaseStepToIStep(&setupCaseStep, tConfig, false, true, r.envVars)
					if err == nil && setupStep != nil {
						steps = append(steps, setupStep)
						r.tcm.SetupCase = true
					}
				}
			}

			// 加载用例关联的步骤
			caseStepList := caseSort(taskCase.AutoCaseID)
			for _, caseStep := range caseStepList {
				var autoCaseStep automation.AutoCaseStep
				err := global.GVA_DB.Model(&automation.AutoCaseStep{}).
					First(&autoCaseStep, "id = ?", caseStep.AutoCaseStepID).Error
				if err != nil {
					global.GVA_LOG.Warn("Failed to load AutoCaseStep",
						zap.Uint("step_id", caseStep.AutoCaseStepID),
						zap.Error(err))
					continue
				}

				step, err := convertAutoCaseStepToIStep(&autoCaseStep, tConfig, caseStep.IsConfig, caseStep.IsStepConfig, r.envVars)
				if err != nil {
					global.GVA_LOG.Warn("Failed to convert AutoCaseStep",
						zap.Uint("step_id", caseStep.AutoCaseStepID),
						zap.Error(err))
					continue
				}
				if step != nil {
					steps = append(steps, step)
				}
			}

			// 构建 YangfanTestCase
			yangfanTestCase := &YangfanTestCase{
				ID:        taskCase.AutoCase.ID,
				Name:      taskCase.AutoCase.CaseName,
				Config:    tConfig,
				TestSteps: steps,
			}

			// 添加到用例列表
			r.tcm.Case = append(r.tcm.Case, yangfanTestCase)
			r.tcm.Config = tConfig

			// 汇总标签下的用例/步骤/接口数量：使用精确的进度计算
			caseTotals := CalcCaseTotalsFromISteps(steps, tConfig)
			r.totalCases += caseTotals.TotalCases
			r.totalSteps += caseTotals.TotalSteps
			r.totalApis += caseTotals.TotalApis
		}
	}

	if len(r.tcm.Case) == 0 {
		return errors.New("未找到关联的用例")
	}

	// 加载已存在的报告并更新状态为运行中
	hostname, _ := os.Hostname()

	r.reportOperation = &ReportOperation{}
	if r.runCaseReq.ReportID != 0 {
		if err := r.reportOperation.LoadReport(r.runCaseReq.ReportID); err != nil {
			global.GVA_LOG.Error("加载报告失败", zap.Uint("report_id", r.runCaseReq.ReportID), zap.Error(err))
		}
	}
	// 补充报告信息
	if r.reportOperation.report != nil {
		r.reportOperation.report.Hostname = &hostname
		r.reportOperation.report.SetupCase = &r.tcm.SetupCase
	}
	if r.reportOperation.report != nil {
		// 使用精确的进度计算，考虑 Parameters 循环
		// 标签包含多个用例，需要累加每个用例的统计
		totals := ReportProgressTotals{
			TotalCases: r.totalCases,
			TotalSteps: r.totalSteps,
			TotalApis:  r.totalApis,
		}
		progressID := initReportProgress(r.reportOperation.report.ID, totals)
		if progressID != 0 {
			r.reportOperation.report.ProgressID = &progressID
		}
	}

	return nil
}

func (r *runTag) RunCase() (err error) {
	var t *testing.T
	defer func() {
		if recovered := recover(); recovered != nil {
			global.GVA_LOG.Error("RunTag panic recovered", zap.Any("panic", recovered))
			if r.reportOperation != nil && r.reportOperation.report != nil {
				failStatus := int64(2)
				success := false
				r.reportOperation.report.Status = &failStatus
				r.reportOperation.report.Success = &success
				r.reportOperation.UpdateReport(r.reportOperation.report)
			}
		}
		r.d.StopDebugTalkFile()
	}()

	reportID := uint(0)
	if r.reportOperation != nil && r.reportOperation.report != nil {
		reportID = r.reportOperation.report.ID
	}

	reportHRP, err := runCasesWithProgress(
		hrp.NewRunner(t).
			SetHTTPStatOn().
			SetFailfast(false),
		reportID,
		r.tcm.Case...,
	)

	if err != nil {
		return err
	}

	global.GVA_LOG.Info("RunTag Finished", zap.String("report", string(reportHRP)))

	var summary hrp.Summary
	if err := json.Unmarshal(reportHRP, &summary); err != nil {
		global.GVA_LOG.Warn("Failed to unmarshal HRP report to Summary", zap.Error(err))
	} else if r.reportOperation != nil {
		r.reportOperation.UpdateFromSummary(&summary)
	}

	return nil
}

func (r *runTag) Report() (interface{}, error) {
	if r.reportOperation == nil || r.reportOperation.report == nil {
		return nil, errors.New("未获取到报告信息")
	}
	return r.reportOperation.report, nil
}
