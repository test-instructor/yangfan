package platform

import (
	"context"
	"errors"
	"fmt"

	"github.com/lingcetech/funplugin/myexec"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	platformReq "github.com/test-instructor/yangfan/server/v2/model/platform/request"
	"go.uber.org/zap"
)

type PythonPackageService struct{}

// CreatePythonPackage 创建py 第三方库记录
// Author [yourname](https://github.com/yourname)
func (ppService *PythonPackageService) CreatePythonPackage(ctx context.Context, pp *platform.PythonPackage) (err error) {
	err = global.GVA_DB.Create(pp).Error
	if err == nil {
		go func() {
			global.GVA_LOG.Info("开始安装")
			err := myexec.InstallPythonPackage(global.Python3Executable, fmt.Sprintf("%s==%s", *pp.Name, pp.Version))
			if err != nil {
				global.GVA_LOG.Error("安装失败", zap.Error(err))
			} else {
				global.GVA_LOG.Debug("安装成功")
			}
		}()
	}
	return err
}

// DeletePythonPackage 删除py 第三方库记录
// Author [yourname](https://github.com/yourname)
func (ppService *PythonPackageService) DeletePythonPackage(ctx context.Context, ID string, projectId int64) (err error) {
	var pp platform.PythonPackage
	err = global.GVA_DB.Where("id = ?", ID).First(&pp).Error
	if err != nil {
		return err
	}
	if pp.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}
	err = global.GVA_DB.Delete(&platform.PythonPackage{}, "id = ?", ID).Error
	return err
}

// DeletePythonPackageByIds 批量删除py 第三方库记录
// Author [yourname](https://github.com/yourname)
func (ppService *PythonPackageService) DeletePythonPackageByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]platform.PythonPackage{}, "id in ?", IDs).Error
	return err
}

// UpdatePythonPackage 更新py 第三方库记录
// Author [yourname](https://github.com/yourname)
func (ppService *PythonPackageService) UpdatePythonPackage(ctx context.Context, pp platform.PythonPackage, projectId int64) (err error) {
	err = global.GVA_DB.Model(&platform.PythonPackage{}).Where("id = ?", pp.ID).Updates(&pp).Error
	if err == nil {
		go func() {
			global.GVA_LOG.Info("开始安装")
			err := myexec.InstallPythonPackage(global.Python3Executable, fmt.Sprintf("%s==%s", *pp.Name, pp.Version))
			if err != nil {
				global.GVA_LOG.Error("安装失败", zap.Error(err))
			} else {
				global.GVA_LOG.Debug("安装成功")
			}
		}()
	}
	return err
}

// GetPythonPackage 根据ID获取py 第三方库记录
// Author [yourname](https://github.com/yourname)
func (ppService *PythonPackageService) GetPythonPackage(ctx context.Context, ID string) (pp platform.PythonPackage, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&pp).Error
	return
}

func (ppService *PythonPackageService) GetPythonPackageVersions(ctx context.Context, pkg string) (versions []string, err error) {
	versions, err = myexec.GetPythonPackageVersions(pkg)
	return
}

// GetPythonPackageInfoList 分页获取py 第三方库记录
// Author [yourname](https://github.com/yourname)
func (ppService *PythonPackageService) GetPythonPackageInfoList(ctx context.Context, info platformReq.PythonPackageSearch) (list []platform.PythonPackage, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&platform.PythonPackage{}).Order("id DESC")
	var pps []platform.PythonPackage
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.Name != nil && *info.Name != "" {
		db = db.Where("name LIKE ?", "%"+*info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&pps).Error
	return pps, total, err
}
func (ppService *PythonPackageService) GetPythonPackagePublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
