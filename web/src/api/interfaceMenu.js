import service from '@/utils/request'

const project = JSON.parse(window.localStorage.getItem('project')).ID
const baseURL = '/case/' + project

export function getTree(params) {
    return service({
        url: baseURL + '/getApiMenuList',
        method: 'get',
        params
    })
}

export function addTree(data, params) {
    return service({
        url: baseURL + '/createApiMenu',
        method: 'post',
        data,
        params
    })
}

export function editTree(data, params) {
    return service({
        url: baseURL + '/updateApiMenu',
        method: 'put',
        data,
        params
    })
}

export function delTree(data, params) {
    return service({
        url: baseURL + '/deleteApiMenu',
        method: 'delete',
        data,
        params
    })
}
