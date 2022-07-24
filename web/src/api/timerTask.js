import service from '@/utils/request'

const project = JSON.parse(window.localStorage.getItem('project')).ID
const baseURL = '/task/' + project


export const findTaskTestCase = (params) => {
    return service({
        url: baseURL + '/findTaskTestCase',
        method: 'get',
        params
    })
}

export const addTaskTestCase = (data) => {
    return service({
        url: baseURL + '/addTaskTestCase',
        method: 'post',
        data
    })
}

export const setTaskCase = (data) => {
    return service({
        url: baseURL + '/setTaskCase',
        method: 'post',
        data: data
    })
}


// @Tags TimerTask
// @Summary 创建TimerTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TimerTask true "创建TimerTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /task/createTimerTask [post]
export const createTimerTask = (data) => {
    return service({
        url: baseURL + '/createTimerTask',
        method: 'post',
        data
    })
}

// @Tags TimerTask
// @Summary 删除TimerTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TimerTask true "删除TimerTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /task/deleteTimerTask [delete]
export const deleteTimerTask = (data) => {
    return service({
        url: baseURL + '/deleteTimerTask',
        method: 'delete',
        data
    })
}

// @Tags TimerTask
// @Summary 删除TimerTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TimerTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /task/deleteTimerTask [delete]
export const deleteTimerTaskByIds = (data) => {
    return service({
        url: baseURL + '/deleteTimerTaskByIds',
        method: 'delete',
        data
    })
}

// @Tags TimerTask
// @Summary 更新TimerTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TimerTask true "更新TimerTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /task/updateTimerTask [put]
export const updateTimerTask = (data) => {
    return service({
        url: baseURL + '/updateTimerTask',
        method: 'put',
        data
    })
}

// @Tags TimerTask
// @Summary 用id查询TimerTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TimerTask true "用id查询TimerTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /task/findTimerTask [get]
export const findTimerTask = (params) => {
    return service({
        url: baseURL + '/findTimerTask',
        method: 'get',
        params
    })
}

// @Tags TimerTask
// @Summary 分页获取TimerTask列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TimerTask列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /task/getTimerTaskList [get]
export const getTimerTaskList = (params) => {
    return service({
        url: baseURL + '/getTimerTaskList',
        method: 'get',
        params
    })
}
