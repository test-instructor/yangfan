package request

import (
	"time"

	"github.com/test-instructor/yangfan/server/v2/model/common/request"
)

type CategoryMenuSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	request.PageInfo
	ProjectId uint   `json:"projectId" form:"projectId"`
	MenuType  string `json:"menuType" form:"menuType"`
}
