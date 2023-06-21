package main

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/test-instructor/yangfan/hrp"
	"github.com/test-instructor/yangfan/hrp/pkg/boomer"
	"github.com/test-instructor/yangfan/server/core"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/grpc/client"
	"github.com/test-instructor/yangfan/server/grpc/pkg"
	"github.com/test-instructor/yangfan/server/initialize"
	"github.com/test-instructor/yangfan/server/source/yangfan"
	"go.uber.org/zap"
	"math/rand"
	"os"
	"sync"
	"time"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

var mutex sync.Mutex

type Agent struct {
	Boom *hrp.HRPBoomer
}

var agent *Agent

func NewAgent(masterBindHost string, masterBindPort int) *Agent {
	if agent == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if agent == nil {
			agent = &Agent{
				Boom: hrp.NewWorkerBoomer(masterBindHost, masterBindPort),
			}
		}
	}
	return agent
}

func (a *Agent) Work() {
	a.Boom.SetProfile(&boomer.Profile{})
	ctx := a.Boom.EnableGracefulQuit(context.Background())
	a.Boom.SetIgnoreQuit()
	go a.Boom.PollTasks(ctx)
	a.Boom.RunWorker()
}

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
	global.HrpMode = global.HrpModeWork
	global.GVA_VP = core.Viper()
	global.GVA_LOG = core.Zap()
	global.GVA_DB = initialize.Gorm()
	if global.GVA_DB.Error != nil {
		global.GVA_LOG.Error("register db", zap.Error(global.GVA_DB.Error))
		os.Exit(0)
	}
	boomer.NewTempConfig()
	zerolog.SetGlobalLevel(zerolog.WarnLevel)
	zap.ReplaceGlobals(global.GVA_LOG)
	go yangfan.InitPythonPackage(true)
	go RunPkgInstallClient()
	NewAgent(global.GVA_CONFIG.YangFan.Master, global.GVA_CONFIG.YangFan.MasterServerProt).Work()
}
