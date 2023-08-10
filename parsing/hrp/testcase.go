package hrp

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/test-instructor/yangfan/parsing/hrp/internal/builtin"
	"github.com/test-instructor/yangfan/parsing/hrp/internal/code"
)

// ITestCase represents interface for testcases,
// includes TestCase and TestCasePath.
type ITestCase interface {
	GetPath() string
	ToTestCase() (*TestCase, error)
}

// TestCase is a container for one testcase, which is used for testcase runner.
// TestCase implements ITestCase interface.
type TestCase struct {
	Config    *TConfig
	TestSteps []IStep
}

func (tc *TestCase) GetPath() string {
	return tc.Config.Path
}

func (tc *TestCase) ToTestCase() (*TestCase, error) {
	return tc, nil
}

func (tc *TestCase) ToTCase() *TCase {
	// 将当前的 TestCase 转换为 TCase 类型
	tCase := &TCase{
		Config: tc.Config, // 将当前的 Config 属性赋值给 tCase 的 Config 属性
	}
	// 遍历当前 TestCase 的所有 TestSteps
	for _, step := range tc.TestSteps {
		// 检查当前 step 的类型是否为 TestCase 类型
		if step.Type() == stepTypeTestCase {
			// 将 step 强制转换为 TestCase 类型，并判断是否成功
			if testcase, ok := step.Struct().TestCase.(*TestCase); ok {
				// 递归调用 ToTCase() 方法将内部的 TestCase 转换为 TCase
				step.Struct().TestCase = testcase.ToTCase()
			}
		}
		// 将转换后的 step 添加到 tCase 的 TestSteps 切片中
		tCase.TestSteps = append(tCase.TestSteps, step.Struct())
	}
	// 返回转换后的 tCase
	return tCase
}

func (tc *TestCase) Dump2JSON(targetPath string) error {
	// 将当前的 json 用例转换为 TCase 类型
	tCase := tc.ToTCase()
	// 使用内置的 Dump2JSON 函数将 tCase 转换为 JSON，并将结果保存到目标路径
	err := builtin.Dump2JSON(tCase, targetPath)
	// 如果转换或保存过程中出现错误，则返回一个错误，并附加错误信息
	if err != nil {
		return errors.Wrap(err, "dump testcase to json failed")
	}
	return nil
}

func (tc *TestCase) Dump2YAML(targetPath string) error {
	// 将当前的 yaml 用例转换为 TCase 类型
	tCase := tc.ToTCase()
	// 使用内置的 Dump2YAML 函数将 tCase 转换为 YAML，并将结果保存到目标路径
	err := builtin.Dump2YAML(tCase, targetPath)
	// 如果转换或保存过程中出现错误，则返回一个错误，并附加错误信息
	if err != nil {
		return errors.Wrap(err, "dump testcase to yaml failed")
	}
	return nil
}

// TestCasePath implements ITestCase interface.
type TestCasePath string

func (path *TestCasePath) GetPath() string {
	return fmt.Sprintf("%v", *path)
}

// ToTestCase loads testcase path and convert to *TestCase
func (path *TestCasePath) ToTestCase() (*TestCase, error) {
	// 创建一个空的 TCase 实例
	tc := &TCase{}
	// 获取 TestCasePath 的路径
	casePath := path.GetPath()
	// 使用内置的 LoadFile 函数从文件中加载内容到测试用例类型`TCase`
	err := builtin.LoadFile(casePath, tc)
	// 如果加载过程中出现错误，则返回错误
	if err != nil {
		return nil, err
	}
	// 调用 tc 的 ToTestCase 方法将其转换为 TestCase 实例，并传递 casePath 作为参数
	return tc.ToTestCase(casePath)
}

// TCase represents testcase data structure.
// Each testcase includes one public config and several sequential teststeps.
type TCase struct {
	Config    *TConfig `json:"config" yaml:"config"`
	TestSteps []*TStep `json:"teststeps" yaml:"teststeps"`
}

