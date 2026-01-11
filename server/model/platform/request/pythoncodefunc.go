package request

import (
	"time"

	"github.com/test-instructor/yangfan/server/v2/model/common/request"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
)

type PythonCodeFuncSearch struct {
	CreatedAtRange []time.Time               `json:"createdAtRange" form:"createdAtRange[]"`
	ProjectId      int64                     `json:"projectId" form:"projectId"`
	Data           []platform.PythonCodeFunc `json:"data" form:"data"`
	request.PageInfo
}
