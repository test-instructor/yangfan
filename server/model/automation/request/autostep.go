package request

import (
	"time"

	"github.com/test-instructor/yangfan/server/v2/model/common/request"
)

type AutoStepSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	StepName       *string     `json:"name" form:"name"`
	ProjectId      uint        `json:"projectId" form:"projectId"`
	Type           *string     `json:"type" form:"type"`
	Sort           string      `json:"sort" form:"sort"`
	Order          string      `json:"order" form:"order"`
	Menu           *uint       `json:"menu" form:"menu"`
	StepType       *int        `json:"step_type" form:"step_type"`
	request.PageInfo
}
