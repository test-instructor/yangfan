package api

import (
	sysModel "github.com/test-instructor/yangfan/server/v2/model/system"
)

func uiApis() []sysModel.SysApi {
	return []sysModel.SysApi{
		{Path: "/ado/createAndroidDeviceOptions", Description: "新增安卓设备", ApiGroup: "安卓设备", Method: "POST"},
		{Path: "/ado/deleteAndroidDeviceOptions", Description: "删除安卓设备", ApiGroup: "安卓设备", Method: "DELETE"},
		{Path: "/ado/deleteAndroidDeviceOptionsByIds", Description: "批量删除安卓设备", ApiGroup: "安卓设备", Method: "DELETE"},
		{Path: "/ado/updateAndroidDeviceOptions", Description: "更新安卓设备", ApiGroup: "安卓设备", Method: "PUT"},
		{Path: "/ado/findAndroidDeviceOptions", Description: "根据ID获取安卓设备", ApiGroup: "安卓设备", Method: "GET"},
		{Path: "/ado/getAndroidDeviceOptionsList", Description: "获取安卓设备列表", ApiGroup: "安卓设备", Method: "GET"},

		{Path: "/ido/createIOSDeviceOptions", Description: "新增iOS设备", ApiGroup: "iOS设备", Method: "POST"},
		{Path: "/ido/deleteIOSDeviceOptions", Description: "删除iOS设备", ApiGroup: "iOS设备", Method: "DELETE"},
		{Path: "/ido/deleteIOSDeviceOptionsByIds", Description: "批量删除iOS设备", ApiGroup: "iOS设备", Method: "DELETE"},
		{Path: "/ido/updateIOSDeviceOptions", Description: "更新iOS设备", ApiGroup: "iOS设备", Method: "PUT"},
		{Path: "/ido/findIOSDeviceOptions", Description: "根据ID获取iOS设备", ApiGroup: "iOS设备", Method: "GET"},
		{Path: "/ido/getIOSDeviceOptionsList", Description: "获取iOS设备列表", ApiGroup: "iOS设备", Method: "GET"},

		{Path: "/hdo/createHarmonyDeviceOptions", Description: "新增鸿蒙设备", ApiGroup: "鸿蒙设备", Method: "POST"},
		{Path: "/hdo/deleteHarmonyDeviceOptions", Description: "删除鸿蒙设备", ApiGroup: "鸿蒙设备", Method: "DELETE"},
		{Path: "/hdo/deleteHarmonyDeviceOptionsByIds", Description: "批量删除鸿蒙设备", ApiGroup: "鸿蒙设备", Method: "DELETE"},
		{Path: "/hdo/updateHarmonyDeviceOptions", Description: "更新鸿蒙设备", ApiGroup: "鸿蒙设备", Method: "PUT"},
		{Path: "/hdo/findHarmonyDeviceOptions", Description: "根据ID获取鸿蒙设备", ApiGroup: "鸿蒙设备", Method: "GET"},
		{Path: "/hdo/getHarmonyDeviceOptionsList", Description: "获取鸿蒙设备列表", ApiGroup: "鸿蒙设备", Method: "GET"},

		{Path: "/bdo/createBrowserDeviceOptions", Description: "新增浏览器设备", ApiGroup: "浏览器设备", Method: "POST"},
		{Path: "/bdo/deleteBrowserDeviceOptions", Description: "删除浏览器设备", ApiGroup: "浏览器设备", Method: "DELETE"},
		{Path: "/bdo/deleteBrowserDeviceOptionsByIds", Description: "批量删除浏览器设备", ApiGroup: "浏览器设备", Method: "DELETE"},
		{Path: "/bdo/updateBrowserDeviceOptions", Description: "更新浏览器设备", ApiGroup: "浏览器设备", Method: "PUT"},
		{Path: "/bdo/findBrowserDeviceOptions", Description: "根据ID获取浏览器设备", ApiGroup: "浏览器设备", Method: "GET"},
		{Path: "/bdo/getBrowserDeviceOptionsList", Description: "获取浏览器设备列表", ApiGroup: "浏览器设备", Method: "GET"},
	}
}
