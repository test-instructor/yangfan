package request

import (
	"github.com/test-instructor/cheetah/server/model/common/request"
	"github.com/test-instructor/cheetah/server/model/interfacecase"
)

type PerformancekSearch struct {
	interfacecase.Performance
	request.PageInfo
}
