// 自动生成模板Project
package system

import (
	"github.com/test-instructor/yangfan/server/global"
)

// Project 结构体
// 如果含有time.Time 请自行import time包
type Project struct {
	global.GVA_MODEL
	Name      string `json:"name" form:"name" gorm:"column:name;comment:;comment:项目名称"`
	Admin     string `json:"admin" form:"admin" gorm:"column:admin;comment:;comment:项目管理员"`
	Creator   string `json:"creator" form:"creator" gorm:"column:creator;comment:;"`
	Describe  string `json:"describe" form:"describe" gorm:"column:describe;comment:;"`
	SysUserID uint
	SysUser   []SysUser `json:"sys_user" gorm:"many2many:sys_user_project;"`
}
