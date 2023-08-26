package interfacecase

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/test-instructor/yangfan/server/api/v1"
	"github.com/test-instructor/yangfan/server/middleware"
)

type ApiCIRouter struct{}

func (ci ApiCIRouter) InitApiCIRouter(Router *gin.RouterGroup) {
	Router.Group("/ci")
	ciRouter := Router.Group("")
	ciRouter.Use(middleware.CIAuth())
	var ciAPi = v1.ApiGroupApp.InterfaceCaseApiGroup.ApiCIApi
	{
		ciRouter.GET("runTag", ciAPi.RunTag)
		ciRouter.POST("runTag", ciAPi.RunTag)

		ciRouter.GET("getReport", ciAPi.GetReport)
	}
}
