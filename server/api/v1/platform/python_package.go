package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/common/response"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	platformReq "github.com/test-instructor/yangfan/server/v2/model/platform/request"
	"github.com/test-instructor/yangfan/server/v2/utils"
	"go.uber.org/zap"
)

type PythonPackageApi struct{}

// CreatePythonPackage 创建py 第三方库
// @Tags PythonPackage
// @Summary 创建py 第三方库
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.PythonPackage true "创建py 第三方库"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /pp/createPythonPackage [post]
func (ppApi *PythonPackageApi) CreatePythonPackage(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pp platform.PythonPackage
	err := c.ShouldBindJSON(&pp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = ppService.CreatePythonPackage(ctx, &pp)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeletePythonPackage 删除py 第三方库
// @Tags PythonPackage
// @Summary 删除py 第三方库
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.PythonPackage true "删除py 第三方库"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /pp/deletePythonPackage [delete]
func (ppApi *PythonPackageApi) DeletePythonPackage(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	projectId := utils.GetProjectIDInt64(c)
	err := ppService.DeletePythonPackage(ctx, ID, projectId)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeletePythonPackageByIds 批量删除py 第三方库
// @Tags PythonPackage
// @Summary 批量删除py 第三方库
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /pp/deletePythonPackageByIds [delete]
func (ppApi *PythonPackageApi) DeletePythonPackageByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := ppService.DeletePythonPackageByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdatePythonPackage 更新py 第三方库
// @Tags PythonPackage
// @Summary 更新py 第三方库
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body platform.PythonPackage true "更新py 第三方库"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /pp/updatePythonPackage [put]
func (ppApi *PythonPackageApi) UpdatePythonPackage(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var pp platform.PythonPackage
	err := c.ShouldBindJSON(&pp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	projectId := utils.GetProjectIDInt64(c)
	err = ppService.UpdatePythonPackage(ctx, pp, projectId)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindPythonPackage 用id查询py 第三方库
// @Tags PythonPackage
// @Summary 用id查询py 第三方库
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询py 第三方库"
// @Success 200 {object} response.Response{data=platform.PythonPackage,msg=string} "查询成功"
// @Router /pp/findPythonPackage [get]
func (ppApi *PythonPackageApi) FindPythonPackage(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	repp, err := ppService.GetPythonPackage(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(repp, c)
}

func (ppApi *PythonPackageApi) FindPythonPackageVersion(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	pkg := c.Query("pkg")
	repp, err := ppService.GetPythonPackageVersions(ctx, pkg)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(repp, c)
}

// GetPythonPackageList 分页获取py 第三方库列表
// @Tags PythonPackage
// @Summary 分页获取py 第三方库列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query platformReq.PythonPackageSearch true "分页获取py 第三方库列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /pp/getPythonPackageList [get]
func (ppApi *PythonPackageApi) GetPythonPackageList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo platformReq.PythonPackageSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := ppService.GetPythonPackageInfoList(ctx, pageInfo)
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

// GetPythonPackagePublic 不需要鉴权的py 第三方库接口
// @Tags PythonPackage
// @Summary 不需要鉴权的py 第三方库接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /pp/getPythonPackagePublic [get]
func (ppApi *PythonPackageApi) GetPythonPackagePublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	ppService.GetPythonPackagePublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的py 第三方库接口信息",
	}, "获取成功", c)
}
