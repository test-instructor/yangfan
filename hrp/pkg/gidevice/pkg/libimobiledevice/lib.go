package libimobiledevice

import (
	"bytes"
	"fmt"

	"github.com/test-instructor/yangfan/server/global"
)

type Packet interface {
	Pack() ([]byte, error)
	Unpack(buffer *bytes.Buffer) (Packet, error)
	Unmarshal(v interface{}) error

	String() string
}

var debugFlag = false

// SetDebug sets debug mode
func SetDebug(debug bool) {
	debugFlag = debug
}

func debugLog(msg string) {
	if !debugFlag {
		return
	}
	global.GVA_LOG.Info(fmt.Sprintf("[%s-debug] %s\n", ProgramName, msg))
}
