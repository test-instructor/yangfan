package request

import (
	"time"

	"github.com/test-instructor/yangfan/server/v2/model/common/request"
)

type PythonCodeDebugSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	Function       string      `json:"function" form:"function"`
	ProjectId      uint64      `json:"projectId" form:"projectId"`
	request.PageInfo
}
