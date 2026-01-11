package platform

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/test-instructor/yangfan/httprunner/hrp"
	"github.com/test-instructor/yangfan/httprunner/python"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	platformReq "github.com/test-instructor/yangfan/server/v2/model/platform/request"
	"gorm.io/datatypes"
)

type PythonCodeDebugService struct{}

// CreatePythonCodeDebug 创建代码调试记录记录
// Author [yourname](https://github.com/yourname)
func (pcdService *PythonCodeDebugService) CreatePythonCodeDebug(ctx context.Context, pcd *platform.PythonCodeDebug) (data interface{}, err error) {
	var pc platform.PythonCode
	db := global.GVA_DB.Model(&platform.PythonCode{})
	db.Select("id, created_at, updated_at, deleted_at, type, project_id, update_by, code")
	db.Where("project_id = ? and type = ?", pcd.ProjectId, pcd.Type)
	db.Order("id DESC")
	if err = db.First(&pc).Error; err != nil {
		return nil, err
	}

	timestamp := time.Now().Format("20060102150405")
	secondDir := fmt.Sprintf("%d_%d_%s", pc.ID, pc.Type, timestamp)
	// 先构建相对路径
	relativePath := filepath.Join("./debugcode", secondDir)
	// 转换为绝对路径
	path, err := filepath.Abs(relativePath)
	defer os.RemoveAll(path)
	if err != nil {
		return nil, fmt.Errorf("获取绝对路径失败: %w", err)
	}

	// 创建目录（使用绝对路径）
	if err = os.MkdirAll(path, 0755); err != nil {
		return nil, fmt.Errorf("创建目录失败: %w", err)
	}

	// 构建文件绝对路径
	filePath := filepath.Join(path, "debugtalk.py")
	if err = os.WriteFile(filePath, []byte(pc.Code), 0644); err != nil {
		return nil, fmt.Errorf("写入文件失败: %w", err)
	}

	// 后续使用绝对路径初始化插件
	plugin, err := python.InitPlugin(path, "", true)
	if err != nil {
		return nil, err
	}
	defer plugin.Quit()
	var parser = hrp.NewParser()
	parser.Plugin = plugin
	var res interface{}
	res, err = parser.Parse(pcd.Function, pcd.Parameters)
	if err != nil {
		return nil, err
	}
	pcd.ReturnValue = datatypes.JSONMap{
		"code": res,
	}

	if err = global.GVA_DB.Create(pcd).Error; err != nil {
		return nil, fmt.Errorf("数据库保存失败: %w", err)
	}

	return res, nil
}

// DeletePythonCodeDebug 删除代码调试记录记录
// Author [yourname](https://github.com/yourname)
func (pcdService *PythonCodeDebugService) DeletePythonCodeDebug(ctx context.Context, ID string, projectId int64) (err error) {
	var pcd platform.PythonCodeDebug
	err = global.GVA_DB.Where("id = ?", ID).First(&pcd).Error
	if err != nil {
		return err
	}
	if pcd.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}
	err = global.GVA_DB.Delete(&platform.PythonCodeDebug{}, "id = ?", ID).Error
	return err
}

// DeletePythonCodeDebugByIds 批量删除代码调试记录记录
// Author [yourname](https://github.com/yourname)
func (pcdService *PythonCodeDebugService) DeletePythonCodeDebugByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]platform.PythonCodeDebug{}, "id in ?", IDs).Error
	return err
}

// UpdatePythonCodeDebug 更新代码调试记录记录
// Author [yourname](https://github.com/yourname)
func (pcdService *PythonCodeDebugService) UpdatePythonCodeDebug(ctx context.Context, pcd platform.PythonCodeDebug, projectId int64) (err error) {
	var oldPythonCodeDebug platform.PythonCodeDebug
	err = global.GVA_DB.Model(&oldPythonCodeDebug).Where("id = ?", pcd.ID).First(&oldPythonCodeDebug).Error
	if err != nil {
		return err
	}
	if oldPythonCodeDebug.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}

	err = global.GVA_DB.Model(&platform.PythonCodeDebug{}).Where("id = ?", pcd.ID).Updates(&pcd).Error
	return err
}

// GetPythonCodeDebug 根据ID获取代码调试记录记录
// Author [yourname](https://github.com/yourname)
func (pcdService *PythonCodeDebugService) GetPythonCodeDebug(ctx context.Context, ID string) (pcd platform.PythonCodeDebug, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&pcd).Error
	return
}

// GetPythonCodeDebugInfoList 分页获取代码调试记录记录
// Author [yourname](https://github.com/yourname)
func (pcdService *PythonCodeDebugService) GetPythonCodeDebugInfoList(ctx context.Context, info platformReq.PythonCodeDebugSearch) (list []platform.PythonCodeDebug, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&platform.PythonCodeDebug{}).Order("id DESC")
	var pcds []platform.PythonCodeDebug
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.ProjectId != 0 {
		db = db.Where("project_id = ?", info.ProjectId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&pcds).Error
	return pcds, total, err
}
func (pcdService *PythonCodeDebugService) GetPythonCodeDebugPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
