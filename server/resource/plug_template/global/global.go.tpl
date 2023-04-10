package global

{{- if .HasGlobal }}

import "github.com/test-instructor/yangfan/server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}