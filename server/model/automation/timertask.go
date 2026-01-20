// 自动生成模板TimerTask
package automation

import (
	"time"

	"github.com/test-instructor/yangfan/server/v2/global"
	"gorm.io/datatypes"
)

// 定时任务 结构体  TimerTask
type TimerTask struct {
	global.GVA_MODEL
	Name           *string        `json:"name" form:"name" gorm:"comment:任务名称;column:name;"`                                         //任务名称
	RunTime        *string        `json:"runTime" form:"runTime" gorm:"comment:运行时间cron;column:run_time;"`                           //运行时间
	NextRunTime    *time.Time     `json:"nextRunTime" form:"nextRunTime" gorm:"comment:下次运行时间;column:next_run_time;"`                //下次运行时间
	Status         *bool          `json:"status" form:"status" gorm:"comment:运行状态;column:status;"`                                   //运行状态
	RunNumber      *int64         `json:"runNumber" form:"runNumber" gorm:"comment:运行次数;column:run_number;"`                         //运行次数
	ConfigName     *string        `json:"configName" form:"configName" gorm:"column:config_name;"`                                   //运行配置
	ConfigID       *uint          `json:"configID" form:"configID" gorm:"comment:运行配置;column:config_id;"`                            //运行配置ID
	Tag            datatypes.JSON `json:"tag" form:"tag" gorm:"comment:API定时任务标签;column:tag;" swaggertype:"array,object"`            //标签
	EnvName        *string        `json:"envName" form:"envName" gorm:"comment:所属环境名称;column:env_name;"`                             //环境名称
	EnvID          *uint          `json:"envID" form:"envID" gorm:"comment:所属环境;column:env_id;"`                                     //环境ID
	MessageName    *string        `json:"messageName" form:"messageName" gorm:"column:message_name;"`                                //消息名称
	MessageID      *uint          `json:"messageID" form:"messageID" gorm:"comment:消息发送;column:message_id;"`                         //消息D
	NotifyEnabled  *bool          `json:"notifyEnabled" form:"notifyEnabled" gorm:"comment:是否发送消息;column:notify_enabled;"`           //是否发送消息
	NotifyRule     *string        `json:"notifyRule" form:"notifyRule" gorm:"comment:发送规则(always/success/fail);column:notify_rule;"` //发送规则
	Describe       *string        `json:"describe" form:"describe" gorm:"comment:备注;column:describe;"`                               //备注
	RunnerNodeName *string        `json:"runnerNodeName" form:"runnerNodeName" gorm:"comment:运行节点;column:runner_node_name;"`         //运行节点
	Failfast       *bool          `json:"failfast" form:"failfast" gorm:"comment:失败停止;column:failfast;"`                             //失败停止
	ProjectId      int64          `json:"projectId" form:"projectId" gorm:"column:project_id;"`                                      //项目信息
}

// TableName 定时任务 TimerTask自定义表名 timer_task
func (TimerTask) TableName() string {
	return "yf_timer_task"
}

type TimerTaskTag struct {
	global.GVA_MODEL
	Name        string `json:"name" form:"name" gorm:"comment:标签名称;column:name;"`                    //标签名称
	Description string `json:"description" form:"description" gorm:"comment:描述;column:description;"` //标签描述
	ProjectId   int64  `json:"projectId" form:"projectId" gorm:"column:project_id;"`                 //项目ID
}

func (TimerTaskTag) TableName() string {
	return "yf_timer_task_tag"
}

type TimerTaskCaseList struct {
	global.GVA_MODEL
	TimerTaskID uint     `json:"task_id" form:"task_id" gorm:"column:task_id;"`
	AutoCaseID  uint     `json:"case_id" form:"case_id" gorm:"column:case_id;"`
	AutoCase    AutoCase `json:"auto_case" form:"auto_case"`
	Sort        int64    `json:"sort" form:"sort" gorm:"column:sort;"`
}

func (TimerTaskCaseList) TableName() string {
	return "yf_timer_task_case_list"
}
