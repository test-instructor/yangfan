package interfacecase

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/test-instructor/yangfan/server/api/v1"
	"github.com/test-instructor/yangfan/server/middleware"
)

type EnvironmentRouter struct {
}

func (env *EnvironmentRouter) InitEnvironmentRouter(Router *gin.RouterGroup) {
	envRouter := Router.Group("").Use(middleware.OperationRecord())
	envRouterWithoutRecord := Router.Group("")
	var envApi = v1.ApiGroupApp.InterfaceCaseApiGroup.EnvironmentAPi
	{
		envRouter.POST("createEnv", envApi.CreateEnv)
		envRouter.DELETE("deleteEnv", envApi.DeleteEnv)
		envRouter.DELETE("deleteEnvByIds", envApi.DeleteEnvByIds)
		envRouter.PUT("updateEnv", envApi.UpdateEnv)

		envRouter.DELETE("deleteEnvVariable", envApi.DeleteEnvVariable)
		envRouter.POST("createEnvVariable", envApi.CreateEnvVariable)
	}
	{
		// mokc api
		envRouter.POST("createEnvMock", envApi.CreateEnvMock)
		envRouter.DELETE("deleteEnvMock", envApi.DeleteEnvMock)
	}
	{
		envRouterWithoutRecord.GET("findEnv", envApi.FindEnv)
		envRouterWithoutRecord.GET("getEnvList", envApi.GetEnvList)

		envRouterWithoutRecord.GET("findEnvVariable", envApi.FindEnvVariable)
		envRouterWithoutRecord.GET("getEnvVariableList", envApi.GetEnvVariableList)
	}
	{
		// mokc api
		envRouterWithoutRecord.GET("findEnvMock", envApi.FindEnvMock)
		envRouterWithoutRecord.GET("getEnvMockList", envApi.GetEnvMockList)
	}
}
