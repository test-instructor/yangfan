package router

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/cheetah/server/plugin/fslogin/api"
)

type FsLoginRouter struct {
}

func (s *FsLoginRouter) InitFsLoginRouter(Router *gin.RouterGroup) {
	plugRouter := Router
	plugApi := api.ApiGroupApp.FsLoginApi
	{
		plugRouter.GET("login", plugApi.Login)
	}
}
