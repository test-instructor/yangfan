package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type RunnerService struct {
	nodeName   string
	nodeAlias  string
	runContent string
	nodeID     string
	ip         string
	port       int
	stopChan   chan struct{}
}

func NewRunnerService(nodeName, nodeAlias, runContent string, port int) *RunnerService {
	return &RunnerService{
		nodeName:   nodeName,
		nodeAlias:  nodeAlias,
		runContent: runContent,
		nodeID:     uuid.New().String(),
		ip:         "127.0.0.1", // TODO: Get actual IP
		port:       port,
		stopChan:   make(chan struct{}),
	}
}

func (s *RunnerService) Register() error {
	var node platform.RunnerNode
	err := global.GVA_DB.Where("node_name = ?", s.nodeName).First(&node).Error

	now := time.Now()
	status := int64(1) // Online
	port := int64(s.port)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Create new
		node = platform.RunnerNode{
			NodeName:      s.nodeName,
			NodeId:        &s.nodeID,
			Ip:            &s.ip,
			Port:          &port,
			Status:        &status,
			LastHeartbeat: &now,
			CreateTime:    &now,
			Alias:         &s.nodeAlias,
			RunContent:    &s.runContent,
		}
		if err := global.GVA_DB.Create(&node).Error; err != nil {
			return err
		}
	} else if err == nil {
		// Update existing
		node.NodeId = &s.nodeID
		node.Alias = &s.nodeAlias
		node.RunContent = &s.runContent
		node.Ip = &s.ip
		node.Port = &port
		node.Status = &status
		node.LastHeartbeat = &now
		if err := global.GVA_DB.Save(&node).Error; err != nil {
			return err
		}
	} else {
		return err
	}

	global.GVA_LOG.Info("Runner registered", zap.String("node", s.nodeName), zap.String("id", s.nodeID))
	return nil
}

func (s *RunnerService) StartHeartbeat() {
	ticker := time.NewTicker(30 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				s.beat()
			case <-s.stopChan:
				ticker.Stop()
				return
			}
		}
	}()
}

func (s *RunnerService) beat() {
	now := time.Now()
	err := global.GVA_DB.Model(&platform.RunnerNode{}).
		Where("node_name = ?", s.nodeName).
		Update("last_heartbeat", now).Error
	if err != nil {
		global.GVA_LOG.Error("Heartbeat failed", zap.Error(err))
	}
}

func (s *RunnerService) Stop() {
	close(s.stopChan)
	// Update status to Offline
	status := int64(0)
	global.GVA_DB.Model(&platform.RunnerNode{}).
		Where("node_name = ?", s.nodeName).
		Update("status", status)
	global.GVA_LOG.Info("Runner stopped", zap.String("node", s.nodeName))
}
