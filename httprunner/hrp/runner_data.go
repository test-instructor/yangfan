package hrp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/test-instructor/yangfan/httprunner/code"
	"github.com/test-instructor/yangfan/httprunner/internal/builtin"
	"github.com/test-instructor/yangfan/httprunner/internal/sdk"
)

// RunJsons starts to run testcases converted from json content.
func (r *HRPRunner) RunJsons(testcases ...ITestCase) (data []byte, err error) {
	event := sdk.Event{
		Name: "RunAPITests",
		Params: map[string]interface{}{
			"action": "hrp run",
		},
	}
	// report start event
	go sdk.SendGA4Event(event.Name, event.Params)
	// report execution timing event
	defer sdk.SendGA4Event("execution_timing", map[string]interface{}{"start": time.Now()})

	// record execution data to summary
	s := NewSummary()

	// load all testcases
	testCases, err := LoadTestCases(testcases...)
	if err != nil {
		log.Error().Err(err).Msg("run json failed to load testcases")
		return
	}

	var runErr error

	for _, testCase := range testCases {
		// each testcase has its own case runner
		caseRunner, err := NewCaseRunner(*testCase, r)
		if err != nil {
			log.Error().Err(err).Msg("run json failed to init case runner")
			return nil, err
		}

		for it := caseRunner.parametersIterator; it.HasNext(); {
			sessionRunner := caseRunner.NewSession()
			caseSummary, err1 := sessionRunner.Start(it.Next())
			if err1 != nil {
				log.Error().Err(err1).Msg("run json failed to run testcase")
				runErr = err1
			}
			if caseSummary != nil {
				caseSummary.Name = testCase.Config.Get().Name
				s.AddCaseSummary(caseSummary)
			}
		}
	}
	s.Time.Duration = time.Since(s.Time.StartAt).Seconds()

	// save summary
	if r.saveTests {
		_, err = s.GenSummary()
		if err != nil {
			return
		}
	}

	// generate HTML report
	if r.genHTMLReport {
		err = s.GenHTMLReport()
		if err != nil {
			return
		}
	}
	sj, _ := json.Marshal(s)
	return sj, runErr
}

// RunJson starts to run testcases converted from json content.
func (r *HRPRunner) RunJson(jsonContent []byte) error {
	log.Info().Msg("start running json testcases")

	// load testcases from json content
	testCases, err := LoadTestCasesFromJSON(jsonContent)
	if err != nil {
		log.Error().Err(err).Msg("failed to load testcases from json")
		return err
	}

	_, err = r.RunJsons(testCases...)
	return err
}

func RunJson(t *testing.T, jsonContent []byte) error {
	return NewRunner(t).SetSaveTests(true).RunJson(jsonContent)
}

// Helper functions from yangfan
func tmpls(relativePath, debugTalkFileName string) string {
	return filepath.Join(debugTalkFileName, relativePath)
}

func BuildHashicorpPyPlugin(debugTalkByte []byte, debugTalkFilePath string) {
	log.Info().Str("debugTalkFilePath", debugTalkFilePath).Msg("prepare hashicorp python plugin")
	err := os.WriteFile(tmpls("debugtalk.py", debugTalkFilePath), debugTalkByte, 0o644)
	if err != nil {
		log.Error().Err(err).Msg("copy hashicorp python plugin failed")
		os.Exit(code.GetErrorCode(err))
	}
}

func RemoveHashicorpPyPlugin(debugTalkFilePath string) {
	log.Info().Str("debugTalkFilePath", debugTalkFilePath).Msg("remove hashicorp python plugin")
	err := os.RemoveAll(debugTalkFilePath)
	if err != nil {
		log.Error().Err(err).Msg("remove hashicorp python plugin failed")
	}
}

// TestCaseJson implements ITestCase interface
type TestCaseJson struct {
	JsonString        string
	ID                uint
	DebugTalkFilePath string
	Config            *TConfig
	Name              string
}

func (testCaseJson *TestCaseJson) GetPath() string {
	return testCaseJson.DebugTalkFilePath
}

func (testCaseJson *TestCaseJson) GetTestCase() (*TestCase, error) {
	return testCaseJson.ToTestCase()
}

