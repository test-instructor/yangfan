package datawarehouse

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/middleware"
)

type DataCategoryManagementRouter struct{}

// InitDataCategoryManagementRouter 初始化 数据分类 路由信息
func (s *DataCategoryManagementRouter) InitDataCategoryManagementRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	dcmRouter := Router.Group("dcm").Use(middleware.OperationRecord())
	dcmRouterWithoutRecord := Router.Group("dcm")
	dcmRouterWithoutAuth := PublicRouter.Group("dcm")
	{
		dcmRouter.POST("createDataCategoryManagement", dcmApi.CreateDataCategoryManagement)             // 新建数据分类
		dcmRouter.DELETE("deleteDataCategoryManagement", dcmApi.DeleteDataCategoryManagement)           // 删除数据分类
		dcmRouter.DELETE("deleteDataCategoryManagementByIds", dcmApi.DeleteDataCategoryManagementByIds) // 批量删除数据分类
		dcmRouter.PUT("updateDataCategoryManagement", dcmApi.UpdateDataCategoryManagement)              // 更新数据分类
	}
	{
		dcmRouterWithoutRecord.GET("findDataCategoryManagement", dcmApi.FindDataCategoryManagement)       // 根据ID获取数据分类
		dcmRouterWithoutRecord.GET("getDataCategoryManagementList", dcmApi.GetDataCategoryManagementList) // 获取数据分类列表
		dcmRouterWithoutRecord.GET("getDataCategoryTypeList", dcmApi.GetDataCategoryTypeList)             // 获取数据分类类型列表
	}
	{
		dcmRouterWithoutAuth.GET("getDataCategoryManagementPublic", dcmApi.GetDataCategoryManagementPublic) // 数据分类开放接口
	}
}
