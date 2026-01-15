package request

import "github.com/test-instructor/yangfan/server/v2/model/common/request"

type ReportNotifyChannelSearch struct {
	request.PageInfo
	ProjectId int64  `json:"projectId" form:"projectId"`
	Provider  string `json:"provider" form:"provider"`
	Enabled   *bool  `json:"enabled" form:"enabled"`
}

type ReportNotifyChannelById struct {
	ID uint `json:"ID" form:"ID"`
}

type ReportNotifyChannelIds struct {
	IDs []uint `json:"IDs" form:"IDs"`
}

type AutoReportNotifyStatusQuery struct {
	ReportId uint `json:"reportId" form:"reportId" binding:"required"`
}
