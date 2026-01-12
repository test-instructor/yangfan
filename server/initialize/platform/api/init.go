package api

import (
	"github.com/test-instructor/yangfan/server/v2/global"
	sysModel "github.com/test-instructor/yangfan/server/v2/model/system"
	"gorm.io/gorm"
)

func GetApis() []sysModel.SysApi {
	apis := make([]sysModel.SysApi, 0, 128)
	apis = append(apis, projectmgrApis()...)
	apis = append(apis, platformApis()...)
	apis = append(apis, automationApis()...)
	apis = append(apis, datawarehouseApis()...)
	apis = append(apis, uiApis()...)
	return apis
}

// Init ensures curated sys_apis routes exist.
// It only runs when global.GVA_DB is initialized.
func Init() error {
	if global.GVA_DB == nil {
		return nil
	}

	apis := GetApis()

	return ensureSysApis(global.GVA_DB, apis)
}

func ensureSysApis(db *gorm.DB, apis []sysModel.SysApi) error {
	if db == nil || len(apis) == 0 {
		return nil
	}

	// Dedup desired APIs and collect unique paths for a single lookup query.
	desired := make(map[string]sysModel.SysApi, len(apis))
	paths := make([]string, 0, len(apis))
	pathSeen := make(map[string]struct{}, len(apis))
	for _, api := range apis {
		if api.Path == "" || api.Method == "" {
			continue
		}
		key := api.Method + " " + api.Path
		if _, ok := desired[key]; ok {
			continue
		}
		desired[key] = api
		if _, ok := pathSeen[api.Path]; !ok {
			pathSeen[api.Path] = struct{}{}
			paths = append(paths, api.Path)
		}
	}
	if len(desired) == 0 {
		return nil
	}

	var existing []sysModel.SysApi
	if err := db.Model(&sysModel.SysApi{}).
		Select("path", "method").
		Where("path IN ?", paths).
		Find(&existing).Error; err != nil {
		return err
	}

	existed := make(map[string]struct{}, len(existing))
	for _, e := range existing {
		if e.Path == "" || e.Method == "" {
			continue
		}
		existed[e.Method+" "+e.Path] = struct{}{}
	}

	toCreate := make([]sysModel.SysApi, 0, len(desired))
	for key, api := range desired {
		if _, ok := existed[key]; ok {
			continue
		}
		toCreate = append(toCreate, api)
	}
	if len(toCreate) == 0 {
		return nil
	}

	return db.Create(&toCreate).Error
}
