package boomer

import (
	"context"

	"github.com/test-instructor/yangfan/server/hrp"
	"github.com/test-instructor/yangfan/server/hrp/pkg/boomer"
)

type BoomerState int

var (
	BoomerStateRunning BoomerState = 1
	BoomerStateStop    BoomerState = 2
)

type B struct {
	Boom     *hrp.HRPBoomer
	State    BoomerState
	OutputDB *boomer.DbOutput
}

func NewB() *B {
	return &B{}
}

func (b *B) Run() {
	masterHttpAddress := "0.0.0.0:9092"
	b.Boom = hrp.NewMasterBoomer("0.0.0.0", 7966)
	ctx := b.Boom.EnableGracefulQuit(context.Background())
	go b.Boom.StartServer(ctx, masterHttpAddress)
	go b.Boom.PollTestCases(ctx)
	b.Boom.RunMaster()
	//./hrp boom --worker --master-host 192.168.0.218 --master-port 7966 --ignore-quit
	//./hrp boom --master --master-bind-host 0.0.0.0 --master-bind-port 7966 --master-http-address "0.0.0.0:9092"
	//127.0.0.1:9092
}

func RunHrpBoomerMaster() {

}
