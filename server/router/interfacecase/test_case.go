package interfacecase

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/cheetah/server/api/v1"
	"github.com/test-instructor/cheetah/server/middleware"
)

type TestCaseRouter struct {
}

// InitTestCaseRouter 初始化 TestCase 路由信息
func (s *TestCaseRouter) InitTestCaseRouter(Router *gin.RouterGroup) {
	testCaseRouter := Router.Group("").Use(middleware.OperationRecord())
	testCaseRouterWithoutRecord := Router.Group("")
	var apicaseApi = v1.ApiGroupApp.InterfaceCaseApiGroup.TestCaseApi
	{
		testCaseRouter.POST("createTestCase", apicaseApi.CreateTestCase)             // 新建TestCase
		testCaseRouter.DELETE("deleteTestCase", apicaseApi.DeleteTestCase)           // 删除TestCase
		testCaseRouter.DELETE("deleteTestCaseByIds", apicaseApi.DeleteTestCaseByIds) // 批量删除TestCase
		testCaseRouter.PUT("updateTestCase", apicaseApi.UpdateTestCase)              // 更新TestCase
		testCaseRouter.POST("sortTestCase", apicaseApi.SortTestCase)                 // 更新TestCase
		testCaseRouter.POST("addTestCase", apicaseApi.AddTestCase)                   // 更新TestCase
		testCaseRouter.DELETE("delTestCase", apicaseApi.DelTestCase)                 // 更新TestCase
	}
	{
		testCaseRouterWithoutRecord.GET("findTestCase", apicaseApi.FindTestCase)       // 根据ID获取TestCase
		testCaseRouterWithoutRecord.GET("getTestCaseList", apicaseApi.GetTestCaseList) // 获取TestCase列表
	}
}
