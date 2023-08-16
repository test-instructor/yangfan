package request

import (
	"github.com/test-instructor/yangfan/server/model/common/request"
	"github.com/test-instructor/yangfan/server/model/system"
)

type ProjectSearch struct {
	system.Project
	request.PageInfo
}

type SysProjectUsers struct {
	ProjectId uint `json:"projectId" form:"projectId"`
	request.PageInfo
}
