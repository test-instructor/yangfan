// 自动生成模板ApiMenu
package interfacecase

import (
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/model/system"
)

type MenuType int

// ApiMenu 结构体
// 如果含有time.Time 请自行import time包
type ApiMenu struct {
	global.GVA_MODEL
	Operator
	Name      string         `json:"name" form:"name" gorm:"column:name;comment:菜单名称;"`
	Parent    uint           `json:"parent" form:"parent" gorm:"column:parent;comment:父节点id;"`
	MenuType  string         `json:"menuType" form:"menuType" gorm:"column:menu_type;comment:菜单类型;"`
	ProjectID uint           `json:"-" gorm:"comment:所属项目"`
	Project   system.Project `json:"-"`
}
