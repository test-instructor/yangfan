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
	"github.com/test-instructor/yangfan/server/source/yangfan"
	"go.uber.org/zap"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

func RunPkgInstallClient() {
	host := fmt.Sprintf("%s:%s", global.GVA_CONFIG.GrpcServer.Background, global.GVA_CONFIG.GrpcServer.BackgroundGrpcPort)
	c, err := client.NewClient(host)
	if err != nil {
		global.GVA_LOG.Error("[RunClient]创建客户端失败", zap.Error(err))
	}
	p := pkg.NewRunInstallPkg(c)
	p.RunClient()
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)

	rand.Seed(time.Now().UnixNano())
	global.GVA_VP = core.Viper()
	global.GVA_LOG = core.Zap()
	global.GVA_DB = initialize.Gorm()
	yangfan.InitPythonPackage(false)
	initialize.TimerTaskCase()
	go RunPkgInstallClient()
	s := <-c
	fmt.Println("exit", s)
}
