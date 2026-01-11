package projectmgr

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/middleware"
)

type UserProjectAccessRouter struct{}

// InitUserProjectAccessRouter 初始化 项目成员与权限 路由信息
func (s *UserProjectAccessRouter) InitUserProjectAccessRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	upaRouter := Router.Group("upa").Use(middleware.OperationRecord())
	upaRouterWithoutRecord := Router.Group("upa")
	upaRouterWithoutAuth := PublicRouter.Group("upa")
	{
		upaRouter.POST("createUserProjectAccess", upaApi.CreateUserProjectAccess)             // 新建项目成员与权限
		upaRouter.DELETE("deleteUserProjectAccess", upaApi.DeleteUserProjectAccess)           // 删除项目成员与权限
		upaRouter.DELETE("deleteUserProjectAccessByIds", upaApi.DeleteUserProjectAccessByIds) // 批量删除项目成员与权限
		upaRouter.PUT("updateUserProjectAccess", upaApi.UpdateUserProjectAccess)              // 更新项目成员与权限
	}
	{
		upaRouterWithoutRecord.GET("findUserProjectAccess", upaApi.FindUserProjectAccess)       // 根据ID获取项目成员与权限
		upaRouterWithoutRecord.GET("getUserProjectAccessList", upaApi.GetUserProjectAccessList) // 获取项目成员与权限列表
	}
	{
		upaRouterWithoutAuth.GET("getUserProjectAccessPublic", upaApi.GetUserProjectAccessPublic) // 项目成员与权限开放接口
	}
}
