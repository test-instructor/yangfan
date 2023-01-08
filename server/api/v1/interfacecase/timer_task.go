package interfacecase

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/model/common/request"
	"github.com/test-instructor/cheetah/server/model/common/response"
	"github.com/test-instructor/cheetah/server/model/interfacecase"
	interfacecaseReq "github.com/test-instructor/cheetah/server/model/interfacecase/request"
	"github.com/test-instructor/cheetah/server/service"
	"github.com/test-instructor/cheetah/server/utils"
	"go.uber.org/zap"
)

type TimerTaskApi struct {
}

var taskService = service.ServiceGroupApp.InterfacecaseServiceGroup.TimerTaskService

// CreateTimerTask 创建TimerTask
// @Tags TimerTask
// @Summary 创建TimerTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body interfacecase.TimerTask true "创建TimerTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /task/createTimerTask [post]
func (taskApi *TimerTaskApi) CreateTimerTask(c *gin.Context) {
	var task interfacecase.ApiTimerTask
	_ = c.ShouldBindJSON(&task)
	task.ProjectID = utils.GetUserProject(c)
	task.CreatedByID = utils.GetUserIDAddress(c)
	if err := taskService.CreateTimerTask(task); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTimerTask 删除TimerTask
// @Tags TimerTask
// @Summary 删除TimerTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body interfacecase.TimerTask true "删除TimerTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /task/deleteTimerTask [delete]
func (taskApi *TimerTaskApi) DeleteTimerTask(c *gin.Context) {
	var task interfacecase.ApiTimerTask
	_ = c.ShouldBindJSON(&task)
	task.ProjectID = utils.GetUserProject(c)
	task.DeleteByID = utils.GetUserIDAddress(c)
	if err := taskService.DeleteTimerTask(task); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTimerTaskByIds 批量删除TimerTask
// @Tags TimerTask
// @Summary 批量删除TimerTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TimerTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /task/deleteTimerTaskByIds [delete]
func (taskApi *TimerTaskApi) DeleteTimerTaskByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := taskService.DeleteTimerTaskByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTimerTask 更新TimerTask
// @Tags TimerTask
// @Summary 更新TimerTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body interfacecase.TimerTask true "更新TimerTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /task/updateTimerTask [put]
func (taskApi *TimerTaskApi) UpdateTimerTask(c *gin.Context) {
	var task interfacecase.ApiTimerTask
	_ = c.ShouldBindJSON(&task)
	task.ProjectID = utils.GetUserProject(c)
	task.UpdateByID = utils.GetUserIDAddress(c)
	if err := taskService.UpdateTimerTask(task); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (taskApi *TimerTaskApi) SortTaskCase(c *gin.Context) {
	var task []interfacecase.ApiTimerTaskRelationship
	_ = c.ShouldBindJSON(&task)
	if err := taskService.SortTaskCase(task); err != nil {
		global.GVA_LOG.Error("排序失败!", zap.Error(err))
		response.FailWithMessage("排序失败", c)
	} else {
		response.OkWithMessage("排序成功", c)
	}
}

type addTaskCaseReq struct {
	TaskID uint   `json:"task_id"`
	CaseID []uint `json:"case_id"`
}

func (taskApi *TimerTaskApi) AddTaskCase(c *gin.Context) {
	var task addTaskCaseReq
	_ = c.ShouldBindJSON(&task)
	if err := taskService.AddTaskCase(task.TaskID, task.CaseID); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (taskApi *TimerTaskApi) DelTaskCase(c *gin.Context) {
	var task interfacecase.ApiTimerTaskRelationship
	_ = c.ShouldBindJSON(&task)
	if err := taskService.DelTaskCase(task); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

type taskTcaseResp struct {
	Name     string     `json:"name"`
	TestCase []testCase `json:"test_case"`
}

type testCase struct {
	ID   uint                  `json:"id"`
	Case interfacecase.ApiCase `json:"case"`
}

func (taskApi *TimerTaskApi) FindTaskTestCase(c *gin.Context) {
	var task interfacecase.ApiTimerTask
	_ = c.ShouldBindQuery(&task)
	global.GVA_DB.First(&task)
	task.ProjectID = utils.GetUserProject(c)
	var reapicase taskTcaseResp
	err, resp := taskService.FindTaskTestCase(task.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		reapicase.Name = task.Name
		if len(resp) > 0 {
			for _, v := range resp {
				var testcase testCase
				testcase.ID = v.ID
				testcase.Case = v.ApiCase
				reapicase.TestCase = append(reapicase.TestCase, testcase)
			}
			response.OkWithData(gin.H{"reapicase": reapicase}, c)
		} else {
			response.OkWithData(gin.H{"reapicase": reapicase}, c)
		}
	}
}

func (taskApi *TimerTaskApi) AddTaskTestCase(c *gin.Context) {
	var apiCaseID request.ApiCaseIdReq
	_ = c.ShouldBindJSON(&apiCaseID)
	caseApiDetail, err := taskService.AddTaskTestCase(apiCaseID)
	if err != nil {
		global.GVA_LOG.Error("添加用例失败!", zap.Error(err))
		response.FailWithMessage("添加用例失败", c)
	} else {
		response.OkWithDetailed(caseApiDetail, "添加用例成功", c)
	}
}

func (taskApi *TimerTaskApi) SetTaskCase(c *gin.Context) {
	var sua interfacecaseReq.SetTimerCares
	_ = c.ShouldBindJSON(&sua)
	if err := taskService.SetTaskCase(sua.ID, sua.CaseIds); err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// FindTimerTask 用id查询TimerTask
// @Tags TimerTask
// @Summary 用id查询TimerTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query interfacecase.TimerTask true "用id查询TimerTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /task/findTimerTask [get]
func (taskApi *TimerTaskApi) FindTimerTask(c *gin.Context) {
	var task interfacecase.ApiTimerTask
	_ = c.ShouldBindQuery(&task)
	task.ProjectID = utils.GetUserProject(c)
	if err, retask := taskService.GetTimerTask(task.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retask": retask}, c)
	}
}

// GetTimerTaskList 分页获取TimerTask列表
// @Tags TimerTask
// @Summary 分页获取TimerTask列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query interfacecaseReq.TimerTaskSearch true "分页获取TimerTask列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /task/getTimerTaskList [get]
func (taskApi *TimerTaskApi) GetTimerTaskList(c *gin.Context) {
	var pageInfo interfacecaseReq.TimerTaskSearch
	_ = c.ShouldBindQuery(&pageInfo)
	pageInfo.ProjectID = utils.GetUserProject(c)
	if err, list, total := taskService.GetTimerTaskInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
