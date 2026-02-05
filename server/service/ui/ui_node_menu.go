package ui

import (
	"errors"

	"github.com/test-instructor/yangfan/server/v2/global"
	sysModel "github.com/test-instructor/yangfan/server/v2/model/system"
	"gorm.io/gorm"
)

type UINodeMenuService struct{}

func (s *UINodeMenuService) GetMenuTree() ([]sysModel.SysUINodeMenu, error) {
	if global.GVA_DB == nil {
		return []sysModel.SysUINodeMenu{}, nil
	}

	var all []sysModel.SysUINodeMenu
	if err := global.GVA_DB.Order("parent_id asc").Order("sort asc").Preload("Parameters").Find(&all).Error; err != nil {
		return nil, err
	}

	treeMap := make(map[uint][]sysModel.SysUINodeMenu)
	for _, m := range all {
		treeMap[m.ParentId] = append(treeMap[m.ParentId], m)
	}

	roots := treeMap[0]
	for i := 0; i < len(roots); i++ {
		fillChildren(&roots[i], treeMap)
	}
	return roots, nil
}

func fillChildren(menu *sysModel.SysUINodeMenu, treeMap map[uint][]sysModel.SysUINodeMenu) {
	if menu == nil {
		return
	}
	menu.Children = treeMap[menu.ID]
	for i := 0; i < len(menu.Children); i++ {
		fillChildren(&menu.Children[i], treeMap)
	}
}

func (s *UINodeMenuService) CreateMenu(menu sysModel.SysUINodeMenu) error {
	if global.GVA_DB == nil {
		return errors.New("db 未初始化")
	}
	if menu.Name == "" || menu.Path == "" {
		return errors.New("name/path 不能为空")
	}

	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var existing sysModel.SysUINodeMenu
		if !errors.Is(tx.Where("name = ?", menu.Name).First(&existing).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在重复name，请修改name")
		}
		if menu.ParentId != 0 {
			var parent sysModel.SysUINodeMenu
			if err := tx.First(&parent, menu.ParentId).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return errors.New("父菜单不存在")
				}
				return err
			}
		}

		params := menu.Parameters
		menu.Parameters = nil
		if err := tx.Create(&menu).Error; err != nil {
			return err
		}
		for _, p := range params {
			if p.Type == "" || p.Key == "" {
				continue
			}
			p.SysUINodeMenuID = menu.ID
			if err := tx.Create(&p).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *UINodeMenuService) UpdateMenu(menu sysModel.SysUINodeMenu) error {
	if global.GVA_DB == nil {
		return errors.New("db 未初始化")
	}
	if menu.ID == 0 {
		return errors.New("id 不能为空")
	}
	if menu.Name == "" || menu.Path == "" {
		return errors.New("name/path 不能为空")
	}

	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var existing sysModel.SysUINodeMenu
		if err := tx.First(&existing, menu.ID).Error; err != nil {
			return err
		}
		if menu.ParentId == menu.ID {
			return errors.New("父菜单不能是自己")
		}
		if menu.ParentId != 0 {
			var parent sysModel.SysUINodeMenu
			if err := tx.First(&parent, menu.ParentId).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return errors.New("父菜单不存在")
				}
				return err
			}
		}

		params := menu.Parameters
		menu.Parameters = nil
		if err := tx.Model(&existing).Updates(map[string]any{
			"parent_id":       menu.ParentId,
			"path":            menu.Path,
			"name":            menu.Name,
			"hidden":          menu.Hidden,
			"component":       menu.Component,
			"sort":            menu.Sort,
			"active_name":     menu.Meta.ActiveName,
			"keep_alive":      menu.Meta.KeepAlive,
			"default_menu":    menu.Meta.DefaultMenu,
			"title":           menu.Meta.Title,
			"icon":            menu.Meta.Icon,
			"close_tab":       menu.Meta.CloseTab,
			"transition_type": menu.Meta.TransitionType,
		}).Error; err != nil {
			return err
		}

		if err := tx.Where("sys_ui_node_menu_id = ?", menu.ID).Delete(&sysModel.SysUINodeMenuParameter{}).Error; err != nil {
			return err
		}
		for _, p := range params {
			if p.Type == "" || p.Key == "" {
				continue
			}
			p.SysUINodeMenuID = menu.ID
			if err := tx.Create(&p).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *UINodeMenuService) DeleteMenu(id uint) error {
	if global.GVA_DB == nil {
		return errors.New("db 未初始化")
	}
	if id == 0 {
		return errors.New("id 不能为空")
	}

	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var childCount int64
		if err := tx.Model(&sysModel.SysUINodeMenu{}).Where("parent_id = ?", id).Count(&childCount).Error; err != nil {
			return err
		}
		if childCount > 0 {
			return errors.New("存在子菜单，无法删除")
		}
		if err := tx.Where("sys_ui_node_menu_id = ?", id).Delete(&sysModel.SysUINodeMenuParameter{}).Error; err != nil {
			return err
		}
		return tx.Delete(&sysModel.SysUINodeMenu{}, id).Error
	})
}

func (s *UINodeMenuService) ListMenus() ([]sysModel.SysUINodeMenu, error) {
	if global.GVA_DB == nil {
		return []sysModel.SysUINodeMenu{}, nil
	}
	var all []sysModel.SysUINodeMenu
	if err := global.GVA_DB.Order("parent_id asc").Order("sort asc").Preload("Parameters").Find(&all).Error; err != nil {
		return nil, err
	}
	return all, nil
}
