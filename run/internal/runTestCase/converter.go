package runTestCase

import (
	"encoding/json"

	"github.com/test-instructor/yangfan/httprunner/hrp"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/automation"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	"go.uber.org/zap"
)

// convertConfigToTConfig 将业务配置转换为 httprunner TConfig
func convertConfigToTConfig(config *platform.RunConfig, environs map[string]string, projectID int64, envID int) *hrp.TConfig {
	if config == nil {
		return hrp.NewConfig("default")
	}

	tConfig := &hrp.TConfig{
		Name:     config.Name,
		BaseURL:  config.BaseUrl,
		Verify:   config.Verify,
		Environs: environs,
	}

	// 转换 Variables（运行配置中显式配置的变量）
	if config.Variables != nil {
		tConfig.Variables = make(map[string]interface{})
		for k, v := range config.Variables {
			tConfig.Variables[k] = v
		}
	}

	// 将环境变量也注入到 TConfig.Variables 中，便于在 base_url 等字段中通过
	// $ENV_xxx / ${ENV_xxx} 这样的占位符直接引用。
	//
	// 例如：
	//   RunConfig.BaseUrl = "${ENV_domain}"
	//   EnvDetail.Key = "ENV_domain"
	//   -> 这里会把 ENV_domain 写入 Variables，避免 "variable ENV_domain not found" 错误。
	if environs != nil {
		if tConfig.Variables == nil {
			tConfig.Variables = make(map[string]interface{})
		}
		for k, v := range environs {
			// 显式配置的 variables 优先级高于环境变量
			if _, exists := tConfig.Variables[k]; !exists {
				tConfig.Variables[k] = v
			}
		}
	}

	// 转换 Headers
	if config.Headers != nil {
		tConfig.Headers = make(map[string]string)
		for k, v := range config.Headers {
			if strVal, ok := v.(string); ok {
				tConfig.Headers[k] = strVal
			}
		}
	}

	// 转换 Parameters
	if config.Parameters != nil {
		tConfig.Parameters = make(map[string]interface{})
		for k, v := range config.Parameters {
			tConfig.Parameters[k] = v
		}
	}

	// 处理 DataWarehouse: 查询数据仓库并转换为 Parameters
	if config.DataWarehouse != nil && len(config.DataWarehouse) > 0 {
		data, err := queryDataWarehouse(config.DataWarehouse, projectID, envID)
		if err != nil {
			global.GVA_LOG.Warn("Failed to query data warehouse for config", zap.Error(err))
		} else if data != nil && len(data) > 0 {
			// 将数据仓库数据转换为Parameters
			if dwParams, ok := convertDataWarehouseToParameters(data); ok {
				if tConfig.Parameters == nil {
					tConfig.Parameters = make(map[string]interface{})
				}
				// 合并数据仓库参数到现有参数中
				for k, v := range dwParams {
					tConfig.Parameters[k] = v
				}
				global.GVA_LOG.Info("Data warehouse parameters applied to config")
			}
		}
	}

	// 设置超时
	if config.Timeout > 0 {
		tConfig.RequestTimeout = config.Timeout
	}

	// 设置重试次数
	if config.Retry > 0 {
		tConfig.ConfigRetry = config.Retry
	}

	return tConfig
}

