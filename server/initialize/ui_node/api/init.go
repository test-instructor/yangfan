package api

import (
	"github.com/test-instructor/yangfan/server/v2/global"
	sysModel "github.com/test-instructor/yangfan/server/v2/model/system"
	"gorm.io/gorm"
)

func GetApis() []sysModel.SysApi {
	return []sysModel.SysApi{
		{Path: "/ui/node/menu/getMenuTree", Description: "获取 UI Node 菜单树", ApiGroup: "UI Node 菜单", Method: "POST"},
		{Path: "/ui/node/menu/create", Description: "创建 UI Node 菜单", ApiGroup: "UI Node 菜单", Method: "POST"},
		{Path: "/ui/node/menu/update", Description: "更新 UI Node 菜单", ApiGroup: "UI Node 菜单", Method: "PUT"},
		{Path: "/ui/node/menu/delete", Description: "删除 UI Node 菜单", ApiGroup: "UI Node 菜单", Method: "DELETE"},
		{Path: "/ui/node/menu/list", Description: "获取 UI Node 菜单列表", ApiGroup: "UI Node 菜单", Method: "GET"},
	}
}

func Init() error {
	if global.GVA_DB == nil {
		return nil
	}
	return ensureSysApis(global.GVA_DB, GetApis())
}

func ensureSysApis(db *gorm.DB, apis []sysModel.SysApi) error {
	if db == nil || len(apis) == 0 {
		return nil
	}

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
