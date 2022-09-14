package runTestCase

import (
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/model/interfacecase"
)

type ReportOperation struct {
	report *interfacecase.ApiReport
}

func (r *ReportOperation) CreateReport() {
	global.GVA_DB.Create(&r.report)
}

func (r *ReportOperation) UpdateReport(reports *interfacecase.ApiReport) {
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
