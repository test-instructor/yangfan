// 自动生成模板CategoryMenu
package platform

import (
	"github.com/test-instructor/yangfan/server/v2/global"
)

// 自动化菜单 结构体  CategoryMenu
type CategoryMenu struct {
	global.GVA_MODEL
	Name     *string `json:"name" form:"name" gorm:"comment:菜单名称;column:name;"`              //名称
	Parent   *uint   `json:"parent" form:"parent" gorm:"comment:父节点id;column:parent;"`       //父节点
	MenuType *string `json:"menuType" form:"menuType" gorm:"comment:菜单类型;column:menu_type;"` //类型

	ProjectId int64 `json:"projectId" form:"projectId" gorm:"column:project_id;"` //项目信息
	Sort      uint  `json:"sort" form:"sort" gorm:"column:sort;"`
}

// TableName 自动化菜单 CategoryMenu自定义表名 category_menu
func (CategoryMenu) TableName() string {
	return "lc_category_menu"
}
