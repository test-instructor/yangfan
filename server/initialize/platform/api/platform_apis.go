package api

import (
	sysModel "github.com/test-instructor/yangfan/server/v2/model/system"
)

func platformApis() []sysModel.SysApi {
	return []sysModel.SysApi{
		{Path: "/env/createEnv", Description: "新增环境配置", ApiGroup: "环境配置", Method: "POST"},
		{Path: "/env/deleteEnv", Description: "删除环境配置", ApiGroup: "环境配置", Method: "DELETE"},
		{Path: "/env/deleteEnvByIds", Description: "批量删除环境配置", ApiGroup: "环境配置", Method: "DELETE"},
		{Path: "/env/updateEnv", Description: "更新环境配置", ApiGroup: "环境配置", Method: "PUT"},
		{Path: "/env/findEnv", Description: "根据ID获取环境配置", ApiGroup: "环境配置", Method: "GET"},
		{Path: "/env/getEnvList", Description: "获取环境配置列表", ApiGroup: "环境配置", Method: "GET"},

		{Path: "/ed/createEnvDetail", Description: "新增环境详情", ApiGroup: "环境详情", Method: "POST"},
		{Path: "/ed/deleteEnvDetail", Description: "删除环境详情", ApiGroup: "环境详情", Method: "DELETE"},
		{Path: "/ed/deleteEnvDetailByIds", Description: "批量删除环境详情", ApiGroup: "环境详情", Method: "DELETE"},
		{Path: "/ed/updateEnvDetail", Description: "更新环境详情", ApiGroup: "环境详情", Method: "PUT"},
		{Path: "/ed/findEnvDetail", Description: "根据ID获取环境详情", ApiGroup: "环境详情", Method: "GET"},
		{Path: "/ed/getEnvDetailList", Description: "获取环境详情列表", ApiGroup: "环境详情", Method: "GET"},

		{Path: "/pcd/createPythonCodeDebug", Description: "新增调试信息", ApiGroup: "调试信息", Method: "POST"},
		{Path: "/pcd/deletePythonCodeDebug", Description: "删除调试信息", ApiGroup: "调试信息", Method: "DELETE"},
		{Path: "/pcd/deletePythonCodeDebugByIds", Description: "批量删除调试信息", ApiGroup: "调试信息", Method: "DELETE"},
		{Path: "/pcd/updatePythonCodeDebug", Description: "更新调试信息", ApiGroup: "调试信息", Method: "PUT"},
		{Path: "/pcd/findPythonCodeDebug", Description: "根据ID获取调试信息", ApiGroup: "调试信息", Method: "GET"},
		{Path: "/pcd/getPythonCodeDebugList", Description: "获取调试信息列表", ApiGroup: "调试信息", Method: "GET"},

		{Path: "/pc/createPythonCode", Description: "新增python 函数", ApiGroup: "python 函数", Method: "POST"},
		{Path: "/pc/deletePythonCode", Description: "删除python 函数", ApiGroup: "python 函数", Method: "DELETE"},
		{Path: "/pc/deletePythonCodeByIds", Description: "批量删除python 函数", ApiGroup: "python 函数", Method: "DELETE"},
		{Path: "/pc/updatePythonCode", Description: "更新python 函数", ApiGroup: "python 函数", Method: "PUT"},
		{Path: "/pc/findPythonCode", Description: "根据ID获取python 函数", ApiGroup: "python 函数", Method: "GET"},
		{Path: "/pc/getPythonCodeList", Description: "获取python 函数列表", ApiGroup: "python 函数", Method: "GET"},

		{Path: "/pp/findPythonPackageVersion", Description: "查询py第三方库版本", ApiGroup: "py 第三方库", Method: "GET"},
		{Path: "/pp/createPythonPackage", Description: "新增py 第三方库", ApiGroup: "py 第三方库", Method: "POST"},
		{Path: "/pp/deletePythonPackage", Description: "删除py 第三方库", ApiGroup: "py 第三方库", Method: "DELETE"},
		{Path: "/pp/deletePythonPackageByIds", Description: "批量删除py 第三方库", ApiGroup: "py 第三方库", Method: "DELETE"},
		{Path: "/pp/updatePythonPackage", Description: "更新py 第三方库", ApiGroup: "py 第三方库", Method: "PUT"},
		{Path: "/pp/findPythonPackage", Description: "根据ID获取py 第三方库", ApiGroup: "py 第三方库", Method: "GET"},
		{Path: "/pp/getPythonPackageList", Description: "获取py 第三方库列表", ApiGroup: "py 第三方库", Method: "GET"},

		{Path: "/pcf/createPythonCodeFunc", Description: "新增python函数详情", ApiGroup: "python函数详情", Method: "POST"},
		{Path: "/pcf/deletePythonCodeFunc", Description: "删除python函数详情", ApiGroup: "python函数详情", Method: "DELETE"},
		{Path: "/pcf/deletePythonCodeFuncByIds", Description: "批量删除python函数详情", ApiGroup: "python函数详情", Method: "DELETE"},
		{Path: "/pcf/updatePythonCodeFunc", Description: "更新python函数详情", ApiGroup: "python函数详情", Method: "PUT"},
		{Path: "/pcf/findPythonCodeFunc", Description: "根据ID获取python函数详情", ApiGroup: "python函数详情", Method: "GET"},
		{Path: "/pcf/getPythonCodeFuncList", Description: "获取python函数详情列表", ApiGroup: "python函数详情", Method: "GET"},

		{Path: "/rc/createRunConfig", Description: "新增运行配置", ApiGroup: "运行配置", Method: "POST"},
		{Path: "/rc/deleteRunConfig", Description: "删除运行配置", ApiGroup: "运行配置", Method: "DELETE"},
		{Path: "/rc/deleteRunConfigByIds", Description: "批量删除运行配置", ApiGroup: "运行配置", Method: "DELETE"},
		{Path: "/rc/updateRunConfig", Description: "更新运行配置", ApiGroup: "运行配置", Method: "PUT"},
		{Path: "/rc/findRunConfig", Description: "根据ID获取运行配置", ApiGroup: "运行配置", Method: "GET"},
		{Path: "/rc/getRunConfigList", Description: "获取运行配置列表", ApiGroup: "运行配置", Method: "GET"},

		{Path: "/cm/createCategoryMenu", Description: "新增自动化菜单", ApiGroup: "自动化菜单", Method: "POST"},
		{Path: "/cm/deleteCategoryMenu", Description: "删除自动化菜单", ApiGroup: "自动化菜单", Method: "DELETE"},
		{Path: "/cm/deleteCategoryMenuByIds", Description: "批量删除自动化菜单", ApiGroup: "自动化菜单", Method: "DELETE"},
		{Path: "/cm/updateCategoryMenu", Description: "更新自动化菜单", ApiGroup: "自动化菜单", Method: "PUT"},
		{Path: "/cm/findCategoryMenu", Description: "根据ID获取自动化菜单", ApiGroup: "自动化菜单", Method: "GET"},
		{Path: "/cm/getCategoryMenuList", Description: "获取自动化菜单列表", ApiGroup: "自动化菜单", Method: "GET"},

		{Path: "/rn/createRunnerNode", Description: "新增节点", ApiGroup: "节点", Method: "POST"},
		{Path: "/rn/deleteRunnerNode", Description: "删除节点", ApiGroup: "节点", Method: "DELETE"},
		{Path: "/rn/deleteRunnerNodeByIds", Description: "批量删除节点", ApiGroup: "节点", Method: "DELETE"},
		{Path: "/rn/updateRunnerNode", Description: "更新节点", ApiGroup: "节点", Method: "PUT"},
		{Path: "/rn/findRunnerNode", Description: "根据ID获取节点", ApiGroup: "节点", Method: "GET"},
		{Path: "/rn/getRunnerNodeList", Description: "获取节点列表", ApiGroup: "节点", Method: "GET"},

		{Path: "/runner/api", Description: "接口自动化运行", ApiGroup: "运行服务", Method: "POST"},
	}
}
