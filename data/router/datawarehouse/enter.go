package datawarehouse

import "github.com/test-instructor/yangfan/data/api"

type RouterGroup struct{ DataQueryRouter }

var dataQueryApi = api.ApiGroupApp.DatawarehouseApiGroup.DataQueryApi
