package request

import (
	"github.com/test-instructor/yangfan/server/v2/model/common/request"
	"github.com/test-instructor/yangfan/server/v2/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
