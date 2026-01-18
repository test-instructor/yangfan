package main

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/test-instructor/yangfan/httprunner/hrp"
	"github.com/test-instructor/yangfan/run/internal/mq"
	"github.com/test-instructor/yangfan/run/internal/service"
	runtimer "github.com/test-instructor/yangfan/run/internal/timer"
	"github.com/test-instructor/yangfan/server/v2/core"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/initialize"
	server_mq "github.com/test-instructor/yangfan/server/v2/utils/mq"
	"go.uber.org/zap"
)

func main() {
	// 1. Initialize System
	initializeSystem()

	mode := strings.ToLower(strings.TrimSpace(os.Getenv("RUN_SERVICE_MODE")))
	if mode == "" {
		mode = "runner"
	}
	enableRunnerConsumer := mode == "runner" || mode == "all"
	enableTimer := mode == "timer" || mode == "all"

	if mode != "runner" && mode != "timer" && mode != "all" {
		global.GVA_LOG.Warn("Unknown RUN_SERVICE_MODE, fallback to runner", zap.String("mode", mode))
		mode = "runner"
		enableRunnerConsumer = true
		enableTimer = false
	}

	runContent := mode

	// 2. Initialize MQ Client
	mqConfig := global.GVA_CONFIG.MQ
	mqLoader := server_mq.NewMQConfigLoader()
	if err := mqLoader.LoadAndValidate(mqConfig); err != nil {
		global.GVA_LOG.Fatal("MQ Config validation failed", zap.Error(err))
	}

	mqClient, err := server_mq.NewMQClient(mqConfig, global.GVA_LOG)
	if err != nil {
		global.GVA_LOG.Fatal("MQ Initialization failed", zap.Error(err))
	}
	defer mqClient.Close()

	// 3. Prepare Runner Config (only NodeName is used for MQ queue naming)
	runnerConfig := global.GVA_CONFIG.Runner
	hostname, err := os.Hostname()
	if err != nil || strings.TrimSpace(hostname) == "" {
		hostname = "unknown"
	}

	// Support Env Override
	if envNode := os.Getenv("NODE_NAME"); envNode != "" {
		runnerConfig.NodeName = envNode
	}

	if runnerConfig.NodeName == "" {
		runnerConfig.NodeName = "yf_runner_" + hostname
		global.GVA_LOG.Info("Runner NodeName not specified, using generated name", zap.String("name", runnerConfig.NodeName))
	}

	nodeAlias := strings.TrimSpace(os.Getenv("NODE_ALIAS"))
	if nodeAlias == "" {
		nodeAlias = hostname
	}

	runnerSvc := service.NewRunnerService(runnerConfig.NodeName, nodeAlias, runContent, runnerConfig.Port)
	if err := runnerSvc.Register(); err != nil {
		global.GVA_LOG.Fatal("Runner registration failed", zap.Error(err))
	}
	runnerSvc.StartHeartbeat()
	defer runnerSvc.Stop()

	var runnerConsumer *mq.RunnerTaskConsumer
	if enableRunnerConsumer {
		runnerConsumer = mq.NewRunnerTaskConsumer(mqClient, runnerConfig.NodeName)
		if err := runnerConsumer.StartListen(); err != nil {
			global.GVA_LOG.Fatal("Failed to start runner consumer", zap.Error(err))
		}
		defer runnerConsumer.Stop()
	} else {
		global.GVA_LOG.Info("Runner task consumer disabled by RUN_SERVICE_MODE", zap.String("mode", mode))
	}

	if enableTimer {
		scheduler := runtimer.NewTimerTaskScheduler(runnerConfig.NodeName)
		scheduler.Start(context.Background())
		defer scheduler.Stop()

		timerConsumer := mq.NewTimerTaskControlConsumer(mqClient, runnerConfig.NodeName, scheduler)
		if err := timerConsumer.StartListen(); err != nil {
			global.GVA_LOG.Fatal("Failed to start timer control consumer", zap.Error(err))
		}
		defer timerConsumer.Stop()

		global.GVA_LOG.Info("Timer mode enabled by RUN_SERVICE_MODE", zap.String("mode", mode))
	}

	// 5. Wait for Signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	global.GVA_LOG.Info("Shutting down runner...")
}

// initializeSystem 初始化系统所有组件
// 提取为单独函数以便于系统重载时调用
func initializeSystem() {
	global.GVA_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	hrp.InitLogger("INFO", true, false)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	if global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
		if global.GVA_CONFIG.System.UseMultipoint {
			initialize.RedisList()
		}
	}
	go func() {
		// 测试平台相关初始化
		//time.Sleep(10 * time.Second)
		initialize.InitPython()
	}()
	os.Setenv("DISABLE_GA", "true") // 禁用GA
	// initialize.Timer() // Runner might not need server timers
	// initialize.DBList()
	// initialize.SetupHandlers() // 注册全局函数
}
