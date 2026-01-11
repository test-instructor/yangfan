// 自动生成模板Project
package projectmgr

import (
	"github.com/test-instructor/yangfan/server/v2/global"
)

// 项目配置 结构体  Project
type Project struct {
	global.GVA_MODEL
	Name     string `json:"name" form:"name" gorm:"comment:项目名称;column:name;" binding:"required"` //名称
	Admin    uint   `json:"admin" form:"admin" gorm:"column:admin;" binding:"required"`           //管理员
	Creator  uint   `json:"creator" form:"creator" gorm:"column:creator;"`                        //创建人
	Describe string `json:"describe" form:"describe" gorm:"column:describe;"`                     //描述
	UUID     string `json:"uuid" form:"uuid" gorm:"column:uuid;"`                                 //唯一标识
	Secret   string `json:"secret" form:"secret" gorm:"column:secret;"`                           //密钥
	Logo     string `json:"logo" form:"logo" gorm:"column:logo;"`                                 //logo
}

// TableName 项目配置 Project自定义表名 project
func (Project) TableName() string {
	return "lc_project"
}
