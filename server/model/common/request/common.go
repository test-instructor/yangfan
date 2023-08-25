package request

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}

// GetById Find by id structure
type GetById struct {
	ID int `json:"id" form:"id"` // 主键ID
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"` // id列表
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId uint `json:"authorityId" form:"authorityId"` // 角色ID
}

type Empty struct{}

type ApiCaseIdReq struct {
	ID     uint   `json:"ID" form:"ID"`         // 主键ID
	ApiID  uint   `json:"apiID" form:"apiID"`   // 接口ID
	CaseID uint   `json:"caseID" form:"caseID"` // 用例ID
	Type   string `json:"type" form:"type"`     // 类型
	Detail bool   `json:"detail" form:"detail"` // 是否详情
}

type RunningType int

var (
	RunningTypeRun       RunningType = 1
	RunningTypeRebalance RunningType = 2
	RunningTypeStop      RunningType = 3
)

type Operation struct {
	Running    RunningType `json:"running"`    // 运行类型
	SpawnCount int64       `json:"spawnCount"` // 并发数
	SpawnRate  float64     `json:"spawnRate"`  // 并发率
	Interval   *Interval   `json:"interval"`   // 间隔
}

type RunCaseReq struct {
	ApiID        uint      `json:"apiID" form:"apiID"`                  // 接口ID
	ConfigID     uint      `json:"configID" form:"configID"`            // 配置ID
	CaseID       uint      `json:"caseID" form:"caseID"`                // 用例ID
	RunType      uint      `json:"run_type" form:"run_type"`            // 运行类型
	TaskID       uint      `json:"taskID" form:"taskID"`                // 任务ID
	Operation    Operation `json:"operation"`                           // 操作
	TagID        uint      `json:"tagID" form:"TagID"`                  // 标签ID
	ProjectID    uint      `json:"-"`                                   // 项目ID
	Env          uint      `json:"env" form:"env"`                      // 环境ID
	ReportID     uint      `json:"reportID" form:"reportID"`            // 报告ID
	ApiMessageID uint      `json:"api_message_id" gorm:"comment:消息发送;"` // 消息ID
	ReportCIID   uint      `json:"ci_id" yaml:"ci_id"`                  // CI ID
}

type Interval struct {
	IntervalTime   int64 `json:"intervalTime"`   // 间隔时间
	IntervalCount  int64 `json:"intervalCount"`  // 间隔次数
	IntervalNumber int64 `json:"intervalNumber"` // 间隔数量
}
