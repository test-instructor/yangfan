package system

import (
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/model/common/request"
	"github.com/test-instructor/cheetah/server/model/system"
	interfacecaseReq "github.com/test-instructor/cheetah/server/model/system/request"
)

type ProjectService struct {
}

// CreateProject 创建Project记录

func (projectService *ProjectService) CreateProject(project system.Project) (err error) {
	err = global.GVA_DB.Create(&project).Error
	return err
}

// DeleteProject 删除Project记录

func (projectService *ProjectService) DeleteProject(project system.Project) (err error) {
	err = global.GVA_DB.Delete(&project).Error
	return err
}

// DeleteProjectByIds 批量删除Project记录

func (projectService *ProjectService) DeleteProjectByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]system.Project{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateProject 更新Project记录

func (projectService *ProjectService) UpdateProject(project system.Project) (err error) {
	err = global.GVA_DB.Save(&project).Error
	return err
}

// GetProject 根据id获取Project记录

func (projectService *ProjectService) GetProject(id uint) (err error, project system.Project) {
	err = global.GVA_DB.Where("id = ?", id).First(&project).Error
	return
}

// GetProjectInfoList 分页获取Project记录

func (projectService *ProjectService) GetProjectInfoList(info interfacecaseReq.ProjectSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&system.Project{})
	var projects []system.Project
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&projects).Error
	return err, projects, total
}
