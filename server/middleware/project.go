package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/model/common/response"
	"github.com/test-instructor/cheetah/server/model/system"
	"github.com/test-instructor/cheetah/server/service"
	"github.com/test-instructor/cheetah/server/utils"
	"strconv"
)

var projectService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

// 项目拦截器
func ProjectHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := utils.GetUserID(c)
		//判断用户在项目中的权限
		var userProject system.SysUseProject
		projectId := c.Param("project")
		projectsInt, _ := strconv.Atoi(projectId)
		db := global.GVA_DB
		db.Where("sys_user_id =?", userID).Where("project_id =?", uint(projectsInt)).Find(&userProject)
		if userProject.SysUserId == 0 {
			response.FailWithDetailed(gin.H{}, "项目权限不足", c)
			c.Abort()
			return
		}
		//设置project 结构体信息
		project := system.Project{GVA_MODEL: global.GVA_MODEL{ID: uint(projectsInt)}}
		c.Set("project", project)
		c.Next()

	}
}
