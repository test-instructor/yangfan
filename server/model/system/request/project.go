package request

import (
	"github.com/test-instructor/yangfan/server/model/common/request"
	"github.com/test-instructor/yangfan/server/model/system"
)

type ProjectSearch struct {
	system.Project
	request.PageInfo
}
