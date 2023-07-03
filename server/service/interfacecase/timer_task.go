package interfacecase

import (
	"github.com/test-instructor/yangfan/proto/tools"
	"github.com/test-instructor/yangfan/server/grpc"
	"gorm.io/gorm"

	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/common/request"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	interfacecaseReq "github.com/test-instructor/yangfan/server/model/interfacecase/request"
)

type TimerTaskService struct {
}

// CreateTimerTask 创建TimerTask记录

func (taskService *TimerTaskService) CreateTimerTask(task interfacecase.ApiTimerTask) (err error) {
	defer func() {
		if err == nil && *task.Status {
			var res = &tools.SetTaskRes{
				ID:          uint64(task.ID),
				TimerStatus: tools.TimerStatusOperate_ADD,
			}
			grpc.ServerToolsObj.SendMessageToTimerTaskClients(res)
		}
	}()
	err = global.GVA_DB.Create(&task).Error
	return err
}

// DeleteTimerTask 删除TimerTask记录

func (taskService *TimerTaskService) DeleteTimerTask(task interfacecase.ApiTimerTask) (err error) {
	defer func() {
		if err == nil {
			var res = &tools.SetTaskRes{
				ID:          uint64(task.ID),
				EntryID:     int64(task.EntryID),
				TimerStatus: tools.TimerStatusOperate_DELETE,
			}
			grpc.ServerToolsObj.SendMessageToTimerTaskClients(res)
		}
	}()
	err = global.GVA_DB.Delete(&task).Error
	return err
}

// DeleteTimerTaskByIds 批量删除TimerTask记录

func (taskService *TimerTaskService) DeleteTimerTaskByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]interfacecase.ApiTimerTask{}, "id in ?", ids.Ids).Error
	return err
}

func (taskService *TimerTaskService) setTaskTag(ids []uint, taskID uint) (err error) {
	if len(ids) == 0 {
		return nil
	}
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Delete(&[]interfacecase.ApiTimerTaskTagRelationship{},
			"api_timer_task_id=?", taskID).Error
		if err != nil {
			return err
		}
		var taskTags []interfacecase.ApiTimerTaskTagRelationship
		for _, v := range ids {
			tags := interfacecase.ApiTimerTaskTagRelationship{
				ApiTimerTaskId:    taskID,
				ApiTimerTaskTagId: v,
			}
			taskTags = append(taskTags, tags)
		}
		err = tx.Create(&taskTags).Error
		return err
	})
}

// UpdateTimerTask 更新TimerTask记录

func (taskService *TimerTaskService) UpdateTimerTask(task interfacecase.ApiTimerTask) (err error) {
	defer func() {
		if err == nil {
			var res = &tools.SetTaskRes{
				ID:          uint64(task.ID),
				TimerStatus: tools.TimerStatusOperate_RESET,
			}
			grpc.ServerToolsObj.SendMessageToTimerTaskClients(res)
		}
	}()

	var oId interfacecase.Operator
	err = taskService.setTaskTag(task.TagIds, task.ID)
	if err != nil {
		return
	}
	global.GVA_DB.Model(interfacecase.ApiTimerTask{}).Where("id = ?", task.ID).First(&oId)
	task.CreatedBy = oId.CreatedBy
	task.TestCase = []*interfacecase.ApiCase{}
	err = global.GVA_DB.Where("id = ?", task.ID).Save(&task).Error
	if err != nil {
		return
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
		var timerCase []interfacecase.ApiTimerTaskRelationship
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
		Preload("ApiTimerTaskTag").
		Where("id = ?", id).
		First(&task).Error
	return
}

// GetTimerTaskInfoList 分页获取TimerTask记录

func (taskService *TimerTaskService) GetTimerTaskInfoList(info interfacecaseReq.TimerTaskSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&interfacecase.ApiTimerTask{}).Preload("RunConfig")
	db2 := global.GVA_DB.Model(&interfacecase.ApiTimerTask{}).Preload("RunConfig").Preload("Project")
	db.Preload("Project").Limit(limit).Offset(offset)

	var tasks []interfacecase.ApiTimerTask
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
		db2 = db2.Where("name LIKE ?", "%"+info.Name+"%")
	}
	db2.Find(nil, projectDB(info.ProjectID))
	err = db2.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Find(&tasks, projectDB(info.ProjectID)).Error
	for i := 0; i < len(tasks); i++ {
		resetRunConfig(&tasks[i].RunConfig)
	}
	return err, tasks, total
}

func (taskService *TimerTaskService) CreateTaskTag(taskTag interfacecase.ApiTimerTaskTag) (taskTags []interfacecase.ApiTimerTaskTag, err error) {
	err = global.GVA_DB.Save(&taskTag).Error
	if err != nil {
		return nil, err
	}
	db := global.GVA_DB.Model(&interfacecase.ApiTimerTaskTag{})
	err = db.Find(&taskTags, projectDB(taskTag.ProjectID)).Error
	return
}

func (taskService *TimerTaskService) GetTimerTaskTagInfoList(info interfacecaseReq.TimerTaskTagSearch) (err error, list interface{}, total int64) {
	// 创建db
	db := global.GVA_DB.Model(&interfacecase.ApiTimerTaskTag{})
	db.Preload("Project")
	var tasks []interfacecase.ApiTimerTaskTag
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Find(&tasks, projectDB(info.ProjectID)).Error
	return err, tasks, total
}

func (taskService *TimerTaskService) DeleteTimerTaskTag(task interfacecase.ApiTimerTaskTag) (err error) {
	err = global.GVA_DB.Delete(&task).Error
	return err
}
