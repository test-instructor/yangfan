package interfacecase

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/model/common/request"
	"github.com/test-instructor/cheetah/server/model/common/response"
	"github.com/test-instructor/cheetah/server/model/interfacecase"
	"github.com/test-instructor/cheetah/server/model/system"
	"github.com/test-instructor/cheetah/server/service"
	"go.uber.org/zap"
)

type ApiMenuApi struct {
}

var apicaseService = service.ServiceGroupApp.InterfacecaseServiceGroup.ApiMenuService

// CreateApiMenu 创建ApiMenu
// @Tags ApiMenu
// @Summary 创建ApiMenu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body interfacecase.ApiMenu true "创建ApiMenu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apicase/createApiMenu [post]
func getMenuList(c *gin.Context) {
	project, _ := c.Get("project")
	menuType := c.Request.FormValue("menutype")
	if treeList, err := apicaseService.GetMenu(0, menuType, project.(system.Project)); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: treeList,
		}, "获取成功", c)
	}
}

func (apicaseApi *ApiMenuApi) CreateApiMenu(c *gin.Context) {
	var apicase interfacecase.ApiMenu
	_ = c.ShouldBindJSON(&apicase)
	apicase.MenuType = c.Request.FormValue("menutype")
	project, _ := c.Get("project")
	apicase.Project = project.(system.Project)
	if err := apicaseService.CreateApiMenu(apicase); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		getMenuList(c)
	}
}

// DeleteApiMenu 删除ApiMenu
// @Tags ApiMenu
// @Summary 删除ApiMenu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body interfacecase.ApiMenu true "删除ApiMenu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /apicase/deleteApiMenu [delete]
func (apicaseApi *ApiMenuApi) DeleteApiMenu(c *gin.Context) {
	var apicase interfacecase.ApiMenu
	_ = c.ShouldBindJSON(&apicase)
	project, _ := c.Get("project")
	apicase.Project = project.(system.Project)

	if err := apicaseService.GetApiMenuInterface(apicase); err != nil {
		global.GVA_LOG.Error("该目录下有api，无法进行删除!", zap.Error(err))
		response.FailWithMessage("该目录下有api，无法进行删除!", c)
	} else if err := apicaseService.DeleteApiMenu(apicase); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		getMenuList(c)
	}
}

// DeleteApiMenuByIds 批量删除ApiMenu
// @Tags ApiMenu
// @Summary 批量删除ApiMenu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ApiMenu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /apicase/deleteApiMenuByIds [delete]
func (apicaseApi *ApiMenuApi) DeleteApiMenuByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := apicaseService.DeleteApiMenuByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		getMenuList(c)
	}
}

// UpdateApiMenu 更新ApiMenu
// @Tags ApiMenu
// @Summary 更新ApiMenu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body interfacecase.ApiMenu true "更新ApiMenu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /apicase/updateApiMenu [put]
func (apicaseApi *ApiMenuApi) UpdateApiMenu(c *gin.Context) {
	var apicase interfacecase.ApiMenu
	_ = c.ShouldBindJSON(&apicase)
	project, _ := c.Get("project")
	apicase.Project = project.(system.Project)
	if err := apicaseService.UpdateApiMenu(apicase); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		getMenuList(c)
	}
}

// FindApiMenu 用id查询ApiMenu
// @Tags ApiMenu
// @Summary 用id查询ApiMenu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query interfacecase.ApiMenu true "用id查询ApiMenu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /apicase/findApiMenu [get]
func (apicaseApi *ApiMenuApi) FindApiMenu(c *gin.Context) {
	var apicase interfacecase.ApiMenu
	_ = c.ShouldBindQuery(&apicase)
	project, _ := c.Get("project")
	apicase.Project = project.(system.Project)
	if err, reapicase := apicaseService.GetApiMenu(apicase.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reapicase": reapicase}, c)
	}
}

// GetApiMenuList 分页获取ApiMenu列表
// @Tags ApiMenu
// @Summary 分页获取ApiMenu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apicase/getApiMenuList [get]
func (apicaseApi *ApiMenuApi) GetApiMenuList(c *gin.Context) {

	/*var interfaceTemp InterfaceTemplate
	interfaceTemp = InterfaceTemplate{
		Name: "接口名称",
		Request: ApiRequest{
			Agreement:"HTTP",
			Method: "POST",
			Url: "/get",
			Params: []ApiKeyToValue{{Key: "Paramskey", Value: "123"}},
			Headers: []ApiKeyToValue{{Key: "Headerskey", Value: "123"}},
			Json: []ApiKeyToValue{{Key: "Jsonkey", Value: "123"}},
			Data: []ApiKeyToValue{{Key: "Datakey", Value: "123"}},
		},
		Variables: []ApiKeyToValue{{Key: "Variableskey", Value: "123"}},
		Extract: []ApiKeyToValue{{Key: "Extractkey", Value: "123"}},
		Validate: []ApiKeyToValue{{Key: "Validatekey", Value: "123"}},
		Project: system.Project{GVA_MODEL:global.GVA_MODEL{ID:1,
		}},

	}
	fmt.Println(interfaceTemp)
	global.GVA_DB.Create(&interfaceTemp)*/

	getMenuList(c)
}
