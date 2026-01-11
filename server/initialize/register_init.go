package initialize

import (
	_ "github.com/test-instructor/yangfan/server/v2/source/example"
	_ "github.com/test-instructor/yangfan/server/v2/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
