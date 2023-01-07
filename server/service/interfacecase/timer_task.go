package interfacecase

import (
	"github.com/robfig/cron/v3"
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/model/common/request"
	"github.com/test-instructor/cheetah/server/model/interfacecase"
	interfacecaseReq "github.com/test-instructor/cheetah/server/model/interfacecase/request"
	"github.com/test-instructor/cheetah/server/service/interfacecase/runTestCase"
	"gorm.io/gorm"
	"strconv"
)

type TimerTaskService struct {
}

// CreateTimerTask 创建TimerTask记录

func (taskService *TimerTaskService) CreateTimerTask(task interfacecase.ApiTimerTask) (err error) {
	err = global.GVA_DB.Create(&task).Error
	return err
}

// DeleteTimerTask 删除TimerTask记录

func (taskService *TimerTaskService) DeleteTimerTask(task interfacecase.ApiTimerTask) (err error) {
	err = global.GVA_DB.Delete(&task).Error
	return err
}

// DeleteTimerTaskByIds 批量删除TimerTask记录

func (taskService *TimerTaskService) DeleteTimerTaskByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]interfacecase.ApiTimerTask{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTimerTask 更新TimerTask记录

func (taskService *TimerTaskService) UpdateTimerTask(task interfacecase.ApiTimerTask) (err error) {
	var oId interfacecase.Operator
	global.GVA_DB.Model(interfacecase.ApiTimerTask{}).Where("id = ?", task.ID).First(&oId)
	task.CreatedByID = oId.CreatedByID
	task.TestCase = []*interfacecase.ApiCase{}
	err = global.GVA_DB.Where("id = ?", task.ID).Save(&task).Error
	if err != nil {
		return
	}
	global.GVA_Timer.Remove(strconv.Itoa(int(task.ID)), task.EntryID)
	if *task.Status {
		id, err := global.GVA_Timer.AddTaskByFunc(strconv.Itoa(int(task.ID)), task.RunTime, runTestCase.RunTimerTaskBack(task.ID), cron.WithSeconds())
		if err != nil {
			return err
		}
		task.EntryID = int(id)
		err = global.GVA_DB.Save(&task).Error
		if err != nil {
			return err
		}
	}
	return err
}

func (taskService *TimerTaskService) AddTaskCase(taskID uint, caseIDs []uint) (err error) {

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		for _, v := range caseIDs {
			var task interfacecase.ApiTimerTaskRelationship
			task.ApiTimerTaskId = taskID
			task.Sort = 9999
			task.ApiCaseId = v
			err := tx.Model(interfacecase.ApiTimerTaskRelationship{}).Create(&task).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func (taskService *TimerTaskService) DelTaskCase(task interfacecase.ApiTimerTaskRelationship) (err error) {
	err = global.GVA_DB.Delete(&task).Error
	return err
}

func (taskService *TimerTaskService) SortTaskCase(task []interfacecase.ApiTimerTaskRelationship) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		for _, v := range task {
			err := tx.Find(&interfacecase.ApiTimerTaskRelationship{
				GVA_MODEL: global.GVA_MODEL{ID: v.ID},
			}).Update("Sort", v.Sort).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func (taskService *TimerTaskService) FindTaskTestCase(id uint) (err error, apiTimerTaskRelationship []interfacecase.ApiTimerTaskRelationship) {
	//timerTask := interfacecase.ApiTimerTask{GVA_MODEL: global.GVA_MODEL{ID: id}}
	//global.GVA_DB.First(&timerTask)
	global.GVA_DB.Model(interfacecase.ApiTimerTaskRelationship{}).
		Where("api_timer_task_id = ?", id).
		Preload("ApiCase").
		Preload("ApiTimerTask").
		Order("Sort").
		Find(&apiTimerTaskRelationship)

	//global.GVA_DB.Where("timer_task_id = ?", id).
	//	Preload("TestCase").
	//	Order("sort").
	//	Find(&apiTimerTaskRelationship)
	return
}

func (taskService *TimerTaskService) AddTaskTestCase(apiCaseID request.ApiCaseIdReq) (caseApiDetail interfacecase.ApiStep, err error) {
	caseApiDetail = interfacecase.ApiStep{GVA_MODEL: global.GVA_MODEL{ID: apiCaseID.ApiID}}
	err = global.GVA_DB.Preload("Request").First(&caseApiDetail).Error
	if err != nil {
		return interfacecase.ApiStep{}, err
	}
	caseApiDetail.Parent = caseApiDetail.ID
	caseApiDetail.ID = 0
	caseApiDetail.Request.ID = 0
	caseApiDetail.ApiType = 2
	caseDetail := interfacecase.ApiCase{GVA_MODEL: global.GVA_MODEL{ID: apiCaseID.CaseID}}
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var err error
		err = tx.Create(&caseApiDetail).Error
		if err != nil {
			return err
		}
		err = tx.Model(&caseDetail).Association("TStep").Append(&caseApiDetail)
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		return interfacecase.ApiStep{}, err
	}
	return
}

func (taskService *TimerTaskService) SetTaskCase(id uint, caseIds []uint) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		TxErr := tx.Delete(&[]interfacecase.ApiTimerTaskRelationship{}, "api_timer_task_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}
		timerCase := []interfacecase.ApiTimerTaskRelationship{}
		for _, caseID := range caseIds {
			timerCase = append(timerCase, interfacecase.ApiTimerTaskRelationship{
				ApiTimerTaskId: id,
				ApiCaseId:      caseID,
			})
		}
		TxErr = tx.Create(&timerCase).Error
		if TxErr != nil {
			return TxErr
		}
		return nil
	})
}

// GetTimerTask 根据id获取TimerTask记录

func (taskService *TimerTaskService) GetTimerTask(id uint) (err error, task interfacecase.ApiTimerTask) {
	err = global.GVA_DB.Preload("Project").
		Preload("RunConfig").
		Preload("TestCase").
		Where("id = ?", id).
		First(&task).Error
	return
}

// GetTimerTaskInfoList 分页获取TimerTask记录

func (taskService *TimerTaskService) GetTimerTaskInfoList(info interfacecaseReq.TimerTaskSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&interfacecase.ApiTimerTask{}).
		Preload("RunConfig").
		Preload("TestCase")
	db.Preload("Project").Limit(limit).Offset(offset)
	var tasks []interfacecase.ApiTimerTask
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Find(&tasks, projectDB(db, info.ProjectID)).Error
	return err, tasks, total
}
