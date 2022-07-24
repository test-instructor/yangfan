package request

import (
	"github.com/test-instructor/cheetah/server/model/common/request"
	"github.com/test-instructor/cheetah/server/model/system"
)

type ProjectSearch struct {
	system.Project
	request.PageInfo
}
