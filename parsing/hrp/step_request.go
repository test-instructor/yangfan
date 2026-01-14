package hrp

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/andybalholm/brotli"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/test-instructor/yangfan/parsing/hrp/internal/builtin"
	"github.com/test-instructor/yangfan/parsing/hrp/internal/code"
	"github.com/test-instructor/yangfan/parsing/hrp/internal/json"
	"github.com/test-instructor/yangfan/parsing/hrp/pkg/httpstat"
)

type HTTPMethod string

// 定义请求方法常量
const (
	httpGET     HTTPMethod = "GET"
	httpHEAD    HTTPMethod = "HEAD"
	httpPOST    HTTPMethod = "POST"
	httpPUT     HTTPMethod = "PUT"
	httpDELETE  HTTPMethod = "DELETE"
	httpOPTIONS HTTPMethod = "OPTIONS"
	httpPATCH   HTTPMethod = "PATCH"
)

// Request represents HTTP request data structure.
// This is used for teststep.
// 请求结构体
type Request struct {
	Method         HTTPMethod             `json:"method" yaml:"method"` // required
	URL            string                 `json:"url" yaml:"url"`       // required
	HTTP2          bool                   `json:"http2,omitempty" yaml:"http2,omitempty"`
	Params         map[string]interface{} `json:"params,omitempty" yaml:"params,omitempty"`
	Headers        map[string]string      `json:"headers,omitempty" yaml:"headers,omitempty"`
	Cookies        map[string]string      `json:"cookies,omitempty" yaml:"cookies,omitempty"`
	Body           interface{}            `json:"body,omitempty" yaml:"body,omitempty"`
	Json           interface{}            `json:"json,omitempty" yaml:"json,omitempty"`
	Data           interface{}            `json:"data,omitempty" yaml:"data,omitempty"`
	Timeout        float64                `json:"timeout,omitempty" yaml:"timeout,omitempty"` // timeout in seconds
	AllowRedirects bool                   `json:"allow_redirects,omitempty" yaml:"allow_redirects,omitempty"`
	Verify         bool                   `json:"verify,omitempty" yaml:"verify,omitempty"`
	Upload         map[string]interface{} `json:"upload,omitempty" yaml:"upload,omitempty"`
}

// 创建requestBuilder对象，初始化数据
func newRequestBuilder(parser *Parser, config *TConfig, stepRequest *Request) *requestBuilder {
	// convert request struct to map
	// 将request对象转换为map
	jsonRequest, _ := json.Marshal(stepRequest)
	var requestMap map[string]interface{}
	_ = json.Unmarshal(jsonRequest, &requestMap)
	// 创建http.Request对象
	request := &http.Request{
		Header: make(http.Header),
	}
	// 设置http版本
	if stepRequest.HTTP2 {
		request.ProtoMajor = 2
		request.ProtoMinor = 0
	} else {
		request.ProtoMajor = 1
		request.ProtoMinor = 1
	}
	// 返回requestBuilder对象
	return &requestBuilder{
		stepRequest: stepRequest,
		req:         request,
		config:      config,
		parser:      parser,
		requestMap:  requestMap,
	}
}

// requestBuilder 结构体
// stepRequest Request 对象
// req 用http.Request 对象做http操作
// parser 解析器，存放plugin
// config 配置文件
type requestBuilder struct {
	stepRequest *Request
	req         *http.Request
	parser      *Parser
	config      *TConfig
	requestMap  map[string]interface{}
}

// prepareHeaders 获取请求头内容
func (r *requestBuilder) prepareHeaders(stepVariables map[string]interface{}) error {
	// prepare request headers
	// 创建stepHeaders临时对象，并将request和config中的headers合并
	stepHeaders := r.stepRequest.Headers
	if r.config.Headers != nil {
		// override headers
		stepHeaders = mergeMap(stepHeaders, r.config.Headers)
	}

	// header 参数化设置，如果没有设置请求头则跳过
	if len(stepHeaders) > 0 {
		//获取header中参数化的值
		headers, err := r.parser.ParseHeaders(stepHeaders, stepVariables)
		if err != nil {
			// 错误日志
			return errors.Wrap(err, "parse headers failed")
		}
		//将header转换成http.Header对象
		for key, value := range headers {
			// omit pseudo header names for HTTP/1, e.g. :authority, :method, :path, :scheme
			if strings.HasPrefix(key, ":") {
				continue
			}
			r.req.Header.Add(key, value)

			// prepare content length
			// 设置 content length 大小
			if strings.EqualFold(key, "Content-Length") && value != "" {
				if l, err := strconv.ParseInt(value, 10, 64); err == nil {
					r.req.ContentLength = l
				}
			}
		}
	}

	// prepare request cookies
	// 获取cookie内容并设置到http.Request对象中
	for cookieName, cookieValue := range r.stepRequest.Cookies {
		value, err := r.parser.Parse(cookieValue, stepVariables)
		if err != nil {
			return errors.Wrap(err, "parse cookie value failed")
		}
		r.req.AddCookie(&http.Cookie{
			Name:  cookieName,
			Value: convertString(value),
		})
	}

	if r.req.Header.Get("User-Agent") == "" {
		r.req.Header.Set("User-Agent", "YangFan-Client/V2")
	}

	// update header
	// 将header设置到requestMap中，requestMap主要用于测试报告中的展示
	headers := make(map[string]string)
	for key, value := range r.req.Header {
		headers[key] = value[0]
	}
	r.requestMap["headers"] = headers
	return nil
}

