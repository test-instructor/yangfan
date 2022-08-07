# 使用config实现例参之间的数传递

> 这几天在群里看到有些人的疑问，就是在各用例间的参数无法传递，比如登录状态、响应结果等。基于下面2种情况，虽然我觉得要保持用例的独立性，但真正要做到用例相对独立是比较困难的。同时我也有第1种情况的需求，所以就开始整理，下面代码会以header为例进行实现，参数导出也可以同样实现
> 1. 用例1设置的token，在用例2上无法使用，导致每个用例都要单独做登录的操作，导致操作繁琐
> 2. 用例1接口的返回结果，无法在用例2上使用，如果把多个用例放在同一个文件，会导致需要引用用例1的用例都需要加上，后期维护是个比较大的问题

# 如何接入

> 下面代码在[测试平台接入HttpRunner V4（一）基本功能接入](https://www.yuque.com/docs/share/bb392180-8ea9-46a0-a27b-bb4fbec3450e?#)基础上修改

1. 使用`config`进行传参，那么就需要使用应用传参，这样才操作后才可以在后续的用例中引用
```go
type TestCaseJson struct {
	JsonString        string
	ID                uint
	DebugTalkFilePath string
	Config            *TConfig   //增加config引用传参
	Name              string
}
```
```go
func (testCaseJson *TestCaseJson) ToTestCase() (*TestCase, error) {
	tc := &TCase{}
	var err error
	casePath := testCaseJson.JsonString
	tc, err = loadFromString(casePath)
	if err != nil {
		return nil, err
	}

	err = tc.MakeCompat()
	if err != nil {
		return nil, err
	}
	
	testCaseJson.Config.Path = testCaseJson.GetPath()
	testCase := &TestCase{
		ID:     testCaseJson.ID,
		Config: testCaseJson.Config,  //设置config
	}

	projectRootDir, err := GetProjectRootDirPath(testCaseJson.GetPath())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get project root dir")
	}

```
2. 获取接口的请求数据，可以用反射（技术没到家，暂时用简单粗暴的方法）取到对应的数据，我的做法比较简单，写一个和数据一样的结构体，然后做一次转换
```go
package hrp
//对应的结构体
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

```
3. 需要在TStep结构体中新增ExportHeader字段，用来提取需要导出的header
```go
Export        []string               `json:"export,omitempty" yaml:"export,omitempty"`
```
4. 获取token内容并且写入到config
```go
//修改session.go Start 函数
//在函数最后面添加以下代码
var StepResults hrp.StepResultStruct
stepResultStr, _ := json.Marshal(stepResult)
json.Unmarshal(stepResultStr, &StepResults)
for _, v := range step.Struct().ExportHeader {
    headerKey := textproto.CanonicalMIMEHeaderKey(v)
    r.testCase.Config.Headers[headerKey] = StepResults.Data.ReqResps.Request.Headers[headerKey]
}
```
5. 调用代码
```go
func RunTask() {

	apiConfig := interfacecase.ApiConfig{ID: timerTask.RunConfig.ID}
	
	var l []hrp.ITestCase
	
	for _, testCase := range timerTask.ApiTestCase {
		toTestCase := ToTestCase{TestSteps: testCase.TStep}
		caseJson, _ := json.Marshal(toTestCase)
		tc := &hrp.TestCaseJson{
			JsonString:        string(caseJson),
			ID:                testCase.ID,
			DebugTalkFilePath: debugTalkFilePath,
			Config:            &tConfig,
			Name:              testCase.Name,
		}
		testCase, _ := tc.ToTestCase()
		l = append(l, testCase)
	}
	reports, _ = hrp.NewRunner(t).
		SetFailfast(false).
		SetHTTPStatOn().
		RunJsons(l...)

}
```