// MakeCompat converts TCase compatible with Golang engine style
// 目的是使TCase结构与Golang引擎风格兼容。它通过对TCase对象及其相关的TestSteps执行某些兼容性转换来实现
func (tc *TCase) MakeCompat() (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("[MakeCompat] convert compat testcase error: %v", p)
		}
	}()
	for _, step := range tc.TestSteps {
		// 1. deal with request body compatibility
		// 将请求体转换为与Golang引擎风格兼容
		convertCompatRequestBody(step.Request)

		// 2. deal with validators compatibility
		// 将断言转换为与Golang引擎风格兼容
		err = convertCompatValidator(step.Validators)
		if err != nil {
			return err
		}

		// 3. deal with extract expr including hyphen
		// 将提取表达式转换为与Golang引擎风格兼容
		convertExtract(step.Extract)
	}
	return nil
}

func (tc *TCase) ToTestCase(casePath string) (*TestCase, error) {
	// 如果 TestSteps 为空时返回一个错误
	if tc.TestSteps == nil {
		return nil, errors.Wrap(code.InvalidCaseFormat,
			"invalid testcase format, missing teststeps!")
	}
	// 如果 Config 为空时，创建一个新的 Config 实例
	if tc.Config == nil {
		tc.Config = &TConfig{Name: "please input testcase name"}
	}
	// 将用例路径写入到 Config 属性中
	tc.Config.Path = casePath
	return tc.toTestCase()
}

