package datawarehouse

import "github.com/gin-gonic/gin"

type DataQueryRouter struct{}

// InitDataQueryRouter 初始化数据查询路由
func (r *DataQueryRouter) InitDataQueryRouter(Router *gin.RouterGroup) {
	dataQueryGroup := Router.Group("datawarehouse")
	{
		dataQueryGroup.POST("query", dataQueryApi.QueryData)
	}
}
