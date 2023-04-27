package main

import (
	"github.com/test-instructor/yangfan/server/core"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/master/boomer"
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
	//global.GVA_VP = core.Viper()
	global.GVA_LOG = core.Zap()
	zap.ReplaceGlobals(global.GVA_LOG)
	boomer.NewB().Run()
}
