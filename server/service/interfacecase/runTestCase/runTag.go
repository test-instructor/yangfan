package runTestCase

import (
	"errors"
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/hrp"
	"github.com/test-instructor/cheetah/server/model/common/request"
	"github.com/test-instructor/cheetah/server/model/interfacecase"
	"testing"
)

func NewRunTag(runCaseReq request.RunCaseReq, runType interfacecase.RunType) TestCase {
	return &runTag{
		CaseID:     runCaseReq.CaseID,
		caseType:   interfacecase.CaseTypeTag,
		runCaseReq: runCaseReq,
		runType:    runType,
	}
}

type runTag struct {
	reportOperation *ReportOperation
	CaseID          uint
	runCaseReq      request.RunCaseReq
	runType         interfacecase.RunType
	caseType        interfacecase.CaseType
	tcm             ApisCaseModel
	d               debugTalkOperation
}

func (r *runTag) LoadCase() (err error) {

	var tags interfacecase.ApiTimerTaskTag
	db := global.GVA_DB.Model(interfacecase.ApiTimerTaskTag{})
	db.Preload("ApiTimerTask")
	db.First(&tags, "id = ?", r.runCaseReq.TagID)

	var testCaseList []interfacecase.HrpCase
	var reportName = tags.Name

	for _, v := range tags.ApiTimerTask {
		taskId := v.ID
		taskCase := taskSort(taskId)

		for _, c := range taskCase {
			var testCase interfacecase.HrpCase
			r.d.ProjectID = c.ApiCase.ProjectID
			r.d.ID = c.ApiTimerTaskId
			cases := caseSort(c.ApiCaseId)
			apiConfig, err := getConfig(c.ApiCase.RunConfigID)
			if err != nil {
				return errors.New("获取配置失败")
			}

			//设置前置套件
			if apiConfig.SetupCaseID != nil && *apiConfig.SetupCaseID != 0 {
				//r.tcm.SetupCase = true
				hrpCaseStep, err := getCaseStepHrp(*apiConfig.SetupCaseID)
				if err != nil {
					return err
				}
				testCase.TestSteps = append(testCase.TestSteps, *hrpCaseStep)
				testCase.Confing = *apiConfig
			}
			r.tcm.Config = *apiConfig
			testCase.Name = c.ApiTimerTask.Name + "-" + c.ApiCase.Name
			testCase.ID = c.ApiCase.ID
			for _, s := range cases {
				hrpCaseStep, err := getCaseStepHrp(s.ApiCaseStepId)
				if err != nil {
					return err
				}
				testCase.TestSteps = append(testCase.TestSteps, *hrpCaseStep)
				testCase.Confing = *apiConfig
			}
			testCaseList = append(testCaseList, testCase)
		}
	}

	r.d.ID = r.runCaseReq.ApiID
	r.d.RunDebugTalkFile()
	err = cheetahTestCaseToHrpCase(testCaseList, r.d.FilePath, &r.tcm)
	if err != nil {
		return errors.New("用例转换失败")
	}
	r.reportOperation = &ReportOperation{
		report: &interfacecase.ApiReport{
			Name:      reportName,
			CaseType:  r.caseType,
			RunType:   r.runType,
			ProjectID: r.d.ProjectID,
			SetupCase: r.tcm.SetupCase,
		},
	}
	r.reportOperation.CreateReport()
	return nil
}

func (r *runTag) RunCase() (err error) {
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

func (r *runTag) Report() (reports *interfacecase.ApiReport, err error) {
	if r.reportOperation.report == nil {
		return nil, errors.New("未获取到报告信息")
	}
	return r.reportOperation.report, nil
}
