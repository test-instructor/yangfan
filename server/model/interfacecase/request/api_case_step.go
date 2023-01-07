package request

import (
	"github.com/test-instructor/cheetah/server/model/common/request"
	"github.com/test-instructor/cheetah/server/model/interfacecase"
)

type TestCaseSearch struct {
	interfacecase.ApiCaseStep
	request.PageInfo
	FrontCase bool                  `json:"front_case" form:"front_case"`
	ApiType   interfacecase.ApiType `json:"type" form:"type"`
}
