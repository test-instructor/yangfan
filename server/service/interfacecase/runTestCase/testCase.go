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

func getCaseStep(id uint) (apiCaseStep interfacecase.ApiCaseStep) {
	global.GVA_DB.Model(&interfacecase.ApiCaseStep{}).
		Preload("TStep", func(db2 *gorm.DB) *gorm.DB {
			return db2.Order("Sort")
		}).
		Preload("TStep.Request").
		First(&apiCaseStep, "id = ?", id)
	return
}

func getCaseStepHrp(stepId uint) (*interfacecase.HrpCaseStep, error) {
	setupCase := getCaseStep(stepId)

	var hrpTestCase interfacecase.HrpTestCase
	hrpTestCase.ID = setupCase.ID
	hrpTestCase.Name = setupCase.Name

	apiConfig, err := getConfig(setupCase.RunConfigID)
	if err != nil {
		return nil, errors.New("获取配置失败")
	}

	hrpTestCase.Confing = *apiConfig
	hrpTestCase.TestSteps = setupCase.TStep
	hrpCase := &interfacecase.HrpCaseStep{
		ID:       setupCase.ID,
		Name:     setupCase.Name,
		TestCase: hrpTestCase,
	}
	return hrpCase, nil
}

func getConfig(id uint) (config *interfacecase.ApiConfig, err error) {
	apiConfig := interfacecase.ApiConfig{GVA_MODEL: global.GVA_MODEL{ID: id}}
	err = global.GVA_DB.Model(&interfacecase.ApiConfig{}).
		Preload("Project").
		First(&apiConfig).Error
	if err != nil {
		return nil, errors.New("获取配置失败")
	}
	return &apiConfig, nil
}

type TestCase interface {
	LoadCase() (err error)
	RunCase() (err error)
	Report() (reports *interfacecase.ApiReport, err error)
}

func RunTestCase(tc TestCase) (reports *interfacecase.ApiReport, err error) {
	err = tc.LoadCase()
	if err != nil {
		return
	}
	err = tc.RunCase()
	if err != nil {
		return
	}
	report, err := tc.Report()
	return report, nil
}

func NewRunApi(runCaseReq request.RunCaseReq, runType interfacecase.RunType) TestCase {
	return &runAPI{
		ApiID:      runCaseReq.ApiID,
		caseType:   interfacecase.CaseTypeApi,
		runCaseReq: runCaseReq,
		runType:    runType,
	}
}

type runAPI struct {
	reportOperation *ReportOperation
	ApiID           uint
	runCaseReq      request.RunCaseReq
	runType         interfacecase.RunType
	caseType        interfacecase.CaseType
	tcm             ApisCaseModel
	d               debugTalkOperation
}

func (r *runAPI) LoadCase() (err error) {
	var apiStep interfacecase.ApiStep
	//var testCaseList []interfacecase.ApiCaseStep
	//var apiCase interfacecase.ApiCaseStep
	//获取运行配置
	var testCase interfacecase.HrpCase
	var testCaseList []interfacecase.HrpCase

	//获取运行配置
	apiConfig, err := getConfig(r.runCaseReq.ConfigID)
	if err != nil {
		return errors.New("获取配置失败")
	}

	//设置前置套件
	if apiConfig.SetupCaseID != nil {
		hrpCaseStep, err := getCaseStepHrp(*apiConfig.SetupCaseID)
		if err != nil {
			return err
		}
		testCase.TestSteps = append(testCase.TestSteps, *hrpCaseStep)
	}
	r.tcm.Config = *apiConfig

	global.GVA_DB.Model(&interfacecase.ApiStep{}).
		Preload("Request").
		First(&apiStep, "id = ?", r.runCaseReq.CaseID)
	testCase.Name = apiStep.Name
	var hrpTestCase interfacecase.HrpTestCase
	hrpTestCase.Name = apiStep.Name
	hrpTestCase.ID = apiStep.ID
	hrpTestCase.Confing = *apiConfig
	hrpTestCase.TestSteps = append(hrpTestCase.TestSteps, apiStep)
	hrpCase := &interfacecase.HrpCaseStep{
		ID:       hrpTestCase.ID,
		Name:     hrpTestCase.Name,
		TestCase: hrpTestCase,
	}
	testCase.TestSteps = append(testCase.TestSteps, *hrpCase)
	testCaseList = append(testCaseList, testCase)
	r.d.ProjectID = apiConfig.ProjectID
	r.d.ID = r.runCaseReq.ApiID
	r.d.RunDebugTalkFile()
	err = cheetahTestCaseToHrpCase(testCaseList, r.d.FilePath, &r.tcm)
	if err != nil {
		return errors.New("用例转换失败")
	}
	r.reportOperation = &ReportOperation{
		report: &interfacecase.ApiReport{
			Name:      apiStep.Name,
			CaseType:  r.caseType,
			RunType:   r.runType,
			ProjectID: apiStep.ProjectID,
			SetupCase: r.tcm.SetupCase,
		},
	}
	r.reportOperation.CreateReport()
	return nil
}

