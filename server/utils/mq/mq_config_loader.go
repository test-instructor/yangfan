package mq

import (
	"errors"
	"fmt"

	"github.com/test-instructor/yangfan/server/v2/config"
)

type MQConfigLoader struct{}

func NewMQConfigLoader() *MQConfigLoader {
	return &MQConfigLoader{}
}

func (l *MQConfigLoader) LoadAndValidate(cfg config.MQ) error {
	if cfg.Type == "" {
		// Empty type means MQ is disabled, which is valid
		return nil
	}

	if cfg.Type != "rabbitmq" && cfg.Type != "rocketmq" && cfg.Type != "kafka" {
		return fmt.Errorf("unsupported MQ type: %s", cfg.Type)
	}

	if cfg.Host == "" {
		return errors.New("MQ host is required")
	}

	if cfg.Port <= 0 || cfg.Port > 65535 {
		return fmt.Errorf("invalid MQ port: %d", cfg.Port)
	}

	if cfg.Type == "rabbitmq" {
		if cfg.VirtualHost == "" {
			return errors.New("RabbitMQ virtual-host is required")
		}
	}

	return nil
}
