import service from '@/utils/request'
import { useUserStore } from '@/pinia/modules/user'
const userStore = useUserStore()
const projectId = userStore.userInfo.projectId
// @Tags EnvDetail
// @Summary 创建环境详情
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.EnvDetail true "创建环境详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ed/createEnvDetail [post]
export const createEnvDetail = (data) => {
  data.projectId = projectId
  return service({
    url: '/ed/createEnvDetail',
    method: 'post',
    data
  })
}

// @Tags EnvDetail
// @Summary 删除环境详情
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.EnvDetail true "删除环境详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ed/deleteEnvDetail [delete]
export const deleteEnvDetail = (params) => {
  params.projectId = projectId
  return service({
    url: '/ed/deleteEnvDetail',
    method: 'delete',
    params
  })
}

// @Tags EnvDetail
// @Summary 批量删除环境详情
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除环境详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ed/deleteEnvDetail [delete]
export const deleteEnvDetailByIds = (params) => {
  params.projectId = projectId
  return service({
    url: '/ed/deleteEnvDetailByIds',
    method: 'delete',
    params
  })
}

// @Tags EnvDetail
// @Summary 更新环境详情
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.EnvDetail true "更新环境详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ed/updateEnvDetail [put]
export const updateEnvDetail = (data) => {
  data.projectId = projectId
  return service({
    url: '/ed/updateEnvDetail',
    method: 'put',
    data
  })
}

// @Tags EnvDetail
// @Summary 用id查询环境详情
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.EnvDetail true "用id查询环境详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ed/findEnvDetail [get]
export const findEnvDetail = (params) => {
  params.projectId = projectId
  return service({
    url: '/ed/findEnvDetail',
    method: 'get',
    params
  })
}

// @Tags EnvDetail
// @Summary 分页获取环境详情列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取环境详情列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ed/getEnvDetailList [get]
export const getEnvDetailList = (params) => {
  params.projectId = projectId
  return service({
    url: '/ed/getEnvDetailList',
    method: 'get',
    params
  })
}

// @Tags EnvDetail
// @Summary 不需要鉴权的环境详情接口
// @Accept application/json
// @Produce application/json
// @Param data query platformReq.EnvDetailSearch true "分页获取环境详情列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ed/getEnvDetailPublic [get]
export const getEnvDetailPublic = () => {
  return service({
    url: '/ed/getEnvDetailPublic',
    method: 'get',
  })
}
