import service from '@/utils/request'
import { useUserStore } from '@/pinia/modules/user'
const userStore = useUserStore()
const projectId = userStore.userInfo.projectId
// @Tags LLMModelConfig
// @Summary 创建大语言模型配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.LLMModelConfig true "创建大语言模型配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /llmconfig/createLLMModelConfig [post]
export const createLLMModelConfig = (data) => {
  data.projectId = projectId
  return service({
    url: '/llmconfig/createLLMModelConfig',
    method: 'post',
    data
  })
}

// @Tags LLMModelConfig
// @Summary 删除大语言模型配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.LLMModelConfig true "删除大语言模型配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /llmconfig/deleteLLMModelConfig [delete]
export const deleteLLMModelConfig = (params) => {
  params.projectId = projectId
  return service({
    url: '/llmconfig/deleteLLMModelConfig',
    method: 'delete',
    params
  })
}

// @Tags LLMModelConfig
// @Summary 批量删除大语言模型配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除大语言模型配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /llmconfig/deleteLLMModelConfig [delete]
export const deleteLLMModelConfigByIds = (params) => {
  params.projectId = projectId
  return service({
    url: '/llmconfig/deleteLLMModelConfigByIds',
    method: 'delete',
    params
  })
}

// @Tags LLMModelConfig
// @Summary 更新大语言模型配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.LLMModelConfig true "更新大语言模型配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /llmconfig/updateLLMModelConfig [put]
export const updateLLMModelConfig = (data) => {
  data.projectId = projectId
  return service({
    url: '/llmconfig/updateLLMModelConfig',
    method: 'put',
    data
  })
}

// @Tags LLMModelConfig
// @Summary 用id查询大语言模型配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.LLMModelConfig true "用id查询大语言模型配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /llmconfig/findLLMModelConfig [get]
export const findLLMModelConfig = (params) => {
  params.projectId = projectId
  return service({
    url: '/llmconfig/findLLMModelConfig',
    method: 'get',
    params
  })
}

// @Tags LLMModelConfig
// @Summary 分页获取大语言模型配置列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取大语言模型配置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /llmconfig/getLLMModelConfigList [get]
export const getLLMModelConfigList = (params) => {
  params.projectId = projectId
  return service({
    url: '/llmconfig/getLLMModelConfigList',
    method: 'get',
    params
  })
}

// @Tags LLMModelConfig
// @Summary 不需要鉴权的大语言模型配置接口
// @Accept application/json
// @Produce application/json
// @Param data query platformReq.LLMModelConfigSearch true "分页获取大语言模型配置列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /llmconfig/getLLMModelConfigPublic [get]
export const getLLMModelConfigPublic = () => {
  return service({
    url: '/llmconfig/getLLMModelConfigPublic',
    method: 'get',
  })
}
