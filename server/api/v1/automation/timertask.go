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

type TimerTaskApi struct{}

// CreateTimerTask 创建定时任务
// @Tags TimerTask
// @Summary 创建定时任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body automation.TimerTask true "创建定时任务"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /tk/createTimerTask [post]
func (tkApi *TimerTaskApi) CreateTimerTask(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var tk automation.TimerTask
	err := c.ShouldBindJSON(&tk)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = tkService.CreateTimerTask(ctx, &tk)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteTimerTask 删除定时任务
// @Tags TimerTask
// @Summary 删除定时任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body automation.TimerTask true "删除定时任务"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /tk/deleteTimerTask [delete]
func (tkApi *TimerTaskApi) DeleteTimerTask(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	projectId := utils.GetProjectIDInt64(c)
	err := tkService.DeleteTimerTask(ctx, ID, projectId)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteTimerTaskByIds 批量删除定时任务
// @Tags TimerTask
// @Summary 批量删除定时任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /tk/deleteTimerTaskByIds [delete]
func (tkApi *TimerTaskApi) DeleteTimerTaskByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := tkService.DeleteTimerTaskByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateTimerTask 更新定时任务
// @Tags TimerTask
// @Summary 更新定时任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body automation.TimerTask true "更新定时任务"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /tk/updateTimerTask [put]
func (tkApi *TimerTaskApi) UpdateTimerTask(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var tk automation.TimerTask
	err := c.ShouldBindJSON(&tk)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	projectId := utils.GetProjectIDInt64(c)
	err = tkService.UpdateTimerTask(ctx, tk, projectId)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindTimerTask 用id查询定时任务
// @Tags TimerTask
// @Summary 用id查询定时任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询定时任务"
// @Success 200 {object} response.Response{data=automation.TimerTask,msg=string} "查询成功"
// @Router /tk/findTimerTask [get]
func (tkApi *TimerTaskApi) FindTimerTask(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	retk, err := tkService.GetTimerTask(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(retk, c)
}

// GetTimerTaskList 分页获取定时任务列表
// @Tags TimerTask
// @Summary 分页获取定时任务列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query automationReq.TimerTaskSearch true "分页获取定时任务列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /tk/getTimerTaskList [get]
func (tkApi *TimerTaskApi) GetTimerTaskList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo automationReq.TimerTaskSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := tkService.GetTimerTaskInfoList(ctx, pageInfo)
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

// GetTimerTaskPublic 不需要鉴权的定时任务接口
// @Tags TimerTask
// @Summary 不需要鉴权的定时任务接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /tk/getTimerTaskPublic [get]
func (tkApi *TimerTaskApi) GetTimerTaskPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	tkService.GetTimerTaskPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的定时任务接口信息",
	}, "获取成功", c)
}

// AddTimerTaskCase 添加任务引用的测试用例
// @Tags     TimerTask
// @Summary  任务添加测试用例
// @Security ApiKeyAuth
// @Accept   application/json
// @Produce  application/json
// @Param    data body automationReq.TimerTaskCaseReq true "任务ID与用例ID"
// @Success  200  {object} response.Response{msg=string} "添加成功"
// @Router   /tk/addTimerTaskCase [post]
func (tkApi *TimerTaskApi) AddTimerTaskCase(c *gin.Context) {
	var req automationReq.TimerTaskCaseReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := tkService.AddTimerTaskCase(c.Request.Context(), req); err != nil {
		global.GVA_LOG.Error("添加失败!", zap.Error(err))
		response.FailWithMessage("添加失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

// SortTimerTaskCase 任务用例排序
// @Tags     TimerTask
// @Summary  任务用例排序
// @Security ApiKeyAuth
// @Accept   application/json
// @Produce  application/json
// @Param    data body automationReq.TimerTaskCaseSort true "排序数据"
// @Success  200  {object} response.Response{msg=string} "排序成功"
// @Router   /tk/sortTimerTaskCase [post]
func (tkApi *TimerTaskApi) SortTimerTaskCase(c *gin.Context) {
	var req automationReq.TimerTaskCaseSort
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := tkService.SortTimerTaskCase(c.Request.Context(), req); err != nil {
		global.GVA_LOG.Error("排序失败!", zap.Error(err))
		response.FailWithMessage("排序失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("排序成功", c)
}

// DelTimerTaskCase 删除任务引用的用例
// @Tags     TimerTask
// @Summary  删除任务引用的用例
// @Security ApiKeyAuth
// @Accept   application/json
// @Produce  application/json
// @Param    ID query string true "关联ID"
// @Success  200  {object} response.Response{msg=string} "删除成功"
// @Router   /tk/delTimerTaskCase [delete]
func (tkApi *TimerTaskApi) DelTimerTaskCase(c *gin.Context) {
	id := c.Query("ID")
	if err := tkService.DelTimerTaskCase(c.Request.Context(), id); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetTimerTaskCases 获取任务引用的用例列表
// @Tags     TimerTask
// @Summary  获取任务引用的用例列表
// @Security ApiKeyAuth
// @Accept   application/json
// @Produce  application/json
// @Param    ID query string true "任务ID"
// @Success  200  {object} response.Response{data=[]automation.AutoCase,msg=string} "获取成功"
// @Router   /tk/getTimerTaskCases [get]
func (tkApi *TimerTaskApi) GetTimerTaskCases(c *gin.Context) {
	id := c.Query("ID")
	list, err := tkService.GetTimerTaskCases(c.Request.Context(), id)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithData(list, c)
}

// CreateTag 创建标签
// @Tags     TimerTaskTag
// @Summary  创建标签
// @Security ApiKeyAuth
// @Accept   application/json
// @Produce  application/json
// @Param    data body automation.TimerTaskTag true "创建标签"
// @Success  200  {object} response.Response{msg=string} "创建成功"
// @Router   /tk/tag/createTag [post]
func (tkApi *TimerTaskApi) CreateTag(c *gin.Context) {
	ctx := c.Request.Context()
	var tag automation.TimerTaskTag
	err := c.ShouldBindJSON(&tag)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	projectId := utils.GetProjectIDInt64(c)
	if projectId != 0 {
		tag.ProjectId = projectId
	}
	err = tkService.CreateTag(ctx, &tag)
	if err != nil {
		global.GVA_LOG.Error("创建标签失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// UpdateTag 更新标签
// @Tags     TimerTaskTag
// @Summary  更新标签
// @Security ApiKeyAuth
// @Accept   application/json
// @Produce  application/json
// @Param    data body automation.TimerTaskTag true "更新标签"
// @Success  200  {object} response.Response{msg=string} "更新成功"
// @Router   /tk/tag/updateTag [put]
func (tkApi *TimerTaskApi) UpdateTag(c *gin.Context) {
	ctx := c.Request.Context()
	var tag automation.TimerTaskTag
	err := c.ShouldBindJSON(&tag)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	projectId := utils.GetProjectIDInt64(c)
	err = tkService.UpdateTag(ctx, tag, projectId)
	if err != nil {
		global.GVA_LOG.Error("更新标签失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// DeleteTag 删除标签
// @Tags     TimerTaskTag
// @Summary  删除标签
// @Security ApiKeyAuth
// @Accept   application/json
// @Produce  application/json
// @Param    ID query uint true "标签ID"
// @Success  200  {object} response.Response{msg=string} "删除成功"
// @Router   /tk/tag/deleteTag [delete]
func (tkApi *TimerTaskApi) DeleteTag(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Query("ID")
	projectId := utils.GetProjectIDInt64(c)
	err := tkService.DeleteTag(ctx, id, projectId)
	if err != nil {
		global.GVA_LOG.Error("删除标签失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetTagList 分页获取标签列表
// @Tags     TimerTaskTag
// @Summary  分页获取标签列表
// @Security ApiKeyAuth
// @Accept   application/json
// @Produce  application/json
// @Param    data query automationReq.TagSearch true "分页获取标签列表"
// @Success  200  {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router   /tk/tag/getTagList [get]
func (tkApi *TimerTaskApi) GetTagList(c *gin.Context) {
	ctx := c.Request.Context()
	var pageInfo automationReq.TagSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := tkService.GetTagList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取标签列表失败!", zap.Error(err))
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
