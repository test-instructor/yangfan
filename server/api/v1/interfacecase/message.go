package interfacecase

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/common/response"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	interfacecaseReq "github.com/test-instructor/yangfan/server/model/interfacecase/request"
	"github.com/test-instructor/yangfan/server/service"
	"github.com/test-instructor/yangfan/server/utils"
	"go.uber.org/zap"
)

type MessageApi struct{}

var messageServer = service.ServiceGroupApp.InterfacecaseServiceGroup.MessageServer

func (m *MessageApi) CreateMessage(c *gin.Context) {
	var msg interfacecase.ApiMessage
	_ = c.ShouldBindJSON(&msg)
	msg.ProjectID = utils.GetUserProject(c)
	msg.CreatedBy = utils.GetUserIDAddress(c)
	if err := messageServer.CreateMessage(msg); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (m *MessageApi) DeleteMessage(c *gin.Context) {
	var msg interfacecase.ApiMessage
	_ = c.ShouldBindJSON(&msg)
	msg.ProjectID = utils.GetUserProject(c)
	msg.DeleteBy = utils.GetUserIDAddress(c)
	if err := messageServer.DeleteMessage(msg); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (m *MessageApi) UpdateMessage(c *gin.Context) {
	var msg interfacecase.ApiMessage
	_ = c.ShouldBindJSON(&msg)
	msg.ProjectID = utils.GetUserProject(c)
	msg.UpdateBy = utils.GetUserIDAddress(c)
	if err := messageServer.UpdateMessage(msg); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (m *MessageApi) GetMessageList(c *gin.Context) {
	var pageInfo interfacecaseReq.ApiMessaceSearch
	_ = c.ShouldBindQuery(&pageInfo)
	pageInfo.ProjectID = utils.GetUserProject(c)
	if err, list, total := messageServer.GetMessageList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (m *MessageApi) FindMessage(c *gin.Context) {
	var msg interfacecase.ApiMessage
	_ = c.ShouldBindQuery(&msg)
	if err, message := messageServer.FindMessage(msg.ID); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(message, "获取成功", c)
	}

}
