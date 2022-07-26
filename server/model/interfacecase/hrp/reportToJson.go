package hrp

type StepResultStruct struct {
	ID          int                    `json:"ID"`
	ParntID     int                    `json:"parntID"`
	Name        string                 `json:"name"`
	StepType    string                 `json:"step_type"`
	Success     bool                   `json:"success"`
	ElapsedMs   int                    `json:"elapsed_ms"`
	Httpstat    Httpstat               `json:"httpstat"`
	Data        Data                   `json:"data"`
	ContentSize int                    `json:"content_size"`
	ExportVars  map[string]interface{} `json:"export_vars"`
}
type Httpstat struct {
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

type Request struct {
	Headers map[string]string `json:"headers"`
	Method  string            `json:"method"`
	URL     string            `json:"url"`
}

type Response struct {
	//Body       string  `json:"body"`
	Cookies    map[string]interface{} `json:"cookies"`
	Headers    map[string]string      `json:"headers"`
	StatusCode int                    `json:"status_code"`
}
type ReqResps struct {
	Request  Request  `json:"request"`
	Response Response `json:"response"`
}
type Validators struct {
	Check       string `json:"check"`
	Assert      string `json:"assert"`
	Expect      string `json:"expect"`
	CheckValue  string `json:"check_value"`
	CheckResult string `json:"check_result"`
}
type Data struct {
	Success    bool         `json:"success"`
	ReqResps   ReqResps     `json:"req_resps"`
	Validators []Validators `json:"validators"`
}
