package service

import (
	"github.com/test-instructor/cheetah/server/service/example"
	"github.com/test-instructor/cheetah/server/service/interfacecase"
	"github.com/test-instructor/cheetah/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup        system.ServiceGroup
	ExampleServiceGroup       example.ServiceGroup
	InterfacecaseServiceGroup interfacecase.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
