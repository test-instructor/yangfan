package request

import (
	"github.com/test-instructor/yangfan/server/v2/model/common/request"
)

type ExaAttachmentCategorySearch struct {
	ClassId int `json:"classId" form:"classId"`
	request.PageInfo
}
