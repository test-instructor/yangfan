// 自动生成模板HarmonyDeviceOptions
package platform

import (
	"github.com/test-instructor/yangfan/server/v2/global"
)

// 设备选项 结构体  HarmonyDeviceOptions
type HarmonyDeviceOptions struct {
	global.GVA_MODEL
	Name        *string `json:"name" form:"name" gorm:"column:name;" binding:"required"`                                 //设备名称
	ConnectKey  *string `json:"connectKey" form:"connectKey" gorm:"comment:连接密钥;column:connect_key;" binding:"required"` //连接密钥
	LogOn       *bool   `json:"logOn" form:"logOn" gorm:"comment:是否开启日志;column:log_on;"`                                 //开启日志
	IgnorePopup *bool   `json:"ignorePopup" form:"ignorePopup" gorm:"comment:是否忽略弹窗;column:ignore_popup;"`               //忽略弹窗

	ProjectId int64 `json:"projectId" form:"projectId" gorm:"column:project_id;"` //项目信息
}

// TableName 设备选项 HarmonyDeviceOptions自定义表名 harmony_device_options
func (HarmonyDeviceOptions) TableName() string {
	return "lc_harmony_device_options"
}
