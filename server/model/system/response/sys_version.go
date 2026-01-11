package response

import (
	"github.com/test-instructor/yangfan/server/v2/model/system"
	"github.com/test-instructor/yangfan/server/v2/model/system/request"
)

// ExportVersionResponse 导出版本响应结构体
type ExportVersionResponse struct {
	Version      request.VersionInfo    `json:"version"`      // 版本信息
	Menus        []system.SysBaseMenu   `json:"menus"`        // 菜单数据，直接复用SysBaseMenu
	Apis         []system.SysApi        `json:"apis"`         // API数据，直接复用SysApi
	Dictionaries []system.SysDictionary `json:"dictionaries"` // 字典数据，直接复用SysDictionary
}
