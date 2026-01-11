package platform

import "github.com/test-instructor/yangfan/server/v2/global"

// PlatformSeedState tracks whether platform seeds have been initialized for a given version.
// This prevents repeated inserts on every startup.
type PlatformSeedState struct {
	global.GVA_MODEL
	Version     string `json:"version" gorm:"size:32;uniqueIndex;comment:platform init version"`
	ApisInited  bool   `json:"apisInited" gorm:"comment:sys_apis initialized"`
	MenusInited bool   `json:"menusInited" gorm:"comment:sys_base_menus initialized"`
}

func (PlatformSeedState) TableName() string {
	return "sys_platform_seed_states"
}
