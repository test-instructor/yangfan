// 自动生成模板UserProjectAccess
package projectmgr

import (
	"github.com/test-instructor/yangfan/server/v2/global"
)

// 项目成员与权限 结构体  UserProjectAccess
type UserProjectAccess struct {
	global.GVA_MODEL
	UserId           uint `json:"userId" form:"userId" gorm:"comment:关联用户ID;column:user_id;" binding:"required"`            //用户ID
	ProjectId        uint `json:"projectId" form:"projectId" gorm:"comment:关联项目ID;column:project_id;" binding:"required"`   //项目ID
	AccessPermission bool `json:"accessPermission" form:"accessPermission" gorm:"comment:访问权限标识;column:access_permission;"` //访问权限
	EditPermission   bool `json:"editPermission" form:"editPermission" gorm:"comment:编辑权限标识;column:edit_permission;"`       //编辑权限
	DeletePermission bool `json:"deletePermission" form:"deletePermission" gorm:"comment:删除权限标识;column:delete_permission;"` //删除权限
}

// TableName 项目成员与权限 UserProjectAccess自定义表名 user_project_access
func (UserProjectAccess) TableName() string {
	return "yf_user_project_access"
}
