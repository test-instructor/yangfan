import service from '@/utils/request'
import { useUserStore } from '@/pinia/modules/user'
const userStore = useUserStore()
const projectId = userStore.userInfo.projectId
// @Tags BrowserDeviceOptions
// @Summary 创建浏览器设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.BrowserDeviceOptions true "创建浏览器设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /bdo/createBrowserDeviceOptions [post]
export const createBrowserDeviceOptions = (data) => {
  data.projectId = projectId
  return service({
    url: '/bdo/createBrowserDeviceOptions',
    method: 'post',
    data
  })
}

// @Tags BrowserDeviceOptions
// @Summary 删除浏览器设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.BrowserDeviceOptions true "删除浏览器设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bdo/deleteBrowserDeviceOptions [delete]
export const deleteBrowserDeviceOptions = (params) => {
  params.projectId = projectId
  return service({
    url: '/bdo/deleteBrowserDeviceOptions',
    method: 'delete',
    params
  })
}

// @Tags BrowserDeviceOptions
// @Summary 批量删除浏览器设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除浏览器设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bdo/deleteBrowserDeviceOptions [delete]
export const deleteBrowserDeviceOptionsByIds = (params) => {
  params.projectId = projectId
  return service({
    url: '/bdo/deleteBrowserDeviceOptionsByIds',
    method: 'delete',
    params
  })
}

// @Tags BrowserDeviceOptions
// @Summary 更新浏览器设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.BrowserDeviceOptions true "更新浏览器设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bdo/updateBrowserDeviceOptions [put]
export const updateBrowserDeviceOptions = (data) => {
  data.projectId = projectId
  return service({
    url: '/bdo/updateBrowserDeviceOptions',
    method: 'put',
    data
  })
}

// @Tags BrowserDeviceOptions
// @Summary 用id查询浏览器设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.BrowserDeviceOptions true "用id查询浏览器设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bdo/findBrowserDeviceOptions [get]
export const findBrowserDeviceOptions = (params) => {
  params.projectId = projectId
  return service({
    url: '/bdo/findBrowserDeviceOptions',
    method: 'get',
    params
  })
}

// @Tags BrowserDeviceOptions
// @Summary 分页获取浏览器设备列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取浏览器设备列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bdo/getBrowserDeviceOptionsList [get]
export const getBrowserDeviceOptionsList = (params) => {
  params.projectId = projectId
  return service({
    url: '/bdo/getBrowserDeviceOptionsList',
    method: 'get',
    params
  })
}

// @Tags BrowserDeviceOptions
// @Summary 不需要鉴权的浏览器设备接口
// @Accept application/json
// @Produce application/json
// @Param data query platformReq.BrowserDeviceOptionsSearch true "分页获取浏览器设备列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /bdo/getBrowserDeviceOptionsPublic [get]
export const getBrowserDeviceOptionsPublic = () => {
  return service({
    url: '/bdo/getBrowserDeviceOptionsPublic',
    method: 'get',
  })
}
