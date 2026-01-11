package config

type Runner struct {
	NodeName string `mapstructure:"node-name" json:"nodeName" yaml:"node-name"` // 节点名称
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`               // 服务端口
}
