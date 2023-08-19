package interfacecase

import (
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
)

type MessageServer struct{}

func (msgs *MessageServer) CreateMessageType(msg interfacecase.ApiMessage) (err error) {
	err = global.GVA_DB.Create(&msg).Error
	return err
}