// prepareUrlParams 获取请求参数,将函数、变量转换为常量
func (r *requestBuilder) prepareUrlParams(stepVariables map[string]interface{}) error {
	// parse step request url
	// 将变量解析为常量
	requestUrl, err := r.parser.ParseString(r.stepRequest.URL, stepVariables)
	if err != nil {
		log.Error().Err(err).Msg("parse request url failed")
		return err
	}
	var baseURL string
	if stepVariables["base_url"] != nil {
		baseURL = stepVariables["base_url"].(string)
	}
	// 合并base_url和request_url
	// 如果request_url中有域名则直接使用request_url
	rawUrl := buildURL(baseURL, convertString(requestUrl))

	// prepare request params
	// 获取请求参数
	var queryParams url.Values
	// 代码检查 r.stepRequest.Params 的长度是否大于 0，如果是，则执行参数解析操作
	if len(r.stepRequest.Params) > 0 {
		// r.parser.Parse 函数用于解析参数，返回解析后的结果 params 和一个可能的错误
		params, err := r.parser.Parse(r.stepRequest.Params, stepVariables)
		if err != nil {
			// 如果解析过程中出现错误，代码返回一个包装了错误信息的错误对象。
			return errors.Wrap(err, "parse request params failed")
		}
		// 将解析后的参数设置为url.Values
		parsedParams := params.(map[string]interface{})
		r.requestMap["params"] = parsedParams
		if len(parsedParams) > 0 {
			queryParams = make(url.Values)
			for k, v := range parsedParams {
				queryParams.Add(k, convertString(v))
			}
		}
	}
	// 如果queryParams有参数，则添加到url中
	if queryParams != nil {
		// append params to url
		// 将queryParams转换为字符串
		paramStr := queryParams.Encode()
		// 根据url是否已有参数添加对应的拼接符号
		if strings.IndexByte(rawUrl, '?') == -1 {
			rawUrl = rawUrl + "?" + paramStr
		} else {
			rawUrl = rawUrl + "&" + paramStr
		}
	}

	// prepare url
	// 将url转换为http.Request对象
	u, err := url.Parse(rawUrl)
	if err != nil {
		return errors.Wrap(err, "parse url failed")
	}
	r.req.URL = u
	r.req.Host = u.Host

	// update url
	// 将url设置到requestMap中，requestMap主要用于测试报告中的展示
	r.requestMap["url"] = u.String()

	return nil
}

// prepareBody 获取请求体内容，将函数、变量转换为常量
func (r *requestBuilder) prepareBody(stepVariables map[string]interface{}) error {
	// prepare request body
	// 如果请求体为空，则直接返回
	if r.stepRequest.Body == nil {
		return nil
	}

	data, err := r.parser.Parse(r.stepRequest.Body, stepVariables)
	if err != nil {
		return err
	}
	// check request body format if Content-Type specified as application/json
	// 如果Content-Type为application/json，则检查请求体格式
	if strings.HasPrefix(r.req.Header.Get("Content-Type"), "application/json") {
		switch data.(type) {
		case bool, float64, string, map[string]interface{}, []interface{}, nil:
			break
		default:
			return errors.Errorf("request body type inconsistent with Content-Type: %v",
				r.req.Header.Get("Content-Type"))
		}
	}
	// 将请求体设置到requestMap中，requestMap主要用于测试报告中的展示
	r.requestMap["body"] = data
	// 将data转换为字节数组
	var dataBytes []byte
	switch vv := data.(type) {
	case map[string]interface{}:
		// 获取Content-Type
		contentType := r.req.Header.Get("Content-Type")
		// 根据Content-Type的不同，将请求体转换为不同的格式
		// 如果Content-Type为application/x-www-form-urlencoded，则转换为表单格式,否则转换为json格式
		if strings.HasPrefix(contentType, "application/x-www-form-urlencoded") {
			// post form data
			// 转换为表单格式
			formData := make(url.Values)
			for k, v := range vv {
				formData.Add(k, convertString(v))
			}
			dataBytes = []byte(formData.Encode())
		} else {
			// post json
			// 转换为json格式
			dataBytes, err = json.Marshal(vv)
			if err != nil {
				return err
			}
			// 设置Content-Type
			if contentType == "" {
				r.req.Header.Set("Content-Type", "application/json; charset=utf-8")
			}
		}
	case []interface{}:
		// 如果为[]interface{}类型，则转换为json格式
		contentType := r.req.Header.Get("Content-Type")
		// post json
		dataBytes, err = json.Marshal(vv)
		if err != nil {
			return err
		}
		// 设置Content-Type
		if contentType == "" {
			r.req.Header.Set("Content-Type", "application/json; charset=utf-8")
		}
	case string:
		dataBytes = []byte(vv)
	case []byte:
		dataBytes = vv
	case bytes.Buffer:
		dataBytes = vv.Bytes()
	case *builtin.TFormDataWriter:
		dataBytes = vv.Payload.Bytes()
	default: // unexpected body type
		return errors.New("unexpected request body type")
	}
	// 将字节数组转换成io.ReadCloser对象
	r.req.Body = io.NopCloser(bytes.NewReader(dataBytes))
	// 设置Content-Length
	r.req.ContentLength = int64(len(dataBytes))

	return nil
}

