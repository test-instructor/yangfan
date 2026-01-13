package global

import (
	"fmt"
	"sync"

	"github.com/mark3labs/mcp-go/server"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/qmgo"

	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/test-instructor/yangfan/server/v2/utils/mq"
	"github.com/test-instructor/yangfan/server/v2/utils/timer"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/test-instructor/yangfan/server/v2/config"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_DB        *gorm.DB
	GVA_DBList    map[string]*gorm.DB
	GVA_REDIS     redis.UniversalClient
	GVA_REDISList map[string]redis.UniversalClient
	GVA_MONGO     *qmgo.QmgoClient
	GVA_CONFIG    config.Server
	GVA_VP        *viper.Viper
	// GVA_LOG    *oplogging.Logger
	GVA_LOG                 *zap.Logger
	GVA_Timer               timer.Timer = timer.NewTimerTask()
	GVA_Concurrency_Control             = &singleflight.Group{}
	GVA_ROUTERS             gin.RoutesInfo
	GVA_ACTIVE_DBNAME       *string
	GVA_MCP_SERVER          *server.MCPServer
	GVA_MQ                  *mq.MQClient
	GVA_MQ_PRODUCER         *mq.RunnerTaskProducer
	BlackCache              local_cache.Cache
	lock                    sync.RWMutex

	PythonVENV        string
	PythonVenvDir     string
	Python3Executable string

	// 数据仓库配置
	GVA_DW_HOST string
	GVA_DW_PORT int
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return GVA_DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := GVA_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}

func GetRedis(name string) redis.UniversalClient {
	redis, ok := GVA_REDISList[name]
	if !ok || redis == nil {
		panic(fmt.Sprintf("redis `%s` no init", name))
	}
	return redis
}
