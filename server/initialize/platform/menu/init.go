package menu

import (
	"fmt"

	"github.com/test-instructor/yangfan/server/v2/global"
	sysModel "github.com/test-instructor/yangfan/server/v2/model/system"
	"gorm.io/gorm"
)

// Init ensures curated sys_base_menus exist.
// It only runs when global.GVA_DB is initialized.
func Init() error {
	if global.GVA_DB == nil {
		return nil
	}
	return ensureSysBaseMenus(global.GVA_DB, Seeds())
}

func ensureSysBaseMenus(db *gorm.DB, seeds []MenuSeed) error {
	if db == nil || len(seeds) == 0 {
		return nil
	}

	seedByPath := make(map[string]MenuSeed, len(seeds))
	paths := make([]string, 0, len(seeds))
	for _, s := range seeds {
		if s.Path == "" {
			continue
		}
		seedByPath[s.Path] = s
		paths = append(paths, s.Path)
	}

	// Load existing menus by path (gorm default scope: excludes soft-deleted rows).
	existingByPath := make(map[string]*sysModel.SysBaseMenu, len(seeds))
	var existing []sysModel.SysBaseMenu
	if err := db.Model(&sysModel.SysBaseMenu{}).
		Select("id", "path", "menu_level", "parent_id").
		Where("path IN ?", paths).
		Find(&existing).Error; err != nil {
		return err
	}
	for i := range existing {
		m := &existing[i]
		if m.Path == "" {
			continue
		}
		existingByPath[m.Path] = m
	}

	// Compute menu level based on seed tree (not DB values).
	levelMemo := map[string]uint{}
	var menuLevel func(path string) uint
	menuLevel = func(path string) uint {
		if v, ok := levelMemo[path]; ok {
			return v
		}
		s, ok := seedByPath[path]
		if !ok || s.ParentPath == "" {
			levelMemo[path] = 0
			return 0
		}
		v := menuLevel(s.ParentPath) + 1
		levelMemo[path] = v
		return v
	}

	creating := map[string]bool{}
	var ensureOne func(path string) error
	ensureOne = func(path string) error {
		if path == "" {
			return nil
		}
		if _, ok := existingByPath[path]; ok {
			return nil
		}
		if creating[path] {
			return fmt.Errorf("cycle detected when creating menu path=%s", path)
		}
		s, ok := seedByPath[path]
		if !ok {
			return nil
		}
		creating[path] = true
		defer func() { delete(creating, path) }()

		parentID := uint(0)
		if s.ParentPath != "" {
			if err := ensureOne(s.ParentPath); err != nil {
				return err
			}
			parent, ok := existingByPath[s.ParentPath]
			if !ok {
				return fmt.Errorf("parent menu not found: child=%s parent=%s", s.Path, s.ParentPath)
			}
			parentID = parent.ID
		}

		m := sysModel.SysBaseMenu{
			MenuLevel: menuLevel(s.Path),
			ParentId:  parentID,
			Path:      s.Path,
			Name:      s.Name,
			Hidden:    s.Hidden,
			Component: s.Component,
			Sort:      s.Sort,
			Meta:      s.Meta,
		}
		if err := db.Create(&m).Error; err != nil {
			return err
		}
		existingByPath[m.Path] = &m
		return nil
	}

	for _, s := range seeds {
		if err := ensureOne(s.Path); err != nil {
			return err
		}
	}
	return nil
}
