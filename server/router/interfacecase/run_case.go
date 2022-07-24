package interfacecase

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/test-instructor/cheetah/server/api/v1"
	"github.com/test-instructor/cheetah/server/middleware"
)

type RunCaseRouter struct {
}

// InitRunCaseRouter RunCaseRouter 初始化 TestCase 路由信息
func (s *RunCaseRouter) InitRunCaseRouter(Router *gin.RouterGroup) {
	runCaseRouter := Router.Group("").Use(middleware.OperationRecord())

	var runCaseApi = v1.ApiGroupApp.InterfaceCaseApiGroup.RunCaseApi
	{
		runCaseRouter.POST("runTestCase", runCaseApi.RunTestCase)
		runCaseRouter.POST("runTimerTask", runCaseApi.RunTimerTask)
	}
}
