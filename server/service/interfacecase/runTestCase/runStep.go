package runTestCase

import (
	"errors"
	"testing"

	"gorm.io/gorm"

	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/hrp"
	"github.com/test-instructor/yangfan/server/model/common/request"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
)

func NewRunStep(runCaseReq request.RunCaseReq, runType interfacecase.RunType) TestCase {
	return &runStep{
		CaseID:     runCaseReq.CaseID,
		caseType:   interfacecase.CaseTypeStep,
		runCaseReq: runCaseReq,
		runType:    runType,
	}
}

type runStep struct {
	reportOperation *ReportOperation
	CaseID          uint
	runCaseReq      request.RunCaseReq
	runType         interfacecase.RunType
	caseType        interfacecase.CaseType
	tcm             ApisCaseModel
	d               debugTalkOperation
	envVars         map[string]string
}

func (r *runStep) LoadCase() (err error) {
	var testCase interfacecase.HrpCase
	var testCaseList []interfacecase.HrpCase
	var apiCases interfacecase.ApiCaseStep
	var envName string
	//var apiCases interfacecase.ApiCaseStep
	//var tcm *ApisCaseModel
	//获取测试套件下对应的配置信息
	{
		var testCaseStep interfacecase.ApiCaseStep
		err := global.GVA_DB.Model(interfacecase.ApiCaseStep{}).Where("id = ? ", r.runCaseReq.CaseID).First(&testCaseStep).Error
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
	//设置前置套件
	if apiConfig.SetupCaseID != nil && *apiConfig.SetupCaseID != 0 {
		//前置用例逻辑需要修改
		//r.tcm.SetupCase = true
		hrpCaseStep, err := getCaseStepHrp(*apiConfig.SetupCaseID)
		if err != nil {
			return err
		}
		if hrpCaseStep != nil {
			testCase.TestSteps = append(testCase.TestSteps, *hrpCaseStep)
		}
		testCase.Confing = *apiConfig
		//if hrpCaseStep.TestCase != nil && hrpCaseStep.Transaction != nil && hrpCaseStep.Rendezvous != nil && hrpCaseStep.ThinkTime != nil {
		//	testCase.TestSteps = append(testCase.TestSteps, *hrpCaseStep)
		//}

	}

	r.tcm.Config = *apiConfig

	//读取用例信息
	global.GVA_DB.Model(&interfacecase.ApiCaseStep{}).
		Preload("TStep", func(db2 *gorm.DB) *gorm.DB {
			return db2.Order("Sort")
		}).
		Preload("TStep.Request").
		Preload("TStep.Grpc").
		First(&apiCases, "id = ?", r.runCaseReq.CaseID)

	{
		hrpCaseStep, err := getCaseStepHrp(r.runCaseReq.CaseID)
		if err != nil {
			return err
		}
		if hrpCaseStep == nil {
			return errors.New("运行失败，请添加用例后再运行")
		}
		if hrpCaseStep != nil {
			testCase.TestSteps = append(testCase.TestSteps, *hrpCaseStep)
			testCase.ID = hrpCaseStep.ID
			testCase.Name = hrpCaseStep.Name
		}
		testCase.Confing = *apiConfig
	}
	testCaseList = append(testCaseList, testCase)
	//testcaseJson, _ := json.Marshal(testCase)
	//fmt.Printf(string(testcaseJson))
	r.d.ProjectID = apiConfig.ProjectID
	r.d.ID = r.runCaseReq.ApiID
	r.d.RunDebugTalkFile()
	err = cheetahTestCaseToHrpCase(testCaseList, r.d.FilePath, &r.tcm)
	if err != nil {
		return errors.New("用例转换失败")
	}
	r.reportOperation = &ReportOperation{
		report: &interfacecase.ApiReport{
			Name:      apiCases.Name,
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

func (r *runStep) RunCase() (err error) {
	var t *testing.T
	defer recoverHrp(r.reportOperation)
	defer r.d.StopDebugTalkFile()
	report, err := hrp.NewRunner(t).
		SetHTTPStatOn().
		SetFailfast(false).
		RunJsons(r.tcm.Case...)
	r.reportOperation.UpdateReport(&report)
	if err != nil {
		return err
	}
	return nil
}

func (r *runStep) Report() (report *interfacecase.ApiReport, err error) {
	if r.reportOperation.report == nil {
		return nil, errors.New("未获取到报告信息")
	}
	return r.reportOperation.report, nil
}
