package runTestCase

import (
	"encoding/json"
	"fmt"

	"go.uber.org/zap"

	"github.com/test-instructor/yangfan/hrp"
	"github.com/test-instructor/yangfan/server/global"
	mic "github.com/test-instructor/yangfan/server/model/interfacecase"
)

func yangfanTestCaseToHrpCase(testCaseList []mic.HrpCase, debugTalkFilePath string, tcm *ApisCaseModel) error {

	for _, testCase := range testCaseList {
		apiConfigJson, _ := json.Marshal(testCase.Confing)
		var tConfig hrp.TConfig
		err := json.Unmarshal(apiConfigJson, &tConfig)
		if err != nil {
			global.GVA_LOG.Error("用例转换出现错误", zap.Error(err))
		}
		global.GVA_LOG.Debug(fmt.Sprintf("用例id", testCase.ID))
		global.GVA_LOG.Debug("case name" + testCase.Name)
		toTestCase := ToTestCase{TestSteps: testCase.TestSteps, Config: tcm.Config}
		caseJson, _ := json.Marshal(toTestCase)
		global.GVA_LOG.Debug("测试用例json格式")
		global.GVA_LOG.Debug("\n" + string(caseJson))
		tConfig.Path = debugTalkFilePath
		tc := &hrp.JsonToCase{
			JsonString:        string(caseJson),
			ID:                testCase.ID,
			DebugTalkFilePath: debugTalkFilePath,
			Name:              testCase.Name,
			Config:            &tConfig,
		}
		testCase, _ := tc.ToTestCase()
		tcm.Case = append(tcm.Case, testCase)
	}
	return nil
}
