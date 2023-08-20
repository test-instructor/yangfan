package interfacecase

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/test-instructor/yangfan/server/api/v1"
	"github.com/test-instructor/yangfan/server/middleware"
)

type MessageRouter struct {
}

func (msg *MessageRouter) InitMessageRouter(Router *gin.RouterGroup) {
	msgRouter := Router.Group("").Use(middleware.OperationRecord())
	msgRouterWithoutRecord := Router.Group("")
	var msgApi = v1.ApiGroupApp.InterfaceCaseApiGroup.MessageApi
	{
		msgRouter.POST("createMessage", msgApi.CreateMessage)   // 新建Message
		msgRouter.DELETE("deleteMessage", msgApi.DeleteMessage) // 删除Message
		msgRouter.PUT("updateMessage", msgApi.UpdateMessage)    // 更新Message
	}
	{
		msgRouterWithoutRecord.GET("getMessageList", msgApi.GetMessageList) // 获取Message列表
		msgRouterWithoutRecord.GET("findMessage", msgApi.FindMessage)       // 根据ID获取Message
	}
}
