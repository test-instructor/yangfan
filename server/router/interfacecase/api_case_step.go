package interfacecase

import (
	"github.com/gin-gonic/gin"

	"github.com/test-instructor/yangfan/server/api/v1"
	"github.com/test-instructor/yangfan/server/middleware"
)

type TestCaseRouter struct {
}

// InitTestCaseRouter 初始化 TestCase 路由信息
func (s *TestCaseRouter) InitTestCaseRouter(Router *gin.RouterGroup) {
	testCaseRouter := Router.Group("step").Use(middleware.OperationRecord())
	testCaseRouterWithoutRecord := Router.Group("step")
	var apiCaseApi = v1.ApiGroupApp.InterfaceCaseApiGroup.ApiCase
	{
		testCaseRouter.POST("createTestCase", apiCaseApi.CreateTestCaseStep)             // 新建TestCase
		testCaseRouter.DELETE("deleteTestCase", apiCaseApi.DeleteTestCaseStep)           // 删除TestCase
		testCaseRouter.DELETE("deleteTestCaseByIds", apiCaseApi.DeleteTestCaseStepByIds) // 批量删除TestCase
		testCaseRouter.PUT("updateTestCase", apiCaseApi.UpdateTestCaseStep)              // 更新TestCase
		testCaseRouter.POST("sortTestCase", apiCaseApi.SortTestCaseStep)                 // 更新TestCase
		testCaseRouter.POST("addTestCase", apiCaseApi.AddTestCaseStep)                   // 更新TestCase
		testCaseRouter.DELETE("delTestCase", apiCaseApi.DelTestCaseStep)                 // 更新TestCase
	}
	{
		testCaseRouterWithoutRecord.GET("findTestCase", apiCaseApi.FindTestCaseStep)       // 根据ID获取TestCase
		testCaseRouterWithoutRecord.GET("getTestCaseList", apiCaseApi.GetTestCaseStepList) // 获取TestCase列表
	}
}
