package request

import (
	"time"

	"github.com/test-instructor/yangfan/server/v2/model/common/request"
)

type PythonCodeSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	Type           int64       `json:"type" form:"type"`
	ProjectId      uint        `json:"projectId" form:"projectId"`
	ID             *uint       `json:"id" form:"id"`
	UniqueKey      string      `json:"uniqueKey" form:"uniqueKey"`
	request.PageInfo
}
