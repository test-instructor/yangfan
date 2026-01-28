package runTestCase

import (
	"errors"
	"strconv"

	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/automation"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	"gorm.io/gorm"
)

func getAutoCaseStepHrp(stepId uint) (*HrpCaseStep, error) {
	var autoCaseStep automation.AutoCaseStep
	err := global.GVA_DB.Model(&automation.AutoCaseStep{}).
		First(&autoCaseStep, "id = ?", stepId).Error
	if err != nil {
		return nil, err
	}

	var relations []automation.AutoCaseStepRelation
	err = global.GVA_DB.Model(&automation.AutoCaseStepRelation{}).
		Preload("AutoStep.Request").
		Where("auto_case_step_id = ?", stepId).
		Order("sort").
		Find(&relations).Error
	if err != nil {
		return nil, err
	}

	var steps []*automation.AutoStep
	for _, r := range relations {
		s := r.AutoStep
		if s.Request == nil && s.RequestID != 0 {
			var req automation.Request
			global.GVA_DB.Model(&automation.Request{}).First(&req, s.RequestID)
			s.Request = &req
		}
		step := s
		steps = append(steps, &step)
	}

	runConfig, err := getRunConfig(uint(autoCaseStep.ConfigID))
	if err != nil {
		return nil, errors.New("获取配置失败")
	}

	var hrpTestCase HrpTestCase
	hrpTestCase.ID = autoCaseStep.ID
	hrpTestCase.Name = autoCaseStep.StepName
	hrpTestCase.Confing = *runConfig
	hrpTestCase.TestSteps = steps

	var hrpCase *HrpCaseStep

	hrpCase = &HrpCaseStep{
		ID:       autoCaseStep.ID,
		Name:     autoCaseStep.StepName,
		TestCase: hrpTestCase,
	}

	return hrpCase, nil
}

func getRunConfig(id uint) (config *platform.RunConfig, err error) {
	apiConfig := platform.RunConfig{GVA_MODEL: global.GVA_MODEL{ID: id}}
	err = global.GVA_DB.Model(&platform.RunConfig{}).
		Preload("AndroidDeviceOptions").
		Preload("IOSDeviceOptions").
		Preload("HarmonyDeviceOptions").
		Preload("BrowserDeviceOptions").
		First(&apiConfig).Error
	if err != nil {
		return nil, errors.New("获取配置失败")
	}
	return &apiConfig, nil
}

func GetEnvVar(projectID int64, envID int64) (envVars map[string]string, envName string, err error) {
	var env platform.Env
	err = global.GVA_DB.Model(&platform.Env{}).Where("id = ? ", envID).First(&env).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
		return
	}
	if err != nil {
		return
	}
	envName = env.Name
	var envDetail []platform.EnvDetail
	envVars = make(map[string]string)
	err = global.GVA_DB.Model(&platform.EnvDetail{}).
		Where("project_id = ? ", projectID).
		Find(&envDetail).Error
	if err != nil {
		return
	}
	for _, detail := range envDetail {
		// detail.Value is already a map[string]interface{} (datatypes.JSONMap)
		key := strconv.Itoa(int(envID))
		if val, ok := detail.Value[key]; ok {
			if strVal, ok := val.(string); ok {
				envVars[detail.Key] = strVal
			}
		}
	}
	return
}
