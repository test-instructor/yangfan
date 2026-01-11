package automation

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/middleware"
)

type AutoCaseStepRouter struct{}

// InitAutoCaseStepRouter 初始化 测试步骤 路由信息
func (s *AutoCaseStepRouter) InitAutoCaseStepRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	acsRouter := Router.Group("acs").Use(middleware.OperationRecord())
	acsRouterWithoutRecord := Router.Group("acs")
	acsRouterWithoutAuth := PublicRouter.Group("acs")
	{
		acsRouter.POST("createAutoCaseStep", acsApi.CreateAutoCaseStep)             // 新建测试步骤
		acsRouter.POST("addAutoCaseStepApi", acsApi.AddAutoCaseStepApi)             // 测试步骤添加API
		acsRouter.POST("sortAutoCaseStepApi", acsApi.SortAutoCaseStepApi)           // 测试步骤API排序
		acsRouter.DELETE("deleteAutoCaseStep", acsApi.DeleteAutoCaseStep)           // 删除测试步骤
		acsRouter.DELETE("deleteAutoCaseStepApi", acsApi.DeleteAutoCaseStepApi)     // 删除测试步骤API
		acsRouter.DELETE("deleteAutoCaseStepByIds", acsApi.DeleteAutoCaseStepByIds) // 批量删除测试步骤
		acsRouter.PUT("updateAutoCaseStep", acsApi.UpdateAutoCaseStep)              // 更新测试步骤
	}
	{
		acsRouterWithoutRecord.GET("findAutoCaseStep", acsApi.FindAutoCaseStep)       // 根据ID获取测试步骤
		acsRouterWithoutRecord.GET("findAutoCaseStepApi", acsApi.FindAutoCaseStepApi) // 根据ID获取测试步骤的API
		acsRouterWithoutRecord.GET("getAutoCaseStepList", acsApi.GetAutoCaseStepList) // 获取测试步骤列表
	}
	{
		acsRouterWithoutAuth.GET("getAutoCaseStepPublic", acsApi.GetAutoCaseStepPublic) // 测试步骤开放接口
	}
}
