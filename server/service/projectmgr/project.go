package projectmgr

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/projectmgr"
	projectmgrReq "github.com/test-instructor/yangfan/server/v2/model/projectmgr/request"
	"github.com/test-instructor/yangfan/server/v2/utils"
)

type ProjectService struct{}

// CreateProject 创建项目配置记录
// Author [yourname](https://github.com/yourname)
func (pjService *ProjectService) CreateProject(c *gin.Context, ctx context.Context, pj *projectmgr.Project) (err error) {
	pj.Creator = utils.GetUserID(c)
	uuid, err := utils.SnowflakeIDString()
	if err != nil {
		return err
	}
	secretSnowflake, err := utils.SnowflakeIDString()
	if err != nil {
		return err
	}
	secretRand, err := utils.SecureRandomString(24)
	if err != nil {
		return err
	}
	pj.UUID = uuid
	pj.Secret = secretSnowflake + "." + secretRand
	err = global.GVA_DB.Create(pj).Error
	return err
}

// DeleteProject 删除项目配置记录
// Author [yourname](https://github.com/yourname)
func (pjService *ProjectService) DeleteProject(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&projectmgr.Project{}, "id = ?", ID).Error
	return err
}

// DeleteProjectByIds 批量删除项目配置记录
// Author [yourname](https://github.com/yourname)
func (pjService *ProjectService) DeleteProjectByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]projectmgr.Project{}, "id in ?", IDs).Error
	return err
}

// UpdateProject 更新项目配置记录
// Author [yourname](https://github.com/yourname)
func (pjService *ProjectService) UpdateProject(ctx context.Context, pj projectmgr.Project) (err error) {
	updates := map[string]interface{}{
		"name":     pj.Name,
		"admin":    pj.Admin,
		"describe": pj.Describe,
		"logo":     pj.Logo,
	}
	err = global.GVA_DB.Model(&projectmgr.Project{}).Where("id = ?", pj.ID).Updates(updates).Error
	return err
}

func (pjService *ProjectService) ResetProjectAuth(ctx context.Context, id uint) (pj projectmgr.Project, err error) {
	uuid, err := utils.SnowflakeIDString()
	if err != nil {
		return pj, err
	}
	secretSnowflake, err := utils.SnowflakeIDString()
	if err != nil {
		return pj, err
	}
	secretRand, err := utils.SecureRandomString(24)
	if err != nil {
		return pj, err
	}
	updates := map[string]interface{}{
		"uuid":   uuid,
		"secret": secretSnowflake + "." + secretRand,
	}
	if err := global.GVA_DB.Model(&projectmgr.Project{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return pj, err
	}
	pj, err = pjService.GetProject(ctx, strconv.FormatUint(uint64(id), 10))
	return pj, err
}

// GetProject 根据ID获取项目配置记录
// Author [yourname](https://github.com/yourname)
func (pjService *ProjectService) GetProject(ctx context.Context, ID string) (pj projectmgr.Project, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&pj).Error
	return
}

// GetProjectInfoList 分页获取项目配置记录
// Author [yourname](https://github.com/yourname)
func (pjService *ProjectService) GetProjectInfoList(ctx context.Context, info projectmgrReq.ProjectSearch) (list []projectmgr.Project, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&projectmgr.Project{})
	var pjs []projectmgr.Project
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

	err = db.Find(&pjs).Error
	return pjs, total, err
}
func (pjService *ProjectService) GetProjectPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
