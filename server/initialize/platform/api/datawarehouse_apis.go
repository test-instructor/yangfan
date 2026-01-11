package api

import (
	sysModel "github.com/test-instructor/yangfan/server/v2/model/system"
)

func datawarehouseApis() []sysModel.SysApi {
	return []sysModel.SysApi{
		{Path: "/dcm/createDataCategoryManagement", Description: "新增数据分类", ApiGroup: "数据分类", Method: "POST"},
		{Path: "/dcm/deleteDataCategoryManagement", Description: "删除数据分类", ApiGroup: "数据分类", Method: "DELETE"},
		{Path: "/dcm/deleteDataCategoryManagementByIds", Description: "批量删除数据分类", ApiGroup: "数据分类", Method: "DELETE"},
		{Path: "/dcm/updateDataCategoryManagement", Description: "更新数据分类", ApiGroup: "数据分类", Method: "PUT"},
		{Path: "/dcm/findDataCategoryManagement", Description: "根据ID获取数据分类", ApiGroup: "数据分类", Method: "GET"},
		{Path: "/dcm/getDataCategoryManagementList", Description: "获取数据分类列表", ApiGroup: "数据分类", Method: "GET"},
		{Path: "/dcm/getDataCategoryTypeList", Description: "获取数据分类类型列表", ApiGroup: "数据分类", Method: "GET"},
	}
}
