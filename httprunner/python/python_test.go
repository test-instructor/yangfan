package python

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/test-instructor/yangfan/httprunner/hrp"
)

func TestPluginInit(t *testing.T) {
	pyScriptPath := "./code"
	ip, err := InitPlugin(pyScriptPath, "", true)
	if err != nil {
		fmt.Println(err)
		fmt.Println("init plugin failed")
		return
	} else {
		fmt.Println("init plugin success")
		_ = ip
	}
	defer func() {
		err = ip.Quit()
		if err != nil {
			fmt.Println(err.Error(), "==================")
		}
	}()

	var parser = hrp.NewParser()
	parser.Plugin = ip
	variables := map[string]interface{}{
		"version": "1.0.0",
		"request": map[string]interface{}{
			"method": "GET",
			"path":   "/test",
		},
	}
	getUserAgent := "${get_user_agent()}"
	res, err := parser.Parse(getUserAgent, variables)
	if err != nil {
		fmt.Println("err=====", err)
	}
	fmt.Println("res=====", res)
	setRequest := "${set_request($request)}"
	res, err = parser.Parse(setRequest, variables)
	if err != nil {
		fmt.Println("err=====", err)
		return
	}
	jsonBytes, err := json.Marshal(res)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	jsonStr := string(jsonBytes)
	fmt.Println("request========", jsonStr)
}
