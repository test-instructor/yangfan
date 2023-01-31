package router

import (
	"github.com/test-instructor/cheetah/server/router/example"
	"github.com/test-instructor/cheetah/server/router/interfacecase"
	"github.com/test-instructor/cheetah/server/router/system"
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
	PyPkg         interfacecase.PyPkgRouter
}

var RouterGroupApp = new(RouterGroup)
