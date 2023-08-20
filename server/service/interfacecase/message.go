package interfacecase

import (
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	interfacecaseReq "github.com/test-instructor/yangfan/server/model/interfacecase/request"
)

type MessageServer struct{}

func (msgs *MessageServer) CreateMessage(msg interfacecase.ApiMessage) (err error) {
	err = global.GVA_DB.Create(&msg).Error
	return err
}

func (msgs *MessageServer) DeleteMessage(msg interfacecase.ApiMessage) (err error) {
	err = global.GVA_DB.Model(&interfacecase.ApiMessage{}).Delete(&msg).Error
	return err
}

func (msgs *MessageServer) UpdateMessage(msg interfacecase.ApiMessage) (err error) {
	err = global.GVA_DB.Save(&msg).Error
	return err
}

func (msgs *MessageServer) GetMessageList(info interfacecaseReq.ApiMessaceSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&interfacecase.ApiMessage{})
	var messages []interfacecase.ApiMessage
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	db.Scopes(projectDB(info.ProjectID))
	err = db.Limit(limit).Offset(offset).Find(&messages).Error
	return err, messages, total
}

func (msgs *MessageServer) FindMessage(id uint) (err error, messages interfacecase.ApiMessage) {
	err = global.GVA_DB.Where("id = ?", id).First(&messages).Error
	return err, messages
}
