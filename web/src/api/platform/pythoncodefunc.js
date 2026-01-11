import service from '@/utils/request'
import { useUserStore } from '@/pinia/modules/user'
const userStore = useUserStore()
const projectId = userStore.userInfo.projectId
// @Tags PythonCodeFunc
// @Summary 创建python函数详情
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PythonCodeFunc true "创建python函数详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pcf/createPythonCodeFunc [post]
export const createPythonCodeFunc = (data) => {
  data.projectId = projectId
  return service({
    url: '/pcf/createPythonCodeFunc',
    method: 'post',
    data
  })
}

// @Tags PythonCodeFunc
// @Summary 删除python函数详情
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PythonCodeFunc true "删除python函数详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pcf/deletePythonCodeFunc [delete]
export const deletePythonCodeFunc = (params) => {
  params.projectId = projectId
  return service({
    url: '/pcf/deletePythonCodeFunc',
    method: 'delete',
    params
  })
}

// @Tags PythonCodeFunc
// @Summary 批量删除python函数详情
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除python函数详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pcf/deletePythonCodeFunc [delete]
export const deletePythonCodeFuncByIds = (params) => {
  params.projectId = projectId
  return service({
    url: '/pcf/deletePythonCodeFuncByIds',
    method: 'delete',
    params
  })
}

// @Tags PythonCodeFunc
// @Summary 更新python函数详情
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PythonCodeFunc true "更新python函数详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pcf/updatePythonCodeFunc [put]
export const updatePythonCodeFunc = (data) => {
  data.projectId = projectId
  return service({
    url: '/pcf/updatePythonCodeFunc',
    method: 'put',
    data
  })
}

// @Tags PythonCodeFunc
// @Summary 用id查询python函数详情
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.PythonCodeFunc true "用id查询python函数详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pcf/findPythonCodeFunc [get]
export const findPythonCodeFunc = (params) => {
  params.projectId = projectId
  return service({
    url: '/pcf/findPythonCodeFunc',
    method: 'get',
    params
  })
}

// @Tags PythonCodeFunc
// @Summary 分页获取python函数详情列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取python函数详情列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pcf/getPythonCodeFuncList [get]
export const getPythonCodeFuncList = (params) => {
  params.projectId = projectId
  return service({
    url: '/pcf/getPythonCodeFuncList',
    method: 'get',
    params
  })
}

// @Tags PythonCodeFunc
// @Summary 不需要鉴权的python函数详情接口
// @Accept application/json
// @Produce application/json
// @Param data query platformReq.PythonCodeFuncSearch true "分页获取python函数详情列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /pcf/getPythonCodeFuncPublic [get]
export const getPythonCodeFuncPublic = () => {
  return service({
    url: '/pcf/getPythonCodeFuncPublic',
    method: 'get',
  })
}
