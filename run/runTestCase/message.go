package runTestCase

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/test-instructor/yangfan/proto/run"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	"go.uber.org/zap"
)

type FSTemplate string

const FSTemplateFail FSTemplate = "ctp_AAmIkXEaoPea"
const FSTemplateSuccess FSTemplate = "ctp_AAmjuIxqEQjK"

type Notifier interface {
	Send() error
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

type WeChatNotifier struct {
	msg     *run.Msg
	reports *interfacecase.ApiReport
}

func (wn WeChatNotifier) Send() error {
	// 实现企业微信消息发送逻辑
	fmt.Println("Sending WeChat message:")
	return nil
}

type DingTalkNotifier struct {
	msg     *run.Msg
	reports *interfacecase.ApiReport
	NotifierDefault
}

func (dn DingTalkNotifier) Send() error {
	if dn.reports.Success != nil && *dn.reports.Success && dn.msg.GetFail() {
		return nil
	}
	global.GVA_LOG.Debug("Sending DingTalk message:")
	body := make(map[string]interface{})
	body["msgtype"] = "actionCard"
	url := dn.msg.GetWebhook()
	if dn.msg.GetSignature() != "" {
		timestamp := time.Now().UnixMilli()
		sign, err := dn.genSignDingTalk(dn.msg.GetSignature(), timestamp)
		if err != nil {
			global.GVA_LOG.Error("签名失败", zap.Error(err))
			return err
		}
		url += fmt.Sprintf("&timestamp=%d&sign=%s", timestamp, sign)
	}
	var actionCard = make(map[string]interface{})
	var success = "失败"
	if *dn.reports.Success {
		success = "成功"
	}
	title := fmt.Sprintf("【%s】%s | %s", dn.reports.Name, success, dn.reports.ApiEnvName)
	actionCard["title"] = title
	text := fmt.Sprintf("# <font color=#FF0000>%s</font>\n\n", title)
	if *dn.reports.Success {
		text = fmt.Sprintf("# <font color=#0000FF>%s</font>\n\n", title)
	}
	data := dn.getCard(dn.reports)
	text += dn.generateTableContent(data.Data.TemplateVariable.Content)
	actionCard["text"] = text

	btn := make(map[string]interface{})
	btn["title"] = "查看详情"
	btn["actionURL"] = ""
	actionCard["btns"] = []interface{}{btn}
	actionCard["btnOrientation"] = "0"
	actionCard["hideAvatar"] = "0"
	actionCard["btnBackgroundColor"] = "#FF9900"
	body["actionCard"] = actionCard

	contentType := "application/json"
	reqJSON, err := json.Marshal(body)
	if err != nil {
		global.GVA_LOG.Error("Error marshaling JSON:", zap.Error(err))
		return err
	}
	// 发起 POST 请求
	global.GVA_LOG.Debug("请求数据", zap.Any("reqJSON", string(reqJSON)))
	bodyByte := bytes.NewBuffer(reqJSON)
	req, err := http.NewRequest(http.MethodPost, url, bodyByte)
	if err != nil {
		global.GVA_LOG.Error("Error creating request:", zap.Error(err))
	}
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		global.GVA_LOG.Error("Error sending POST request:", zap.Error(err))
	}
	defer resp.Body.Close()

	global.GVA_LOG.Debug("请求数据", zap.Any("header", resp.Request.Header))
	global.GVA_LOG.Debug("响应数据", zap.Any("Status", resp.Status))
	return nil
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

type FeishuNotifier struct {
	msg     *run.Msg
	reports *interfacecase.ApiReport
	NotifierDefault
}

func (fn FeishuNotifier) Send() error {
	// 实现飞书消息发送逻辑
	if fn.reports.Success != nil && *fn.reports.Success && fn.msg.GetFail() {
		return nil
	}
	global.GVA_LOG.Debug("Sending Feishu message:")
	body := make(map[string]interface{})
	body["msg_type"] = "interactive"
	if fn.msg.GetSignature() != "" {
		timestamp := time.Now().Unix()
		sign, err := fn.genSign(fn.msg.GetSignature(), timestamp)
		if err != nil {
			global.GVA_LOG.Error("签名失败", zap.Error(err))
			return err
		}
		body["sign"] = sign
		body["timestamp"] = timestamp
	}
	body["card"] = fn.getCard(fn.reports)

	contentType := "application/json"
	reqJSON, err := json.Marshal(body)
	if err != nil {
		global.GVA_LOG.Error("Error marshaling JSON:", zap.Error(err))
		return err
	}
	// 发起 POST 请求
	global.GVA_LOG.Debug("请求数据", zap.Any("reqJSON", string(reqJSON)))

	bodyByte := bytes.NewBuffer(reqJSON)
	req, err := http.NewRequest(http.MethodPost, fn.msg.GetWebhook(), bodyByte)
	if err != nil {
		global.GVA_LOG.Error("Error creating request:", zap.Error(err))
	}
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		global.GVA_LOG.Error("Error sending POST request:", zap.Error(err))
	}
	defer resp.Body.Close()

	global.GVA_LOG.Debug("请求数据", zap.Any("header", resp.Request.Header))
	global.GVA_LOG.Debug("响应数据", zap.Any("Status", resp.Status))

	return nil
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
