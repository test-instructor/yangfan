package request

import (
	"github.com/test-instructor/yangfan/server/v2/model/common/request"
	"time"
)

type TimerTaskSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	ConfigName     *string     `json:"configName" form:"configName"`
	EnvName        *string     `json:"envName" form:"envName"`
	MessageName    *string     `json:"messageName" form:"messageName"`
	Type           *string     `json:"type" form:"type"`
	request.PageInfo
	ProjectId uint `json:"projectId" form:"projectId"`
}

type TagSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	Name           *string     `json:"name" form:"name"`
	request.PageInfo
	ProjectId int64 `json:"projectId" form:"projectId"`
}

// 任务-用例关联请求
type TimerTaskCaseReq struct {
	TaskID uint `json:"taskId" form:"taskId"`
	CaseID uint `json:"caseId" form:"caseId"`
}

type TimerTaskCaseListUpdate struct {
	ID   uint `json:"id" form:"id"`
	Sort uint `json:"sort" form:"sort"`
}

type TimerTaskCaseSort struct {
	Data []TimerTaskCaseListUpdate `json:"data" form:"data"`
}
