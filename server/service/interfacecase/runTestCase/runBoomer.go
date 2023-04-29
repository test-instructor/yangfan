package runTestCase

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/hrp"
	"github.com/test-instructor/yangfan/server/hrp/pkg/boomer"
	"github.com/test-instructor/yangfan/server/model/common/request"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	"gorm.io/gorm"
)

type BoomerState int

var (
	BoomerStateRunning BoomerState = 1
	BoomerStateStop    BoomerState = 2
)

type B struct {
	Boom     *hrp.HRPBoomer
	State    BoomerState
	OutputDB *boomer.DbOutput
	r        *RunBoomerStandalone
}

var b *B

func RunYangfanBoomer(r *RunBoomerStandalone, pReport *interfacecase.PerformanceReport, pTask *interfacecase.Performance, runCaseReq request.RunCaseReq) {
	spawnCount := runCaseReq.Operation.SpawnCount
	spawnRate := runCaseReq.Operation.SpawnRate
	profile := &boomer.Profile{}
	profile.SpawnCount = spawnCount
	profile.SpawnRate = spawnRate
	switch runCaseReq.Operation.Running {
	case request.RunningTypeRun:
		if b != nil {
			return
		}
		b = &B{}
		b.r = r
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
		b.Boom = hrp.NewStandaloneBoomer(spawnCount, spawnRate)
		b.OutputDB = boomer.NewDbOutput(*pReport, *pTask)
		//b.OutputDB.OnStart()
		b.Boom.SetProfile(profile)
		b.Boom.InitBoomerYangfan()
		b.Boom.AddOutput(b.OutputDB)
		b.State = BoomerStateRunning
		go b.Boom.Run(r.tcm.Case...)
	case request.RunningTypeRebalance:
		profile = b.Boom.GetProfile()
		profile.SpawnCount = spawnCount
		profile.SpawnRate = spawnRate
		err := b.Boom.ReBalance(profile)
		if err != nil {
			return
		}
	case request.RunningTypeStop:
		if b == nil {
			return
		}
		b.Boom.Quit()
		b.OutputDB.PReport.State = interfacecase.StateStopped
		b.State = BoomerStateStop
		b.r.d.StopDebugTalkFile()
		b = nil
	}
}

func NewBoomer(runCaseReq request.RunCaseReq, runType interfacecase.RunType) TestCase {
	return &RunBoomerStandalone{
		CaseID:     runCaseReq.CaseID,
		caseType:   interfacecase.CaseTypeBoomerDebug,
		runCaseReq: runCaseReq,
		runType:    runType,
	}
}

type RunBoomerStandalone struct {
	reportID   uint
	pTask      interfacecase.Performance
	CaseID     uint
	runCaseReq request.RunCaseReq
	runType    interfacecase.RunType
	caseType   interfacecase.CaseType
	tcm        ApisCaseModel
	d          debugTalkOperation
	envVars    map[string]string
}

