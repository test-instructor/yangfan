package request

import (
	"github.com/test-instructor/yangfan/server/model/common/request"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
)

type HrpPyPkgRequest struct {
	interfacecase.HrpPyPkg
	request.PageInfo
}
