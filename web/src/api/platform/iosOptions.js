import service from '@/utils/request'
import { useUserStore } from '@/pinia/modules/user'
const userStore = useUserStore()
const projectId = userStore.userInfo.projectId
// @Tags IOSDeviceOptions
// @Summary 创建iOS设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.IOSDeviceOptions true "创建iOS设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ido/createIOSDeviceOptions [post]
export const createIOSDeviceOptions = (data) => {
  data.projectId = projectId
  return service({
    url: '/ido/createIOSDeviceOptions',
    method: 'post',
    data
  })
}

// @Tags IOSDeviceOptions
// @Summary 删除iOS设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.IOSDeviceOptions true "删除iOS设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ido/deleteIOSDeviceOptions [delete]
export const deleteIOSDeviceOptions = (params) => {
  params.projectId = projectId
  return service({
    url: '/ido/deleteIOSDeviceOptions',
    method: 'delete',
    params
  })
}

// @Tags IOSDeviceOptions
// @Summary 批量删除iOS设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除iOS设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ido/deleteIOSDeviceOptions [delete]
export const deleteIOSDeviceOptionsByIds = (params) => {
  params.projectId = projectId
  return service({
    url: '/ido/deleteIOSDeviceOptionsByIds',
    method: 'delete',
    params
  })
}

// @Tags IOSDeviceOptions
// @Summary 更新iOS设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.IOSDeviceOptions true "更新iOS设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ido/updateIOSDeviceOptions [put]
export const updateIOSDeviceOptions = (data) => {
  data.projectId = projectId
  return service({
    url: '/ido/updateIOSDeviceOptions',
    method: 'put',
    data
  })
}

// @Tags IOSDeviceOptions
// @Summary 用id查询iOS设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.IOSDeviceOptions true "用id查询iOS设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ido/findIOSDeviceOptions [get]
export const findIOSDeviceOptions = (params) => {
  params.projectId = projectId
  return service({
    url: '/ido/findIOSDeviceOptions',
    method: 'get',
    params
  })
}

// @Tags IOSDeviceOptions
// @Summary 分页获取iOS设备列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取iOS设备列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ido/getIOSDeviceOptionsList [get]
export const getIOSDeviceOptionsList = (params) => {
  params.projectId = projectId
  return service({
    url: '/ido/getIOSDeviceOptionsList',
    method: 'get',
    params
  })
}

// @Tags IOSDeviceOptions
// @Summary 不需要鉴权的iOS设备接口
// @Accept application/json
// @Produce application/json
// @Param data query platformReq.IOSDeviceOptionsSearch true "分页获取iOS设备列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ido/getIOSDeviceOptionsPublic [get]
export const getIOSDeviceOptionsPublic = () => {
  return service({
    url: '/ido/getIOSDeviceOptionsPublic',
    method: 'get',
  })
}
