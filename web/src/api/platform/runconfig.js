import service from '@/utils/request'
import { useUserStore } from '@/pinia/modules/user'
const userStore = useUserStore()
const projectId = userStore.userInfo.projectId
// @Tags RunConfig
// @Summary 创建运行配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.RunConfig true "创建运行配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /rc/createRunConfig [post]
export const createRunConfig = (data) => {
  data.projectId = projectId
  return service({
    url: '/rc/createRunConfig',
    method: 'post',
    data
  })
}

// @Tags RunConfig
// @Summary 删除运行配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.RunConfig true "删除运行配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rc/deleteRunConfig [delete]
export const deleteRunConfig = (params) => {
  params.projectId = projectId
  return service({
    url: '/rc/deleteRunConfig',
    method: 'delete',
    params
  })
}

// @Tags RunConfig
// @Summary 批量删除运行配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除运行配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rc/deleteRunConfig [delete]
export const deleteRunConfigByIds = (params) => {
  params.projectId = projectId
  return service({
    url: '/rc/deleteRunConfigByIds',
    method: 'delete',
    params
  })
}

// @Tags RunConfig
// @Summary 更新运行配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.RunConfig true "更新运行配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /rc/updateRunConfig [put]
export const updateRunConfig = (data) => {
  data.projectId = projectId
  return service({
    url: '/rc/updateRunConfig',
    method: 'put',
    data
  })
}

// @Tags RunConfig
// @Summary 用id查询运行配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.RunConfig true "用id查询运行配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /rc/findRunConfig [get]
export const findRunConfig = (params) => {
  params.projectId = projectId
  return service({
    url: '/rc/findRunConfig',
    method: 'get',
    params
  })
}

// @Tags RunConfig
// @Summary 分页获取运行配置列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取运行配置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rc/getRunConfigList [get]
export const getRunConfigList = (params) => {
  params.projectId = projectId
  return service({
    url: '/rc/getRunConfigList',
    method: 'get',
    params
  })
}

// @Tags RunConfig
// @Summary 不需要鉴权的运行配置接口
// @Accept application/json
// @Produce application/json
// @Param data query platformReq.RunConfigSearch true "分页获取运行配置列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /rc/getRunConfigPublic [get]
export const getRunConfigPublic = () => {
  return service({
    url: '/rc/getRunConfigPublic',
    method: 'get',
  })
}
