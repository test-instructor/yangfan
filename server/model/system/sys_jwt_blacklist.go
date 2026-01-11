package system

import (
	"github.com/test-instructor/yangfan/server/v2/global"
)

type JwtBlacklist struct {
	global.GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
