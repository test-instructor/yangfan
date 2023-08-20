package router

import (
	"github.com/test-instructor/yangfan/server/router/example"
	"github.com/test-instructor/yangfan/server/router/interfacecase"
	"github.com/test-instructor/yangfan/server/router/system"
)

type RouterGroup struct {
	System        system.RouterGroup
	Example       example.RouterGroup
	Interfacecase interfacecase.RouterGroup
	RunCase       interfacecase.RunCaseRouterGroup
	ApiConfig     interfacecase.ApiConfigRouterGroup
	Report        interfacecase.ReportRouter
	ApiCase       interfacecase.ApiCaseGroup
	TimerTask     interfacecase.TimerTaskGroup
	Performance   interfacecase.PerformanceRouterGroup
	Environment   interfacecase.EnvironmentRouterGroup
	Message       interfacecase.MessageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
