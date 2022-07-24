# 测试平台接入HttpRunner V4

## 1、如何接入

###  1.1 v4 版本支持`go 形态`、`json`、`yaml`等多种数据运行，那么接入可以从这几方面入手

1. `go 形态` 需要将测试用例转成`go`代码，实现起来比较麻烦，所以不合适
2. `json`、`yaml`需要将测试用例转成文件，再去执行、收集测试报告，效率比较低，在有其他条件的情况下，尽量不考虑

### 1.2 考虑直接用代码调用

> 既然支持`json`格式，那么是否可以直接将测试用例转成`json`，然后直接塞入`ITestCase`中，这时候看稍微看了一下源码，好像也行不通，这时候看一下官方文档，发现一缕曙光，`TestSteps：有序步骤的集合；采用了 go interface 的设计理念，支持进行任意协议和测试类型的拓展；步骤内容统一在 Run 方法中进行实现。`最后通过这部分`go interface`的设计理念实现接入





## 2、接入流程

1. 实现`ITestCase`，通过源码发现`ITestCase`接口实现了`GetPath`、`ToTestCase`两个方法，那么只需要写一个`struct`实现`GetPath`、`ToTestCase`两个方法就可以
2. 增加`id`字段，关联至已有的已有数据，方便统计用例运行情况
3. 获取测试报告，v4 报告通类型为`Summary`，创建一个相同的结构体用来保存测试报告
4. 函数驱动，由于文件读写问题，需要每个`用例/任务/接口`+`id`放在一个目录去运行debugtalk，否则没办法同时运行多个`用例/任务/接口`,会出现1用例直接结束，停止debugtalk后导致异常

## 3、代码实现
```go
// 代码调用，以下只列出调用测试用例的关键代码
func RunCase(apiCaseID request.RunCaseReq) (reports interfacecase.ApiReport, err error) {
    toTestCase := ToTestCase{Config: apiConfig, TestSteps: apiCases.TStep}
    caseJson, _ := json.Marshal(toTestCase)
    
    tc := &hrp.TestCaseJson{
        JsonString:        string(caseJson),
        ID:                apiCaseID.CaseID,
        DebugTalkFilePath: debugTalkFilePath,
    }
    reports, errs := hrp.NewRunner(t).
    SetHTTPStatOn().
    SetFailfast(false).
    RunJsons(testCase)
    if errs != nil {
        t.Fatalf("run testcase error: %v", err)
    }
    return
}
```

```go
// 路径拼接
func tmpls(relativePath, debugTalkFileName string) string {
	return filepath.Join(debugTalkFileName, relativePath)
}
```

```go
// 读取测试用例，并转换成httprunner格式
type TestCaseJson struct {
	JsonString        string
	ID                uint
	DebugTalkFilePath string
}

func (testCaseJson *TestCaseJson) GetPath() string {
	return testCaseJson.DebugTalkFilePath
}

func (testCaseJson *TestCaseJson) ToTestCase() (*TestCase, error) {
	tc := &TCase{}
	var err error
	casePath := testCaseJson.JsonString
	tc, err = loadFromString(casePath)
	if err != nil {
		return nil, err
	}

	err = tc.MakeCompat()
	if err != nil {
		return nil, err
	}

	tc.Config.Path = testCaseJson.GetPath()

	testCase := &TestCase{
		ID:     testCaseJson.ID,
		Config: tc.Config,
	}

	projectRootDir, err := GetProjectRootDirPath(testCaseJson.GetPath())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get project root dir")
	}

	// load .env file
	dotEnvPath := filepath.Join(projectRootDir, ".env")
	if builtin.IsFilePathExists(dotEnvPath) {
		envVars := make(map[string]string)
		err = builtin.LoadFile(dotEnvPath, envVars)
		if err != nil {
			return nil, errors.Wrap(err, "failed to load .env file")
		}

		// override testcase config env with variables loaded from .env file
		// priority: .env file > testcase config env
		if testCase.Config.Environs == nil {
			testCase.Config.Environs = make(map[string]string)
		}
		for key, value := range envVars {
			testCase.Config.Environs[key] = value
		}
	}

	for _, step := range tc.TestSteps {
		step.ParntID = step.ID
		step.ID = 0
		if step.API != nil {
			apiPath, ok := step.API.(string)
			if !ok {
				return nil, fmt.Errorf("referenced api path should be string, got %v", step.API)
			}
			path := filepath.Join(projectRootDir, apiPath)
			if !builtin.IsFilePathExists(path) {
				return nil, errors.New("referenced api file not found: " + path)
			}

			refAPI := APIPath(path)
			apiContent, err := refAPI.ToAPI()
			if err != nil {
				return nil, err
			}
			step.API = apiContent

			testCase.TestSteps = append(testCase.TestSteps, &StepAPIWithOptionalArgs{
				step: step,
			})
		} else if step.TestCase != nil {
			casePath, ok := step.TestCase.(string)
			if !ok {
				return nil, fmt.Errorf("referenced testcase path should be string, got %v", step.TestCase)
			}
			path := filepath.Join(projectRootDir, casePath)
			if !builtin.IsFilePathExists(path) {
				return nil, errors.New("referenced testcase file not found: " + path)
			}

			refTestCase := TestCasePath(path)
			tc, err := refTestCase.ToTestCase()
			if err != nil {
				return nil, err
			}
			step.TestCase = tc
			testCase.TestSteps = append(testCase.TestSteps, &StepTestCaseWithOptionalArgs{
				step: step,
			})
		} else if step.ThinkTime != nil {
			testCase.TestSteps = append(testCase.TestSteps, &StepThinkTime{
				step: step,
			})
		} else if step.Request != nil {
			testCase.TestSteps = append(testCase.TestSteps, &StepRequestWithOptionalArgs{
				step: step,
			})
		} else if step.Transaction != nil {
			testCase.TestSteps = append(testCase.TestSteps, &StepTransaction{
				step: step,
			})
		} else if step.Rendezvous != nil {
			testCase.TestSteps = append(testCase.TestSteps, &StepRendezvous{
				step: step,
			})
		} else if step.WebSocket != nil {
			testCase.TestSteps = append(testCase.TestSteps, &StepWebSocket{
				step: step,
			})
		} else {
			log.Warn().Interface("step", step).Msg("[convertTestCase] unexpected step")
		}
	}
	return testCase, nil
}

func loadFromString(jsonString string) (*TCase, error) {
	tc := &TCase{}
	decoder := json.NewDecoder(bytes.NewReader([]byte(jsonString)))
	decoder.UseNumber()
	err := decoder.Decode(tc)
	return tc, err
}
```

