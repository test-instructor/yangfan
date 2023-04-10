import service from '@/utils/request'

const project = JSON.parse(window.localStorage.getItem('project')).ID
const baseURL = '/env/' + project

// 创建环境
export const createEnv = (data) => {
    return service({
        url: baseURL + '/createEnv',
        method: 'post',
        data
    })
}

// 更新环境
export const updateEnv = (data) => {
    return service({
        url: baseURL + '/updateEnv',
        method: 'put',
        data
    })
}

// 删除环境
export const deleteEnv = (data) => {
    return service({
        url: baseURL + '/deleteEnv',
        method: 'delete',
        data
    })
}

// 通过id查找环境
export const findEnv = (data) => {
    return service({
        url: baseURL + '/findEnv',
        method: 'post',
        data
    })
}

// 查询环境列表
export const getEnvList = (data) => {
    return service({
        url: baseURL + '/getEnvList',
        method: 'get',
        data
    })
}


// 创建变量
export const createEnvVariable = (data) => {
    return service({
        url: baseURL + '/createEnvVariable',
        method: 'post',
        data
    })
}

// 删除变量
export const deleteEnvVariable = (data) => {
    return service({
        url: baseURL + '/deleteEnvVariable',
        method: 'delete',
        data
    })
}

// 通过id查找变量
export const findEnvVariable = (params) => {
    return service({
        url: baseURL + '/findEnvVariable',
        method: 'get',
        params
    })
}

// 查询变量列表
export const getEnvVariableList = (params) => {
    return service({
        url: baseURL + '/getEnvVariableList',
        method: 'get',
        params
    })
}