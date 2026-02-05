package menu

import (
	"errors"
	"fmt"

	"github.com/test-instructor/yangfan/server/v2/global"
	sysModel "github.com/test-instructor/yangfan/server/v2/model/system"
	"gorm.io/gorm"
)

func Init() error {
	if global.GVA_DB == nil {
		return nil
	}
	if err := ensureMenus(global.GVA_DB, Seeds()); err != nil {
		return err
	}
	if err := ensureActiveNameDefaults(global.GVA_DB); err != nil {
		return err
	}
	return ensureMenuParameters(global.GVA_DB, ParameterSeeds())
}

func ensureActiveNameDefaults(db *gorm.DB) error {
	if db == nil {
		return nil
	}
	return db.Model(&sysModel.SysUINodeMenu{}).
		Where("name = ? AND (active_name = '' OR active_name IS NULL)", "autoReportAndroidDetail").
		Update("active_name", "autoReportAndroid").Error
}

func ensureMenus(db *gorm.DB, seeds []MenuSeed) error {
	if db == nil || len(seeds) == 0 {
		return nil
	}

	seedByName := make(map[string]MenuSeed, len(seeds))
	names := make([]string, 0, len(seeds))
	for _, s := range seeds {
		if s.Name == "" {
			continue
		}
		seedByName[s.Name] = s
		names = append(names, s.Name)
	}

	existingByName := make(map[string]*sysModel.SysUINodeMenu, len(seeds))
	var existing []sysModel.SysUINodeMenu
	if err := db.Model(&sysModel.SysUINodeMenu{}).
		Select("id", "name", "menu_level", "parent_id").
		Where("name IN ?", names).
		Find(&existing).Error; err != nil {
		return err
	}
	for i := range existing {
		m := &existing[i]
		if m.Name == "" {
			continue
		}
		existingByName[m.Name] = m
	}

	levelMemo := map[string]uint{}
	var menuLevel func(name string) uint
	menuLevel = func(name string) uint {
		if v, ok := levelMemo[name]; ok {
			return v
		}
		s, ok := seedByName[name]
		if !ok || s.ParentName == "" {
			levelMemo[name] = 0
			return 0
		}
		v := menuLevel(s.ParentName) + 1
		levelMemo[name] = v
		return v
	}

	creating := map[string]bool{}
	var ensureOne func(name string) error
	ensureOne = func(name string) error {
		if name == "" {
			return nil
		}
		if _, ok := existingByName[name]; ok {
			return nil
		}
		if creating[name] {
			return fmt.Errorf("cycle detected when creating ui node menu name=%s", name)
		}
		s, ok := seedByName[name]
		if !ok {
			return nil
		}
		creating[name] = true
		defer func() { delete(creating, name) }()

		parentID := uint(0)
		if s.ParentName != "" {
			if err := ensureOne(s.ParentName); err != nil {
				return err
			}
			parent, ok := existingByName[s.ParentName]
			if !ok {
				return fmt.Errorf("parent menu not found: child=%s parent=%s", s.Name, s.ParentName)
			}
			parentID = parent.ID
		}

		m := sysModel.SysUINodeMenu{
			MenuLevel: menuLevel(s.Name),
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
		existingByName[m.Name] = &m
		return nil
	}

	for _, s := range seeds {
		if err := ensureOne(s.Name); err != nil {
			return err
		}
	}
	return nil
}

func ensureMenuParameters(db *gorm.DB, seeds []MenuParameterSeed) error {
	if db == nil || len(seeds) == 0 {
		return nil
	}

	for _, s := range seeds {
		if s.MenuName == "" || s.Type == "" || s.Key == "" {
			continue
		}

		query := db.Model(&sysModel.SysUINodeMenu{}).Select("id").Where("name = ?", s.MenuName)
		if s.Component != "" {
			query = query.Where("component = ?", s.Component)
		}

		var menus []sysModel.SysUINodeMenu
		if err := query.Find(&menus).Error; err != nil {
			return err
		}
		if len(menus) == 0 {
			return fmt.Errorf("ui node menu not found when ensuring menu parameter: name=%s component=%s", s.MenuName, s.Component)
		}

		for _, m := range menus {
			if m.ID == 0 {
				continue
			}

			var existing sysModel.SysUINodeMenuParameter
			err := db.Where(map[string]any{
				"sys_ui_node_menu_id": m.ID,
				"type":                s.Type,
				"key":                 s.Key,
			}).First(&existing).Error
			if err == nil {
				if existing.Value != s.Value {
					if err := db.Model(&existing).Update("value", s.Value).Error; err != nil {
						return err
					}
				}
				continue
			}
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}

			toCreate := sysModel.SysUINodeMenuParameter{
				SysUINodeMenuID: m.ID,
				Type:            s.Type,
				Key:             s.Key,
				Value:           s.Value,
			}
			if err := db.Create(&toCreate).Error; err != nil {
				return err
			}
		}
	}
	return nil
}