// toTestCase converts *TCase to *TestCase
func (tc *TCase) toTestCase() (*TestCase, error) {
	testCase := &TestCase{
		Config: tc.Config,
	}
	// 使TCase结构与Golang引擎风格兼容
	err := tc.MakeCompat()
	if err != nil {
		return nil, err
	}

	// locate project root dir by plugin path
	// 根据插件路径定位项目根目录
	// 环境变量和函数插件都是通过项目根目录来定位的
	projectRootDir, err := GetProjectRootDirPath(tc.Config.Path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get project root dir")
	}

	// load .env file
	// 加载 .env 文件
	dotEnvPath := filepath.Join(projectRootDir, ".env")
	if builtin.IsFilePathExists(dotEnvPath) {
		envVars := make(map[string]string)
		err = builtin.LoadFile(dotEnvPath, envVars)
		if err != nil {
			return nil, errors.Wrap(err, "failed to load .env file")
		}

		// override testcase config env with variables loaded from .env file
		// priority: .env file > testcase config env
		// 使用从 .env 文件加载的变量覆盖测试用例配置中的环境变量
		// 优先级：.env 文件 > 测试用例配置中的环境变量
		// 3.x 版本中，.env 会加载到系统的临时环境变量中，4.x 则是加载到配置中
		if testCase.Config.Environs == nil {
			testCase.Config.Environs = make(map[string]string)
		}
		for key, value := range envVars {
			testCase.Config.Environs[key] = value
		}
	}

	// 遍历测试步骤，根据类型进行处理并添加到 testCase 中
	// 所有类型的测试步骤都会被转换为 APITestStep 类型
	// 如果要添加自定义协议，则需要在这里增加对应的类型
	for _, step := range tc.TestSteps {

		if step.API != nil {
			// 处理api步骤, 如果 api 为引用其他json/yaml 文件
			apiPath, ok := step.API.(string)
			if ok {
				// 获取文件并转换成用例
				path := filepath.Join(projectRootDir, apiPath)
				if !builtin.IsFilePathExists(path) {
					return nil, errors.Wrap(code.ReferencedFileNotFound,
						fmt.Sprintf("referenced api file not found: %s", path))
				}

				refAPI := APIPath(path)
				apiContent, err := refAPI.ToAPI()
				if err != nil {
					return nil, err
				}
				step.API = apiContent
			} else {
				// 如果 api 不是引用其他文件，则直接转换成 API 实例
				apiMap, ok := step.API.(map[string]interface{})
				if !ok {
					return nil, errors.Wrap(code.InvalidCaseFormat,
						fmt.Sprintf("referenced api should be map or path(string), got %v", step.API))
				}
				api := &API{}
				err = mapstructure.Decode(apiMap, api)
				if err != nil {
					return nil, err
				}
				step.API = api
			}
			_, ok = step.API.(*API)
			if !ok {
				return nil, errors.Wrap(code.InvalidCaseFormat,
					fmt.Sprintf("failed to handle referenced API, got %v", step.TestCase))
			}
			testCase.TestSteps = append(testCase.TestSteps, &StepAPIWithOptionalArgs{
				step: step,
			})
		} else if step.TestCase != nil {
			// 处理 testcase 步骤
			// 这里和处理 api 步骤类似
			casePath, ok := step.TestCase.(string)
			if ok {
				path := filepath.Join(projectRootDir, casePath)
				if !builtin.IsFilePathExists(path) {
					return nil, errors.Wrap(code.ReferencedFileNotFound,
						fmt.Sprintf("referenced testcase file not found: %s", path))
				}

				refTestCase := TestCasePath(path)
				tc, err := refTestCase.ToTestCase()
				if err != nil {
					return nil, err
				}
				step.TestCase = tc
			} else {
				testCaseMap, ok := step.TestCase.(map[string]interface{})
				if !ok {
					return nil, errors.Wrap(code.InvalidCaseFormat,
						fmt.Sprintf("referenced testcase should be map or path(string), got %v", step.TestCase))
				}
				tCase := &TCase{}
				err = mapstructure.Decode(testCaseMap, tCase)
				if err != nil {
					return nil, err
				}
				tc, err := tCase.toTestCase()
				if err != nil {
					return nil, err
				}
				step.TestCase = tc
			}
			_, ok = step.TestCase.(*TestCase)
			if !ok {
				return nil, errors.Wrap(code.InvalidCaseFormat,
					fmt.Sprintf("failed to handle referenced testcase, got %v", step.TestCase))
			}
			testCase.TestSteps = append(testCase.TestSteps, &StepTestCaseWithOptionalArgs{
				step: step,
			})
		} else if step.ThinkTime != nil {
			// 处理 thinktime 步骤
			testCase.TestSteps = append(testCase.TestSteps, &StepThinkTime{
				step: step,
			})
		} else if step.Request != nil {
			// init upload
			// 处理 request 步骤
			if len(step.Request.Upload) != 0 {
				initUpload(step)
			}
			testCase.TestSteps = append(testCase.TestSteps, &StepRequestWithOptionalArgs{
				step: step,
			})
		} else if step.Transaction != nil {
			// 处理 transaction 步骤
			testCase.TestSteps = append(testCase.TestSteps, &StepTransaction{
				step: step,
			})
		} else if step.Rendezvous != nil {
			// 处理 rendezvous 步骤
			testCase.TestSteps = append(testCase.TestSteps, &StepRendezvous{
				step: step,
			})
		} else if step.WebSocket != nil {
			// 处理 websocket 步骤
			testCase.TestSteps = append(testCase.TestSteps, &StepWebSocket{
				step: step,
			})
		} else if step.IOS != nil {
			// 处理 ios 步骤
			testCase.TestSteps = append(testCase.TestSteps, &StepMobile{
				step: step,
			})
		} else if step.Android != nil {
			// 处理 android 步骤
			testCase.TestSteps = append(testCase.TestSteps, &StepMobile{
				step: step,
			})
		} else {
			// 处理未知步骤
			log.Warn().Interface("step", step).Msg("[convertTestCase] unexpected step")
		}
	}
	return testCase, nil
}