// convertRequestToHrpRequest 将业务 Request 转换为 httprunner Request
func convertRequestToHrpRequest(req *automation.Request) *hrp.Request {
	if req == nil {
		return nil
	}

	hrpReq := &hrp.Request{
		Method:         hrp.HTTPMethod(req.Method),
		URL:            req.URL,
		HTTP2:          req.HTTP2,
		Timeout:        req.Timeout,
		AllowRedirects: req.AllowRedirects,
		Verify:         req.Verify,
	}

	// 转换 Params
	if req.Params != nil {
		hrpReq.Params = make(map[string]interface{})
		for k, v := range req.Params {
			hrpReq.Params[k] = v
		}
	}

	// 转换 Headers
	if req.Headers != nil {
		hrpReq.Headers = make(map[string]string)
		for k, v := range req.Headers {
			if strVal, ok := v.(string); ok {
				hrpReq.Headers[k] = strVal
			}
		}
	}

	// 转换 Body
	// 约定：
	// - Json 映射到 hrp.Request.Body，并设置 Content-Type header
	// - Data 映射到 hrp.Request.Body（与 convertCompatRequestBody 保持一致）
	if req.Json != nil && len(req.Json) > 0 {
		jsonBody := make(map[string]interface{})
		for k, v := range req.Json {
			jsonBody[k] = v
		}
		// Json 直接作为请求体
		hrpReq.Body = jsonBody
		// 确保设置 Content-Type header（与 convertCompatRequestBody 逻辑一致）
		if hrpReq.Headers == nil {
			hrpReq.Headers = make(map[string]string)
		}
		if _, ok := hrpReq.Headers["Content-Type"]; !ok {
			hrpReq.Headers["Content-Type"] = "application/json; charset=utf-8"
		}
	} else if req.Data != nil && len(req.Data) > 0 {
		// Data 映射到 Body（与 convertCompatRequestBody 保持一致）
		dataBody := make(map[string]interface{})
		for k, v := range req.Data {
			dataBody[k] = v
		}
		hrpReq.Body = dataBody
	}

	// 转换 Upload
	if req.Upload != nil && len(req.Upload) > 0 {
		hrpReq.Upload = make(map[string]interface{})
		for k, v := range req.Upload {
			hrpReq.Upload[k] = v
		}
	}

	// 转换 DataWarehouse
	if req.DataWarehouse != nil && len(req.DataWarehouse) > 0 {
		hrpReq.DataWarehouse = make(map[string]interface{})
		for k, v := range req.DataWarehouse {
			hrpReq.DataWarehouse[k] = v
		}
	}

	return hrpReq
}

// convertStepConfigToHrpStepConfig 将业务 StepConfig 转换为 httprunner StepConfig
func convertStepConfigToHrpStepConfig(stepConfig *automation.StepConfig, projectID int64, envID int) hrp.StepConfig {
	hrpStepConfig := hrp.StepConfig{
		StepName: stepConfig.StepName,
		Loops:    int(stepConfig.Loops),
		Retry:    stepConfig.Retry,
	}

	// 转换 Variables
	hrpStepConfig.Variables = make(map[string]interface{})
	if stepConfig.Variables != nil {
		for k, v := range stepConfig.Variables {
			hrpStepConfig.Variables[k] = v
		}
	}

	// 转换 Parameters
	if stepConfig.Parameters != nil {
		hrpStepConfig.Parameters = make(map[string]interface{})
		for k, v := range stepConfig.Parameters {
			hrpStepConfig.Parameters[k] = v
		}
	}

	// 处理 DataWarehouse: 查询数据仓库并转换为 Parameters
	if stepConfig.DataWarehouse != nil && len(stepConfig.DataWarehouse) > 0 {
		data, err := queryDataWarehouse(stepConfig.DataWarehouse, projectID, envID)
		if err != nil {
			global.GVA_LOG.Warn("Failed to query data warehouse for step", zap.Error(err))
		} else if data != nil && len(data) > 0 {
			// 将数据仓库数据转换为Parameters
			if dwParams, ok := convertDataWarehouseToParameters(data); ok {
				if hrpStepConfig.Parameters == nil {
					hrpStepConfig.Parameters = make(map[string]interface{})
				}
				// 合并数据仓库参数到现有参数中
				for k, v := range dwParams {
					hrpStepConfig.Parameters[k] = v
				}
				global.GVA_LOG.Info("Data warehouse parameters applied to step config")
			}
		}
	}

	// 转换 Extract
	if stepConfig.Extract != nil {
		hrpStepConfig.Extract = make(map[string]string)
		for k, v := range stepConfig.Extract {
			if strVal, ok := v.(string); ok {
				hrpStepConfig.Extract[k] = strVal
			}
		}
	}

	// 转换 SetupHooks
	if stepConfig.SetupHooks != nil {
		var hooks []string
		if err := json.Unmarshal(stepConfig.SetupHooks, &hooks); err == nil {
			hrpStepConfig.SetupHooks = hooks
		}
	}

	// 转换 TeardownHooks
	if stepConfig.TeardownHooks != nil {
		var hooks []string
		if err := json.Unmarshal(stepConfig.TeardownHooks, &hooks); err == nil {
			hrpStepConfig.TeardownHooks = hooks
		}
	}

	// 转换 Validators
	if stepConfig.Validators != nil {
		var validators []interface{}
		if err := json.Unmarshal(stepConfig.Validators, &validators); err == nil {
			hrpStepConfig.Validators = validators
		}
	}

	// 转换 Export
	if stepConfig.StepExport != nil {
		var exports []string
		if err := json.Unmarshal(stepConfig.StepExport, &exports); err == nil {
			hrpStepConfig.StepExport = exports
		}
	}

	return hrpStepConfig
}

