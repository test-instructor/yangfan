package runTestCase

import (
	"fmt"
	"time"

	"github.com/test-instructor/yangfan/proto/run"
	"github.com/test-instructor/yangfan/server/global"
	"go.uber.org/zap"
)

type DingTalkNotifier struct {
	NotifierDefault
	Type run.NotifierType
}

func (dn DingTalkNotifier) Send() error {
	if dn.reports.Success != nil && *dn.reports.Success && dn.msg.GetFail() {
		return nil
	}
	global.GVA_LOG.Debug("Sending DingTalk message:")
	body := make(map[string]interface{})
	body["msgtype"] = "actionCard"

	if dn.msg.GetSignature() != "" {
		timestamp := time.Now().UnixMilli()
		sign, err := dn.genSignDingTalk(dn.msg.GetSignature(), timestamp)
		if err != nil {
			global.GVA_LOG.Error("签名失败", zap.Error(err))
			return err
		}
		dn.msg.Webhook += fmt.Sprintf("&timestamp=%d&sign=%s", timestamp, sign)
	}
	var actionCard = make(map[string]interface{})
	var success = "失败"
	if *dn.reports.Success {
		success = "成功"
	}
	title := fmt.Sprintf("【%s】%s | %s", dn.reports.Name, success, dn.reports.ApiEnvName)
	actionCard["title"] = title
	if dn.Type == run.NotifierType_Dingtalk {
		text := fmt.Sprintf("# <font color=#FF0000>%s</font>\n\n", title)
		if *dn.reports.Success {
			text = fmt.Sprintf("# <font color=#0000FF>%s</font>\n\n", title)
		}
		data := dn.getCard()
		text += dn.generateTableContent(data.Data.TemplateVariable.Content)
		actionCard["text"] = text
	}
	if dn.Type == run.NotifierType_DingtalkText {
		text := fmt.Sprintf("# <font color=#FF0000>%s</font>\n\n", title)
		if *dn.reports.Success {
			text = fmt.Sprintf("# <font color=#0000FF>%s</font>\n\n", title)
		}
		card := dn.getCard()
		for _, content := range card.Data.TemplateVariable.Content {
			text += fmt.Sprintf("用例：%s,成功用例：%d,失败用例：%d,耗时：%ds秒\n\n", content.Name, content.Success, content.Fail, content.Time)
		}
		actionCard["text"] = text
	}

	btn := make(map[string]interface{})
	btn["title"] = "查看详情"
	btn["actionURL"] = dn.getReportUrl()
	actionCard["btns"] = []interface{}{btn}
	actionCard["btnOrientation"] = "0"
	actionCard["hideAvatar"] = "0"
	actionCard["btnBackgroundColor"] = "#FF9900"
	body["actionCard"] = actionCard

	err := dn.SendMessage(body)
	return err
}
