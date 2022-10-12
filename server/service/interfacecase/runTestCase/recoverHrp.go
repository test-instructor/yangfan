package runTestCase

import (
	"fmt"
	"github.com/test-instructor/cheetah/server/global"
)

func recoverHrp(reportOperation *ReportOperation) {
	if msg := recover(); msg != nil {
		global.GVA_LOG.Error("测试用例运行时报错：")
		global.GVA_LOG.Error(fmt.Sprintln(msg))
		reportOperation.Recover(fmt.Sprintln(msg))
	}
}
