package request

import (
	"github.com/test-instructor/yangfan/server/model/common/request"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
)

type ReportSearch struct {
	interfacecase.ApiReport
	request.PageInfo
	ApiType int `json:"type" form:"type"`
}

type PReportSearch struct {
	interfacecase.PerformanceReport
	request.PageInfo
	ApiType int `json:"type" form:"type"`
}

type PReportDetail struct {
	ID       uint `json:"ID"`
	DetailID uint `json:"DetailID"`
}

//type InterfaceTemplateApi struct {
//	Name   string
//	Method string
//}
//
//type InterfaceTemplateList struct {
//	ID   uint
//	Name string
//	//Request InterfaceTemplateApi `json:"request"`
//}