// prepareHeaders 设置content tpe
func initUpload(step *TStep) {
	if step.Request.Headers == nil {
		step.Request.Headers = make(map[string]string)
	}
	step.Request.Headers["Content-Type"] = "${multipart_content_type($m_encoder)}"
	step.Request.Body = "$m_encoder"
}

// prepareUpload 设置上传文件内容
func prepareUpload(parser *Parser, step *TStep, stepVariables map[string]interface{}) (err error) {
	if len(step.Request.Upload) == 0 {
		return
	}
	uploadMap, err := parser.Parse(step.Request.Upload, stepVariables)
	if err != nil {
		return
	}
	stepVariables["m_upload"] = uploadMap
	mEncoder, err := parser.Parse("${multipart_encoder($m_upload)}", stepVariables)
	if err != nil {
		return
	}
	stepVariables["m_encoder"] = mEncoder
	return
}

// runStepRequest 执行请求
func runStepRequest(r *SessionRunner, step *TStep) (stepResult *StepResult, err error) {
	stepResult = &StepResult{
		Name:        step.Name,
		StepType:    stepTypeRequest,
		Success:     false,
		ContentSize: 0,
	}

	// merge step variables with session variables
	// 合并step、运行和config变量
	stepVariables, err := r.ParseStepVariables(step.Variables)
	if err != nil {
		err = errors.Wrap(err, "parse step variables failed")
		return
	}

	defer func() {
		// update testcase summary
		// 更新测试报告的错误状态
		if err != nil {
			stepResult.Attachments = err.Error()
		}
	}()
	// 如果请求体中包含上传文件，则设置上传文件内容
	err = prepareUpload(r.caseRunner.parser, step, stepVariables)
	if err != nil {
		return
	}

	// 创建测试报告内容
	sessionData := newSessionData()
	parser := r.caseRunner.parser
	config := r.caseRunner.parsedConfig

	// 创建请求构建器
	rb := newRequestBuilder(parser, config, step.Request)
	rb.req.Method = strings.ToUpper(string(step.Request.Method))

	// 将params、headers、body的函数、变量等解析为具体的值
	err = rb.prepareUrlParams(stepVariables)
	if err != nil {
		return
	}

	err = rb.prepareHeaders(stepVariables)
	if err != nil {
		return
	}

	err = rb.prepareBody(stepVariables)
	if err != nil {
		return
	}

	// add request object to step variables, could be used in setup hooks
	// 将请求对象添加到step变量中，hook中可以直接使用hrp_step_name、hrp_step_request、request 变量
	stepVariables["hrp_step_name"] = step.Name
	stepVariables["hrp_step_request"] = rb.requestMap
	stepVariables["request"] = rb.requestMap

	// deal with setup hooks
	// 遍历setp hook并执行
	for _, setupHook := range step.SetupHooks {
		// 调用hook
		req, err := parser.Parse(setupHook, stepVariables)
		if err != nil {
			return stepResult, errors.Wrap(err, "run setup hooks failed")
		}
		// 处理hook返回的结果
		reqMap, ok := req.(map[string]interface{})
		if ok && reqMap != nil {
			rb.requestMap = reqMap
			stepVariables["request"] = reqMap
		}
	}
	// 将hook处理后的结果回写到请求对象中
	if len(step.SetupHooks) > 0 {
		// 更新body的内容
		requestBody, ok := rb.requestMap["body"].(map[string]interface{})
		if ok {
			body, err := json.Marshal(requestBody)
			if err == nil {
				rb.req.Body = io.NopCloser(bytes.NewReader(body))
				rb.req.ContentLength = int64(len(body))
			}
		}
		// 更新header的内容
		headers, ok := rb.requestMap["headers"].(map[string]string)
		rb.req.Header = map[string][]string{}
		for key, value := range headers {
			rb.req.Header.Set(key, value)
		}
	}

	// log & print request
	// 打印req内容
	if r.caseRunner.hrpRunner.requestsLogOn {
		if err := printRequest(rb.req); err != nil {
			return stepResult, err
		}
	}

	// stat HTTP request
	// 统计http请求
	var httpStat httpstat.Stat
	if r.caseRunner.hrpRunner.httpStatOn {
		ctx := httpstat.WithHTTPStat(rb.req, &httpStat)
		rb.req = rb.req.WithContext(ctx)
	}

	// select HTTP client
	// 设置http client版本
	var client *http.Client
	if step.Request.HTTP2 {
		client = r.caseRunner.hrpRunner.http2Client
	} else {
		client = r.caseRunner.hrpRunner.httpClient
	}

	// set step timeout
	// 设置超时时间
	if step.Request.Timeout != 0 {
		client.Timeout = time.Duration(step.Request.Timeout*1000) * time.Millisecond
	}

	// do request action
	// 发起http请求
	start := time.Now()
	resp, err := client.Do(rb.req)
	if err != nil {
		return stepResult, errors.Wrap(err, "do request failed")
	}
	if resp != nil {
		// 释放资源
		defer resp.Body.Close()
	}

	// decode response body in br/gzip/deflate formats
	// 解析响应体
	err = decodeResponseBody(resp)
	if err != nil {
		return stepResult, errors.Wrap(err, "decode response body failed")
	}
	// 释放资源
	defer resp.Body.Close()

	// log & print response
	// 打印响应体
	if r.caseRunner.hrpRunner.requestsLogOn {
		if err := printResponse(resp); err != nil {
			return stepResult, err
		}
	}

	// new response object
	// 创建测试报告中的返回对象
	respObj, err := newHttpResponseObject(r.caseRunner.hrpRunner.t, parser, resp)
	if err != nil {
		err = errors.Wrap(err, "init ResponseObject error")
		return
	}
	// 统计请求消耗的时间
	stepResult.Elapsed = time.Since(start).Milliseconds()
	if r.caseRunner.hrpRunner.httpStatOn {
		// resp.Body has been ReadAll
		// 获取http请求的时间
		httpStat.Finish()
		stepResult.HttpStat = httpStat.Durations()
		httpStat.Print()
	}

	// add response object to step variables, could be used in teardown hooks
	// 将响应对象添加到step变量中，hrp_step_response、response 变量
	stepVariables["hrp_step_response"] = respObj.respObjMeta
	stepVariables["response"] = respObj.respObjMeta

	// deal with teardown hooks
	// 遍历teardown hook并执行
	for _, teardownHook := range step.TeardownHooks {
		// 调用hook
		res, err := parser.Parse(teardownHook, stepVariables)
		if err != nil {
			return stepResult, errors.Wrap(err, "run teardown hooks failed")
		}
		// 处理hook返回的结果
		resMpa, ok := res.(map[string]interface{})
		if ok {
			stepVariables["response"] = resMpa
			respObj.respObjMeta = resMpa
		}
	}
	// 将hook处理后的结果回写到响应对象中
	sessionData.ReqResps.Request = rb.requestMap
	sessionData.ReqResps.Response = builtin.FormatResponse(respObj.respObjMeta)

	// extract variables from response
	// 从响应体中提取变量
	extractors := step.Extract
	extractMapping := respObj.Extract(extractors, stepVariables)
	stepResult.ExportVars = extractMapping

	// override step variables with extracted variables
	// 将提取的变量覆盖到step变量中
	stepVariables = mergeVariables(stepVariables, extractMapping)

	// validate response
	// 断言
	err = respObj.Validate(step.Validators, stepVariables)
	// 设置测试报告中的断言结果
	sessionData.Validators = respObj.validationResults
	// 设置测试报告中的断言结果
	if err == nil {
		sessionData.Success = true
		stepResult.Success = true
	}
	stepResult.ContentSize = resp.ContentLength
	stepResult.Data = sessionData

	return stepResult, err
}

