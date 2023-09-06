package interfacecase

import (
	"gorm.io/datatypes"

	"github.com/test-instructor/yangfan/server/global"
)

type ApiEnv struct {
	global.GVA_MODEL
	Operator
	Default bool   `json:"default,omitempty" form:"default" gorm:"column:default;comment:默认环境;"`
	Key     string `json:"key,omitempty" form:"key"`
	Name    string `json:"name,omitempty" form:"name"`
	Remarks string `json:"remarks"`
}

type ApiEnvDetail struct {
	global.GVA_MODEL
	Operator
	Key   string            `json:"key"  form:"key"`
	Name  string            `json:"name,omitempty" form:"name"`
	Value datatypes.JSONMap `json:"value"`
}

type ApiEnvMock struct {
	global.GVA_MODEL
	Operator
	Name       string `json:"name,omitempty" form:"name"`
	Url        string `json:"url,omitempty" form:"url"`
	StatusCode int    `json:"status_code,omitempty" form:"status_code"`
}
