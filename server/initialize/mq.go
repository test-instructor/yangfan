package initialize

import (
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/utils/mq"
	"go.uber.org/zap"
)

func InitMQ() {
	config := global.GVA_CONFIG.MQ
	if config.Type == "" {
		return
	}

	loader := mq.NewMQConfigLoader()
	if err := loader.LoadAndValidate(config); err != nil {
		global.GVA_LOG.Error("MQ Config validation failed", zap.Error(err))
		return
	}

	client, err := mq.NewMQClient(config, global.GVA_LOG)
	if err != nil {
		global.GVA_LOG.Error("MQ Initialization failed", zap.Error(err))
		return
	}

	global.GVA_MQ = client
	global.GVA_MQ_PRODUCER = mq.NewRunnerTaskProducer(client, global.GVA_DB, config, global.GVA_LOG)
	global.GVA_LOG.Info("MQ Initialized successfully", zap.String("type", config.Type))
}
