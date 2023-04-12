package api

import "github.com/test-instructor/yangfan/server/plugin/fslogin/passport"

type ApiGroup struct {
	FsLoginApi
}

var ApiGroupApp = new(ApiGroup)

var (
	fsLoginPassPort = new(passport.FsLoginPassport)
)
