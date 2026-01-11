package initialize

import (
	"github.com/test-instructor/yangfan/server/v2/global"
	"go.uber.org/zap"
)

// Reload 优雅地重新加载系统配置
func Reload() error {
	global.GVA_LOG.Info("正在重新加载系统配置...")

	// 重新加载配置文件
	if err := global.GVA_VP.ReadInConfig(); err != nil {
		global.GVA_LOG.Error("重新读取配置文件失败!", zap.Error(err))
		return err
	}
	if err := global.GVA_VP.Unmarshal(&global.GVA_CONFIG); err != nil {
		global.GVA_LOG.Error("重新解析配置文件失败!", zap.Error(err))
		return err
	}

	// 重新初始化数据库连接
	if global.GVA_DB != nil {
		db, _ := global.GVA_DB.DB()
		err := db.Close()
		if err != nil {
			global.GVA_LOG.Error("关闭原数据库连接失败!", zap.Error(err))
			return err
		}
	}

	// 重新建立数据库连接
	global.GVA_DB = Gorm()

	// 重新初始化其他配置
	OtherInit()
	DBList()

	if global.GVA_DB != nil {
		// 确保数据库表结构是最新的
		RegisterTables()
	}

	// 重新初始化定时任务
	Timer()

	// 重新初始化MQ连接
	if global.GVA_MQ != nil {
		global.GVA_MQ.Close()
		global.GVA_MQ = nil
		global.GVA_MQ_PRODUCER = nil
	}
	InitMQ()

	global.GVA_LOG.Info("系统配置重新加载完成")
	return nil
}
