import service from '@/utils/request'
import { useUserStore } from '@/pinia/modules/user'
const userStore = useUserStore()
const projectId = userStore.userInfo.projectId
// @Tags AndroidDeviceOptions
// @Summary 创建安卓设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AndroidDeviceOptions true "创建安卓设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ado/createAndroidDeviceOptions [post]
export const createAndroidDeviceOptions = (data) => {
  data.projectId = projectId
  return service({
    url: '/ado/createAndroidDeviceOptions',
    method: 'post',
    data
  })
}

// @Tags AndroidDeviceOptions
// @Summary 删除安卓设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AndroidDeviceOptions true "删除安卓设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ado/deleteAndroidDeviceOptions [delete]
export const deleteAndroidDeviceOptions = (params) => {
  params.projectId = projectId
  return service({
    url: '/ado/deleteAndroidDeviceOptions',
    method: 'delete',
    params
  })
}

// @Tags AndroidDeviceOptions
// @Summary 批量删除安卓设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除安卓设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ado/deleteAndroidDeviceOptions [delete]
export const deleteAndroidDeviceOptionsByIds = (params) => {
  params.projectId = projectId
  return service({
    url: '/ado/deleteAndroidDeviceOptionsByIds',
    method: 'delete',
    params
  })
}

// @Tags AndroidDeviceOptions
// @Summary 更新安卓设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AndroidDeviceOptions true "更新安卓设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ado/updateAndroidDeviceOptions [put]
export const updateAndroidDeviceOptions = (data) => {
  data.projectId = projectId
  return service({
    url: '/ado/updateAndroidDeviceOptions',
    method: 'put',
    data
  })
}

// @Tags AndroidDeviceOptions
// @Summary 用id查询安卓设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.AndroidDeviceOptions true "用id查询安卓设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ado/findAndroidDeviceOptions [get]
export const findAndroidDeviceOptions = (params) => {
  params.projectId = projectId
  return service({
    url: '/ado/findAndroidDeviceOptions',
    method: 'get',
    params
  })
}

// @Tags AndroidDeviceOptions
// @Summary 分页获取安卓设备列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取安卓设备列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ado/getAndroidDeviceOptionsList [get]
export const getAndroidDeviceOptionsList = (params) => {
  params.projectId = projectId
  return service({
    url: '/ado/getAndroidDeviceOptionsList',
    method: 'get',
    params
  })
}

// @Tags AndroidDeviceOptions
// @Summary 不需要鉴权的安卓设备接口
// @Accept application/json
// @Produce application/json
// @Param data query platformReq.AndroidDeviceOptionsSearch true "分页获取安卓设备列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ado/getAndroidDeviceOptionsPublic [get]
export const getAndroidDeviceOptionsPublic = () => {
  return service({
    url: '/ado/getAndroidDeviceOptionsPublic',
    method: 'get',
  })
}
