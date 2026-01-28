package initialize

import (
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/automation"
	"github.com/test-instructor/yangfan/server/v2/model/datawarehouse"
	"github.com/test-instructor/yangfan/server/v2/model/example"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	"github.com/test-instructor/yangfan/server/v2/model/projectmgr"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(projectmgr.UserProjectAccess{}, projectmgr.Project{}, projectmgr.ProjectReportNotifyChannel{}, projectmgr.ProjectReportNotifyLog{}, platform.Env{}, platform.EnvDetail{}, platform.PythonCode{}, platform.PythonCodeDebug{}, platform.PythonPackage{}, platform.PythonCodeFunc{}, platform.RunConfig{}, platform.CategoryMenu{}, automation.AutoStep{}, automation.Request{}, automation.AutoCaseStepRelation{}, automation.AutoCaseStep{}, automation.AutoCaseStepList{}, automation.AutoCase{}, automation.TimerTask{}, automation.TimerTaskTag{}, automation.TimerTaskCaseList{}, platform.RunnerNode{}, automation.AutoReport{}, automation.AutoReportStat{}, automation.AutoReportStatTestcases{}, automation.AutoReportStatTeststeps{}, automation.AutoReportTime{}, automation.AutoReportDetail{}, automation.AutoReportRecord{}, automation.AutoReportProgress{}, datawarehouse.DataCategoryManagement{}, datawarehouse.DataCategoryData{}, example.ExaFile{}, example.ExaFileChunk{}, example.ExaAttachmentCategory{}, platform.LLMModelConfig{}, platform.AndroidDeviceOptions{}, platform.IOSDeviceOptions{}, platform.HarmonyDeviceOptions{}, // 自动迁移模型
		platform.BrowserDeviceOptions{})
	if err != nil {
		return err
	}
	if err := migrateDCMCountColumns(db); err != nil {
		return err
	}
	return nil
}