// convertCaseStepConfigToHrpStepConfig 将 CaseStepConfig 转换为 httprunner StepConfig
func convertCaseStepConfigToHrpStepConfig(stepConfig *automation.StepConfig, projectID int64, envID int) hrp.StepConfig {
	hrpStepConfig := hrp.StepConfig{
		StepName: stepConfig.StepName,
		Loops:    int(stepConfig.Loops),
		Retry:    stepConfig.Retry,
	}

	// 转换 Variables
	hrpStepConfig.Variables = make(map[string]interface{})
	if stepConfig.Variables != nil {
		for k, v := range stepConfig.Variables {
			hrpStepConfig.Variables[k] = v
		}
	}

	// 转换 Parameters
	if stepConfig.Parameters != nil {
		hrpStepConfig.Parameters = make(map[string]interface{})
		for k, v := range stepConfig.Parameters {
			hrpStepConfig.Parameters[k] = v
		}
	}

	// 处理 DataWarehouse: 查询数据仓库并转换为 Parameters
	if stepConfig.DataWarehouse != nil && len(stepConfig.DataWarehouse) > 0 {
		data, err := queryDataWarehouse(stepConfig.DataWarehouse, projectID, envID)
		if err != nil {
			global.GVA_LOG.Warn("Failed to query data warehouse for case step", zap.Error(err))
		} else if data != nil && len(data) > 0 {
			// 将数据仓库数据转换为Parameters
			if dwParams, ok := convertDataWarehouseToParameters(data); ok {
				if hrpStepConfig.Parameters == nil {
					hrpStepConfig.Parameters = make(map[string]interface{})
				}
				// 合并数据仓库参数到现有参数中
				for k, v := range dwParams {
					hrpStepConfig.Parameters[k] = v
				}
				global.GVA_LOG.Info("Data warehouse parameters applied to case step config")
			}
		}
	}

	// 转换 Extract
	if stepConfig.Extract != nil {
		hrpStepConfig.Extract = make(map[string]string)
		for k, v := range stepConfig.Extract {
			if strVal, ok := v.(string); ok {
				hrpStepConfig.Extract[k] = strVal
			}
		}
	}

	// 转换 SetupHooks
	if stepConfig.SetupHooks != nil {
		var hooks []string
		if err := json.Unmarshal(stepConfig.SetupHooks, &hooks); err == nil {
			hrpStepConfig.SetupHooks = hooks
		}
	}

	// 转换 TeardownHooks
	if stepConfig.TeardownHooks != nil {
		var hooks []string
		if err := json.Unmarshal(stepConfig.TeardownHooks, &hooks); err == nil {
			hrpStepConfig.TeardownHooks = hooks
		}
	}

	// 转换 Validators
	if stepConfig.Validators != nil {
		var validators []interface{}
		if err := json.Unmarshal(stepConfig.Validators, &validators); err == nil {
			hrpStepConfig.Validators = validators
		}
	}

	// 转换 Export
	if stepConfig.StepExport != nil {
		var exports []string
		if err := json.Unmarshal(stepConfig.StepExport, &exports); err == nil {
			hrpStepConfig.StepExport = exports
		}
	}

	return hrpStepConfig
}

