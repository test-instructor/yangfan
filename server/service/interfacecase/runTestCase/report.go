package runTestCase

import (
	"encoding/json"
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/model/interfacecase"
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
	successNum := 0
	failNum := 0

	for k, v := range reports.Details {
		//stat := v.Stat.(map[string]uint)
		var stat Stat
		json.Unmarshal(v.Stat, &stat)
		if stat.Total != stat.Successes {
			//v.Success = false
			reports.Details[k].Success = false
			failNum++
		} else {
			successNum++
		}
	}
	success := failNum == 0
	reports.Success = &success
	reports.Stat.TestCases.Fail = failNum
	reports.Stat.TestCases.Success = successNum
}

func (r *ReportOperation) UpdateReport(reports *interfacecase.ApiReport) {
	resetReport(reports)
	reports.Name = r.report.Name
	reports.ID = r.report.ID
	reports.CaseType = r.report.CaseType
	reports.RunType = r.report.RunType
	reports.CreatedAt = r.report.CreatedAt
	reports.Project.ID = r.report.ProjectID
	reports.Status = 1
	reports.SetupCase = r.report.SetupCase
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