func (testCaseJson *TestCaseJson) ToTestCase() (*TestCase, error) {
	tc := &TestCaseDef{}
	var err error
	casePath := testCaseJson.JsonString
	tc, err = loadFromString(casePath)
	if err != nil {
		return nil, err
	}

	if tc.Config == nil {
		tc.Config = &TConfig{}
	}
	tc.Config.Path = testCaseJson.GetPath()
	if testCaseJson.Config != nil {
		testCaseJson.Config.Path = testCaseJson.GetPath()
	}

	testCase := &TestCase{
		Config: testCaseJson.Config,
	}
	if testCase.Config == nil {
		testCase.Config = tc.Config
	}
	if testCaseJson.Name != "" && testCase.Config != nil {
		testCase.Config.Get().Name = testCaseJson.Name
	}

	projectRootDir, err := GetProjectRootDirPath(testCaseJson.GetPath())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get project root dir")
	}

	dotEnvPath := filepath.Join(projectRootDir, ".env")
	if builtin.IsFilePathExists(dotEnvPath) {
		envVars := make(map[string]string)
		envVarsByte, err := builtin.LoadFile(dotEnvPath)
		if err != nil {
			return nil, errors.Wrap(err, "failed to load .env file")
		}
		// parse .env file
		envVarsMap, err := godotenv.Unmarshal(string(envVarsByte))
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse .env file")
		}
		for key, value := range envVarsMap {
			envVars[key] = value
		}

		if testCase.Config.Get().Environs == nil {
			testCase.Config.Get().Environs = make(map[string]string)
		}
		for key, value := range envVars {
			testCase.Config.Get().Environs[key] = value
		}
	}

	for _, step := range tc.Steps {
		if step.TestCase != nil {
			testcaseYangfanStr, _ := json.Marshal(step.TestCase)
			tcj := &TestCaseJson{
				JsonString:        string(testcaseYangfanStr),
				ID:                0,
				DebugTalkFilePath: testCaseJson.GetPath(),
				Config:            testCaseJson.Config,
				Name:              testCase.Config.Get().Name,
			}
			tc, err := tcj.ToTestCase()
			if err != nil {
				return nil, err
			}
			step.TestCase = tc
			_, ok := step.TestCase.(*TestCase)
			if !ok {
				return nil, fmt.Errorf("failed to handle referenced testcase, got %v", step.TestCase)
			}
			testCase.TestSteps = append(testCase.TestSteps, &StepTestCaseWithOptionalArgs{
				StepConfig: step.StepConfig,
				TestCase:   step.TestCase,
			})
		} else if step.ThinkTime != nil {
			testCase.TestSteps = append(testCase.TestSteps, &StepThinkTime{
				StepConfig: step.StepConfig,
				ThinkTime:  step.ThinkTime,
			})
		} else if step.Request != nil {
			testCase.TestSteps = append(testCase.TestSteps, &StepRequestWithOptionalArgs{
				StepRequest: &StepRequest{
					StepConfig: step.StepConfig,
					Request:    step.Request,
				},
			})
		} else if step.Transaction != nil {
			testCase.TestSteps = append(testCase.TestSteps, &StepTransaction{
				StepConfig:  step.StepConfig,
				Transaction: step.Transaction,
			})
		} else if step.Rendezvous != nil {
			testCase.TestSteps = append(testCase.TestSteps, &StepRendezvous{
				StepConfig: step.StepConfig,
				Rendezvous: step.Rendezvous,
			})
		} else if step.WebSocket != nil {
			testCase.TestSteps = append(testCase.TestSteps, &StepWebSocket{
				StepConfig: step.StepConfig,
				WebSocket:  step.WebSocket,
			})
		} else if step.Android != nil {
			testCase.TestSteps = append(testCase.TestSteps, &StepMobile{
				StepConfig: step.StepConfig,
				Android:    step.Android,
			})
		} else if step.IOS != nil {
			testCase.TestSteps = append(testCase.TestSteps, &StepMobile{
				StepConfig: step.StepConfig,
				IOS:        step.IOS,
			})
		} else {
			log.Warn().Interface("step", step).Msg("unexpected step")
		}
	}
	return testCase, nil
}

func loadFromString(jsonString string) (*TestCaseDef, error) {
	tc := &TestCaseDef{}
	decoder := json.NewDecoder(bytes.NewReader([]byte(jsonString)))
	decoder.UseNumber()
	err := decoder.Decode(tc)
	return tc, err
}

// JsonToCase implements ITestCase interface
type JsonToCase struct {
	JsonString        string
	ID                uint
	DebugTalkFilePath string
	Name              string
	Config            *TConfig
}

func (testCaseJson *JsonToCase) GetPath() string {
	return testCaseJson.DebugTalkFilePath
}

func (testCaseJson *JsonToCase) GetTestCase() (*TestCase, error) {
	return testCaseJson.ToTestCase()
}

