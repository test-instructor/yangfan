package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/test-instructor/yangfan/run/server"
	"github.com/test-instructor/yangfan/server/core"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/grpc/client"
	"github.com/test-instructor/yangfan/server/grpc/pkg"
	"github.com/test-instructor/yangfan/server/initialize"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func RunPkgInstallClient() {
	host := fmt.Sprintf("%s:%s", global.GVA_CONFIG.YangFan.Background, global.GVA_CONFIG.YangFan.BackgroundGrpcPort)
	c, err := client.NewClient(host)
	if err != nil {
		global.GVA_LOG.Error("[RunClient]创建客户端失败", zap.Error(err))
	}
	p := pkg.NewRunInstallPkg(c)
	p.RunInstallPkg()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	global.GVA_VP = core.Viper()
	global.GVA_LOG = core.Zap()
	global.GVA_DB = initialize.Gorm()
	zap.ReplaceGlobals(global.GVA_LOG)
	if global.GVA_DB.Error != nil {
		global.GVA_LOG.Error("register db", zap.Error(global.GVA_DB.Error))
		os.Exit(0)
	}
	//yangfan.InitPythonPackage(true)
	//go RunPkgInstallClient()
	server.StartRunServer(":9099")
}
