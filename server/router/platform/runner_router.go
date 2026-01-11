package platform

import (
	"github.com/gin-gonic/gin"
)

type RunnerRouter struct{}

func (r *RunnerRouter) InitRunnerRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	runnerRouter := Router.Group("runner")
	{
		runnerRouter.POST("api", runnerApi.RunTask)
	}
}
