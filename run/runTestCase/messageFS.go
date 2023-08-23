package runTestCase

import (
	"time"

	"github.com/test-instructor/yangfan/proto/run"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	"go.uber.org/zap"
)

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
	err := fn.SendMessage(body, fn.msg)

	return err
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
