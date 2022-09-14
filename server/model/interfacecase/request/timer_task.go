package request

import (
	"github.com/test-instructor/cheetah/server/model/common/request"
	"github.com/test-instructor/cheetah/server/model/interfacecase"
)

type TimerTaskSearch struct {
	interfacecase.ApiTimerTask
	request.PageInfo
}

type SetTimerCares struct {
	ID      uint
	CaseIds []uint `json:"caseIds"` // 角色ID
}
