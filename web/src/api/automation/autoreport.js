import service from '@/utils/request'
import { useUserStore } from '@/pinia/modules/user'
const userStore = useUserStore()
const projectId = userStore.userInfo.projectId
// @Tags AutoReport
// @Summary 创建自动报告
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AutoReport true "创建自动报告"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ar/createAutoReport [post]
export const createAutoReport = (data) => {
  data.projectId = projectId
  return service({
    url: '/ar/createAutoReport',
    method: 'post',
    data
  })
}

// @Tags AutoReport
// @Summary 删除自动报告
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AutoReport true "删除自动报告"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ar/deleteAutoReport [delete]
export const deleteAutoReport = (params) => {
  params.projectId = projectId
  return service({
    url: '/ar/deleteAutoReport',
    method: 'delete',
    params
  })
}

// @Tags AutoReport
// @Summary 批量删除自动报告
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除自动报告"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ar/deleteAutoReport [delete]
export const deleteAutoReportByIds = (params) => {
  params.projectId = projectId
  return service({
    url: '/ar/deleteAutoReportByIds',
    method: 'delete',
    params
  })
}

// @Tags AutoReport
// @Summary 更新自动报告
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AutoReport true "更新自动报告"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ar/updateAutoReport [put]
export const updateAutoReport = (data) => {
  data.projectId = projectId
  return service({
    url: '/ar/updateAutoReport',
    method: 'put',
    data
  })
}

// @Tags AutoReport
// @Summary 用id查询自动报告
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.AutoReport true "用id查询自动报告"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ar/findAutoReport [get]
export const findAutoReport = (params) => {
  params.projectId = projectId
  return service({
    url: '/ar/findAutoReport',
    method: 'get',
    params
  })
}

// @Tags AutoReport
// @Summary 分页获取自动报告列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取自动报告列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ar/getAutoReportList [get]
export const getAutoReportList = (params) => {
  params.projectId = projectId
  return service({
    url: '/ar/getAutoReportList',
    method: 'get',
    params
  })
}

// @Tags AutoReport
// @Summary 不需要鉴权的自动报告接口
// @Accept application/json
// @Produce application/json
// @Param data query automationReq.AutoReportSearch true "分页获取自动报告列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ar/getAutoReportPublic [get]
export const getAutoReportPublic = () => {
  return service({
    url: '/ar/getAutoReportPublic',
    method: 'get',
  })
}
