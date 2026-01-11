package request

import (
	"time"

	"github.com/test-instructor/yangfan/server/v2/model/common/request"
)

type RunConfigSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	Name           *string     `json:"name" form:"name"`
	ProjectId      int64       `json:"projectId" form:"projectId"`
	request.PageInfo
}