// printRequest 打印请求数据
func printRequest(req *http.Request) error {
	reqContentType := req.Header.Get("Content-Type")
	printBody := shouldPrintBody(reqContentType)
	reqDump, err := httputil.DumpRequest(req, printBody)
	if err != nil {
		return errors.Wrap(err, "dump request failed")
	}
	fmt.Println("-------------------- request --------------------")
	reqContent := string(reqDump)
	if reqContentType != "" && !printBody {
		reqContent += fmt.Sprintf("(request body omitted for Content-Type: %v)", reqContentType)
	}
	fmt.Println(reqContent)
	return nil
}

func printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(color.Output, format, a...)
}

// printResponse 打印响应数据
func printResponse(resp *http.Response) error {
	fmt.Println("==================== response ====================")
	connectedVia := "plaintext"
	if resp.TLS != nil {
		switch resp.TLS.Version {
		case tls.VersionTLS12:
			connectedVia = "TLSv1.2"
		case tls.VersionTLS13:
			connectedVia = "TLSv1.3"
		}
	}
	printf("%s %s\n", color.CyanString("Connected via"), color.BlueString("%s", connectedVia))
	respContentType := resp.Header.Get("Content-Type")
	printBody := shouldPrintBody(respContentType)
	respDump, err := httputil.DumpResponse(resp, printBody)
	if err != nil {
		return errors.Wrap(err, "dump response failed")
	}
	respContent := string(respDump)
	if respContentType != "" && !printBody {
		respContent += fmt.Sprintf("(response body omitted for Content-Type: %v)", respContentType)
	}
	fmt.Println(respContent)
	fmt.Println("--------------------------------------------------")
	return nil
}

// decodeResponseBody 解析响应体
func decodeResponseBody(resp *http.Response) (err error) {
	switch resp.Header.Get("Content-Encoding") {
	case "br":
		resp.Body = io.NopCloser(brotli.NewReader(resp.Body))
	case "gzip":
		resp.Body, err = gzip.NewReader(resp.Body)
		if err != nil {
			return err
		}
		resp.ContentLength = -1 // set to unknown to avoid Content-Length mismatched
	case "deflate":
		resp.Body, err = zlib.NewReader(resp.Body)
		if err != nil {
			return err
		}
		resp.ContentLength = -1 // set to unknown to avoid Content-Length mismatched
	}
	return nil
}

