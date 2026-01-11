// 自动生成模板EnvDetail
package platform

import (
	"github.com/test-instructor/yangfan/server/v2/global"
	"gorm.io/datatypes"
)

// 环境详情 结构体  EnvDetail
type EnvDetail struct {
	global.GVA_MODEL
	Key       string            `json:"key" form:"key" gorm:"column:key;" binding:"required"`                                                                     //变量名称
	Name      string            `json:"name" form:"name" gorm:"column:name;" binding:"required"`                                                                  //中文名
	Value     datatypes.JSONMap `json:"value" form:"value" gorm:"type:JSON;comment:json格式，key为环境id，值为变量的值;column:value;" swaggertype:"object" binding:"required"` //变量值
	ProjectId int64             `json:"projectId" form:"projectId" gorm:"column:project_id;"`                                                                     //项目信息
}

// TableName 环境详情 EnvDetail自定义表名 env_detail
func (EnvDetail) TableName() string {
	return "lc_env_detail"
}
