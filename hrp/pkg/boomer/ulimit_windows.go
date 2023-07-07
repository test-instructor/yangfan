//go:build windows

package boomer

import (
	"github.com/test-instructor/yangfan/server/global"
)

// set resource limit
func SetUlimit(limit uint64) {
	global.GVA_LOG.Warn("windows does not support setting ulimit")
}
