package runTestCase

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
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

type Image struct {
	Base64 string `json:"base64"`
	MD5    string `json:"md5"`
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
	NotifierDefault
}

func (wn WeChatNotifier) Send() error {
	// 实现企业微信消息发送逻辑
	data := wn.getCard(wn.reports)
	htmlContent := wn.generateTableContent(data.Data.TemplateVariable.Content)
	outputPath := "output1.png"
	height := wn.getImageSize(data.Data.TemplateVariable.Content)
	err := wn.html2png(htmlContent, outputPath, height)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		contentType := "application/json"
		body := make(map[string]interface{})
		body["msgtype"] = "image"
		body["image"] = wn.getImage(outputPath)
		reqJSON, err := json.Marshal(body)
		if err != nil {
			global.GVA_LOG.Error("Error marshaling JSON:", zap.Error(err))
		}
		// 发起 POST 请求
		global.GVA_LOG.Debug("请求数据", zap.Any("reqJSON", string(reqJSON)))

		bodyByte := bytes.NewBuffer(reqJSON)
		req, err := http.NewRequest(http.MethodPost, wn.msg.Webhook, bodyByte)
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
		fmt.Println(resp.Status)

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
		fmt.Println(string(jsonResponse))
	}
	return nil
}

func (wn WeChatNotifier) html2png(html string, path string, height int) error {
	var deviceToUse = devices.NewDeviceByName("iPhone 8")

	ctx := context.Background()
	ctx, cancel := chromedp.NewContext(ctx)
	defer cancel()

	err := chromedp.Run(
		ctx,
		chromedp.Emulate(deviceToUse),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// Load the provided HTML content using JavaScript
			loadHTML := fmt.Sprintf(`document.open(); document.write(%s); document.close();`, jsonStringify(html))
			return chromedp.Evaluate(loadHTML, nil).Do(ctx)
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// Capture screenshot
			_, _, _, _, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}
			var buf []byte

			width, heights := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))
			err = emulation.SetDeviceMetricsOverride(width, heights, 1, false).Do(ctx)
			if err != nil {
				return err
			}
			buf, err = page.CaptureScreenshot().WithClip(&page.Viewport{
				X:      contentSize.X,
				Y:      contentSize.Y,
				Width:  float64(525),
				Height: float64(height),
				Scale:  deviceToUse.info.Scale,
			}).Do(ctx)
			//	第一行50，第二行90

			if err != nil {
				return err
			}
			return ioutil.WriteFile(path, buf, 0644)
		}),
	)

	return err
}

func (wn WeChatNotifier) generateTableContent(data []Content) (tableContent string) {

	for _, row := range data {
		tableContent += "<tr>"
		tableContent += fmt.Sprintf("<td>%s</td>", row.Name)
		tableContent += fmt.Sprintf("<td>%d</td>", row.Total)
		tableContent += fmt.Sprintf("<td>%d</td>", row.Fail)
		tableContent += fmt.Sprintf("<td>%d秒</td>", row.Time)
		tableContent += "</tr>"
	}
	tableContent = fmt.Sprintf("<!DOCTYPE html><html><head><style>body {margin: 20; padding: 0;}"+
		"table {font-size: 15;border-collapse: collapse;}"+
		"th, td {border: 1px solid black;padding: 8px;text-align: center;}</style></head><body><table>"+
		"<tr><th style=\"width: 200px;\">用例名称</th><th style=\"width: 80px;\">成功数</th>"+
		"<th style=\"width: 80px;\">失败数</th><th style=\"width: 80px;\">耗时</th></tr>%s"+
		"</table></body></html>", tableContent)
	return tableContent
}

func (wn WeChatNotifier) StringWidth(s string) int {
	length := 0
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		if r == utf8.RuneError {
			// Invalid rune found, skip it
			s = s[size:]
			continue
		}
		if utf8.RuneLen(r) == 1 {
			length++
		} else {
			length += 2
		}
		s = s[size:]
	}
	return length
}

func (wn WeChatNotifier) getImageSize(datas []Content) (height int) {
	height = 55 + 38*len(datas)
	for _, data := range datas {
		wn.StringWidth(data.Name)
		lineCount := int(math.Ceil(float64(wn.StringWidth(data.Name))/float64(24))) - 1
		height += lineCount * 21
	}
	return
}

func (wn WeChatNotifier) getImage(filePath string) (image Image) {
	// 修改为实际图片的路径
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 读取图片内容
	imageData, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 计算图片内容的MD5哈希
	md5Hash := md5.Sum(imageData)
	md5HashString := fmt.Sprintf("%x", md5Hash)

	fmt.Println("MD5 hash of image data:")
	fmt.Println(md5HashString)

	// 将图片内容转换为Base64编码
	base64Data := base64.StdEncoding.EncodeToString(imageData)

	//fmt.Println("Base64 image data:")
	//fmt.Println(base64Data)

	image.Base64 = base64Data
	image.MD5 = md5HashString
	return
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
