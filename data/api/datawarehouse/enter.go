package datawarehouse

import "github.com/test-instructor/yangfan/data/service"

type ApiGroup struct{ DataQueryApi }

var dataQueryService = service.ServiceGroupApp.DatawarehouseServiceGroup.DataQueryService
