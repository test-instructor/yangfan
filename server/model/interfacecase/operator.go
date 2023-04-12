package interfacecase

import "github.com/test-instructor/yangfan/server/model/system"

type Operator struct {
	CreatedBy *uint `gorm:"column:created_by_id;comment:创建者"`
	UpdateBy  *uint `gorm:"column:update_by_id;comment:更新者"`
	DeleteBy  *uint `gorm:"column:delete_by_id;comment:删除者"`

	Project   system.Project `json:"-"`
	ProjectID uint           `json:"-"`
	//CreatedBy   *system.SysUser
	//UpdateBy    *system.SysUser
	//DeleteBy    *system.SysUser
}
