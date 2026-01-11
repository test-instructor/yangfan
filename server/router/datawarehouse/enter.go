package datawarehouse

import api "github.com/test-instructor/yangfan/server/v2/api/v1"

type RouterGroup struct{ DataCategoryManagementRouter }

var dcmApi = api.ApiGroupApp.DatawarehouseApiGroup.DataCategoryManagementApi
