package menu

import sysModel "github.com/test-instructor/yangfan/server/v2/model/system"

type MenuSeed struct {
	Path       string
	ParentName string
	Name       string
	Hidden     bool
	Component  string
	Sort       int
	Meta       sysModel.Meta
}

func Seeds() []MenuSeed {
	return []MenuSeed{
		{Path: "/home", ParentName: "", Name: "home", Hidden: false, Component: "home/index", Sort: 100, Meta: sysModel.Meta{Title: "menu.common.home", Icon: "home"}},
		{Path: "/settings", ParentName: "", Name: "settings", Hidden: false, Component: "settings/index", Sort: 200, Meta: sysModel.Meta{Title: "menu.common.settings", Icon: "settings"}},
		{Path: "/person", ParentName: "", Name: "person", Hidden: true, Component: "person/index", Sort: 300, Meta: sysModel.Meta{Title: "menu.common.profile", Icon: "user"}},

		{Path: "/androidui", ParentName: "", Name: "androidui", Hidden: false, Component: "RouterView", Sort: 400, Meta: sysModel.Meta{Title: "menu.androidui.root", Icon: "android"}},
		{Path: "/androidui/androidDeviceOptions", ParentName: "androidui", Name: "androidDeviceOptions", Hidden: false, Component: "androidui/deviceOptions/index", Sort: 100, Meta: sysModel.Meta{Title: "menu.androidui.deviceManagement", Icon: "android"}},
		{Path: "/androidui/runConfig", ParentName: "androidui", Name: "runConfigAndroid", Hidden: false, Component: "androidui/runConfig/index", Sort: 200, Meta: sysModel.Meta{Title: "menu.androidui.runConfig", Icon: "setting"}},
		{Path: "/androidui/autoStep", ParentName: "androidui", Name: "AutoStepAndroid", Hidden: false, Component: "androidui/autoStep/index", Sort: 300, Meta: sysModel.Meta{Title: "menu.androidui.elementAction", Icon: "action"}},
		{Path: "/androidui/autoCaseStep", ParentName: "androidui", Name: "autoCaseStepAndroid", Hidden: false, Component: "androidui/autoCaseStep/index", Sort: 400, Meta: sysModel.Meta{Title: "menu.androidui.testSteps", Icon: "case-step"}},
		{Path: "/androidui/autoCase", ParentName: "androidui", Name: "autoCaseAndroid", Hidden: false, Component: "androidui/autoCase/index", Sort: 500, Meta: sysModel.Meta{Title: "menu.androidui.testCases", Icon: "testcase"}},
		{Path: "/androidui/timerTask", ParentName: "androidui", Name: "timerTaskAndroid", Hidden: false, Component: "androidui/timerTask/index", Sort: 600, Meta: sysModel.Meta{Title: "menu.androidui.scheduledTasks", Icon: "time-task"}},
		{Path: "/androidui/autoReport", ParentName: "androidui", Name: "autoReportAndroid", Hidden: false, Component: "androidui/autoReport/index", Sort: 700, Meta: sysModel.Meta{Title: "menu.androidui.reports", Icon: "bxs-report"}},
		{Path: "/androidui/autoReport/:id", ParentName: "androidui", Name: "autoReportAndroidDetail", Hidden: true, Component: "androidui/autoReport/detail", Sort: 10000, Meta: sysModel.Meta{Title: "menu.androidui.reportDetail", Icon: ""}},
	}
}
