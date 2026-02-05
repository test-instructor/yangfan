package menu

type MenuParameterSeed struct {
	MenuName  string
	Component string
	Type      string
	Key       string
	Value     string
}

func ParameterSeeds() []MenuParameterSeed {
	return []MenuParameterSeed{
		{MenuName: "AutoStepAndroid", Component: "androidui/autoStep/index", Type: "query", Key: "type", Value: "android"},
		{MenuName: "runConfigAndroid", Component: "androidui/runConfig/index", Type: "query", Key: "type", Value: "android"},
		{MenuName: "autoCaseStepAndroid", Component: "androidui/autoCaseStep/index", Type: "query", Key: "type", Value: "android"},
		{MenuName: "autoCaseAndroid", Component: "androidui/autoCase/index", Type: "query", Key: "type", Value: "android"},
		{MenuName: "timerTaskAndroid", Component: "androidui/timerTask/index", Type: "query", Key: "type", Value: "android"},
		{MenuName: "autoReportAndroid", Component: "androidui/autoReport/index", Type: "query", Key: "type", Value: "android"},
	}
}
