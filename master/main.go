package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/test-instructor/yangfan/master/server"
	"github.com/test-instructor/yangfan/server/core"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/grpc/client"
	"github.com/test-instructor/yangfan/server/grpc/pkg"
	"github.com/test-instructor/yangfan/server/initialize"
	"github.com/test-instructor/yangfan/server/source/yangfan"
	"go.uber.org/zap"
	"math/rand"
	"os"
	"time"
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
	global.HrpMode = global.HrpModeMaster
	zerolog.SetGlobalLevel(zerolog.WarnLevel)
	zap.ReplaceGlobals(global.GVA_LOG)
	if global.GVA_DB.Error != nil {
		global.GVA_LOG.Error("register db", zap.Error(global.GVA_DB.Error))
		os.Exit(0)
	}
	go yangfan.InitPythonPackage(true)
	go RunPkgInstallClient()
	b := server.NewMasterBoom()
	b.Run()

	//./hrp boom --worker --master-host 0.0.0.0 --master-port 7966 --ignore-quit
	//./hrp boom --master --master-bind-host 0.0.0.0 --master-bind-port 7966 --master-http-address "0.0.0.0:9092"
	//127.0.0.1:9092
}
