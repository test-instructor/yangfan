package platform

import (
	"errors"

	"github.com/test-instructor/yangfan/server/v2/global"
	apiInit "github.com/test-instructor/yangfan/server/v2/initialize/platform/api"
	menuInit "github.com/test-instructor/yangfan/server/v2/initialize/platform/menu"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Init initializes platform-related seed data (sys_apis + sys_base_menus) once per version.
// If the current version is already marked as initialized, it will be skipped.
func Init() {
	if global.GVA_DB == nil {
		return
	}

	// Ensure state table exists.
	if err := global.GVA_DB.AutoMigrate(&PlatformSeedState{}); err != nil {
		global.GVA_LOG.Error("auto migrate platform seed state failed", zap.Error(err))
		return
	}

	version := global.Version
	state, err := getOrInitState(global.GVA_DB, version)
	if err != nil {
		global.GVA_LOG.Error("load platform seed state failed", zap.Error(err))
		return
	}

	if state.ApisInited && state.MenusInited {
		return
	}

	apisOk := state.ApisInited
	menusOk := state.MenusInited

	if !apisOk {
		if err := apiInit.Init(); err != nil {
			global.GVA_LOG.Error("init sys_apis failed", zap.Error(err))
		} else {
			apisOk = true
		}
	}

	if !menusOk {
		if err := menuInit.Init(); err != nil {
			global.GVA_LOG.Error("init sys_base_menus failed", zap.Error(err))
		} else {
			menusOk = true
		}
	}

	if err := upsertState(global.GVA_DB, version, apisOk, menusOk); err != nil {
		global.GVA_LOG.Error("save platform seed state failed", zap.Error(err))
	}
}

func getOrInitState(db *gorm.DB, version string) (PlatformSeedState, error) {
	var state PlatformSeedState
	err := db.Where("version = ?", version).First(&state).Error
	if err == nil {
		return state, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Return zero state; caller will upsert after running validations.
		return PlatformSeedState{Version: version}, nil
	}
	return PlatformSeedState{}, err
}

func upsertState(db *gorm.DB, version string, apisInited, menusInited bool) error {
	// Try update first.
	res := db.Model(&PlatformSeedState{}).Where("version = ?", version).Updates(map[string]any{
		"apis_inited":  apisInited,
		"menus_inited": menusInited,
	})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected > 0 {
		return nil
	}
	// No existing row, create.
	state := PlatformSeedState{Version: version, ApisInited: apisInited, MenusInited: menusInited}
	return db.Create(&state).Error
}
