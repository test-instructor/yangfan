package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/common/response"
	"github.com/test-instructor/yangfan/server/model/system"
	"github.com/test-instructor/yangfan/server/utils"
)

//var projectService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

// ProjectHandler 项目拦截器
func ProjectHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := utils.GetUserID(c)
		var userProject system.SysUseProject
		project := utils.GetUserProject(c)
		db := global.GVA_DB
		db.Where("sys_user_id =?", userID).Where("project_id =?", project).Find(&userProject)
		if userProject.SysUserId == 0 {
			response.FailWithDetailed(gin.H{}, "项目权限不足", c)
			c.Abort()
			return
		}
		//设置project 结构体信息
		c.Set("project", project)
		c.Next()
	}
}
