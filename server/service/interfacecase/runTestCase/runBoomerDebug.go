package runTestCase

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"gorm.io/gorm"

	"github.com/test-instructor/yangfan/hrp"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/common/request"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
)

func NewBoomerDebug(runCaseReq request.RunCaseReq, runType interfacecase.RunType) TestCase {
	return &runBoomerDebug{
		CaseID:     runCaseReq.CaseID,
		caseType:   interfacecase.CaseTypeBoomerDebug,
		runCaseReq: runCaseReq,
		runType:    runType,
	}
}

type runBoomerDebug struct {
	reportOperation *ReportOperation
	CaseID          uint
	runCaseReq      request.RunCaseReq
	runType         interfacecase.RunType
	caseType        interfacecase.CaseType
	tcm             ApisCaseModel
	d               debugTalkOperation
	envVars         map[string]string
}

func (r *runBoomerDebug) LoadCase() (err error) {
	//获取运行配置

	var testCase interfacecase.HrpCase
	var testCaseList []interfacecase.HrpCase
	var apiCase interfacecase.Performance
	var apiCaseCase []interfacecase.PerformanceRelationship
	var envName string

	{
		var testCaseStep interfacecase.Performance
		err := global.GVA_DB.Model(interfacecase.Performance{}).Where("id = ? ", r.runCaseReq.CaseID).First(&testCaseStep).Error
		if err != nil {
			return err
		}
		r.runCaseReq.ConfigID = testCaseStep.RunConfigID
		r.runCaseReq.Env = testCaseStep.ApiEnvID
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
	global.GVA_LOG.Debug(fmt.Sprintf("boomer debug 1 apiConfig:%d", apiConfig.ID))

	//设置前置套件
	if apiConfig.SetupCaseID != nil && *apiConfig.SetupCaseID != 0 {
		global.GVA_LOG.Debug(fmt.Sprintf("boomer debug 2 apiConfig.SetupCaseID %d", *apiConfig.SetupCaseID))
		hrpCaseStep, err := getCaseStepHrp(*apiConfig.SetupCaseID)
		if err != nil {
			global.GVA_LOG.Debug(fmt.Sprintf("boomer debug 8 已设置前置套件，%s", err))
		}
		testCase.Confing = *apiConfig
		testCase.TestSteps = append(testCase.TestSteps, hrpCaseStep)
	}

	global.GVA_LOG.Debug(fmt.Sprintf("boomer debug 3 已设置前置套件，%d", apiConfig.ID))
	r.tcm.Config = *apiConfig

	//读取用例信息
	apiCase.ID = r.runCaseReq.CaseID
	err = global.GVA_DB.Model(interfacecase.Performance{}).First(&apiCase).Error
	global.GVA_LOG.Debug(fmt.Sprintf("boomer debug 4 apiCase，%v", apiCase))
	caseDB := global.GVA_DB.Model(interfacecase.PerformanceRelationship{}).
		Preload("ApiCaseStep").
		Preload("ApiCaseStep.TStep", func(db2 *gorm.DB) *gorm.DB {
			return db2.Order("Sort")
		}).
		Preload("ApiCaseStep.TStep.Request").
		Where("performance_id = ?", r.runCaseReq.CaseID).
		Order("Sort")
	caseDB.Find(&apiCaseCase)
	global.GVA_LOG.Debug(fmt.Sprintf("boomer debug 5 获取用例信息，%v", apiCaseCase))
	for _, v := range apiCaseCase {
		hrpCaseStep, err := getCaseStepHrp(v.ApiCaseStepId)
		if err != nil {
			return err
		}
		testCase.TestSteps = append(testCase.TestSteps, hrpCaseStep)
		global.GVA_LOG.Debug(fmt.Sprintf("boomer debug 8 *apiConfig.SetupCaseID:%v", hrpCaseStep))
	}
	testCase.Confing = *apiConfig
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
			Name:      apiCase.Name,
			CaseType:  r.caseType,
			RunType:   r.runType,
			SetupCase: r.tcm.SetupCase,
			Operator: interfacecase.Operator{
				ProjectID: apiConfig.ProjectID,
			},
			ApiEnvName: envName,
			ApiEnvID:   r.runCaseReq.Env,
		},
	}
	r.reportOperation.CreateReport()
	return nil
}

func (r *runBoomerDebug) RunCase() (err error) {
	var t *testing.T
	defer recoverHrp(r.reportOperation)
	defer r.d.StopDebugTalkFile()
	global.GVA_LOG.Debug(fmt.Sprintf("r.tcm.Case,%v", r.tcm.Case))
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

func (r *runBoomerDebug) Report() (report *interfacecase.ApiReport, err error) {
	if r.reportOperation.report == nil {
		return nil, errors.New("未获取到报告信息")
	}
	return r.reportOperation.report, nil
}
