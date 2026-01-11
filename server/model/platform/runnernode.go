// 自动生成模板RunnerNode
package platform

import (
	"time"

	"github.com/test-instructor/yangfan/server/v2/global"
)

// 节点 结构体  RunnerNode
type RunnerNode struct {
	global.GVA_MODEL
	NodeName      string     `json:"nodeName" form:"nodeName" gorm:"comment:节点名称（Docker固定名称）;column:node_name;" binding:"required"` //节点名称
	NodeId        *string    `json:"nodeId" form:"nodeId" gorm:"comment:唯一主键（Docker固定名称）;column:node_id;" binding:"required"`       //节点ID
	Alias         *string    `json:"alias" form:"alias" gorm:"comment:节点别名（可修改）;column:alias;" binding:"required"`                  //别名
	Ip            *string    `json:"ip" form:"ip" gorm:"comment:Run服务所在IP（Docker内网/宿主机IP）;column:ip;" binding:"required"`           //IP
	Port          *int64     `json:"port" form:"port" gorm:"comment:Run服务端口（如无暴露则存0）;column:port;"`                                 //端口
	Status        *int64     `json:"status" form:"status" gorm:"comment:0 = 离线，1 = 在线;column:status;"`                              //状态
	LastHeartbeat *time.Time `json:"lastHeartbeat" form:"lastHeartbeat" gorm:"comment:最后心跳时间;column:last_heartbeat;"`               //最后心跳
	CreateTime    *time.Time `json:"createTime" form:"createTime" gorm:"comment:注册时间;column:create_time;"`                          //注册时间

	ProjectId int64 `json:"projectId" form:"projectId" gorm:"column:project_id;"` //项目信息
}

// TableName 节点 RunnerNode自定义表名 runner_node
func (RunnerNode) TableName() string {
	return "lc_runner_node"
}
