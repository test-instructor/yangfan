package automation

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/middleware"
)

type RequestRouter struct{}

// InitRequestRouter 初始化 请求 路由信息
func (s *RequestRouter) InitRequestRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	reqRouter := Router.Group("req").Use(middleware.OperationRecord())
	reqRouterWithoutRecord := Router.Group("req")
	reqRouterWithoutAuth := PublicRouter.Group("req")
	{
		reqRouter.POST("createRequest", reqApi.CreateRequest)             // 新建请求
		reqRouter.DELETE("deleteRequest", reqApi.DeleteRequest)           // 删除请求
		reqRouter.DELETE("deleteRequestByIds", reqApi.DeleteRequestByIds) // 批量删除请求
		reqRouter.PUT("updateRequest", reqApi.UpdateRequest)              // 更新请求
	}
	{
		reqRouterWithoutRecord.GET("findRequest", reqApi.FindRequest)       // 根据ID获取请求
		reqRouterWithoutRecord.GET("getRequestList", reqApi.GetRequestList) // 获取请求列表
	}
	{
		reqRouterWithoutAuth.GET("getRequestPublic", reqApi.GetRequestPublic) // 请求开放接口
	}
}