func (r *RunBoomerStandalone) LoadCase() (err error) {
	//获取运行配置

	var testCase interfacecase.HrpCase
	var testCaseList []interfacecase.HrpCase
	var apiCase interfacecase.Performance
	var apiCaseCase []interfacecase.PerformanceRelationship

	{
		var testCaseStep interfacecase.Performance
		err := global.GVA_DB.Model(interfacecase.Performance{}).Where("id = ? ", r.runCaseReq.CaseID).First(&testCaseStep).Error
		if err != nil {
			return err
		}
		r.runCaseReq.ConfigID = testCaseStep.RunConfigID
		r.runCaseReq.Env = testCaseStep.ApiEnvID
	}
	//获取运行配置
	apiConfig, err := getConfig(r.runCaseReq.ConfigID)
	if err != nil {
		return errors.New("获取配置失败")
	}
	r.envVars, _, err = GetEnvVar(apiConfig.ProjectID, r.runCaseReq.Env)
	if err != nil {
		return errors.New("获取环境变量失败")
	}
	apiConfig.Environs = r.envVars
	global.GVA_LOG.Debug(fmt.Sprintf("boomer debug 1 apiConfig:%d", apiConfig.ID))

	//设置前置套件
	if apiConfig.SetupCaseID != nil && *apiConfig.SetupCaseID != 0 {
		global.GVA_LOG.Debug(fmt.Sprintf("boomer debug 2 apiConfig.SetupCaseID %d", *apiConfig.SetupCaseID))
		hrpCaseStep, err := getCaseStepHrp(*apiConfig.SetupCaseID)
		if err != nil {
			global.GVA_LOG.Debug(fmt.Sprintf("boomer debug 8 已设置前置套件，%s", err))
		}
		testCase.Confing = *apiConfig
		testCase.TestSteps = append(testCase.TestSteps, hrpCaseStep)
	}

	global.GVA_LOG.Debug(fmt.Sprintf("boomer debug 3 已设置前置套件，%d", apiConfig.ID))
	r.tcm.Config = *apiConfig

	//读取用例信息
	apiCase.ID = r.runCaseReq.CaseID
	err = global.GVA_DB.Model(interfacecase.Performance{}).First(&apiCase).Error
	global.GVA_LOG.Debug(fmt.Sprintf("boomer debug 4 apiCase，%v", apiCase))
	caseDB := global.GVA_DB.Model(interfacecase.PerformanceRelationship{}).
		Preload("ApiCaseStep").
		Preload("ApiCaseStep.TStep", func(db2 *gorm.DB) *gorm.DB {
			return db2.Order("Sort")
		}).
		Preload("ApiCaseStep.TStep.Request").
		Where("performance_id = ?", r.runCaseReq.CaseID).
		Order("Sort")
	caseDB.Find(&apiCaseCase)
	global.GVA_LOG.Debug(fmt.Sprintf("boomer debug 5 获取用例信息，%v", apiCaseCase))
	for _, v := range apiCaseCase {
		hrpCaseStep, err := getCaseStepHrp(v.ApiCaseStepId)
		if err != nil {
			return err
		}
		testCase.TestSteps = append(testCase.TestSteps, hrpCaseStep)
		global.GVA_LOG.Debug(fmt.Sprintf("boomer debug 8 *apiConfig.SetupCaseID:%v", hrpCaseStep))
	}
	testCase.Confing = *apiConfig
	testCaseList = append(testCaseList, testCase)
	r.d.ProjectID = apiConfig.ProjectID
	r.d.ID = r.runCaseReq.ApiID
	r.d.RunDebugTalkFile()
	err = yangfanTestCaseToHrpCase(testCaseList, r.d.FilePath, &r.tcm)
	if err != nil {
		return errors.New("用例转换失败")
	}
	return nil
}

func (r *RunBoomerStandalone) RunCase() (err error) {
	//defer r.d.StopDebugTalkFile()
	var pTask interfacecase.Performance
	var pReport interfacecase.PerformanceReport
	err = global.GVA_DB.Model(&interfacecase.Performance{}).
		Where("id = ?", r.CaseID).First(&pTask).Error
	if err != nil {
		return err
	}
	pReport.Name = pTask.Name
	pReport.PerformanceID = pTask.ID
	pReport.ProjectID = pTask.ProjectID
	pReport.State = 1
	err = global.GVA_DB.Save(&pReport).Error
	if err != nil {
		return err
	}
	r.reportID = pReport.ID
	pTask.PerformanceReportId = pReport.ID
	pTask.State = interfacecase.StateInit
	err = global.GVA_DB.Save(&pTask).Error
	if err != nil {
		return err
	}
	r.pTask = pTask
	defer func() {
		if msg := recover(); msg != nil {
			pReport.State = interfacecase.StateError
			global.GVA_DB.Save(&pReport)
			err = errors.New(fmt.Sprintln(msg))
			global.GVA_LOG.Error(fmt.Sprintln(msg))
			return
		}

	}()
	RunYangfanBoomer(r, &pReport, &pTask, r.runCaseReq)

	//var t *testing.T
	//r.reportOperation.CreateReport()
	//defer recoverHrp(r.reportOperation)
	//defer r.d.StopDebugTalkFile()
	//report, err := hrp.NewRunner(t).
	//	SetHTTPStatOn().
	//	SetFailfast(false).
	//	RunJsons(r.tcm.Case...)
	//r.reportOperation.UpdateReport(&report)

	return nil
}

func (r *RunBoomerStandalone) Report() (report *interfacecase.ApiReport, err error) {
	return
}
