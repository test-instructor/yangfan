package system

import (
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/model/common/request"
	"github.com/test-instructor/cheetah/server/model/interfacecase"
	"github.com/test-instructor/cheetah/server/model/system"
	interfacecaseReq "github.com/test-instructor/cheetah/server/model/system/request"
	"gorm.io/gorm"
)

type ProjectService struct {
}

// CreateProject 创建Project记录

func (projectService *ProjectService) CreateProject(project system.Project) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {

		err = tx.Create(&project).Error
		if err != nil {
			return err
		}
		content := `import logging
import time
from typing import List


# commented out function will be filtered
# def get_headers():
#     return {"User-Agent": "hrp"}


def get_user_agent():
   return "hrp/funppy"


def sleep(n_secs):
   time.sleep(n_secs)


def sum(*args):
   result = 0
   for arg in args:
       result += arg
   return result


def sum_ints(*args: List[int]) -> int:
   result = 0
   for arg in args:
       result += arg
   return result


def sum_two_int(a: int, b: int) -> int:
   return a + b


def sum_two_string(a: str, b: str) -> str:
   return a + b


def sum_strings(*args: List[str]) -> str:
   result = ""
   for arg in args:
       result += arg
   return result


def concatenate(*args: List[str]) -> str:
   result = ""
   for arg in args:
       result += str(arg)
   return result


def setup_hook_example(name):
   logging.warning("setup_hook_example")
   return f"setup_hook_example: {name}"


def teardown_hook_example(name):
   logging.warning("teardown_hook_example")
   return f"teardown_hook_example: {name}"

def return_string():
   return "demo"
`
		var debugtalk *interfacecase.ApiDebugTalk
		debugtalk = new(interfacecase.ApiDebugTalk)
		debugtalk.ProjectID = project.ID
		debugtalk.Content = content
		debugtalk.FileType = interfacecase.FileDebugTalk
		err = global.GVA_DB.Model(interfacecase.ApiDebugTalk{}).Create(debugtalk).Error

		return err
	})
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
