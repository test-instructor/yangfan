package config

import (
	"fmt"
	"strings"
)

type Mysql struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (m *Mysql) Dsn() string {
	cfg := ensureMysqlKV(ensureMysqlKV(ensureMysqlKV(m.Config, "timeout", "10s"), "readTimeout", "30s"), "writeTimeout", "30s")
	query := ""
	if strings.TrimSpace(cfg) != "" {
		query = "?" + cfg
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s", m.Username, m.Password, m.Path, m.Port, m.Dbname, query)
}

func ensureMysqlKV(config string, key string, value string) string {
	if key == "" || value == "" {
		return config
	}
	if strings.Contains(config, key+"=") {
		return config
	}
	if strings.TrimSpace(config) == "" {
		return key + "=" + value
	}
	return config + "&" + key + "=" + value
}
