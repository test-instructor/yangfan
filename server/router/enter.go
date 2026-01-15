package router

import (
	"github.com/test-instructor/yangfan/server/v2/router/automation"
	"github.com/test-instructor/yangfan/server/v2/router/datawarehouse"
	"github.com/test-instructor/yangfan/server/v2/router/example"
	"github.com/test-instructor/yangfan/server/v2/router/performance"
	"github.com/test-instructor/yangfan/server/v2/router/platform"
	"github.com/test-instructor/yangfan/server/v2/router/projectmgr"
	"github.com/test-instructor/yangfan/server/v2/router/system"
	"github.com/test-instructor/yangfan/server/v2/router/ui"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System           system.RouterGroup
	Example          example.RouterGroup
	Automation       automation.RouterGroup
	Performance      performance.RouterGroup
	Ui               ui.RouterGroup
	Platform         platform.RouterGroup
	Projectmgr       projectmgr.RouterGroup
	Datawarehouse    datawarehouse.RouterGroup
	OpenRunnerRouter platform.OpenRunnerRouter
}
