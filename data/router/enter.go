package router

import "github.com/test-instructor/yangfan/data/router/datawarehouse"

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	Datawarehouse datawarehouse.RouterGroup
}
