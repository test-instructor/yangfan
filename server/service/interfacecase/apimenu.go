package interfacecase

import (
	"errors"
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/model/common/request"
	"github.com/test-instructor/cheetah/server/model/interfacecase"
	interfacecaseReq "github.com/test-instructor/cheetah/server/model/interfacecase/request"
)

/*
递归获取树形菜单
*/
type TreeList struct {
	Id       uint        `json:"id"`
	Key      uint        `json:"key"`
	Label    string      `json:"label"`
	Title    string      `json:"title"`
	Parent   uint        `json:"pid"`
	Children []*TreeList `json:"children"`
}

func (apicaseService *ApiMenuService) GetMenu(pid uint, menuType string, project uint) ([]*TreeList, error) {
	var menu []interfacecase.ApiMenu

	db := global.GVA_DB.Model(&interfacecase.ApiMenu{})

	db.Where("Parent = ?", pid).Where("menu_type = ?", menuType).Find(&menu, projectDB(db, project))
	treeList := []*TreeList{}
	for _, v := range menu {
		child, _ := apicaseService.GetMenu(v.ID, v.MenuType, project)
		node := &TreeList{
			Id:     v.ID,
			Key:    v.ID,
			Label:  v.Name,
			Title:  v.Name,
			Parent: v.Parent,
		}
		node.Children = child
		treeList = append(treeList, node)
	}
	return treeList, nil
}

type ApiMenuService struct {
}

// CreateApiMenu 创建ApiMenu记录

func (apicaseService *ApiMenuService) CreateApiMenu(apicase interfacecase.ApiMenu) (err error) {
	err = global.GVA_DB.Create(&apicase).Error
	return err
}

// DeleteApiMenu 删除ApiMenu记录

func (apicaseService *ApiMenuService) DeleteApiMenu(apicase interfacecase.ApiMenu) (err error) {
	err = global.GVA_DB.Delete(&apicase).Error
	return err
}

// DeleteApiMenuByIds 批量删除ApiMenu记录

func (apicaseService *ApiMenuService) DeleteApiMenuByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]interfacecase.ApiMenu{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateApiMenu 更新ApiMenu记录

func (apicaseService *ApiMenuService) UpdateApiMenu(apicase interfacecase.ApiMenu) (err error) {
	var apicaseTemp interfacecase.ApiMenu
	var oId getOperationId
	global.GVA_DB.Model(interfacecase.ApiMenu{}).Where("id = ?", apicase.ID).First(&oId)
	apicaseTemp.CreatedByID = oId.CreatedByID
	apicaseTemp.UpdateByID = apicase.UpdateByID

	global.GVA_DB.Preload("Project").Where("id = ?", apicase.ID).First(&apicaseTemp)
	apicaseTemp.Name = apicase.Name
	err = global.GVA_DB.Save(&apicaseTemp).Error
	return err
}

// GetApiMenu 根据id获取ApiMenu记录

func (apicaseService *ApiMenuService) GetApiMenu(id uint) (err error, apicase interfacecase.ApiMenu) {
	err = global.GVA_DB.Preload("Project").Where("id = ?", id).First(&apicase).Error
	return
}

// GetApiMenuInfoList 分页获取ApiMenu记录

func (apicaseService *ApiMenuService) GetApiMenuInfoList(info interfacecaseReq.ApiMenuSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&interfacecase.ApiMenu{})
	var apicases []interfacecase.ApiMenu
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Preload("Project").Limit(limit).Offset(offset).Find(&apicases, projectDB(db, info.ProjectID)).Error
	return err, apicases, total
}

func (apicaseService *ApiMenuService) GetApiMenuInterface(apicase interfacecase.ApiMenu) (err error) {

	// 创建db
	db := global.GVA_DB.Model(&interfacecase.ApiStep{})
	var interfaceTemplate interfacecase.ApiStep
	// 如果有条件搜索 下方会自动创建搜索语句
	db.Where("Api_Menu_ID = ?", apicase.ID).Where("Project_ID = ?", apicase.Project.ID).First(&interfaceTemplate)
	if interfaceTemplate.ID > 0 {
		err = errors.New("以存在数据，无法删除")
	} else {
		err = nil
	}

	return err
}
