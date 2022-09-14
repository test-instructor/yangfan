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
	var apiCaseApi = v1.ApiGroupApp.InterfaceCaseApiGroup.ApiCase
	var stepTitle = "step/"
	{
		testCaseRouter.POST(stepTitle+"createTestCase", apiCaseApi.CreateTestCaseStep)             // 新建TestCase
		testCaseRouter.DELETE(stepTitle+"deleteTestCase", apiCaseApi.DeleteTestCaseStep)           // 删除TestCase
		testCaseRouter.DELETE(stepTitle+"deleteTestCaseByIds", apiCaseApi.DeleteTestCaseStepByIds) // 批量删除TestCase
		testCaseRouter.PUT(stepTitle+"updateTestCase", apiCaseApi.UpdateTestCaseStep)              // 更新TestCase
		testCaseRouter.POST(stepTitle+"sortTestCase", apiCaseApi.SortTestCaseStep)                 // 更新TestCase
		testCaseRouter.POST(stepTitle+"addTestCase", apiCaseApi.AddTestCaseStep)                   // 更新TestCase
		testCaseRouter.DELETE(stepTitle+"delTestCase", apiCaseApi.DelTestCaseStep)                 // 更新TestCase
	}
	{
		testCaseRouterWithoutRecord.GET(stepTitle+"findTestCase", apiCaseApi.FindTestCaseStep)       // 根据ID获取TestCase
		testCaseRouterWithoutRecord.GET(stepTitle+"getTestCaseList", apiCaseApi.GetTestCaseStepList) // 获取TestCase列表
	}
}
