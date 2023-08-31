package hrp

import (
	"crypto/tls"
	_ "embed"
	"fmt"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/httprunner/funplugin"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"

	"github.com/test-instructor/yangfan/parsing/hrp/internal/builtin"
	"github.com/test-instructor/yangfan/parsing/hrp/internal/sdk"
	"github.com/test-instructor/yangfan/parsing/hrp/internal/version"
	"github.com/test-instructor/yangfan/parsing/hrp/pkg/uixt"
)

// Run starts to run API test with default configs.
func Run(testcases ...ITestCase) error {
	t := &testing.T{}
	return NewRunner(t).SetRequestsLogOn().Run(testcases...)
}

// NewRunner 构造一个新的运行器实例。
func NewRunner(t *testing.T) *HRPRunner {
	if t == nil {
		t = &testing.T{}
	}
	jar, _ := cookiejar.New(nil)
	return &HRPRunner{
		t:             t,
		failfast:      true, // 默认为 failfast 模式
		genHTMLReport: false,
		httpClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			Jar:     jar, // 将响应的 cookies 插入请求中
			Timeout: 120 * time.Second,
		},
		http2Client: &http.Client{
			Transport: &http2.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			Jar:     jar, // 将响应的 cookies 插入请求中
			Timeout: 120 * time.Second,
		},
		// 在此处使用默认的握手超时时间（没有超时限制），在步骤级别启用超时
		wsDialer: &websocket.Dialer{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
}

type HRPRunner struct {
	t             *testing.T                 // testing.T 的引用，用于报告测试失败和管理测试状态
	failfast      bool                       // 在第一个测试步骤失败后是否立即停止的标志。如果为 true，则在第一个失败步骤后停止测试，否则继续执行所有步骤。
	httpStatOn    bool                       // 是否启用 HTTP 统计跟踪的标志。如果为 true，则在测试执行期间收集和显示与 HTTP 请求相关的统计信息。
	requestsLogOn bool                       // 是否启用 HTTP 请求日志的标志。如果为 true，则在测试执行期间记录每个 HTTP 请求的详细信息。
	pluginLogOn   bool                       // 是否启用插件日志的标志。如果为 true，则在测试执行期间记录与插件相关的信息。
	venv          string                     // 虚拟环境路径的字符串。用于指定测试将在其中执行的虚拟环境的路径。
	saveTests     bool                       // 是否保存测试结果的标志。如果为 true，则会保存测试结果。
	genHTMLReport bool                       // 是否生成测试执行的 HTML 报告的标志。如果为 true，则会生成包含测试结果总结的 HTML 报告。
	httpClient    *http.Client               // 指向 HTTP 客户端实例的指针。用于在测试执行期间进行 HTTP 请求。
	http2Client   *http.Client               // 指向 HTTP/2 客户端实例的指针。用于在测试执行期间进行 HTTP/2 请求。
	wsDialer      *websocket.Dialer          // 指向 WebSocket 拨号器实例的指针。用于在测试执行期间建立 WebSocket 连接。
	uiClients     map[string]*uixt.DriverExt // UI 客户端的映射。用于管理具有唯一键作为标识符的 UI 测试驱动程序。
}

// SetClientTransport 配置 HTTP 客户端的传输以进行高并发负载测试。
func (r *HRPRunner) SetClientTransport(maxConns int, disableKeepAlive bool, disableCompression bool) *HRPRunner {
	log.Info().
		Int("maxConns", maxConns).
		Bool("disableKeepAlive", disableKeepAlive).
		Bool("disableCompression", disableCompression).
		Msg("[init] SetClientTransport")
	r.httpClient.Transport = &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		DialContext:         (&net.Dialer{}).DialContext,
		MaxIdleConns:        0,
		MaxIdleConnsPerHost: maxConns,
		DisableKeepAlives:   disableKeepAlive,
		DisableCompression:  disableCompression,
	}
	r.http2Client.Transport = &http2.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: disableCompression,
	}
	r.wsDialer.EnableCompression = !disableCompression
	return r
}

// SetFailfast 配置是否在步骤失败时停止运行。
func (r *HRPRunner) SetFailfast(failfast bool) *HRPRunner {
	log.Info().Bool("failfast", failfast).Msg("[init] SetFailfast")
	r.failfast = failfast
	return r
}

// SetRequestsLogOn 打开请求和响应详细信息记录。
func (r *HRPRunner) SetRequestsLogOn() *HRPRunner {
	log.Info().Msg("[init] SetRequestsLogOn")
	r.requestsLogOn = true
	return r
}

