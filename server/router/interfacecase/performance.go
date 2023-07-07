package interfacecase

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/test-instructor/yangfan/server/api/v1"
	"github.com/test-instructor/yangfan/server/middleware"
)

type PerformanceRouter struct {
}

// InitPerformanceRouter 初始化 TimerTask 路由信息
func (s *PerformanceRouter) InitPerformanceRouter(Router *gin.RouterGroup) {
	performanceRouter := Router.Group("").Use(middleware.OperationRecord())
	performanceRouterWithoutRecord := Router.Group("")
	var performanceApi = v1.ApiGroupApp.InterfaceCaseApiGroup.PerformanceApi
	//fmt.Println(performanceRouter, performanceRouterWithoutRecord, performanceRouterWithoutRecord)
	{
		performanceRouter.POST("createPerformance", performanceApi.CreatePerformance)   // 新建Performance
		performanceRouter.DELETE("deletePerformance", performanceApi.DeletePerformance) // 删除Performance
		//performanceRouter.DELETE("deleteTimerTaskByIds", taskApi.DeleteTimerTaskByIds) // 批量删除TimerTask
		performanceRouter.PUT("updatePerformance", performanceApi.UpdatePerformance) // 更新Performance
		performanceRouter.POST("sortPerformanceCase", performanceApi.SortPerformanceCase)
		performanceRouter.POST("addPerformanceCase", performanceApi.AddPerformanceCase)
		performanceRouter.POST("addOperation", performanceApi.AddOperation)
		performanceRouter.DELETE("delPerformanceCase", performanceApi.DelPerformanceCase)
		performanceRouter.GET("findPerformance", performanceApi.FindPerformance)
		performanceRouter.GET("findPerformanceCase", performanceApi.FindPerformanceCase)
		performanceRouter.GET("findPerformanceStep", performanceApi.FindPerformanceStep)
		performanceRouter.GET("getReportList", performanceApi.GetReportList)
		performanceRouter.DELETE("deleteReport", performanceApi.DeleteReport)
		//performanceRouter.POST("addTaskTestCase", taskApi.AddTaskTestCase)
		//performanceRouter.POST("setTaskCase", taskApi.SetTaskCase)
	}
	{
		//performanceRouterWithoutRecord.GET("findTimerTask", taskApi.FindTimerTask)       // 根据ID获取TimerTask
		//performanceRouterWithoutRecord.GET("getTimerTaskList", taskApi.GetTimerTaskList) // 获取TimerTask列表
		performanceRouterWithoutRecord.GET("getPerformanceList", performanceApi.GetPerformanceList) // 新建Performance
		performanceRouterWithoutRecord.GET("findReport", performanceApi.FindReport)
	}
}
