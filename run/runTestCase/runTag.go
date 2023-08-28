package runTestCase

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/test-instructor/yangfan/proto/run"
	"go.uber.org/zap"

	"github.com/test-instructor/yangfan/hrp"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/common/request"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
)

func NewRunTag(runCaseReq request.RunCaseReq, runType interfacecase.RunType, msg *run.Msg) TestCase {
	caseType := interfacecase.CaseTypeTag
	if runCaseReq.ReportCIID > 0 {
		runType = interfacecase.RunTypeCI
	}
	return &runTag{
		CaseID:     runCaseReq.CaseID,
		caseType:   caseType,
		runCaseReq: runCaseReq,
		runType:    runType,
		msg:        msg,
		ReportCIID: runCaseReq.ReportCIID,
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
	envVars         map[string]string
	msg             *run.Msg
	ReportCIID      uint
}

func (r *runTag) LoadCase() (err error) {

	var tag interfacecase.ApiTimerTaskTag
	var envName string
	db := global.GVA_DB.Model(interfacecase.ApiTimerTaskTag{})
	db.Preload("ApiTimerTask")
	db.First(&tag, "id = ?", r.runCaseReq.TagID)

	var testCaseList []interfacecase.HrpCase
	var reportName = tag.Name

	r.envVars, envName, err = GetEnvVar(tag.ProjectID, r.runCaseReq.Env)
	if err != nil {
		return errors.New("获取环境变量失败")
	}
	for _, v := range tag.ApiTimerTask {

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
			testCase.Name = fmt.Sprintln("【任务：", c.ApiTimerTask.Name, "】-", c.ApiCase.Name)
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

	}

	r.d.ID = r.runCaseReq.ApiID
	r.d.RunDebugTalkFile()
	err = yangfanTestCaseToHrpCase(testCaseList, r.d.FilePath, &r.tcm)
	if err != nil {
		return errors.New("用例转换失败")
	}
	hostname, _ := os.Hostname()
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
			ApiEnvID:   r.runCaseReq.Env,
			Hostname:   hostname,
		},
		msg: r.msg,
	}
	r.reportOperation.CreateReport()
	return nil
}

func (r *runTag) setCIReport() {
	if r.runType != interfacecase.RunTypeCI {
		return
	}
	var ci interfacecase.ApiReportCI
	err := global.GVA_DB.Model(interfacecase.ApiReportCI{}).First(&ci, "id = ?", r.ReportCIID).Error
	if err != nil {
		return
	}
	ci.ReportID = r.reportOperation.report.ID
	err = global.GVA_DB.Model(interfacecase.ApiReportCI{}).Where("id = ?", ci.ID).Save(&ci).Error
	if err != nil {
		global.GVA_LOG.Error("CI报告更新失败", zap.Error(err))
	}
	if ci.CallbackUrl != "" {
		err := global.GVA_DB.Preload("Report").First(&ci, "id = ? ", ci.ID).Error
		if err != nil {
			global.GVA_LOG.Error("CI报告获取失败", zap.Error(err))
			return
		}
		resp := sendReport(ci.Report, ci.CallbackUrl)
		global.GVA_DB.Model(&ci).UpdateColumns(interfacecase.ApiReportCI{CallbackResponse: resp})
	}
}

func sendReport(body interface{}, url string) string {
	reqJSON, err := json.Marshal(body)
	if err != nil {
		global.GVA_LOG.Error("Error marshaling JSON:", zap.Error(err))
	}

	bodyByte := bytes.NewBuffer(reqJSON)
	req, err := http.NewRequest(http.MethodPost, url, bodyByte)
	if err != nil {
		global.GVA_LOG.Error("Error creating request:", zap.Error(err))
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		global.GVA_LOG.Error("Error sending POST request:", zap.Error(err))
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	var responseBody map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		global.GVA_LOG.Error("Error decoding JSON response body:", zap.Error(err))
		return ""
	}
	jsonResponse, err := json.MarshalIndent(responseBody, "", "    ")
	if err != nil {
		global.GVA_LOG.Error("Error encoding JSON response body:", zap.Error(err))
		return ""
	}
	global.GVA_LOG.Debug(string(jsonResponse))
	return string(jsonResponse)
}

func (r *runTag) RunCase() (err error) {
	var t *testing.T
	defer recoverHrp(r.reportOperation)
	defer r.d.StopDebugTalkFile()
	reportHRP, err := hrp.NewRunner(t).
		SetHTTPStatOn().
		SetFailfast(false).
		RunJsons(r.tcm.Case...)
	var report interfacecase.ApiReport
	err = json.Unmarshal(reportHRP, &report)
	if err != nil {
		global.GVA_LOG.Error("报告解析失败", zap.Error(err))
	}
	r.reportOperation.UpdateReport(&report)
	if err != nil {
		return err
	}
	r.setCIReport()
	return nil
}

func (r *runTag) Report() (reports *interfacecase.ApiReport, err error) {
	if r.reportOperation.report == nil {
		return nil, errors.New("未获取到报告信息")
	}
	return r.reportOperation.report, nil
}
