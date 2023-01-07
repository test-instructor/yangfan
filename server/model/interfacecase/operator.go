package interfacecase

import "github.com/test-instructor/cheetah/server/model/system"

type Operator struct {
	CreatedByID *uint `gorm:"comment:创建者"`
	UpdateByID  *uint `gorm:"comment:更新者"`
	DeleteByID  *uint `gorm:"comment:删除者"`
	CreatedBy   *system.SysUser
	UpdateBy    *system.SysUser
	DeleteBy    *system.SysUser
}
