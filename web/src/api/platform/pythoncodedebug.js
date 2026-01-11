import service from '@/utils/request'
import { useUserStore } from '@/pinia/modules/user'
const userStore = useUserStore()
const projectId = userStore.userInfo.projectId
// @Tags PythonCodeDebug
// @Summary 创建调试信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PythonCodeDebug true "创建调试信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pcd/createPythonCodeDebug [post]
export const createPythonCodeDebug = (data) => {
  data.projectId = projectId
  return service({
    url: '/pcd/createPythonCodeDebug',
    method: 'post',
    data
  })
}

// @Tags PythonCodeDebug
// @Summary 删除调试信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PythonCodeDebug true "删除调试信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pcd/deletePythonCodeDebug [delete]
export const deletePythonCodeDebug = (params) => {
  params.projectId = projectId
  return service({
    url: '/pcd/deletePythonCodeDebug',
    method: 'delete',
    params
  })
}

// @Tags PythonCodeDebug
// @Summary 批量删除调试信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除调试信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pcd/deletePythonCodeDebug [delete]
export const deletePythonCodeDebugByIds = (params) => {
  params.projectId = projectId
  return service({
    url: '/pcd/deletePythonCodeDebugByIds',
    method: 'delete',
    params
  })
}

// @Tags PythonCodeDebug
// @Summary 更新调试信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PythonCodeDebug true "更新调试信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pcd/updatePythonCodeDebug [put]
export const updatePythonCodeDebug = (data) => {
  data.projectId = projectId
  return service({
    url: '/pcd/updatePythonCodeDebug',
    method: 'put',
    data
  })
}

// @Tags PythonCodeDebug
// @Summary 用id查询调试信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.PythonCodeDebug true "用id查询调试信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pcd/findPythonCodeDebug [get]
export const findPythonCodeDebug = (params) => {
  params.projectId = projectId
  return service({
    url: '/pcd/findPythonCodeDebug',
    method: 'get',
    params
  })
}

// @Tags PythonCodeDebug
// @Summary 分页获取调试信息列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取调试信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pcd/getPythonCodeDebugList [get]
export const getPythonCodeDebugList = (params) => {
  params.projectId = projectId
  return service({
    url: '/pcd/getPythonCodeDebugList',
    method: 'get',
    params
  })
}

// @Tags PythonCodeDebug
// @Summary 不需要鉴权的调试信息接口
// @Accept application/json
// @Produce application/json
// @Param data query platformReq.PythonCodeDebugSearch true "分页获取调试信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /pcd/getPythonCodeDebugPublic [get]
export const getPythonCodeDebugPublic = () => {
  return service({
    url: '/pcd/getPythonCodeDebugPublic',
    method: 'get',
  })
}
