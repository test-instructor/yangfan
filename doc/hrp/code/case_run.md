
在前面的章节中，我们了解了每个测试用例的执行都是由一个运行对象`CaseRunner`来管理的。以下是`CaseRunner`的主要执行步骤：

## 用例执行流程
1. **判断接口可用性：** 通过调用`HasNext`方法判断是否还有待执行的接口。
2. **创建运行会话：** 通过`NewSession`方法创建一个运行会话，用于执行当前接口。
3. **启动接口运行：** 通过`Start`方法开始运行当前接口，这包括了解析和处理请求参数、执行`SetupHooks`、发起实际请求、解析响应、执行`TeardownHooks`等步骤。
4. **获取运行结果：** 通过`GetSummary`方法获取接口运行结果，并将其汇总组装到测试报告列表中。
5. **退出机制：** 在接口运行失败的情况下，根据设置的中断条件，可以终止整个用例的运行。

通过以上流程，`CaseRunner`保证了每个测试用例的顺序执行，将每个接口的运行结果汇总到测试报告中，并提供了灵活的退出机制，以满足不同运行需求。

```go
for it := caseRunner.parametersIterator; it.HasNext(); {
			// 每次运行都有自己的会话运行器
			sessionRunner := caseRunner.NewSession()
			err1 := sessionRunner.Start(it.Next())
			if err1 != nil {
				log.Error().Err(err1).Msg("[Run] run testcase failed")
				runErr = err1
			}
			// 获取用例的执行结果
			caseSummary, err2 := sessionRunner.GetSummary()
			s.appendCaseSummary(caseSummary)
			if err2 != nil {
				log.Error().Err(err2).Msg("[Run] get summary failed")
				if err1 != nil {
					runErr = errors.Wrap(err1, err2.Error())
				} else {
					runErr = err2
				}
			}

			// 如果发生错误且设置了failfast标志，则终止运行
			// 实际操作一般都会执行所有用例后，最后获取用例执行情况
			if runErr != nil && r.failfast {
				break
			}
		}
```



`SessionRunner` 是一个结构体，它用于管理测试用例执行的会话过程，包含了多个成员变量和方法。以下是 `SessionRunner` 结构体中的成员变量和方法，以及它们的大致作用，由于代码篇幅太长，如果需要看具体方法的源码解析，可以到项目查看：

## CaseRunner 变量解析：

1. `caseRunner`: 指向执行当前会话的 `CaseRunner` 实例。
2. `sessionVariables`: 保存会话中的变量和数据，可以在不同步骤之间共享。
3. `transactions`: 记录事务的时间信息，以事务名为键，存储事务类型和时间信息。
4. `startTime`: 记录会话开始的时间。
5. `summary`: 保存会话的测试结果和摘要信息。
6. `wsConnMap`: 保存所有` WebSocket` 连接的映射。
7. `inheritWsConnMap`: 保存继承的 `WebSocke`t 连接的映射。
8. `pongResponseChan`: 用于接收` WebSocket` 的 `Pong` 响应消息的通道。
9. `closeResponseChan`: 用于接收 `WebSocket` 的关闭响应消息的通道。

## CaseRunner 方法解析：

1. `resetSession()`: 重置会话，清空会话变量、事务信息等，用于开始新的测试会话。
2. `inheritConnection(src *SessionRunner)`: 继承另一个 `SessionRunner` 的 `WebSocket` 连接。
3. `Start(givenVars map[string]interface{}) error`: 开始执行会话的测试步骤，按顺序运行每个步骤。
4. `ParseStepVariables(stepVariables map[string]interface{}) (map[string]interface{}, error)`: 解析步骤中的变量，合并并解析会话、步骤和配置中的变量。
5. `InitWithParameters(parameters map[string]interface{})`: 使用给定的参数更新会话变量，用于数据驱动。
6. `GetSummary() (*TestCaseSummary, error)`: 获取会话的测试结果摘要，包括执行时间、导出变量等。
7. `updateSummary(stepResult *StepResult)`: 更新会话摘要信息，根据步骤执行结果进行统计。
8. `addSingleStepResult(stepResult *StepResult)`: 添加单个步骤的执行结果到会话摘要中。
9. `releaseResources()`: 释放会话使用的资源，关闭 `WebSocket` 连接。
10. `getWsClient(url string) *websocket.Conn`: 获取指定 URL 对应的 `WebSocket` 连接。

