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
	//msgRouterWithoutRecord := Router.Group("")
	var magApi = v1.ApiGroupApp.InterfaceCaseApiGroup.MessageApi
	{
		msgRouter.POST("createMessage", magApi.CreateMessageType) // 新建Message
	}
	{
		//msgRouterWithoutRecord.GET("getMessageList", GetMessageList)
	}
}
