package runTestCase

import (
	"errors"
	"fmt"
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/model/interfacecase"
	"gorm.io/gorm"
)

func getCaseStep(id uint) (apiCaseStep interfacecase.ApiCaseStep) {
	global.GVA_DB.Model(&interfacecase.ApiCaseStep{}).
		Preload("TStep", func(db2 *gorm.DB) *gorm.DB {
			return db2.Order("Sort")
		}).
		Preload("TStep.Request").
		Preload("TStep.Transaction").
		Preload("TStep.Rendezvous").
		First(&apiCaseStep, "id = ?", id)
	return
}

func getCaseStepHrp(stepId uint) (*interfacecase.HrpCaseStep, error) {
	global.GVA_LOG.Debug(fmt.Sprintf("stepId:%v", stepId))
	setupCase := getCaseStep(stepId)
	var hrpTestCase interfacecase.HrpTestCase
	hrpTestCase.ID = setupCase.ID
	hrpTestCase.Name = setupCase.Name

	apiConfig, err := getConfig(setupCase.RunConfigID)
	if err != nil {
		return nil, errors.New("获取配置失败")
	}
	var hrpCase *interfacecase.HrpCaseStep
	hrpTestCase.Confing = *apiConfig
	hrpTestCase.TestSteps = setupCase.TStep
	global.GVA_LOG.Debug("hrpTestCase")
	global.GVA_LOG.Debug(fmt.Sprintf("getCaseStepHrp 1 %v", setupCase))
	for k, _ := range setupCase.TStep {
		if setupCase.TStep[k].Transaction != nil || setupCase.TStep[k].Rendezvous != nil || setupCase.TStep[k].ThinkTime != nil {
			hrpCase = &interfacecase.HrpCaseStep{
				ID:          setupCase.ID,
				Name:        setupCase.Name,
				TestCase:    nil,
				Transaction: setupCase.TStep[k].Transaction,
				Rendezvous:  setupCase.TStep[k].Rendezvous,
				ThinkTime:   setupCase.TStep[k].ThinkTime,
			}

		} else {
			hrpCase = &interfacecase.HrpCaseStep{
				ID:       setupCase.ID,
				Name:     setupCase.Name,
				TestCase: hrpTestCase,
			}
		}
	}

	global.GVA_LOG.Debug(fmt.Sprintf("hrpCase: %v", hrpCase))
	return hrpCase, nil
}

func getConfig(id uint) (config *interfacecase.ApiConfig, err error) {
	global.GVA_LOG.Debug(fmt.Sprintf("获取配置id：%d", id))
	global.GVA_LOG.Debug(fmt.Sprintln(id))
	apiConfig := interfacecase.ApiConfig{GVA_MODEL: global.GVA_MODEL{ID: id}}
	err = global.GVA_DB.Model(&interfacecase.ApiConfig{}).
		Preload("Project").
		Preload("SetupCase").
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