func (testCaseJson *JsonToCase) ToTestCase() (*TestCase, error) {
	tc := &TestCaseDef{}
	var err error
	casePath := testCaseJson.JsonString
	tc, err = loadFromString(casePath)
	if err != nil {
		return nil, err
	}

	if tc.Config == nil {
		tc.Config = &TConfig{}
	}
	tc.Config.Path = testCaseJson.GetPath()
	testCase := &TestCase{
		Config: testCaseJson.Config,
	}
	if testCase.Config == nil {
		testCase.Config = tc.Config
	}
	if testCaseJson.Name != "" {
		testCase.Config.Get().Name = testCaseJson.Name
	}

	projectRootDir, err := GetProjectRootDirPath(testCaseJson.GetPath())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get project root dir")
	}

	dotEnvPath := filepath.Join(projectRootDir, ".env")
	if builtin.IsFilePathExists(dotEnvPath) {
		envVars := make(map[string]string)
		envVarsByte, err := builtin.LoadFile(dotEnvPath)
		if err != nil {
			return nil, errors.Wrap(err, "failed to load .env file")
		}
		// parse .env file
		envVarsMap, err := godotenv.Unmarshal(string(envVarsByte))
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse .env file")
		}
		for key, value := range envVarsMap {
			envVars[key] = value
		}

		if testCase.Config.Get().Environs == nil {
			testCase.Config.Get().Environs = make(map[string]string)
		}
		for key, value := range envVars {
			testCase.Config.Get().Environs[key] = value
		}
	}

	for _, step := range tc.Steps {
		if step.TestCase != nil {
			caseStr, _ := json.Marshal(step.TestCase)
			jtc := &JsonToCase{
				JsonString:        string(caseStr),
				ID:                0,
				Name:              testCase.Config.Get().Name,
				DebugTalkFilePath: testCaseJson.GetPath(),
				Config:            testCaseJson.Config,
			}
			tc, err := jtc.ToTestCase()
			if err != nil {
				return nil, err
			}
			step.TestCase = tc
			testCase.TestSteps = append(testCase.TestSteps, &StepTestCaseWithOptionalArgs{
				StepConfig: step.StepConfig,
				TestCase:   step.TestCase,
			})
		} else if step.ThinkTime != nil {
			testCase.TestSteps = append(testCase.TestSteps, &StepThinkTime{
				StepConfig: step.StepConfig,
				ThinkTime:  step.ThinkTime,
			})
		} else if step.Request != nil {
			testCase.TestSteps = append(testCase.TestSteps, &StepRequestWithOptionalArgs{
				StepRequest: &StepRequest{
					StepConfig: step.StepConfig,
					Request:    step.Request,
				},
			})
		} else if step.Transaction != nil {
			testCase.TestSteps = append(testCase.TestSteps, &StepTransaction{
				StepConfig:  step.StepConfig,
				Transaction: step.Transaction,
			})
		} else if step.Rendezvous != nil {
			testCase.TestSteps = append(testCase.TestSteps, &StepRendezvous{
				StepConfig: step.StepConfig,
				Rendezvous: step.Rendezvous,
			})
		} else if step.WebSocket != nil {
			testCase.TestSteps = append(testCase.TestSteps, &StepWebSocket{
				StepConfig: step.StepConfig,
				WebSocket:  step.WebSocket,
			})
		} else {
			log.Warn().Interface("step", step).Msg("unexpected step")
		}
	}
	return testCase, nil
}

type HttpRunnerConfig struct {
	ID uint `json:"id,omitempty"`
	*TConfig
}

type HttpRunnerTestCase struct {
	Config    *HttpRunnerConfig `json:"config"`
	TestSteps []*TStep          `json:"teststeps"`
}

func LoadTestCasesFromJSON(jsonContent []byte) ([]ITestCase, error) {
	// Try to decode as a single test case first (your data format)
	var singleCase HttpRunnerTestCase
	decoder := json.NewDecoder(bytes.NewReader(jsonContent))
	decoder.UseNumber()
	if err := decoder.Decode(&singleCase); err == nil && singleCase.Config != nil {
		// Successfully decoded as single case
		var testCases []ITestCase
		var config *TConfig
		if singleCase.Config != nil {
			config = singleCase.Config.TConfig
		} else {
			config = NewConfig("default")
		}

		tcDef := &TestCaseDef{
			Config: config,
			Steps:  singleCase.TestSteps,
		}

		testCase, err := tcDef.loadISteps()
		if err != nil {
			return nil, errors.Wrap(err, "load ISteps failed")
		}
		testCases = append(testCases, testCase)
		return testCases, nil
	}

	// If not a single case, try decoding as a list (original logic)
	var data []*HttpRunnerTestCase
	decoder = json.NewDecoder(bytes.NewReader(jsonContent))
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil {
		return nil, errors.Wrap(err, "decode json content failed")
	}

	configCache := make(map[uint]*TConfig)
	var testCases []ITestCase

	for _, caseData := range data {
		var config *TConfig
		if caseData.Config != nil {
			if caseData.Config.ID > 0 {
				if cachedConfig, ok := configCache[caseData.Config.ID]; ok {
					config = cachedConfig
				} else {
					config = caseData.Config.TConfig
					configCache[caseData.Config.ID] = config
				}
			} else {
				config = caseData.Config.TConfig
			}
		} else {
			config = NewConfig("default")
		}

		tcDef := &TestCaseDef{
			Config: config,
			Steps:  caseData.TestSteps,
		}

		// load ISteps from TSteps
		testCase, err := tcDef.loadISteps()
		if err != nil {
			return nil, errors.Wrap(err, "load ISteps failed")
		}
		testCases = append(testCases, testCase)
	}

	return testCases, nil
}
