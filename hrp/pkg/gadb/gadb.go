package gadb

import (
	"github.com/test-instructor/yangfan/server/global"
)

var debugFlag = false

// SetDebug set debug mode
func SetDebug(debug bool) {
	debugFlag = debug
}

func debugLog(msg string) {
	if !debugFlag {
		return
	}
	global.GVA_LOG.Info("[DEBUG] [gadb] " + msg)
}
