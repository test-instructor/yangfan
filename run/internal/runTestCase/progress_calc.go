package runTestCase

import (
	"github.com/test-instructor/yangfan/httprunner/hrp"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/automation"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	"gorm.io/datatypes"
)

// ProgressTotals 进度统计结果
type ProgressTotals struct {
	TotalCases int // 用例总数
	TotalSteps int // 步骤总数（考虑循环）
	TotalApis  int // 接口总数（考虑循环）
}

// calcParametersCount 计算 Parameters 的循环次数
// Parameters 是一个 map，每个 key 对应一个数组，循环次数等于数组长度的乘积
// 例如: {"a": [1,2], "b": [3,4,5]} 循环次数为 2 * 3 = 6
func calcParametersCount(params datatypes.JSONMap) int {
	if params == nil || len(params) == 0 {
		return 1
	}

	count := 1
	for _, v := range params {
		// Parameters 的值应该是数组
		if arr, ok := v.([]interface{}); ok && len(arr) > 0 {
			count *= len(arr)
		}
	}

	if count == 0 {
		return 1
	}
	return count
}

// calcApiCountFromAutoStep 计算单个 AutoStep（接口）的实际执行次数
// 接口的 Parameters 只循环单个接口
func calcApiCountFromAutoStep(step *automation.AutoStep) int {
	if step == nil {
		return 0
	}
	return calcParametersCount(step.Parameters)
}

// calcStepApiCount 计算步骤集合（AutoCaseStep）的接口执行次数
// 步骤集合的 Parameters 循环当前步骤下的所有接口
func calcStepApiCount(stepID uint, stepParams datatypes.JSONMap) int {
	// 加载步骤集合关联的所有接口
	var relations []automation.AutoCaseStepRelation
	err := global.GVA_DB.Model(&automation.AutoCaseStepRelation{}).
		Preload("AutoStep").
		Where("auto_case_step_id = ?", stepID).
		Find(&relations).Error
	if err != nil {
		return 0
	}

	// 步骤集合的循环次数
	stepLoopCount := calcParametersCount(stepParams)

	// 统计所有接口的执行次数
	totalApiCount := 0
	for _, r := range relations {
		// 每个接口的执行次数 = 步骤循环次数 * 接口自身的循环次数
		apiLoopCount := calcParametersCount(r.AutoStep.Parameters)
		totalApiCount += stepLoopCount * apiLoopCount
	}

	return totalApiCount
}

// calcStepTotals 计算步骤集合的统计信息
// 返回: (步骤执行次数, 接口执行次数)
func calcStepTotals(stepID uint, stepParams datatypes.JSONMap) (int, int) {
	// 加载步骤集合关联的所有接口
	var relations []automation.AutoCaseStepRelation
	err := global.GVA_DB.Model(&automation.AutoCaseStepRelation{}).
		Preload("AutoStep").
		Where("auto_case_step_id = ?", stepID).
		Find(&relations).Error
	if err != nil {
		return 0, 0
	}

	// 步骤集合的循环次数
	stepLoopCount := calcParametersCount(stepParams)

	// 统计所有接口的执行次数
	totalApiCount := 0
	for _, r := range relations {
		// 每个接口的执行次数 = 步骤循环次数 * 接口自身的循环次数
		apiLoopCount := calcParametersCount(r.AutoStep.Parameters)
		totalApiCount += stepLoopCount * apiLoopCount
	}

	// 步骤执行次数 = 步骤循环次数
	return stepLoopCount, totalApiCount
}

// calcCaseTotals 计算用例的总执行次数
// configParams: RunConfig 的 Parameters，会循环执行所有步骤集合下所有接口
// caseSteps: 用例关联的步骤列表
func calcCaseTotals(configParams datatypes.JSONMap, caseSteps []automation.AutoCaseStepList) ProgressTotals {
	// RunConfig 的循环次数
	configLoopCount := calcParametersCount(configParams)

	totalSteps := 0
	totalApis := 0

	for _, caseStep := range caseSteps {
		// 加载步骤集合详情获取其 Parameters
		var autoCaseStep automation.AutoCaseStep
		err := global.GVA_DB.Model(&automation.AutoCaseStep{}).
			Where("id = ?", caseStep.AutoCaseStepID).
			First(&autoCaseStep).Error
		if err != nil {
			continue
		}

		// 计算步骤集合的统计
		stepCount, apiCount := calcStepTotals(caseStep.AutoCaseStepID, autoCaseStep.Parameters)

		// 乘以配置的循环次数
		totalSteps += configLoopCount * stepCount
		totalApis += configLoopCount * apiCount
	}

	return ProgressTotals{
		TotalCases: 1,
		TotalSteps: totalSteps,
		TotalApis:  totalApis,
	}
}

// CalcCaseTotalsFromISteps 从 IStep 列表计算进度统计
// 用于在 LoadCase 时根据已组装的步骤计算总量
func CalcCaseTotalsFromISteps(steps []hrp.IStep, config *hrp.TConfig) ProgressTotals {
	configLoopCount := 1
	if config != nil && config.Parameters != nil {
		configLoopCount = calcParametersCountFromInterface(config.Parameters)
	}

	totalSteps := 0
	totalApis := 0

	for _, step := range steps {
		stepCount, apiCount := calcIStepTotals(step)
		totalSteps += configLoopCount * stepCount
		totalApis += configLoopCount * apiCount
	}

	// TotalCases 应该考虑配置的参数化循环次数
	return ProgressTotals{
		TotalCases: configLoopCount,
		TotalSteps: totalSteps,
		TotalApis:  totalApis,
	}
}

