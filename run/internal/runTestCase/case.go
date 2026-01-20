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

func NewRunCase(runCaseReq request.RunnerRequest, msg *string) TestCase {
	return &runCase{
		CaseID:     uint(runCaseReq.CaseID),
		runCaseReq: runCaseReq,
		msg:        msg,
	}
}

type runCase struct {
	reportOperation *ReportOperation
	CaseID          uint
	runCaseReq      request.RunnerRequest
	tcm             ApisCaseModel
	d               debugTalkOperation
	envVars         map[string]string
	msg             *string
}

func (r *runCase) LoadCase() (err error) {
	var autoCase automation.AutoCase

	// Load AutoCase
	err = global.GVA_DB.Model(automation.AutoCase{}).Where("id = ? ", r.runCaseReq.CaseID).First(&autoCase).Error
	if err != nil {
		return err
	}

	// Config: 优先使用 mq 消息中的 ConfigID，否则使用用例自身的配置
	if r.runCaseReq.ConfigID == 0 {
		r.runCaseReq.ConfigID = int(autoCase.ConfigID)
	}
	if r.runCaseReq.EnvID == 0 {
		r.runCaseReq.EnvID = int(autoCase.EnvID)
	}

	apiConfig, err := getRunConfig(uint(r.runCaseReq.ConfigID))
	if err != nil {
		return errors.New("获取配置失败")
	}

	// Env Vars
	var envName string
	r.envVars, envName, err = GetEnvVar(autoCase.ProjectId, int64(r.runCaseReq.EnvID))
	if err != nil {
		return errors.New("获取环境变量失败")
	}

	// 转换配置为 httprunner TConfig
	tConfig := convertConfigToTConfig(apiConfig, r.envVars, autoCase.ProjectId, r.runCaseReq.EnvID)

	// 传递运行时上下文到 Variables
	if tConfig.Variables == nil {
		tConfig.Variables = make(map[string]interface{})
	}
	tConfig.Variables["platformProjectId"] = autoCase.ProjectId
	tConfig.Variables["platformEnvId"] = r.runCaseReq.EnvID

	// DebugTalk
	r.d.ProjectID = uint(autoCase.ProjectId)
	r.d.ID = uint(r.runCaseReq.CaseID)
	if err := r.d.RunDebugTalkFile(); err != nil {
		return errors.New("准备DebugTalk环境失败")
	}
	tConfig.Path = r.d.FilePath

	// 构建步骤列表
	var steps []hrp.IStep

	// SetupCase - 前置步骤
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

	// Load Case Steps - 加载用例关联的步骤
	caseStepList := caseSort(uint(r.runCaseReq.CaseID))
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
		ID:        autoCase.ID,
		Name:      autoCase.CaseName,
		Config:    tConfig,
		TestSteps: steps,
	}

	// 添加到用例列表
	r.tcm.Case = append(r.tcm.Case, yangfanTestCase)
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
		if r.reportOperation.report.NodeName == nil && r.runCaseReq.NodeName != "" {
			v := r.runCaseReq.NodeName
			r.reportOperation.report.NodeName = &v
		}
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

func (r *runCase) RunCase() (err error) {
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

	failfast := false
	if r.runCaseReq.Failfast != nil {
		failfast = *r.runCaseReq.Failfast
	}

	hrpRunner := hrp.NewRunner(t).
		SetHTTPStatOn().
		SetFailfast(failfast)

	reportID := uint(0)
	if r.reportOperation != nil && r.reportOperation.report != nil {
		reportID = r.reportOperation.report.ID
	}

	reportHRP, err := runCasesWithProgress(hrpRunner, reportID, r.tcm.Case...)
	if err != nil {
		return err
	}

	global.GVA_LOG.Debug("RunCase Finished", zap.String("report", string(reportHRP)))

	var summary hrp.Summary
	if err := json.Unmarshal(reportHRP, &summary); err != nil {
		global.GVA_LOG.Warn("Failed to unmarshal HRP report to Summary", zap.Error(err))
	} else if r.reportOperation != nil {
		r.reportOperation.UpdateFromSummary(&summary)
	}

	global.GVA_LOG.Info("RunCase Finished")
	return nil
}

func (r *runCase) Report() (interface{}, error) {
	if r.reportOperation == nil || r.reportOperation.report == nil {
		return nil, errors.New("未获取到报告信息")
	}
	return r.reportOperation.report, nil
}
