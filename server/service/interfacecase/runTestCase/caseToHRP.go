package runTestCase

import (
	"encoding/json"
	"fmt"
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/hrp"
	mic "github.com/test-instructor/cheetah/server/model/interfacecase"
)

func cheetahCaseToHrpCase(config mic.ApiConfig, testCaseList []mic.ApiCaseStep, debugTalkFilePath string, tcm *ApisCaseModel) error {
	apiConfig_json, _ := json.Marshal(config)
	var tConfig hrp.TConfig
	json.Unmarshal(apiConfig_json, &tConfig)
	for _, testCase := range testCaseList {
		fmt.Println("用例id", testCase.ID)
		fmt.Println("case name", testCase.Name)
		toTestCase := ToTestCase{TestSteps: testCase.TStep}
		caseJson, _ := json.Marshal(toTestCase)
		global.GVA_LOG.Debug("测试用例json格式")
		global.GVA_LOG.Debug("\n" + string(caseJson))
		tc := &hrp.TestCaseJson{
			JsonString:        string(caseJson),
			ID:                testCase.ID,
			DebugTalkFilePath: debugTalkFilePath,
			Config:            &tConfig,
			Name:              testCase.Name,
		}
		testCase, _ := tc.ToTestCase()
		tcm.Case = append(tcm.Case, testCase)
	}
	return nil
}

func cheetahTaskToHrpCase(testCaseList []ApiCaseStep, debugTalkFilePath string, tcm *ApisCaseModel) error {
	for _, testCase := range testCaseList {
		var tConfig hrp.TConfig
		apiConfigJson, _ := json.Marshal(testCase.Config)
		json.Unmarshal(apiConfigJson, &tConfig)
		fmt.Println("用例id", testCase.ID)
		fmt.Println("case name", testCase.Name)
		toTestCase := ToTestCase{TestSteps: testCase.TStep}
		caseJson, _ := json.Marshal(toTestCase)
		global.GVA_LOG.Debug("测试用例json格式")
		global.GVA_LOG.Debug("\n" + string(caseJson))
		tc := &hrp.TestCaseJson{
			JsonString:        string(caseJson),
			ID:                testCase.ID,
			DebugTalkFilePath: debugTalkFilePath,
			Config:            &tConfig,
			Name:              testCase.Name,
		}
		testCase, _ := tc.ToTestCase()
		tcm.Case = append(tcm.Case, testCase)
	}
	return nil
}
