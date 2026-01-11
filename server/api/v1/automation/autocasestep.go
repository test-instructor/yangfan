package automation

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/automation"
	automationReq "github.com/test-instructor/yangfan/server/v2/model/automation/request"
	"github.com/test-instructor/yangfan/server/v2/model/common/response"
	"github.com/test-instructor/yangfan/server/v2/utils"
	"go.uber.org/zap"
)

type AutoCaseStepApi struct{}

// CreateAutoCaseStep 创建测试步骤
// @Tags AutoCaseStep
// @Summary 创建测试步骤
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body automation.AutoCaseStep true "创建测试步骤"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /acs/createAutoCaseStep [post]
func (acsApi *AutoCaseStepApi) CreateAutoCaseStep(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var acs automation.AutoCaseStep
	err := c.ShouldBindJSON(&acs)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = acsService.CreateAutoCaseStep(ctx, &acs)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (acsApi *AutoCaseStepApi) AddAutoCaseStepApi(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var acs automationReq.AutoCaseStepSearchApi
	err := c.ShouldBindJSON(&acs)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var data interface{}
	data, err = acsService.AddAutoCaseStepApi(ctx, &acs)
	if err != nil {
		global.GVA_LOG.Error("添加接口失败!", zap.Error(err))
		response.FailWithMessage("添加接口失败:"+err.Error(), c)
		return
	}
	response.OkWithData(data, c)
}

func (acsApi *AutoCaseStepApi) SortAutoCaseStepApi(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var acs automationReq.AutoCaseStepSearchApi
	err := c.ShouldBindJSON(&acs)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = acsService.SortAutoCaseStepApi(ctx, &acs)
	if err != nil {
		global.GVA_LOG.Error("添加接口失败!", zap.Error(err))
		response.FailWithMessage("添加接口失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("排序成功", c)
}

// DeleteAutoCaseStep 删除测试步骤
// @Tags AutoCaseStep
// @Summary 删除测试步骤
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body automation.AutoCaseStep true "删除测试步骤"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /acs/deleteAutoCaseStep [delete]
func (acsApi *AutoCaseStepApi) DeleteAutoCaseStep(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	projectId := utils.GetProjectIDInt64(c)
	err := acsService.DeleteAutoCaseStep(ctx, ID, projectId)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (acsApi *AutoCaseStepApi) DeleteAutoCaseStepApi(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	projectId := utils.GetProjectIDInt64(c)
	err := acsService.DeleteAutoCaseStepApi(ctx, ID, projectId)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAutoCaseStepByIds 批量删除测试步骤
// @Tags AutoCaseStep
// @Summary 批量删除测试步骤
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /acs/deleteAutoCaseStepByIds [delete]
func (acsApi *AutoCaseStepApi) DeleteAutoCaseStepByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := acsService.DeleteAutoCaseStepByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAutoCaseStep 更新测试步骤
// @Tags AutoCaseStep
// @Summary 更新测试步骤
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body automation.AutoCaseStep true "更新测试步骤"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /acs/updateAutoCaseStep [put]
func (acsApi *AutoCaseStepApi) UpdateAutoCaseStep(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var acs automation.AutoCaseStep
	err := c.ShouldBindJSON(&acs)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	projectId := utils.GetProjectIDInt64(c)
	err = acsService.UpdateAutoCaseStep(ctx, acs, projectId)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAutoCaseStep 用id查询测试步骤
// @Tags AutoCaseStep
// @Summary 用id查询测试步骤
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询测试步骤"
// @Success 200 {object} response.Response{data=automation.AutoCaseStep,msg=string} "查询成功"
// @Router /acs/findAutoCaseStep [get]
func (acsApi *AutoCaseStepApi) FindAutoCaseStep(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reacs, err := acsService.GetAutoCaseStep(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reacs, c)
}

func (acsApi *AutoCaseStepApi) FindAutoCaseStepApi(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reacs, err := acsService.GetAutoCaseStepApi(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reacs, c)
}

// GetAutoCaseStepList 分页获取测试步骤列表
// @Tags AutoCaseStep
// @Summary 分页获取测试步骤列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query automationReq.AutoCaseStepSearch true "分页获取测试步骤列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /acs/getAutoCaseStepList [get]
func (acsApi *AutoCaseStepApi) GetAutoCaseStepList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo automationReq.AutoCaseStepSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := acsService.GetAutoCaseStepInfoList(ctx, pageInfo)
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

// GetAutoCaseStepPublic 不需要鉴权的测试步骤接口
// @Tags AutoCaseStep
// @Summary 不需要鉴权的测试步骤接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /acs/getAutoCaseStepPublic [get]
func (acsApi *AutoCaseStepApi) GetAutoCaseStepPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	acsService.GetAutoCaseStepPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的测试步骤接口信息",
	}, "获取成功", c)
}
