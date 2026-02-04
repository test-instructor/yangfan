package menu

type MenuParameterSeed struct {
	MenuPath  string
	Component string
	Type      string
	Key       string
	Value     string
}

func ParameterSeeds() []MenuParameterSeed {
	return []MenuParameterSeed{
		{MenuPath: "AutoStepAndroid", Component: "view/automation/autostep/autostep.vue", Type: "query", Key: "type", Value: "android"},
		{MenuPath: "runConfigAndroid", Component: "view/platform/runconfig/runconfig.vue", Type: "query", Key: "type", Value: "android"},
		{MenuPath: "autoCaseStepAndroid", Component: "view/automation/autocasestep/autocasestep.vue", Type: "query", Key: "type", Value: "android"},
		{MenuPath: "autoCaseAndroid", Component: "view/automation/autocase/autocase.vue", Type: "query", Key: "type", Value: "android"},
		{MenuPath: "timerTaskAndroid", Component: "view/automation/timertask/timertask.vue", Type: "query", Key: "type", Value: "android"},
		{MenuPath: "autoReportAndroid", Component: "view/automation/autoreport/autoreport.vue", Type: "query", Key: "type", Value: "android"},

		{MenuPath: "rc", Component: "view/platform/runconfig/runconfig.vue", Type: "query", Key: "type", Value: "api"},
		{MenuPath: "as", Component: "view/automation/autostep/autostep.vue", Type: "query", Key: "type", Value: "api"},
		{MenuPath: "acs", Component: "view/automation/autocasestep/autocasestep.vue", Type: "query", Key: "type", Value: "api"},
		{MenuPath: "ac", Component: "view/automation/autocase/autocase.vue", Type: "query", Key: "type", Value: "api"},
		{MenuPath: "tk", Component: "view/automation/timertask/timertask.vue", Type: "query", Key: "type", Value: "api"},
		{MenuPath: "auto-report-detail/:id", Component: "view/automation/autoreport/AutoReportDetail.vue", Type: "query", Key: "type", Value: "api"},
	}
}
