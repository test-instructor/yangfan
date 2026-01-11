package platform

import (
	"context"
	"errors"

	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	platformReq "github.com/test-instructor/yangfan/server/v2/model/platform/request"
)

type CategoryMenuService struct{}

// CreateCategoryMenu 创建自动化菜单记录
// Author [yourname](https://github.com/yourname)
func (cmService *CategoryMenuService) CreateCategoryMenu(ctx context.Context, cm *platform.CategoryMenu) (err error) {
	err = global.GVA_DB.Create(cm).Error
	return err
}

// DeleteCategoryMenu 删除自动化菜单记录
// Author [yourname](https://github.com/yourname)
func (cmService *CategoryMenuService) DeleteCategoryMenu(ctx context.Context, ID string, projectId int64) (err error) {
	var cm platform.CategoryMenu
	err = global.GVA_DB.Where("id = ?", ID).First(&cm).Error
	if err != nil {
		return err
	}
	if cm.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}
	err = global.GVA_DB.Delete(&platform.CategoryMenu{}, "id = ?", ID).Error
	return err
}

// DeleteCategoryMenuByIds 批量删除自动化菜单记录
// Author [yourname](https://github.com/yourname)
func (cmService *CategoryMenuService) DeleteCategoryMenuByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]platform.CategoryMenu{}, "id in ?", IDs).Error
	return err
}

// UpdateCategoryMenu 更新自动化菜单记录
// Author [yourname](https://github.com/yourname)
func (cmService *CategoryMenuService) UpdateCategoryMenu(ctx context.Context, cm platform.CategoryMenu, projectId int64) (err error) {
	var oldCategoryMenu platform.CategoryMenu
	err = global.GVA_DB.Model(&oldCategoryMenu).Where("id = ?", cm.ID).First(&oldCategoryMenu).Error
	if err != nil {
		return err
	}
	if oldCategoryMenu.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}

	err = global.GVA_DB.Model(&platform.CategoryMenu{}).Where("id = ?", cm.ID).Updates(&cm).Error
	return err
}

// GetCategoryMenu 根据ID获取自动化菜单记录
// Author [yourname](https://github.com/yourname)
func (cmService *CategoryMenuService) GetCategoryMenu(ctx context.Context, ID string) (cm platform.CategoryMenu, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&cm).Error
	return
}

// GetCategoryMenuInfoList 分页获取自动化菜单记录
// Author [yourname](https://github.com/yourname)
func (cmService *CategoryMenuService) GetCategoryMenuInfoList(ctx context.Context, info platformReq.CategoryMenuSearch) (list []CustomCategoryMenuTree, err error) {

	// 创建db
	db := global.GVA_DB.Model(&platform.CategoryMenu{})
	var cms []platform.CategoryMenu

	db.Where("project_id = ? AND menu_type = ?", info.ProjectId, info.MenuType)
	//db.Order("id asc")
	db.Debug()
	err = db.Find(&cms).Error
	list = cmService.GetCategoryMenuInfoTree(cms)
	return
}
func (cmService *CategoryMenuService) GetCategoryMenuPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// 自定义树形菜单结构体，仅包含需要返回的字段
type CustomCategoryMenuTree struct {
	ID       uint                     `json:"id"` // 菜单ID
	Key      uint                     `json:"key"`
	Name     string                   `json:"name"` // 菜单名称
	Label    string                   `json:"label"`
	Title    string                   `json:"title"`
	Parent   uint                     `json:"parent"`   // 父菜单ID
	Children []CustomCategoryMenuTree `json:"children"` // 子菜单列表，支持N层嵌套
}

// 递归构建树形结构
func (cmService *CategoryMenuService) GetCategoryMenuInfoTree(list []platform.CategoryMenu) []CustomCategoryMenuTree {
	// 1. 先将所有菜单转换为自定义节点并分组
	nodeMap := make(map[uint]CustomCategoryMenuTree)
	for _, menu := range list {
		// 注意：这里存储的是值类型，后续通过ID查找时会复制
		node := CustomCategoryMenuTree{
			ID:       menu.ID,
			Key:      menu.ID, // 补充Key字段（与ID一致）
			Name:     *menu.Name,
			Title:    *menu.Name,
			Label:    *menu.Name, // 补充Label字段（与Name一致）
			Parent:   *menu.Parent,
			Children: []CustomCategoryMenuTree{},
		}
		nodeMap[menu.ID] = node
	}

	// 2. 递归构建子树
	var buildTree func(parentID uint) []CustomCategoryMenuTree
	buildTree = func(parentID uint) []CustomCategoryMenuTree {
		var children []CustomCategoryMenuTree
		// 遍历所有节点，找出父ID匹配的子节点
		for _, node := range nodeMap {
			if node.Parent == parentID {
				// 递归获取当前节点的子节点
				node.Children = buildTree(node.ID)
				children = append(children, node)
			}
		}
		return children
	}

	// 3. 从根节点（Parent=0）开始构建
	return buildTree(0)
}
