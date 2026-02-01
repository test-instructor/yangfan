package request

import (
	"time"

	"github.com/test-instructor/yangfan/server/v2/model/common/request"
)

type AutoCaseSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	CaseName       *string     `json:"caseName" form:"caseName"`
	Status         string      `json:"status" form:"status"`
	EnvName        *string     `json:"envName" form:"envName"`
	ConfigName     *string     `json:"configName" form:"configName"`
	Type           *string     `json:"type" form:"type"`
	Menu           *int64      `json:"menu" form:"menu"`
	ProjectId      uint        `json:"projectId" form:"projectId"`
	Sort           string      `json:"sort" form:"sort"`
	Order          string      `json:"order" form:"order"`
	request.PageInfo
}

type AutoCaseStepReq struct {
	CaseID uint `json:"caseId" form:"caseId"`
	StepID uint `json:"stepId" form:"stepId"`
}

type AutoCaseStepSort struct {
	Data []AutoCaseStepListUpdate `json:"data" form:"data"`
}

type AutoCaseStepListUpdate struct {
	ID   uint `json:"id" form:"id"`
	Sort uint `json:"sort" form:"sort"`
}

type SetStepConfigReq struct {
	ID           uint `json:"ID" form:"ID"`
	IsConfig     bool `json:"isConfig" form:"isConfig"`
	IsStepConfig bool `json:"isStepConfig" form:"isStepConfig"`
	ProjectId    uint `json:"projectId" form:"projectId"`
}
