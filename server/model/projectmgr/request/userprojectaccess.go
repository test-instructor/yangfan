package request

import (
	"github.com/test-instructor/yangfan/server/v2/model/common/request"
	"time"
)

type UserProjectAccessSearch struct {
	CreatedAtRange   []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	UserId           *int        `json:"userId" form:"userId"`
	ProjectId        *int        `json:"projectId" form:"projectId"`
	AccessPermission *bool       `json:"accessPermission" form:"accessPermission"`
	request.PageInfo
}
