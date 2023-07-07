package main

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod downloa

import (
	"fmt"
	"github.com/test-instructor/yangfan/server/grpc/client"
	"github.com/test-instructor/yangfan/server/grpc/pkg"

	"github.com/test-instructor/yangfan/server/core"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/initialize"
	"go.uber.org/zap"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

func RunSetTimerTaskClient() {
	host := fmt.Sprintf("%s:%s", global.GVA_CONFIG.YangFan.Background, global.GVA_CONFIG.YangFan.BackgroundGrpcPort)
	c, err := client.NewClient(host)
	if err != nil {
		global.GVA_LOG.Error("[RunClient]创建客户端失败", zap.Error(err))
	}
	p := pkg.NewRunInstallPkg(c)
	go p.RunSetTimerTask()
	go p.RunInstallPkg()
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)

	rand.Seed(time.Now().UnixNano())
	global.GVA_VP = core.Viper()
	global.GVA_LOG = core.Zap()
	global.GVA_DB = initialize.Gorm()
	if global.GVA_DB.Error != nil {
		global.GVA_LOG.Error("register db", zap.Error(global.GVA_DB.Error))
		os.Exit(0)
	}
	initialize.TimerTaskCase()
	go RunSetTimerTaskClient()
	s := <-c
	fmt.Println("exit", s)
}
