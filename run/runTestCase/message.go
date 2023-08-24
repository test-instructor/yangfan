package runTestCase

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/chromedp/chromedp/device"
	"github.com/test-instructor/yangfan/proto/run"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	"go.uber.org/zap"
)

type FSTemplate string

const FSTemplateFail FSTemplate = "ctp_AAmIkXEaoPea"
const FSTemplateSuccess FSTemplate = "ctp_AAmjuIxqEQjK"

type Devices []Device
type Device struct {
	id   int
	info device.Info
}
type Notifier interface {
	Send() error
}
type Image struct {
	Base64 string `json:"base64"`
	MD5    string `json:"md5"`
}
type FSCard struct {
	Type string `json:"type"`
	Data FSData `json:"data"`
	NotifierDefault
}
type FSData struct {
	TemplateID       FSTemplate       `json:"template_id"`
	TemplateVariable TemplateVariable `json:"template_variable"`
}
type Content struct {
	Name    string `json:"name"`
	Success int    `json:"success"`
	Fail    int    `json:"fail"`
	Time    int    `json:"time"`
	Total   int    `json:"total"`
}
type TemplateVariable struct {
	Env     string    `json:"env"`
	Detail  string    `json:"detail"`
	Content []Content `json:"content"`
	Title   string    `json:"title"`
}

func init() {
	t := reflect.TypeOf(device.Reset)
	defer func() { recover() }() // exit loop if "index-out-of-range"
	for i := 1; ; i++ {
		infoType := reflect.New(t)
		infoType.Elem().SetInt(int64(i))
		devices = append(devices, Device{
			id:   i,
			info: infoType.MethodByName("Device").Interface().(func() device.Info)(),
		})
	}
}

func jsonStringify(s string) string {
	b, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return string(b)
}

var devices Devices

func (d *Device) Device() device.Info {
	return d.info
}

func (d *Device) Set(i string) error { // flag.Value
	_d := devices.NewDeviceByName(i)
	if _d == nil {
		n, _ := strconv.Atoi(i)
		_d = devices.NewDeviceById(n)
	}
	if _d == nil {
		query, _ := url.ParseQuery(i)
		for key := range query {
			value := query.Get(key)
			switch strings.ToLower(key) {
			case "name":
				d.info.Name = value
			case "useragent":
				d.info.UserAgent = value
			case "width":
				n, _ := strconv.Atoi(value)
				d.info.Width = int64(n)
			case "height":
				n, _ := strconv.Atoi(value)
				d.info.Height = int64(n)
			case "scale":
				n, _ := strconv.ParseFloat(value, 64)
				d.info.Scale = n
			case "landscape":
				d.info.Landscape = value == "true"
			case "mobile":
				d.info.Mobile = value == "true"
			case "touch":
				d.info.Touch = value == "true"
			}
		}
	} else {
		d.updateFrom(_d)
	}
	return nil
}

func (d *Device) String() string { // flag.Value
	return ""
}

func (d *Device) MultilineString() string {
	return "Name: " + d.info.Name + "\n" +
		"User-Agent: " + d.info.UserAgent + "\n" +
		"Width: " + strconv.Itoa(int(d.info.Width)) + "\n" +
		"Height: " + strconv.Itoa(int(d.info.Height)) + "\n" +
		"Scale: " + strconv.FormatFloat(d.info.Scale, 'f', -1, 64) + "\n" +
		"Landscape: " + strconv.FormatBool(d.info.Landscape) + "\n" +
		"Mobile: " + strconv.FormatBool(d.info.Mobile) + "\n" +
		"Touch: " + strconv.FormatBool(d.info.Touch)
}

func (d *Device) MultilineStringIndent(n int) string {
	return regexp.MustCompile("(?m)^").ReplaceAllString(d.MultilineString(), strings.Repeat(" ", n))
}

func (d *Device) updateFrom(o *Device) {
	d.info.Name = o.info.Name
	d.info.UserAgent = o.info.UserAgent
	d.info.Width = o.info.Width
	d.info.Height = o.info.Height
	d.info.Scale = o.info.Scale
	d.info.Landscape = o.info.Landscape
	d.info.Mobile = o.info.Mobile
	d.info.Touch = o.info.Touch
}

func (d Devices) NewDeviceByName(name string) *Device {
	for _, _d := range d {
		if _d.info.Name == name {
			d := &Device{}
			d.updateFrom(&_d)
			return d
		}
	}
	return nil
}

func (d Devices) NewDeviceById(i int) *Device {
	for _, _d := range d {
		if _d.id == i {
			d := &Device{}
			d.updateFrom(&_d)
			return d
		}
	}
	return nil
}

func (d Devices) String() string {
	var ss []string
	for _, d := range d {
		s := fmt.Sprintf("%2d  %-40s  %-10s  %0.2fx", d.id, d.info.Name, fmt.Sprintf("%dx%d", d.info.Width, d.info.Height), d.info.Scale)
		ss = append(ss, s)
	}
	return strings.Join(ss, "\n")
}

type NotifierDefault struct{}

