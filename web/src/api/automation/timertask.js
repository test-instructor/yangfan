import service from '@/utils/request'
import { useUserStore } from '@/pinia/modules/user'
const userStore = useUserStore()
const projectId = userStore.userInfo.projectId
// @Tags TimerTask
// @Summary 创建定时任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.TimerTask true "创建定时任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /tk/createTimerTask [post]
export const createTimerTask = (data) => {
  data.projectId = projectId
  return service({
    url: '/tk/createTimerTask',
    method: 'post',
    data
  })
}

// @Tags TimerTask
// @Summary 删除定时任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.TimerTask true "删除定时任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tk/deleteTimerTask [delete]
export const deleteTimerTask = (params) => {
  params.projectId = projectId
  return service({
    url: '/tk/deleteTimerTask',
    method: 'delete',
    params
  })
}

// @Tags TimerTask
// @Summary 批量删除定时任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除定时任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tk/deleteTimerTask [delete]
export const deleteTimerTaskByIds = (params) => {
  params.projectId = projectId
  return service({
    url: '/tk/deleteTimerTaskByIds',
    method: 'delete',
    params
  })
}

// @Tags TimerTask
// @Summary 更新定时任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.TimerTask true "更新定时任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tk/updateTimerTask [put]
export const updateTimerTask = (data) => {
  data.projectId = projectId
  return service({
    url: '/tk/updateTimerTask',
    method: 'put',
    data
  })
}

// @Tags TimerTask
// @Summary 用id查询定时任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.TimerTask true "用id查询定时任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tk/findTimerTask [get]
export const findTimerTask = (params) => {
  params.projectId = projectId
  return service({
    url: '/tk/findTimerTask',
    method: 'get',
    params
  })
}

// @Tags TimerTask
// @Summary 分页获取定时任务列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取定时任务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tk/getTimerTaskList [get]
export const getTimerTaskList = (params) => {
  params.projectId = projectId
  return service({
    url: '/tk/getTimerTaskList',
    method: 'get',
    params
  })
}

// @Tags TimerTask
// @Summary 不需要鉴权的定时任务接口
// @Accept application/json
// @Produce application/json
// @Param data query automationReq.TimerTaskSearch true "分页获取定时任务列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /tk/getTimerTaskPublic [get]
export const getTimerTaskPublic = () => {
  return service({
    url: '/tk/getTimerTaskPublic',
    method: 'get',
  })
}

// 任务-用例关联：添加
export const addTimerTaskCase = (data) => {
  data.projectId = projectId
  return service({
    url: '/tk/addTimerTaskCase',
    method: 'post',
    data
  })
}

// 任务-用例关联：排序
export const sortTimerTaskCase = (data) => {
  data.projectId = projectId
  return service({
    url: '/tk/sortTimerTaskCase',
    method: 'post',
    data
  })
}

// 任务-用例关联：删除
export const delTimerTaskCase = (params) => {
  params.projectId = projectId
  return service({
    url: '/tk/delTimerTaskCase',
    method: 'delete',
    params
  })
}

// 任务-用例关联：获取任务引用的用例列表
export const getTimerTaskCases = (params) => {
  params.projectId = projectId
  return service({
    url: '/tk/getTimerTaskCases',
    method: 'get',
    params
  })
}
