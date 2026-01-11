import service from '@/utils/request'
import { useUserStore } from '@/pinia/modules/user'
const userStore = useUserStore()
const projectId = userStore.userInfo.projectId
// @Tags AutoCase
// @Summary 创建测试用例
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AutoCase true "创建测试用例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ac/createAutoCase [post]
export const createAutoCase = (data) => {
  data.projectId = projectId
  return service({
    url: '/ac/createAutoCase',
    method: 'post',
    data
  })
}

// @Tags AutoCase
// @Summary 删除测试用例
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AutoCase true "删除测试用例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ac/deleteAutoCase [delete]
export const deleteAutoCase = (params) => {
  params.projectId = projectId
  return service({
    url: '/ac/deleteAutoCase',
    method: 'delete',
    params
  })
}

// @Tags AutoCase
// @Summary 批量删除测试用例
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除测试用例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ac/deleteAutoCase [delete]
export const deleteAutoCaseByIds = (params) => {
  params.projectId = projectId
  return service({
    url: '/ac/deleteAutoCaseByIds',
    method: 'delete',
    params
  })
}

// @Tags AutoCase
// @Summary 更新测试用例
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AutoCase true "更新测试用例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ac/updateAutoCase [put]
export const updateAutoCase = (data) => {
  data.projectId = projectId
  return service({
    url: '/ac/updateAutoCase',
    method: 'put',
    data
  })
}

// @Tags AutoCase
// @Summary 用id查询测试用例
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.AutoCase true "用id查询测试用例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ac/findAutoCase [get]
export const findAutoCase = (params) => {
  params.projectId = projectId
  return service({
    url: '/ac/findAutoCase',
    method: 'get',
    params
  })
}

// @Tags AutoCase
// @Summary 分页获取测试用例列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取测试用例列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ac/getAutoCaseList [get]
export const getAutoCaseList = (params) => {
  params.projectId = projectId
  return service({
    url: '/ac/getAutoCaseList',
    method: 'get',
    params
  })
}

// @Tags AutoCase
// @Summary 不需要鉴权的测试用例接口
// @Accept application/json
// @Produce application/json
// @Param data query automationReq.AutoCaseSearch true "分页获取测试用例列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ac/getAutoCasePublic [get]
export const getAutoCasePublic = () => {
  return service({
    url: '/ac/getAutoCasePublic',
    method: 'get',
  })
}

// @Tags AutoCase
// @Summary 添加测试步骤
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body automationReq.AutoCaseStepReq true "添加测试步骤"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /ac/addAutoCaseStep [post]
export const addAutoCaseStep = (data) => {
  data.projectId = projectId
  return service({
    url: '/ac/addAutoCaseStep',
    method: 'post',
    data
  })
}

// @Tags AutoCase
// @Summary 测试步骤排序
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body automationReq.AutoCaseStepSort true "测试步骤排序"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"排序成功"}"
// @Router /ac/sortAutoCaseStep [post]
export const sortAutoCaseStep = (data) => {
  data.projectId = projectId
  return service({
    url: '/ac/sortAutoCaseStep',
    method: 'post',
    data
  })
}

// @Tags AutoCase
// @Summary 删除测试步骤
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query string true "删除测试步骤"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ac/delAutoCaseStep [delete]
export const delAutoCaseStep = (params) => {
  params.projectId = projectId
  return service({
    url: '/ac/delAutoCaseStep',
    method: 'delete',
    params
  })
}

// @Tags AutoCase
// @Summary 获取测试用例步骤
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query string true "获取测试用例步骤"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ac/getAutoCaseSteps [get]
export const getAutoCaseSteps = (params) => {
  params.projectId = projectId
  return service({
    url: '/ac/getAutoCaseSteps',
    method: 'get',
    params
  })
}

// @Tags AutoCase
// @Summary 设置步骤配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body automationReq.SetStepConfigReq true "设置步骤配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /ac/setStepConfig [post]
export const setStepConfig = (data) => {
  data.projectId = projectId
  return service({
    url: '/ac/setStepConfig',
    method: 'put',
    data
  })
}
