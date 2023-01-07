package runTestCase

import (
	"errors"
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/hrp"
	"github.com/test-instructor/cheetah/server/model/common/request"
	"github.com/test-instructor/cheetah/server/model/interfacecase"
	"gorm.io/gorm"
	"testing"
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
}

func (r *runStep) LoadCase() (err error) {
	var testCase interfacecase.HrpCase
	var testCaseList []interfacecase.HrpCase
	var apiCases interfacecase.ApiCaseStep
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
	}
	//获取运行配置
	apiConfig, err := getConfig(r.runCaseReq.ConfigID)
	if err != nil {
		return errors.New("获取配置失败")
	}

	//设置前置套件
	if apiConfig.SetupCaseID != nil && *apiConfig.SetupCaseID != 0 {
		//前置用例逻辑需要修改
		//r.tcm.SetupCase = true
		hrpCaseStep, err := getCaseStepHrp(*apiConfig.SetupCaseID)
		if err != nil {
			return err
		}
		testCase.TestSteps = append(testCase.TestSteps, *hrpCaseStep)
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
		First(&apiCases, "id = ?", r.runCaseReq.CaseID)

	{
		hrpCaseStep, err := getCaseStepHrp(r.runCaseReq.CaseID)
		if err != nil {
			return err
		}
		testCase.TestSteps = append(testCase.TestSteps, *hrpCaseStep)

		testCase.ID = hrpCaseStep.ID
		testCase.Name = hrpCaseStep.Name
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
			ProjectID: apiConfig.ProjectID,
			SetupCase: r.tcm.SetupCase,
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
