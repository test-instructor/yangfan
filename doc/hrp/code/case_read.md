根据之前的介绍，我们了解到hrp用例的文件读取是由`loader.go`文件中的`LoadTestCases`函数完成的。该函数接收一个`ITestCase`切片作为输入，并返回一个`TestCase`切片作为输出。在函数内部，通过遍历`ITestCase`切片，将文件内容转换成可执行的测试用例数据。

# 用例类型

```go
func LoadTestCases(iTestCases ...ITestCase) ([]*TestCase, error) {
	for _, iTestCase := range iTestCases {
	}
}
```

在Go语言中，我们可以编写代码将不同文件格式（json、yaml、yml）的内容解析成hrp可以读取的测试用例格式。为此，我们需要定义一个`ITestCase`接口，其中包含两个方法：`GetPath`用于获取用例文件路径，`ToTestCase`用于将用例转换成`TestCase`对象，后者包含配置和测试步骤列表。在上一节中，我们对用例类型进行讲解，最重要的就是鸭子类型，只要实现接口`ITestCase`的所有方法，那么这个结构体就实现了`ITestCase`。如果你还没学会，可以咨询作者了解更多，也可以再看一遍，有时候多看几遍你会发现有不一样的惊喜。

```go
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
```

在处理可以直接转换成`TestCase`对象的用例类型时，我们可以直接进行格式转换，并将其添加到用例列表中。通常这些用例类型是直接用Go语言编写的，因此转换过程相对简单直接。一旦转换成功，我们便可将它们无缝地集成到用例列表中，为后续的测试执行做好准备。这样，我们可以轻松地管理不同格式的用例，并确保它们能够被hrp正确地识别和执行。

```go
		if _, ok := iTestCase.(*TestCase); ok {
			testcase, err := iTestCase.ToTestCase()
			if err != nil {
				log.Error().Err(err).Msg("failed to convert ITestCase interface to TestCase struct")
				return nil, err
			}
			testCases = append(testCases, testcase)
			continue
		}
```

## 读取文件类型的用例

如果`iTestCase`是`*TestCasePath`类型的实例，那么它表示一个文件或文件夹的路径，需要读取其中的测试用例数据。为了实现这一目的，我们使用`os.DirFS`来创建一个虚拟文件系统，并通过`fs.WalkDir`遍历指定路径下的文件和文件夹。在遍历过程中，我们会忽略隐藏文件夹（以点"."开头的文件夹）和非测试用例文件（非 `.yml`、`.yaml` 和 `.json` 后缀的文件），而将符合条件的测试用例文件转换为`*TestCase`结构，并将其添加到`testCases`切片中。这样，我们就能够方便地从指定的文件路径读取测试用例数据，并进行后续的测试执行。

```go
		// 否则，iTestCase应该是一个TestCasePath，表示文件路径或文件夹路径
		tcPath, ok := iTestCase.(*TestCasePath)
		if !ok {
			return nil, errors.New("invalid iTestCase type")
		}
		// 获取测试用例路径
		casePath := tcPath.GetPath()
		// 使用fs.WalkDir函数遍历目录结构，处理每个测试用例文件
		err := fs.WalkDir(os.DirFS(casePath), ".", func(path string, dir fs.DirEntry, e error) error {
			if dir == nil {
				// casePath是文件而不是目录
				path = casePath
			} else if dir.IsDir() && path != "." && strings.HasPrefix(path, ".") {
				// 跳过隐藏文件夹
				return fs.SkipDir
			} else {
				// casePath是目录
				path = filepath.Join(casePath, path)
			}

			// 忽略非测试用例文件
			ext := filepath.Ext(path)
			if ext != ".yml" && ext != ".yaml" && ext != ".json" {
				return nil
			}

			// 获取TestCasePath并转换为TestCase结构，然后添加到testCases切片中
			testCasePath := TestCasePath(path)
			tc, err := testCasePath.ToTestCase()
			if err != nil {
				return nil
			}
			testCases = append(testCases, tc)
			return nil
		})
```

在上述操作中，我们不仅需要通过`os.DirFS`和`fs.WalkDir`遍历文件和文件夹路径，还需要对文件内容进行读取。读取后的内容需要进一步转换成hrp可以运行的测试用例对象。为此，我们使用了`builtin.LoadFile`来获取类型为`TCase`的测试用例，并通过`testCasePath.ToTestCase()`调用进行用例转换。

这些步骤相互配合，使我们能够从指定的文件路径中获取测试用例数据，并将其转换为hrp可运行的测试用例对象。这样，我们便能够顺利地将文件中的用例内容读取并集成到整个测试用例列表中，为后续的测试执行做好准备。

## 用例转换

