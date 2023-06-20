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
	Ids []int `json:"ids" form:"ids"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId uint `json:"authorityId" form:"authorityId"` // 角色ID
}

type Empty struct{}

type ApiCaseIdReq struct {
	ApiID  uint   `json:"apiID" form:"apiID"`
	CaseID uint   `json:"caseID" form:"caseID"`
	Type   string `json:"type" form:"type"`
}

type RunningType int

var (
	RunningTypeRun       RunningType = 1
	RunningTypeRebalance RunningType = 2
	RunningTypeStop      RunningType = 3
)

type Operation struct {
	Running    RunningType `json:"running"`
	SpawnCount int64       `json:"spawnCount"`
	SpawnRate  float64     `json:"spawnRate"`
}

type RunCaseReq struct {
	ApiID     uint      `json:"apiID" form:"apiID"`
	ConfigID  uint      `json:"configID" form:"configID"`
	CaseID    uint      `json:"caseID" form:"caseID"`
	RunType   uint      `json:"run_type" form:"run_type"`
	TaskID    uint      `json:"taskID" form:"taskID"`
	Operation Operation `json:"operation"`
	TagID     uint      `json:"tagID" form:"TagID"`
	ProjectID uint      `json:"-"`
	Env       uint      `json:"env" form:"env"`
	ReportID  uint      `json:"reportID" form:"reportID"`
}
