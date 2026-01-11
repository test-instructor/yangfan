package ast

import (
	"github.com/test-instructor/yangfan/server/v2/global"
	"path/filepath"
)

func init() {
	global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("../../../")
	global.GVA_CONFIG.AutoCode.Server = "server"
}
