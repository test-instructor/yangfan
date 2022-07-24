
import service from '@/utils/request'

const project = JSON.parse(window.localStorage.getItem('project')).ID
const baseURL = '/ac/' + project

// @Tags ApiConfig
// @Summary 创建ApiConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ApiConfig true "创建ApiConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ac/createApiConfig [post]
export const createApiConfig = (data) => {
    return service({
        url: baseURL + '/createApiConfig',
        method: 'post',
        data
    })
}

// @Tags ApiConfig
// @Summary 删除ApiConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ApiConfig true "删除ApiConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ac/deleteApiConfig [delete]
export const deleteApiConfig = (data) => {
    return service({
        url: baseURL + '/deleteApiConfig',
        method: 'delete',
        data
    })
}

// @Tags ApiConfig
// @Summary 删除ApiConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ApiConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ac/deleteApiConfig [delete]
export const deleteApiConfigByIds = (data) => {
    return service({
        url: baseURL + '/deleteApiConfigByIds',
        method: 'delete',
        data
    })
}

// @Tags ApiConfig
// @Summary 更新ApiConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ApiConfig true "更新ApiConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ac/updateApiConfig [put]
export const updateApiConfig = (data) => {
    return service({
        url: baseURL + '/updateApiConfig',
        method: 'put',
        data
    })
}

// @Tags ApiConfig
// @Summary 用id查询ApiConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ApiConfig true "用id查询ApiConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ac/findApiConfig [get]
export const findApiConfig = (params) => {
    return service({
        url: baseURL + '/findApiConfig',
        method: 'get',
        params
    })
}

// @Tags ApiConfig
// @Summary 分页获取ApiConfig列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ApiConfig列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ac/getApiConfigList [get]
export const getApiConfigList = (params) => {
    return service({
        url: baseURL + '/getApiConfigList',
        method: 'get',
        params
    })
}
