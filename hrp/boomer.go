package hrp

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/test-instructor/yangfan/server/global"
	"go.uber.org/zap"

	"github.com/httprunner/funplugin"
	"golang.org/x/net/context"

	"github.com/test-instructor/yangfan/hrp/internal/builtin"
	"github.com/test-instructor/yangfan/hrp/internal/code"
	"github.com/test-instructor/yangfan/hrp/internal/json"
	"github.com/test-instructor/yangfan/hrp/internal/sdk"
	"github.com/test-instructor/yangfan/hrp/pkg/boomer"
)

func NewStandaloneBoomer(spawnCount int64, spawnRate float64) *HRPBoomer {
	b := &HRPBoomer{
		Boomer:       boomer.NewStandaloneBoomer(spawnCount, spawnRate),
		pluginsMutex: new(sync.RWMutex),
	}

	b.hrpRunner = NewRunner(nil)
	return b
}

var HrpBoomer *HRPBoomer
var mutexBoomer sync.Mutex

func NewMasterBoomerSingleton(masterBindHost string, masterBindPort int) *HRPBoomer {
	if HrpBoomer == nil {
		mutexBoomer.Lock()
		defer mutexBoomer.Unlock()
		if HrpBoomer == nil {
			HrpBoomer = &HRPBoomer{
				Boomer:       boomer.NewMasterBoomer(masterBindHost, masterBindPort),
				pluginsMutex: new(sync.RWMutex),
			}
			HrpBoomer.hrpRunner = NewRunner(nil)
		}
	}
	return HrpBoomer
}

func NewMasterBoomer(masterBindHost string, masterBindPort int) *HRPBoomer {
	b := &HRPBoomer{
		Boomer:       boomer.NewMasterBoomer(masterBindHost, masterBindPort),
		pluginsMutex: new(sync.RWMutex),
	}
	b.hrpRunner = NewRunner(nil)
	return b
}

func NewWorkerBoomer(masterHost string, masterPort int) *HRPBoomer {
	b := &HRPBoomer{
		Boomer:       boomer.NewWorkerBoomer(masterHost, masterPort),
		pluginsMutex: new(sync.RWMutex),
	}

	b.hrpRunner = NewRunner(nil)
	// set client transport for high concurrency load testing
	b.hrpRunner.SetClientTransport(b.GetSpawnCount(), b.GetDisableKeepAlive(), b.GetDisableCompression())
	return b
}

type HRPBoomer struct {
	*boomer.Boomer
	hrpRunner       *HRPRunner
	plugins         []funplugin.IPlugin // each task has its own plugin process
	pluginsMutex    *sync.RWMutex       // avoid data race
	debugtalk       *debugTalkOperation
	runBoomerMaster *RunBoomerMaster
	OutputDB        *boomer.DbOutput
	ReportName      string
}

func (b *HRPBoomer) InitBoomer() {
	if !b.GetProfile().DisableConsoleOutput {
		b.AddOutput(boomer.NewConsoleOutput())
	}
	b.SetSpawnCount(b.GetProfile().SpawnCount)
	b.SetSpawnRate(b.GetProfile().SpawnRate)
	b.SetRunTime(b.GetProfile().RunTime)
	if b.GetProfile().LoopCount > 0 {
		b.SetLoopCount(b.GetProfile().LoopCount)
	}
	b.SetRateLimiter(b.GetProfile().MaxRPS, b.GetProfile().RequestIncreaseRate)
	b.SetDisableKeepAlive(b.GetProfile().DisableKeepalive)
	b.SetDisableCompression(b.GetProfile().DisableCompression)
	b.SetClientTransport()
	b.EnableCPUProfile(b.GetProfile().CPUProfile, b.GetProfile().CPUProfileDuration)
	b.EnableMemoryProfile(b.GetProfile().MemoryProfile, b.GetProfile().MemoryProfileDuration)
}

func (b *HRPBoomer) SetClientTransport() *HRPBoomer {
	// set client transport for high concurrency load testing
	b.hrpRunner.SetClientTransport(b.GetSpawnCount(), b.GetDisableKeepAlive(), b.GetDisableCompression())
	return b
}

