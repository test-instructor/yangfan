import service from '@/utils/request'

const project = JSON.parse(window.localStorage.getItem('project')).ID
const baseURL = '/apicase/' + project

// @Tags InterfaceTemplate
// @Summary 创建InterfaceTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InterfaceTemplate true "创建InterfaceTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apicase/createInterfaceTemplate [post]
export const createInterfaceTemplate = (data, params) => {
    return service({
        url: baseURL + '/createInterfaceTemplate',
        method: 'post',
        data,
        params
    })
}

// @Tags InterfaceTemplate
// @Summary 删除InterfaceTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InterfaceTemplate true "删除InterfaceTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /apicase/deleteInterfaceTemplate [delete]
export const deleteInterfaceTemplate = (data) => {
    return service({
        url: baseURL + '/deleteInterfaceTemplate',
        method: 'delete',
        data
    })
}

// @Tags InterfaceTemplate
// @Summary 删除InterfaceTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除InterfaceTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /apicase/deleteInterfaceTemplate [delete]
export const deleteInterfaceTemplateByIds = (data) => {
    return service({
        url: '/apicase/deleteInterfaceTemplateByIds',
        method: 'delete',
        data
    })
}

// @Tags InterfaceTemplate
// @Summary 更新InterfaceTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InterfaceTemplate true "更新InterfaceTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /apicase/updateInterfaceTemplate [put]
export const updateInterfaceTemplate = (data, params) => {
    return service({
        url: baseURL + '/updateInterfaceTemplate',
        method: 'put',
        data,
        params,
    })
}

// @Tags InterfaceTemplate
// @Summary 用id查询InterfaceTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.InterfaceTemplate true "用id查询InterfaceTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /apicase/findInterfaceTemplate [get]
export const findInterfaceTemplate = (params) => {
    return service({
        url: baseURL + '/findInterfaceTemplate',
        method: 'get',
        params
    })
}

// @Tags InterfaceTemplate
// @Summary 分页获取InterfaceTemplate列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取InterfaceTemplate列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apicase/getInterfaceTemplateList [get]
export const getInterfaceTemplateList = (params) => {
    return service({
        url: baseURL + '/getInterfaceTemplateList',
        method: 'get',
        params
    })
}

export const getDebugTalk = (data) => {
    return service({
        url: baseURL + '/getDebugTalk',
        method: 'post',
        data
    })
}

export const updateDebugTalk = (data) => {
    return service({
        url: baseURL + '/updateDebugTalk',
        method: 'put',
        data
    })
}
