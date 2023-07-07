package sdk

import (
	"fmt"

	"github.com/denisbrodbeck/machineid"
	"github.com/getsentry/sentry-go"
	uuid "github.com/satori/go.uuid"
	"github.com/test-instructor/yangfan/server/global"
	"go.uber.org/zap"

	"github.com/test-instructor/yangfan/hrp/internal/env"
	"github.com/test-instructor/yangfan/hrp/internal/version"
)

const (
	trackingID = "UA-114587036-1" // Tracking ID for Google Analytics
	sentryDSN  = "https://cff5efc69b1a4325a4cf873f1e70c13a@o334324.ingest.sentry.io/6070292"
)

var gaClient *GAClient

func init() {
	// init GA client
	clientID, err := machineid.ProtectedID("hrp")
	if err != nil {
		clientID = uuid.NewV1().String()
	}
	gaClient = NewGAClient(trackingID, clientID)

	// init sentry sdk
	if env.DISABLE_SENTRY == "true" {
		return
	}
	err = sentry.Init(sentry.ClientOptions{
		Dsn:              sentryDSN,
		Release:          fmt.Sprintf("httprunner@%s", version.VERSION),
		AttachStacktrace: true,
	})
	if err != nil {
		global.GVA_LOG.Error("init sentry sdk failed!", zap.Error(err))
		return
	}
	sentry.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetLevel(sentry.LevelError)
		scope.SetUser(sentry.User{
			ID: clientID,
		})
	})
}

func SendEvent(e IEvent) error {
	if env.DISABLE_GA == "true" {
		// do not send GA events in CI environment
		return nil
	}
	return gaClient.SendEvent(e)
}
