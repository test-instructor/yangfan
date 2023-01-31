package interfacecase

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/cheetah/server/model/common/response"
	"github.com/test-instructor/cheetah/server/model/interfacecase/request"
)

type PyPkg struct {
}

//var pyPkgService = service.ServiceGroupApp.InterfacecaseServiceGroup

func (p *PyPkg) GetPyPkgList(ctx *gin.Context) {
	var pageInfo request.HrpPyPkgRequest
	err := ctx.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	list, total, err := pyPkgService.PyPkgListService(pageInfo)
	//global.GVA_LOG.Info("获取Python包列表成功!", zap.Any("service", service))
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}

func (p *PyPkg) InstallPyPkg(ctx *gin.Context) {
	var pyPkg request.HrpPyPkgRequest
	_ = ctx.ShouldBindJSON(&pyPkg.HrpPyPkg)
	if err := pyPkgService.PyPkgInstallService(pyPkg); err != nil {
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.Ok(ctx)
	}
}

func (p *PyPkg) UninstallPyPkg(ctx *gin.Context) {
	var pyPkg request.HrpPyPkgRequest
	_ = ctx.ShouldBindJSON(&pyPkg)
	if err := pyPkgService.UnInstallService(pyPkg); err != nil {
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.Ok(ctx)
	}
}

func (p *PyPkg) UpdatePyPkg(ctx *gin.Context) {
	var pyPkg request.HrpPyPkgRequest
	_ = ctx.ShouldBindJSON(&pyPkg)
	err := pyPkgService.UpdateService(pyPkg)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.Ok(ctx)
	}
}

func (p *PyPkg) SearchPyPkg(ctx *gin.Context) {
	var pageInfo request.HrpPyPkgRequest
	err := ctx.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	list, total, err := pyPkgService.PyPkgListService(pageInfo)
	//global.GVA_LOG.Info("获取Python包列表成功!", zap.Any("service", service))
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}
