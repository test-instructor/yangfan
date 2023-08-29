package hrp

import (
	"crypto/tls"
	_ "embed"
	"encoding/json"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/textproto"
	"net/url"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/httprunner/funplugin"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/interfacecase/hrp"
	"go.uber.org/zap"
	"golang.org/x/net/http2"

	"github.com/test-instructor/yangfan/hrp/internal/builtin"
	"github.com/test-instructor/yangfan/hrp/internal/sdk"
	"github.com/test-instructor/yangfan/hrp/pkg/uixt"
)

var retrySleepTimer = 3 * time.Second

// Run starts to run API test with default configs.
func Run(testcases ...ITestCase) error {
	t := &testing.T{}
	return NewRunner(t).SetRequestsLogOn().Run(testcases...)
}

// NewRunner constructs a new runner instance.
func NewRunner(t *testing.T) *HRPRunner {
	if t == nil {
		t = &testing.T{}
	}
	jar, _ := cookiejar.New(nil)
	return &HRPRunner{
		t:             t,
		failfast:      true, // default to failfast
		genHTMLReport: false,
		httpClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			Jar:     jar, // insert response cookies into request
			Timeout: 120 * time.Second,
		},
		http2Client: &http.Client{
			Transport: &http2.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			Jar:     jar, // insert response cookies into request
			Timeout: 120 * time.Second,
		},
		// use default handshake timeout (no timeout limit) here, enable timeout at step level
		wsDialer: &websocket.Dialer{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
}

type HRPRunner struct {
	t             *testing.T
	failfast      bool
	httpStatOn    bool
	requestsLogOn bool
	pluginLogOn   bool
	venv          string
	saveTests     bool
	genHTMLReport bool
	httpClient    *http.Client
	http2Client   *http.Client
	wsDialer      *websocket.Dialer
	uiClients     map[string]*uixt.DriverExt // UI automation clients for iOS and Android, key is udid/serial
}

// SetClientTransport configures transport of http client for high concurrency load testing
func (r *HRPRunner) SetClientTransport(maxConns int, disableKeepAlive bool, disableCompression bool) *HRPRunner {
	global.GVA_LOG.Info("[init] SetClientTransport", zap.Int("maxConns", maxConns), zap.Bool("disableKeepAlive", disableKeepAlive), zap.Bool("disableCompression", disableCompression))
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

// SetFailfast configures whether to stop running when one step fails.
func (r *HRPRunner) SetFailfast(failfast bool) *HRPRunner {
	global.GVA_LOG.Info("[init] SetFailfast", zap.Bool("failfast", failfast))
	r.failfast = failfast
	return r
}

// SetRequestsLogOn turns on request & response details logging.
func (r *HRPRunner) SetRequestsLogOn() *HRPRunner {
	global.GVA_LOG.Info("[init] SetRequestsLogOn")
	r.requestsLogOn = true
	return r
}

// SetHTTPStatOn turns on HTTP latency stat.
func (r *HRPRunner) SetHTTPStatOn() *HRPRunner {
	global.GVA_LOG.Info("[init] SetHTTPStatOn")
	r.httpStatOn = true
	return r
}

// SetPluginLogOn turns on plugin logging.
func (r *HRPRunner) SetPluginLogOn() *HRPRunner {
	global.GVA_LOG.Info("[init] SetPluginLogOn")
	r.pluginLogOn = true
	return r
}

// SetPython3Venv specifies python3 venv.
func (r *HRPRunner) SetPython3Venv(venv string) *HRPRunner {
	global.GVA_LOG.Info("[init] SetPython3Venv", zap.String("venv", venv))
	r.venv = venv
	return r
}

// SetProxyUrl configures the proxy URL, which is usually used to capture HTTP packets for debugging.
func (r *HRPRunner) SetProxyUrl(proxyUrl string) *HRPRunner {
	global.GVA_LOG.Info("[init] SetProxyUrl", zap.String("proxyUrl", proxyUrl))
	p, err := url.Parse(proxyUrl)
	if err != nil {
		global.GVA_LOG.Error("[init] invalid proxyUrl", zap.String("proxyUrl", proxyUrl), zap.Error(err))
		return r
	}
	r.httpClient.Transport = &http.Transport{
		Proxy:           http.ProxyURL(p),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	r.wsDialer.Proxy = http.ProxyURL(p)
	return r
}

// SetTimeout configures global timeout in seconds.
func (r *HRPRunner) SetTimeout(timeout time.Duration) *HRPRunner {
	global.GVA_LOG.Info("[init] SetTimeout", zap.Float64("timeout(seconds)", timeout.Seconds()))
	r.httpClient.Timeout = timeout
	return r
}

// SetSaveTests configures whether to save summary of tests.
func (r *HRPRunner) SetSaveTests(saveTests bool) *HRPRunner {
	global.GVA_LOG.Info("[init] SetSaveTests", zap.Bool("saveTests", saveTests))
	r.saveTests = saveTests
	return r
}

// GenHTMLReport configures whether to gen html report of api tests.
func (r *HRPRunner) GenHTMLReport() *HRPRunner {
	global.GVA_LOG.Info("[init] SetgenHTMLReport", zap.Bool("genHTMLReport", true))
	r.genHTMLReport = true
	return r
}

// Run starts to execute one or multiple testcases.
func (r *HRPRunner) Run(testcases ...ITestCase) error {
	global.GVA_LOG.Info("[init] Run", zap.Int("testcases", len(testcases)))
	event := sdk.EventTracking{
		Category: "RunAPITests",
		Action:   "hrp run",
	}
	// report start event
	go sdk.SendEvent(event)
	// report execution timing event
	defer sdk.SendEvent(event.StartTiming("execution"))
	// record execution data to summary
	s := newOutSummary()

	// load all testcases
	testCases, err := LoadTestCases(testcases...)
	if err != nil {
		global.GVA_LOG.Info("[init] Run", zap.Error(err))
		return err
	}

	// quit all plugins
	defer func() {
		pluginMap.Range(func(key, value interface{}) bool {
			if plugin, ok := value.(funplugin.IPlugin); ok {
				plugin.Quit()
			}
			return true
		})
	}()

	var runErr error
	// run testcase one by one
	for _, testcase := range testCases {
		// each testcase has its own case runner
		caseRunner, err := r.NewCaseRunner(testcase)
		if err != nil {
			global.GVA_LOG.Error("[init] Run", zap.Error(err))
			return err
		}

		// release UI driver session
		defer func() {
			for _, client := range r.uiClients {
				client.Driver.DeleteSession()
			}
		}()

		for it := caseRunner.parametersIterator; it.HasNext(); {
			// case runner can run multiple times with different parameters
			// each run has its own session runner
			sessionRunner := caseRunner.NewSession()
			err1 := sessionRunner.Start(it.Next())
			if err1 != nil {
				global.GVA_LOG.Error("[init] Run", zap.Error(err1))
				runErr = err1
			}
			caseSummary, err2 := sessionRunner.GetSummary()
			s.appendCaseSummary(caseSummary)
			if err2 != nil {
				global.GVA_LOG.Error("[init] Run", zap.Error(err2))
				if err1 != nil {
					runErr = errors.Wrap(err1, err2.Error())
				} else {
					runErr = err2
				}
			}

			if runErr != nil && r.failfast {
				break
			}
		}
	}
	s.Time.Duration = time.Since(s.Time.StartAt).Seconds()

	// save summary
	if r.saveTests {
		err := s.genSummary()
		if err != nil {
			return err
		}
	}

	// generate HTML report
	if r.genHTMLReport {
		err := s.genHTMLReport()
		if err != nil {
			return err
		}
	}

	return runErr
}

// NewCaseRunner creates a new case runner for testcase.
// each testcase has its own case runner
func (r *HRPRunner) NewCaseRunner(testcase *TestCase) (*CaseRunner, error) {
	caseRunner := &CaseRunner{
		testCase:  testcase,
		hrpRunner: r,
		parser:    newParser(),
	}

	// init parser plugin
	//plugin, err := initPlugin(testcase.Config.Path, r.venv, r.pluginLogOn)
	//压测运行时会同时运行多个plugin，用单例方式控制每次压测任务只能运行一个plugin
	plugin, err := initPlugin(testcase.Config.Path, r.venv, r.pluginLogOn)

	if err != nil {
		return nil, errors.Wrap(err, "init plugin failed")
	}
	if plugin != nil {
		caseRunner.parser.plugin = plugin
		caseRunner.rootDir = filepath.Dir(plugin.Path())
	}

	// parse testcase config
	if err := caseRunner.parseConfig(); err != nil {
		return nil, errors.Wrap(err, "parse testcase config failed")
	}

	// set testcase timeout in seconds
	if testcase.Config.Timeout != 0 {
		timeout := time.Duration(testcase.Config.Timeout*1000) * time.Millisecond
		r.SetTimeout(timeout)
	}

	// load plugin info to testcase config
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

	return caseRunner, nil
}

type CaseRunner struct {
	testCase  *TestCase
	hrpRunner *HRPRunner
	parser    *Parser

	parsedConfig       *TConfig
	parametersIterator *ParametersIterator
	rootDir            string // project root dir
}

// parseConfig parses testcase config, stores to parsedConfig.
func (r *CaseRunner) parseConfig() error {
	cfg := r.testCase.Config

	r.parsedConfig = &TConfig{}
	// deep copy config to avoid data racing
	if err := copier.Copy(r.parsedConfig, cfg); err != nil {
		global.GVA_LOG.Error("[init] Run", zap.Error(err))
		return err
	}

	// parse config variables
	parsedVariables, err := r.parser.ParseVariables(cfg.Variables)
	if err != nil {
		global.GVA_LOG.Error("[init] Run", zap.Any("variables", cfg.Variables), zap.Error(err))
		return err
	}
	r.parsedConfig.Variables = parsedVariables

	for k, _ := range cfg.Environs {
		parsedVariables[k] = cfg.Environs[k]
	}

	// parse config name
	parsedName, err := r.parser.ParseString(cfg.Name, parsedVariables)
	if err != nil {
		return errors.Wrap(err, "parse config name failed")
	}
	r.parsedConfig.Name = convertString(parsedName)

	// parse config base url
	parsedBaseURL, err := r.parser.ParseString(cfg.BaseURL, parsedVariables)
	if err != nil {
		return errors.Wrap(err, "parse config base url failed")
	}
	r.parsedConfig.BaseURL = convertString(parsedBaseURL)

	// merge config environment variables with base_url
	// priority: env base_url > base_url
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

	// merge config variables with environment variables
	// priority: env > config variables
	for k, v := range r.parsedConfig.Environs {
		r.parsedConfig.Variables[k] = v
	}

	// ensure correction of think time config
	r.parsedConfig.ThinkTimeSetting.checkThinkTime()

	// ensure correction of websocket config
	r.parsedConfig.WebSocketSetting.checkWebSocket()

	// parse testcase config parameters
	parametersIterator, err := initParametersIterator(r.parsedConfig)
	if err != nil {
		global.GVA_LOG.Error("[init] Run", zap.Any("parameters", r.parsedConfig.Parameters), zap.Any("parametersSetting", r.parsedConfig.ParametersSetting), zap.Error(err))
		return errors.Wrap(err, "parse testcase config parameters failed")
	}
	r.parametersIterator = parametersIterator

	// init iOS/Android clients
	if r.hrpRunner.uiClients == nil {
		r.hrpRunner.uiClients = make(map[string]*uixt.DriverExt)
	}
	for _, iosDeviceConfig := range r.parsedConfig.IOS {
		if iosDeviceConfig.UDID != "" {
			udid, err := r.parser.ParseString(iosDeviceConfig.UDID, parsedVariables)
			if err != nil {
				return errors.Wrap(err, "failed to parse ios device udid")
			}
			iosDeviceConfig.UDID = udid.(string)
		}
		device, err := uixt.NewIOSDevice(uixt.GetIOSDeviceOptions(iosDeviceConfig)...)
		if err != nil {
			return errors.Wrap(err, "init iOS device failed")
		}
		client, err := device.NewDriver(nil)
		if err != nil {
			return errors.Wrap(err, "init iOS WDA client failed")
		}
		r.hrpRunner.uiClients[device.UDID] = client
	}
	for _, androidDeviceConfig := range r.parsedConfig.Android {
		if androidDeviceConfig.SerialNumber != "" {
			sn, err := r.parser.ParseString(androidDeviceConfig.SerialNumber, parsedVariables)
			if err != nil {
				return errors.Wrap(err, "failed to parse android device serial")
			}
			androidDeviceConfig.SerialNumber = sn.(string)
		}
		device, err := uixt.NewAndroidDevice(uixt.GetAndroidDeviceOptions(androidDeviceConfig)...)
		if err != nil {
			return errors.Wrap(err, "init iOS device failed")
		}
		client, err := device.NewDriver(nil)
		if err != nil {
			return errors.Wrap(err, "init Android UIAutomator client failed")
		}
		r.hrpRunner.uiClients[device.SerialNumber] = client
	}

	return nil
}

// each boomer task initiates a new session
// in order to avoid data racing
func (r *CaseRunner) NewSession() *SessionRunner {
	sessionRunner := &SessionRunner{
		caseRunner: r,
	}
	sessionRunner.resetSession()
	return sessionRunner
}

// SessionRunner is used to run testcase and its steps.
// each testcase has its own SessionRunner instance and share session variables.
type SessionRunner struct {
	caseRunner       *CaseRunner
	sessionVariables map[string]interface{}
	// transactions stores transaction timing info.
	// key is transaction name, value is map of transaction type and time, e.g. start time and end time.
	transactions      map[string]map[transactionType]time.Time
	startTime         time.Time                  // record start time of the testcase
	summary           *TestCaseSummary           // record test case summary
	wsConnMap         map[string]*websocket.Conn // save all websocket connections
	pongResponseChan  chan string                // channel used to receive pong response message
	closeResponseChan chan *wsCloseRespObject    // channel used to receive close response message
}

func (r *SessionRunner) resetSession() {
	global.GVA_LOG.Info("reset session runner")
	r.sessionVariables = make(map[string]interface{})
	r.transactions = make(map[string]map[transactionType]time.Time)
	r.startTime = time.Now()
	r.summary = newSummary()
	r.wsConnMap = make(map[string]*websocket.Conn)
	r.pongResponseChan = make(chan string, 1)
	r.closeResponseChan = make(chan *wsCloseRespObject, 1)
}

// Start runs the test steps in sequential order.
// givenVars is used for data driven
func (r *SessionRunner) Start(givenVars map[string]interface{}) error {
	config := r.caseRunner.testCase.Config
	global.GVA_LOG.Info("run testcase start", zap.String("testcase", config.Name))

	// reset session runner
	r.resetSession()

	// update config variables with given variables
	r.InitWithParameters(givenVars)

	// run step in sequential order
	for _, step := range r.caseRunner.testCase.TestSteps {
		// TODO: parse step struct
		// parse step name
		parsedName, err := r.caseRunner.parser.ParseString(step.Name(), r.sessionVariables)
		if err != nil {
			parsedName = step.Name()
		}
		stepName := convertString(parsedName)
		global.GVA_LOG.Info("run step start", zap.String("step", stepName), zap.String("type", string(step.Type())))

		// run step
		// 失败重试
		var retry = r.caseRunner.parsedConfig.Retry
		var retryNum uint = 0
		if step.Struct().Retry > 0 {
			retry = step.Struct().Retry
		}
		var stepResult *StepResult
		for retryNum <= retry {
			stepResult, _ = step.Run(r)
			if stepResult.Success || retryNum >= retry {
				break
			}
			retryNum++
			time.Sleep(retrySleepTimer)
		}
		stepResult.Name = stepName
		stepResult.Retry = retryNum
		// update summary
		r.summary.Records = append(r.summary.Records, stepResult)
		r.summary.Stat.Total += 1
		if stepResult.Success {
			r.summary.Stat.Successes += 1
		} else {
			r.summary.Stat.Failures += 1
			// update summary result to failed
			r.summary.Success = false
		}

		// update extracted variables
		for k, v := range stepResult.ExportVars {
			r.sessionVariables[k] = v
		}

		var StepResults hrp.StepResultStruct
		stepResultStr, _ := json.Marshal(stepResult)
		json.Unmarshal(stepResultStr, &StepResults)
		for _, v := range step.Struct().ExportHeader {
			headerKey := textproto.CanonicalMIMEHeaderKey(v)
			if r.caseRunner.testCase.Config.Headers == nil {
				r.caseRunner.testCase.Config.Headers = make(map[string]string)
			}
			if r.caseRunner.parsedConfig.Headers == nil {
				r.caseRunner.parsedConfig.Headers = make(map[string]string)
			}
			r.caseRunner.testCase.Config.Headers[headerKey] = StepResults.Data.ReqResps.Request.Headers[headerKey]
			r.caseRunner.parsedConfig.Headers[headerKey] = StepResults.Data.ReqResps.Request.Headers[headerKey]
		}
		//for _, v := range step.Struct().ExportParameter {
		//	r.caseRunner.testCase.Config.Variables[v] = StepResults.ExportVars[v]
		//}

		if err == nil {
			global.GVA_LOG.Error("run step end",
				zap.String("step", stepResult.Name), zap.String("type", string(stepResult.StepType)),
				zap.Bool("success", true), zap.Any("exportVars", stepResult.ExportVars))
			continue
		}

		// failed
		global.GVA_LOG.Error("run step end", zap.String("step", stepResult.Name), zap.String("type", string(stepResult.StepType)), zap.Bool("success", false))
		// check if failfast
		if r.caseRunner.hrpRunner.failfast {
			return errors.Wrap(err, "abort running due to failfast setting")
		}
	}

	// close websocket connection after all steps done
	defer func() {
		for _, wsConn := range r.wsConnMap {
			if wsConn != nil {
				global.GVA_LOG.Info("websocket disconnected", zap.String("testcase", config.Name))
				err := wsConn.Close()
				if err != nil {
					global.GVA_LOG.Error("websocket disconnection failed", zap.Error(err), zap.String("testcase", config.Name))
				}
			}
		}
	}()

	global.GVA_LOG.Error("run testcase end", zap.String("testcase", config.Name))
	return nil
}

// ParseStepVariables merges step variables with config variables and session variables
func (r *SessionRunner) ParseStepVariables(stepVariables map[string]interface{}) (map[string]interface{}, error) {
	// override variables
	// step variables > session variables (extracted variables from previous steps)
	overrideVars := mergeVariables(stepVariables, r.sessionVariables)
	// step variables > testcase config variables
	overrideVars = mergeVariables(overrideVars, r.caseRunner.parsedConfig.Variables)

	// parse step variables
	parsedVariables, err := r.caseRunner.parser.ParseVariables(overrideVars)
	if err != nil {

		global.GVA_LOG.Error("parse step variables failed", zap.Any("variables", r.caseRunner.parsedConfig.Variables), zap.Error(err))
		return nil, errors.Wrap(err, "parse step variables failed")
	}
	return parsedVariables, nil
}

// InitWithParameters updates session variables with given parameters.
// this is used for data driven
func (r *SessionRunner) InitWithParameters(parameters map[string]interface{}) {
	if len(parameters) == 0 {
		return
	}

	global.GVA_LOG.Info("update session variables", zap.Any("parameters", parameters))
	for k, v := range parameters {
		r.sessionVariables[k] = v
	}
}

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
		// add WDA/UIA logs to summary
		Log, err := client.Driver.StopCaptureLog()
		if err != nil {
			return caseSummary, err
		}
		logs := map[string]interface{}{
			"uuid":    uuid,
			"content": Log,
		}

		// stop performance monitor
		logs["performance"] = client.GetPerfData()

		caseSummary.Logs = append(caseSummary.Logs, logs)
	}

	return caseSummary, nil
}
