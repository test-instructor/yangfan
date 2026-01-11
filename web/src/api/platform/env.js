import service from '@/utils/request'
import { useUserStore } from '@/pinia/modules/user'
const userStore = useUserStore()
const projectId = userStore.userInfo.projectId

// @Tags Env
// @Summary 创建环境配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Env true "创建环境配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /env/createEnv [post]
export const createEnv = (data) => {
  data.projectId = projectId
  return service({
    url: '/env/createEnv',
    method: 'post',
    data
  })
}

// @Tags Env
// @Summary 删除环境配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Env true "删除环境配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /env/deleteEnv [delete]
  export const deleteEnv = (params) => {
    params.projectId = projectId
  return service({
    url: '/env/deleteEnv',
    method: 'delete',
    params
  })
}

// @Tags Env
// @Summary 批量删除环境配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除环境配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /env/deleteEnv [delete]
export const deleteEnvByIds = (params) => {
  params.projectId = projectId
  return service({
    url: '/env/deleteEnvByIds',
    method: 'delete',
    params
  })
}

// @Tags Env
// @Summary 更新环境配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Env true "更新环境配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /env/updateEnv [put]
export const updateEnv = (data) => {
  data.projectId = projectId
  return service({
    url: '/env/updateEnv',
    method: 'put',
    data
  })
}

// @Tags Env
// @Summary 用id查询环境配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Env true "用id查询环境配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /env/findEnv [get]
export const findEnv = (params) => {
  params.projectId = projectId
  return service({
    url: '/env/findEnv',
    method: 'get',
    params
  })
}

// @Tags Env
// @Summary 分页获取环境配置列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取环境配置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /env/getEnvList [get]
export const getEnvList = (params) => {
  params.projectId = projectId
  return service({
    url: '/env/getEnvList',
    method: 'get',
    params
  })
}

// @Tags Env
// @Summary 不需要鉴权的环境配置接口
// @Accept application/json
// @Produce application/json
// @Param data query platformReq.EnvSearch true "分页获取环境配置列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /env/getEnvPublic [get]
export const getEnvPublic = () => {
  return service({
    url: '/env/getEnvPublic',
    method: 'get',
  })
}
