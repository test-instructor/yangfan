package request

import (
	"github.com/test-instructor/yangfan/server/v2/model/common/request"
	"github.com/test-instructor/yangfan/server/v2/model/system"
	"time"
)

type SysExportTemplateSearch struct {
	system.SysExportTemplate
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