// SetPython3Venv specifies python3 venv.
func (b *HRPBoomer) SetPython3Venv(venv string) *HRPBoomer {
	b.hrpRunner.SetPython3Venv(venv)
	return b
}

// Run starts to run load test for one or multiple testcases.
func (b *HRPBoomer) Run(testcases ...ITestCase) {
	event := sdk.EventTracking{
		Category: "RunLoadTests",
		Action:   "hrp boom",
	}
	// report start event
	go sdk.SendEvent(event)
	// report execution timing event
	defer sdk.SendEvent(event.StartTiming("execution"))

	// quit all plugins
	defer func() {
		pluginMap.Range(func(key, value interface{}) bool {
			if plugin, ok := value.(funplugin.IPlugin); ok {
				plugin.Quit()
			}
			return true
		})
	}()

	taskSlice := b.ConvertTestCasesToBoomerTasks(testcases...)

	b.Boomer.Run(taskSlice...)
}

func (b *HRPBoomer) ConvertTestCasesToBoomerTasks(testcases ...ITestCase) (taskSlice []*boomer.Task) {
	// load all testcases
	testCases, err := LoadTestCases(testcases...)
	if err != nil {
		global.GVA_LOG.Error("failed to load testcases", zap.Any("err", err))
		os.Exit(code.GetErrorCode(err))
	}

	for _, testcase := range testCases {
		rendezvousList := initRendezvous(testcase, int64(b.GetSpawnCount()))
		task := b.convertBoomerTask(testcase, rendezvousList)
		taskSlice = append(taskSlice, task)
		waitRendezvous(rendezvousList, b)
	}
	return taskSlice
}

func (b *HRPBoomer) ParseTestCases(testCases []*TestCase) []*TCase {
	var parsedTestCases []*TCase
	var config *TConfig
	for _, tc := range testCases {
		if config == nil {
			config = tc.Config
		}
		tc.Config = config
		if global.HrpMode == global.HrpModeMaster {
			b.initPluginMaster(tc.Config.ProjectID)
			tc.Config.Path = b.debugtalk.FilePath
			tc.Config.ReportID = b.OutputDB.PReport.ID
			b.Boomer.SetServerReportID(b.OutputDB.PReport.ID)
			reportName := fmt.Sprintf("%s_id_%d", b.OutputDB.PReport.Name, b.OutputDB.PReport.ID)
			stats := boomer.NewPrometheusPusherStats(global.GVA_CONFIG.YangFan.PrometheusPushgatewayURL, "yangfan", "stats", reportName)
			stats.OnStart()
		}
		caseRunner, err := b.hrpRunner.NewCaseRunner(tc)
		if err != nil {
			global.GVA_LOG.Error("failed to create runner", zap.Any("err", err))
			os.Exit(code.GetErrorCode(err))
		}
		caseRunner.parsedConfig.Parameters = caseRunner.parametersIterator.outParameters()

		parsedTestCases = append(parsedTestCases, &TCase{
			Config:    caseRunner.parsedConfig,
			TestSteps: caseRunner.testCase.ToTCase().TestSteps,
		})
	}
	return parsedTestCases
}

func (b *HRPBoomer) initPlugin(path string) string {
	paths := strings.Split(path, "/")
	ids := strings.Split(paths[1], "_")
	id, _ := strconv.Atoi(ids[0])
	b.debugtalk = new(debugTalkOperation)
	b.debugtalk.ProjectID = uint(id)
	b.debugtalk.FilePath = path
	b.debugtalk.CreateDebugTalkWork()
	b.debugtalk.RunDebugTalkFile()
	return b.debugtalk.FilePath
}

func (b *HRPBoomer) initPluginMaster(id uint) {
	b.debugtalk = new(debugTalkOperation)
	b.debugtalk.ProjectID = id
	b.debugtalk.CreateDebugTalkMaster()
	b.debugtalk.RunDebugTalkFile()
}

