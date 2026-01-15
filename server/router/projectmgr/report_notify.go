package projectmgr

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/middleware"
)

type ReportNotifyRouter struct{}

func (s *ReportNotifyRouter) InitReportNotifyRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	ntRouter := Router.Group("nt").Use(middleware.OperationRecord())
	ntRouterWithoutRecord := Router.Group("nt")
	{
		ntRouter.POST("createReportNotifyChannel", ntApi.CreateReportNotifyChannel)
		ntRouter.DELETE("deleteReportNotifyChannel", ntApi.DeleteReportNotifyChannel)
		ntRouter.DELETE("deleteReportNotifyChannelByIds", ntApi.DeleteReportNotifyChannelByIds)
		ntRouter.PUT("updateReportNotifyChannel", ntApi.UpdateReportNotifyChannel)
	}
	{
		ntRouterWithoutRecord.GET("findReportNotifyChannel", ntApi.FindReportNotifyChannel)
		ntRouterWithoutRecord.GET("getReportNotifyChannelList", ntApi.GetReportNotifyChannelList)
		ntRouterWithoutRecord.GET("getAutoReportNotifyStatus", ntApi.GetAutoReportNotifyStatus)
	}
}
