// 自动生成模板PythonCode
package platform

import (
	"github.com/test-instructor/yangfan/server/v2/global"
)

// python 函数 结构体  PythonCode
type PythonCode struct {
	global.GVA_MODEL
	Type      int64  `json:"type" form:"type" gorm:"comment:1为自动化代码,2为数据分类代码;column:type;index;"`              //函数类型
	UniqueKey string `json:"uniqueKey" form:"uniqueKey" gorm:"comment:唯一标识;column:unique_key;size:100;index;"` //唯一标识
	ProjectId int64  `json:"projectId" form:"projectId" gorm:"column:project_id;index;"`                       //项目信息
	Code      string `json:"code" form:"code" gorm:"column:code;type:mediumtext;"`                             //代码
	UpdateBy  uint   `json:"updateBy" form:"updateBy" gorm:"column:update_by;"`                                //更新人
}

// TableName python 函数 PythonCode自定义表名 python_code
func (PythonCode) TableName() string {
	return "lc_python_code"
}
