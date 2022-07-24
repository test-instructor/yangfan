import service from '@/utils/request'

const project = JSON.parse(window.localStorage.getItem('project')).ID
const baseURL = '/apicase/report/' + project


export const deleteReport = (data) => {
    return service({
        url: baseURL + '/deleteReport',
        method: 'delete',
        data
    })
}

// @Tags TestCase
// @Summary 删除TestCase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TestCase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /apicase/deleteTestCase [delete]
export const deleteReportIds = (data) => {
    return service({
        url: baseURL + '/deleteReportIds',
        method: 'delete',
        data
    })
}


// @Tags TestCase
// @Summary 用id查询TestCase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TestCase true "用id查询TestCase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /apicase/findTestCase [get]
export const findReport = (params) => {
    return service({
        url: baseURL + '/findReport',
        method: 'get',
        params
    })
}


export const delReport = (data) => {
    return service({
        url: baseURL + '/delReport',
        method: 'delete',
        data
    })
}


// @Tags TestCase
// @Summary 分页获取TestCase列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TestCase列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apicase/getTestCaseList [get]
export const getReportList = (params) => {
    return service({
        url: baseURL + '/getReportList',
        method: 'get',
        params
    })
}
