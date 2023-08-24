package request

type CIRun struct {
	TagID     uint   `json:"tag" form:"tag"`
	TaskID    uint   `json:"task" form:"task"`
	EnvID     uint   `json:"env" form:"env"`
	UUID      string `json:"uuid" form:"uuid"`
	Secret    string `json:"secret" form:"secret"`
	ProjectID uint   `json:"project" form:"project"`
	MessageID uint   `json:"message" form:"message"`
	ReportID  uint   `json:"report" form:"report"`
}
