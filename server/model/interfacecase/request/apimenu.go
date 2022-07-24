package request

import (
	"github.com/test-instructor/cheetah/server/model/common/request"
	"github.com/test-instructor/cheetah/server/model/interfacecase"
)

type ApiMenuSearch struct {
	interfacecase.ApiMenu
	request.PageInfo
}
