package interfacecase

import (
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/model/system"
	"gorm.io/datatypes"
	"time"
)

type CaseType int
type RunType int

var (
	CaseTypeApi     CaseType = 1 // api 运行
	CaseTypeStep    CaseType = 2 // 套件运行
	CaseTypeCases   CaseType = 3 // 用例运行
	CaseTypeTask    CaseType = 4 // 定时任务运行
	RunTypeDebug    RunType  = 1 // 调试模式
	RunTypeRuning   RunType  = 2 // 立即运行
	RunTypeRunBack  RunType  = 3 // 后台运行
	RunTypeRunTimer RunType  = 4 // 定时执行
	RunTypeRunSave  RunType  = 5 // 调试并保存
)

type ApiReport struct {
	global.GVA_MODEL
	Name      string             `json:"name"`
	Success   *bool              `json:"success"`
	Stat      ApiReportStat      `json:"stat"`
	StatID    uint               `json:"-"`
	Time      ApiReportTime      `json:"time"`
	TimeID    uint               `json:"-"`
	Platform  datatypes.JSON     `json:"platform"`
	Details   []ApiReportDetails `json:"details"`
	CaseType  CaseType           `json:"runType"` //1、api，2、case，3、task
	RunType   RunType            `json:"type"`    //1、调试，2、立即运行，3、后台运行
	Status    int                `json:"status"`
	Project   system.Project     `json:"-"`
	ProjectID uint               `json:"-"`
	SetupCase bool               `json:"setup_case"`
}

type ApiReportTime struct {
	global.GVA_MODEL
	StartAt  time.Time `json:"start_at"`
	Duration float64   `json:"duration"`
}

type ApiReportStat struct {
	global.GVA_MODEL
	TestCases   ApiReportStatTestcases `json:"testcases"`
	TestCasesID uint                   `json:"-"`
	TestSteps   ApiReportStatTeststeps `json:"teststeps"`
	TestStepsID uint                   `json:"-"`
}

type ApiReportStatTestcases struct {
	global.GVA_MODEL
	Total   int `json:"total"`
	Success int `json:"success"`
	Fail    int `json:"fail"`
}

type ApiReportStatTeststeps struct {
	global.GVA_MODEL
	Total     int `json:"total"`
	Successes int `json:"successes"`
	Failures  int `json:"failures"`
}

type ApiReportDetails struct {
	global.GVA_MODEL
	Name         string                    `json:"name"`
	Success      bool                      `json:"success"`
	Stat         datatypes.JSON            `json:"stat"`
	Time         datatypes.JSON            `json:"time"`
	InOut        datatypes.JSON            `json:"in_out"`
	Records      []ApiReportDetailsRecords `json:"records"`
	ApiRecordsID uint                      `json:"-"`
	RootDir      string                    `json:"root_dir"`
	ApiReportID  uint                      `json:"-"`
	CaseID       uint                      `json:"caseID"`
}

type ApiReportDetailsRecords struct {
	global.GVA_MODEL
	ParntID            uint                            `json:"parntID"`
	Name               string                          `json:"name"`
	StepType           string                          `json:"step_type"`
	Success            bool                            `json:"success"`
	ElapsedMs          int                             `json:"elapsed_ms"`
	HttpStat           ApiReportDetailsRecordsHttpstat `json:"httpstat"`
	ValidateNumber     uint                            `json:"validate_number" form:"validate_number"`
	HttpStatID         uint                            `json:"-"`
	Data               datatypes.JSON                  `json:"data"`
	ExportVars         datatypes.JSON                  `json:"export_vars"`
	ContentSize        int                             `json:"content_size"`
	ApiReportDetailsID uint                            `json:"-"`
}

type ApiReportDetailsRecordsHttpstat struct {
	global.GVA_MODEL
	Connect          int `json:"Connect"`
	ContentTransfer  int `json:"ContentTransfer"`
	DNSLookup        int `json:"DNSLookup"`
	NameLookup       int `json:"NameLookup"`
	Pretransfer      int `json:"Pretransfer"`
	ServerProcessing int `json:"ServerProcessing"`
	StartTransfer    int `json:"StartTransfer"`
	TCPConnection    int `json:"TCPConnection"`
	TLSHandshake     int `json:"TLSHandshake"`
	Total            int `json:"Total"`
}
