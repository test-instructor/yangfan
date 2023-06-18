package runTestCase

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/test-instructor/yangfan/hrp"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/common/request"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
)

func NewRunApi(runCaseReq request.RunCaseReq, runType interfacecase.RunType) TestCase {
	return &runAPI{
		ApiID:      runCaseReq.ApiID,
		caseType:   interfacecase.CaseTypeApi,
		runCaseReq: runCaseReq,
		runType:    runType,
	}
}

type runAPI struct {
	reportOperation *ReportOperation
	ApiID           uint
	runCaseReq      request.RunCaseReq
	runType         interfacecase.RunType
	caseType        interfacecase.CaseType
	tcm             ApisCaseModel
	d               debugTalkOperation
	envVars         map[string]string
}

func (r *runAPI) LoadCase() (err error) {
	var apiStep interfacecase.ApiStep
	//获取运行配置
	var testCase interfacecase.HrpCase
	var testCaseList []interfacecase.HrpCase
	var envName string

	//获取运行配置
	apiConfig, err := getConfig(r.runCaseReq.ConfigID)
	if err != nil {
		return errors.New("获取配置失败")
	}
	//获取环境变量
	r.envVars, envName, err = GetEnvVar(apiConfig.ProjectID, r.runCaseReq.Env)
	if err != nil {
		return errors.New("获取环境变量失败")
	}
	apiConfig.Environs = r.envVars
	//设置前置步骤
	if apiConfig.SetupCaseID != nil && *apiConfig.SetupCaseID != 0 {
		hrpCaseStep, err := getCaseStepHrp(*apiConfig.SetupCaseID)
		if err != nil {
			return err
		}
		if hrpCaseStep != nil {
			testCase.TestSteps = append(testCase.TestSteps, *hrpCaseStep)
		}
		testCase.Confing = *apiConfig
	}
	r.tcm.Config = *apiConfig

	global.GVA_DB.Model(&interfacecase.ApiStep{}).
		Preload("Request").
		Preload("Grpc").
		First(&apiStep, "id = ?", r.runCaseReq.CaseID)
	testCase.Name = apiStep.Name
	var hrpTestCase interfacecase.HrpTestCase
	hrpTestCase.Name = apiStep.Name
	hrpTestCase.ID = apiStep.ID
	hrpTestCase.Confing = *apiConfig
	hrpTestCase.TestSteps = append(hrpTestCase.TestSteps, apiStep)
	hrpCase := &interfacecase.HrpCaseStep{
		ID:       hrpTestCase.ID,
		Name:     hrpTestCase.Name,
		TestCase: hrpTestCase,
	}
	testCase.Confing = *apiConfig
	testCase.TestSteps = append(testCase.TestSteps, *hrpCase)
	testCaseList = append(testCaseList, testCase)
	r.d.ProjectID = apiConfig.ProjectID
	r.d.ID = r.runCaseReq.ApiID
	r.d.RunDebugTalkFile()
	err = yangfanTestCaseToHrpCase(testCaseList, r.d.FilePath, &r.tcm)
	if err != nil {
		return errors.New("用例转换失败")
	}
	r.reportOperation = &ReportOperation{
		report: &interfacecase.ApiReport{
			Name:      apiStep.Name,
			CaseType:  r.caseType,
			RunType:   r.runType,
			SetupCase: r.tcm.SetupCase,
			Operator: interfacecase.Operator{
				ProjectID: apiStep.ProjectID,
			},
			ApiEnvName: envName,
			ApiEnvID:   r.runCaseReq.Env,
		},
	}
	r.reportOperation.CreateReport()
	return nil
}

func (r *runAPI) RunCase() (err error) {
	var t *testing.T
	defer recoverHrp(r.reportOperation)
	defer r.d.StopDebugTalkFile()
	reportHRP, err := hrp.NewRunner(t).
		SetHTTPStatOn().
		SetFailfast(false).
		RunJsons(r.tcm.Case...)
	var report interfacecase.ApiReport
	json.Unmarshal(reportHRP, &report)
	r.reportOperation.UpdateReport(&report)
	global.GVA_LOG.Debug("debugtalk 目录")
	global.GVA_LOG.Debug(r.d.FilePath)
	return nil
}

func (r *runAPI) Report() (report *interfacecase.ApiReport, err error) {
	if r.reportOperation.report == nil {
		return nil, errors.New("未获取到报告信息")
	}
	return r.reportOperation.report, nil
}
