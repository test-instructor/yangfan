package runTestCase

import (
	"encoding/json"
	"errors"
	"os"
	"testing"

	"gorm.io/gorm"

	"github.com/test-instructor/yangfan/hrp"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/common/request"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
)

func NewRunCase(runCaseReq request.RunCaseReq, runType interfacecase.RunType) TestCase {
	return &runCase{
		CaseID:     runCaseReq.CaseID,
		caseType:   interfacecase.CaseTypeCases,
		runCaseReq: runCaseReq,
		runType:    runType,
	}
}

type runCase struct {
	reportOperation *ReportOperation
	CaseID          uint
	runCaseReq      request.RunCaseReq
	runType         interfacecase.RunType
	caseType        interfacecase.CaseType
	tcm             ApisCaseModel
	d               debugTalkOperation
	envVars         map[string]string
}

func (r *runCase) LoadCase() (err error) {
	var testCase interfacecase.HrpCase
	var testCaseList []interfacecase.HrpCase
	var apiCase interfacecase.ApiCase
	var apiCaseCase []interfacecase.ApiCaseRelationship
	var envName string

	//获取测试用例下对应的配置信息
	{
		var testCase interfacecase.ApiCase
		err := global.GVA_DB.Model(interfacecase.ApiCase{}).Where("id = ? ", r.runCaseReq.CaseID).First(&testCase).Error
		if err != nil {
			return err
		}
		r.runCaseReq.ConfigID = testCase.RunConfigID
		r.runCaseReq.Env = testCase.ApiEnvID
	}
	//获取运行配置
	apiConfig, err := getConfig(r.runCaseReq.ConfigID)
	if err != nil {
		return errors.New("获取配置失败")
	}
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

	//读取用例信息
	apiCase.ID = r.runCaseReq.CaseID
	err = global.GVA_DB.Model(interfacecase.ApiCase{}).First(&apiCase).Error
	caseDB := global.GVA_DB.Model(interfacecase.ApiCaseRelationship{}).
		Preload("ApiCaseStep").
		Preload("ApiCaseStep.TStep", func(db2 *gorm.DB) *gorm.DB {
			return db2.Order("Sort")
		}).
		Preload("ApiCaseStep.TStep.Request").
		Where("api_case_id = ?", r.runCaseReq.CaseID).
		Order("Sort")
	caseDB.Find(&apiCaseCase)
	for _, v := range apiCaseCase {
		//testCaseList = append(testCaseList, v.ApiCaseStep)
		hrpCaseStep, err := getCaseStepHrp(v.ApiCaseStepId)
		if err != nil {
			return err
		}
		if hrpCaseStep != nil {
			testCase.TestSteps = append(testCase.TestSteps, *hrpCaseStep)
		}
		testCase.Confing = *apiConfig
	}
	testCaseList = append(testCaseList, testCase)
	r.d.ProjectID = apiConfig.ProjectID
	r.d.ID = r.runCaseReq.ApiID
	r.d.RunDebugTalkFile()
	err = yangfanTestCaseToHrpCase(testCaseList, r.d.FilePath, &r.tcm)
	if err != nil {
		return errors.New("用例转换失败")
	}
	hostname, _ := os.Hostname()
	r.reportOperation = &ReportOperation{
		report: &interfacecase.ApiReport{
			Name:      apiCase.Name,
			CaseType:  r.caseType,
			RunType:   r.runType,
			SetupCase: r.tcm.SetupCase,
			Operator: interfacecase.Operator{
				ProjectID: apiConfig.ProjectID,
			},
			ApiEnvName: envName,
			ApiEnvID:   r.runCaseReq.Env,
			Hostname:   hostname,
		},
	}
	r.reportOperation.CreateReport()
	return nil
}

func (r *runCase) RunCase() (err error) {
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
	if err != nil {
		return err
	}
	return nil
}

func (r *runCase) Report() (report *interfacecase.ApiReport, err error) {
	if r.reportOperation.report == nil {
		return nil, errors.New("未获取到报告信息")
	}
	return r.reportOperation.report, nil
}
