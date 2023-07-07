package wiki

import (
	"github.com/test-instructor/yangfan/server/global"
	"go.uber.org/zap"

	"github.com/test-instructor/yangfan/hrp/internal/myexec"
	"github.com/test-instructor/yangfan/hrp/internal/sdk"
)

func OpenWiki() error {
	sdk.SendEvent(sdk.EventTracking{
		Category: "OpenWiki",
		Action:   "hrp wiki",
	})
	global.GVA_LOG.Info("open wiki", zap.String("url", openCmd))
	return myexec.RunCommand(openCmd, "https://httprunner.com")
}
