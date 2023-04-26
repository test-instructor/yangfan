package boomer

import (
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

func (b *B) run() {
	b.Boom = hrp.NewMasterBoomer("0.0.0.0", 7966)
}

func RunHrpBoomerMaster() {

}
