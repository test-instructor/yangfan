package request

import (
	"github.com/test-instructor/yangfan/server/model/common/request"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
)

type TestCaseSearch struct {
	interfacecase.ApiCaseStep
	request.PageInfo
	FrontCase bool                  `json:"front_case,omitempty" form:"front_case"`
	ApiType   interfacecase.ApiType `json:"type,omitempty" form:"type"`
}
