package interfacecase

import "github.com/test-instructor/yangfan/server/global"

type ApiUserConfig struct {
	global.GVA_MODEL
	Operator
	ApiConfig   *ApiConfig `json:"api_config"`
	ApiConfigID *uint      `json:"api_config_id" gorm:"comment:接口配置ID"`
	ApiEnv      *ApiEnv    `json:"api_env"`
	ApiEnvID    *uint      `json:"api_env_id"`
	UserID      uint       `json:"-" gorm:"comment:用户ID"`
}