func convertCompatRequestBody(request *Request) {
	// 检查 request 是否为空，以及是否已经设置了 Body
	if request != nil && request.Body == nil {
		// 设置请求参数，并根据设置请求头内容
		if request.Json != nil {
			// 如果 Headers 为空，则创建一个空的 Headers map
			if request.Headers == nil {
				request.Headers = make(map[string]string)
			}
			request.Headers["Content-Type"] = "application/json; charset=utf-8"

			request.Body = request.Json
			request.Json = nil
		} else if request.Data != nil {
			// 如果存在 Data 数据，则将 Data 设置为 Body，并将 Data 字段置为空
			request.Body = request.Data
			request.Data = nil
		}
	}
}

func convertCompatValidator(Validators []interface{}) (err error) {
	for i, iValidator := range Validators {
		// 检查当前 Validator 是否已经是 Validator 类型，如果是则跳过处理
		if _, ok := iValidator.(Validator); ok {
			continue
		}
		// 将当前 Validator 转换为 map[string]interface{} 类型
		validatorMap := iValidator.(map[string]interface{})
		validator := Validator{}
		iCheck, checkExisted := validatorMap["check"]
		iAssert, assertExisted := validatorMap["assert"]
		iExpect, expectExisted := validatorMap["expect"]
		// validator check priority: Golang > Python engine style
		// 断言风格，python 风格则和 v3.x 一致，golang 风格则和 v4.x 一致
		// 个人比较喜欢golang 风格，因此在测试平台的实现上优先使用 golang 风格
		if checkExisted && assertExisted && expectExisted {
			// Golang engine style
			validator.Check = iCheck.(string)
			validator.Assert = iAssert.(string)
			validator.Expect = iExpect
			if iMsg, msgExisted := validatorMap["msg"]; msgExisted {
				validator.Message = iMsg.(string)
			}
			validator.Check = convertJmespathExpr(validator.Check)
			Validators[i] = validator
			continue
		}
		if len(validatorMap) == 1 {
			// Python engine style
			for assertMethod, iValidatorContent := range validatorMap {
				validatorContent := iValidatorContent.([]interface{})
				if len(validatorContent) > 3 {
					return errors.Wrap(code.InvalidCaseFormat,
						fmt.Sprintf("unexpected validator format: %v", validatorMap))
				}
				validator.Check = validatorContent[0].(string)
				validator.Assert = assertMethod
				validator.Expect = validatorContent[1]
				if len(validatorContent) == 3 {
					validator.Message = validatorContent[2].(string)
				}
			}
			validator.Check = convertJmespathExpr(validator.Check)
			Validators[i] = validator
			continue
		}
		return errors.Wrap(code.InvalidCaseFormat,
			fmt.Sprintf("unexpected validator format: %v", validatorMap))
	}
	return nil
}

// convertExtract deals with extract expr including hyphen
func convertExtract(extract map[string]string) {
	// 设置导出参数内容
	for key, value := range extract {
		extract[key] = convertJmespathExpr(value)
	}
}

// convertJmespathExpr deals with limited jmespath expression conversion
func convertJmespathExpr(checkExpr string) string {
	// 检查检查表达式中是否包含文本提取器的子正则表达式，如果包含则直接返回原始表达式
	if strings.Contains(checkExpr, textExtractorSubRegexp) {
		return checkExpr
	}
	// 将检查表达式按照 "." 分割为多个项
	checkItems := strings.Split(checkExpr, ".")
	for i, checkItem := range checkItems {
		// 去除检查项的首尾双引号
		checkItem = strings.Trim(checkItem, "\"")
		// 将检查项转换为小写字母形式
		lowerItem := strings.ToLower(checkItem)
		// 如果检查项以 "content-" 开头或者为 "user-agent"，则在检查项两侧添加双引号
		if strings.HasPrefix(lowerItem, "content-") || lowerItem == "user-agent" {
			checkItems[i] = fmt.Sprintf("\"%s\"", checkItem)
		}
	}
	// 将处理后的检查项重新组合为检查表达式
	return strings.Join(checkItems, ".")
}
