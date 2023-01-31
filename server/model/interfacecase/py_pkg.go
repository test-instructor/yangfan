package interfacecase

import "github.com/test-instructor/cheetah/server/global"

type HrpPyPkg struct {
	global.GVA_MODEL
	Name        string `json:"name" form:"name" gorm:"uniqueIndex:idx_name"`
	Version     string `json:"version" form:"version" gorm:"column:version;comment:;"`
	IsUninstall *bool  `json:"isUninstall" form:"isUninstall" gorm:"column:is_uninstall;comment:;"`
}

func (HrpPyPkg) TableName() string {
	return "api_py_pkg"
}