func (b *HRPBoomer) TestCasesToBytes(testcases ...ITestCase) []byte {
	// load all testcases
	testCases, err := LoadTestCases(testcases...)
	if err != nil {
		global.GVA_LOG.Error("failed to load testcases", zap.Any("err", err))
		os.Exit(code.GetErrorCode(err))
	}
	tcs := b.ParseTestCases(testCases)
	testCasesBytes, err := json.Marshal(tcs)
	if err != nil {
		global.GVA_LOG.Error("failed to marshal testcases", zap.Any("err", err))
		return nil
	}
	return testCasesBytes
}

func (b *HRPBoomer) BytesToTCases(testCasesBytes []byte) []*TCase {
	var testcase []*TCase
	err := json.Unmarshal(testCasesBytes, &testcase)
	if err != nil {
		global.GVA_LOG.Error("failed to unmarshal testcases", zap.Any("err", err))
	}
	return testcase
}

func (b *HRPBoomer) Quit() {
	b.Boomer.Quit()
}

var mutex sync.Mutex

func (b *HRPBoomer) parseTCases(testCases []*TCase) (testcases []ITestCase) {
	if testCases != nil && len(testCases) > 0 {
		tc := testCases[0]
		if global.HrpMode == global.HrpModeWork {
			global.IgnoreInstall = true
			b.initPlugin(tc.Config.Path)
			if b.OutputDB == nil {
				mutex.Lock()
				defer mutex.Unlock()
				if b.OutputDB == nil {
					b.OutputDB = boomer.NewDbOutputWork(tc.Config.ReportID, tc.Config.CaseID)
					b.Boomer.AddOutput(b.OutputDB)
				}
			}
		}
		if b.GetProfile().PrometheusPushgatewayURL != "" && (global.HrpMode == global.HrpModeMaster || global.HrpMode == global.HrpModeWork) {
			b.ReportName = fmt.Sprintf("%s_id_%d", b.OutputDB.PReport.Name, b.OutputDB.PReport.ID)
			b.AddOutput(boomer.NewPrometheusPusherOutput(b.GetProfile().PrometheusPushgatewayURL, "hrp", b.GetMode(), b.ReportName))
		}
	}
	for _, tc := range testCases {
		// create temp dir to save testcase
		tempDir, err := ioutil.TempDir("", "hrp_testcases")
		if err != nil {
			global.GVA_LOG.Error("failed to create hrp testcases directory", zap.Any("err", err))
			return
		}

		if tc.Config.PluginSetting != nil {
			tc.Config.PluginSetting.Path = filepath.Join(tempDir, fmt.Sprintf("debugtalk.%s", tc.Config.PluginSetting.Type))
			err = builtin.Bytes2File(tc.Config.PluginSetting.Content, tc.Config.PluginSetting.Path)
			if err != nil {
				global.GVA_LOG.Error("failed to save plugin file", zap.Any("err", err))
				return
			}
			tc.Config.PluginSetting.Content = nil // remove the content in testcase
		}

		if tc.Config.Environs != nil {
			envContent := ""
			for k, v := range tc.Config.Environs {
				envContent += fmt.Sprintf("%s=%s\n", k, v)
			}
			err = os.WriteFile(filepath.Join(tempDir, ".env"), []byte(envContent), 0o644)
			if err != nil {
				global.GVA_LOG.Error("failed to dump environs", zap.Any("err", err))
				return
			}
		}

		tc.Config.Path = filepath.Join(tempDir, "test-case.json")
		err = builtin.Dump2JSON(tc, tc.Config.Path)
		if err != nil {
			global.GVA_LOG.Error("failed to dump testcases", zap.Any("err", err))
			return
		}

		tesecase, err := tc.toTestCase()
		if err != nil {
			global.GVA_LOG.Error("failed to load testcases", zap.Any("err", err))
			return
		}

		testcases = append(testcases, tesecase)
	}
	return testcases
}

func (b *HRPBoomer) initWorker(profile *boomer.Profile) {
	// if no IP address is specified, the default IP address is that of the master
	if profile.PrometheusPushgatewayURL != "" {
		urlSlice := strings.Split(profile.PrometheusPushgatewayURL, ":")
		if len(urlSlice) != 2 {
			profile.PrometheusPushgatewayURL = ""
		} else {
			if urlSlice[0] == "" {
				urlSlice[0] = b.Boomer.GetMasterHost()
			}
		}
		profile.PrometheusPushgatewayURL = strings.Join(urlSlice, ":")
	}

	b.SetProfile(profile)
	b.InitBoomer()
}

