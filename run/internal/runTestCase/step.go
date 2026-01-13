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

func NewRunStep(runCaseReq request.RunnerRequest, msg *string) TestCase {
	return &runStep{
		CaseID:     uint(runCaseReq.CaseID),
		runCaseReq: runCaseReq,
		msg:        msg,
	}
}

type runStep struct {
	reportOperation *ReportOperation
	CaseID          uint
	runCaseReq      request.RunnerRequest
	tcm             ApisCaseModel
	d               debugTalkOperation
	envVars         map[string]string
	msg             *string
}

func (r *runStep) LoadCase() (err error) {
	var autoCaseStep automation.AutoCaseStep

	var envName string

	// Get AutoCaseStep to determine Config and Env
	err = global.GVA_DB.Model(&automation.AutoCaseStep{}).
		Where("id = ?", r.CaseID).
		First(&autoCaseStep).Error
	if err != nil {
		return err
	}

	//r.runCaseReq.ConfigID = int(autoCaseStep.ConfigID)
	//r.runCaseReq.EnvID = int(autoCaseStep.EnvID)

	// Get Config
	apiConfig, err := getRunConfig(uint(r.runCaseReq.ConfigID))
	if err != nil {
		return errors.New("获取配置失败")
	}

	// Get Env Vars
	r.envVars, envName, err = GetEnvVar(apiConfig.ProjectId, int64(r.runCaseReq.EnvID))
	if err != nil {
		return errors.New("获取环境变量失败")
	}

	// 转换配置为 httprunner TConfig
	tConfig := convertConfigToTConfig(apiConfig, r.envVars, apiConfig.ProjectId, r.runCaseReq.EnvID)

	// DebugTalk
	r.d.ProjectID = uint(apiConfig.ProjectId)
	r.d.ID = r.CaseID
	r.d.RunDebugTalkFile()
	tConfig.Path = r.d.FilePath

	// 构建步骤列表
	var steps []hrp.IStep

	// Setup Case (Preparatory Step) - 前置步骤
	if apiConfig.PreparatoryStepsID != 0 {
		var setupCaseStep automation.AutoCaseStep
		err := global.GVA_DB.Model(&automation.AutoCaseStep{}).
			First(&setupCaseStep, "id = ?", apiConfig.PreparatoryStepsID).Error
		if err == nil {
			// 确保有正确的 ProjectId 和 EnvID
			if setupCaseStep.ProjectId == 0 {
				setupCaseStep.ProjectId = apiConfig.ProjectId
			}
			if setupCaseStep.EnvID == 0 {
				setupCaseStep.EnvID = int64(r.runCaseReq.EnvID)
			}
			setupStep, err := convertAutoCaseStepToIStep(&setupCaseStep, tConfig, false, true, r.envVars)
			if err == nil && setupStep != nil {
				steps = append(steps, setupStep)
				r.tcm.SetupCase = true
			}
		}
	}

	// 确保 autoCaseStep 有正确的 ProjectId 和 EnvID
	if autoCaseStep.ProjectId == 0 {
		autoCaseStep.ProjectId = apiConfig.ProjectId
	}
	if autoCaseStep.EnvID == 0 {
		autoCaseStep.EnvID = int64(r.runCaseReq.EnvID)
	}

	// 转换主步骤 (作为嵌套步骤)
	mainStep, err := convertAutoCaseStepToIStep(&autoCaseStep, tConfig, false, true, r.envVars)
	if err != nil {
		return errors.New("转换步骤失败")
	}
	if mainStep == nil {
		return errors.New("运行失败，请添加用例后再运行")
	}
	steps = append(steps, mainStep)

	// 包装为虚拟用例
	yangfanTestCase := wrapStepsInVirtualTestCase(steps, tConfig, autoCaseStep.StepName)
	yangfanTestCase.ID = autoCaseStep.ID

	// 添加到用例列表
	r.tcm.Case = append(r.tcm.Case, yangfanTestCase)
	r.tcm.Config = tConfig

	// Report
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

func (r *runStep) RunCase() (err error) {
	var t *testing.T
	defer func() {
		if recovered := recover(); recovered != nil {
			global.GVA_LOG.Error("RunCase panic recovered", zap.Any("panic", recovered))
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

	global.GVA_LOG.Info("RunStep Finished", zap.String("report", string(reportHRP)))

	var summary hrp.Summary
	if err := json.Unmarshal(reportHRP, &summary); err != nil {
		global.GVA_LOG.Warn("Failed to unmarshal HRP report to Summary", zap.Error(err))
	} else if r.reportOperation != nil {
		r.reportOperation.UpdateFromSummary(&summary)
	}

	return nil
}

func (r *runStep) Report() (interface{}, error) {
	if r.reportOperation == nil || r.reportOperation.report == nil {
		return nil, errors.New("未获取到报告信息")
	}
	return r.reportOperation.report, nil
}
