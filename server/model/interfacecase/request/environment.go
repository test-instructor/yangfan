package request

import (
	"github.com/test-instructor/yangfan/server/model/common/request"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
)

type EnvSearch struct {
	interfacecase.ApiEnv
	request.PageInfo
}

type EnvVariableSearch struct {
	ShowKey bool `json:"show_key,omitempty" form:"show_key"`
	interfacecase.ApiEnvDetail
	request.PageInfo
}
