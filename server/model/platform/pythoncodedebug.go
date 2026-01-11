// 自动生成模板PythonCodeDebug
package platform

import (
	"github.com/test-instructor/yangfan/server/v2/global"
	"gorm.io/datatypes"
)

// 调试信息 结构体  PythonCodeDebug
type PythonCodeDebug struct {
	global.GVA_MODEL
	Function     string            `json:"function" form:"function" gorm:"column:function;"`                                  //函数
	Parameters   datatypes.JSONMap `json:"parameters" form:"parameters" gorm:"column:parameters;" swaggertype:"object"`       //参数
	ReturnValue  datatypes.JSONMap `json:"return_value" form:"return_value" gorm:"column:return_value;" swaggertype:"object"` //返回值
	Type         int64             `json:"type" form:"type" gorm:"column:type;"`                                              //类型
	Debugger     uint              `json:"debugger" form:"debugger" gorm:"column:debugger;"`                                  //调试人
	ProjectId    int64             `json:"projectId" form:"projectId" gorm:"column:project_id;"`                              //项目信息
	PythonCodeId int64             `json:"pythonCodeId" form:"pythonCodeId" gorm:"column:python_code_id;"`                    //代码id
}

// TableName 调试信息 PythonCodeDebug自定义表名 python_code_debug
func (PythonCodeDebug) TableName() string {
	return "lc_python_code_debug"
}
