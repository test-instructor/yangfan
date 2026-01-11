package request

import (
	"time"

	"github.com/test-instructor/yangfan/server/v2/model/common/request"
)

type AutoCaseStepSearch struct {
	request.PageInfo
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	StepName       *string     `json:"name" form:"name"`
	ProjectId      uint        `json:"projectId" form:"projectId"`
	Sort           string      `json:"sort" form:"sort"`
	Order          string      `json:"order" form:"order"`
	Menu           uint        `json:"menu" form:"menu"`
}

type AutoCaseStepSearchApi struct {
	ID    uint                         `json:"id" form:"id"`         // 步骤集合ID
	ApiID uint                         `json:"api_id" form:"api_id"` // 步骤ID（API、UI等步骤）
	Sort  uint                         `json:"sort" form:"sort"`
	Data  []AutoCaseStepRelationUpdate `json:"data" form:"data"`
	AutoCaseStepSearch
}

type AutoCaseStepRelationUpdate struct {
	ID   uint `json:"id" form:"id"`
	Sort uint `json:"sort" form:"sort"`
}