// shouldPrintBody return true if the Content-Type is printable
// including text/*, application/json, application/xml, application/www-form-urlencoded
func shouldPrintBody(contentType string) bool {
	if strings.HasPrefix(contentType, "text/") {
		return true
	}
	if strings.HasPrefix(contentType, "application/json") {
		return true
	}
	if strings.HasPrefix(contentType, "application/xml") {
		return true
	}
	if strings.HasPrefix(contentType, "application/x-www-form-urlencoded") {
		return true
	}
	return false
}

// NewStep returns a new constructed teststep with specified step name.
// 创建一个新的测试步骤
func NewStep(name string) *StepRequest {
	return &StepRequest{
		step: &TStep{
			Name:      name,
			Variables: make(map[string]interface{}),
		},
	}
}

// StepRequest 设置测试步骤类型
type StepRequest struct {
	step *TStep
}

// WithVariables 设置当前测试步骤的变量
func (s *StepRequest) WithVariables(variables map[string]interface{}) *StepRequest {
	s.step.Variables = variables
	return s
}

// SetupHook 设置当前测试步骤的setup hook
func (s *StepRequest) SetupHook(hook string) *StepRequest {
	s.step.SetupHooks = append(s.step.SetupHooks, hook)
	return s
}

// HTTP2 设置当前测试步骤的HTTP2协议
func (s *StepRequest) HTTP2() *StepRequest {
	s.step.Request = &Request{
		HTTP2: true,
	}
	return s
}

// Loop 设置当前测试步骤的循环次数
func (s *StepRequest) Loop(times int) *StepRequest {
	s.step.Loops = times
	return s
}

// GET 设置当前测试步骤的请求方法为GET、设置url
func (s *StepRequest) GET(url string) *StepRequestWithOptionalArgs {
	if s.step.Request != nil {
		s.step.Request.Method = httpGET
		s.step.Request.URL = url
	} else {
		s.step.Request = &Request{
			Method: httpGET,
			URL:    url,
		}
	}
	return &StepRequestWithOptionalArgs{
		step: s.step,
	}
}

// HEAD 设置当前测试步骤的请求方法为HEAD、设置url
func (s *StepRequest) HEAD(url string) *StepRequestWithOptionalArgs {
	if s.step.Request != nil {
		s.step.Request.Method = httpHEAD
		s.step.Request.URL = url
	} else {
		s.step.Request = &Request{
			Method: httpHEAD,
			URL:    url,
		}
	}
	return &StepRequestWithOptionalArgs{
		step: s.step,
	}
}

// POST 设置当前测试步骤的请求方法为POST、设置url
func (s *StepRequest) POST(url string) *StepRequestWithOptionalArgs {
	if s.step.Request != nil {
		s.step.Request.Method = httpPOST
		s.step.Request.URL = url
	} else {
		s.step.Request = &Request{
			Method: httpPOST,
			URL:    url,
		}
	}
	return &StepRequestWithOptionalArgs{
		step: s.step,
	}
}

// PUT 设置当前测试步骤的请求方法为PUT、设置url
func (s *StepRequest) PUT(url string) *StepRequestWithOptionalArgs {
	if s.step.Request != nil {
		s.step.Request.Method = httpPUT
		s.step.Request.URL = url
	} else {
		s.step.Request = &Request{
			Method: httpPUT,
			URL:    url,
		}
	}
	return &StepRequestWithOptionalArgs{
		step: s.step,
	}
}

// DELETE 设置当前测试步骤的请求方法为DELETE、设置url
func (s *StepRequest) DELETE(url string) *StepRequestWithOptionalArgs {
	if s.step.Request != nil {
		s.step.Request.Method = httpDELETE
		s.step.Request.URL = url
	} else {
		s.step.Request = &Request{
			Method: httpDELETE,
			URL:    url,
		}
	}
	return &StepRequestWithOptionalArgs{
		step: s.step,
	}
}

// OPTIONS 设置当前测试步骤的请求方法为OPTIONS、设置url
func (s *StepRequest) OPTIONS(url string) *StepRequestWithOptionalArgs {
	if s.step.Request != nil {
		s.step.Request.Method = httpOPTIONS
		s.step.Request.URL = url
	} else {
		s.step.Request = &Request{
			Method: httpOPTIONS,
			URL:    url,
		}
	}
	return &StepRequestWithOptionalArgs{
		step: s.step,
	}
}

// PATCH 设置当前测试步骤的请求方法为PATCH、设置url
func (s *StepRequest) PATCH(url string) *StepRequestWithOptionalArgs {
	if s.step.Request != nil {
		s.step.Request.Method = httpPATCH
		s.step.Request.URL = url
	} else {
		s.step.Request = &Request{
			Method: httpPATCH,
			URL:    url,
		}
	}
	return &StepRequestWithOptionalArgs{
		step: s.step,
	}
}

// CallRefCase 调用一个引用的测试用例。
func (s *StepRequest) CallRefCase(tc ITestCase) *StepTestCaseWithOptionalArgs {
	var err error
	s.step.TestCase, err = tc.ToTestCase()
	if err != nil {
		log.Error().Err(err).Msg("failed to load testcase")
		os.Exit(code.GetErrorCode(err))
	}
	return &StepTestCaseWithOptionalArgs{
		step: s.step,
	}
}

