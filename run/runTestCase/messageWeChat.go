package runTestCase

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"unicode/utf8"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/test-instructor/yangfan/proto/run"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	"go.uber.org/zap"
)

type WeChatNotifier struct {
	msg     *run.Msg
	reports *interfacecase.ApiReport
	NotifierDefault
}

func (wn WeChatNotifier) Send() error {
	// 实现企业微信消息发送逻辑
	if wn.reports.Success != nil && *wn.reports.Success && wn.msg.GetFail() {
		return nil
	}
	data := wn.getCard(wn.reports)
	htmlContent := wn.generateTableContent(data.Data.TemplateVariable.Content)
	outputPath := "output1.png"
	height := wn.getImageSize(data.Data.TemplateVariable.Content)
	err := wn.html2png(htmlContent, outputPath, height)
	if err != nil {
		global.GVA_LOG.Error("Error:", zap.Error(err))
	} else {
		body := make(map[string]interface{})
		body["msgtype"] = "image"
		body["image"] = wn.getImage(outputPath)
		err = wn.SendMessage(body, wn.msg, wn.reports.ProjectID)
	}
	return err
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
		global.GVA_LOG.Error("Error opening file:", zap.Error(err))
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// 读取图片内容
	imageData, err := ioutil.ReadAll(file)
	if err != nil {
		global.GVA_LOG.Error("Error reading file:", zap.Error(err))
		return
	}

	// 计算图片内容的MD5哈希
	md5Hash := md5.Sum(imageData)
	image.MD5 = fmt.Sprintf("%x", md5Hash)

	// 将图片内容转换为Base64编码
	image.Base64 = base64.StdEncoding.EncodeToString(imageData)
	return
}
