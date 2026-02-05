package system

import "github.com/test-instructor/yangfan/server/v2/global"

type SysUINodeMenu struct {
	global.GVA_MODEL
	MenuLevel  uint   `json:"-"`
	ParentId   uint   `json:"parentId" gorm:"comment:父菜单ID"`
	Path       string `json:"path" gorm:"comment:路由path"`
	Name       string `json:"name" gorm:"comment:路由name"`
	Hidden     bool   `json:"hidden" gorm:"comment:是否在列表隐藏"`
	Component  string `json:"component" gorm:"comment:UI node前端组件key"`
	Sort       int    `json:"sort" gorm:"comment:排序标记"`
	Meta       `json:"meta" gorm:"embedded;comment:附加属性"`
	Children   []SysUINodeMenu          `json:"children" gorm:"-"`
	Parameters []SysUINodeMenuParameter `json:"parameters"`
}

type SysUINodeMenuParameter struct {
	global.GVA_MODEL
	SysUINodeMenuID uint
	Type            string `json:"type" gorm:"comment:地址栏携带参数为params还是query"`
	Key             string `json:"key" gorm:"comment:地址栏携带参数的key"`
	Value           string `json:"value" gorm:"comment:地址栏携带参数的值"`
}

func (SysUINodeMenu) TableName() string {
	return "sys_ui_node_menus"
}

func (SysUINodeMenuParameter) TableName() string {
	return "sys_ui_node_menu_parameters"
}
