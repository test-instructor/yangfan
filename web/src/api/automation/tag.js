import service from '@/utils/request'
import { useUserStore } from '@/pinia/modules/user'

const userStore = useUserStore()
const projectId = userStore.userInfo.projectId

export const createTag = (data) => {
  data.projectId = projectId
  return service({
    url: '/tk/tag/createTag',
    method: 'post',
    data
  })
}

export const updateTag = (data) => {
  data.projectId = projectId
  return service({
    url: '/tk/tag/updateTag',
    method: 'put',
    data
  })
}

export const deleteTag = (params) => {
  params.projectId = projectId
  return service({
    url: '/tk/tag/deleteTag',
    method: 'delete',
    params
  })
}

export const getTagList = (params) => {
  params.projectId = projectId
  return service({
    url: '/tk/tag/getTagList',
    method: 'get',
    params
  })
}