// convertAutoStepToIStep 将单个 AutoStep 转换为 httprunner IStep
func convertAutoStepToIStep(autoStep *automation.AutoStep, projectID int64, envID int) hrp.IStep {
	if autoStep == nil {
		return nil
	}

	stepConfig := convertStepConfigToHrpStepConfig(&autoStep.StepConfig, projectID, envID)
	hrpRequest := convertRequestToHrpRequest(autoStep.Request)

	if hrpRequest == nil {
		global.GVA_LOG.Warn("AutoStep has no request", zap.Uint("step_id", autoStep.ID))
		return nil
	}

	stepRequest := &hrp.StepRequest{
		StepConfig: stepConfig,
		Request:    hrpRequest,
	}

	step := &hrp.StepRequestWithOptionalArgs{
		StepRequest: stepRequest,
	}

	// 处理上传
	if hrpRequest.Upload != nil && len(hrpRequest.Upload) > 0 {
		if step.Request.Headers == nil {
			step.Request.Headers = make(map[string]string)
		}
		step.Request.Headers["Content-Type"] = "${multipart_content_type($m_encoder)}"
		step.Request.Body = "$m_encoder"
	}

	return step
}

// convertAutoStepsToISteps 批量转换 AutoStep 为 IStep
func convertAutoStepsToISteps(autoSteps []*automation.AutoStep, projectID int64, envID int) []hrp.IStep {
	var steps []hrp.IStep
	for _, autoStep := range autoSteps {
		if step := convertAutoStepToIStep(autoStep, projectID, envID); step != nil {
			steps = append(steps, step)
		}
	}
	return steps
}

// loadAutoCaseStepSteps 加载 AutoCaseStep 关联的所有 AutoStep
func loadAutoCaseStepSteps(autoCaseStepID uint) ([]*automation.AutoStep, error) {
	var relations []automation.AutoCaseStepRelation
	err := global.GVA_DB.Model(&automation.AutoCaseStepRelation{}).
		Preload("AutoStep.Request").
		Where("auto_case_step_id = ?", autoCaseStepID).
		Order("sort").
		Find(&relations).Error
	if err != nil {
		return nil, err
	}

	var steps []*automation.AutoStep
	for _, r := range relations {
		s := r.AutoStep
		// 确保 Request 被加载
		if s.Request == nil && s.RequestID != 0 {
			var req automation.Request
			global.GVA_DB.Model(&automation.Request{}).First(&req, s.RequestID)
			s.Request = &req
		}
		step := s
		steps = append(steps, &step)
	}
	return steps, nil
}

// convertAutoCaseStepToSteps 将 AutoCaseStep 转换为扁平的 IStep 列表
// 不会包装为子用例
func convertAutoCaseStepToSteps(autoCaseStep *automation.AutoCaseStep, projectID int64, envID int) ([]hrp.IStep, error) {
	// 加载关联的 AutoSteps
	autoSteps, err := loadAutoCaseStepSteps(autoCaseStep.ID)
	if err != nil {
		return nil, err
	}

	// 转换 AutoSteps 为 ISteps
	return convertAutoStepsToISteps(autoSteps, projectID, envID), nil
}

