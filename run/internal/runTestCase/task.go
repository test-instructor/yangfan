package runTestCase

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/test-instructor/yangfan/httprunner/hrp"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/automation"
	"github.com/test-instructor/yangfan/server/v2/model/platform/request"
	"go.uber.org/zap"
)

// runTask.go logic

func NewRunTask(runCaseReq request.RunnerRequest, msg *string) TestCase {
	return &runTask{
		CaseID:     uint(runCaseReq.CaseID),
		runCaseReq: runCaseReq,
		msg:        msg,
	}
}

type runTask struct {
	reportOperation *ReportOperation
	CaseID          uint
	runCaseReq      request.RunnerRequest
	tcm             ApisCaseModel
	d               debugTalkOperation
	envVars         map[string]string
	msg             *string

	// 任务运行维度统计
	totalCases int
	totalSteps int
	totalApis  int
	envName    string
}

func (r *runTask) LoadCase() (err error) {
	var task automation.TimerTask
	// 定时任务本身使用 RunnerRequest.CaseID 作为任务 ID
	err = global.GVA_DB.Model(automation.TimerTask{}).
		Where("id = ? ", r.runCaseReq.CaseID).
		First(&task).Error
	if err != nil {
		return errors.New("获取定时任务失败")
	}

	// 获取任务关联的用例列表
	// yf_timer_task_case_list.task_id -> TimerTask.ID
	taskCaseList := taskSort(uint(r.CaseID))
	if len(taskCaseList) == 0 {
		return fmt.Errorf("定时任务未关联任何用例（task_id=%d）", r.runCaseReq.CaseID)
	}

	debugTalkStarted := false
	defer func() {
		if err != nil && debugTalkStarted {
			r.d.StopDebugTalkFile()
		}
	}()

	// 计算运行使用的环境 ID：
	// 优先级：
	// 1. RunnerRequest.EnvID（MQ 消息里显式指定的 env_id，例如调试任务时）；
	// 2. 定时任务本身配置的 EnvID；
	// 3. 第一个用例自身的 EnvID（兼容老数据）。
	var envID int64
	if r.runCaseReq.EnvID != 0 {
		// 来自 MQ 消息的 env_id
		envID = int64(r.runCaseReq.EnvID)
	} else if task.EnvID != nil && *task.EnvID != 0 {
		// 定时任务上配置的环境
		envID = int64(*task.EnvID)
	} else if len(taskCaseList) > 0 {
		// 回退到第一个用例自身的环境
		if taskCaseList[0].AutoCase.EnvID != 0 {
			envID = taskCaseList[0].AutoCase.EnvID
		}
	}

	// 根据最终确定的 envID 加载环境变量（可能为空）
	if envID != 0 {
		var envName string
		r.envVars, envName, err = GetEnvVar(task.ProjectId, envID)
		if err != nil {
			return errors.New("获取环境变量失败")
		}
		r.envName = envName
	}

	// DebugTalk，按项目维度创建调试脚本目录
	r.d.ProjectID = uint(task.ProjectId)
	r.d.ID = uint(r.runCaseReq.CaseID)
	if err := r.d.RunDebugTalkFile(); err != nil {
		return errors.New("准备DebugTalk环境失败")
	}
	debugTalkStarted = true

	// 遍历每个用例，转换为 ITestCase
	for _, taskCase := range taskCaseList {
		requestConfigID := uint(r.runCaseReq.ConfigID)
		caseConfigID := uint(taskCase.AutoCase.ConfigID)
		configID := requestConfigID
		if configID == 0 {
			configID = caseConfigID
		}

		apiConfig, configErr := getRunConfig(configID)
		if configErr != nil && requestConfigID != 0 && caseConfigID != 0 && caseConfigID != requestConfigID {
			global.GVA_LOG.Warn("RunTask config_id invalid, fallback to case config",
				zap.Uint("task_id", uint(r.runCaseReq.CaseID)),
				zap.Uint("case_id", taskCase.AutoCaseID),
				zap.Uint("request_config_id", requestConfigID),
				zap.Uint("case_config_id", caseConfigID),
				zap.Error(configErr))
			configID = caseConfigID
			apiConfig, configErr = getRunConfig(configID)
		}
		if configErr == nil && requestConfigID != 0 && apiConfig != nil {
			caseProjectID := taskCase.AutoCase.ProjectId
			if apiConfig.ProjectId != 0 && apiConfig.ProjectId != caseProjectID {
				global.GVA_LOG.Warn("RunTask config project mismatch, fallback to case config",
					zap.Uint("task_id", uint(r.runCaseReq.CaseID)),
					zap.Uint("case_id", taskCase.AutoCaseID),
					zap.Int64("config_project_id", apiConfig.ProjectId),
					zap.Int64("case_project_id", caseProjectID),
					zap.Uint("request_config_id", requestConfigID),
					zap.Uint("case_config_id", caseConfigID))
				if caseConfigID != 0 && caseConfigID != requestConfigID {
					apiConfig, configErr = getRunConfig(caseConfigID)
				} else {
					configErr = errors.New("配置与用例项目不匹配")
				}
			}
		}
		if configErr != nil {
			global.GVA_LOG.Warn("Failed to get config for case",
				zap.Uint("task_id", uint(r.runCaseReq.CaseID)),
				zap.Uint("case_id", taskCase.AutoCaseID),
				zap.Uint("config_id", configID),
				zap.Error(configErr))
			continue
		}

		// 转换配置为 httprunner TConfig
		tConfig := convertConfigToTConfig(apiConfig, r.envVars, taskCase.AutoCase.ProjectId, int(envID))
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

		// 汇总统计：使用精确的进度计算，考虑 Parameters 循环
		caseTotals := CalcCaseTotalsFromISteps(steps, tConfig)
		r.totalCases += caseTotals.TotalCases
		r.totalSteps += caseTotals.TotalSteps
		r.totalApis += caseTotals.TotalApis
	}
	if len(r.tcm.Case) == 0 {
		return fmt.Errorf("定时任务未生成任何可运行用例（task_id=%d, 关联用例数=%d），请检查用例配置或消息中的 config_id", r.runCaseReq.CaseID, len(taskCaseList))
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
		r.reportOperation.report.EnvName = r.envName
		r.reportOperation.report.Hostname = &hostname
	}
	if r.reportOperation.report != nil {
		// 使用精确的进度计算，考虑 Parameters 循环
		// 任务包含多个用例，需要累加每个用例的统计
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

func (r *runTask) RunCase() (err error) {
	var t *testing.T
	defer func() {
		if recovered := recover(); recovered != nil {
			global.GVA_LOG.Error("RunTask panic recovered", zap.Any("panic", recovered))
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

	global.GVA_LOG.Info("RunTask Finished", zap.String("report", string(reportHRP)))

	var summary hrp.Summary
	if err := json.Unmarshal(reportHRP, &summary); err != nil {
		global.GVA_LOG.Warn("Failed to unmarshal HRP report to Summary (task)", zap.Error(err))
	} else if r.reportOperation != nil {
		r.reportOperation.UpdateFromSummary(&summary)
	}

	return nil
}

func (r *runTask) Report() (interface{}, error) {
	if r.reportOperation == nil || r.reportOperation.report == nil {
		return nil, errors.New("未获取到报告信息")
	}
	return r.reportOperation.report, nil
}
