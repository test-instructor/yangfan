package request

import (
	"github.com/test-instructor/yangfan/server/model/{{.Package}}"
	"github.com/test-instructor/yangfan/server/model/common/request"
)

type {{.StructName}}Search struct{
    {{.Package}}.{{.StructName}}
    request.PageInfo
}
