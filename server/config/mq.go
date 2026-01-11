package config

type MQ struct {
	Type          string `mapstructure:"type" json:"type" yaml:"type"`                               // MQ类型: rabbitmq, rocketmq, kafka
	Host          string `mapstructure:"host" json:"host" yaml:"host"`                               // MQ服务器地址
	Port          int    `mapstructure:"port" json:"port" yaml:"port"`                               // MQ端口
	Username      string `mapstructure:"username" json:"username" yaml:"username"`                   // 用户名
	Password      string `mapstructure:"password" json:"password" yaml:"password"`                   // 密码
	VirtualHost   string `mapstructure:"virtual-host" json:"virtual-host" yaml:"virtual-host"`       // RabbitMQ虚拟主机
	Exchange      string `mapstructure:"exchange" json:"exchange" yaml:"exchange"`                   // 交换机名称
	QueuePrefix   string `mapstructure:"queue-prefix" json:"queue-prefix" yaml:"queue-prefix"`       // 队列名称前缀
	RetryCount    int    `mapstructure:"retry-count" json:"retry-count" yaml:"retry-count"`          // 消息重试次数
	Timeout       int    `mapstructure:"timeout" json:"timeout" yaml:"timeout"`                      // 连接超时时间(秒)
	Heartbeat     int    `mapstructure:"heartbeat" json:"heartbeat" yaml:"heartbeat"`                // 心跳间隔(秒)
	PrefetchCount int    `mapstructure:"prefetch-count" json:"prefetch-count" yaml:"prefetch-count"` // 预取消息数量
	Durable       bool   `mapstructure:"durable" json:"durable" yaml:"durable"`                      // 队列是否持久化
	AutoDelete    bool   `mapstructure:"auto-delete" json:"auto-delete" yaml:"auto-delete"`          // 队列是否自动删除
}
