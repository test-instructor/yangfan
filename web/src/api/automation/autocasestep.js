import service from '@/utils/request'
import { useUserStore } from '@/pinia/modules/user'
const userStore = useUserStore()
const projectId = userStore.userInfo.projectId
// @Tags AutoCaseStep
// @Summary 创建测试步骤
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AutoCaseStep true "创建测试步骤"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /acs/createAutoCaseStep [post]
export const createAutoCaseStep = (data) => {
  data.projectId = projectId
  return service({
    url: '/acs/createAutoCaseStep',
    method: 'post',
    data
  })
}

export const addAutoCaseStepApi = (data) => {
  data.projectId = projectId
  return service({
    url: '/acs/addAutoCaseStepApi',
    method: 'post',
    data
  })
}

export const sortAutoCaseStepApi = (data) => {
  data.projectId = projectId
  return service({
    url: '/acs/sortAutoCaseStepApi',
    method: 'post',
    data
  })
}


// @Tags AutoCaseStep
// @Summary 删除测试步骤
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AutoCaseStep true "删除测试步骤"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /acs/deleteAutoCaseStep [delete]
export const deleteAutoCaseStep = (params) => {
  params.projectId = projectId
  return service({
    url: '/acs/deleteAutoCaseStep',
    method: 'delete',
    params
  })
}

export const deleteAutoCaseStepApi = (params) => {
  params.projectId = projectId
  return service({
    url: '/acs/deleteAutoCaseStepApi',
    method: 'delete',
    params
  })
}


// @Tags AutoCaseStep
// @Summary 批量删除测试步骤
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除测试步骤"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /acs/deleteAutoCaseStep [delete]
export const deleteAutoCaseStepByIds = (params) => {
  params.projectId = projectId
  return service({
    url: '/acs/deleteAutoCaseStepByIds',
    method: 'delete',
    params
  })
}

// @Tags AutoCaseStep
// @Summary 更新测试步骤
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AutoCaseStep true "更新测试步骤"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /acs/updateAutoCaseStep [put]
export const updateAutoCaseStep = (data) => {
  data.projectId = projectId
  return service({
    url: '/acs/updateAutoCaseStep',
    method: 'put',
    data
  })
}

// @Tags AutoCaseStep
// @Summary 用id查询测试步骤
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.AutoCaseStep true "用id查询测试步骤"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /acs/findAutoCaseStep [get]
export const findAutoCaseStep = (params) => {
  params.projectId = projectId
  return service({
    url: '/acs/findAutoCaseStep',
    method: 'get',
    params
  })
}

export const findAutoCaseStepApi = (params) => {
  params.projectId = projectId
  return service({
    url: '/acs/findAutoCaseStepApi',
    method: 'get',
    params
  })
}

// @Tags AutoCaseStep
// @Summary 分页获取测试步骤列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取测试步骤列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /acs/getAutoCaseStepList [get]
export const getAutoCaseStepList = (params) => {
  params.projectId = projectId
  return service({
    url: '/acs/getAutoCaseStepList',
    method: 'get',
    params
  })
}

// @Tags AutoCaseStep
// @Summary 不需要鉴权的测试步骤接口
// @Accept application/json
// @Produce application/json
// @Param data query automationReq.AutoCaseStepSearch true "分页获取测试步骤列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /acs/getAutoCaseStepPublic [get]
export const getAutoCaseStepPublic = () => {
  return service({
    url: '/acs/getAutoCaseStepPublic',
    method: 'get',
  })
}
