import service from '@/utils/request'
import { useUserStore } from '@/pinia/modules/user'
const userStore = useUserStore()
const projectId = userStore.userInfo.projectId
// @Tags Request
// @Summary 创建请求
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Request true "创建请求"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /req/createRequest [post]
export const createRequest = (data) => {
  data.projectId = projectId
  return service({
    url: '/req/createRequest',
    method: 'post',
    data
  })
}

// @Tags Request
// @Summary 删除请求
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Request true "删除请求"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /req/deleteRequest [delete]
export const deleteRequest = (params) => {
  params.projectId = projectId
  return service({
    url: '/req/deleteRequest',
    method: 'delete',
    params
  })
}

// @Tags Request
// @Summary 批量删除请求
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除请求"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /req/deleteRequest [delete]
export const deleteRequestByIds = (params) => {
  params.projectId = projectId
  return service({
    url: '/req/deleteRequestByIds',
    method: 'delete',
    params
  })
}

// @Tags Request
// @Summary 更新请求
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Request true "更新请求"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /req/updateRequest [put]
export const updateRequest = (data) => {
  data.projectId = projectId
  return service({
    url: '/req/updateRequest',
    method: 'put',
    data
  })
}

// @Tags Request
// @Summary 用id查询请求
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Request true "用id查询请求"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /req/findRequest [get]
export const findRequest = (params) => {
  params.projectId = projectId
  return service({
    url: '/req/findRequest',
    method: 'get',
    params
  })
}

// @Tags Request
// @Summary 分页获取请求列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取请求列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /req/getRequestList [get]
export const getRequestList = (params) => {
  params.projectId = projectId
  return service({
    url: '/req/getRequestList',
    method: 'get',
    params
  })
}

// @Tags Request
// @Summary 不需要鉴权的请求接口
// @Accept application/json
// @Produce application/json
// @Param data query automationReq.RequestSearch true "分页获取请求列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /req/getRequestPublic [get]
export const getRequestPublic = () => {
  return service({
    url: '/req/getRequestPublic',
    method: 'get',
  })
}