// CallRefAPI 调用一个引用的 API。
func (s *StepRequest) CallRefAPI(api IAPI) *StepAPIWithOptionalArgs {
	var err error
	s.step.API, err = api.ToAPI()
	if err != nil {
		log.Error().Err(err).Msg("failed to load api")
		os.Exit(code.GetErrorCode(err))
	}
	return &StepAPIWithOptionalArgs{
		step: s.step,
	}
}

// StartTransaction 开始一个事务。
func (s *StepRequest) StartTransaction(name string) *StepTransaction {
	s.step.Transaction = &Transaction{
		Name: name,
		Type: transactionStart,
	}
	return &StepTransaction{
		step: s.step,
	}
}

// EndTransaction 结束一个事务。
func (s *StepRequest) EndTransaction(name string) *StepTransaction {
	s.step.Transaction = &Transaction{
		Name: name,
		Type: transactionEnd,
	}
	return &StepTransaction{
		step: s.step,
	}
}

// SetThinkTime 设置思考时间。
func (s *StepRequest) SetThinkTime(time float64) *StepThinkTime {
	s.step.ThinkTime = &ThinkTime{
		Time: time,
	}
	return &StepThinkTime{
		step: s.step,
	}
}

// SetRendezvous 创建一个新的会合点。
func (s *StepRequest) SetRendezvous(name string) *StepRendezvous {
	s.step.Rendezvous = &Rendezvous{
		Name: name,
	}
	return &StepRendezvous{
		step: s.step,
	}
}

// WebSocket 创建一个新的 WebSocket 动作。
func (s *StepRequest) WebSocket() *StepWebSocket {
	s.step.WebSocket = &WebSocketAction{}
	return &StepWebSocket{
		step: s.step,
	}
}

// Android 创建一个新的 Android 动作。
func (s *StepRequest) Android() *StepMobile {
	s.step.Android = &MobileStep{}
	return &StepMobile{
		step: s.step,
	}
}

// IOS 创建一个新的 iOS 动作。
func (s *StepRequest) IOS() *StepMobile {
	s.step.IOS = &MobileStep{}
	return &StepMobile{
		step: s.step,
	}
}

// StepRequestWithOptionalArgs implements IStep interface.
type StepRequestWithOptionalArgs struct {
	step *TStep
}

// SetVerify sets whether to verify SSL for current HTTP request.
func (s *StepRequestWithOptionalArgs) SetVerify(verify bool) *StepRequestWithOptionalArgs {
	log.Info().Bool("verify", verify).Msg("set step request verify")
	s.step.Request.Verify = verify
	return s
}

// SetTimeout sets timeout for current HTTP request.
func (s *StepRequestWithOptionalArgs) SetTimeout(timeout time.Duration) *StepRequestWithOptionalArgs {
	log.Info().Float64("timeout(seconds)", timeout.Seconds()).Msg("set step request timeout")
	s.step.Request.Timeout = timeout.Seconds()
	return s
}

// SetProxies sets proxies for current HTTP request.
func (s *StepRequestWithOptionalArgs) SetProxies(proxies map[string]string) *StepRequestWithOptionalArgs {
	log.Info().Interface("proxies", proxies).Msg("set step request proxies")
	// TODO
	return s
}

// SetAllowRedirects sets whether to allow redirects for current HTTP request.
func (s *StepRequestWithOptionalArgs) SetAllowRedirects(allowRedirects bool) *StepRequestWithOptionalArgs {
	log.Info().Bool("allowRedirects", allowRedirects).Msg("set step request allowRedirects")
	s.step.Request.AllowRedirects = allowRedirects
	return s
}

// SetAuth sets auth for current HTTP request.
func (s *StepRequestWithOptionalArgs) SetAuth(auth map[string]string) *StepRequestWithOptionalArgs {
	log.Info().Interface("auth", auth).Msg("set step request auth")
	// TODO
	return s
}

// WithParams sets HTTP request params for current step.
func (s *StepRequestWithOptionalArgs) WithParams(params map[string]interface{}) *StepRequestWithOptionalArgs {
	s.step.Request.Params = params
	return s
}

// WithHeaders sets HTTP request headers for current step.
func (s *StepRequestWithOptionalArgs) WithHeaders(headers map[string]string) *StepRequestWithOptionalArgs {
	s.step.Request.Headers = headers
	return s
}

// WithCookies sets HTTP request cookies for current step.
func (s *StepRequestWithOptionalArgs) WithCookies(cookies map[string]string) *StepRequestWithOptionalArgs {
	s.step.Request.Cookies = cookies
	return s
}

// WithBody sets HTTP request body for current step.
func (s *StepRequestWithOptionalArgs) WithBody(body interface{}) *StepRequestWithOptionalArgs {
	s.step.Request.Body = body
	return s
}

// WithUpload sets HTTP request body for uploading file(s).
func (s *StepRequestWithOptionalArgs) WithUpload(upload map[string]interface{}) *StepRequestWithOptionalArgs {
	// init upload
	initUpload(s.step)
	s.step.Request.Upload = upload
	return s
}