// convertAutoCaseStepToIStep 将 AutoCaseStep 转换为嵌套的 TestCase 步骤
// AutoCaseStep 包含多个 AutoStep，作为一个整体步骤执行
func convertAutoCaseStepToIStep(autoCaseStep *automation.AutoCaseStep, config *hrp.TConfig, isConfig bool, isStepConfig bool, envVars map[string]string) (hrp.IStep, error) {
	// 加载关联的 AutoSteps
	autoSteps, err := loadAutoCaseStepSteps(autoCaseStep.ID)
	if err != nil {
		return nil, err
	}

	// 转换 AutoSteps 为 ISteps
	iSteps := convertAutoStepsToISteps(autoSteps, autoCaseStep.ProjectId, int(autoCaseStep.EnvID))
	if len(iSteps) == 0 {
		return nil, nil
	}

	// 获取 CaseStepConfig
	var stepConfig hrp.StepConfig
	if isStepConfig {
		stepConfig = convertCaseStepConfigToHrpStepConfig(&autoCaseStep.StepConfig, autoCaseStep.ProjectId, int(autoCaseStep.EnvID))
	} else {
		stepConfig = hrp.StepConfig{
			StepName: autoCaseStep.StepConfig.StepName,
		}
	}

	// 创建嵌套的 TestCase
	// 组装用例（包括用例、任务、标签等场景）时，步骤内部不再携带独立的 Config，
	// 统一使用外层用例的 Config，避免出现多层 Config 嵌套导致的配置混淆。
	// 除非 isConfig 为 true 且步骤自身有关联配置。
	nestedTestCase := &hrp.TestCase{
		Config:    hrp.NewConfig(autoCaseStep.StepName),
		TestSteps: iSteps,
	}

	if isConfig && autoCaseStep.ConfigID != 0 {
		apiConfig, err := getRunConfig(uint(autoCaseStep.ConfigID))
		if err == nil && apiConfig != nil {
			stepTConfig := convertConfigToTConfig(apiConfig, envVars, autoCaseStep.ProjectId, int(autoCaseStep.EnvID))
			// 保持路径一致
			if config != nil {
				stepTConfig.Path = config.Path
			}
			nestedTestCase.Config = stepTConfig
		}
	}

	// 返回 StepTestCaseWithOptionalArgs
	return &hrp.StepTestCaseWithOptionalArgs{
		StepConfig: stepConfig,
		TestCase:   nestedTestCase,
	}, nil
}

// wrapStepsInVirtualTestCase 将步骤列表包装为虚拟用例
// 用于仅获取 "步骤" 而无外层用例结构的场景
func wrapStepsInVirtualTestCase(steps []hrp.IStep, config *hrp.TConfig, name string) *LingceTestCase {
	if config == nil {
		config = hrp.NewConfig(name)
	}
	if config.Name == "" {
		config.Name = name
	}

	return &LingceTestCase{
		Name:      name,
		Config:    config,
		TestSteps: steps,
	}
}

// buildTestCaseFromAutoCaseSteps 从 AutoCaseStepList 构建 TestCase 的步骤
func buildTestCaseFromAutoCaseSteps(caseStepList []automation.AutoCaseStepList, config *hrp.TConfig, envVars map[string]string) ([]hrp.IStep, error) {
	var steps []hrp.IStep

	for _, caseStep := range caseStepList {
		// 加载 AutoCaseStep
		var autoCaseStep automation.AutoCaseStep
		err := global.GVA_DB.Model(&automation.AutoCaseStep{}).
			First(&autoCaseStep, "id = ?", caseStep.AutoCaseStepID).Error
		if err != nil {
			global.GVA_LOG.Warn("Failed to load AutoCaseStep",
				zap.Uint("step_id", caseStep.AutoCaseStepID),
				zap.Error(err))
			continue
		}

		// 转换为 IStep
		step, err := convertAutoCaseStepToIStep(&autoCaseStep, config, caseStep.IsConfig, caseStep.IsStepConfig, envVars)
		if err != nil {
			global.GVA_LOG.Warn("Failed to convert AutoCaseStep",
				zap.Uint("step_id", caseStep.AutoCaseStepID),
				zap.Error(err))
			continue
		}
		if step != nil {
			steps = append(steps, step)
		}
	}

	return steps, nil
}

// buildTestCaseWithSetupStep 构建包含前置步骤的 TestCase
func buildTestCaseWithSetupStep(setupStepID uint, caseSteps []hrp.IStep, config *hrp.TConfig, envVars map[string]string) ([]hrp.IStep, bool, error) {
	var allSteps []hrp.IStep
	hasSetup := false

	if setupStepID != 0 {
		// 加载前置步骤
		var autoCaseStep automation.AutoCaseStep
		err := global.GVA_DB.Model(&automation.AutoCaseStep{}).
			First(&autoCaseStep, "id = ?", setupStepID).Error
		if err == nil {
			setupStep, err := convertAutoCaseStepToIStep(&autoCaseStep, config, false, true, envVars)
			if err == nil && setupStep != nil {
				allSteps = append(allSteps, setupStep)
				hasSetup = true
			}
		}
	}

	allSteps = append(allSteps, caseSteps...)
	return allSteps, hasSetup, nil
}
