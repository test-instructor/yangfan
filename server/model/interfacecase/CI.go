package interfacecase

import (
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/system"
)

type ApiReportCI struct {
	global.GVA_MODEL
	Key       string         `json:"key" yaml:"key"`
	TagID     uint           `json:"tag" form:"tag"`
	TaskID    uint           `json:"task" form:"task"`
	EnvID     uint           `json:"env" form:"env"`
	ReportID  uint           `json:"report_id" form:"report_id"`
	Report    *ApiReport     `json:"report"`
	Project   system.Project `json:"-"`
	ProjectID uint           `json:"project_id,omitempty" yaml:"project_id"`
}
