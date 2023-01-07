import service from '@/utils/request'

const project = JSON.parse(window.localStorage.getItem('project')).ID
const baseURL = '/case/run/' + project

export const runTestCaseStep = (data) => {
    return service({
        url: baseURL + '/runTestCaseStep',
        method: 'post',
        data
    })
}

export const runBoomer = (data) => {
    return service({
        url: baseURL + '/runBoomer',
        method: 'post',
        data
    })
}

export const runBoomerDebug = (data) => {
    return service({
        url: baseURL + '/runBoomerDebug',
        method: 'post',
        data
    })
}

export const runApiCase = (data) => {
    return service({
        url: baseURL + '/runApiCase',
        method: 'post',
        data
    })
}

export const runApi = (data) => {
    return service({
        url: baseURL + '/runApi',
        method: 'post',
        data
    })
}

export const runTimerTask = (data) => {
    return service({
        url: baseURL + '/runTimerTask',
        method: 'post',
        data
    })
}


