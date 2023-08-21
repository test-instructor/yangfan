package runTestCase

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
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
}

func (dn DingTalkNotifier) Send() error {
	// 实现钉钉消息发送逻辑
	fmt.Println("Sending DingTalk message:")
	return nil
}

type FSCard struct {
	Type string `json:"type"`
	Data FSData `json:"data"`
}

type FSData struct {
	TemplateID       FSTemplate       `json:"template_id"`
	TemplateVariable TemplateVariable `json:"template_variable"`
}
type Content struct {
	Name    string      `json:"name"`
	Success interface{} `json:"success"`
	Fail    interface{} `json:"fail"`
	Time    interface{} `json:"time"`
	Total   interface{} `json:"total"`
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
}

func (fn FeishuNotifier) Send() error {
	// 实现飞书消息发送逻辑
	if fn.reports.Success != nil && *fn.reports.Success && fn.msg.GetFail() {
		return nil
	}
	fmt.Println("Sending Feishu message:")
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
	body["card"] = fn.getCard()

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

func (fn FeishuNotifier) getCard() FSCard {
	var templateID FSTemplate
	if fn.reports.Success != nil && *fn.reports.Success {
		templateID = FSTemplateSuccess
	} else {
		templateID = FSTemplateFail
	}
	data := FSData{
		TemplateID: templateID,
		TemplateVariable: TemplateVariable{
			Env:    fn.reports.ApiEnvName,
			Detail: "",
			Title:  fn.reports.Name,
		},
	}
	var contents []Content
	for _, v := range fn.reports.Details {
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
			Total:   statMap["total"],
			Fail:    statMap["failures"],
			Success: statMap["successes"],
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

func (fn FeishuNotifier) genSign(secret string, timestamp int64) (string, error) {
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
