package platform

import "github.com/gin-gonic/gin"

type OpenRunnerRouter struct{}

func (s *OpenRunnerRouter) InitOpenRunnerRouter(OpenRouter *gin.RouterGroup) {
	if OpenRouter == nil {
		return
	}
	runnerRouter := OpenRouter.Group("runner")
	{
		runnerRouter.GET("run", runnerApi.OpenRun)
		runnerRouter.POST("run", runnerApi.OpenRun)
	}
}
