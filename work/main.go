package main

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/test-instructor/yangfan/server/hrp"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

type Agent struct {
	Boom *hrp.HRPBoomer
}

func NewAgent() *Agent {
	return &Agent{}
}

func (a *Agent) Work() {
	a.Boom = hrp.NewWorkerBoomer("0.0.0.0", 7966)
	ctx := a.Boom.EnableGracefulQuit(context.Background())
	//go a.Boom.PollTestCases(ctx)
	go a.Boom.PollTasks(ctx)
	a.Boom.RunWorker()
}

func main() {
	zerolog.SetGlobalLevel(zerolog.WarnLevel)
	NewAgent().Work()
}
