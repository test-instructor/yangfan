package platform

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/projectmgr"
)

// GetProjectID 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserProject(c *gin.Context) uint {
	var userProject projectmgr.Project
	project, _ := strconv.ParseInt(c.Param("project"), 10, 64)
	userProject.ID = uint(project)
	db := global.GVA_DB
	db.Model(projectmgr.Project{}).First(&userProject)
	return uint(project)
}