func (r *runAPI) RunCase() (err error) {
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

func (r *runAPI) Report() (report *interfacecase.ApiReport, err error) {
	if r.reportOperation.report == nil {
		return nil, errors.New("未获取到报告信息")
	}
	return r.reportOperation.report, nil
}

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

	//获取运行配置
	apiConfig, err := getConfig(r.runCaseReq.ConfigID)
	if err != nil {
		return errors.New("获取配置失败")
	}

	//设置前置套件
	if apiConfig.SetupCaseID != nil {
		//前置用例逻辑需要修改
		//r.tcm.SetupCase = true
		hrpCaseStep, err := getCaseStepHrp(*apiConfig.SetupCaseID)
		if err != nil {
			return err
		}
		testCase.TestSteps = append(testCase.TestSteps, *hrpCaseStep)
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

func NewRunCase(runCaseReq request.RunCaseReq, runType interfacecase.RunType) TestCase {
	return &runCase{
		CaseID:     runCaseReq.CaseID,
		caseType:   interfacecase.CaseTypeCases,
		runCaseReq: runCaseReq,
		runType:    runType,
	}
}

type runCase struct {
	reportOperation *ReportOperation
	CaseID          uint
	runCaseReq      request.RunCaseReq
	runType         interfacecase.RunType
	caseType        interfacecase.CaseType
	tcm             ApisCaseModel
	d               debugTalkOperation
}

func (r *runCase) LoadCase() (err error) {
	var testCase interfacecase.HrpCase
	var testCaseList []interfacecase.HrpCase
	var apiCase interfacecase.ApiCase
	var apiCaseCase []interfacecase.ApiCaseRelationship
	//var apiCases interfacecase.ApiCaseStep
	//var tcm *ApisCaseModel

	//获取运行配置
	apiConfig, err := getConfig(r.runCaseReq.ConfigID)
	if err != nil {
		return errors.New("获取配置失败")
	}

	//设置前置套件
	if apiConfig.SetupCaseID != nil {
		hrpCaseStep, err := getCaseStepHrp(*apiConfig.SetupCaseID)
		if err != nil {
			return err
		}
		testCase.TestSteps = append(testCase.TestSteps, *hrpCaseStep)
	}
	r.tcm.Config = *apiConfig

	//读取用例信息
	apiCase.ID = r.runCaseReq.CaseID
	err = global.GVA_DB.Model(interfacecase.ApiCase{}).First(&apiCase).Error
	caseDB := global.GVA_DB.Model(interfacecase.ApiCaseRelationship{}).
		Preload("ApiCaseStep").
		Preload("ApiCaseStep.TStep", func(db2 *gorm.DB) *gorm.DB {
			return db2.Order("Sort")
		}).
		Preload("ApiCaseStep.TStep.Request").
		Where("api_case_id = ?", r.runCaseReq.CaseID).
		Order("Sort")
	caseDB.Find(&apiCaseCase)
	for _, v := range apiCaseCase {
		//testCaseList = append(testCaseList, v.ApiCaseStep)

		hrpCaseStep, err := getCaseStepHrp(v.ApiCaseStepId)
		if err != nil {
			return err
		}
		testCase.TestSteps = append(testCase.TestSteps, *hrpCaseStep)
	}
	testCaseList = append(testCaseList, testCase)
	r.d.ProjectID = apiConfig.ProjectID
	r.d.ID = r.runCaseReq.ApiID
	r.d.RunDebugTalkFile()
	err = cheetahTestCaseToHrpCase(testCaseList, r.d.FilePath, &r.tcm)
	if err != nil {
		return errors.New("用例转换失败")
	}
	r.reportOperation = &ReportOperation{
		report: &interfacecase.ApiReport{
			Name:      apiCase.Name,
			CaseType:  r.caseType,
			RunType:   r.runType,
			ProjectID: apiConfig.ProjectID,
			SetupCase: r.tcm.SetupCase,
		},
	}
	r.reportOperation.CreateReport()
	return nil
}

func (r *runCase) RunCase() (err error) {
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

func (r *runCase) Report() (report *interfacecase.ApiReport, err error) {
	if r.reportOperation.report == nil {
		return nil, errors.New("未获取到报告信息")
	}
	return r.reportOperation.report, nil
}

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
}

func (r *runTask) LoadCase() (err error) {
	var testCaseList []interfacecase.HrpCase
	//var apiCases interfacecase.ApiCaseStep
	//var tcm *ApisCaseModel
	var reportName string
	taskCase := taskSort(r.runCaseReq.TaskID)

	for _, c := range taskCase {
		var testCase interfacecase.HrpCase
		reportName = c.ApiTimerTask.Name
		r.d.ProjectID = c.ApiCase.ProjectID
		r.d.ID = c.ApiTimerTaskId
		cases := caseSort(c.ApiCaseId)
		apiConfig, err := getConfig(r.runCaseReq.ConfigID)
		if err != nil {
			return errors.New("获取配置失败")
		}

		//设置前置套件
		if apiConfig.SetupCaseID != nil {
			//r.tcm.SetupCase = true
			hrpCaseStep, err := getCaseStepHrp(*apiConfig.SetupCaseID)
			if err != nil {
				return err
			}
			testCase.TestSteps = append(testCase.TestSteps, *hrpCaseStep)
		}
		r.tcm.Config = *apiConfig
		testCase.Name = c.ApiCase.Name
		testCase.ID = c.ApiCase.ID
		for _, s := range cases {
			hrpCaseStep, err := getCaseStepHrp(s.ApiCaseStepId)
			if err != nil {
				return err
			}

			testCase.TestSteps = append(testCase.TestSteps, *hrpCaseStep)
		}
		testCaseList = append(testCaseList, testCase)
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

func (r *runTask) RunCase() (err error) {
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

func (r *runTask) Report() (report *interfacecase.ApiReport, err error) {
	if r.reportOperation.report == nil {
		return nil, errors.New("未获取到报告信息")
	}
	return r.reportOperation.report, nil
}
