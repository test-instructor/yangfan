package interfacecase

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/test-instructor/yangfan/server/api/v1"
	"github.com/test-instructor/yangfan/server/middleware"
)

type ApiCIRouter struct{}

func (ci ApiCIRouter) InitApiCIRouter(Router *gin.RouterGroup) {
	ciRouter := Router.Group("/ci")
	ciRouter.Use(middleware.CIAuth())
	var ciAPi = v1.ApiGroupApp.InterfaceCaseApiGroup.ApiCIApi
	{
		ciRouter.GET("runTag", ciAPi.RunTag)
		ciRouter.POST("runTag", ciAPi.RunTag)

		ciRouter.GET("getReport", ciAPi.GetReport)
		//ciRouter.GET("findReport", ciAPi.FindReport)
		//ciRouter.GET("getReportDetail", ciAPi.GetReportDetail)
	}
}

type ApiCIRespRouter struct{}

func (ci ApiCIRespRouter) InitApiCIRouter(Router *gin.RouterGroup) {
	ciRouter := Router.Group("/ci/resp")
	var ciAPi = v1.ApiGroupApp.InterfaceCaseApiGroup.ApiCIApi
	{
		ciRouter.GET("findReport", ciAPi.FindReport)
		ciRouter.GET("getReportDetail", ciAPi.GetReportDetail)
	}
}
