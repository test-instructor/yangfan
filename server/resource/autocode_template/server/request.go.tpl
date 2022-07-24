package request

import (
	"github.com/test-instructor/cheetah/server/model/{{.Package}}"
	"github.com/test-instructor/cheetah/server/model/common/request"
)

type {{.StructName}}Search struct{
    {{.Package}}.{{.StructName}}
    request.PageInfo
}
