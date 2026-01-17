import service from '@/utils/request'
import { useUserStore } from '@/pinia/modules/user'

const getProjectId = () => useUserStore().userInfo?.projectId || 0

export const createTag = (data) => {
  data.projectId = getProjectId()
  return service({
    url: '/tk/tag/createTag',
    method: 'post',
    data
  })
}

export const updateTag = (data) => {
  data.projectId = getProjectId()
  return service({
    url: '/tk/tag/updateTag',
    method: 'put',
    data
  })
}

export const deleteTag = (params) => {
  params.projectId = getProjectId()
  return service({
    url: '/tk/tag/deleteTag',
    method: 'delete',
    params
  })
}

export const getTagList = (params) => {
  params.projectId = getProjectId()
  return service({
    url: '/tk/tag/getTagList',
    method: 'get',
    params
  })
}
