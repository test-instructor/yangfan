package runTestCase

import (
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/platform/request"
	"go.uber.org/zap"
)

func Entry(req request.RunnerRequest, msg *string) (reports interface{}, err error) {
	var tc TestCase
	switch req.CaseType {
	case "接口", "api":
		tc = NewRunApi(req, msg)
	case "用例", "case":
		tc = NewRunCase(req, msg)
	case "任务", "task":
		tc = NewRunTask(req, msg)
	case "步骤", "step":
		tc = NewRunStep(req, msg)
	case "标签", "tag":
		tc = NewRunTag(req, msg)
	default:
		global.GVA_LOG.Warn("Unknown CaseType", zap.String("type", req.CaseType))
		tc = NewRunCase(req, msg)
	}
	return RunTestCase(tc)
}
