package projectmgr

import (
	"time"

	"github.com/test-instructor/yangfan/server/v2/global"
	"gorm.io/datatypes"
)

type ReportNotifyReportResult string

const (
	ReportNotifyReportResultSuccess ReportNotifyReportResult = "success"
	ReportNotifyReportResultFail    ReportNotifyReportResult = "fail"
)

type ProjectReportNotifyLog struct {
	global.GVA_MODEL
	ProjectId      int64                    `json:"projectId" form:"projectId" gorm:"column:project_id;uniqueIndex:uniq_report_channel_result,priority:1"`
	ReportId       uint                     `json:"report_id" form:"report_id" gorm:"column:report_id;uniqueIndex:uniq_report_channel_result,priority:2"`
	ChannelId      uint                     `json:"channel_id" form:"channel_id" gorm:"column:channel_id;uniqueIndex:uniq_report_channel_result,priority:3"`
	Provider       ReportNotifyProvider     `json:"provider" form:"provider" gorm:"column:provider;type:varchar(16)"`
	ReportResult   ReportNotifyReportResult `json:"report_result" form:"report_result" gorm:"column:report_result;type:varchar(16);uniqueIndex:uniq_report_channel_result,priority:4"`
	Ok             bool                     `json:"ok" form:"ok" gorm:"column:ok"`
	Error          string                   `json:"error" form:"error" gorm:"column:error;type:text"`
	SentAt         time.Time                `json:"sent_at" form:"sent_at" gorm:"column:sent_at"`
	RequestPayload datatypes.JSON           `json:"request_payload" form:"request_payload" gorm:"column:request_payload" swaggertype:"object"`
}

func (ProjectReportNotifyLog) TableName() string {
	return "yf_project_report_notify_logs"
}
