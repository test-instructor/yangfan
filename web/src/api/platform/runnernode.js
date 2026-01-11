import service from '@/utils/request'
import { useUserStore } from '@/pinia/modules/user'
const userStore = useUserStore()
const projectId = userStore.userInfo.projectId
// @Tags RunnerNode
// @Summary 创建节点
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.RunnerNode true "创建节点"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /rn/createRunnerNode [post]
export const createRunnerNode = (data) => {
  data.projectId = projectId
  return service({
    url: '/rn/createRunnerNode',
    method: 'post',
    data
  })
}

// @Tags RunnerNode
// @Summary 删除节点
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.RunnerNode true "删除节点"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rn/deleteRunnerNode [delete]
export const deleteRunnerNode = (params) => {
  params.projectId = projectId
  return service({
    url: '/rn/deleteRunnerNode',
    method: 'delete',
    params
  })
}

// @Tags RunnerNode
// @Summary 批量删除节点
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除节点"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rn/deleteRunnerNode [delete]
export const deleteRunnerNodeByIds = (params) => {
  params.projectId = projectId
  return service({
    url: '/rn/deleteRunnerNodeByIds',
    method: 'delete',
    params
  })
}

// @Tags RunnerNode
// @Summary 更新节点
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.RunnerNode true "更新节点"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /rn/updateRunnerNode [put]
export const updateRunnerNode = (data) => {
  data.projectId = projectId
  return service({
    url: '/rn/updateRunnerNode',
    method: 'put',
    data
  })
}

// @Tags RunnerNode
// @Summary 用id查询节点
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.RunnerNode true "用id查询节点"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /rn/findRunnerNode [get]
export const findRunnerNode = (params) => {
  params.projectId = projectId
  return service({
    url: '/rn/findRunnerNode',
    method: 'get',
    params
  })
}

// @Tags RunnerNode
// @Summary 分页获取节点列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取节点列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rn/getRunnerNodeList [get]
export const getRunnerNodeList = (params) => {
  params.projectId = projectId
  return service({
    url: '/rn/getRunnerNodeList',
    method: 'get',
    params
  })
}

// @Tags RunnerNode
// @Summary 不需要鉴权的节点接口
// @Accept application/json
// @Produce application/json
// @Param data query platformReq.RunnerNodeSearch true "分页获取节点列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /rn/getRunnerNodePublic [get]
export const getRunnerNodePublic = () => {
  return service({
    url: '/rn/getRunnerNodePublic',
    method: 'get',
  })
}
