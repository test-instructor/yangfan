package request

import (
	"github.com/test-instructor/yangfan/server/v2/model/common/request"
	"time"
)

type RequestSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	request.PageInfo
	ProjectId uint `json:"projectId" form:"projectId"`
}