// SetHTTPStatOn 打开 HTTP 延迟统计。
func (r *HRPRunner) SetHTTPStatOn() *HRPRunner {
	log.Info().Msg("[init] SetHTTPStatOn")
	r.httpStatOn = true
	return r
}

// SetPluginLogOn 打开插件日志记录。
func (r *HRPRunner) SetPluginLogOn() *HRPRunner {
	log.Info().Msg("[init] SetPluginLogOn")
	r.pluginLogOn = true
	return r
}

// SetPython3Venv 指定 Python3 虚拟环境。
func (r *HRPRunner) SetPython3Venv(venv string) *HRPRunner {
	log.Info().Str("venv", venv).Msg("[init] SetPython3Venv")
	r.venv = venv
	return r
}

// SetProxyUrl 配置代理 URL，通常用于捕获 HTTP 数据包进行调试。
func (r *HRPRunner) SetProxyUrl(proxyUrl string) *HRPRunner {
	log.Info().Str("proxyUrl", proxyUrl).Msg("[init] SetProxyUrl")
	p, err := url.Parse(proxyUrl)
	if err != nil {
		log.Error().Err(err).Str("proxyUrl", proxyUrl).Msg("[init] invalid proxyUrl")
		return r
	}
	r.httpClient.Transport = &http.Transport{
		Proxy:           http.ProxyURL(p),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	r.wsDialer.Proxy = http.ProxyURL(p)
	return r
}

// SetTimeout 配置全局超时时间（秒）。
func (r *HRPRunner) SetTimeout(timeout time.Duration) *HRPRunner {
	log.Info().Float64("timeout(seconds)", timeout.Seconds()).Msg("[init] SetTimeout")
	r.httpClient.Timeout = timeout
	return r
}

// SetSaveTests 配置是否保存测试摘要。
func (r *HRPRunner) SetSaveTests(saveTests bool) *HRPRunner {
	log.Info().Bool("saveTests", saveTests).Msg("[init] SetSaveTests")
	r.saveTests = saveTests
	return r
}

// GenHTMLReport 配置是否生成 API 测试的 HTML 报告。
func (r *HRPRunner) GenHTMLReport() *HRPRunner {
	log.Info().Bool("genHTMLReport", true).Msg("[init] SetgenHTMLReport")
	r.genHTMLReport = true
	return r
}

// Run 开始执行一个或多个测试用例。
func (r *HRPRunner) Run(testcases ...ITestCase) error {
	// 输出HRP版本号
	log.Info().Str("hrp_version", version.VERSION).Msg("start running")
	// 初始化事件跟踪对象，用于统计执行过程中的事件
	event := sdk.EventTracking{
		Category: "RunAPITests",
		Action:   "hrp run",
	}
	go sdk.SendEvent(event)
	defer sdk.SendEvent(event.StartTiming("execution"))
	// 记录执行数据到总结对象中
	s := newOutSummary()

	// 加载所有测试用例
	testCases, err := LoadTestCases(testcases...)
	if err != nil {
		log.Error().Err(err).Msg("failed to load testcases")
		return err
	}

	// 退出所有函数插件
	defer func() {
		pluginMap.Range(func(key, value interface{}) bool {
			if plugin, ok := value.(funplugin.IPlugin); ok {
				plugin.Quit()
			}
			return true
		})
	}()

	var runErr error
	// 遍历所有用例
	for _, testcase := range testCases {
		// 每个测试用例都有自己的用例运行器，把用例转换成可迭代的用例
		caseRunner, err := r.NewCaseRunner(testcase)
		if err != nil {
			log.Error().Err(err).Msg("[Run] init case runner failed")
			return err
		}

		// 释放UI驱动程序的会话
		defer func() {
			for _, client := range r.uiClients {
				client.Driver.DeleteSession()
			}
		}()

		// 使用参数迭代器执行用例多次，每次执行有不同的参数
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
	}

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

	return runErr
}

// NewCaseRunner 创建一个新的用例运行器（CaseRunner）用于指定的测试用例（testcase）。
// 每个测试用例都有自己的用例运行器。
func (r *HRPRunner) NewCaseRunner(testcase *TestCase) (*CaseRunner, error) {
	// 创建一个新的用例运行器（CaseRunner）对象
	caseRunner := &CaseRunner{
		testCase:  testcase,    // 设置用例运行器的测试用例字段
		hrpRunner: r,           // 设置用例运行器的HRPRunner字段
		parser:    newParser(), // 创建并初始化一个新的解析器（parser）对象
	}

	// 初始化函数插件
	plugin, err := initPlugin(testcase.Config.Path, r.venv, r.pluginLogOn)
	if err != nil {
		return nil, errors.Wrap(err, "init plugin failed")
	}
	if plugin != nil {
		caseRunner.parser.plugin = plugin                // 将解析器插件设置到用例运行器的解析器对象中
		caseRunner.rootDir = filepath.Dir(plugin.Path()) // 设置用例运行器的根目录（rootDir），即插件所在的目录
	}

	// 解析测试用例的配置
	if err := caseRunner.parseConfig(); err != nil {
		return nil, errors.Wrap(err, "parse testcase config failed")
	}

	// 设置测试用例的超时时间（timeout）（单位：秒）
	if testcase.Config.Timeout != 0 {
		timeout := time.Duration(testcase.Config.Timeout*1000) * time.Millisecond
		r.SetTimeout(timeout)
	}

	// 将插件信息加载到测试用例的配置中
	if plugin != nil {
		pluginPath, _ := locatePlugin(testcase.Config.Path)
		if caseRunner.parsedConfig.PluginSetting == nil {
			pluginContent, err := builtin.ReadFile(pluginPath)
			if err != nil {
				return nil, err
			}
			tp := strings.Split(plugin.Path(), ".")
			caseRunner.parsedConfig.PluginSetting = &PluginConfig{
				Path:    pluginPath,
				Content: pluginContent,
				Type:    tp[len(tp)-1],
			}
		}
	}

	return caseRunner, nil // 返回用例运行器对象和可能出现的错误
}

// CaseRunner 结构体用于执行单个测试用例的运行器
type CaseRunner struct {
	testCase  *TestCase
	hrpRunner *HRPRunner
	parser    *Parser

	parsedConfig       *TConfig
	parametersIterator *ParametersIterator
	rootDir            string // 项目根目录
}

// parseConfig 方法解析测试用例的配置，将解析结果存储到 parsedConfig 中
func (r *CaseRunner) parseConfig() error {
	cfg := r.testCase.Config

	r.parsedConfig = &TConfig{} // 创建一个新的 TConfig 对象
	// 深拷贝配置以避免数据竞争
	if err := copier.Copy(r.parsedConfig, cfg); err != nil {
		log.Error().Err(err).Msg("copy testcase config failed")
		return err
	}

	// 解析配置变量
	parsedVariables, err := r.parser.ParseVariables(cfg.Variables)
	if err != nil {
		log.Error().Interface("variables", cfg.Variables).Err(err).Msg("parse config variables failed")
		return err
	}
	r.parsedConfig.Variables = parsedVariables

	// 解析配置名称
	parsedName, err := r.parser.ParseString(cfg.Name, parsedVariables)
	if err != nil {
		return errors.Wrap(err, "parse config name failed")
	}
	r.parsedConfig.Name = convertString(parsedName)

	// 解析配置基础URL
	parsedBaseURL, err := r.parser.ParseString(cfg.BaseURL, parsedVariables)
	if err != nil {
		return errors.Wrap(err, "parse config base url failed")
	}
	r.parsedConfig.BaseURL = convertString(parsedBaseURL)

	// 合并配置环境变量和基础URL
	// 优先级：env base_url > base_url
	if cfg.Environs != nil {
		r.parsedConfig.Environs = cfg.Environs
	} else {
		r.parsedConfig.Environs = make(map[string]string)
	}
	if value, ok := r.parsedConfig.Environs["base_url"]; !ok || value == "" {
		if r.parsedConfig.BaseURL != "" {
			r.parsedConfig.Environs["base_url"] = r.parsedConfig.BaseURL
		}
	}

	// 合并配置变量和环境变量
	// 优先级：env > config variables
	for k, v := range r.parsedConfig.Environs {
		r.parsedConfig.Variables[k] = v
	}

	// 确保思考时间配置的正确性
	r.parsedConfig.ThinkTimeSetting.checkThinkTime()

	// 确保WebSocket配置的正确性
	r.parsedConfig.WebSocketSetting.checkWebSocket()

	// 解析测试用例配置参数
	parametersIterator, err := r.parser.initParametersIterator(r.parsedConfig)
	if err != nil {
		log.Error().Err(err).
			Interface("parameters", r.parsedConfig.Parameters).
			Interface("parametersSetting", r.parsedConfig.ParametersSetting).
			Msg("parse config parameters failed")
		return errors.Wrap(err, "parse testcase config parameters failed")
	}
	r.parametersIterator = parametersIterator

	// 初始化iOS/Android客户端
	if r.hrpRunner.uiClients == nil {
		r.hrpRunner.uiClients = make(map[string]*uixt.DriverExt)
	}
	for _, iosDeviceConfig := range r.parsedConfig.IOS {
		// 解析iOS设备的UDID
		if iosDeviceConfig.UDID != "" {
			udid, err := r.parser.ParseString(iosDeviceConfig.UDID, parsedVariables)
			if err != nil {
				return errors.Wrap(err, "failed to parse ios device udid")
			}
			iosDeviceConfig.UDID = udid.(string)
		}
		// 创建iOS设备驱动
		device, err := uixt.NewIOSDevice(uixt.GetIOSDeviceOptions(iosDeviceConfig)...)
		if err != nil {
			return errors.Wrap(err, "init iOS device failed")
		}
		// 创建iOS设备UI驱动客户端
		client, err := device.NewDriver(nil)
		if err != nil {
			return errors.Wrap(err, "init iOS WDA client failed")
		}
		r.hrpRunner.uiClients[device.UDID] = client
	}
	for _, androidDeviceConfig := range r.parsedConfig.Android {
		// 解析Android设备的序列号
		if androidDeviceConfig.SerialNumber != "" {
			sn, err := r.parser.ParseString(androidDeviceConfig.SerialNumber, parsedVariables)
			if err != nil {
				return errors.Wrap(err, "failed to parse android device serial")
			}
			androidDeviceConfig.SerialNumber = sn.(string)
		}
		// 创建Android设备驱动
		device, err := uixt.NewAndroidDevice(uixt.GetAndroidDeviceOptions(androidDeviceConfig)...)
		if err != nil {
			return errors.Wrap(err, "init Android device failed")
		}
		// 创建Android设备UIAutomator客户端
		client, err := device.NewDriver(nil)
		if err != nil {
			return errors.Wrap(err, "init Android UIAutomator client failed")
		}
		r.hrpRunner.uiClients[device.SerialNumber] = client
	}

	return nil
}

// NewSession 方法用于创建一个新的会话运行器（SessionRunner）
func (r *CaseRunner) NewSession() *SessionRunner {
	sessionRunner := &SessionRunner{
		caseRunner: r,
	}
	sessionRunner.resetSession() // 重置会话状态
	return sessionRunner
}

// SessionRunner 用于运行测试用例及其步骤。
// 每个测试用例都有自己的 SessionRunner 实例，并共享会话变量。
type SessionRunner struct {
	caseRunner        *CaseRunner
	sessionVariables  map[string]interface{}
	transactions      map[string]map[transactionType]time.Time
	startTime         time.Time                  // 记录测试用例的开始时间
	summary           *TestCaseSummary           // 记录测试用例的摘要信息
	wsConnMap         map[string]*websocket.Conn // 保存所有 WebSocket 连接
	inheritWsConnMap  map[string]*websocket.Conn // 继承的 WebSocket 连接
	pongResponseChan  chan string                // 通道用于接收 Pong 响应消息
	closeResponseChan chan *wsCloseRespObject    // 通道用于接收关闭响应消息
}

// resetSession 重置会话状态，用于初始化一个新的会话
func (r *SessionRunner) resetSession() {
	log.Info().Msg("重置会话运行器")
	r.sessionVariables = make(map[string]interface{})
	r.transactions = make(map[string]map[transactionType]time.Time)
	r.startTime = time.Now()
	r.summary = newSummary()
	r.wsConnMap = make(map[string]*websocket.Conn)
	r.inheritWsConnMap = make(map[string]*websocket.Conn)
	r.pongResponseChan = make(chan string, 1)
	r.closeResponseChan = make(chan *wsCloseRespObject, 1)
}

// inheritConnection 从另一个 SessionRunner 实例继承 WebSocket 连接
func (r *SessionRunner) inheritConnection(src *SessionRunner) {
	log.Info().Msg("继承会话运行器")
	r.inheritWsConnMap = make(map[string]*websocket.Conn, len(src.wsConnMap)+len(src.inheritWsConnMap))
	for k, v := range src.wsConnMap {
		r.inheritWsConnMap[k] = v
	}
	for k, v := range src.inheritWsConnMap {
		r.inheritWsConnMap[k] = v
	}
}

// Start 按顺序运行测试步骤。
// givenVars 用于数据驱动
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

// ParseStepVariables 将步骤变量与配置变量和会话变量合并
func (r *SessionRunner) ParseStepVariables(stepVariables map[string]interface{}) (map[string]interface{}, error) {
	// 覆盖变量
	// 步骤变量 > 会话变量（从之前的步骤中提取的变量）
	overrideVars := mergeVariables(stepVariables, r.sessionVariables)
	// 步骤变量 > 测试用例配置变量
	overrideVars = mergeVariables(overrideVars, r.caseRunner.parsedConfig.Variables)

	// 解析步骤变量
	parsedVariables, err := r.caseRunner.parser.ParseVariables(overrideVars)
	if err != nil {
		log.Error().Interface("variables", r.caseRunner.parsedConfig.Variables).
			Err(err).Msg("解析步骤变量失败")
		return nil, errors.Wrap(err, "解析步骤变量失败")
	}
	return parsedVariables, nil
}

// InitWithParameters 使用给定参数更新会话变量。
// 这用于数据驱动
func (r *SessionRunner) InitWithParameters(parameters map[string]interface{}) {
	if len(parameters) == 0 {
		return
	}

	log.Info().Interface("parameters", parameters).Msg("更新会话变量")
	for k, v := range parameters {
		r.sessionVariables[k] = v
	}
}

// GetSummary 获取测试用例的摘要信息
func (r *SessionRunner) GetSummary() (*TestCaseSummary, error) {
	caseSummary := r.summary
	caseSummary.Name = r.caseRunner.parsedConfig.Name
	caseSummary.Time.StartAt = r.startTime
	caseSummary.Time.Duration = time.Since(r.startTime).Seconds()
	exportVars := make(map[string]interface{})
	for _, value := range r.caseRunner.parsedConfig.Export {
		exportVars[value] = r.sessionVariables[value]
	}
	caseSummary.InOut.ExportVars = exportVars
	caseSummary.InOut.ConfigVars = r.caseRunner.parsedConfig.Variables

	for uuid, client := range r.caseRunner.hrpRunner.uiClients {
		// 将 WDA/UIA 日志添加到摘要信息中
		logs := map[string]interface{}{
			"uuid": uuid,
		}

		if client.Device.LogEnabled() {
			log, err := client.Driver.StopCaptureLog()
			if err != nil {
				return caseSummary, err
			}
			logs["content"] = log
		}

		// 停止性能监控
		logs["performance"] = client.Device.StopPerf()
		logs["pcap"] = client.Device.StopPcap()

		caseSummary.Logs = append(caseSummary.Logs, logs)
	}

	return caseSummary, nil
}

// updateSummary 更新 StepResult 的摘要信息。
func (r *SessionRunner) updateSummary(stepResult *StepResult) {
	switch stepResult.StepType {
	case stepTypeTestCase:
		// 记录测试用例步骤的请求
		if records, ok := stepResult.Data.([]*StepResult); ok {
			for _, result := range records {
				r.addSingleStepResult(result)
			}
		} else {
			r.addSingleStepResult(stepResult)
		}
	default:
		r.addSingleStepResult(stepResult)
	}
}

// addSingleStepResult 更新单个步骤的摘要信息
func (r *SessionRunner) addSingleStepResult(stepResult *StepResult) {
	// 更新摘要信息
	r.summary.Records = append(r.summary.Records, stepResult)
	r.summary.Stat.Total += 1
	if stepResult.Success {
		r.summary.Stat.Successes += 1
	} else {
		r.summary.Stat.Failures += 1
		// 将摘要信息结果更新为失败
		r.summary.Success = false
	}
}

// releaseResources 释放会话运行器使用的资源
func (r *SessionRunner) releaseResources() {
	// 关闭 WebSocket 连接
	for _, wsConn := range r.wsConnMap {
		if wsConn != nil {
			log.Info().Str("testcase", r.caseRunner.testCase.Config.Name).Msg("WebSocket 连接已断开")
			err := wsConn.Close()
			if err != nil {
				log.Error().Err(err).Msg("WebSocket 连接关闭失败")
			}
		}
	}
}

// getWsClient 获取指定 URL 的 WebSocket 客户端连接
func (r *SessionRunner) getWsClient(url string) *websocket.Conn {
	if client, ok := r.wsConnMap[url]; ok {
		return client
	}

	if client, ok := r.inheritWsConnMap[url]; ok {
		return client
	}

	return nil
}
