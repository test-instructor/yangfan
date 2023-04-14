package interfacecase

import (
	"gorm.io/datatypes"

	"github.com/test-instructor/yangfan/server/global"
)

type ApiEnv struct {
	global.GVA_MODEL
	Operator
	Default bool   `json:"default" form:"default" gorm:"column:default;comment:默认环境;"`
	Key     string `json:"key" form:"key"`
	Name    string `json:"name" form:"name"`
	Remarks string `json:"remarks"`
}

type ApiEnvDetail struct {
	global.GVA_MODEL
	Operator
	Key   string            `json:"key"  form:"key"`
	Name  string            `json:"name" form:"name"`
	Value datatypes.JSONMap `json:"value"`
}
