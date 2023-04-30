package main

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/test-instructor/yangfan/hrp"
	"github.com/test-instructor/yangfan/hrp/pkg/boomer"
	"sync"
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

func main() {
	zerolog.SetGlobalLevel(zerolog.WarnLevel)
	NewAgent("0.0.0.0", 7966).Work()
}
