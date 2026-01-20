package mq

import (
	"fmt"

	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/test-instructor/yangfan/run/internal/runTestCase"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/platform/request"
	"github.com/test-instructor/yangfan/server/v2/utils/mq"
	"go.uber.org/zap"
)

type RunnerTaskConsumer struct {
	client   *mq.MQClient
	nodeName string
	stopChan chan struct{}
}

func NewRunnerTaskConsumer(client *mq.MQClient, nodeName string) *RunnerTaskConsumer {
	return &RunnerTaskConsumer{
		client:   client,
		nodeName: nodeName,
		stopChan: make(chan struct{}),
	}
}

func (c *RunnerTaskConsumer) StartListen() error {
	// 1. Determine Queue Name
	queueName := c.nodeName
	if global.GVA_CONFIG.MQ.QueuePrefix != "" {
		queueName = global.GVA_CONFIG.MQ.QueuePrefix + c.nodeName
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

	// 4. Start Consuming
	msgs, err := c.client.Consume(queueName, c.nodeName)
	if err != nil {
		return fmt.Errorf("failed to start consuming: %w", err)
	}

	global.GVA_LOG.Info("Started listening for tasks", zap.String("queue", queueName))

	go func() {
		for {
			select {
			case msg, ok := <-msgs:
				if !ok {
					global.GVA_LOG.Warn("MQ channel closed")
					return
				}
				c.handleMessage(msg)
			case <-c.stopChan:
				global.GVA_LOG.Info("Stopping consumer")
				return
			}
		}
	}()

	return nil
}

func (c *RunnerTaskConsumer) handleMessage(msg amqp.Delivery) {
	fmt.Println("Received MQ message:", string(msg.Body))

	// 先解析 TaskMessage 获取完整信息（包含 ReportID）
	var taskMsg mq.TaskMessage
	err := json.Unmarshal(msg.Body, &taskMsg)
	if err != nil {
		global.GVA_LOG.Error("Failed to unmarshal message", zap.Error(err))
		_ = msg.Ack(false)
		return
	}

	// 构建 RunnerRequest
	runnerReq := request.RunnerRequest{
		CaseType:         taskMsg.CaseType,
		CaseID:           taskMsg.CaseID,
		ReportID:         taskMsg.ReportID,
		RunMode:          taskMsg.RunMode,
		NodeName:         taskMsg.NodeName,
		EnvID:            taskMsg.EnvID,
		ConfigID:         taskMsg.ConfigID,
		ProjectId:        taskMsg.ProjectID,
		NotifyEnabled:    taskMsg.NotifyEnabled,
		NotifyRule:       taskMsg.NotifyRule,
		NotifyChannelIDs: taskMsg.NotifyChannelIDs,
		Failfast:         taskMsg.Failfast,
	}

	msgBody := string(msg.Body)

	_, err = runTestCase.Entry(runnerReq, &msgBody)
	if err != nil {
		global.GVA_LOG.Error("Failed to run test case", zap.Error(err))
	}

	_ = msg.Ack(false)
}

func (c *RunnerTaskConsumer) Stop() {
	close(c.stopChan)
}
