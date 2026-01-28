package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/middleware"
)

type LLMModelConfigRouter struct{}

// InitLLMModelConfigRouter 初始化 大语言模型配置 路由信息
func (s *LLMModelConfigRouter) InitLLMModelConfigRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	llmconfigRouter := Router.Group("llmconfig").Use(middleware.OperationRecord())
	llmconfigRouterWithoutRecord := Router.Group("llmconfig")
	llmconfigRouterWithoutAuth := PublicRouter.Group("llmconfig")
	{
		llmconfigRouter.POST("createLLMModelConfig", llmconfigApi.CreateLLMModelConfig)             // 新建大语言模型配置
		llmconfigRouter.DELETE("deleteLLMModelConfig", llmconfigApi.DeleteLLMModelConfig)           // 删除大语言模型配置
		llmconfigRouter.DELETE("deleteLLMModelConfigByIds", llmconfigApi.DeleteLLMModelConfigByIds) // 批量删除大语言模型配置
		llmconfigRouter.PUT("updateLLMModelConfig", llmconfigApi.UpdateLLMModelConfig)              // 更新大语言模型配置
	}
	{
		llmconfigRouterWithoutRecord.GET("findLLMModelConfig", llmconfigApi.FindLLMModelConfig)       // 根据ID获取大语言模型配置
		llmconfigRouterWithoutRecord.GET("getLLMModelConfigList", llmconfigApi.GetLLMModelConfigList) // 获取大语言模型配置列表
	}
	{
		llmconfigRouterWithoutAuth.GET("getLLMModelConfigPublic", llmconfigApi.GetLLMModelConfigPublic) // 大语言模型配置开放接口
	}
}
