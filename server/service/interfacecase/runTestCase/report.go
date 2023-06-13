package runTestCase

import (
	"encoding/json"

	"github.com/test-instructor/yangfan/hrp"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
)

type ReportOperation struct {
	report *interfacecase.ApiReport
}

func (r *ReportOperation) CreateReport() {
	global.GVA_DB.Create(&r.report)
}

type Stat struct {
	Total     int `json:"total"`
	Failures  int `json:"failures"`
	Successes int `json:"successes"`
}

func resetReport(reports *interfacecase.ApiReport) {

	//修正测试报告
	var statTestcases interfacecase.ApiReportStatTestcases
	var statTeststeps interfacecase.ApiReportStatTeststeps
	testcaseStatus := true
	for k, v := range reports.Details {
		var statStep hrp.TestStepStat
		stepStatus := true
		for _, v2 := range v.Records {
			apiSuccess := 0
			apiFail := 0
			for _, v2 := range v2.Data {
				if v2.Success == true {
					apiSuccess++
				} else {
					apiFail++
					stepStatus = false
				}
			}
			statStep.Successes += apiSuccess
			statStep.Failures += apiFail
			statStep.Total = apiSuccess + apiFail
		}
		statString, _ := json.Marshal(statStep)
		reports.Details[k].Success = stepStatus
		reports.Details[k].Stat = statString
		if stepStatus == true {
			statTestcases.Success++
		} else {
			statTestcases.Fail++
			testcaseStatus = false
		}
		statTeststeps.Successes += statStep.Successes
		statTeststeps.Failures += statStep.Failures
		statTeststeps.Total = statTeststeps.Successes + statTeststeps.Failures
		statTestcases.Total = statTestcases.Success + statTestcases.Fail
	}
	*reports.Success = testcaseStatus
	reports.Stat.TestSteps = &statTeststeps
	reports.Stat.TestCases = &statTestcases
}

func (r *ReportOperation) UpdateReport(reports *interfacecase.ApiReport) {
	//修正测试报告，hrp的测试报告数据不兼容
	resetReport(reports)
	reports.Name = r.report.Name
	reports.ID = r.report.ID
	reports.CaseType = r.report.CaseType
	reports.RunType = r.report.RunType
	reports.CreatedAt = r.report.CreatedAt
	reports.Project.ID = r.report.ProjectID
	reports.Status = 1
	reports.SetupCase = r.report.SetupCase
	reports.ApiEnvName = r.report.ApiEnvName
	reports.ApiEnvID = r.report.ApiEnvID
	reports.Hostname = r.report.Hostname
	for i, v := range reports.Details {
		if v.Name == "" {
			testCase := interfacecase.ApiCase{
				GVA_MODEL: global.GVA_MODEL{ID: v.CaseID},
			}
			global.GVA_DB.Model(&interfacecase.ApiCase{}).First(&testCase)
			reports.Details[i].Name = testCase.Name
		}
	}
	global.GVA_DB.Save(&reports)
}

func (r *ReportOperation) Recover(msg string) {
	r.report.Status = 2
	r.report.Describe = msg
	global.GVA_DB.Save(&r.report)
}
