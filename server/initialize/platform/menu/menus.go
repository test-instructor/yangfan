package menu

import (
	sysModel "github.com/test-instructor/yangfan/server/v2/model/system"
)

type MenuSeed struct {
	Path       string
	ParentPath string
	Name       string
	Hidden     bool
	Component  string
	Sort       int
	Meta       sysModel.Meta
}

func Seeds() []MenuSeed {
	return []MenuSeed{
		// Root menus
		{Path: "platform", ParentPath: "", Name: "platform", Hidden: false, Component: "view/routerHolder.vue", Sort: 200, Meta: sysModel.Meta{Title: "配置管理", Icon: "setting"}},
		{Path: "APIAutomation", ParentPath: "", Name: "APIAutomation", Hidden: false, Component: "view/routerHolder.vue", Sort: 400, Meta: sysModel.Meta{Title: "接口自动化", Icon: "connection"}},
		{Path: "pm", ParentPath: "", Name: "pm", Hidden: false, Component: "view/routerHolder.vue", Sort: 100, Meta: sysModel.Meta{Title: "项目管理", Icon: "project"}},
		{Path: "pc3", ParentPath: "", Name: "pc3", Hidden: false, Component: "view/platform/pc/pc3.vue", Sort: 0, Meta: sysModel.Meta{Title: "代码测试3", Icon: ""}},
		{Path: "dataWarehouse", ParentPath: "", Name: "dataWarehouse", Hidden: false, Component: "view/routerHolder.vue", Sort: 300, Meta: sysModel.Meta{Title: "数据仓库", Icon: "dataWarehouse"}},

		// pm children
		{Path: "pj", ParentPath: "pm", Name: "pj", Hidden: false, Component: "view/projectmgr/project/project.vue", Sort: 100, Meta: sysModel.Meta{Title: "项目配置", Icon: "project"}},
		{Path: "upa", ParentPath: "pm", Name: "upa", Hidden: false, Component: "view/projectmgr/userprojectaccess/userprojectaccess.vue", Sort: 200, Meta: sysModel.Meta{Title: "项目成员与权限", Icon: "auth"}},
		{Path: "reportNotify", ParentPath: "pm", Name: "reportNotify", Hidden: false, Component: "view/projectmgr/reportNotify/reportNotify.vue", Sort: 300, Meta: sysModel.Meta{Title: "报告通知", Icon: "warn"}},

		// platform children
		{Path: "envDetail", ParentPath: "platform", Name: "envDetail", Hidden: false, Component: "view/platform/envdetail/envdetail.vue", Sort: 200, Meta: sysModel.Meta{Title: "环境变量管理", Icon: "env"}},
		{Path: "FunctionPlugin", ParentPath: "platform", Name: "FunctionPlugin", Hidden: false, Component: "view/routerHolder.vue", Sort: 300, Meta: sysModel.Meta{Title: "函数插件", Icon: "Plugin"}},
		{Path: "rc", ParentPath: "platform", Name: "rc", Hidden: false, Component: "view/platform/runconfig/runconfig.vue", Sort: 100, Meta: sysModel.Meta{Title: "运行配置", Icon: "un-config-o"}},
		{Path: "rn", ParentPath: "platform", Name: "rn", Hidden: false, Component: "view/platform/runnernode/runnernode.vue", Sort: 499, Meta: sysModel.Meta{Title: "运行节点", Icon: ""}},

		// FunctionPlugin children
		{Path: "pc", ParentPath: "FunctionPlugin", Name: "pc", Hidden: false, Component: "view/platform/pc/pc.vue", Sort: 200, Meta: sysModel.Meta{Title: "python 函数", Icon: "Function"}},
		{Path: "pcd", ParentPath: "FunctionPlugin", Name: "pcd", Hidden: false, Component: "view/platform/pythoncodedebug/pythoncodedebug.vue", Sort: 300, Meta: sysModel.Meta{Title: "调试信息", Icon: "debug"}},
		{Path: "pp", ParentPath: "FunctionPlugin", Name: "pp", Hidden: false, Component: "view/platform/pythonPackage/pythonPackage.vue", Sort: 300, Meta: sysModel.Meta{Title: "py 第三方库", Icon: "package-line"}},

		// APIAutomation children
		{Path: "as", ParentPath: "APIAutomation", Name: "as", Hidden: false, Component: "view/automation/autostep/autostep.vue", Sort: 100, Meta: sysModel.Meta{Title: "接口管理", Icon: "api"}},
		{Path: "acs", ParentPath: "APIAutomation", Name: "acs", Hidden: false, Component: "view/automation/autocasestep/autocasestep.vue", Sort: 200, Meta: sysModel.Meta{Title: "测试步骤", Icon: "case-step"}},
		{Path: "ac", ParentPath: "APIAutomation", Name: "ac", Hidden: false, Component: "view/automation/autocase/autocase.vue", Sort: 300, Meta: sysModel.Meta{Title: "测试用例", Icon: "testcase"}},
		{Path: "tk", ParentPath: "APIAutomation", Name: "tk", Hidden: false, Component: "view/automation/timertask/timertask.vue", Sort: 400, Meta: sysModel.Meta{Title: "定时任务", Icon: "time-task"}},
		{Path: "ar", ParentPath: "APIAutomation", Name: "ar", Hidden: false, Component: "view/automation/autoreport/autoreport.vue", Sort: 500, Meta: sysModel.Meta{Title: "自动报告", Icon: "bxs-report"}},
		{Path: "auto-report-detail/:id", ParentPath: "APIAutomation", Name: "auto-report-detail", Hidden: true, Component: "view/automation/autoreport/AutoReportDetail.vue", Sort: 10000, Meta: sysModel.Meta{Title: "运行报告详情", Icon: ""}},

		// dataWarehouse children
		{Path: "dcm", ParentPath: "dataWarehouse", Name: "dcm", Hidden: false, Component: "view/datawarehouse/dataCategoryManagement/dataCategoryManagement.vue", Sort: 100, Meta: sysModel.Meta{Title: "数据分类", Icon: "dataType"}},
	}
}
