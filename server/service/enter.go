package service

import (
	"github.com/test-instructor/yangfan/server/v2/service/automation"
	"github.com/test-instructor/yangfan/server/v2/service/datawarehouse"
	"github.com/test-instructor/yangfan/server/v2/service/example"
	"github.com/test-instructor/yangfan/server/v2/service/performance"
	"github.com/test-instructor/yangfan/server/v2/service/platform"
	"github.com/test-instructor/yangfan/server/v2/service/projectmgr"
	"github.com/test-instructor/yangfan/server/v2/service/system"
	"github.com/test-instructor/yangfan/server/v2/service/ui"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup        system.ServiceGroup
	ExampleServiceGroup       example.ServiceGroup
	AutomationServiceGroup    automation.ServiceGroup
	PerformanceServiceGroup   performance.ServiceGroup
	UiServiceGroup            ui.ServiceGroup
	PlatformServiceGroup      platform.ServiceGroup
	ProjectmgrServiceGroup    projectmgr.ServiceGroup
	DatawarehouseServiceGroup datawarehouse.ServiceGroup
}
