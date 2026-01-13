package hrp

import (
	"github.com/test-instructor/yangfan/server/v2/global"
	"go.uber.org/zap"
)

// hrpLogger returns a zap.Logger for hrp internal logging.
//
// If the main server has initialized global.GVA_LOG (via zap.ReplaceGlobals),
// we reuse it so that hrp logs are unified with the rest of the system logs.
// Otherwise we fall back to zap.L(), which is a no-op logger by default,
// avoiding nil-pointer panics in standalone tests.
func hrpLogger() *zap.Logger {
	if global.GVA_LOG != nil {
		return global.GVA_LOG
	}
	return zap.L()
}