总体而言，`SessionRunner` 负责管理测试会话的整个执行过程，包括变量处理、步骤运行、结果统计等，保证了测试执行的可控性和数据共享。
## 部分源码解析
```go
// SessionRunner 用于运行测试用例及其步骤。
// 每个测试用例都有自己的 SessionRunner 实例，并共享会话变量。
type SessionRunner struct {
	caseRunner         *CaseRunner
	sessionVariables  map[string]interface{}
	transactions      map[string]map[transactionType]time.Time
	startTime         time.Time
	summary           *TestCaseSummary
	wsConnMap         map[string]*websocket.Conn
	inheritWsConnMap  map[string]*websocket.Conn
	pongResponseChan  chan string
	closeResponseChan chan *wsCloseRespObject
}

// resetSession 重置会话状态，用于初始化一个新的会话
func (r *SessionRunner) resetSession() {
	// 初始化会话变量、事务、时间等
}

// inheritConnection 从另一个 SessionRunner 实例继承 WebSocket 连接
func (r *SessionRunner) inheritConnection(src *SessionRunner) {
	// 继承 WebSocket 连接
}

// Start 顺序执行测试用例的测试步骤
func (r *SessionRunner) Start(givenVars map[string]interface{}) error {
	// 获取测试用例配置信息
	config := r.caseRunner.testCase.Config
	log.Info().Str("testcase", config.Name).Msg("运行测试用例开始")

	// 使用给定的变量更新配置变量
	r.InitWithParameters(givenVars)

	defer func() {
		// 在所有步骤完成或出现快速失败时释放会话资源
		r.releaseResources()
	}()

	// 按顺序运行每个步骤
	for _, step := range r.caseRunner.testCase.TestSteps {
		// TODO: 解析步骤结构
		// 解析步骤名称
		parsedName, err := r.caseRunner.parser.ParseString(step.Name(), r.sessionVariables)
		if err != nil {
			parsedName = step.Name()
		}
		stepName := convertString(parsedName)
		log.Info().Str("step", stepName).
			Str("type", string(step.Type())).Msg("运行步骤开始")

		// 获取步骤运行次数
		loopTimes := step.Struct().Loops
		if loopTimes < 0 {
			log.Warn().Int("loops", loopTimes).Msg("循环次数应为正数，设置为 1")
			loopTimes = 1
		} else if loopTimes == 0 {
			loopTimes = 1
		} else if loopTimes > 1 {
			log.Info().Int("loops", loopTimes).Msg("按指定的循环次数运行步骤")
		}

		// 按指定的循环次数运行步骤
		var stepResult *StepResult
		for i := 1; i <= loopTimes; i++ {
			var loopIndex string
			if loopTimes > 1 {
				log.Info().Int("index", i).Msg("在循环中开始运行步骤")
				loopIndex = fmt.Sprintf("_loop_%d", i)
			}

			// 运行步骤
			stepResult, err = step.Run(r)
			stepResult.Name = stepName + loopIndex

			r.updateSummary(stepResult)
		}

		// 更新提取的变量
		for k, v := range stepResult.ExportVars {
			r.sessionVariables[k] = v
		}

		if err == nil {
			log.Info().Str("step", stepResult.Name).
				Str("type", string(stepResult.StepType)).
				Bool("success", true).
				Interface("exportVars", stepResult.ExportVars).
				Msg("运行步骤结束")
			continue
		}

		// 运行失败
		log.Error().Err(err).Str("step", stepResult.Name).
			Str("type", string(stepResult.StepType)).
			Bool("success", false).
			Msg("运行步骤结束")

		// 检查是否设置了 failfast
		if r.caseRunner.hrpRunner.failfast {
			return errors.Wrap(err, "由于设置了 failfast，中止运行")
		}
	}

	log.Info().Str("testcase", config.Name).Msg("运行测试用例结束")
	return nil
}

// ParseStepVariables 将步骤变量与配置变量和会话变量合并并解析
func (r *SessionRunner) ParseStepVariables(stepVariables map[string]interface{}) (map[string]interface{}, error) {
	// 解析步骤变量并与其他变量合并
}

// InitWithParameters 使用给定参数更新会话变量，用于数据驱动
func (r *SessionRunner) InitWithParameters(parameters map[string]interface{}) {
	// 使用给定参数更新会话变量
}

// GetSummary 获取测试用例的摘要信息
func (r *SessionRunner) GetSummary() (*TestCaseSummary, error) {
	// 获取测试用例的摘要信息，包括步骤结果、配置变量、日志等
}

// updateSummary 更新测试摘要中的步骤结果
func (r *SessionRunner) updateSummary(stepResult *StepResult) {
	// 更新测试摘要中的步骤结果
}

// addSingleStepResult 将单个步骤结果添加到测试摘要中
func (r *SessionRunner) addSingleStepResult(stepResult *StepResult) {
	// 将单个步骤结果添加到测试摘要中
}

// releaseResources 释放会话运行器使用的资源
func (r *SessionRunner) releaseResources() {
	// 关闭 WebSocket 连接
}

// getWsClient 获取指定 URL 对应的 WebSocket 连接
func (r *SessionRunner) getWsClient(url string) *websocket.Conn {
	// 获取 WebSocket 连接
}
```

通过`Start`方法，测试用例会按照预定的步骤（如`http`、`websocket`、`rpc`、`ui`等）有序地执行，每个步骤的测试报告都会逐一汇总到总体测试报告中。同时，参数提取功能允许将关键结果保存在配置中，以便后续接口调用时使用。若循环次数设置，每个接口将按照指定次数重复运行。这一流程实现了有序、自动化的用例执行与结果记录。

运行测试用例后，我们可以根据运行时设置的参数来对测试报告进行二次开发：

1. **增加其他类型的测试报告：** 如果现有的测试报告格式不满足需求，我们可以根据实际情况，自定义并添加其他类型的测试报告。通过在`CaseRunner`运行完成后，将运行结果整理成不同格式的报告，如Markdown。
2. **收集测试报告内容并存入数据库：** 在运行后，我们可以将测试报告的内容整理并存入数据库，以便后续查询和分析。通过将测试结果与数据库相结合，可以实现更深入的数据分析和业务探索。
3. **根据执行情况进行消息发送：** 在测试用例执行完成后，根据执行情况，我们可以通过消息发送的方式，将运行结果通知相关人员。这可以是通过邮件、飞书、钉钉、企业微信等方式进行通知，以便团队成员及时了解测试结果。

通过这些二次开发操作，我们可以根据实际需求定制化测试报告的内容和格式，将测试结果与其他业务流程结合起来，从而更好地支持项目的测试和管理工作。

```go
	// 计算执行时间
	s.Time.Duration = time.Since(s.Time.StartAt).Seconds()

	// 保存测试报告
	if r.saveTests {
		err := s.genSummary()
		if err != nil {
			return err
		}
	}

	// 生成HTML报告
	if r.genHTMLReport {
		err := s.genHTMLReport()
		if err != nil {
			return err
		}
	}
```

