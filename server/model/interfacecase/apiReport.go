package interfacecase

import (
	"time"

	"gorm.io/datatypes"

	"github.com/test-instructor/yangfan/server/global"
)

type CaseType int
type RunType int

const (
	CaseTypeApi         CaseType = 1 // api 运行
	CaseTypeStep        CaseType = 2 // 步骤运行
	CaseTypeCases       CaseType = 3 // 用例运行
	CaseTypeTask        CaseType = 4 // 定时任务运行
	CaseTypeBoomerDebug CaseType = 5 // 性能任务调试
	CaseTypeBoomer      CaseType = 6 // 性能任务运行
	CaseTypeTag         CaseType = 7 // 标签运行
	RunTypeDebug        RunType  = 1 // 调试模式
	RunTypeRuning       RunType  = 2 // 立即运行
	RunTypeRunBack      RunType  = 3 // 后台运行
	RunTypeRunTimer     RunType  = 4 // 定时执行
	RunTypeRunSave      RunType  = 5 // 调试并保存
	RunTypeCI           RunType  = 6 // 调试并保存
)

var _ = CaseTypeBoomer
var _ = RunTypeRunSave

type ApiReport struct {
	global.GVA_MODEL
	Operator
	Name       string             `json:"name,omitempty"`
	Success    *bool              `json:"success,omitempty"`
	Stat       *ApiReportStat     `json:"stat,omitempty"`
	StatID     uint               `json:"-"`
	Time       *ApiReportTime     `json:"time,omitempty"`
	TimeID     uint               `json:"-"`
	Platform   datatypes.JSON     `json:"platform,omitempty"`
	Details    []ApiReportDetails `json:"details,omitempty"`
	CaseType   CaseType           `json:"runType,omitempty"` //1、api，2、case，3、task
	RunType    RunType            `json:"type,omitempty"`    //1、调试，2、立即运行，3、后台运行
	Status     int                `json:"status,omitempty"`
	SetupCase  bool               `json:"setup_case,omitempty"`
	Describe   string             `json:"describe,omitempty" form:"describe" gorm:"column:describe;comment:;"`
	ApiEnvName string             `json:"api_env_name,omitempty" gorm:"comment:所属环境名称;"`
	ApiEnvID   uint               `json:"api_env_id,omitempty" gorm:"comment:所属环境;"`
	Hostname   string             `json:"hostname,omitempty" gorm:"comment:主机名;"`
}

type ApiReportTime struct {
	global.GVA_MODEL
	StartAt  time.Time `json:"start_at,omitempty"`
	Duration float64   `json:"duration,omitempty"`
}

type ApiReportStat struct {
	global.GVA_MODEL
	TestCases   *ApiReportStatTestcases `json:"testcases,omitempty"`
	TestCasesID uint                    `json:"-"`
	TestSteps   *ApiReportStatTeststeps `json:"teststeps,omitempty"`
	TestStepsID uint                    `json:"-"`
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
	CaseID       uint                      `json:"caseID,omitempty"`
	Name         string                    `json:"name,omitempty"`
	Success      bool                      `json:"success,omitempty"`
	Stat         datatypes.JSON            `json:"stat,omitempty"`
	Time         datatypes.JSON            `json:"time,omitempty"`
	InOut        datatypes.JSON            `json:"in_out,omitempty"`
	Records      []ApiReportDetailsRecords `json:"records,omitempty"`
	ApiRecordsID uint                      `json:"-"`
	RootDir      string                    `json:"root_dir,omitempty"`
	ApiReportID  uint                      `json:"-"`
}

type ApiReportDetailsRecords struct {
	global.GVA_MODEL
	ParntID            uint                          `json:"parntID,omitempty"`
	Name               string                        `json:"name,omitempty"`
	StepType           string                        `json:"step_type,omitempty"`
	Success            bool                          `json:"success,omitempty"`
	ElapsedMs          int                           `json:"elapsed_ms,omitempty"`
	ValidateNumber     uint                          `json:"validate_number,omitempty" form:"validate_number"`
	Data               []ApiReportDetailsRecordsData `json:"data,omitempty"`
	ExportVars         datatypes.JSON                `json:"export_vars,omitempty"`
	ContentSize        int                           `json:"content_size,omitempty"`
	ApiReportDetailsID uint                          `json:"-"`
}

type ApiReportDetailsRecordsData struct {
	ID                        int                                  `json:"ID,omitempty"`
	ParntID                   int                                  `json:"parntID,omitempty"`
	Name                      string                               `json:"name,omitempty"`
	StepType                  string                               `json:"step_type,omitempty"`
	Success                   bool                                 `json:"success,omitempty"`
	ElapsedMs                 int                                  `json:"elapsed_ms,omitempty"`
	HttpStat                  *ApiReportDetailsRecordsDataHttpstat `json:"httpstat,omitempty"`
	Attachment                string                               `json:"attachments,omitempty"`
	HttpStatID                uint                                 `json:"-"`
	Data                      datatypes.JSON                       `json:"data,omitempty"`
	ExportVars                datatypes.JSON                       `json:"export_vars,omitempty"`
	ContentSize               int                                  `json:"content_size,omitempty"`
	ApiReportDetailsRecordsID uint                                 `json:"-"`
}

type ApiReportDetailsRecordsDataHttpstat struct {
	global.GVA_MODEL
	Connect          int `json:"Connect,omitempty"`
	ContentTransfer  int `json:"ContentTransfer,omitempty"`
	DNSLookup        int `json:"DNSLookup,omitempty"`
	NameLookup       int `json:"NameLookup,omitempty"`
	Pretransfer      int `json:"Pretransfer,omitempty"`
	ServerProcessing int `json:"ServerProcessing,omitempty"`
	StartTransfer    int `json:"StartTransfer,omitempty"`
	TCPConnection    int `json:"TCPConnection,omitempty"`
	TLSHandshake     int `json:"TLSHandshake,omitempty"`
	Total            int `json:"Total,omitempty"`
}
