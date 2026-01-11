package api

import (
	sysModel "github.com/test-instructor/yangfan/server/v2/model/system"
)

func automationApis() []sysModel.SysApi {
	return []sysModel.SysApi{
		{Path: "/as/createAutoStep", Description: "新增自动化步骤", ApiGroup: "接口管理骤", Method: "POST"},
		{Path: "/as/deleteAutoStep", Description: "删除自动化步骤", ApiGroup: "接口管理骤", Method: "DELETE"},
		{Path: "/as/deleteAutoStepByIds", Description: "批量删除自动化步骤", ApiGroup: "接口管理骤", Method: "DELETE"},
		{Path: "/as/updateAutoStep", Description: "更新自动化步骤", ApiGroup: "接口管理骤", Method: "PUT"},
		{Path: "/as/findAutoStep", Description: "根据ID获取自动化步骤", ApiGroup: "接口管理骤", Method: "GET"},
		{Path: "/as/getAutoStepList", Description: "获取自动化步骤列表", ApiGroup: "接口管理骤", Method: "GET"},

		{Path: "/acs/createAutoCaseStep", Description: "新增测试步骤", ApiGroup: "测试步骤", Method: "POST"},
		{Path: "/acs/deleteAutoCaseStep", Description: "删除测试步骤", ApiGroup: "测试步骤", Method: "DELETE"},
		{Path: "/acs/deleteAutoCaseStepByIds", Description: "批量删除测试步骤", ApiGroup: "测试步骤", Method: "DELETE"},
		{Path: "/acs/updateAutoCaseStep", Description: "更新测试步骤", ApiGroup: "测试步骤", Method: "PUT"},
		{Path: "/acs/findAutoCaseStep", Description: "根据ID获取测试步骤", ApiGroup: "测试步骤", Method: "GET"},
		{Path: "/acs/getAutoCaseStepList", Description: "获取测试步骤列表", ApiGroup: "测试步骤", Method: "GET"},
		{Path: "/acs/sortAutoCaseStepApi", Description: "测试步骤API排序", ApiGroup: "测试步骤", Method: "POST"},
		{Path: "/acs/addAutoCaseStepApi", Description: "测试步骤添加API", ApiGroup: "测试步骤", Method: "POST"},
		{Path: "/acs/findAutoCaseStepApi", Description: "测试步骤查询API信息", ApiGroup: "测试步骤", Method: "GET"},
		{Path: "/acs/deleteAutoCaseStepApi", Description: "删除测试步骤API", ApiGroup: "测试步骤", Method: "DELETE"},

		{Path: "/ac/createAutoCase", Description: "新增测试用例", ApiGroup: "测试用例", Method: "POST"},
		{Path: "/ac/deleteAutoCase", Description: "删除测试用例", ApiGroup: "测试用例", Method: "DELETE"},
		{Path: "/ac/deleteAutoCaseByIds", Description: "批量删除测试用例", ApiGroup: "测试用例", Method: "DELETE"},
		{Path: "/ac/updateAutoCase", Description: "更新测试用例", ApiGroup: "测试用例", Method: "PUT"},
		{Path: "/ac/findAutoCase", Description: "根据ID获取测试用例", ApiGroup: "测试用例", Method: "GET"},
		{Path: "/ac/getAutoCaseList", Description: "获取测试用例列表", ApiGroup: "测试用例", Method: "GET"},
		{Path: "/ac/sortAutoCaseStep", Description: "测试步骤排序", ApiGroup: "测试用例", Method: "POST"},
		{Path: "/ac/addAutoCaseStep", Description: "添加测试步骤", ApiGroup: "测试用例", Method: "POST"},
		{Path: "/ac/delAutoCaseStep", Description: "删除测试步骤", ApiGroup: "测试用例", Method: "DELETE"},
		{Path: "/ac/getAutoCaseSteps", Description: "获取测试用例步骤", ApiGroup: "测试用例", Method: "GET"},
		{Path: "/ac/setStepConfig", Description: "设置步骤配置", ApiGroup: "测试用例", Method: "PUT"},

		{Path: "/tk/createTimerTask", Description: "新增定时任务", ApiGroup: "定时任务", Method: "POST"},
		{Path: "/tk/deleteTimerTask", Description: "删除定时任务", ApiGroup: "定时任务", Method: "DELETE"},
		{Path: "/tk/deleteTimerTaskByIds", Description: "批量删除定时任务", ApiGroup: "定时任务", Method: "DELETE"},
		{Path: "/tk/updateTimerTask", Description: "更新定时任务", ApiGroup: "定时任务", Method: "PUT"},
		{Path: "/tk/findTimerTask", Description: "根据ID获取定时任务", ApiGroup: "定时任务", Method: "GET"},
		{Path: "/tk/getTimerTaskList", Description: "获取定时任务列表", ApiGroup: "定时任务", Method: "GET"},
		{Path: "/tk/tag/createTag", Description: "新建标签", ApiGroup: "定时任务", Method: "POST"},
		{Path: "/tk/tag/updateTag", Description: "更新标签", ApiGroup: "定时任务", Method: "PUT"},
		{Path: "/tk/tag/deleteTag", Description: "删除标签", ApiGroup: "定时任务", Method: "DELETE"},
		{Path: "/tk/tag/getTagList", Description: "获取标签列表", ApiGroup: "定时任务", Method: "GET"},
		{Path: "/tk/addTimerTaskCase", Description: "新增任务引用用例", ApiGroup: "定时任务", Method: "POST"},
		{Path: "/tk/sortTimerTaskCase", Description: "任务用例排序", ApiGroup: "定时任务", Method: "POST"},
		{Path: "/tk/delTimerTaskCase", Description: "删除任务引用用例", ApiGroup: "定时任务", Method: "DELETE"},
		{Path: "/tk/getTimerTaskCases", Description: "获取任务引用的用例列表", ApiGroup: "定时任务", Method: "GET"},

		{Path: "/ar/createAutoReport", Description: "新增自动报告", ApiGroup: "自动报告", Method: "POST"},
		{Path: "/ar/deleteAutoReport", Description: "删除自动报告", ApiGroup: "自动报告", Method: "DELETE"},
		{Path: "/ar/deleteAutoReportByIds", Description: "批量删除自动报告", ApiGroup: "自动报告", Method: "DELETE"},
		{Path: "/ar/updateAutoReport", Description: "更新自动报告", ApiGroup: "自动报告", Method: "PUT"},
		{Path: "/ar/findAutoReport", Description: "根据ID获取自动报告", ApiGroup: "自动报告", Method: "GET"},
		{Path: "/ar/getAutoReportList", Description: "获取自动报告列表", ApiGroup: "自动报告", Method: "GET"},
	}
}
