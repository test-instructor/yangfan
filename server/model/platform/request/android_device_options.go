package request

import (
	"github.com/test-instructor/yangfan/server/v2/model/common/request"
	"time"
)

type AndroidDeviceOptionsSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	Name           *string     `json:"name" form:"name"`
	request.PageInfo
	ProjectId uint   `json:"projectId" form:"projectId"`
	Sort      string `json:"sort" form:"sort"`
	Order     string `json:"order" form:"order"`
}
