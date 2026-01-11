import service from '@/utils/request'
import { useUserStore } from '@/pinia/modules/user'
const userStore = useUserStore()
const projectId = userStore.userInfo.projectId
// @Tags CategoryMenu
// @Summary 创建自动化菜单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.CategoryMenu true "创建自动化菜单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cm/createCategoryMenu [post]
export const createCategoryMenu = (data) => {
  data.projectId = projectId
  return service({
    url: '/cm/createCategoryMenu',
    method: 'post',
    data
  })
}

// @Tags CategoryMenu
// @Summary 删除自动化菜单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.CategoryMenu true "删除自动化菜单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cm/deleteCategoryMenu [delete]
export const deleteCategoryMenu = (params) => {
  params.projectId = projectId
  return service({
    url: '/cm/deleteCategoryMenu',
    method: 'delete',
    params
  })
}

// @Tags CategoryMenu
// @Summary 批量删除自动化菜单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除自动化菜单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cm/deleteCategoryMenu [delete]
export const deleteCategoryMenuByIds = (params) => {
  params.projectId = projectId
  return service({
    url: '/cm/deleteCategoryMenuByIds',
    method: 'delete',
    params
  })
}

// @Tags CategoryMenu
// @Summary 更新自动化菜单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.CategoryMenu true "更新自动化菜单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cm/updateCategoryMenu [put]
export const updateCategoryMenu = (data) => {
  data.projectId = projectId
  return service({
    url: '/cm/updateCategoryMenu',
    method: 'put',
    data
  })
}

// @Tags CategoryMenu
// @Summary 用id查询自动化菜单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.CategoryMenu true "用id查询自动化菜单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cm/findCategoryMenu [get]
export const findCategoryMenu = (params) => {
  params.projectId = projectId
  return service({
    url: '/cm/findCategoryMenu',
    method: 'get',
    params
  })
}

// @Tags CategoryMenu
// @Summary 分页获取自动化菜单列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取自动化菜单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cm/getCategoryMenuList [get]
export const getCategoryMenuList = (params) => {
  params.projectId = projectId
  return service({
    url: '/cm/getCategoryMenuList',
    method: 'get',
    params
  })
}

// @Tags CategoryMenu
// @Summary 不需要鉴权的自动化菜单接口
// @Accept application/json
// @Produce application/json
// @Param data query platformReq.CategoryMenuSearch true "分页获取自动化菜单列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cm/getCategoryMenuPublic [get]
export const getCategoryMenuPublic = () => {
  return service({
    url: '/cm/getCategoryMenuPublic',
    method: 'get',
  })
}
