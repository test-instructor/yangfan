// 自动生成模板DataCategoryManagement
package datawarehouse

import (
	"time"

	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	"gorm.io/datatypes"
)

// PythonCodeInfo 用于前端交互的 Python 代码信息
// 包含代码内容和最后一次更新时间（来自 yf_python_code.UpdatedAt）
type PythonCodeInfo struct {
	Code     string     `json:"code"`                // 代码内容
	UpdateAt *time.Time `json:"update_at,omitempty"` // 最后更新时间
}

// 数据分类 结构体  DataCategoryManagement
type DataCategoryManagement struct {
	global.GVA_MODEL
	Name              *string           `json:"name" form:"name" gorm:"comment:数据分类名称;column:name;size:255;"`                                                                       // 数据分类名称
	Type              *string           `json:"type" form:"type" gorm:"comment:数据类型;column:type;size:255;"`                                                                         // 数据类型
	Count             datatypes.JSONMap `json:"count" form:"count" gorm:"type:JSON;comment:各环境数据总数量，key为env_id字符串;column:count;" swaggertype:"object"`                              // 各环境总数量 {env_id: 数量}
	AvailableCount    datatypes.JSONMap `json:"availableCount" form:"availableCount" gorm:"type:JSON;comment:各环境可用数据数量，key为env_id字符串;column:available_count;" swaggertype:"object"` // 各环境可用数量 {env_id: 数量}
	CreateCallType    *int64            `json:"createCallType" form:"createCallType" gorm:"comment:创建数据的调用类型(1测试步骤,2Python);column:create_call_type;"`                              // 创建数据的调用类型
	CreateTestStepId  *uint             `json:"createTestStepId" form:"createTestStepId" gorm:"comment:创建数据的测试步骤ID;column:create_test_step_id;"`                                    // 创建数据的测试步骤ID
	CleanCallType     *int64            `json:"cleanCallType" form:"cleanCallType" gorm:"comment:清洗数据的调用类型(1测试步骤,2Python,3直接删除);column:clean_call_type;"`                           // 清洗数据的调用类型
	CleanTestStepId   *uint             `json:"cleanTestStepId" form:"cleanTestStepId" gorm:"comment:清洗数据的测试步骤ID;column:clean_test_step_id;"`                                       // 清洗数据的测试步骤ID
	DirectDelete      *bool             `json:"directDelete" form:"directDelete" gorm:"comment:是否直接删除数据;column:direct_delete;"`                                                     // 是否直接删除数据
	ProjectId         int64             `json:"projectId" form:"projectId" gorm:"column:project_id;index;"`                                                                         // 项目信息
	CreateRunConfigId *uint             `json:"createRunConfigId" form:"createRunConfigId" gorm:"comment:创建数据的运行配置ID;column:create_run_config_id;"`                                 // 创建数据的运行配置ID
	CleanRunConfigId  *uint             `json:"cleanRunConfigId" form:"cleanRunConfigId" gorm:"comment:清洗数据的运行配置ID;column:clean_run_config_id;"`                                    // 清洗数据的运行配置ID
	LastData          datatypes.JSONMap `json:"lastData" form:"lastData" gorm:"type:JSON;comment:最后一次生成的数据(按环境隔离);column:last_data;" swaggertype:"object"`                          // 最后一次生成的数据 {env_id: {data}}

	// 非存储字段，用于前端交互
	EnvList     []platform.Env            `json:"envList" gorm:"-"`               // 项目下所有环境
	PythonCodes map[string]PythonCodeInfo `json:"pythonCodes,omitempty" gorm:"-"` // key=envId, value=代码内容及元信息
}

// TableName 数据分类 DataCategoryManagement自定义表名 data_category_management
func (DataCategoryManagement) TableName() string {
	return "yf_data_category_management"
}

// ==================== 数据池记录表 ====================

// DataCategoryDataStatus 数据状态常量
const (
	DataStatusAvailable = 0 // 可用
	DataStatusUsed      = 1 // 已占用
	DataStatusCleaned   = 2 // 已清洗/失效
)

// DataCategoryData 数据分类的数据记录（数据池）
// 每条记录代表某个数据分类在某个环境下的一条具体数据
type DataCategoryData struct {
	global.GVA_MODEL
	DataCategoryId uint              `json:"dataCategoryId" form:"dataCategoryId" gorm:"column:data_category_id;index;comment:数据分类ID"` // 关联 DataCategoryManagement.ID
	Type           *string           `json:"type" form:"type" gorm:"column:type;size:255;index;comment:数据类型"`                          // 数据类型
	EnvId          uint              `json:"envId" form:"envId" gorm:"column:env_id;index;comment:环境ID"`                               // 环境 ID
	Value          datatypes.JSONMap `json:"value" form:"value" gorm:"type:JSON;comment:数据内容(JSON格式)" swaggertype:"object"`            // 数据内容，JSON 格式便于存储多字段
	Status         int               `json:"status" form:"status" gorm:"column:status;default:0;index;comment:状态(0可用,1已占用,2已清洗)"`      // 数据状态
	ProjectId      int64             `json:"projectId" form:"projectId" gorm:"column:project_id;index;comment:项目ID"`                   // 项目 ID（冗余便于查询）
	UsedBy         *string           `json:"usedBy" form:"usedBy" gorm:"column:used_by;size:255;comment:占用者(用例名/任务名等)"`                // 被谁占用
	UsedAt         *time.Time        `json:"usedAt" form:"usedAt" gorm:"column:used_at;comment:占用时间"`                                  // 占用时间
	ExpireAt       *time.Time        `json:"expireAt" form:"expireAt" gorm:"column:expire_at;comment:过期时间"`                            // 过期时间（可选）
	CleanedAt      *time.Time        `json:"cleanedAt" form:"cleanedAt" gorm:"column:cleaned_at;comment:清洗时间"`                         // 清洗时间
	Remark         *string           `json:"remark" form:"remark" gorm:"column:remark;size:500;comment:备注"`                            // 备注
}

// TableName 数据池记录表
func (DataCategoryData) TableName() string {
	return "yf_data_category_data"
}