// calcParametersCountFromInterface 从 map[string]interface{} 计算循环次数
func calcParametersCountFromInterface(params map[string]interface{}) int {
	if params == nil || len(params) == 0 {
		return 1
	}

	count := 1
	for _, v := range params {
		if arr, ok := v.([]interface{}); ok && len(arr) > 0 {
			count *= len(arr)
		}
	}

	if count == 0 {
		return 1
	}
	return count
}

// calcIStepTotals 计算单个 IStep 的执行次数
// 返回: (步骤集合执行次数, 接口执行次数)
// 步骤集合 = StepTestCaseWithOptionalArgs（嵌套用例）
// 接口 = StepRequestWithOptionalArgs（单个请求）
func calcIStepTotals(step hrp.IStep) (int, int) {
	if step == nil {
		return 0, 0
	}

	// 尝试断言为不同类型的步骤
	switch s := step.(type) {
	case *hrp.StepRequestWithOptionalArgs:
		// 单个请求步骤（接口）
		// stepCount = 0（不是步骤集合）
		// apiCount = 循环次数
		loopCount := 1
		if s.StepConfig.Parameters != nil {
			loopCount = calcParametersCountFromInterface(s.StepConfig.Parameters)
		}
		return 0, loopCount

	case *hrp.StepTestCaseWithOptionalArgs:
		// 嵌套用例步骤（步骤集合）
		// stepCount = 循环次数
		// apiCount = 内部接口总数 * 循环次数
		loopCount := 1
		if s.StepConfig.Parameters != nil {
			loopCount = calcParametersCountFromInterface(s.StepConfig.Parameters)
		}

		// 计算嵌套用例内部的接口数
		innerApiCount := 0
		if s.TestCase != nil {
			// TestCase 是 interface{} 类型，需要断言为 *hrp.TestCase
			if tc, ok := s.TestCase.(*hrp.TestCase); ok && tc != nil {
				for _, innerStep := range tc.TestSteps {
					_, apiCount := calcIStepTotals(innerStep)
					innerApiCount += apiCount
				}
			}
		}

		return loopCount, loopCount * innerApiCount

	default:
		// 其他类型步骤，返回 0, 1
		return 0, 1
	}
}

// CalcTaskTotals 计算任务的总执行次数
// 任务包含多个用例，每个用例可能有不同的配置
func CalcTaskTotals(taskCases []automation.TimerTaskCaseList, configParams datatypes.JSONMap) ProgressTotals {
	totalCases := len(taskCases)
	totalSteps := 0
	totalApis := 0

	configLoopCount := calcParametersCount(configParams)

	for _, taskCase := range taskCases {
		// 获取用例关联的步骤
		caseStepList := caseSort(taskCase.AutoCaseID)

		// 计算每个用例的步骤和接口数
		caseTotals := calcCaseTotals(nil, caseStepList)
		totalSteps += configLoopCount * caseTotals.TotalSteps
		totalApis += configLoopCount * caseTotals.TotalApis
	}

	return ProgressTotals{
		TotalCases: totalCases,
		TotalSteps: totalSteps,
		TotalApis:  totalApis,
	}
}

// CalcApiTotals 计算单接口运行的统计
func CalcApiTotals(apiStep *automation.AutoStep, setupStepID uint, config *platform.RunConfig) ProgressTotals {
	configLoopCount := 1
	if config != nil {
		configLoopCount = calcParametersCount(config.Parameters)
	}

	totalApis := 0

	// 前置步骤
	if setupStepID != 0 {
		var setupStep automation.AutoCaseStep
		if err := global.GVA_DB.Model(&automation.AutoCaseStep{}).First(&setupStep, "id = ?", setupStepID).Error; err == nil {
			_, setupApiCount := calcStepTotals(setupStepID, setupStep.Parameters)
			totalApis += configLoopCount * setupApiCount
		}
	}

	// 主接口
	apiLoopCount := calcApiCountFromAutoStep(apiStep)
	totalApis += configLoopCount * apiLoopCount

	return ProgressTotals{
		TotalCases: 1,
		TotalSteps: totalApis, // 对于单接口，步骤数等于接口数
		TotalApis:  totalApis,
	}
}

// CalcStepTotals 计算单步骤集合运行的统计
func CalcStepTotals(stepID uint, setupStepID uint, config *platform.RunConfig) ProgressTotals {
	configLoopCount := 1
	if config != nil {
		configLoopCount = calcParametersCount(config.Parameters)
	}

	totalSteps := 0
	totalApis := 0

	// 前置步骤
	if setupStepID != 0 {
		var setupStep automation.AutoCaseStep
		if err := global.GVA_DB.Model(&automation.AutoCaseStep{}).First(&setupStep, "id = ?", setupStepID).Error; err == nil {
			setupStepCount, setupApiCount := calcStepTotals(setupStepID, setupStep.Parameters)
			totalSteps += configLoopCount * setupStepCount
			totalApis += configLoopCount * setupApiCount
		}
	}

	// 主步骤集合
	var mainStep automation.AutoCaseStep
	if err := global.GVA_DB.Model(&automation.AutoCaseStep{}).First(&mainStep, "id = ?", stepID).Error; err == nil {
		mainStepCount, mainApiCount := calcStepTotals(stepID, mainStep.Parameters)
		totalSteps += configLoopCount * mainStepCount
		totalApis += configLoopCount * mainApiCount
	}

	return ProgressTotals{
		TotalCases: 1,
		TotalSteps: totalSteps,
		TotalApis:  totalApis,
	}
}
