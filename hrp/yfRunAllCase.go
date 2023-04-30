package hrp

import (
	"github.com/test-instructor/yangfan/server/model/interfacecase"
)

type ToTestCase struct {
	Config    interfacecase.ApiConfig
	TestSteps []interface{}
}
