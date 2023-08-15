package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/common/response"
	"github.com/test-instructor/yangfan/server/model/system"
	"github.com/test-instructor/yangfan/server/utils"
	"net/http"
)

//var projectService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

// ProjectHandler 项目拦截器
func ProjectHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := utils.GetUserID(c)
		var userProject system.SysUserProject
		project := utils.GetUserProject(c)
		db := global.GVA_DB
		db.Where("sys_user_id =?", userID).Where("project_id =?", project).Find(&userProject)
		if userProject.SysUserID == 0 {
			response.FailWithDetailed(gin.H{}, "项目权限不足", c)
			c.Abort()
			return
		}
		//设置project 结构体信息
		c.Set("project", project)
		c.Next()
	}
}

func ProjectAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := utils.GetUserID(c)
		projectID := utils.GetUserProject(c) // Assuming you get the project ID from utils

		var userProject system.SysUserProject
		db := global.GVA_DB
		db.Where("sys_user_id = ?", userID).Where("project_id = ?", projectID).Find(&userProject)

		if userProject.SysUserID == 0 {
			response.FailWithDetailed(gin.H{}, "项目权限不足", c)
			c.Abort()
			return
		}

		// Check the HTTP method and corresponding permission
		switch c.Request.Method {
		case http.MethodGet:
			// For GET requests, allow access without further checks
			break
		case http.MethodPost, http.MethodPut:
			// For POST and PUT requests, check the 'save' permission
			if !userProject.Save {
				response.FailWithDetailed(gin.H{}, "权限不足，请联系管理员增加权限", c)
				c.Abort()
				return
			}
		case http.MethodDelete:
			// For DELETE requests, check the 'delete' permission
			if !userProject.Delete {
				response.FailWithDetailed(gin.H{}, "权限不足，请联系管理员增加权限", c)
				c.Abort()
				return
			}
		default:
			// Handle other HTTP methods if needed
		}

		c.Next()
	}
}
