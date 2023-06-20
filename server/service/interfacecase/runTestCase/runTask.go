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

func NewRunTask(runCaseReq request.RunCaseReq, runType interfacecase.RunType) TestCase {
	return &runTask{
		CaseID:     runCaseReq.CaseID,
		caseType:   interfacecase.CaseTypeTask,
		runCaseReq: runCaseReq,
		runType:    runType,
	}
}

type runTask struct {
	reportOperation *ReportOperation
	CaseID          uint
	runCaseReq      request.RunCaseReq
	runType         interfacecase.RunType
	caseType        interfacecase.CaseType
	tcm             ApisCaseModel
	d               debugTalkOperation
	envVars         map[string]string
}

func (r *runTask) LoadCase() (err error) {
	var testCaseList []interfacecase.HrpCase
	var task interfacecase.ApiTimerTask
	err = global.GVA_DB.Model(interfacecase.ApiTimerTask{}).Where("id = ? ", r.runCaseReq.TaskID).First(&task).Error
	if err != nil {
		return errors.New("获取定时任务失败")
	}
	var reportName string
	var envName string
	taskCase := taskSort(r.runCaseReq.TaskID)
	r.envVars, envName, err = GetEnvVar(task.ProjectID, task.ApiEnvID)
	if err != nil {
		return errors.New("获取环境变量失败")
	}
	for _, c := range taskCase {
		var testCase interfacecase.HrpCase
		reportName = c.ApiTimerTask.Name
		r.d.ProjectID = c.ApiCase.ProjectID
		r.d.ID = c.ApiTimerTaskId
		cases := caseSort(c.ApiCaseId)
		apiConfig, err := getConfig(c.ApiCase.RunConfigID)
		if err != nil {
			return errors.New("获取配置失败")
		}
		apiConfig.Environs = r.envVars

		//设置前置步骤
		if apiConfig.SetupCaseID != nil && *apiConfig.SetupCaseID != 0 {
			//r.tcm.SetupCase = true
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
		testCase.Name = c.ApiCase.Name
		testCase.ID = c.ApiCase.ID
		for _, s := range cases {
			hrpCaseStep, err := getCaseStepHrp(s.ApiCaseStepId)
			if err != nil {
				return err
			}
			if hrpCaseStep != nil {
				testCase.TestSteps = append(testCase.TestSteps, *hrpCaseStep)
			}
			testCase.Confing = *apiConfig
		}
		testCaseList = append(testCaseList, testCase)
	}

	r.d.ID = r.runCaseReq.ApiID
	r.d.RunDebugTalkFile()
	err = yangfanTestCaseToHrpCase(testCaseList, r.d.FilePath, &r.tcm)
	if err != nil {
		return errors.New("用例转换失败")
	}
	r.reportOperation = &ReportOperation{
		report: &interfacecase.ApiReport{
			Name:      reportName,
			CaseType:  r.caseType,
			RunType:   r.runType,
			SetupCase: r.tcm.SetupCase,
			Operator: interfacecase.Operator{
				ProjectID: r.d.ProjectID,
			},
			ApiEnvName: envName,
			ApiEnvID:   task.ApiEnvID,
		},
	}
	r.reportOperation.CreateReport()
	return nil
}

func (r *runTask) RunCase() (err error) {
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

func (r *runTask) Report() (report *interfacecase.ApiReport, err error) {
	if r.reportOperation.report == nil {
		return nil, errors.New("未获取到报告信息")
	}
	return r.reportOperation.report, nil
}
