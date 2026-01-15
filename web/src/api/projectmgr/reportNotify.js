import service from '@/utils/request'

export const createReportNotifyChannel = (data) => {
  return service({
    url: '/nt/createReportNotifyChannel',
    method: 'post',
    data
  })
}

export const deleteReportNotifyChannel = (params) => {
  return service({
    url: '/nt/deleteReportNotifyChannel',
    method: 'delete',
    params
  })
}

export const deleteReportNotifyChannelByIds = (params) => {
  return service({
    url: '/nt/deleteReportNotifyChannelByIds',
    method: 'delete',
    params
  })
}

export const updateReportNotifyChannel = (data) => {
  return service({
    url: '/nt/updateReportNotifyChannel',
    method: 'put',
    data
  })
}

export const findReportNotifyChannel = (params) => {
  return service({
    url: '/nt/findReportNotifyChannel',
    method: 'get',
    params
  })
}

export const getReportNotifyChannelList = (params) => {
  return service({
    url: '/nt/getReportNotifyChannelList',
    method: 'get',
    params
  })
}

export const getAutoReportNotifyStatus = (params) => {
  return service({
    url: '/nt/getAutoReportNotifyStatus',
    method: 'get',
    params
  })
}

