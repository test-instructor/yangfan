package config

type Grpc struct {
	Background               string `mapstructure:"background" json:"background" yaml:"background"`
	BackgroundGrpcPort       string `mapstructure:"background-grpc-port" json:"background-grpc-port" yaml:"background-grpc-port"`
	Master                   string `mapstructure:"master" json:"master" yaml:"master"`
	MasterBoomerProt         string `mapstructure:"master-boomer-prot" json:"master-boomer-prot" yaml:"master-boomer-prot"`
	MasterServerProt         int    `mapstructure:"master-server-prot" json:"master-server-prot" yaml:"master-server-prot"`
	PrometheusPushgatewayURL string `mapstructure:"prometheus-pushgateway-url" json:"prometheus-pushgateway-url" yaml:"prometheus-pushgateway-url"`
}
