// 自动生成模板PythonCodeFunc
package platform

import (
	"github.com/test-instructor/yangfan/server/v2/global"
	"gorm.io/datatypes"
)

// python函数详情 结构体  PythonCodeFunc
type PythonCodeFunc struct {
	global.GVA_MODEL
	Name       string         `json:"name" form:"name" gorm:"column:name;" binding:"required"`                    //名称
	Params     datatypes.JSON `json:"params" form:"params" gorm:"column:params;" swaggertype:"array,object"`      //参数
	FullCode   string         `json:"fullCode" form:"fullCode" gorm:"column:full_code;" binding:"required"`       //完整代码
	StartIndex int64          `json:"startIndex" form:"startIndex" gorm:"column:start_index;" binding:"required"` //起始索引
	ProjectId  int64          `json:"projectId" form:"projectId" gorm:"column:project_id;"`                       //项目信息
}

// TableName python函数详情 PythonCodeFunc自定义表名 python_code_func
func (PythonCodeFunc) TableName() string {
	return "yf_python_code_func"
}
