package api

import "github.com/test-instructor/yangfan/data/api/datawarehouse"

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	DatawarehouseApiGroup datawarehouse.ApiGroup
}
