package automation

import (
	"context"
	"errors"
	"time"

	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/automation"
	automationReq "github.com/test-instructor/yangfan/server/v2/model/automation/request"
	projectmgrService "github.com/test-instructor/yangfan/server/v2/service/projectmgr"
)

type AutoReportService struct{}

// CreateAutoReport 创建自动报告记录
// Author [yourname](https://github.com/yourname)
func (arService *AutoReportService) CreateAutoReport(ctx context.Context, ar *automation.AutoReport) (err error) {
	err = global.GVA_DB.Create(ar).Error
	return err
}

// DeleteAutoReport 删除自动报告记录
// Author [yourname](https://github.com/yourname)
func (arService *AutoReportService) DeleteAutoReport(ctx context.Context, ID string, projectId int64) (err error) {
	var ar automation.AutoReport
	err = global.GVA_DB.Where("id = ?", ID).First(&ar).Error
	if err != nil {
		return err
	}
	if ar.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}
	err = global.GVA_DB.Delete(&automation.AutoReport{}, "id = ?", ID).Error
	return err
}

// DeleteAutoReportByIds 批量删除自动报告记录
// Author [yourname](https://github.com/yourname)
func (arService *AutoReportService) DeleteAutoReportByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]automation.AutoReport{}, "id in ?", IDs).Error
	return err
}

// UpdateAutoReport 更新自动报告记录
// Author [yourname](https://github.com/yourname)
func (arService *AutoReportService) UpdateAutoReport(ctx context.Context, ar automation.AutoReport, projectId int64) (err error) {
	var oldAutoReport automation.AutoReport
	err = global.GVA_DB.Model(&oldAutoReport).Where("id = ?", ar.ID).First(&oldAutoReport).Error
	if err != nil {
		return err
	}
	if oldAutoReport.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}

	err = global.GVA_DB.Model(&automation.AutoReport{}).Where("id = ?", ar.ID).Omit("node_name").Updates(&ar).Error
	if err == nil {
		oldStatus := int64(-1)
		if oldAutoReport.Status != nil {
			oldStatus = *oldAutoReport.Status
		}
		newStatus := oldStatus
		if ar.Status != nil {
			newStatus = *ar.Status
		}
		if newStatus != oldStatus && (newStatus == int64(automation.ReportStatusSuccess) || newStatus == int64(automation.ReportStatusFailed)) {
			go func(reportId uint) {
				nctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
				defer cancel()
				_ = (&projectmgrService.ReportNotifyService{}).NotifyAutoReport(nctx, reportId)
			}(ar.ID)
		}
	}
	return err
}

// GetAutoReport 根据ID获取自动报告记录
// Author [yourname](https://github.com/yourname)
func (arService *AutoReportService) GetAutoReport(ctx context.Context, ID string) (ar automation.AutoReport, err error) {
	err = global.GVA_DB.
		Preload("Stat").
		Preload("Stat.Testcases").
		Preload("Stat.Teststeps").
		Preload("Stat.Teststepapi").
		Preload("Time").
		Preload("Progress").
		Preload("Details").
		Preload("Details.Records").
		Where("id = ?", ID).
		First(&ar).Error
	if err == nil {
		// 如果 Progress 为空，尝试从 Redis 加载
		ar.LoadProgressFromRedis()
	}
	return
}

// GetAutoReportInfoList 分页获取自动报告记录
// Author [yourname](https://github.com/yourname)
func (arService *AutoReportService) GetAutoReportInfoList(ctx context.Context, info automationReq.AutoReportSearch) (list []automation.AutoReport, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&automation.AutoReport{}).Preload("Progress")
	var ars []automation.AutoReport
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}
	db.Order("id desc")
	db.Where("project_id = ? ", info.ProjectId)

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&ars).Error
	if err == nil {
		// 如果 Progress 为空，尝试从 Redis 加载
		for i := range ars {
			ars[i].LoadProgressFromRedis()
		}
	}
	return ars, total, err
}
func (arService *AutoReportService) GetAutoReportPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
