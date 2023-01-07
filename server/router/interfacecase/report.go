package interfacecase

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/test-instructor/cheetah/server/api/v1"
	"github.com/test-instructor/cheetah/server/middleware"
)

type ReportRouter struct {
}

// InitReportRouter ReportRouter 初始化 TestCase 路由信息
func (s *ReportRouter) InitReportRouter(Router *gin.RouterGroup) {
	reportRouter := Router.Group("").Use(middleware.OperationRecord())

	var reportApi = v1.ApiGroupApp.InterfaceCaseApiGroup.ReportApi
	{
		reportRouter.GET("getReportList", reportApi.GetReportList)
		reportRouter.GET("findReport", reportApi.FindReport)
		reportRouter.DELETE("delReport", reportApi.DelReport)
	}
}
