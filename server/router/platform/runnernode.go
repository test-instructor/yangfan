package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/middleware"
)

type RunnerNodeRouter struct{}

// InitRunnerNodeRouter 初始化 节点 路由信息
func (s *RunnerNodeRouter) InitRunnerNodeRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	rnRouter := Router.Group("rn").Use(middleware.OperationRecord())
	rnRouterWithoutRecord := Router.Group("rn")
	rnRouterWithoutAuth := PublicRouter.Group("rn")
	{
		rnRouter.POST("createRunnerNode", rnApi.CreateRunnerNode)             // 新建节点
		rnRouter.DELETE("deleteRunnerNode", rnApi.DeleteRunnerNode)           // 删除节点
		rnRouter.DELETE("deleteRunnerNodeByIds", rnApi.DeleteRunnerNodeByIds) // 批量删除节点
		rnRouter.PUT("updateRunnerNode", rnApi.UpdateRunnerNode)              // 更新节点
	}
	{
		rnRouterWithoutRecord.GET("findRunnerNode", rnApi.FindRunnerNode)       // 根据ID获取节点
		rnRouterWithoutRecord.GET("getRunnerNodeList", rnApi.GetRunnerNodeList) // 获取节点列表
	}
	{
		rnRouterWithoutAuth.GET("getRunnerNodePublic", rnApi.GetRunnerNodePublic) // 节点开放接口
	}
}
