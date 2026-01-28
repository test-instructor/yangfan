// 自动生成模板BrowserDeviceOptions
package platform

import (
	"github.com/test-instructor/yangfan/server/v2/global"
)

// 浏览器设备选项 结构体  BrowserDeviceOptions
type BrowserDeviceOptions struct {
	global.GVA_MODEL
	BrowserID   *string `json:"browserId" form:"browserId" gorm:"comment:浏览器标识ID;column:browser_id;size:255;"`    //浏览器标识
	LogOn       *bool   `json:"logOn" form:"logOn" gorm:"comment:是否登录状态;column:log_on;size:1;"`                   //登录状态
	IgnorePopup *bool   `json:"ignorePopup" form:"ignorePopup" gorm:"comment:是否忽略弹窗;column:ignore_popup;size:1;"` //忽略弹窗
	Width       *int64  `json:"width" form:"width" gorm:"comment:浏览器宽度;column:width;"`                            //宽度
	Height      *int64  `json:"height" form:"height" gorm:"comment:浏览器高度;column:height;"`                         //高度

	ProjectId int64 `json:"projectId" form:"projectId" gorm:"column:project_id;"` //项目信息
}

// TableName 浏览器设备选项 BrowserDeviceOptions自定义表名 browser_device_options
func (BrowserDeviceOptions) TableName() string {
	return "lc_browser_device_options"
}
