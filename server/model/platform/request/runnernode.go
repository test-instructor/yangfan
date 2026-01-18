package request

import (
	"github.com/test-instructor/yangfan/server/v2/model/common/request"
	"time"
)

type RunnerNodeSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	request.PageInfo
	ProjectId   uint     `json:"projectId" form:"projectId"`
	RunContent  string   `json:"runContent" form:"runContent"`
	RunContents []string `json:"runContents" form:"runContents[]"`
}
