package projectmgr

import (
	"context"
	"errors"

	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/projectmgr"
	projectmgrReq "github.com/test-instructor/yangfan/server/v2/model/projectmgr/request"
)

type UserProjectAccessService struct{}

// CreateUserProjectAccess 创建项目成员与权限记录
// Author [yourname](https://github.com/yourname)
func (upaService *UserProjectAccessService) CreateUserProjectAccess(ctx context.Context, upa *projectmgr.UserProjectAccess) (err error) {
	var upaOld *projectmgr.UserProjectAccess
	err = global.GVA_DB.Debug().Where("user_id = ? and project_id = ?", upa.UserId, upa.ProjectId).First(&upaOld).Error
	if err == nil && upaOld != nil && upaOld.ID != 0 {
		return errors.New("用户项目访问已存在")
	}
	err = global.GVA_DB.Create(upa).Error
	return err
}

// DeleteUserProjectAccess 删除项目成员与权限记录
// Author [yourname](https://github.com/yourname)
func (upaService *UserProjectAccessService) DeleteUserProjectAccess(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&projectmgr.UserProjectAccess{}, "id = ?", ID).Error
	return err
}

// DeleteUserProjectAccessByIds 批量删除项目成员与权限记录
// Author [yourname](https://github.com/yourname)
func (upaService *UserProjectAccessService) DeleteUserProjectAccessByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]projectmgr.UserProjectAccess{}, "id in ?", IDs).Error
	return err
}

// UpdateUserProjectAccess 更新项目成员与权限记录
// Author [yourname](https://github.com/yourname)
func (upaService *UserProjectAccessService) UpdateUserProjectAccess(ctx context.Context, upa projectmgr.UserProjectAccess) (err error) {
	err = global.GVA_DB.Model(&projectmgr.UserProjectAccess{}).Where("id = ?", upa.ID).Updates(&upa).Error
	return err
}

// GetUserProjectAccess 根据ID获取项目成员与权限记录
// Author [yourname](https://github.com/yourname)
func (upaService *UserProjectAccessService) GetUserProjectAccess(ctx context.Context, ID string) (upa projectmgr.UserProjectAccess, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&upa).Error
	return
}

// GetUserProjectAccessInfoList 分页获取项目成员与权限记录
// Author [yourname](https://github.com/yourname)
func (upaService *UserProjectAccessService) GetUserProjectAccessInfoList(ctx context.Context, info projectmgrReq.UserProjectAccessSearch) (list []projectmgr.UserProjectAccess, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&projectmgr.UserProjectAccess{})
	var upas []projectmgr.UserProjectAccess
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.UserId != nil {
		db = db.Where("user_id = ?", *info.UserId)
	}
	if info.ProjectId != nil {
		db = db.Where("project_id = ?", *info.ProjectId)
	}
	if info.AccessPermission != nil {
		db = db.Where("access_permission = ?", *info.AccessPermission)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&upas).Error
	return upas, total, err
}
func (upaService *UserProjectAccessService) GetUserProjectAccessPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
