package global

import (
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/test-instructor/yangfan/server/utils/timer"
	"sync"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/test-instructor/yangfan/server/config"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_DBList map[string]*gorm.DB
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	// GVA_LOG    *oplogging.Logger
	GVA_LOG                 *zap.Logger
	GVA_Timer               timer.Timer = timer.NewApiCase()
	GVA_Concurrency_Control             = &singleflight.Group{}

	BlackCache        local_cache.Cache
	lock              sync.RWMutex
	DebugTalkLock     = make(map[string]*sync.Mutex)
	DebugTalkFileLock = sync.RWMutex{}

	HrpMode       HrpModes
	IgnoreInstall bool
)

type HrpModes int

const (
	HrpModeMaster HrpModes = 1
	HrpModeWork   HrpModes = 2
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
