package datawarehouse

import "github.com/test-instructor/yangfan/server/v2/service"

type ApiGroup struct{ DataCategoryManagementApi }

var dcmService = service.ServiceGroupApp.DatawarehouseServiceGroup.DataCategoryManagementService
