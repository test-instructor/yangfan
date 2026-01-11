package request

import (
	"github.com/test-instructor/yangfan/server/v2/model/common/request"
	"time"
)

type EnvSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	ProjectId      *int        `json:"projectId" form:"projectId"`
	request.PageInfo
}
