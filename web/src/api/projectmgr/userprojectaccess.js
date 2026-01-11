import service from '@/utils/request'
// @Tags UserProjectAccess
// @Summary 创建项目成员与权限
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserProjectAccess true "创建项目成员与权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /upa/createUserProjectAccess [post]
export const createUserProjectAccess = (data) => {
  return service({
    url: '/upa/createUserProjectAccess',
    method: 'post',
    data
  })
}

// @Tags UserProjectAccess
// @Summary 删除项目成员与权限
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserProjectAccess true "删除项目成员与权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /upa/deleteUserProjectAccess [delete]
export const deleteUserProjectAccess = (params) => {
  return service({
    url: '/upa/deleteUserProjectAccess',
    method: 'delete',
    params
  })
}

// @Tags UserProjectAccess
// @Summary 批量删除项目成员与权限
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除项目成员与权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /upa/deleteUserProjectAccess [delete]
export const deleteUserProjectAccessByIds = (params) => {
  return service({
    url: '/upa/deleteUserProjectAccessByIds',
    method: 'delete',
    params
  })
}

// @Tags UserProjectAccess
// @Summary 更新项目成员与权限
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserProjectAccess true "更新项目成员与权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /upa/updateUserProjectAccess [put]
export const updateUserProjectAccess = (data) => {
  return service({
    url: '/upa/updateUserProjectAccess',
    method: 'put',
    data
  })
}

// @Tags UserProjectAccess
// @Summary 用id查询项目成员与权限
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.UserProjectAccess true "用id查询项目成员与权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /upa/findUserProjectAccess [get]
export const findUserProjectAccess = (params) => {
  return service({
    url: '/upa/findUserProjectAccess',
    method: 'get',
    params
  })
}

// @Tags UserProjectAccess
// @Summary 分页获取项目成员与权限列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取项目成员与权限列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /upa/getUserProjectAccessList [get]
export const getUserProjectAccessList = (params) => {
  return service({
    url: '/upa/getUserProjectAccessList',
    method: 'get',
    params
  })
}

// @Tags UserProjectAccess
// @Summary 不需要鉴权的项目成员与权限接口
// @Accept application/json
// @Produce application/json
// @Param data query projectmgrReq.UserProjectAccessSearch true "分页获取项目成员与权限列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /upa/getUserProjectAccessPublic [get]
export const getUserProjectAccessPublic = () => {
  return service({
    url: '/upa/getUserProjectAccessPublic',
    method: 'get',
  })
}
