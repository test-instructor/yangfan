// 自动生成模板Env
package platform

import (
	"github.com/test-instructor/yangfan/server/v2/global"
)

// 环境配置 结构体  Env
type Env struct {
	global.GVA_MODEL
	Name      string `json:"name" form:"name" gorm:"column:name;" binding:"required"` //名称
	Remarks   string `json:"remarks" form:"remarks" gorm:"column:remarks;"`           //备注
	ProjectId int64  `json:"projectId" form:"projectId" gorm:"column:project_id;"`    //项目id
}

// TableName 环境配置 Env自定义表名 env
func (Env) TableName() string {
	return "yf_env"
}
