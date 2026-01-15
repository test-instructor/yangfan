package projectmgr

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/common/response"
	"github.com/test-instructor/yangfan/server/v2/model/projectmgr"
	projectmgrReq "github.com/test-instructor/yangfan/server/v2/model/projectmgr/request"
	"github.com/test-instructor/yangfan/server/v2/utils"
	"go.uber.org/zap"
)

type ReportNotifyApi struct{}

func (a *ReportNotifyApi) CreateReportNotifyChannel(c *gin.Context) {
	ctx := c.Request.Context()
	projectId := utils.GetProjectIDInt64(c)
	var channel projectmgr.ProjectReportNotifyChannel
	if err := c.ShouldBindJSON(&channel); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ntService.CreateReportNotifyChannel(ctx, &channel, projectId); err != nil {
		global.GVA_LOG.Error("CreateReportNotifyChannel failed", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (a *ReportNotifyApi) DeleteReportNotifyChannel(c *gin.Context) {
	ctx := c.Request.Context()
	projectId := utils.GetProjectIDInt64(c)
	idStr := c.Query("ID")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	if err := ntService.DeleteReportNotifyChannel(ctx, uint(id), projectId); err != nil {
		global.GVA_LOG.Error("DeleteReportNotifyChannel failed", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (a *ReportNotifyApi) DeleteReportNotifyChannelByIds(c *gin.Context) {
	ctx := c.Request.Context()
	projectId := utils.GetProjectIDInt64(c)
	idStrs := c.QueryArray("IDs[]")
	ids := make([]uint, 0, len(idStrs))
	for _, s := range idStrs {
		v, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			continue
		}
		ids = append(ids, uint(v))
	}
	if err := ntService.DeleteReportNotifyChannelByIds(ctx, ids, projectId); err != nil {
		global.GVA_LOG.Error("DeleteReportNotifyChannelByIds failed", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

func (a *ReportNotifyApi) UpdateReportNotifyChannel(c *gin.Context) {
	ctx := c.Request.Context()
	projectId := utils.GetProjectIDInt64(c)
	var channel projectmgr.ProjectReportNotifyChannel
	if err := c.ShouldBindJSON(&channel); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ntService.UpdateReportNotifyChannel(ctx, channel, projectId); err != nil {
		global.GVA_LOG.Error("UpdateReportNotifyChannel failed", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (a *ReportNotifyApi) FindReportNotifyChannel(c *gin.Context) {
	ctx := c.Request.Context()
	projectId := utils.GetProjectIDInt64(c)
	idStr := c.Query("ID")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	channel, err := ntService.GetReportNotifyChannel(ctx, uint(id), projectId)
	if err != nil {
		global.GVA_LOG.Error("FindReportNotifyChannel failed", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(channel, c)
}

func (a *ReportNotifyApi) GetReportNotifyChannelList(c *gin.Context) {
	ctx := c.Request.Context()
	projectId := utils.GetProjectIDInt64(c)
	var pageInfo projectmgrReq.ReportNotifyChannelSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.ProjectId = projectId
	list, total, err := ntService.GetReportNotifyChannelList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("GetReportNotifyChannelList failed", zap.Error(err))
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

func (a *ReportNotifyApi) GetAutoReportNotifyStatus(c *gin.Context) {
	ctx := c.Request.Context()
	var q projectmgrReq.AutoReportNotifyStatusQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := ntService.GetAutoReportNotifyStatus(ctx, q.ReportId)
	if err != nil {
		global.GVA_LOG.Error("GetAutoReportNotifyStatus failed", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithData(res, c)
}
