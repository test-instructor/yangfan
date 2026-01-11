// 自动生成模板PythonPackage
package platform

import (
	"github.com/test-instructor/yangfan/server/v2/global"
)

// py 第三方库 结构体  PythonPackage
type PythonPackage struct {
	global.GVA_MODEL
	Name        *string `json:"name" form:"name" gorm:"column:name;" binding:"required"`             //包名
	Version     string  `json:"version" form:"version" gorm:"column:version;" binding:"required"`    //版本
	Description string  `json:"description" form:"description" gorm:"column:description;type:text;"` //描述
	UpdateBy    int64   `json:"updateBy" form:"updateBy" gorm:"column:update_by;"`                   //更新人
	Type        int64   `json:"type" form:"type" gorm:"column:type;"`                                //类型
	ProjectId   int64   `json:"projectId" form:"projectId" gorm:"column:project_id;"`
}

// TableName py 第三方库 PythonPackage自定义表名 python_package
func (PythonPackage) TableName() string {
	return "lc_python_package"
}
