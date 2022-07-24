package global

{{- if .HasGlobal }}

import "github.com/test-instructor/cheetah/server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}