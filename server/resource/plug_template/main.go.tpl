package {{ .Snake}}

import (
{{- if .HasGlobal }}
	"github.com/test-instructor/yangfan/server/plugin/{{ .Snake}}/global"
{{- end }}
	"github.com/test-instructor/yangfan/server/plugin/{{ .Snake}}/router"
	"github.com/gin-gonic/gin"
)

type {{ .PlugName}}Plugin struct {
}

func Create{{ .PlugName}}Plug({{- range .Global}} {{.Key}} {{.Type}}, {{- end }})*{{ .PlugName}}Plugin {
{{- if .HasGlobal }}
	{{- range .Global}}
	    global.GlobalConfig.{{.Key}} = {{.Key}}
	{{- end }}
{{ end }}
	return &{{ .PlugName}}Plugin{}
}

func (*{{ .PlugName}}Plugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.Init{{ .PlugName}}Router(group)
}

func (*{{ .PlugName}}Plugin) RouterPath() string {
	return "{{ .RouterGroup}}"
}