```go
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
// tc.ToTestCase
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
```
### 读取json、yaml格式的用例
在`LoadFile`这一步中，逻辑相对较简单。主要是根据不同的文件类型（通过后缀识别），读取文件内容后通过`JSON`转换成对应的结构体。`structObj`作为一个`any`类型参数，可以传入任意类型，在执行时会自动解析到实际参数类型。传入的类型为`TCase`，会自动将配置和测试步骤解析到`structObj`中。

这一步的优化使得我们能够轻松地根据文件后缀识别文件类型，读取文件内容，并将其转换为对应的结构体。这样，我们可以有效地将不同文件格式的用例数据转化成hrp可以理解的内部结构，为后续的用例执行和测试准备奠定了基础。

```go
type TCase struct {
	Config    *TConfig `json:"config" yaml:"config"`
	TestSteps []*TStep `json:"teststeps" yaml:"teststeps"`
}
// LoadFile loads file content with file extension and assigns to structObj
func LoadFile(path string, structObj interface{}) (err error) {
	log.Info().Str("path", path).Msg("load file")
	// 使用ReadFile函数读取文件内容
	file, err := ReadFile(path)
	if err != nil {
		return errors.Wrap(err, "read file failed")
	}
	// 移除文件开头的BOM（字节顺序标记）
	file = bytes.TrimLeft(file, "\xef\xbb\xbf")
	// 获取文件扩展名
	ext := filepath.Ext(path)
	switch ext {
	case ".json", ".har":
		// 如果文件扩展名是.json或.har，使用json.NewDecoder和Decode函数解码JSON格式文件内容
		// 并将解码后的结果赋值给structObj对象
		decoder := json.NewDecoder(bytes.NewReader(file))
		decoder.UseNumber()
		err = decoder.Decode(structObj)
		if err != nil {
			err = errors.Wrap(code.LoadJSONError, err.Error())
		}
	case ".yaml", ".yml":
		// 如果文件扩展名是.yaml或.yml，使用yaml.Unmarshal函数解析YAML格式文件内容
		// 并将解析后的结果赋值给structObj对象
		err = yaml.Unmarshal(file, structObj)
		if err != nil {
			err = errors.Wrap(code.LoadYAMLError, err.Error())
		}
	case ".env":
		// 如果文件扩展名是.env，使用parseEnvContent函数解析环境变量文件内容
		// 并将解析后的结果赋值给structObj对象
		err = parseEnvContent(file, structObj)
		if err != nil {
			err = errors.Wrap(code.LoadEnvError, err.Error())
		}
	default:
		err = code.UnsupportedFileExtension
	}
	return err
}
```

在`tc.toTestCase` 这一步，将`TCase`转换成读取到的`TestCase`测试用例。在这个过程中，我们遍历`tc.TestSteps`，在hrp中，`Step`可以是多种类型，比如`API`、Request、WebSocket等，只要它们实现了`IStep`接口，就可以作为一个类型。我们的目标是支持`MQTT`协议，因此我们需要在这里添加对应的判断，以便正确地读取我们的自定义类型。

这样，通过`tc.toTestCase`的处理，我们可以将`TCase`对象转换成hrp可以识别的`TestCase`测试用例，其中包含了我们所需的自定义类型支持，包括`MQTT`协议。这个过程为我们的测试框架添加了更多灵活性和扩展性，让我们可以更好地适应不同协议和类型的测试需求。

### toTestCase 源码解析
```go
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

// MakeCompat converts TCase compatible with Golang engine style
// 目的是使TCase结构与Golang引擎风格兼容。它通过对TCase对象及其相关的TestSteps执行某些兼容性转换来实现
// tc.MakeCompat
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
```
## 小结
现在我们对用例的读取和转换流程有了清晰的梳理。让我们再次总结一下整个流程：

1. `LoadTestCases`函数作为操作入口，接收一个文件列表作为输入参数，然后遍历该列表中的每个文件路径。
2. 对于每个文件路径，我们使用`ToTestCase`方法将其转换为一个`TestCase`测试用例对象，并将其添加到用例列表中。
3. `ToTestCase`方法内部，使用`LoadFile`函数读取文件内容，并根据文件类型通过`JSON`转换成对应的`TCase`结构体对象。
4. `TCase`对象中包含了配置和测试步骤列表。我们需要调用其`toTestCase`方法，将其转换成`TestCase`对象，以便hrp可以正确地执行测试。

代码篇幅较长，为了方便阅读和理解，我将带有注释的代码上传到项目仓库中。这样我们可以更方便地查阅详细的代码实现，更好地理解整个用例读取和转换的流程。如果还有其他需要帮助的问题，随时向我询问。