// TeardownHook adds a teardown hook for current teststep.
func (s *StepRequestWithOptionalArgs) TeardownHook(hook string) *StepRequestWithOptionalArgs {
	s.step.TeardownHooks = append(s.step.TeardownHooks, hook)
	return s
}

// Validate switches to step validation.
func (s *StepRequestWithOptionalArgs) Validate() *StepRequestValidation {
	return &StepRequestValidation{
		step: s.step,
	}
}

// Extract switches to step extraction.
func (s *StepRequestWithOptionalArgs) Extract() *StepRequestExtraction {
	s.step.Extract = make(map[string]string)
	return &StepRequestExtraction{
		step: s.step,
	}
}

func (s *StepRequestWithOptionalArgs) Name() string {
	if s.step.Name != "" {
		return s.step.Name
	}
	return fmt.Sprintf("%v %s", s.step.Request.Method, s.step.Request.URL)
}

func (s *StepRequestWithOptionalArgs) Type() StepType {
	return StepType(fmt.Sprintf("request-%v", s.step.Request.Method))
}

func (s *StepRequestWithOptionalArgs) Struct() *TStep {
	return s.step
}

func (s *StepRequestWithOptionalArgs) Run(r *SessionRunner) (*StepResult, error) {
	return runStepRequest(r, s.step)
}

// StepRequestExtraction implements IStep interface.
type StepRequestExtraction struct {
	step *TStep
}

// WithJmesPath sets the JMESPath expression to extract from the response.
func (s *StepRequestExtraction) WithJmesPath(jmesPath string, varName string) *StepRequestExtraction {
	s.step.Extract[varName] = jmesPath
	return s
}

// Validate switches to step validation.
func (s *StepRequestExtraction) Validate() *StepRequestValidation {
	return &StepRequestValidation{
		step: s.step,
	}
}

func (s *StepRequestExtraction) Name() string {
	return s.step.Name
}

func (s *StepRequestExtraction) Type() StepType {
	if s.step.Request != nil {
		return StepType(fmt.Sprintf("request-%v", s.step.Request.Method))
	}
	if s.step.WebSocket != nil {
		return StepType(fmt.Sprintf("websocket-%v", s.step.WebSocket.Type))
	}
	return "extraction"
}

func (s *StepRequestExtraction) Struct() *TStep {
	return s.step
}

func (s *StepRequestExtraction) Run(r *SessionRunner) (*StepResult, error) {
	if s.step.Request != nil {
		return runStepRequest(r, s.step)
	}
	if s.step.WebSocket != nil {
		return runStepWebSocket(r, s.step)
	}
	return nil, errors.New("unexpected protocol type")
}

// StepRequestValidation implements IStep interface.
type StepRequestValidation struct {
	step *TStep
}

func (s *StepRequestValidation) Name() string {
	if s.step.Name != "" {
		return s.step.Name
	}
	return fmt.Sprintf("%s %s", s.step.Request.Method, s.step.Request.URL)
}

func (s *StepRequestValidation) Type() StepType {
	if s.step.Request != nil {
		return StepType(fmt.Sprintf("request-%v", s.step.Request.Method))
	}
	if s.step.WebSocket != nil {
		return StepType(fmt.Sprintf("websocket-%v", s.step.WebSocket.Type))
	}
	return "validation"
}

func (s *StepRequestValidation) Struct() *TStep {
	return s.step
}

func (s *StepRequestValidation) Run(r *SessionRunner) (*StepResult, error) {
	if s.step.Request != nil {
		return runStepRequest(r, s.step)
	}
	if s.step.WebSocket != nil {
		return runStepWebSocket(r, s.step)
	}
	return nil, errors.New("unexpected protocol type")
}

func (s *StepRequestValidation) AssertEqual(jmesPath string, expected interface{}, msg string) *StepRequestValidation {
	v := Validator{
		Check:   jmesPath,
		Assert:  "equals",
		Expect:  expected,
		Message: msg,
	}
	s.step.Validators = append(s.step.Validators, v)
	return s
}

func (s *StepRequestValidation) AssertGreater(jmesPath string, expected interface{}, msg string) *StepRequestValidation {
	v := Validator{
		Check:   jmesPath,
		Assert:  "greater_than",
		Expect:  expected,
		Message: msg,
	}
	s.step.Validators = append(s.step.Validators, v)
	return s
}

func (s *StepRequestValidation) AssertLess(jmesPath string, expected interface{}, msg string) *StepRequestValidation {
	v := Validator{
		Check:   jmesPath,
		Assert:  "less_than",
		Expect:  expected,
		Message: msg,
	}
	s.step.Validators = append(s.step.Validators, v)
	return s
}

func (s *StepRequestValidation) AssertGreaterOrEqual(jmesPath string, expected interface{}, msg string) *StepRequestValidation {
	v := Validator{
		Check:   jmesPath,
		Assert:  "greater_or_equals",
		Expect:  expected,
		Message: msg,
	}
	s.step.Validators = append(s.step.Validators, v)
	return s
}

func (s *StepRequestValidation) AssertLessOrEqual(jmesPath string, expected interface{}, msg string) *StepRequestValidation {
	v := Validator{
		Check:   jmesPath,
		Assert:  "less_or_equals",
		Expect:  expected,
		Message: msg,
	}
	s.step.Validators = append(s.step.Validators, v)
	return s
}

