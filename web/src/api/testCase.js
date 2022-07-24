import service from '@/utils/request'

const project = JSON.parse(window.localStorage.getItem('project')).ID
const baseURL = '/apicase/' + project

// @Tags TestCase
// @Summary 创建TestCase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestCase true "创建TestCase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apicase/createTestCase [post]
export const createTestCase = (data, params) => {
    return service({
        url: baseURL + '/createTestCase',
        method: 'post',
        data,
        params
    })
}

export const sortTestCase = (data) => {
    return service({
        url: baseURL + '/sortTestCase',
        method: 'post',
        data
    })
}


// @Tags TestCase
// @Summary 删除TestCase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestCase true "删除TestCase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /apicase/deleteTestCase [delete]
export const deleteTestCase = (data) => {
    return service({
        url: baseURL + '/deleteTestCase',
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
export const deleteTestCaseByIds = (data) => {
    return service({
        url: baseURL + '/deleteTestCaseByIds',
        method: 'delete',
        data
    })
}

// @Tags TestCase
// @Summary 更新TestCase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestCase true "更新TestCase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /apicase/updateTestCase [put]
export const updateTestCase = (data, params) => {
    return service({
        url: baseURL + '/updateTestCase',
        method: 'put',
        data,
        params
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
export const findTestCase = (params) => {
    return service({
        url: baseURL + '/findTestCase',
        method: 'get',
        params
    })
}

export const addTestCase = (data) => {
    return service({
        url: baseURL + '/addTestCase',
        method: 'post',
        data
    })
}

export const delTestCase = (data) => {
    return service({
        url: baseURL + '/delTestCase',
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
export const getTestCaseList = (params) => {
    return service({
        url: baseURL + '/getTestCaseList',
        method: 'get',
        params
    })
}
