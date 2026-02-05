package ui_node

import (
	"fmt"

	"github.com/test-instructor/yangfan/server/v2/global"
	apiInit "github.com/test-instructor/yangfan/server/v2/initialize/ui_node/api"
	menuInit "github.com/test-instructor/yangfan/server/v2/initialize/ui_node/menu"
	sysModel "github.com/test-instructor/yangfan/server/v2/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func Init() {
	if global.GVA_DB == nil {
		return
	}

	apisOk := true
	if err := apiInit.Init(); err != nil {
		apisOk = false
		global.GVA_LOG.Error("init ui node sys_apis failed", zap.Error(err))
	}

	menusOk := true
	if err := menuInit.Init(); err != nil {
		menusOk = false
		global.GVA_LOG.Error("init ui node menus failed", zap.Error(err))
	}

	if apisOk && menusOk {
		assignUINodeMenuReadApiToAllAuthorities(global.GVA_DB)
		assignUINodeMenuAdminApisToRole888(global.GVA_DB)
	}
}

func assignUINodeMenuReadApiToAllAuthorities(db *gorm.DB) {
	var authorityIDs []uint
	if err := db.Model(&sysModel.SysAuthority{}).Pluck("authority_id", &authorityIDs).Error; err != nil {
		global.GVA_LOG.Error("assignUINodeMenuReadApiToAllAuthorities: failed to load authority ids", zap.Error(err))
		return
	}
	if len(authorityIDs) == 0 {
		return
	}

	readApis := []struct {
		Method string
		Path   string
	}{
		{Method: "POST", Path: "/ui/node/menu/getMenuTree"},
	}

	casbinRules := make([]gormadapter.CasbinRule, 0, len(authorityIDs)*len(readApis))
	for _, aid := range authorityIDs {
		sub := fmt.Sprintf("%d", aid)
		for _, api := range readApis {
			casbinRules = append(casbinRules, gormadapter.CasbinRule{
				Ptype: "p",
				V0:    sub,
				V1:    api.Path,
				V2:    api.Method,
			})
		}
	}
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&casbinRules).Error; err != nil {
		global.GVA_LOG.Error("assignUINodeMenuReadApiToAllAuthorities: failed to create casbin rules", zap.Error(err))
	}
}

func assignUINodeMenuAdminApisToRole888(db *gorm.DB) {
	var authority sysModel.SysAuthority
	if err := db.Where("authority_id = ?", 888).First(&authority).Error; err != nil {
		global.GVA_LOG.Error("assignUINodeMenuAdminApisToRole888: failed to find authority 888", zap.Error(err))
		return
	}

	apis := apiInit.GetApis()
	if len(apis) == 0 {
		return
	}

	casbinRules := make([]gormadapter.CasbinRule, 0, len(apis))
	authorityIdStr := "888"
	for _, api := range apis {
		if api.Path == "/ui/node/menu/getMenuTree" {
			continue
		}
		if api.Path == "" || api.Method == "" {
			continue
		}
		casbinRules = append(casbinRules, gormadapter.CasbinRule{
			Ptype: "p",
			V0:    authorityIdStr,
			V1:    api.Path,
			V2:    api.Method,
		})
	}
	if len(casbinRules) == 0 {
		return
	}
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&casbinRules).Error; err != nil {
		global.GVA_LOG.Error("assignUINodeMenuAdminApisToRole888: failed to create casbin rules", zap.Error(err))
	}
}
