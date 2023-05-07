package wiki

import (
	"github.com/rs/zerolog/log"

	"github.com/test-instructor/yangfan/hrp/internal/myexec"
	"github.com/test-instructor/yangfan/hrp/internal/sdk"
)

func OpenWiki() error {
	sdk.SendEvent(sdk.EventTracking{
		Category: "OpenWiki",
		Action:   "hrp wiki",
	})
	log.Info().Msgf("%s https://httprunner.com", openCmd)
	return myexec.RunCommand(openCmd, "https://httprunner.com")
}
