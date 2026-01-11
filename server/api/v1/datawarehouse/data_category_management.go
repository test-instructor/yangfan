package datawarehouse

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/common/response"
	datawarehouseReq "github.com/test-instructor/yangfan/server/v2/model/datawarehouse/request"
	"github.com/test-instructor/yangfan/server/v2/utils"
	"go.uber.org/zap"
)

type DataCategoryManagementApi struct{}

// CreateDataCategoryManagement 创建数据分类
// @Tags DataCategoryManagement
// @Summary 创建数据分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body datawarehouseReq.DataCategoryManagementSave true "创建数据分类"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /dcm/createDataCategoryManagement [post]
func (dcmApi *DataCategoryManagementApi) CreateDataCategoryManagement(c *gin.Context) {
	ctx := c.Request.Context()

	var req datawarehouseReq.DataCategoryManagementSave
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 从上下文获取 projectId
	req.ProjectId = utils.GetProjectIDInt64(c)

	if err := dcmService.CreateDataCategoryManagement(ctx, req); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteDataCategoryManagement 删除数据分类
// @Tags DataCategoryManagement
// @Summary 删除数据分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body datawarehouse.DataCategoryManagement true "删除数据分类"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /dcm/deleteDataCategoryManagement [delete]
func (dcmApi *DataCategoryManagementApi) DeleteDataCategoryManagement(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	projectId := utils.GetProjectIDInt64(c)
	err := dcmService.DeleteDataCategoryManagement(ctx, ID, projectId)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteDataCategoryManagementByIds 批量删除数据分类
// @Tags DataCategoryManagement
// @Summary 批量删除数据分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /dcm/deleteDataCategoryManagementByIds [delete]
func (dcmApi *DataCategoryManagementApi) DeleteDataCategoryManagementByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := dcmService.DeleteDataCategoryManagementByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateDataCategoryManagement 更新数据分类
// @Tags DataCategoryManagement
// @Summary 更新数据分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body datawarehouseReq.DataCategoryManagementSave true "更新数据分类"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /dcm/updateDataCategoryManagement [put]
func (dcmApi *DataCategoryManagementApi) UpdateDataCategoryManagement(c *gin.Context) {
	ctx := c.Request.Context()

	var req datawarehouseReq.DataCategoryManagementSave
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	projectId := utils.GetProjectIDInt64(c)
	userId := utils.GetUserID(c)
	if err := dcmService.UpdateDataCategoryManagement(ctx, req, projectId, userId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindDataCategoryManagement 用id查询数据分类
// @Tags DataCategoryManagement
// @Summary 用id查询数据分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询数据分类"
// @Success 200 {object} response.Response{data=datawarehouse.DataCategoryManagement,msg=string} "查询成功"
// @Router /dcm/findDataCategoryManagement [get]
func (dcmApi *DataCategoryManagementApi) FindDataCategoryManagement(c *gin.Context) {
	ctx := c.Request.Context()

	ID := c.Query("ID")
	projectId := utils.GetProjectIDInt64(c)
	redcm, err := dcmService.GetDataCategoryManagement(ctx, ID, projectId)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(redcm, c)
}

// GetDataCategoryManagementList 分页获取数据分类列表
// @Tags DataCategoryManagement
// @Summary 分页获取数据分类列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query datawarehouseReq.DataCategoryManagementSearch true "分页获取数据分类列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /dcm/getDataCategoryManagementList [get]
func (dcmApi *DataCategoryManagementApi) GetDataCategoryManagementList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo datawarehouseReq.DataCategoryManagementSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := dcmService.GetDataCategoryManagementInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetDataCategoryTypeList 获取数据分类类型列表
// @Tags DataCategoryManagement
// @Summary 获取数据分类类型列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /dcm/getDataCategoryTypeList [get]
func (dcmApi *DataCategoryManagementApi) GetDataCategoryTypeList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	projectId := utils.GetProjectIDInt64(c)

	list, err := dcmService.GetDataCategoryTypeList(ctx, projectId)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithData(list, c)
}

// GetDataCategoryManagementPublic 不需要鉴权的数据分类接口
// @Tags DataCategoryManagement
// @Summary 不需要鉴权的数据分类接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /dcm/getDataCategoryManagementPublic [get]
func (dcmApi *DataCategoryManagementApi) GetDataCategoryManagementPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	dcmService.GetDataCategoryManagementPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的数据分类接口信息",
	}, "获取成功", c)
}
