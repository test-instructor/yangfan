package automation

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/middleware"
)

type TimerTaskRouter struct{}

// InitTimerTaskRouter 初始化 定时任务 路由信息
func (s *TimerTaskRouter) InitTimerTaskRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	tkRouter := Router.Group("tk").Use(middleware.OperationRecord())
	tkRouterWithoutRecord := Router.Group("tk")
	tkRouterWithoutAuth := PublicRouter.Group("tk")
	{
		tkRouter.POST("createTimerTask", tkApi.CreateTimerTask)             // 新建定时任务
		tkRouter.DELETE("deleteTimerTask", tkApi.DeleteTimerTask)           // 删除定时任务
		tkRouter.DELETE("deleteTimerTaskByIds", tkApi.DeleteTimerTaskByIds) // 批量删除定时任务
		tkRouter.PUT("updateTimerTask", tkApi.UpdateTimerTask)              // 更新定时任务
		// 任务-用例写操作
		tkRouter.POST("addTimerTaskCase", tkApi.AddTimerTaskCase)   // 新增任务引用用例
		tkRouter.POST("sortTimerTaskCase", tkApi.SortTimerTaskCase) // 任务用例排序
		tkRouter.DELETE("delTimerTaskCase", tkApi.DelTimerTaskCase) // 删除任务引用用例
		// Tag 写操作
		tkRouter.POST("tag/createTag", tkApi.CreateTag)   // 新建标签
		tkRouter.PUT("tag/updateTag", tkApi.UpdateTag)    // 更新标签
		tkRouter.DELETE("tag/deleteTag", tkApi.DeleteTag) // 删除标签
	}
	{
		tkRouterWithoutRecord.GET("findTimerTask", tkApi.FindTimerTask)       // 根据ID获取定时任务
		tkRouterWithoutRecord.GET("getTimerTaskList", tkApi.GetTimerTaskList) // 获取定时任务列表
		// 任务-用例读操作
		tkRouterWithoutRecord.GET("getTimerTaskCases", tkApi.GetTimerTaskCases) // 获取任务引用的用例列表
		// Tag 读操作
		tkRouterWithoutRecord.GET("tag/getTagList", tkApi.GetTagList) // 获取标签列表
	}
	{
		tkRouterWithoutAuth.GET("getTimerTaskPublic", tkApi.GetTimerTaskPublic) // 定时任务开放接口
	}
}
