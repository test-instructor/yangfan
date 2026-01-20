package mq

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/test-instructor/yangfan/server/v2/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type RunnerTaskProducer struct {
	client *MQClient
	db     *gorm.DB
	cfg    config.MQ
	logger *zap.Logger
}

func NewRunnerTaskProducer(client *MQClient, db *gorm.DB, cfg config.MQ, logger *zap.Logger) *RunnerTaskProducer {
	return &RunnerTaskProducer{
		client: client,
		db:     db,
		cfg:    cfg,
		logger: logger,
	}
}

// SendTask 发送任务
// nodeName: 指定节点名称 (e.g. "run1"), 如果为空则随机选择
// reportID: 测试报告ID
// Returns: (selectedNodeName, error)
func (p *RunnerTaskProducer) SendTask(taskID string, reportID uint, caseType string, caseID uint, runMode string, nodeName string, envID int, configID int, projectID uint, notifyEnabled *bool, notifyRule string, notifyChannelIDs []uint, failfast *bool) (string, error) {
	// 1. Node Selection
	targetNode := nodeName
	if targetNode == "" {
		var err error
		targetNode, err = p.selectRandomNode()
		if err != nil {
			return "", fmt.Errorf("failed to select node: %w", err)
		}
	}

	// 2. Construct Message
	msg := TaskMessage{
		TaskID:           taskID,
		ReportID:         reportID,
		CaseType:         caseType,
		CaseID:           caseID,
		RunMode:          runMode,
		NodeName:         targetNode,
		EnvID:            envID,
		ConfigID:         configID,
		ProjectID:        projectID,
		NotifyEnabled:    notifyEnabled,
		NotifyRule:       notifyRule,
		NotifyChannelIDs: notifyChannelIDs,
		Failfast:         failfast,
		CreateTime:       time.Now().Unix(),
		ExpireTime:       time.Now().Add(24 * time.Hour).Unix(), // Default expiration 24h
	}

	body, err := json.Marshal(msg)
	if err != nil {
		return "", fmt.Errorf("failed to marshal message: %w", err)
	}

	// 3. Determine Queue/Routing Key
	// Queue Name = Prefix + NodeName
	queueName := targetNode
	if p.cfg.QueuePrefix != "" {
		queueName = p.cfg.QueuePrefix + targetNode
	}

	// 4. Publish
	// Using the configured Exchange.
	var lastErr error
	retryCount := p.cfg.RetryCount
	if retryCount <= 0 {
		retryCount = 1
	}

	for i := 0; i < retryCount; i++ {
		lastErr = p.client.Publish(p.cfg.Exchange, queueName, body)
		if lastErr == nil {
			p.logger.Info("Task sent to MQ",
				zap.String("task_id", taskID),
				zap.String("node", targetNode),
				zap.String("queue", queueName),
			)
			return targetNode, nil
		}

		p.logger.Warn("Failed to send task to MQ, retrying...",
			zap.String("task_id", taskID),
			zap.Int("attempt", i+1),
			zap.Error(lastErr),
		)
		time.Sleep(1 * time.Second)
	}

	p.logger.Error("Failed to send task to MQ after retries",
		zap.String("task_id", taskID),
		zap.String("node", targetNode),
		zap.Error(lastErr),
	)
	return "", lastErr
}

// local struct to avoid importing model package
type runnerNode struct {
	NodeName string `gorm:"column:node_name"`
}

// selectRandomNode 从数据库中随机选择一个在线节点
func (p *RunnerTaskProducer) selectRandomNode() (string, error) {
	var nodes []runnerNode
	// Status 1 = Online
	// Table name: yf_runner_node
	err := p.db.Table("yf_runner_node").Where("status = ?", 1).Find(&nodes).Error
	if err != nil {
		return "", err
	}

	if len(nodes) == 0 {
		return "", fmt.Errorf("no online nodes available")
	}

	idx := rand.Intn(len(nodes))
	return nodes[idx].NodeName, nil
}
