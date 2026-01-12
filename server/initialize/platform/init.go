package platform

import (
	"errors"
	"strconv"

	"github.com/test-instructor/yangfan/server/v2/global"
	apiInit "github.com/test-instructor/yangfan/server/v2/initialize/platform/api"
	menuInit "github.com/test-instructor/yangfan/server/v2/initialize/platform/menu"
	sysModel "github.com/test-instructor/yangfan/server/v2/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	gormadapter "github.com/casbin/gorm-adapter/v3"
)

// Init initializes and keeps platform-related seed data (sys_apis + sys_base_menus) in sync.
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

	if _, err := getOrInitState(global.GVA_DB, version); err != nil {
		global.GVA_LOG.Error("load platform seed state failed", zap.Error(err))
	}

	apisOk := true
	if err := apiInit.Init(); err != nil {
		apisOk = false
		global.GVA_LOG.Error("init sys_apis failed", zap.Error(err))
	}

	menusOk := true
	if err := menuInit.Init(); err != nil {
		menusOk = false
		global.GVA_LOG.Error("init sys_base_menus failed", zap.Error(err))
	}

	if apisOk && menusOk {
		assignPermissionsToRole888(global.GVA_DB)
	}

	if err := upsertState(global.GVA_DB, version, apisOk, menusOk); err != nil {
		global.GVA_LOG.Error("save platform seed state failed", zap.Error(err))
	}
}

func assignPermissionsToRole888(db *gorm.DB) {
	// 1. Assign Menus
	var authority sysModel.SysAuthority
	if err := db.Where("authority_id = ?", 888).First(&authority).Error; err != nil {
		global.GVA_LOG.Error("assignPermissionsToRole888: failed to find authority 888", zap.Error(err))
		return
	}

	seeds := menuInit.Seeds()
	var menuPaths []string
	for _, s := range seeds {
		menuPaths = append(menuPaths, s.Path)
	}

	var menus []sysModel.SysBaseMenu
	if len(menuPaths) > 0 {
		if err := db.Where("path IN ?", menuPaths).Find(&menus).Error; err != nil {
			global.GVA_LOG.Error("assignPermissionsToRole888: failed to find menus", zap.Error(err))
			return
		}

		menuIdStrs := make([]string, 0, len(menus))
		for _, m := range menus {
			if m.ID == 0 {
				continue
			}
			menuIdStrs = append(menuIdStrs, strconv.Itoa(int(m.ID)))
		}

		if len(menuIdStrs) > 0 {
			var existingMenuIds []string
			if err := db.Model(&sysModel.SysAuthorityMenu{}).
				Where("sys_authority_authority_id = ?", "888").
				Where("sys_base_menu_id IN ?", menuIdStrs).
				Pluck("sys_base_menu_id", &existingMenuIds).Error; err != nil {
				global.GVA_LOG.Error("assignPermissionsToRole888: failed to load existing authority menus", zap.Error(err))
				return
			}

			existing := make(map[string]struct{}, len(existingMenuIds))
			for _, id := range existingMenuIds {
				if id == "" {
					continue
				}
				existing[id] = struct{}{}
			}

			toCreate := make([]sysModel.SysAuthorityMenu, 0, len(menuIdStrs))
			for _, id := range menuIdStrs {
				if _, ok := existing[id]; ok {
					continue
				}
				toCreate = append(toCreate, sysModel.SysAuthorityMenu{
					MenuId:      id,
					AuthorityId: "888",
				})
			}

			if len(toCreate) > 0 {
				if err := db.Create(&toCreate).Error; err != nil {
					global.GVA_LOG.Error("assignPermissionsToRole888: failed to create authority menus", zap.Error(err))
				}
			}
		}
	}

	// 2. Assign APIs (Casbin)
	apis := apiInit.GetApis()
	if len(apis) > 0 {
		var casbinRules []gormadapter.CasbinRule
		authorityIdStr := "888"
		for _, api := range apis {
			casbinRules = append(casbinRules, gormadapter.CasbinRule{
				Ptype: "p",
				V0:    authorityIdStr,
				V1:    api.Path,
				V2:    api.Method,
			})
		}

		if len(casbinRules) > 0 {
			if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&casbinRules).Error; err != nil {
				global.GVA_LOG.Error("assignPermissionsToRole888: failed to create casbin rules", zap.Error(err))
			}
		}
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
