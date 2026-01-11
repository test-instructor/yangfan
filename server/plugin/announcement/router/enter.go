package router

import "github.com/test-instructor/yangfan/server/v2/plugin/announcement/api"

var (
	Router  = new(router)
	apiInfo = api.Api.Info
)

type router struct{ Info info }
