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
	Name        string         `json:"name" form:"name" gorm:"column:name;comment:;"`
	Parent      uint           `json:"parent" form:"parent" gorm:"column:parent;comment:;"`
	MenuType    string         `json:"menuType" form:"menuType" gorm:"column:menu_type;comment:;"`
	ProjectID   uint           `json:"-"`
	Project     system.Project `json:"-"`
	CreatedBy   system.SysUser `json:"-"`
	CreatedByID uint           `json:"-"`
	UpdateBy    system.SysUser `json:"-"`
	UpdateByID  uint           `json:"-"`
	DeleteBy    system.SysUser `json:"-"`
	DeleteByID  uint           `json:"-"`
}