func (NotifierDefault) genSign(secret string, timestamp int64) (string, error) {
	//timestamp + key 做sha256, 再进行base64 encode
	stringToSign := fmt.Sprintf("%v", timestamp) + "\n" + secret
	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	_, err := h.Write(data)
	if err != nil {
		return "", err
	}
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature, nil
}

func (NotifierDefault) genSignDingTalk(secret string, timestamp int64) (string, error) {
	//timestamp + key 做sha256, 再进行base64 encode
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, secret)
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(stringToSign))
	signData := mac.Sum(nil)
	signature := url.QueryEscape(base64.StdEncoding.EncodeToString(signData))
	return signature, nil
}

func (NotifierDefault) getCard(reports *interfacecase.ApiReport) FSCard {
	var templateID FSTemplate
	if reports.Success != nil && *reports.Success {
		templateID = FSTemplateSuccess
	} else {
		templateID = FSTemplateFail
	}
	data := FSData{
		TemplateID: templateID,
		TemplateVariable: TemplateVariable{
			Env:    reports.ApiEnvName,
			Detail: "",
			Title:  reports.Name,
		},
	}
	var contents []Content
	for _, v := range reports.Details {
		var statMap, timerMap map[string]interface{}
		err := json.Unmarshal(v.Stat, &statMap)
		if err != nil {
			fmt.Println("Error converting JSON to map:", err)
			global.GVA_LOG.Error("[getCard]用例运行数据出错", zap.Error(err))
			continue
		}
		err = json.Unmarshal(v.Time, &timerMap)
		if err != nil {
			global.GVA_LOG.Error("[getCard]用例时间数据出错", zap.Error(err))
			continue
		}
		durationValue, ok := timerMap["duration"].(float64)
		if !ok {
			global.GVA_LOG.Error("[getCard]运行时间转换成整数时报错", zap.Error(err))
			continue
		}
		content := Content{
			Name:    v.Name,
			Total:   int(statMap["total"].(float64)),
			Fail:    int(statMap["failures"].(float64)),
			Success: int(statMap["successes"].(float64)),
			Time:    int(durationValue),
		}
		contents = append(contents, content)
	}
	data.TemplateVariable.Content = contents
	card := FSCard{
		Type: "template",
		Data: data,
	}
	return card
}

func (NotifierDefault) generateTableContent(data []Content) (tableContent string) {

	for _, row := range data {
		tableContent += "<tr>"
		tableContent += fmt.Sprintf("<td>%s</td>", row.Name)
		tableContent += fmt.Sprintf("<td>%d</td>", row.Total)
		tableContent += fmt.Sprintf("<td>%d</td>", row.Fail)
		tableContent += fmt.Sprintf("<td>%d秒</td>", row.Time)
		tableContent += "</tr>"
	}
	tableContent = fmt.Sprintf("<table><tr><th style=\"min-width: 80px; max-width: 240px;\">用例名称</th>"+
		"<th style=\"min-width: 80px; max-width: 240px;\">成功数</th>"+
		"<th style=\"min-width: 80px; max-width: 240px;\">失败数</th>"+
		"<th style=\"min-width: 80px; max-width: 240px;\">耗时</th></tr>%s</table>", tableContent)

	return tableContent
}

func (n NotifierDefault) SendMessage(body interface{}, msg *run.Msg, projectID uint) error {
	reqJSON, err := json.Marshal(body)
	if err != nil {
		global.GVA_LOG.Error("Error marshaling JSON:", zap.Error(err))
	}

	bodyByte := bytes.NewBuffer(reqJSON)
	req, err := http.NewRequest(http.MethodPost, msg.Webhook, bodyByte)
	if err != nil {
		global.GVA_LOG.Error("Error creating request:", zap.Error(err))
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		global.GVA_LOG.Error("Error sending POST request:", zap.Error(err))
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var responseBody map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		global.GVA_LOG.Error("Error decoding JSON response body:", zap.Error(err))
		return err
	}

	// 将内容转换为 JSON 格式的字符串
	jsonResponse, err := json.MarshalIndent(responseBody, "", "    ")
	if err != nil {
		global.GVA_LOG.Error("Error encoding JSON response body:", zap.Error(err))
		return err
	}
	global.GVA_LOG.Debug(string(jsonResponse))
	n.msgLog(msg, resp.StatusCode, projectID, string(jsonResponse))
	return nil
}

func (NotifierDefault) msgLog(msg *run.Msg, status int, projectID uint, respMessage string) {
	var msgLog = interfacecase.ApiMessageLog{}
	msgLog.ProjectID = projectID
	msgLog.StatusCode = status
	msgLog.ApiMessageID = uint(msg.Id)
	msgLog.Message = respMessage
	if status == 200 {
		msgLog.Status = true
	}
	global.GVA_DB.Create(&msgLog)

}

func NewNotifier(msg *run.Msg, reports *interfacecase.ApiReport) Notifier {
	switch msg.GetType() {
	case run.NotifierType_Wechat:
		return WeChatNotifier{msg: msg, reports: reports}
	case run.NotifierType_Dingtalk:
		return DingTalkNotifier{msg: msg, reports: reports}
	case run.NotifierType_Feishu:
		return FeishuNotifier{msg: msg, reports: reports}
	default:
		return nil
	}
}
