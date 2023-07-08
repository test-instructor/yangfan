package interfacecase

import (
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	interfacecaseReq "github.com/test-instructor/yangfan/server/model/interfacecase/request"
)

type ReportService struct {
}

func (reportService *ReportService) GetReportList(info interfacecaseReq.ReportSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.
		Model(&interfacecase.ApiReport{}).
		Preload("Stat.TestCases").
		Preload("Time").
		Preload("Project").Joins("Project").Where("Project.ID = ?", info.ProjectID)

	var apiReport []interfacecase.ApiReport

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.ApiType > 0 {
		db.Where("case_type = ?", info.ApiType)
	}
	if info.RunType > 0 {
		db.Where("run_type = ?", info.RunType)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("ID desc").Find(&apiReport).Error

	return err, apiReport, total
}

func (reportService *ReportService) GetReportDetail(id uint) (err error, detail interfacecase.ApiReportDetailsRecordsData) {

	// 创建db
	db := global.GVA_DB.
		Model(&interfacecase.ApiReportDetailsRecordsData{}).
		Preload("HttpStat")

	var apiReport interfacecase.ApiReportDetailsRecordsData

	err = db.Find(&apiReport, "id = ?", id).Error

	return err, apiReport
}

func (reportService *ReportService) FindReport(apiReport interfacecase.ApiReport) (err error, list interface{}) {

	// 创建db
	db := global.GVA_DB.
		Model(&interfacecase.ApiReport{}).
		Preload("Time").
		Preload("Stat.TestCases").
		Preload("Stat.TestSteps").
		Preload("Details").
		Preload("Details.Records").
		Preload("Details.Records.Data").
		Preload("Details.Records.Data.HttpStat")

	err = db.Find(&apiReport).Error
	resetReport(&apiReport)
	return err, apiReport
}

func (reportService *ReportService) DelReport(report interfacecase.ApiReport) (err error) {
	err = global.GVA_DB.Delete(&report).Error
	return err
}
