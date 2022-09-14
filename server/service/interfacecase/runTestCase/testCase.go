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

//type Report interface {
//	Create()
//	Update(reports *interfacecase.ApiReport)
//}

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
	reportOperation ReportOperation
	ApiID           uint
	runCaseReq      request.RunCaseReq
	runType         interfacecase.RunType
	caseType        interfacecase.CaseType
	tcm             ApisCaseModel
	d               debugTalkOperation
}

func (r *runAPI) LoadCase() (err error) {
	var apiStep interfacecase.ApiStep
	var testCaseList []interfacecase.ApiCaseStep
	var apiCase interfacecase.ApiCaseStep
	//获取运行配置
	apiConfig := interfacecase.ApiConfig{GVA_MODEL: global.GVA_MODEL{ID: r.runCaseReq.ConfigID}}
	err = global.GVA_DB.Model(&interfacecase.ApiConfig{}).
		Preload("Project").
		Preload("SetupCase").
		Preload("SetupCase.TStep.Request").
		First(&apiConfig).Error
	if err != nil {
		return errors.New("获取配置失败")
	}

	//设置前置套件
	if apiConfig.SetupCase != nil {
		r.tcm.SetupCase = true
		testCaseList = append(testCaseList, *apiConfig.SetupCase)
	}
	global.GVA_DB.Model(&interfacecase.ApiStep{}).
		Preload("Request").
		First(&apiStep, "id = ?", r.runCaseReq.CaseID)
	apiCase.Name = apiStep.Name
	apiCase.ProjectID = apiStep.ProjectID
	apiCase.TStep = append(apiCase.TStep, apiStep)
	testCaseList = append(testCaseList, apiCase)
	r.d.ProjectID = apiConfig.ProjectID
	r.d.ID = r.runCaseReq.ApiID
	r.d.RunDebugTalkFile()
	err = cheetahCaseToHrpCase(apiConfig, testCaseList, r.d.FilePath, &r.tcm)
	if err != nil {
		return errors.New("用例转换失败")
	}
	r.reportOperation = ReportOperation{
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
	report, err := hrp.NewRunner(t).
		SetHTTPStatOn().
		SetFailfast(false).
		RunJsons(r.tcm.Case...)
	r.reportOperation.UpdateReport(&report)
	r.d.StopDebugTalkFile()
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
	var testCaseList []interfacecase.ApiCaseStep
	var apiCases interfacecase.ApiCaseStep
	//var apiCases interfacecase.ApiCaseStep
	//var tcm *ApisCaseModel

	//获取运行配置
	apiConfig := interfacecase.ApiConfig{GVA_MODEL: global.GVA_MODEL{ID: r.runCaseReq.ConfigID}}
	err = global.GVA_DB.Model(&interfacecase.ApiConfig{}).
		Preload("Project").
		Preload("SetupCase").
		Preload("SetupCase.TStep.Request").
		First(&apiConfig).Error
	if err != nil {
		return errors.New("获取配置失败")
	}

	//设置前置套件
	if apiConfig.SetupCase != nil {
		r.tcm.SetupCase = true
		testCaseList = append(testCaseList, *apiConfig.SetupCase)
	}
	r.tcm.Config = apiConfig

	//读取用例信息
	global.GVA_DB.Model(&interfacecase.ApiCaseStep{}).
		Preload("TStep", func(db2 *gorm.DB) *gorm.DB {
			return db2.Order("Sort")
		}).
		Preload("TStep.Request").
		First(&apiCases, "id = ?", r.runCaseReq.CaseID)
	testCaseList = append(testCaseList, apiCases)

	r.d.ProjectID = apiConfig.ProjectID
	r.d.ID = r.runCaseReq.ApiID
	r.d.RunDebugTalkFile()
	err = cheetahCaseToHrpCase(apiConfig, testCaseList, r.d.FilePath, &r.tcm)
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
	report, err := hrp.NewRunner(t).
		SetHTTPStatOn().
		SetFailfast(false).
		RunJsons(r.tcm.Case...)
	r.reportOperation.UpdateReport(&report)
	r.d.StopDebugTalkFile()
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
	var testCaseList []interfacecase.ApiCaseStep
	var apiCase interfacecase.ApiCase
	var apiCaseCase []interfacecase.ApiCaseRelationship
	//var apiCases interfacecase.ApiCaseStep
	//var tcm *ApisCaseModel

	//获取运行配置
	apiConfig := interfacecase.ApiConfig{GVA_MODEL: global.GVA_MODEL{ID: r.runCaseReq.ConfigID}}
	err = global.GVA_DB.Model(&interfacecase.ApiConfig{}).
		Preload("Project").
		Preload("SetupCase").
		Preload("SetupCase.TStep.Request").
		First(&apiConfig).Error
	if err != nil {
		return errors.New("获取配置失败")
	}

	//设置前置套件
	if apiConfig.SetupCase != nil {
		r.tcm.SetupCase = true
		testCaseList = append(testCaseList, *apiConfig.SetupCase)
	}
	r.tcm.Config = apiConfig

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
		testCaseList = append(testCaseList, v.ApiCaseStep)
	}

	r.d.ProjectID = apiConfig.ProjectID
	r.d.ID = r.runCaseReq.ApiID
	r.d.RunDebugTalkFile()
	err = cheetahCaseToHrpCase(apiConfig, testCaseList, r.d.FilePath, &r.tcm)
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
	report, err := hrp.NewRunner(t).
		SetHTTPStatOn().
		SetFailfast(false).
		RunJsons(r.tcm.Case...)
	r.reportOperation.UpdateReport(&report)
	r.d.StopDebugTalkFile()
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
	var testCaseList []ApiCaseStep
	//var apiCases interfacecase.ApiCaseStep
	//var tcm *ApisCaseModel
	var reportName string
	taskCase := taskSort(r.runCaseReq.TaskID)

	for _, c := range taskCase {
		reportName = c.ApiTimerTask.Name
		r.d.ProjectID = c.ApiCase.ProjectID
		r.d.ID = c.ApiTimerTaskId
		cases := caseSort(c.ApiCaseId)
		apiConfig := interfacecase.ApiConfig{GVA_MODEL: global.GVA_MODEL{ID: r.runCaseReq.ConfigID}}
		err = global.GVA_DB.Model(&interfacecase.ApiConfig{}).
			Preload("Project").
			Preload("SetupCase").
			Preload("SetupCase.TStep", func(db2 *gorm.DB) *gorm.DB {
				return db2.Order("Sort")
			}).
			Preload("SetupCase.TStep.Request").
			First(&apiConfig).Error
		if err != nil {
			return errors.New("获取配置失败")
		}

		//设置前置套件
		if apiConfig.SetupCase != nil {
			//r.tcm.SetupCase = true
			var apiCaseStep ApiCaseStep
			apiCaseStep.ID = *apiConfig.SetupCaseID
			apiCaseStep.Name = apiConfig.SetupCase.Name
			apiCaseStep.TStep = apiConfig.SetupCase.TStep
			apiCaseStep.ApiCase = apiConfig.SetupCase.ApiCase
			apiCaseStep.ProjectID = apiConfig.SetupCase.ProjectID
			apiCaseStep.Config = &apiConfig
			testCaseList = append(testCaseList, apiCaseStep)
		}
		r.tcm.Config = apiConfig

		for _, s := range cases {
			//for j, _ := range s.ApiCaseStep.TStep {
			//	cases[i].ApiCaseStep.TStep[j].Name = cases[i].ApiCaseStep.Name + " - " + cases[i].ApiCaseStep.TStep[j].Name
			//}
			var apiCaseStep ApiCaseStep
			apiCaseStep.ID = s.ApiCaseStep.ID
			apiCaseStep.Name = s.ApiCase.Name + " - " + s.ApiCaseStep.Name
			apiCaseStep.TStep = s.ApiCaseStep.TStep
			apiCaseStep.ApiCase = s.ApiCaseStep.ApiCase
			apiCaseStep.ProjectID = s.ApiCaseStep.ProjectID
			apiCaseStep.Config = &apiConfig
			testCaseList = append(testCaseList, apiCaseStep)
		}
	}

	r.d.ID = r.runCaseReq.ApiID
	r.d.RunDebugTalkFile()
	err = cheetahTaskToHrpCase(testCaseList, r.d.FilePath, &r.tcm)
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
	report, err := hrp.NewRunner(t).
		SetHTTPStatOn().
		SetFailfast(false).
		RunJsons(r.tcm.Case...)
	r.reportOperation.UpdateReport(&report)
	r.d.StopDebugTalkFile()
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
