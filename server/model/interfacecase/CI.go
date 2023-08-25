package interfacecase

import (
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/system"
	"gorm.io/datatypes"
)

type ApiReportCI struct {
	global.GVA_MODEL
	Key              string            `json:"key" yaml:"key"`
	TagID            uint              `json:"tag" form:"tag"`
	TaskID           uint              `json:"task" form:"task"`
	EnvID            uint              `json:"env" form:"env"`
	ReportID         uint              `json:"report_id" form:"report_id"`
	Report           *ApiReport        `json:"report"`
	Project          system.Project    `json:"-"`
	ProjectID        uint              `json:"project_id,omitempty" yaml:"project_id"`
	CallbackUrl      string            `json:"callback_url" form:"callback_url"`
	CallbackResponse string            `json:"callback_response" form:"callback_response" gorm:"column:callback_response;type:text;comment:回调返回内容"`
	Other            datatypes.JSONMap `json:"other" form:"other"`
}
