package hrp

import (
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/rs/zerolog/log"

	"github.com/test-instructor/yangfan/parsing/hrp/internal/builtin"
)

// IAPI represents interface for api,
// includes API and APIPath.
// API实现有2中方式，一种是直接定义API结构体，另一种是定义APIPath结构体，然后通过APIPath结构体的GetPath()方法获取API结构体
type IAPI interface {
	GetPath() string
	ToAPI() (*API, error)
}

// API 直接定义结构体
type API struct {
	Name          string                 `json:"name" yaml:"name"` // required
	Request       *Request               `json:"request,omitempty" yaml:"request,omitempty"`
	Variables     map[string]interface{} `json:"variables,omitempty" yaml:"variables,omitempty"`
	SetupHooks    []string               `json:"setup_hooks,omitempty" yaml:"setup_hooks,omitempty"`
	TeardownHooks []string               `json:"teardown_hooks,omitempty" yaml:"teardown_hooks,omitempty"`
	Extract       map[string]string      `json:"extract,omitempty" yaml:"extract,omitempty"`
	Validators    []interface{}          `json:"validate,omitempty" yaml:"validate,omitempty"`
	Export        []string               `json:"export,omitempty" yaml:"export,omitempty"`
	Path          string
}

func (api *API) GetPath() string {
	return api.Path
}

// ToAPI 直接定义的结构体的API不需要进行转换
func (api *API) ToAPI() (*API, error) {
	return api, nil
}

// APIPath implements IAPI interface.
type APIPath string

func (path *APIPath) GetPath() string {
	return fmt.Sprintf("%v", *path)
}

// ToAPI 通过路径引用的API需要将文件内容转换为API结构体
func (path *APIPath) ToAPI() (*API, error) {
	api := &API{}
	// 获取文件路径
	apiPath := path.GetPath()
	// 加载文件内容并转换成API类型
	err := builtin.LoadFile(apiPath, api)
	if err != nil {
		return nil, err
	}
	// 1. deal with request body compatibility
	// 处理请求数据
	convertCompatRequestBody(api.Request)
	// 2. deal with validators compatibility
	// 处理断言数据
	err = convertCompatValidator(api.Validators)
	// 3. deal with extract expr including hyphen
	// 处理导出数据
	convertExtract(api.Extract)
	return api, err
}

// StepAPIWithOptionalArgs implements IStep interface.
// 定义API类型的步骤
type StepAPIWithOptionalArgs struct {
	step *TStep
}

// TeardownHook adds a teardown hook for current teststep.
// 设置步骤的teardown hook，主要用户清理数据
func (s *StepAPIWithOptionalArgs) TeardownHook(hook string) *StepAPIWithOptionalArgs {
	s.step.TeardownHooks = append(s.step.TeardownHooks, hook)
	return s
}

// Export specifies variable names to export from referenced api for current step.
// 设置步骤导出的变量
func (s *StepAPIWithOptionalArgs) Export(names ...string) *StepAPIWithOptionalArgs {
	api, ok := s.step.API.(*API)
	if ok {
		s.step.Export = append(api.Export, names...)
	}
	return s
}

func (s *StepAPIWithOptionalArgs) Name() string {
	if s.step.Name != "" {
		return s.step.Name
	}
	api, ok := s.step.API.(*API)
	if ok {
		return api.Name
	}
	return ""
}

// Type 设置类型
func (s *StepAPIWithOptionalArgs) Type() StepType {
	return stepTypeAPI
}

// Struct 返回结构体
func (s *StepAPIWithOptionalArgs) Struct() *TStep {
	return s.step
}

// Run 执行步骤
func (s *StepAPIWithOptionalArgs) Run(r *SessionRunner) (stepResult *StepResult, err error) {
	defer func() {
		stepResult.StepType = stepTypeAPI
	}()
	// extend request with referenced API
	// 获取类型为interface的API数据并将其转换成API类型
	api, _ := s.step.API.(*API)
	step := &TStep{}
	// deep copy step to avoid data racing
	// 深拷贝步骤，避免数据竞争，主要是性能测试并发执行时会出现数据竞争
	if err = copier.Copy(step, s.step); err != nil {
		log.Error().Err(err).Msg("copy step failed")
		return
	}
	// 将api转换成可执行的步骤步骤
	extendWithAPI(step, api)
	// 执行并收集测试报告
	stepResult, err = runStepRequest(r, step)
	return
}

// extend teststep with api, teststep will merge and override referenced api
// 将api转换成步骤
func extendWithAPI(testStep *TStep, overriddenStep *API) {
	// override api name
	if testStep.Name == "" {
		testStep.Name = overriddenStep.Name
	}
	// merge & override request
	// 获取步骤的请求数据
	testStep.Request = overriddenStep.Request
	// init upload
	// 判断api中是否有上传文件
	if len(testStep.Request.Upload) != 0 {
		initUpload(testStep)
	}
	// merge & override variables
	// 设置步骤的变量
	testStep.Variables = mergeVariables(testStep.Variables, overriddenStep.Variables)
	// merge & override extractors
	// 设置步骤的提取数据
	testStep.Extract = mergeMap(testStep.Extract, overriddenStep.Extract)
	// merge & override validators
	// 设置步骤的断言
	testStep.Validators = mergeValidators(testStep.Validators, overriddenStep.Validators)
	// merge & override setupHooks
	// 设置步骤的setup hook
	testStep.SetupHooks = mergeSlices(testStep.SetupHooks, overriddenStep.SetupHooks)
	// merge & override teardownHooks
	// 设置步骤的teardown hook
	testStep.TeardownHooks = mergeSlices(testStep.TeardownHooks, overriddenStep.TeardownHooks)
}
