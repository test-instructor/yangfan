package mq

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/test-instructor/yangfan/server/v2/config"
	"go.uber.org/zap"
)

type TimerTaskControlProducer struct {
	client *MQClient
	cfg    config.MQ
	logger *zap.Logger
}

func NewTimerTaskControlProducer(client *MQClient, cfg config.MQ, logger *zap.Logger) *TimerTaskControlProducer {
	return &TimerTaskControlProducer{
		client: client,
		cfg:    cfg,
		logger: logger,
	}
}

func (p *TimerTaskControlProducer) Send(action string, taskID uint, nodeName string) error {
	action = strings.ToLower(strings.TrimSpace(action))
	if action != "upsert" && action != "delete" {
		return fmt.Errorf("invalid action: %s", action)
	}
	nodeName = strings.TrimSpace(nodeName)
	if nodeName == "" {
		return fmt.Errorf("node_name is required")
	}

	msg := TimerTaskControlMessage{
		Action:   action,
		TaskID:   taskID,
		SendTime: time.Now().Unix(),
	}
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	queueName := TimerTaskControlQueueName(p.cfg.QueuePrefix, nodeName)
	if err := p.client.CreateExchange(p.cfg.Exchange, "direct", p.cfg.Durable, p.cfg.AutoDelete); err != nil {
		return err
	}
	if err := p.client.Publish(p.cfg.Exchange, queueName, body); err != nil {
		p.logger.Error("failed to publish timer control message",
			zap.String("action", action),
			zap.Uint("task_id", taskID),
			zap.String("node_name", nodeName),
			zap.Error(err))
		return err
	}
	return nil
}

func TimerTaskControlQueueName(queuePrefix, nodeName string) string {
	base := strings.TrimSpace(nodeName)
	if base == "" {
		return ""
	}
	if queuePrefix != "" {
		base = queuePrefix + base
	}
	return base + ".timer"
}
