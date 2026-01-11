package config

type DataWarehouse struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"` // 数据仓库服务器IP
	Port int    `mapstructure:"port" json:"port" yaml:"port"` // 数据仓库端口
}
