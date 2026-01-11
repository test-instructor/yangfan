package api

import "github.com/test-instructor/yangfan/server/v2/plugin/announcement/service"

var (
	Api         = new(api)
	serviceInfo = service.Service.Info
)

type api struct{ Info info }
