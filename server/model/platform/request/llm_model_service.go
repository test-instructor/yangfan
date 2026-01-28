package request

import (
	"github.com/test-instructor/yangfan/server/v2/model/common/request"
	"time"
)

type LLMModelConfigSearch struct {
	CreatedAtRange  []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	Name            *string     `json:"name" form:"name"`
	RequestSchema   *string     `json:"requestSchema" form:"requestSchema"`
	Model           *string     `json:"model" form:"model"`
	ReasoningEffort *string     `json:"reasoningEffort" form:"reasoningEffort"`
	Enabled         *bool       `json:"enabled" form:"enabled"`
	request.PageInfo
	ProjectId uint `json:"projectId" form:"projectId"`
}
