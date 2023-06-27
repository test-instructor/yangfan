package request

import (
	"github.com/test-instructor/yangfan/server/model/common/request"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
)

type ApiCaseSearch struct {
	interfacecase.ApiCase
	request.PageInfo
	FrontCase bool `json:"front_case,omitempty" form:"front_case"`
}
