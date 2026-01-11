package request

import (
	"time"

	"github.com/test-instructor/yangfan/server/v2/model/common/request"
	datawarehouse "github.com/test-instructor/yangfan/server/v2/model/datawarehouse"
)

type DataCategoryManagementSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	request.PageInfo
	ProjectId int64 `json:"projectId" form:"projectId"`
}

// DataCategoryManagementSave 创建/更新请求
type DataCategoryManagementSave struct {
	ID                uint                                    `json:"ID"`
	Name              *string                                 `json:"name"`
	Type              *string                                 `json:"type"`  // 数据类型
	Count             map[int]int64                           `json:"count"` // 各环境总数量 {env_id: 数量}
	CreateCallType    *int64                                  `json:"createCallType"`
	CreateTestStepId  *uint                                   `json:"createTestStepId"`
	CleanCallType     *int64                                  `json:"cleanCallType"`
	CleanTestStepId   *uint                                   `json:"cleanTestStepId"`
	DirectDelete      *bool                                   `json:"directDelete"`
	PythonCodes       map[string]datawarehouse.PythonCodeInfo `json:"pythonCodes"` // key=envId, value=代码及元信息
	ProjectId         int64                                   `json:"projectId"`
	CreateRunConfigId *uint                                   `json:"createRunConfigId"`
	CleanRunConfigId  *uint                                   `json:"cleanRunConfigId"`
}
