package interfacecase

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/test-instructor/yangfan/server/api/v1"
	"github.com/test-instructor/yangfan/server/middleware"
)

type RunCaseRouter struct {
}

// InitRunCaseRouter RunCaseRouter 初始化 TestCase 路由信息
func (s *RunCaseRouter) InitRunCaseRouter(Router *gin.RouterGroup) {
	runCaseRouter := Router.Group("").Use(middleware.OperationRecord())

	var runCaseApi = v1.ApiGroupApp.InterfaceCaseApiGroup.RunCaseApi
	{
		runCaseRouter.POST("runTestCaseStep", runCaseApi.RunTestCaseStep)
		runCaseRouter.POST("runApiCase", runCaseApi.RunApiCase)
		runCaseRouter.POST("runBoomerDebug", runCaseApi.RunBoomerDebug)
		runCaseRouter.POST("runBoomer", runCaseApi.RunBoomer)
		runCaseRouter.POST("rebalance", runCaseApi.Rebalance)
		runCaseRouter.GET("stop", runCaseApi.Stop)
		runCaseRouter.POST("runApi", runCaseApi.RunApi)
		runCaseRouter.POST("runTimerTask", runCaseApi.RunTimerTask)
	}
}
