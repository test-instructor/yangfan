package request

import (
	"github.com/test-instructor/yangfan/server/v2/model/common/request"
	"github.com/test-instructor/yangfan/server/v2/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
