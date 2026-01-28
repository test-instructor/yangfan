// 自动生成模板LLMModelConfig
package platform

import (
	"github.com/test-instructor/yangfan/server/v2/global"
)

// 大语言模型配置 结构体  LLMModelConfig
type LLMModelConfig struct {
	global.GVA_MODEL
	Name                *string  `json:"name" form:"name" gorm:"comment:模型名称;column:name;size:255;"`                                            //模型名称
	RequestSchema       *string  `json:"requestSchema" form:"requestSchema" gorm:"comment:请求模式;column:request_schema;"`                         //请求模式
	Model               *string  `json:"model" form:"model" gorm:"comment:模型标识;column:model;"`                                                  //模型标识
	BaseURL             *string  `json:"baseURL" form:"baseURL" gorm:"comment:基础地址;column:base_url;size:500;"`                                  //基础地址
	APIKey              *string  `json:"apiKey" form:"apiKey" gorm:"comment:接口密钥;column:api_key;size:500;"`                                     //密钥
	SupportFormatOutput *bool    `json:"supportFormatOutput" form:"supportFormatOutput" gorm:"comment:是否支持格式化输出;column:support_format_output;"` //格式化输出
	ReasoningEffort     *string  `json:"reasoningEffort" form:"reasoningEffort" gorm:"comment:推理强度配置;column:reasoning_effort;"`                 //推理强度
	MaxTokens           *int64   `json:"maxTokens" form:"maxTokens" gorm:"comment:最大生成token数;column:max_tokens;"`                               //最大tokens
	Temperature         *float64 `json:"temperature" form:"temperature" gorm:"comment:随机性参数0-2;column:temperature;"`                            //随机性参数
	TopP                *float64 `json:"topP" form:"topP" gorm:"comment:核心采样参数;column:top_p;"`                                                  //核心采样
	Enabled             *bool    `json:"enabled" form:"enabled" gorm:"comment:模型是否启用;column:enabled;"`                                          //启用状态
	Timeout             *int64   `json:"timeout" form:"timeout" gorm:"comment:API调用超时时间（秒）;column:timeout;"`                                    //超时时间
	Description         *string  `json:"description" form:"description" gorm:"comment:模型描述信息;column:description;type:text;"`                    //模型描述

	ProjectId int64 `json:"projectId" form:"projectId" gorm:"column:project_id;"` //项目信息
}

// TableName 大语言模型配置 LLMModelConfig自定义表名 llm_model_configs
func (LLMModelConfig) TableName() string {
	return "lc_llm_model_configs"
}
