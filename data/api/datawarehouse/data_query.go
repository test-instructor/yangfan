package datawarehouse

import (
	"net/http"

	"github.com/gin-gonic/gin"
	svcDw "github.com/test-instructor/yangfan/data/service/datawarehouse"
	"github.com/test-instructor/yangfan/server/v2/global"
	"go.uber.org/zap"
)

type DataQueryApi struct{}

// QueryData 数据查询接口
// @Tags DataWarehouse
// @Summary 数据仓库数据查询
// @Accept application/json
// @Produce application/json
// @Param data body svcDw.DataQueryRequest true "查询参数"
// @Success 200 {object} svcDw.DataQueryResponse "查询成功"
// @Router /api/datawarehouse/query [post]
func (a *DataQueryApi) QueryData(c *gin.Context) {
	var req svcDw.DataQueryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// 验证请求参数
	if err := dataQueryService.ValidateRequest(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// 执行查询
	list, total, err := dataQueryService.QueryData(req)
	if err != nil {
		global.GVA_LOG.Error("查询数据失败", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "查询数据失败"})
		return
	}

	c.JSON(http.StatusOK, svcDw.DataQueryResponse{
		Count: req.Count,
		Total: total,
		List:  list,
	})
}
