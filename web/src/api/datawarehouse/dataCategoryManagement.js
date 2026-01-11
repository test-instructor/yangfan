import service from '@/utils/request'
import { useUserStore } from '@/pinia/modules/user'
const userStore = useUserStore()
const projectId = userStore.userInfo.projectId
// @Tags DataCategoryManagement
// @Summary 创建数据分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DataCategoryManagement true "创建数据分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dcm/createDataCategoryManagement [post]
export const createDataCategoryManagement = (data) => {
  data.projectId = projectId
  return service({
    url: '/dcm/createDataCategoryManagement',
    method: 'post',
    data
  })
}

// @Tags DataCategoryManagement
// @Summary 删除数据分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DataCategoryManagement true "删除数据分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dcm/deleteDataCategoryManagement [delete]
export const deleteDataCategoryManagement = (params) => {
  params.projectId = projectId
  return service({
    url: '/dcm/deleteDataCategoryManagement',
    method: 'delete',
    params
  })
}

// @Tags DataCategoryManagement
// @Summary 批量删除数据分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dcm/deleteDataCategoryManagement [delete]
export const deleteDataCategoryManagementByIds = (params) => {
  params.projectId = projectId
  return service({
    url: '/dcm/deleteDataCategoryManagementByIds',
    method: 'delete',
    params
  })
}

// @Tags DataCategoryManagement
// @Summary 更新数据分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DataCategoryManagement true "更新数据分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dcm/updateDataCategoryManagement [put]
export const updateDataCategoryManagement = (data) => {
  data.projectId = projectId
  return service({
    url: '/dcm/updateDataCategoryManagement',
    method: 'put',
    data
  })
}

// @Tags DataCategoryManagement
// @Summary 用id查询数据分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.DataCategoryManagement true "用id查询数据分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dcm/findDataCategoryManagement [get]
export const findDataCategoryManagement = (params) => {
  params.projectId = projectId
  return service({
    url: '/dcm/findDataCategoryManagement',
    method: 'get',
    params
  })
}

// @Tags DataCategoryManagement
// @Summary 分页获取数据分类列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取数据分类列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dcm/getDataCategoryManagementList [get]
export const getDataCategoryManagementList = (params) => {
  params.projectId = projectId
  return service({
    url: '/dcm/getDataCategoryManagementList',
    method: 'get',
    params
  })
}

// @Tags DataCategoryManagement
// @Summary 获取数据分类类型列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}"
// @Router /dcm/getDataCategoryTypeList [get]
export const getDataCategoryTypeList = () => {
  const params = {}
  params.projectId = projectId
  return service({
    url: '/dcm/getDataCategoryTypeList',
    method: 'get',
    params
  })
}
