import service from '@/utils/request'
import { useUserStore } from '@/pinia/modules/user'
const userStore = useUserStore()
const projectId = userStore.userInfo.projectId
// @Tags HarmonyDeviceOptions
// @Summary 创建设备选项
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.HarmonyDeviceOptions true "创建设备选项"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /hdo/createHarmonyDeviceOptions [post]
export const createHarmonyDeviceOptions = (data) => {
  data.projectId = projectId
  return service({
    url: '/hdo/createHarmonyDeviceOptions',
    method: 'post',
    data
  })
}

// @Tags HarmonyDeviceOptions
// @Summary 删除设备选项
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.HarmonyDeviceOptions true "删除设备选项"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hdo/deleteHarmonyDeviceOptions [delete]
export const deleteHarmonyDeviceOptions = (params) => {
  params.projectId = projectId
  return service({
    url: '/hdo/deleteHarmonyDeviceOptions',
    method: 'delete',
    params
  })
}

// @Tags HarmonyDeviceOptions
// @Summary 批量删除设备选项
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除设备选项"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hdo/deleteHarmonyDeviceOptions [delete]
export const deleteHarmonyDeviceOptionsByIds = (params) => {
  params.projectId = projectId
  return service({
    url: '/hdo/deleteHarmonyDeviceOptionsByIds',
    method: 'delete',
    params
  })
}

// @Tags HarmonyDeviceOptions
// @Summary 更新设备选项
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.HarmonyDeviceOptions true "更新设备选项"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hdo/updateHarmonyDeviceOptions [put]
export const updateHarmonyDeviceOptions = (data) => {
  data.projectId = projectId
  return service({
    url: '/hdo/updateHarmonyDeviceOptions',
    method: 'put',
    data
  })
}

// @Tags HarmonyDeviceOptions
// @Summary 用id查询设备选项
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.HarmonyDeviceOptions true "用id查询设备选项"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hdo/findHarmonyDeviceOptions [get]
export const findHarmonyDeviceOptions = (params) => {
  params.projectId = projectId
  return service({
    url: '/hdo/findHarmonyDeviceOptions',
    method: 'get',
    params
  })
}

// @Tags HarmonyDeviceOptions
// @Summary 分页获取设备选项列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取设备选项列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hdo/getHarmonyDeviceOptionsList [get]
export const getHarmonyDeviceOptionsList = (params) => {
  params.projectId = projectId
  return service({
    url: '/hdo/getHarmonyDeviceOptionsList',
    method: 'get',
    params
  })
}

// @Tags HarmonyDeviceOptions
// @Summary 不需要鉴权的设备选项接口
// @Accept application/json
// @Produce application/json
// @Param data query platformReq.HarmonyDeviceOptionsSearch true "分页获取设备选项列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /hdo/getHarmonyDeviceOptionsPublic [get]
export const getHarmonyDeviceOptionsPublic = () => {
  return service({
    url: '/hdo/getHarmonyDeviceOptionsPublic',
    method: 'get',
  })
}