func (b *HRPBoomer) rebalanceRunner(profile *boomer.Profile) {
	b.SetSpawnCount(profile.SpawnCount)
	b.SetSpawnRate(profile.SpawnRate)
	b.GetRebalanceChan() <- true
	global.GVA_LOG.Info("rebalance tasks successfully", zap.Any("profile", profile))
}

func (b *HRPBoomer) PollTasks(ctx context.Context) {
	for {
		select {
		case task := <-b.Boomer.GetTasksChan():
			// 清理过时测试用例任务
			if len(b.Boomer.GetTasksChan()) > 0 {
				continue
			}
			// Todo: 过滤掉已经传输过的task
			if task.TestCasesBytes != nil {
				// init boomer with profile
				b.initWorker(task.Profile)
				// get testcases
				testcases := b.parseTCases(b.BytesToTCases(task.TestCasesBytes))
				global.GVA_LOG.Info("starting to run tasks", zap.Any("testcases", testcases), zap.Any("profile", b.GetProfile()))
				// run testcases
				go b.Run(testcases...)
			} else {
				// rebalance runner with profile
				go b.rebalanceRunner(task.Profile)
			}

		case <-b.Boomer.GetCloseChan():
			return
		case <-ctx.Done():
			return
		}
	}
}

func (b *HRPBoomer) PollTestCases(ctx context.Context) {
	// quit all plugins
	defer func() {
		pluginMap.Range(func(key, value interface{}) bool {
			if plugin, ok := value.(funplugin.IPlugin); ok {
				plugin.Quit()
			}
			return true
		})
	}()

	for {
		select {
		case <-b.Boomer.ParseTestCasesChan():
			var tcs []ITestCase
			for _, tc := range b.GetTestCasesPath() {
				tcp := TestCasePath(tc)
				tcs = append(tcs, &tcp)
			}
			b.TestCaseBytesChan() <- b.TestCasesToBytes(tcs...)
			global.GVA_LOG.Info("put testcase successfully")
		case <-b.Boomer.GetCloseChan():
			return
		case <-ctx.Done():
			return
		}
	}
}

