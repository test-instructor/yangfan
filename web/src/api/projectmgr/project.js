import service from '@/utils/request'
// @Tags Project
// @Summary 创建项目配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Project true "创建项目配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pj/createProject [post]
export const createProject = (data) => {
  return service({
    url: '/pj/createProject',
    method: 'post',
    data
  })
}

// @Tags Project
// @Summary 删除项目配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Project true "删除项目配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pj/deleteProject [delete]
export const deleteProject = (params) => {
  return service({
    url: '/pj/deleteProject',
    method: 'delete',
    params
  })
}

// @Tags Project
// @Summary 批量删除项目配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除项目配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pj/deleteProject [delete]
export const deleteProjectByIds = (params) => {
  return service({
    url: '/pj/deleteProjectByIds',
    method: 'delete',
    params
  })
}

// @Tags Project
// @Summary 更新项目配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Project true "更新项目配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pj/updateProject [put]
export const updateProject = (data) => {
  return service({
    url: '/pj/updateProject',
    method: 'put',
    data
  })
}

export const resetProjectAuth = (data) => {
  return service({
    url: '/pj/resetProjectAuth',
    method: 'put',
    data
  })
}

// @Tags Project
// @Summary 用id查询项目配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Project true "用id查询项目配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pj/findProject [get]
export const findProject = (params) => {
  return service({
    url: '/pj/findProject',
    method: 'get',
    params
  })
}

// @Tags Project
// @Summary 分页获取项目配置列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取项目配置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pj/getProjectList [get]
export const getProjectList = (params) => {
  return service({
    url: '/pj/getProjectList',
    method: 'get',
    params
  })
}

// @Tags Project
// @Summary 不需要鉴权的项目配置接口
// @Accept application/json
// @Produce application/json
// @Param data query projectmgrReq.ProjectSearch true "分页获取项目配置列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /pj/getProjectPublic [get]
export const getProjectPublic = () => {
  return service({
    url: '/pj/getProjectPublic',
    method: 'get',
  })
}
