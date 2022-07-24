import service from '@/utils/request'

const project = JSON.parse(window.localStorage.getItem('project')).ID
const baseURL = '/apicase/run/' + project

export const runTestCase = (data) => {
    return service({
        url: baseURL + '/runTestCase',
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