func (s *StepRequestValidation) AssertNotEqual(jmesPath string, expected interface{}, msg string) *StepRequestValidation {
	v := Validator{
		Check:   jmesPath,
		Assert:  "not_equal",
		Expect:  expected,
		Message: msg,
	}
	s.step.Validators = append(s.step.Validators, v)
	return s
}

func (s *StepRequestValidation) AssertContains(jmesPath string, expected interface{}, msg string) *StepRequestValidation {
	v := Validator{
		Check:   jmesPath,
		Assert:  "contains",
		Expect:  expected,
		Message: msg,
	}
	s.step.Validators = append(s.step.Validators, v)
	return s
}

func (s *StepRequestValidation) AssertTypeMatch(jmesPath string, expected interface{}, msg string) *StepRequestValidation {
	v := Validator{
		Check:   jmesPath,
		Assert:  "type_match",
		Expect:  expected,
		Message: msg,
	}
	s.step.Validators = append(s.step.Validators, v)
	return s
}

func (s *StepRequestValidation) AssertRegexp(jmesPath string, expected interface{}, msg string) *StepRequestValidation {
	v := Validator{
		Check:   jmesPath,
		Assert:  "regex_match",
		Expect:  expected,
		Message: msg,
	}
	s.step.Validators = append(s.step.Validators, v)
	return s
}

func (s *StepRequestValidation) AssertStartsWith(jmesPath string, expected interface{}, msg string) *StepRequestValidation {
	v := Validator{
		Check:   jmesPath,
		Assert:  "startswith",
		Expect:  expected,
		Message: msg,
	}
	s.step.Validators = append(s.step.Validators, v)
	return s
}

func (s *StepRequestValidation) AssertEndsWith(jmesPath string, expected interface{}, msg string) *StepRequestValidation {
	v := Validator{
		Check:   jmesPath,
		Assert:  "endswith",
		Expect:  expected,
		Message: msg,
	}
	s.step.Validators = append(s.step.Validators, v)
	return s
}

func (s *StepRequestValidation) AssertLengthEqual(jmesPath string, expected interface{}, msg string) *StepRequestValidation {
	v := Validator{
		Check:   jmesPath,
		Assert:  "length_equals",
		Expect:  expected,
		Message: msg,
	}
	s.step.Validators = append(s.step.Validators, v)
	return s
}

func (s *StepRequestValidation) AssertContainedBy(jmesPath string, expected interface{}, msg string) *StepRequestValidation {
	v := Validator{
		Check:   jmesPath,
		Assert:  "contained_by",
		Expect:  expected,
		Message: msg,
	}
	s.step.Validators = append(s.step.Validators, v)
	return s
}

func (s *StepRequestValidation) AssertLengthLessThan(jmesPath string, expected interface{}, msg string) *StepRequestValidation {
	v := Validator{
		Check:   jmesPath,
		Assert:  "length_less_than",
		Expect:  expected,
		Message: msg,
	}
	s.step.Validators = append(s.step.Validators, v)
	return s
}

func (s *StepRequestValidation) AssertStringEqual(jmesPath string, expected interface{}, msg string) *StepRequestValidation {
	v := Validator{
		Check:   jmesPath,
		Assert:  "string_equals",
		Expect:  expected,
		Message: msg,
	}
	s.step.Validators = append(s.step.Validators, v)
	return s
}

func (s *StepRequestValidation) AssertEqualFold(jmesPath string, expected interface{}, msg string) *StepRequestValidation {
	v := Validator{
		Check:   jmesPath,
		Assert:  "equal_fold",
		Expect:  expected,
		Message: msg,
	}
	s.step.Validators = append(s.step.Validators, v)
	return s
}

func (s *StepRequestValidation) AssertLengthLessOrEquals(jmesPath string, expected interface{}, msg string) *StepRequestValidation {
	v := Validator{
		Check:   jmesPath,
		Assert:  "length_less_or_equals",
		Expect:  expected,
		Message: msg,
	}
	s.step.Validators = append(s.step.Validators, v)
	return s
}

func (s *StepRequestValidation) AssertLengthGreaterThan(jmesPath string, expected interface{}, msg string) *StepRequestValidation {
	v := Validator{
		Check:   jmesPath,
		Assert:  "length_greater_than",
		Expect:  expected,
		Message: msg,
	}
	s.step.Validators = append(s.step.Validators, v)
	return s
}

func (s *StepRequestValidation) AssertLengthGreaterOrEquals(jmesPath string, expected interface{}, msg string) *StepRequestValidation {
	v := Validator{
		Check:   jmesPath,
		Assert:  "length_greater_or_equals",
		Expect:  expected,
		Message: msg,
	}
	s.step.Validators = append(s.step.Validators, v)
	return s
}

// Validator represents validator for one HTTP response.
type Validator struct {
	Check   string      `json:"check" yaml:"check"` // get value with jmespath
	Assert  string      `json:"assert" yaml:"assert"`
	Expect  interface{} `json:"expect" yaml:"expect"`
	Message string      `json:"msg,omitempty" yaml:"msg,omitempty"` // optional
}
