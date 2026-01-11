package projectmgr

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/common/response"
	"github.com/test-instructor/yangfan/server/v2/model/projectmgr"
	projectmgrReq "github.com/test-instructor/yangfan/server/v2/model/projectmgr/request"
	"go.uber.org/zap"
)

type UserProjectAccessApi struct{}

// CreateUserProjectAccess 创建项目成员与权限
// @Tags UserProjectAccess
// @Summary 创建项目成员与权限
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body projectmgr.UserProjectAccess true "创建项目成员与权限"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /upa/createUserProjectAccess [post]
func (upaApi *UserProjectAccessApi) CreateUserProjectAccess(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var upa projectmgr.UserProjectAccess
	err := c.ShouldBindJSON(&upa)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = upaService.CreateUserProjectAccess(ctx, &upa)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteUserProjectAccess 删除项目成员与权限
// @Tags UserProjectAccess
// @Summary 删除项目成员与权限
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body projectmgr.UserProjectAccess true "删除项目成员与权限"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /upa/deleteUserProjectAccess [delete]
func (upaApi *UserProjectAccessApi) DeleteUserProjectAccess(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := upaService.DeleteUserProjectAccess(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteUserProjectAccessByIds 批量删除项目成员与权限
// @Tags UserProjectAccess
// @Summary 批量删除项目成员与权限
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /upa/deleteUserProjectAccessByIds [delete]
func (upaApi *UserProjectAccessApi) DeleteUserProjectAccessByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := upaService.DeleteUserProjectAccessByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateUserProjectAccess 更新项目成员与权限
// @Tags UserProjectAccess
// @Summary 更新项目成员与权限
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body projectmgr.UserProjectAccess true "更新项目成员与权限"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /upa/updateUserProjectAccess [put]
func (upaApi *UserProjectAccessApi) UpdateUserProjectAccess(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var upa projectmgr.UserProjectAccess
	err := c.ShouldBindJSON(&upa)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = upaService.UpdateUserProjectAccess(ctx, upa)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindUserProjectAccess 用id查询项目成员与权限
// @Tags UserProjectAccess
// @Summary 用id查询项目成员与权限
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询项目成员与权限"
// @Success 200 {object} response.Response{data=projectmgr.UserProjectAccess,msg=string} "查询成功"
// @Router /upa/findUserProjectAccess [get]
func (upaApi *UserProjectAccessApi) FindUserProjectAccess(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reupa, err := upaService.GetUserProjectAccess(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reupa, c)
}

// GetUserProjectAccessList 分页获取项目成员与权限列表
// @Tags UserProjectAccess
// @Summary 分页获取项目成员与权限列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query projectmgrReq.UserProjectAccessSearch true "分页获取项目成员与权限列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /upa/getUserProjectAccessList [get]
func (upaApi *UserProjectAccessApi) GetUserProjectAccessList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo projectmgrReq.UserProjectAccessSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := upaService.GetUserProjectAccessInfoList(ctx, pageInfo)
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

// GetUserProjectAccessPublic 不需要鉴权的项目成员与权限接口
// @Tags UserProjectAccess
// @Summary 不需要鉴权的项目成员与权限接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /upa/getUserProjectAccessPublic [get]
func (upaApi *UserProjectAccessApi) GetUserProjectAccessPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	upaService.GetUserProjectAccessPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的项目成员与权限接口信息",
	}, "获取成功", c)
}