func (b *HRPBoomer) convertBoomerTask(testcase *TestCase, rendezvousList []*Rendezvous) *boomer.Task {
	// init case runner for testcase
	// this runner is shared by multiple session runners
	caseRunner, err := b.hrpRunner.NewCaseRunner(testcase)
	if err != nil {
		global.GVA_LOG.Error("failed to create runner", zap.Any("err", err))
		os.Exit(code.GetErrorCode(err))
	}
	if caseRunner.parser.plugin != nil {
		b.pluginsMutex.Lock()
		b.plugins = append(b.plugins, caseRunner.parser.plugin)
		b.pluginsMutex.Unlock()
	}

	// broadcast to all rendezvous at once when spawn done
	go func() {
		<-b.GetSpawnDoneChan()
		for _, rendezvous := range rendezvousList {
			rendezvous.setSpawnDone()
		}
	}()

	// set paramters mode for load testing
	parametersIterator := caseRunner.parametersIterator
	parametersIterator.SetUnlimitedMode()

	// reset start time only once
	once := sync.Once{}

	// update session variables mutex
	mutex := sync.Mutex{}

	return &boomer.Task{
		Name:   testcase.Config.Name,
		Weight: testcase.Config.Weight,
		Fn: func() {
			testcaseSuccess := true    // flag whole testcase result
			transactionSuccess := true // flag current transaction result

			// init session runner
			sessionRunner := caseRunner.NewSession()

			mutex.Lock()
			if parametersIterator.HasNext() {
				sessionRunner.InitWithParameters(parametersIterator.Next())
			}
			mutex.Unlock()

			startTime := time.Now()
			for _, step := range testcase.TestSteps {
				// TODO: parse step struct
				// parse step name
				parsedName, err := caseRunner.parser.ParseString(step.Name(), sessionRunner.sessionVariables)
				if err != nil {
					parsedName = step.Name()
				}
				stepName := convertString(parsedName)
				// reset start time only once before step
				once.Do(func() {
					b.Boomer.ResetStartTime()
				})
				stepResult, err := step.Run(sessionRunner)
				// update step result name with parsed step name
				stepResult.Name = stepName
				// record requests result of the step if step type is testcase
				if stepResult.StepType == stepTypeTestCase && stepResult.Data != nil {
					// record requests of testcase step
					for _, result := range stepResult.Data.([]*StepResult) {
						if result.Success {
							b.RecordSuccess(string(result.StepType), result.Name, result.Elapsed, result.ContentSize)
						} else {
							exception, _ := result.Attachments.(string)
							b.RecordFailure(string(result.StepType), result.Name, result.Elapsed, exception)
						}
					}
				}
				// record step failure
				if err != nil {
					// step failed
					var elapsed int64
					if stepResult != nil {
						elapsed = stepResult.Elapsed
					}
					b.RecordFailure(string(step.Type()), stepResult.Name, elapsed, err.Error())

					// update flag
					testcaseSuccess = false
					transactionSuccess = false

					if b.hrpRunner.failfast {
						global.GVA_LOG.Error("abort running due to failfast setting", zap.Any("err", err))
						break
					}
					global.GVA_LOG.Warn("run step failed, continue next step", zap.Any("err", err))
					continue
				}

				// record step success
				if stepResult.StepType == stepTypeTransaction {
					// transaction
					// FIXME: support nested transactions
					if step.Struct().Transaction.Type == transactionEnd { // only record when transaction ends
						b.RecordTransaction(step.Struct().Transaction.Name, transactionSuccess, stepResult.Elapsed, 0)
						transactionSuccess = true // reset flag for next transaction
					}
				} else if stepResult.StepType == stepTypeRendezvous {
					// rendezvous
				} else if stepResult.StepType == stepTypeThinkTime {
					// think time
					// no record required
				} else {
					// request or testcase step
					b.RecordSuccess(string(step.Type()), stepResult.Name, stepResult.Elapsed, stepResult.ContentSize)
					// update extracted variables
					for k, v := range stepResult.ExportVars {
						sessionRunner.sessionVariables[k] = v
					}
				}
			}
			endTime := time.Now()

			// report duration for transaction without end
			for name, transaction := range sessionRunner.transactions {
				if len(transaction) == 1 {
					// if transaction end time not exists, use testcase end time instead
					duration := endTime.Sub(transaction[transactionStart])
					b.RecordTransaction(name, transactionSuccess, duration.Milliseconds(), 0)
				}
			}

			// report testcase as a whole Action transaction, inspired by LoadRunner
			b.RecordTransaction("Action", testcaseSuccess, endTime.Sub(startTime).Milliseconds(), 0)
		},
	}
}

func (b *HRPBoomer) PollTestCasesPlatform(ctx context.Context) {
	// quit all plugins
	defer func() {
		pluginMap.Range(func(key, value interface{}) bool {
			if plugin, ok := value.(funplugin.IPlugin); ok {
				plugin.Quit()
			}
			return true
		})
	}()

	for {
		select {
		case <-b.Boomer.ParseTestCasesChan():
			var tcs []ITestCase
			id := b.GetTestCasesID()
			_ = id
			b.runBoomerMaster = NewBoomerMaster(id)
			err := b.runBoomerMaster.LoadCase()
			if err != nil {
				global.GVA_LOG.Error("获取用例失败", zap.Any("err", err))
				return
			}
			if err != nil {
				return
			}
			b.OutputDB = boomer.NewDbOutput(b.runBoomerMaster.PReport, b.runBoomerMaster.PTask)
			b.Boomer.AddOutput(b.OutputDB)
			tcs = append(tcs, b.runBoomerMaster.TCM.Case...)
			b.TestCaseBytesChan() <- b.TestCasesToBytes(tcs...)
			global.GVA_LOG.Info("put testcase successfully")
		case <-b.Boomer.GetCloseChan():
			return
		case <-ctx.Done():
			return
		}
	}
}