```go
// 在HRPRunner下增加运行测试用例，并输出测试报告
func (r *HRPRunner) RunJsons(testcases ...ITestCase) (interfacecase.ApiReport, error) {
	event := sdk.EventTracking{
		Category: "RunAPITests",
		Action:   "hrp run",
	}
	// report start event
	go func() {
		err := sdk.SendEvent(event)
		if err != nil {

		}
	}()
	// report execution timing event
	defer func(e sdk.IEvent) {
		err := sdk.SendEvent(e)
		if err != nil {

		}
	}(event.StartTiming("execution"))
	// record execution data to summary
	s := newOutSummary()

	// load all testcases
	testCases, err := LoadTestCases(testcases...)
	if err != nil {
		log.Error().Err(err).Msg("run json failed to load testcases")
		return interfacecase.ApiReport{}, err
	}

	// run testcase one by one
	for _, testcase := range testCases {
		sessionRunner, err := r.NewSessionRunner(testcase)
		if err != nil {
			log.Error().Err(err).Msg("[Run] init session runner failed")
			return interfacecase.ApiReport{}, err
		}
		defer func() {
			if sessionRunner.parser.plugin != nil {
				sessionRunner.parser.plugin.Quit()
			}
		}()

		for it := sessionRunner.parametersIterator; it.HasNext(); {
			if err = sessionRunner.Start(it.Next()); err != nil {
				log.Error().Err(err).Msg("[Run] run testcase failed")
				return interfacecase.ApiReport{}, err
			}
			caseSummary := sessionRunner.GetSummary()
			caseSummary.CaseID = testcase.ID
			s.appendCaseSummary(caseSummary)
		}
	}
	s.Time.Duration = time.Since(s.Time.StartAt).Seconds()

	// save summary
	if r.saveTests {
		err := s.genSummary()
		if err != nil {
			return interfacecase.ApiReport{}, err
		}
	}

	// generate HTML report
	if r.genHTMLReport {
		err := s.genHTMLReport()
		if err != nil {
			return interfacecase.ApiReport{}, err
		}
	}
	sj, _ := json.Marshal(s)
	global.GVA_LOG.Debug("测试报告json格式")
	global.GVA_LOG.Debug("\n" + string(sj))
	var reportsStruct interfacecase.ApiReport
	err = json.Unmarshal(sj, &reportsStruct)
	return reportsStruct, nil
}
```

```go
// python函数驱动方法
func BuildHashicorpPyPlugin(debugTalkByte []byte, debugTalkFilePath string) {
	log.Info().Msg("[init] prepare hashicorp python plugin")
	//src, _ := ioutil.ReadFile(tmpls("plugin/debugtalk.py"))
	err := ioutil.WriteFile(tmpls("debugtalk.py", debugTalkFilePath), debugTalkByte, 0o644)
	if err != nil {
		log.Error().Err(err).Msg("copy hashicorp python plugin failed")
		os.Exit(1)
	}
}

func RemoveHashicorpPyPlugin(debugTalkFilePath string) {
	log.Info().Msg("[teardown] remove hashicorp python plugin")
	// on v4.1^, running case will generate .debugtalk_gen.py used by python plugin
	os.Remove(tmpls(PluginPySourceFile, debugTalkFilePath))
	os.Remove(tmpls(PluginPySourceGenFile, debugTalkFilePath))
}
```

```go
// 增加id ，只列出修改部分，根据之前git提交记录，不确定是否齐全，有问题欢迎联系
type TStep struct {
    ID            uint                   `json:"ID"`
    ParntID       uint                   `json:"parntID"`
}

type Summary struct {
    CaseID   uint               `json:"caseID"`
}

type TestCase struct {
    ID        uint
}
func (r *SessionRunner) Start(givenVars map[string]interface{}) error {
    stepResult, err := step.Run(r)  // 在这一行下增加下面一行
    stepResult.ParntID = step.Struct().ParntID
}
```
