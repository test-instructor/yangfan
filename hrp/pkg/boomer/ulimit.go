//go:build !windows

package boomer

import (
	"github.com/test-instructor/yangfan/server/global"
	"go.uber.org/zap"
	"syscall"
)

// set resource limit
// ulimit -n 10240
func SetUlimit(limit uint64) {
	var rLimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		global.GVA_LOG.Error("failed to convert data", zap.Error(err))
		return
	}
	global.GVA_LOG.Info("set ulimit", zap.Uint64("limit", limit))
	if rLimit.Cur >= limit {
		return
	}

	rLimit.Cur = limit
	rLimit.Max = limit
	global.GVA_LOG.Info("set current ulimit", zap.Uint64("limit", rLimit.Cur))
	err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		global.GVA_LOG.Error("failed to convert data", zap.Error(err))
		return
	}
}
