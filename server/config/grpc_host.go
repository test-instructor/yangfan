package config

type Grpc struct {
	Background string `mapstructure:"background" json:"background" yaml:"background"`
	Master     string `mapstructure:"master" json:"master" yaml:"master"`
}
