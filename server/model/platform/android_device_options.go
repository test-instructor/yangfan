// 自动生成模板AndroidDeviceOptions
package platform

import (
	"github.com/test-instructor/yangfan/server/v2/global"
)

// 设备选项 结构体  AndroidDeviceOptions
type AndroidDeviceOptions struct {
	global.GVA_MODEL
	Name                      *string `json:"name" form:"name" gorm:"column:name;"`                                                                                                      //设备名称
	SerialNumber              *string `json:"serial" form:"serial" gorm:"comment:设备序列号;column:serial_number;size:255;" binding:"required"`                                               //序列号
	LogOn                     *bool   `json:"logOn" form:"logOn" gorm:"comment:是否开启日志;column:log_on;size:1;"`                                                                            //日志开关
	IgnorePopup               *bool   `json:"ignorePopup" form:"ignorePopup" gorm:"comment:是否忽略弹窗;column:ignore_popup;size:1;"`                                                          //忽略弹窗
	AdbServerHost             *string `json:"adbServerHost" form:"adbServerHost" gorm:"comment:ADB服务器主机地址;column:adb_server_host;size:255;"`                                             //ADB主机
	AdbServerPort             *int64  `json:"adbServerPort" form:"adbServerPort" gorm:"comment:ADB服务器端口;column:adb_server_port;"`                                                        //ADB端口
	Composite                 *bool   `json:"composite" form:"composite" gorm:"comment:是否启用复合驱动;column:composite;size:1;"`                                                               //复合驱动
	UIA2                      *bool   `json:"uia2" form:"uia2" gorm:"comment:是否使用UIAutomator2;column:uia2;size:1;"`                                                                      //UIA2开关
	UIA2IP                    *string `json:"uia2Ip" form:"uia2Ip" gorm:"comment:UIAutomator2服务器IP;column:uia2_ip;size:255;"`                                                            //UIA2地址
	UIA2Port                  *int64  `json:"uia2Port" form:"uia2Port" gorm:"comment:UIAutomator2服务器端口;column:uia2_port;"`                                                               //UIA2端口
	UIA2ServerPackageName     *string `json:"uia2ServerPackageName" form:"uia2ServerPackageName" gorm:"comment:UIAutomator2服务器包名;column:uia2_server_package_name;size:255;"`             //UIA2服务包名
	UIA2ServerTestPackageName *string `json:"uia2ServerTestPackageName" form:"uia2ServerTestPackageName" gorm:"comment:UIAutomator2测试包名;column:uia2_server_test_package_name;size:255;"` //UIA2测试包名

	ProjectId int64 `json:"projectId" form:"projectId" gorm:"column:project_id;"` //项目信息
}

// TableName 设备选项 AndroidDeviceOptions自定义表名 android_device_options
func (AndroidDeviceOptions) TableName() string {
	return "lc_android_device_options"
}
