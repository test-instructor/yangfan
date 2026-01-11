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

type AutoCaseApi struct{}

// CreateAutoCase 创建测试用例
// @Tags AutoCase
// @Summary 创建测试用例
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body automation.AutoCase true "创建测试用例"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /ac/createAutoCase [post]
func (acApi *AutoCaseApi) CreateAutoCase(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var ac automation.AutoCase
	err := c.ShouldBindJSON(&ac)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = acService.CreateAutoCase(ctx, &ac)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteAutoCase 删除测试用例
// @Tags AutoCase
// @Summary 删除测试用例
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body automation.AutoCase true "删除测试用例"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /ac/deleteAutoCase [delete]
func (acApi *AutoCaseApi) DeleteAutoCase(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	projectId := utils.GetProjectIDInt64(c)
	err := acService.DeleteAutoCase(ctx, ID, projectId)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAutoCaseByIds 批量删除测试用例
// @Tags AutoCase
// @Summary 批量删除测试用例
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /ac/deleteAutoCaseByIds [delete]
func (acApi *AutoCaseApi) DeleteAutoCaseByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := acService.DeleteAutoCaseByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAutoCase 更新测试用例
// @Tags AutoCase
// @Summary 更新测试用例
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body automation.AutoCase true "更新测试用例"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /ac/updateAutoCase [put]
func (acApi *AutoCaseApi) UpdateAutoCase(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var ac automation.AutoCase
	err := c.ShouldBindJSON(&ac)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	projectId := utils.GetProjectIDInt64(c)
	err = acService.UpdateAutoCase(ctx, ac, projectId)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAutoCase 用id查询测试用例
// @Tags AutoCase
// @Summary 用id查询测试用例
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询测试用例"
// @Success 200 {object} response.Response{data=automation.AutoCase,msg=string} "查询成功"
// @Router /ac/findAutoCase [get]
func (acApi *AutoCaseApi) FindAutoCase(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reac, err := acService.GetAutoCase(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reac, c)
}

// GetAutoCaseList 分页获取测试用例列表
// @Tags AutoCase
// @Summary 分页获取测试用例列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query automationReq.AutoCaseSearch true "分页获取测试用例列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /ac/getAutoCaseList [get]
func (acApi *AutoCaseApi) GetAutoCaseList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo automationReq.AutoCaseSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := acService.GetAutoCaseInfoList(ctx, pageInfo)
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

// GetAutoCasePublic 不需要鉴权的测试用例接口
// @Tags AutoCase
// @Summary 不需要鉴权的测试用例接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ac/getAutoCasePublic [get]
func (acApi *AutoCaseApi) GetAutoCasePublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	acService.GetAutoCasePublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的测试用例接口信息",
	}, "获取成功", c)
}

func (acApi *AutoCaseApi) AddAutoCaseStep(c *gin.Context) {
	var req automationReq.AutoCaseStepReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = acService.AddAutoCaseStep(c.Request.Context(), req)
	if err != nil {
		global.GVA_LOG.Error("添加失败!", zap.Error(err))
		response.FailWithMessage("添加失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

func (acApi *AutoCaseApi) SortAutoCaseStep(c *gin.Context) {
	var req automationReq.AutoCaseStepSort
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = acService.SortAutoCaseStep(c.Request.Context(), req)
	if err != nil {
		global.GVA_LOG.Error("排序失败!", zap.Error(err))
		response.FailWithMessage("排序失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("排序成功", c)
}

func (acApi *AutoCaseApi) DelAutoCaseStep(c *gin.Context) {
	ID := c.Query("ID")
	err := acService.DelAutoCaseStep(c.Request.Context(), ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (acApi *AutoCaseApi) GetAutoCaseSteps(c *gin.Context) {
	ID := c.Query("ID")
	list, err := acService.GetAutoCaseSteps(c.Request.Context(), ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithData(list, c)
}

func (acApi *AutoCaseApi) SetStepConfig(c *gin.Context) {
	var req automationReq.SetStepConfigReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = acService.SetStepConfig(c.Request.Context(), req)
	if err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("设置成功", c)
}
