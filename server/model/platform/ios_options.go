// 自动生成模板IOSDeviceOptions
package platform

import (
	"github.com/test-instructor/yangfan/server/v2/global"
)

// iOS设备选项 结构体  IOSDeviceOptions
type IOSDeviceOptions struct {
	global.GVA_MODEL
	Name                       *string `json:"name" form:"name" gorm:"column:name;" binding:"required"`                                                                                       //设备名称
	UDID                       *string `json:"udid" form:"udid" gorm:"comment:设备唯一标识符;column:udid;size:255;"`                                                                                 //设备标识
	Wireless                   *bool   `json:"wireless" form:"wireless" gorm:"comment:是否使用无线连接;column:wireless;size:1;"`                                                                      //无线连接
	WDAPort                    *int64  `json:"port" form:"port" gorm:"comment:WDA远程端口;column:wda_port;"`                                                                                      //WDA端口
	WDAMjpegPort               *int64  `json:"mjpeg_port" form:"mjpeg_port" gorm:"comment:WDA远程MJPEG端口;column:wda_mjpeg_port;"`                                                               //WDA MJPEG 端口
	LogOn                      *bool   `json:"log_on" form:"log_on" gorm:"comment:是否开启日志;column:log_on;size:1;"`                                                                              //开启日志
	IgnorePopup                *bool   `json:"ignore_popup" form:"ignore_popup" gorm:"comment:是否忽略弹窗;column:ignore_popup;size:1;"`                                                            //忽略弹窗
	ResetHomeOnStartup         *bool   `json:"reset_home_on_startup" form:"reset_home_on_startup" gorm:"comment:启动时是否重置到主页;column:reset_home_on_startup;size:1;"`                             //重置主页
	SnapshotMaxDepth           *int64  `json:"snapshot_max_depth" form:"snapshot_max_depth" gorm:"comment:元素快照的最大深度;column:snapshot_max_depth;"`                                              //快照深度
	AcceptAlertButtonSelector  *string `json:"accept_alert_button_selector" form:"accept_alert_button_selector" gorm:"comment:接受警告弹窗按钮选择器;column:accept_alert_button_selector;size:1000;"`    //接受按钮选择器
	DismissAlertButtonSelector *string `json:"dismiss_alert_button_selector" form:"dismiss_alert_button_selector" gorm:"comment:取消警告弹窗按钮选择器;column:dismiss_alert_button_selector;size:1000;"` //取消按钮选择器

	ProjectId int64 `json:"projectId" form:"projectId" gorm:"column:project_id;"` //项目信息
}

// TableName iOS设备选项 IOSDeviceOptions自定义表名 ios_device_options
func (IOSDeviceOptions) TableName() string {
	return "lc_ios_device_options"
}
