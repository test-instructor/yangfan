package automation

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/middleware"
)

type AutoCaseRouter struct{}

// InitAutoCaseRouter 初始化 测试用例 路由信息
func (s *AutoCaseRouter) InitAutoCaseRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	acRouter := Router.Group("ac").Use(middleware.OperationRecord())
	acRouterWithoutRecord := Router.Group("ac")
	acRouterWithoutAuth := PublicRouter.Group("ac")
	{
		acRouter.POST("createAutoCase", acApi.CreateAutoCase)             // 新建测试用例
		acRouter.DELETE("deleteAutoCase", acApi.DeleteAutoCase)           // 删除测试用例
		acRouter.DELETE("deleteAutoCaseByIds", acApi.DeleteAutoCaseByIds) // 批量删除测试用例
		acRouter.PUT("updateAutoCase", acApi.UpdateAutoCase)              // 更新测试用例
		acRouter.POST("addAutoCaseStep", acApi.AddAutoCaseStep)           // 添加测试步骤
		acRouter.POST("sortAutoCaseStep", acApi.SortAutoCaseStep)         // 测试步骤排序
		acRouter.DELETE("delAutoCaseStep", acApi.DelAutoCaseStep)         // 删除测试步骤
		acRouter.PUT("setStepConfig", acApi.SetStepConfig)                // 设置步骤配置
	}
	{
		acRouterWithoutRecord.GET("findAutoCase", acApi.FindAutoCase)         // 根据ID获取测试用例
		acRouterWithoutRecord.GET("getAutoCaseList", acApi.GetAutoCaseList)   // 获取测试用例列表
		acRouterWithoutRecord.GET("getAutoCaseSteps", acApi.GetAutoCaseSteps) // 获取测试用例步骤
	}
	{
		acRouterWithoutAuth.GET("getAutoCasePublic", acApi.GetAutoCasePublic) // 测试用例开放接口
	}
}
