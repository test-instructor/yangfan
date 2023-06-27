package interfacecase

import "github.com/test-instructor/yangfan/server/model/system"

type Operator struct {
	CreatedBy *uint `json:"CreatedBy,omitempty" gorm:"column:created_by_id;comment:创建者"`
	UpdateBy  *uint `json:"UpdateBy,omitempty" gorm:"column:update_by_id;comment:更新者"`
	DeleteBy  *uint `json:"DeleteBy,omitempty" gorm:"column:delete_by_id;comment:删除者"`

	Project   system.Project `json:"-"`
	ProjectID uint           `json:"project_id",omitempty`
	//CreatedBy   *system.SysUser
	//UpdateBy    *system.SysUser
	//DeleteBy    *system.SysUser
}
