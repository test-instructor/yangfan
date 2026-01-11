package automation

import (
	"context"
	"errors"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/automation"
	automationReq "github.com/test-instructor/yangfan/server/v2/model/automation/request"
)

type RequestService struct{}

// CreateRequest 创建请求记录
// Author [yourname](https://github.com/yourname)
func (reqService *RequestService) CreateRequest(ctx context.Context, req *automation.Request) (err error) {
	err = global.GVA_DB.Create(req).Error
	return err
}

// DeleteRequest 删除请求记录
// Author [yourname](https://github.com/yourname)
func (reqService *RequestService) DeleteRequest(ctx context.Context, ID string, projectId int64) (err error) {
	var req automation.Request
	err = global.GVA_DB.Where("id = ?", ID).First(&req).Error
	if err != nil {
		return err
	}
	if req.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}
	err = global.GVA_DB.Delete(&automation.Request{}, "id = ?", ID).Error
	return err
}

// DeleteRequestByIds 批量删除请求记录
// Author [yourname](https://github.com/yourname)
func (reqService *RequestService) DeleteRequestByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]automation.Request{}, "id in ?", IDs).Error
	return err
}

// UpdateRequest 更新请求记录
// Author [yourname](https://github.com/yourname)
func (reqService *RequestService) UpdateRequest(ctx context.Context, req automation.Request, projectId int64) (err error) {
	var oldRequest automation.Request
	err = global.GVA_DB.Model(&oldRequest).Where("id = ?", req.ID).First(&oldRequest).Error
	if err != nil {
		return err
	}
	if oldRequest.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}

	err = global.GVA_DB.Model(&automation.Request{}).Where("id = ?", req.ID).Updates(&req).Error
	return err
}

// GetRequest 根据ID获取请求记录
// Author [yourname](https://github.com/yourname)
func (reqService *RequestService) GetRequest(ctx context.Context, ID string) (req automation.Request, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&req).Error
	return
}

// GetRequestInfoList 分页获取请求记录
// Author [yourname](https://github.com/yourname)
func (reqService *RequestService) GetRequestInfoList(ctx context.Context, info automationReq.RequestSearch) (list []automation.Request, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&automation.Request{})
	var reqs []automation.Request
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&reqs).Error
	return reqs, total, err
}
func (reqService *RequestService) GetRequestPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
