package platform

import (
	"context"
	"errors"

	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	platformReq "github.com/test-instructor/yangfan/server/v2/model/platform/request"
)

type PythonCodeFuncService struct{}

// CreatePythonCodeFunc 创建python函数详情记录
// Author [yourname](https://github.com/yourname)
func (pcfService *PythonCodeFuncService) CreatePythonCodeFunc(ctx context.Context, pcf *platform.PythonCodeFunc) (err error) {
	err = global.GVA_DB.Create(pcf).Error
	return err
}

// DeletePythonCodeFunc 删除python函数详情记录
// Author [yourname](https://github.com/yourname)
func (pcfService *PythonCodeFuncService) DeletePythonCodeFunc(ctx context.Context, ID string, projectId int64) (err error) {
	var pcf platform.PythonCodeFunc
	err = global.GVA_DB.Where("id = ?", ID).First(&pcf).Error
	if err != nil {
		return err
	}
	if pcf.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}
	err = global.GVA_DB.Delete(&platform.PythonCodeFunc{}, "id = ?", ID).Error
	return err
}

// DeletePythonCodeFuncByIds 批量删除python函数详情记录
// Author [yourname](https://github.com/yourname)
func (pcfService *PythonCodeFuncService) DeletePythonCodeFuncByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]platform.PythonCodeFunc{}, "id in ?", IDs).Error
	return err
}

// UpdatePythonCodeFunc 更新python函数详情记录
// Author [yourname](https://github.com/yourname)
func (pcfService *PythonCodeFuncService) UpdatePythonCodeFunc(ctx context.Context, pcfs platformReq.PythonCodeFuncSearch, projectId int64) (err error) {

	if pcfs.Data == nil {
		return errors.New("函数列表为空")
	}
	err = global.GVA_DB.Model([]platform.PythonCodeFunc{}).Where("project_id = ?", projectId).Delete(&platform.PythonCodeFunc{}).Error
	if err != nil {
		return err
	}
	var data []platform.PythonCodeFunc
	for _, v := range pcfs.Data {
		var oldPythonCodeFunc platform.PythonCodeFunc
		oldPythonCodeFunc.ProjectId = projectId
		oldPythonCodeFunc.Name = v.Name
		oldPythonCodeFunc.Params = v.Params
		oldPythonCodeFunc.FullCode = v.FullCode
		oldPythonCodeFunc.StartIndex = v.StartIndex
		data = append(data, oldPythonCodeFunc)
	}
	err = global.GVA_DB.Model(&platform.PythonCodeFunc{}).Save(&data).Error
	return err
}

// GetPythonCodeFunc 根据ID获取python函数详情记录
// Author [yourname](https://github.com/yourname)
func (pcfService *PythonCodeFuncService) GetPythonCodeFunc(ctx context.Context, ID string) (pcf platform.PythonCodeFunc, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&pcf).Error
	return
}

// GetPythonCodeFuncInfoList 分页获取python函数详情记录
// Author [yourname](https://github.com/yourname)
func (pcfService *PythonCodeFuncService) GetPythonCodeFuncInfoList(ctx context.Context, info platformReq.PythonCodeFuncSearch) (list []platform.PythonCodeFunc, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&platform.PythonCodeFunc{})
	var pcfs []platform.PythonCodeFunc
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

	err = db.Find(&pcfs).Error
	return pcfs, total, err
}
func (pcfService *PythonCodeFuncService) GetPythonCodeFuncPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
