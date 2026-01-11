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

func NewRunApi(runCaseReq request.RunnerRequest, msg *string) TestCase {
	return &runAPI{
		ApiID:      uint(runCaseReq.CaseID),
		runCaseReq: runCaseReq,
		msg:        msg,
	}
}

type runAPI struct {
	reportOperation *ReportOperation
	ApiID           uint
	runCaseReq      request.RunnerRequest
	tcm             ApisCaseModel
	d               debugTalkOperation
	envVars         map[string]string
	msg             *string
}

func (r *runAPI) LoadCase() (err error) {
	var apiStep automation.AutoStep
	var envName string

	// 获取运行配置
	apiConfig, err := getRunConfig(uint(r.runCaseReq.ConfigID))
	if err != nil {
		return errors.New("获取配置失败")
	}

	// 获取环境变量
	r.envVars, envName, err = GetEnvVar(apiConfig.ProjectId, int64(r.runCaseReq.EnvID))
	if err != nil {
		return errors.New("获取环境变量失败")
	}

	// 转换配置为 httprunner TConfig
	tConfig := convertConfigToTConfig(apiConfig, r.envVars, apiConfig.ProjectId, r.runCaseReq.EnvID)

	// 传递运行时上下文到 Variables
	if tConfig.Variables == nil {
		tConfig.Variables = make(map[string]interface{})
	}
	tConfig.Variables["platformProjectId"] = apiConfig.ProjectId
	tConfig.Variables["platformEnvId"] = r.runCaseReq.EnvID

	// DebugTalk
	r.d.ProjectID = uint(apiConfig.ProjectId)
	r.d.ID = r.ApiID
	r.d.RunDebugTalkFile()
	tConfig.Path = r.d.FilePath

	// 加载 API Step
	err = global.GVA_DB.Model(&automation.AutoStep{}).
		Preload("Request").
		First(&apiStep, "id = ?", r.runCaseReq.CaseID).Error
	if err != nil {
		return err
	}

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

	// 转换 API 步骤为 IStep
	apiIStep := convertAutoStepToIStep(&apiStep, apiConfig.ProjectId, r.runCaseReq.EnvID)
	if apiIStep == nil {
		return errors.New("转换 API 步骤失败")
	}
	steps = append(steps, apiIStep)

	// 包装为虚拟用例
	lingceTestCase := wrapStepsInVirtualTestCase(steps, tConfig, apiStep.StepName)
	lingceTestCase.ID = apiStep.ID

	// 添加到用例列表
	r.tcm.Case = append(r.tcm.Case, lingceTestCase)
	r.tcm.Config = tConfig

	hostname, _ := os.Hostname()

	r.reportOperation = &ReportOperation{}
	// 加载已存在的报告并更新状态为运行中
	if r.runCaseReq.ReportID != 0 {
		if err := r.reportOperation.LoadReport(r.runCaseReq.ReportID); err != nil {
			global.GVA_LOG.Error("加载报告失败", zap.Uint("report_id", r.runCaseReq.ReportID), zap.Error(err))
		}
	}
	// 补充报告信息
	if r.reportOperation.report != nil {
		r.reportOperation.report.EnvName = envName
		r.reportOperation.report.Hostname = &hostname
		r.reportOperation.report.SetupCase = &r.tcm.SetupCase
		r.reportOperation.report.ConfigName = r.tcm.Config.Name
	}
	if r.reportOperation.report != nil {
		// 使用精确的进度计算，考虑 Parameters 循环
		progressTotals := CalcCaseTotalsFromISteps(steps, r.tcm.Config)
		totals := ReportProgressTotals{
			TotalCases: progressTotals.TotalCases,
			TotalSteps: progressTotals.TotalSteps,
			TotalApis:  progressTotals.TotalApis,
		}
		progressID := initReportProgress(r.reportOperation.report.ID, totals)
		if progressID != 0 {
			r.reportOperation.report.ProgressID = &progressID
		}
	}
	return nil
}

func (r *runAPI) RunCase() (err error) {
	var t *testing.T

	defer func() {
		if recovered := recover(); recovered != nil {
			global.GVA_LOG.Error("RunCase panic recovered", zap.Any("panic", recovered))
			if r.reportOperation != nil && r.reportOperation.report != nil {
				failStatus := int64(2) // Fail
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

	global.GVA_LOG.Info("RunApi Finished", zap.String("report", string(reportHRP)))

	var summary hrp.Summary
	if err := json.Unmarshal(reportHRP, &summary); err != nil {
		global.GVA_LOG.Warn("Failed to unmarshal HRP report to Summary", zap.Error(err))
	} else if r.reportOperation != nil {
		r.reportOperation.UpdateFromSummary(&summary)
	}

	global.GVA_LOG.Debug("debugtalk 目录")
	global.GVA_LOG.Debug(r.d.FilePath)

	return nil
}

func (r *runAPI) Report() (interface{}, error) {
	if r.reportOperation == nil || r.reportOperation.report == nil {
		return nil, errors.New("未获取到报告信息")
	}
	return r.reportOperation.report, nil
}
