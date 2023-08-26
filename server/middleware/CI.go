package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/common/response"
	interfacecaseReq "github.com/test-instructor/yangfan/server/model/interfacecase/request"
	"github.com/test-instructor/yangfan/server/model/system"
)

func CIAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tagReq interfacecaseReq.CIRun
		_ = c.ShouldBindQuery(&tagReq)
		if tagReq.ProjectID == 0 || tagReq.Key == "" || tagReq.UUID == "" {
			response.FailWithDetailed(gin.H{}, "鉴权信息不完整，请检查project、uuid和secret参数", c)
			c.Abort()
			return
		}
		var project system.Project
		err := global.GVA_DB.Where("id = ?", tagReq.ProjectID).First(&project)
		if err != nil {
			response.FailWithDetailed(gin.H{}, "无法查询到对应的项目", c)
			c.Abort()
			return
		}
		if tagReq.UUID == project.UUID && tagReq.Secret == project.Secret {
			c.Next()
		} else {
			response.FailWithDetailed(gin.H{}, "鉴权信息错误，请检查project、uuid和secret参数", c)
			c.Abort()
			return
		}
	}
}
