package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/test-instructor/yangfan/data/internal/httpserver"
	"github.com/test-instructor/yangfan/data/internal/timer"
	"github.com/test-instructor/yangfan/httprunner/hrp"

	"github.com/test-instructor/yangfan/server/v2/core"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/initialize"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 全局调度器实例
var scheduler *timer.Scheduler

// 全局 HTTP Server 实例
var httpSrv *httpserver.Server

func main() {
	// 1. Initialize System
	initializeSystem()

	// 2. Initialize HTTP Server (no auth, fixed port)
	initializeHTTP()

	// 3. Initialize and Start Timer
	initializeTimer()

	// 4. Initialize DB/Redis/Python in background (avoid blocking HTTP/timer startup)
	initializeBackendAsync()

	// 4. Wait for Signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 5. Graceful Shutdown
	if scheduler != nil {
		scheduler.Stop()
	}
	if httpSrv != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = httpSrv.Shutdown(ctx)
	}
	global.GVA_LOG.Info("Shutting down dataWarehouse...")
}

// initializeSystem 初始化系统所有组件
// 提取为单独函数以便于系统重载时调用
func initializeSystem() {
	global.GVA_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	hrp.InitLogger("INFO", true, false)
	os.Setenv("DISABLE_GA", "true") // 禁用GA
}

func initializeBackendAsync() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				global.GVA_LOG.Error("后端初始化异常退出", zap.Any("recover", r))
			}
		}()

		for {
			global.GVA_LOG.Info("数据库初始化开始")
			db, err := tryInitDB()
			if err == nil && db != nil {
				global.GVA_DB = db
				global.GVA_LOG.Info("数据库初始化成功")
				break
			}
			global.GVA_LOG.Warn("数据库初始化失败，稍后重试", zap.Error(err))
			time.Sleep(5 * time.Second)
		}

		if global.GVA_CONFIG.System.UseRedis {
			global.GVA_LOG.Info("Redis 初始化开始")
			initialize.Redis()
			if global.GVA_CONFIG.System.UseMultipoint {
				initialize.RedisList()
			}
			global.GVA_LOG.Info("Redis 初始化结束")
		}

		global.GVA_LOG.Info("Python 初始化开始")
		initialize.InitPython()
		global.GVA_LOG.Info("Python 初始化结束")
	}()
}

func tryInitDB() (db *gorm.DB, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
			db = nil
		}
	}()
	return initialize.Gorm(), nil
}

// initializeTimer 初始化定时任务
// 环境变量 DEBUG_MODE=true 启用调试模式（每 10 分钟执行）
func initializeTimer() {
	if os.Getenv("DEBUG_MODE") == "true" {
		global.GVA_LOG.Info("启用调试模式，定时任务每 10 分钟执行一次")
		scheduler = timer.NewDebugScheduler()
	} else {
		scheduler = timer.NewScheduler()
	}
	if err := scheduler.Start(); err != nil {
		global.GVA_LOG.Error("启动定时任务失败", zap.Error(err))
	} else {
		global.GVA_LOG.Info("定时任务启动完成")
	}
}

// initializeHTTP 初始化 HTTP 服务（无需鉴权）
func initializeHTTP() {
	httpSrv = httpserver.NewServer()
	// 使用配置的数据仓库端口
	addr := fmt.Sprintf(":%d", global.GVA_DW_PORT)
	go func() {
		global.GVA_LOG.Info("HTTP 服务启动", zap.String("addr", addr))
		if err := httpSrv.Start(addr); err != nil {
			global.GVA_LOG.Error("HTTP 服务启动失败", zap.Error(err))
		}
	}()
}
