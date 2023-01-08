package v1

import (
	"github.com/test-instructor/cheetah/server/api/v1/example"
	"github.com/test-instructor/cheetah/server/api/v1/interfacecase"

	"github.com/test-instructor/cheetah/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup        system.ApiGroup
	ExampleApiGroup       example.ApiGroup
	InterfaceCaseApiGroup interfacecase.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
