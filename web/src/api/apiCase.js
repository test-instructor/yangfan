import service from '@/utils/request'

const project = JSON.parse(window.localStorage.getItem('project')).ID
const baseURL = '/testcase/' + project


export const findApiTestCase = (params) => {
    return service({
        url: baseURL + '/findApiTestCase',
        method: 'get',
        params
    })
}

export const addApiTestCase = (data) => {
    return service({
        url: baseURL + '/addApiTestCase',
        method: 'post',
        data
    })
}

export const addApisCase = (data) => {
    return service({
        url: baseURL + '/addApisCase',
        method: 'post',
        data
    })
}

export const setApisCase = (data) => {
    return service({
        url: baseURL + '/setApisCase',
        method: 'post',
        data: data
    })
}

export const sortApisCase = (data) => {
    return service({
        url: baseURL + '/sortApisCase',
        method: 'post',
        data
    })
}

export const delApisCase = (data) => {
    return service({
        url: baseURL + '/delApisCase',
        method: 'delete',
        data
    })
}



// @Tags ApiCase
// @Summary 创建ApiCase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ApiCase true "创建ApiCase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /task/createApiCase [post]
export const createApiCase = (data,params) => {
    return service({
        url: baseURL + '/createApiCase',
        method: 'post',
        data,
        params
    })
}

// @Tags ApiCase
// @Summary 删除ApiCase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ApiCase true "删除ApiCase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /task/deleteApiCase [delete]
export const deleteApiCase = (data) => {
    return service({
        url: baseURL + '/deleteApiCase',
        method: 'delete',
        data
    })
}

// @Tags ApiCase
// @Summary 删除ApiCase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ApiCase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /task/deleteApiCase [delete]
export const deleteApiCaseByIds = (data) => {
    return service({
        url: baseURL + '/deleteApiCaseByIds',
        method: 'delete',
        data
    })
}

// @Tags ApiCase
// @Summary 更新ApiCase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ApiCase true "更新ApiCase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /task/updateApiCase [put]
export const updateApiCase = (data, params) => {
    return service({
        url: baseURL + '/updateApiCase',
        method: 'put',
        data,
        params
    })
}

// @Tags ApiCase
// @Summary 用id查询ApiCase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ApiCase true "用id查询ApiCase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /task/findApiCase [get]
export const findApiCase = (params) => {
    return service({
        url: baseURL + '/findApiCase',
        method: 'get',
        params
    })
}

// @Tags ApiCase
// @Summary 分页获取ApiCase列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ApiCase列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /task/getApiCaseList [get]
export const getApiCaseList = (params) => {
    return service({
        url: baseURL + '/getApiCaseList',
        method: 'get',
        params
    })
}
