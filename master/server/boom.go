package server

import (
	"context"
	"github.com/test-instructor/yangfan/hrp"
	"github.com/test-instructor/yangfan/hrp/pkg/boomer"
	"sync"
)

type State int

var b *MasterBoom
var mutex sync.Mutex

var (
	StateRunning State = 1
	StateStop    State = 2
)

var _ = []State{StateRunning, StateStop}

type MasterBoom struct {
	*hrp.HRPBoomer
	State    State
	OutputDB *boomer.DbOutput
}

func NewMasterBoom() *MasterBoom {
	if b == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if b == nil {
			b = new(MasterBoom)
		}
	}
	return b
}

func (b *MasterBoom) Run() {

	masterHttpAddress := "0.0.0.0:9092"
	b.HRPBoomer = hrp.NewMasterBoomerSingleton("0.0.0.0", 7966)
	ctx := b.HRPBoomer.EnableGracefulQuit(context.Background())
	go b.StartServer(ctx, masterHttpAddress)
	go StartGrpc("0.0.0.0:9093")
	go b.HRPBoomer.PollTestCasesPlatform(ctx)
	b.HRPBoomer.RunMaster()
}

func RunHrpBoomerMaster() {

}
