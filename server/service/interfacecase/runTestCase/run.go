package runTestCase

import (
	"encoding/json"
	"fmt"
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/hrp"
	"github.com/test-instructor/cheetah/server/model/common/request"
	"github.com/test-instructor/cheetah/server/model/interfacecase"
	"gorm.io/gorm"
	"sync"
	"testing"
)

type ToTestCase struct {
	Config    interfacecase.ApiConfig
	TestSteps []interfacecase.ApiStep
}

func getDebugTalkFile(projectID uint) (debugTalkByte []byte, err error) {
	var debugTalkFirst interfacecase.ApiDebugTalk
	db := global.GVA_DB.
		Model(&interfacecase.ApiDebugTalk{}).
		Preload("Project").Joins("Project").Where("Project.ID = ?", projectID)
	//查询对应的类型
	db.Where("file_type = ?", interfacecase.FileDebugTalk).Order("id desc")
	err = db.First(&debugTalkFirst).Error
	if err != nil {
		defaultDB := global.GVA_DB.Model(&interfacecase.ApiDebugTalk{}).
			Preload("Project").Joins("Project").Where("Project.ID = ?", 2)
		defaultDB.Where("file_type = ?", interfacecase.FileDebugTalk)
		err = defaultDB.First(&debugTalkFirst).Error
	}
	return []byte(debugTalkFirst.Content), err
}

func RunTask(timerTaskID uint, castType interfacecase.CaseType, runType interfacecase.RunType) {
	fmt.Println("开始定时任务")
	var err error
	var reports interfacecase.ApiReport
	var report interfacecase.ApiReport
	debugTalkFilePath := CreateDebugTalk(fmt.Sprintf("Task_%d_debugTalk/", timerTaskID))
	if global.DebugTalkLock[debugTalkFilePath] == nil {
		global.DebugTalkLock[debugTalkFilePath] = &sync.Mutex{}
	}
	global.DebugTalkLock[debugTalkFilePath].Lock()
	defer global.DebugTalkLock[debugTalkFilePath].Unlock()
	global.GVA_DB.Model(interfacecase.ApiDebugTalk{}).Where("")

	var t *testing.T

	var timerTask interfacecase.TimerTask
	var testCaseList interfacecase.ApiTestCase

	timerTask.ID = timerTaskID
	err = global.GVA_DB.Model(interfacecase.TimerTask{}).
		Preload("RunConfig").
		Preload("ApiTestCase", func(caseDB *gorm.DB) *gorm.DB {
			return caseDB.Preload("TStep", func(db2 *gorm.DB) *gorm.DB {
				return db2.Order("Sort")
			}).
				Preload("TStep.Request")
		}).
		First(&timerTask).Error

	caseDB := global.GVA_DB.Model(interfacecase.ApiTestCase{})
	caseDB.Find(&testCaseList)

	if err != nil {
		return
	}

	apiConfig := interfacecase.ApiConfig{GVA_MODEL: global.GVA_MODEL{ID: timerTask.RunConfig.ID}}
	global.GVA_DB.Model(&interfacecase.ApiConfig{}).Preload("Project").First(&apiConfig)
	report.Name = timerTask.Name
	report.CaseType = castType
	report.RunType = runType
	report.Project.ID = apiConfig.ProjectID
	global.GVA_DB.Create(&report)

	var apiConfigName = apiConfig.Name
	var l []hrp.ITestCase
	debugTalkByte, _ := getDebugTalkFile(apiConfig.ProjectID)
	hrp.BuildHashicorpPyPlugin(debugTalkByte, debugTalkFilePath)
	defer hrp.RemoveHashicorpPyPlugin(debugTalkFilePath)
	fmt.Println("用例数", len(timerTask.ApiTestCase))
	for _, testCase := range timerTask.ApiTestCase {
		fmt.Println("用例id", testCase.ID)
		apiConfig.Name = apiConfigName + "-" + testCase.Name
		toTestCase := ToTestCase{Config: apiConfig, TestSteps: testCase.TStep}
		caseJson, _ := json.Marshal(toTestCase)
		global.GVA_LOG.Debug("测试用例json格式")
		global.GVA_LOG.Debug("\n" + string(caseJson))
		tc := &hrp.TestCaseJson{
			JsonString:        string(caseJson),
			ID:                testCase.ID,
			DebugTalkFilePath: debugTalkFilePath,
		}
		testCase, _ := tc.ToTestCase()
		l = append(l, testCase)
	}
	reports, _ = hrp.NewRunner(t).
		SetFailfast(false).
		SetHTTPStatOn().
		RunJsons(l...)
	reports.Name = report.Name
	reports.ID = report.ID
	reports.CaseType = castType
	reports.RunType = runType
	reports.CreatedAt = report.CreatedAt
	reports.Project.ID = apiConfig.ProjectID
	reports.Status = 1
	global.GVA_DB.Save(&reports)
}

