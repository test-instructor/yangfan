import service from '@/utils/request'
import { useUserStore } from '@/pinia/modules/user'
const userStore = useUserStore()
const projectId = userStore.userInfo.projectId
// @Tags PythonPackage
// @Summary 创建py 第三方库
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PythonPackage true "创建py 第三方库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pp/createPythonPackage [post]
export const createPythonPackage = (data) => {
  data.projectId = projectId
  return service({
    url: '/pp/createPythonPackage',
    method: 'post',
    data
  })
}

// @Tags PythonPackage
// @Summary 删除py 第三方库
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PythonPackage true "删除py 第三方库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pp/deletePythonPackage [delete]
export const deletePythonPackage = (params) => {
  params.projectId = projectId
  return service({
    url: '/pp/deletePythonPackage',
    method: 'delete',
    params
  })
}

// @Tags PythonPackage
// @Summary 批量删除py 第三方库
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除py 第三方库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pp/deletePythonPackage [delete]
export const deletePythonPackageByIds = (params) => {
  params.projectId = projectId
  return service({
    url: '/pp/deletePythonPackageByIds',
    method: 'delete',
    params
  })
}

// @Tags PythonPackage
// @Summary 更新py 第三方库
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PythonPackage true "更新py 第三方库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pp/updatePythonPackage [put]
export const updatePythonPackage = (data) => {
  data.projectId = projectId
  return service({
    url: '/pp/updatePythonPackage',
    method: 'put',
    data
  })
}

// @Tags PythonPackage
// @Summary 用id查询py 第三方库
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.PythonPackage true "用id查询py 第三方库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pp/findPythonPackage [get]
export const findPythonPackage = (params) => {
  params.projectId = projectId
  return service({
    url: '/pp/findPythonPackage',
    method: 'get',
    params
  })
}


export const findPythonPackageVersion = (params) => {
  params.projectId = projectId
  return service({
    url: '/pp/findPythonPackageVersion',
    method: 'get',
    params
  })
}

// @Tags PythonPackage
// @Summary 分页获取py 第三方库列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取py 第三方库列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pp/getPythonPackageList [get]
export const getPythonPackageList = (params) => {
  params.projectId = projectId
  return service({
    url: '/pp/getPythonPackageList',
    method: 'get',
    params
  })
}

// @Tags PythonPackage
// @Summary 不需要鉴权的py 第三方库接口
// @Accept application/json
// @Produce application/json
// @Param data query platformReq.PythonPackageSearch true "分页获取py 第三方库列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /pp/getPythonPackagePublic [get]
export const getPythonPackagePublic = () => {
  return service({
    url: '/pp/getPythonPackagePublic',
    method: 'get',
  })
}
