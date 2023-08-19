package interfacecase

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/common/response"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	"github.com/test-instructor/yangfan/server/service"
	"github.com/test-instructor/yangfan/server/utils"
	"go.uber.org/zap"
)

type MessageApi struct{}

var messageServer = service.ServiceGroupApp.InterfacecaseServiceGroup.MessageServer

func (m *MessageApi) CreateMessageType(c *gin.Context) {
	var msg interfacecase.ApiMessage
	_ = c.ShouldBindJSON(&msg)
	msg.ProjectID = utils.GetUserProject(c)
	msg.CreatedBy = utils.GetUserIDAddress(c)
	if err := messageServer.CreateMessageType(msg); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