func RunTimerTask(timerTaskID uint) func() {
	return func() {
		RunTask(timerTaskID, 3, 3)
	}
}

func RunCase(apiCaseID request.RunCaseReq) (reports interfacecase.ApiReport, err error) {
	var t *testing.T

	var report interfacecase.ApiReport
	var apiCases interfacecase.ApiTestCase
	debugTalkFilePath := CreateDebugTalk(fmt.Sprintf("Task_%d_debugTalk/", apiCaseID))
	if global.DebugTalkLock[debugTalkFilePath] == nil {
		global.DebugTalkLock[debugTalkFilePath] = &sync.Mutex{}
	}
	global.DebugTalkLock[debugTalkFilePath].Lock()
	defer global.DebugTalkLock[debugTalkFilePath].Unlock()
	apiConfig := interfacecase.ApiConfig{GVA_MODEL: global.GVA_MODEL{ID: apiCaseID.ConfigID}}
	global.GVA_DB.Model(&interfacecase.ApiConfig{}).Preload("Project").First(&apiConfig)
	global.GVA_DB.Model(&interfacecase.ApiTestCase{}).
		Preload("TStep", func(db2 *gorm.DB) *gorm.DB {
			return db2.Order("Sort")
		}).
		Preload("TStep.Request").
		First(&apiCases, "id = ?", apiCaseID.CaseID)
	report.Name = apiCases.Name
	report.CaseType = interfacecase.CaseTypeCases
	report.RunType = interfacecase.RunTypeRuning
	report.Project.ID = apiConfig.ProjectID
	global.GVA_DB.Create(&report)
	//暂时注释
	//hrp.BuildHashicorpPyPlugin(apiConfig.ProjectID)
	//defer hrp.RemoveHashicorpPyPlugin()
	apiConfig.Name = apiCases.Name + "-" + apiConfig.Name
	toTestCase := ToTestCase{Config: apiConfig, TestSteps: apiCases.TStep}
	caseJson, _ := json.Marshal(toTestCase)

	tc := &hrp.TestCaseJson{
		JsonString:        string(caseJson),
		ID:                apiCaseID.CaseID,
		DebugTalkFilePath: debugTalkFilePath,
	}
	testCase, _ := tc.ToTestCase()
	fmt.Println("------------------------------------")
	fmt.Println(string(caseJson))
	fmt.Println("==========================")
	reports, errs := hrp.NewRunner(t).
		SetHTTPStatOn().
		SetFailfast(false).
		RunJsons(testCase)
	if errs != nil {
		t.Fatalf("run testcase error: %v", err)
	}

	reports.Name = report.Name
	reports.ID = report.ID
	report.CaseType = interfacecase.CaseTypeCases
	report.RunType = interfacecase.RunTypeRuning
	reports.CreatedAt = report.CreatedAt
	reports.Project.ID = apiConfig.ProjectID
	reports.Status = 1
	global.GVA_DB.Save(&reports)
	return
}
