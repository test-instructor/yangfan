package v1

import (
	"github.com/test-instructor/yangfan/server/v2/api/v1/automation"
	"github.com/test-instructor/yangfan/server/v2/api/v1/datawarehouse"
	"github.com/test-instructor/yangfan/server/v2/api/v1/example"
	"github.com/test-instructor/yangfan/server/v2/api/v1/performance"
	"github.com/test-instructor/yangfan/server/v2/api/v1/platform"
	"github.com/test-instructor/yangfan/server/v2/api/v1/projectmgr"
	"github.com/test-instructor/yangfan/server/v2/api/v1/system"
	"github.com/test-instructor/yangfan/server/v2/api/v1/ui"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup        system.ApiGroup
	ExampleApiGroup       example.ApiGroup
	AutomationApiGroup    automation.ApiGroup
	PerformanceApiGroup   performance.ApiGroup
	UiApiGroup            ui.ApiGroup
	PlatformApiGroup      platform.ApiGroup
	ProjectmgrApiGroup    projectmgr.ApiGroup
	DatawarehouseApiGroup datawarehouse.ApiGroup
}
