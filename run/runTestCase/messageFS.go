package runTestCase

import (
	"time"

	"github.com/test-instructor/yangfan/proto/run"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	"go.uber.org/zap"
)

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
	err := fn.SendMessage(body, fn.msg, fn.reports.ProjectID)

	return err
}
