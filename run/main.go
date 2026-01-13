package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/test-instructor/yangfan/httprunner/hrp"
	"github.com/test-instructor/yangfan/run/internal/mq"
	"github.com/test-instructor/yangfan/run/internal/service"
	"github.com/test-instructor/yangfan/server/v2/core"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/initialize"
	server_mq "github.com/test-instructor/yangfan/server/v2/utils/mq"
	"go.uber.org/zap"
)

func main() {
	// 1. Initialize System
	initializeSystem()

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
	// Support Env Override
	if envNode := os.Getenv("NODE_NAME"); envNode != "" {
		runnerConfig.NodeName = envNode
	}

	if runnerConfig.NodeName == "" {
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "unknown"
		}
		runnerConfig.NodeName = "yf_runner_" + hostname
		global.GVA_LOG.Info("Runner NodeName not specified, using generated name", zap.String("name", runnerConfig.NodeName))
	}

	runnerSvc := service.NewRunnerService(runnerConfig.NodeName, runnerConfig.Port)
	if err := runnerSvc.Register(); err != nil {
		global.GVA_LOG.Fatal("Runner registration failed", zap.Error(err))
	}
	runnerSvc.StartHeartbeat()
	defer runnerSvc.Stop()

	// 4. Initialize MQ Consumer
	consumer := mq.NewRunnerTaskConsumer(mqClient, runnerConfig.NodeName)
	if err := consumer.StartListen(); err != nil {
		global.GVA_LOG.Fatal("Failed to start consumer", zap.Error(err))
	}
	defer consumer.Stop()

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
	os.Setenv("DISABLE_GA", "true") // 禁用GA
	// initialize.Timer() // Runner might not need server timers
	// initialize.DBList()
	// initialize.SetupHandlers() // 注册全局函数
}
