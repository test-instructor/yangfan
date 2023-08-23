package interfacecase

import "github.com/test-instructor/yangfan/server/global"

type MessageType string

const MessageTypeFeishu MessageType = "feishu"
const MessageTypeDingtalk MessageType = "dingtalk"
const MessageTypeWechat MessageType = "wechat"

var _ = []MessageType{MessageTypeFeishu, MessageTypeDingtalk}

type ApiMessage struct {
	global.GVA_MODEL
	Operator
	Name      string      `json:"name" form:"name" gorm:"column:name;comment:消息名称;"`
	Type      MessageType `json:"type" form:"type" gorm:"column:type;comment:消息类型;"`
	TypeName  string      `json:"type_name" form:"type_name" gorm:"column:type_name;comment:消息类型名称;"`
	WebHook   string      `json:"webhook" form:"webhook" gorm:"column:webhook;comment:webhook地址;"`
	Signature string      `json:"signature" form:"signature" gorm:"column:signature;comment:签名;"`
	Fail      bool        `json:"fail" form:"fail" gorm:"column:fail;comment:仅失败时发送;"`
}

type ApiMessageLog struct {
	global.GVA_MODEL
	Operator
	ApiMessageID uint       `json:"api_message_id,omitempty" form:"api_message_id" gorm:"column:api_message_id;comment:消息ID;"`
	ApiMessage   ApiMessage `json:"api_message,omitempty" form:"api_message"`
	Message      string     `json:"message,omitempty" form:"message" gorm:"column:message;comment:消息内容;type:text"`
	Status       bool       `json:"status,omitempty" form:"status" gorm:"column:status;comment:消息发送状态;"`
	StatusCode   int        `json:"status_code,omitempty" form:"status_code" gorm:"column:status_code;comment:消息发送状态码;"`
}
