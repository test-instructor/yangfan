import service from '@/utils/request'

const project = JSON.parse(window.localStorage.getItem('project')).ID
const baseURL = '/case/' + project

// @Tags ApiMenu
// @Summary 创建ApiMenu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ApiMenu true "创建ApiMenu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apicase/createApiMenu [post]
export const createApiMenu = (data) => {
    return service({
        url: baseURL + '/createApiMenu',
        method: 'post',
        data
    })
}

// @Tags ApiMenu
// @Summary 删除ApiMenu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ApiMenu true "删除ApiMenu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /apicase/deleteApiMenu [delete]
export const deleteApiMenu = (data) => {
    return service({
        url: baseURL + '/deleteApiMenu',
        method: 'delete',
        data
    })
}

// @Tags ApiMenu
// @Summary 删除ApiMenu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ApiMenu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /apicase/deleteApiMenu [delete]
export const deleteApiMenuByIds = (data) => {
    return service({
        url: baseURL + '/deleteApiMenuByIds',
        method: 'delete',
        data
    })
}

// @Tags ApiMenu
// @Summary 更新ApiMenu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ApiMenu true "更新ApiMenu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /apicase/updateApiMenu [put]
export const updateApiMenu = (data) => {
    return service({
        url: baseURL + '/updateApiMenu',
        method: 'put',
        data
    })
}

// @Tags ApiMenu
// @Summary 用id查询ApiMenu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ApiMenu true "用id查询ApiMenu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /apicase/findApiMenu [get]
export const findApiMenu = (params) => {
    return service({
        url: baseURL + '/findApiMenu',
        method: 'get',
        params
    })
}

// @Tags ApiMenu
// @Summary 分页获取ApiMenu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ApiMenu列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apicase/getApiMenuList [get]
export const getApiMenuList = (params) => {
    return service({
        url: baseURL + '/getApiMenuList',
        method: 'get',
        params
    })
}
