import service from '@/utils/request'
import { useUserStore } from '@/pinia/modules/user'
const userStore = useUserStore()
const projectId = userStore.userInfo.projectId
// @Tags AutoStep
// @Summary 创建自动化步骤
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AutoStep true "创建自动化步骤"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /as/createAutoStep [post]
export const createAutoStep = (data) => {
  data.projectId = projectId
  return service({
    url: '/as/createAutoStep',
    method: 'post',
    data
  })
}

// @Tags AutoStep
// @Summary 删除自动化步骤
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AutoStep true "删除自动化步骤"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /as/deleteAutoStep [delete]
export const deleteAutoStep = (params) => {
  params.projectId = projectId
  return service({
    url: '/as/deleteAutoStep',
    method: 'delete',
    params
  })
}

// @Tags AutoStep
// @Summary 批量删除自动化步骤
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除自动化步骤"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /as/deleteAutoStep [delete]
export const deleteAutoStepByIds = (params) => {
  params.projectId = projectId
  return service({
    url: '/as/deleteAutoStepByIds',
    method: 'delete',
    params
  })
}

// @Tags AutoStep
// @Summary 更新自动化步骤
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AutoStep true "更新自动化步骤"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /as/updateAutoStep [put]
export const updateAutoStep = (data) => {
  data.projectId = projectId
  return service({
    url: '/as/updateAutoStep',
    method: 'put',
    data
  })
}

// @Tags AutoStep
// @Summary 用id查询自动化步骤
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.AutoStep true "用id查询自动化步骤"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /as/findAutoStep [get]
export const findAutoStep = (params) => {
  params.projectId = projectId
  return service({
    url: '/as/findAutoStep',
    method: 'get',
    params
  })
}

// @Tags AutoStep
// @Summary 分页获取自动化步骤列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取自动化步骤列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /as/getAutoStepList [get]
export const getAutoStepList = (params) => {
  params.projectId = projectId
  return service({
    url: '/as/getAutoStepList',
    method: 'get',
    params
  })
}

// @Tags AutoStep
// @Summary 不需要鉴权的自动化步骤接口
// @Accept application/json
// @Produce application/json
// @Param data query automationReq.AutoStepSearch true "分页获取自动化步骤列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /as/getAutoStepPublic [get]
export const getAutoStepPublic = () => {
  return service({
    url: '/as/getAutoStepPublic',
    method: 'get',
  })
}
