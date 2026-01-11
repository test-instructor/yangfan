import service from '@/utils/request'
import { useUserStore } from '@/pinia/modules/user'
const userStore = useUserStore()
const projectId = userStore.userInfo.projectId
// @Tags PythonCode
// @Summary 创建python 函数
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PythonCode true "创建python 函数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pc/createPythonCode [post]
export const createPythonCode = (data) => {
  data.projectId = projectId
  return service({
    url: '/pc/createPythonCode',
    method: 'post',
    data
  })
}

// @Tags PythonCode
// @Summary 删除python 函数
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PythonCode true "删除python 函数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pc/deletePythonCode [delete]
export const deletePythonCode = (params) => {
  params.projectId = projectId
  return service({
    url: '/pc/deletePythonCode',
    method: 'delete',
    params
  })
}

// @Tags PythonCode
// @Summary 批量删除python 函数
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除python 函数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pc/deletePythonCode [delete]
export const deletePythonCodeByIds = (params) => {
  params.projectId = projectId
  return service({
    url: '/pc/deletePythonCodeByIds',
    method: 'delete',
    params
  })
}

// @Tags PythonCode
// @Summary 更新python 函数
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PythonCode true "更新python 函数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pc/updatePythonCode [put]
export const updatePythonCode = (data) => {
  data.projectId = projectId
  return service({
    url: '/pc/updatePythonCode',
    method: 'put',
    data
  })
}

// @Tags PythonCode
// @Summary 用id查询python 函数
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.PythonCode true "用id查询python 函数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pc/findPythonCode [get]
export const findPythonCode = (params) => {
  params.projectId = projectId
  return service({
    url: '/pc/findPythonCode',
    method: 'get',
    params
  })
}

// @Tags PythonCode
// @Summary 分页获取python 函数列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取python 函数列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pc/getPythonCodeList [get]
export const getPythonCodeList = (params) => {
  params.projectId = projectId
  return service({
    url: '/pc/getPythonCodeList',
    method: 'get',
    params
  })
}

// @Tags PythonCode
// @Summary 不需要鉴权的python 函数接口
// @Accept application/json
// @Produce application/json
// @Param data query platformReq.PythonCodeSearch true "分页获取python 函数列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /pc/getPythonCodePublic [get]
export const getPythonCodePublic = () => {
  return service({
    url: '/pc/getPythonCodePublic',
    method: 'get',
  })
}
