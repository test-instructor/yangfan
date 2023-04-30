package main

import (
	"github.com/rs/zerolog"
	"github.com/test-instructor/yangfan/master/boomer"
	"github.com/test-instructor/yangfan/server/core"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/initialize"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	rand.Seed(time.Now().UnixNano())
	global.GVA_VP = core.Viper()
	global.GVA_LOG = core.Zap()
	global.GVA_DB = initialize.Gorm()
	zerolog.SetGlobalLevel(zerolog.WarnLevel)
	zap.ReplaceGlobals(global.GVA_LOG)
	b := boomer.NewMasterBoom()
	b.Run()

	//./hrp boom --worker --master-host 0.0.0.0 --master-port 7966 --ignore-quit
	//./hrp boom --master --master-bind-host 0.0.0.0 --master-bind-port 7966 --master-http-address "0.0.0.0:9092"
	//127.0.0.1:9092
}
