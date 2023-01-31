import service from '@/utils/request'

const project = JSON.parse(window.localStorage.getItem('project')).ID
const baseURL = '/case/' + project

// @Tags HrpPyPkg
// @Summary 创建HrpPyPkg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HrpPyPkg true "创建HrpPyPkg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PyPkg/installHrpPyPkg [post]
export const installHrpPyPkg = (data) => {
    return service({
        url: baseURL + "/pyPkg/installPyPkg",
        method: 'post',
        data
    })
}

// @Tags HrpPyPkg
// @Summary 删除HrpPyPkg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HrpPyPkg true "删除HrpPyPkg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PyPkg/uninstallHrpPyPkg [delete]
export const uninstallHrpPyPkg = (data) => {
    return service({
        url: baseURL + "/pyPkg/uninstallPyPkg",
        method: 'post',
        data
    })
}

// @Tags HrpPyPkg
// @Summary 删除HrpPyPkg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HrpPyPkg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PyPkg/uninstallHrpPyPkg [delete]
export const uninstallHrpPyPkgByIds = (data) => {
    return service({
        url: baseURL + "/pyPkg/uninstallPyPkg",
        method: 'post',
        data
    })
}

// @Tags HrpPyPkg
// @Summary 更新HrpPyPkg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HrpPyPkg true "更新HrpPyPkg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PyPkg/updateHrpPyPkg [put]
export const updateHrpPyPkg = (data) => {
    return service({
        url: baseURL + "/pyPkg/updatePyPkg",
        method: 'post',
        data
    })
}

// @Tags HrpPyPkg
// @Summary 用id查询HrpPyPkg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HrpPyPkg true "用id查询HrpPyPkg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PyPkg/searchHrpPyPkg [post]
export const searchHrpPyPkg = (data,params) => {
    return service({
        url: baseURL + '/pyPkg/searchPyPkg',
        method: 'post',
        params,
        data
    })
}

// @Tags HrpPyPkg
// @Summary 分页获取HrpPyPkg列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HrpPyPkg列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PyPkg/getHrpPyPkgList [get]
export const getHrpPyPkgList = (params) => {
    return service({
        url: baseURL + '/pyPkg/pyPkgList',
        method: 'get',
        params
    })
}