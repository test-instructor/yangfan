package request

import (
	"github.com/test-instructor/yangfan/server/v2/model/common/request"
	"time"
)

type ProjectSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	Name           *string     `json:"name" form:"name"`
	request.PageInfo
}

type ResetProjectAuthReq struct {
	ID uint `json:"ID" binding:"required"`
}
