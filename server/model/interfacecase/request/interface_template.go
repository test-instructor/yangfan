package request

import (
	"github.com/test-instructor/yangfan/server/model/common/request"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
)

type InterfaceTemplateSearch struct {
	interfacecase.ApiStep
	request.PageInfo
	ApiType int `json:"type" form:"type"`
}

type InterfaceTemplateApi struct {
	Name   string
	Method string
}

type InterfaceTemplateList struct {
	ID   uint
	Name string
	//Request InterfaceTemplateApi `json:"request"`
}
