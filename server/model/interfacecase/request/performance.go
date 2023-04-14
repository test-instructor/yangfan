package request

import (
	"github.com/test-instructor/yangfan/server/model/common/request"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
)

type PerformancekSearch struct {
	interfacecase.Performance
	request.PageInfo
}
