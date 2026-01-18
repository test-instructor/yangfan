package mq

import (
	"context"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	runtimer "github.com/test-instructor/yangfan/run/internal/timer"
	"github.com/test-instructor/yangfan/server/v2/global"
	servermq "github.com/test-instructor/yangfan/server/v2/utils/mq"
	"go.uber.org/zap"
)

type TimerTaskControlConsumer struct {
	client    *servermq.MQClient
	nodeName  string
	scheduler *runtimer.TimerTaskScheduler
	stopChan  chan struct{}
}

func NewTimerTaskControlConsumer(client *servermq.MQClient, nodeName string, scheduler *runtimer.TimerTaskScheduler) *TimerTaskControlConsumer {
	return &TimerTaskControlConsumer{
		client:    client,
		nodeName:  nodeName,
		scheduler: scheduler,
		stopChan:  make(chan struct{}),
	}
}

func (c *TimerTaskControlConsumer) StartListen() error {
	queueName := servermq.TimerTaskControlQueueName(global.GVA_CONFIG.MQ.QueuePrefix, c.nodeName)
	if queueName == "" {
		return fmt.Errorf("invalid timer control queue name")
	}

	_, err := c.client.CreateQueue(queueName, global.GVA_CONFIG.MQ.Durable, global.GVA_CONFIG.MQ.AutoDelete)
	if err != nil {
		return fmt.Errorf("failed to create queue: %w", err)
	}

	err = c.client.CreateExchange(global.GVA_CONFIG.MQ.Exchange, "direct", global.GVA_CONFIG.MQ.Durable, global.GVA_CONFIG.MQ.AutoDelete)
	if err != nil {
		return fmt.Errorf("failed to create exchange: %w", err)
	}

	err = c.client.BindQueue(queueName, queueName, global.GVA_CONFIG.MQ.Exchange)
	if err != nil {
		return fmt.Errorf("failed to bind queue: %w", err)
	}

	msgs, err := c.client.Consume(queueName, queueName)
	if err != nil {
		return fmt.Errorf("failed to start consuming: %w", err)
	}

	global.GVA_LOG.Info("Started listening for timer control", zap.String("queue", queueName))

	go func() {
		for {
			select {
			case msg, ok := <-msgs:
				if !ok {
					global.GVA_LOG.Warn("Timer control MQ channel closed")
					return
				}
				c.handleMessage(msg)
			case <-c.stopChan:
				global.GVA_LOG.Info("Stopping timer control consumer")
				return
			}
		}
	}()

	return nil
}

func (c *TimerTaskControlConsumer) handleMessage(msg amqp.Delivery) {
	var m servermq.TimerTaskControlMessage
	if err := json.Unmarshal(msg.Body, &m); err != nil {
		global.GVA_LOG.Error("Failed to unmarshal timer control message", zap.Error(err))
		_ = msg.Ack(false)
		return
	}

	switch m.Action {
	case "upsert":
		if c.scheduler != nil {
			c.scheduler.UpsertTask(context.Background(), m.TaskID)
		}
	case "delete":
		if c.scheduler != nil {
			c.scheduler.RemoveTask(m.TaskID)
		}
	default:
		global.GVA_LOG.Warn("Unknown timer control action", zap.String("action", m.Action))
	}

	_ = msg.Ack(false)
}

func (c *TimerTaskControlConsumer) Stop() {
	close(c.stopChan)
}
