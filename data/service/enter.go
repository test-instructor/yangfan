package service

import "github.com/test-instructor/yangfan/data/service/datawarehouse"

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	DatawarehouseServiceGroup datawarehouse.ServiceGroup
}
