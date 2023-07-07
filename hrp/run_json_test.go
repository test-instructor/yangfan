package hrp

import (
	"testing"
)

/*func TestJsonRunner(t *testing.T) {
	jsonString := `{
    "config": {
        "name": "demo with complex mechanisms",
        "base_url": "https://postman-echo.com",
        "variables": {
            "a": "${sum(10, 2.3)}",
            "b": 3.45,
            "n": "${sum_ints(1, 2, 2)}",
            "varFoo1": "${gen_random_string($n)}",
            "varFoo2": "${max($a, $b)}"
        }
    },
    "teststeps": [
        {
            "name": "transaction 1 start",
            "transaction": {
                "name": "tran1",
                "type": "start"
            }
        },
        {
            "name": "get with params",
            "request": {
                "method": "GET",
                "url": "/get",
                "params": {
                    "foo1": "$varFoo1",
                    "foo2": "$varFoo2"
                },
                "headers": {
                    "User-Agent": "HttpRunnerPlus"
                }
            },
            "variables": {
                "b": 34.5,
                "n": 3,
                "name": "get with params",
                "varFoo2": "${max($a, $b)}"
            },
            "setup_hooks": [
                "${setup_hook_example($name)}"
            ],
            "teardown_hooks": [
                "${teardown_hook_example($name)}"
            ],
            "extract": {
                "varFoo1": "body.args.foo1"
            },
            "validate": [
                {
                    "check": "status_code",
                    "assert": "equals",
                    "expect": 200,
                    "msg": "check response status code"
                },
                {
                    "check": "headers.\"Content-Type\"",
                    "assert": "startswith",
                    "expect": "application/json"
                },
                {
                    "check": "body.args.foo1",
                    "assert": "length_equals",
                    "expect": 5,
                    "msg": "check args foo1"
                },
                {
                    "check": "$varFoo1",
                    "assert": "length_equals",
                    "expect": 5,
                    "msg": "check args foo1"
                },
                {
                    "check": "body.args.foo2",
                    "assert": "equals",
                    "expect": "34.5",
                    "msg": "check args foo2"
                }
            ]
        },
        {
            "name": "transaction 1 end",
            "transaction": {
                "name": "tran1",
                "type": "end"
            }
        },
        {
            "name": "post json data",
            "request": {
                "method": "POST",
                "url": "/post",
                "body": {
                    "foo1": "$varFoo1",
                    "foo2": "${max($a, $b)}"
                }
            },
            "validate": [
                {
                    "check": "status_code",
                    "assert": "equals",
                    "expect": 200,
                    "msg": "check status code"
                },
                {
                    "check": "body.json.foo1",
                    "assert": "length_equals",
                    "expect": 5,
                    "msg": "check args foo1"
                },
                {
                    "check": "body.json.foo2",
                    "assert": "equals",
                    "expect": 12.3,
                    "msg": "check args foo2"
                }
            ]
        },
        {
            "name": "post form data",
            "request": {
                "method": "POST",
                "url": "/post",
                "headers": {
                    "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8"
                },
                "body": {
                    "foo1": "$varFoo1",
                    "foo2": "${max($a, $b)}",
                    "time": "${get_timestamp()}"
                }
            },
            "extract": {
                "varTime": "body.form.time"
            },
            "validate": [
                {
                    "check": "status_code",
                    "assert": "equals",
                    "expect": 200,
                    "msg": "check status code"
                },
                {
                    "check": "body.form.foo1",
                    "assert": "length_equals",
                    "expect": 5,
                    "msg": "check args foo1"
                },
                {
                    "check": "body.form.foo2",
                    "assert": "equals",
                    "expect": "12.3",
                    "msg": "check args foo2"
                }
            ]
        },
        {
            "name": "get with timestamp",
            "request": {
                "method": "GET",
                "url": "/get",
                "params": {
                    "time": "$varTime"
                }
            },
            "validate": [
                {
                    "check": "body.args.time",
                    "assert": "length_equals",
                    "expect": 13,
                    "msg": "check extracted var timestamp"
                }
            ]
        }
    ]
}`
	BuildHashicorpPyPlugin([]byte{}, "")
	defer RemoveHashicorpPyPlugin("")
	testcase3 := &TestCaseJson{
		JsonString: jsonString,
		ID:         1,
	}
	testCase, _ := testcase3.ToTestCase()
	err := NewRunner(t).Run(testCase)
	if err != nil {
		t.Fatalf("run testcase error: %v", err)
	}
}*/

func TestJsonRunnerCase(t *testing.T) {
	jsonString := `{
  "ID": 2,
  "Name": "httpbin 步骤1",
  "config": {
    "ID": 1,
    "CreatedAt": "2022-08-22T19:40:06.294+08:00",
    "UpdatedAt": "2022-09-21T17:17:41.7+08:00",
    "name": "httpbin 配置1",
    "base_url": "http://httpbin.org",
    "variables": {},
    "headers": {},
    "parameters": {},
    "variables_json": [],
    "headers_json": [],
    "weight": 0,
    "default": true,
    "verify": false,
    "setup_case": null,
    "setup_case_id": 1
  },
  "teststeps": [
    {
      "ID": 1,
      "Name": "httpbin 步骤2",
      "testcase": {
        "ID": 1,
        "Name": "httpbin 步骤2",
        "config": {
          "ID": 1,
          "CreatedAt": "2022-08-22T19:40:06.294+08:00",
          "UpdatedAt": "2022-09-21T17:17:41.7+08:00",
          "name": "httpbin 配置1",
          "base_url": "http://httpbin.org",
          "variables": {},
          "headers": {},
          "parameters": {},
          "variables_json": [],
          "headers_json": [],
          "weight": 0,
          "default": true,
          "verify": false,
          "setup_case": null,
          "setup_case_id": 1
        },
        "teststeps": [
          {
            "ID": 26,
            "CreatedAt": "2022-08-22T19:44:50.753+08:00",
            "UpdatedAt": "2022-09-15T18:18:16.157+08:00",
            "name": "headers",
            "type": 2,
            "request": {
              "ID": 26,
              "CreatedAt": "2022-08-22T19:44:50.71+08:00",
              "UpdatedAt": "2022-09-14T18:38:21.756+08:00",
              "agreement": "",
              "method": "GET",
              "url": "/headers",
              "params": {},
              "headers": {},
              "data": {},
              "params_json": [],
              "headers_json": [],
              "data_json": [],
              "json": {},
              "verify": false
            },
            "ThinkTimeID": 0,
            "TransactionID": 0,
            "RendezvousID": 0,
            "variables": {},
            "extract": {},
            "validate": [
              {
                "assert": "equals",
                "check": "status_code",
                "expect": 2010,
                "expectTemp": "200",
                "msg": "check status code",
                "type": 2
              }
            ],
            "validate_number": 1,
            "validate_json": null,
            "extract_json": [],
            "variables_json": [],
            "hooks": "",
            "testCase": null,
            "sort": 1,
            "export_header": [],
            "export_parameter": []
          },
          {
            "ID": 28,
            "CreatedAt": "2022-09-03T00:59:15.708+08:00",
            "UpdatedAt": "2022-09-15T18:18:16.185+08:00",
            "name": "delete 请求",
            "type": 2,
            "request": {
              "ID": 28,
              "CreatedAt": "2022-09-03T00:59:15.675+08:00",
              "UpdatedAt": "2022-09-03T00:59:15.675+08:00",
              "agreement": "",
              "method": "DELETE",
              "url": "/delete",
              "params": {},
              "headers": {},
              "data": {},
              "params_json": [],
              "headers_json": [],
              "data_json": [],
              "json": {},
              "verify": false
            },
            "ThinkTimeID": 0,
            "TransactionID": 0,
            "RendezvousID": 0,
            "variables": {},
            "extract": {},
            "validate": [],
            "validate_number": 0,
            "validate_json": null,
            "extract_json": [],
            "variables_json": [],
            "hooks": "",
            "testCase": null,
            "sort": 2,
            "export_header": [],
            "export_parameter": []
          },
          {
            "ID": 22,
            "CreatedAt": "2022-09-03T00:58:06.501+08:00",
            "UpdatedAt": "2022-09-15T18:18:16.216+08:00",
            "name": "get 请求",
            "type": 2,
            "request": {
              "ID": 22,
              "CreatedAt": "2022-09-03T00:58:06.468+08:00",
              "UpdatedAt": "2022-09-03T01:04:13.497+08:00",
              "agreement": "",
              "method": "GET",
              "url": "/get",
              "params": {},
              "headers": {},
              "data": {},
              "params_json": [],
              "headers_json": [],
              "data_json": [],
              "json": {},
              "verify": false
            },
            "ThinkTimeID": 0,
            "TransactionID": 0,
            "RendezvousID": 0,
            "variables": {},
            "extract": {
              "url": "body.url"
            },
            "validate": [
              {
                "assert": "equals",
                "check": "body.url",
                "expect": "http://httpbin.org/get",
                "expectTemp": "http://httpbin.org/get",
                "msg": "断言url",
                "type": 1
              }
            ],
            "validate_number": 1,
            "validate_json": null,
            "extract_json": [
              {
                "key": "url",
                "desc": "提取url变量",
                "value": "body.url",
                "rowKeyID": 0
              }
            ],
            "variables_json": [],
            "hooks": "",
            "testCase": null,
            "sort": 3,
            "export_header": [],
            "export_parameter": []
          },
          {
            "ID": 23,
            "CreatedAt": "2022-09-03T00:58:23.974+08:00",
            "UpdatedAt": "2022-09-15T18:18:16.246+08:00",
            "name": "post 请求",
            "type": 2,
            "request": {
              "ID": 23,
              "CreatedAt": "2022-09-03T00:58:23.94+08:00",
              "UpdatedAt": "2022-09-03T00:58:31.442+08:00",
              "agreement": "",
              "method": "POST",
              "url": "/post",
              "params": {},
              "headers": {},
              "data": {},
              "params_json": [],
              "headers_json": [],
              "data_json": [],
              "json": {},
              "verify": false
            },
            "ThinkTimeID": 0,
            "TransactionID": 0,
            "RendezvousID": 0,
            "variables": {},
            "extract": {},
            "validate": [],
            "validate_number": 0,
            "validate_json": null,
            "extract_json": [],
            "variables_json": [],
            "hooks": "",
            "testCase": null,
            "sort": 4,
            "export_header": [],
            "export_parameter": []
          },
          {
            "ID": 27,
            "CreatedAt": "2022-09-03T00:58:06.501+08:00",
            "UpdatedAt": "2022-09-15T18:18:16.274+08:00",
            "name": "get 请求",
            "type": 2,
            "request": {
              "ID": 27,
              "CreatedAt": "2022-09-03T00:58:06.468+08:00",
              "UpdatedAt": "2022-09-03T00:58:06.468+08:00",
              "agreement": "",
              "method": "GET",
              "url": "/get",
              "params": {},
              "headers": {},
              "data": {},
              "params_json": [],
              "headers_json": [],
              "data_json": [],
              "json": {},
              "verify": false
            },
            "ThinkTimeID": 0,
            "TransactionID": 0,
            "RendezvousID": 0,
            "variables": {},
            "extract": {},
            "validate": [],
            "validate_number": 0,
            "validate_json": null,
            "extract_json": [],
            "variables_json": [],
            "hooks": "",
            "testCase": null,
            "sort": 5,
            "export_header": [],
            "export_parameter": []
          },
          {
            "ID": 24,
            "CreatedAt": "2022-09-03T00:58:53.068+08:00",
            "UpdatedAt": "2022-09-15T18:18:16.304+08:00",
            "name": "put 请求",
            "type": 2,
            "request": {
              "ID": 24,
              "CreatedAt": "2022-09-03T00:58:53.034+08:00",
              "UpdatedAt": "2022-09-03T00:58:53.034+08:00",
              "agreement": "",
              "method": "PUT",
              "url": "/put",
              "params": {},
              "headers": {},
              "data": {},
              "params_json": [],
              "headers_json": [],
              "data_json": [],
              "json": {},
              "verify": false
            },
            "ThinkTimeID": 0,
            "TransactionID": 0,
            "RendezvousID": 0,
            "variables": {},
            "extract": {},
            "validate": [],
            "validate_number": 0,
            "validate_json": null,
            "extract_json": [],
            "variables_json": [],
            "hooks": "",
            "testCase": null,
            "sort": 6,
            "export_header": [],
            "export_parameter": []
          },
          {
            "ID": 25,
            "CreatedAt": "2022-09-03T00:59:15.708+08:00",
            "UpdatedAt": "2022-09-15T18:18:16.331+08:00",
            "name": "delete 请求",
            "type": 2,
            "request": {
              "ID": 25,
              "CreatedAt": "2022-09-03T00:59:15.675+08:00",
              "UpdatedAt": "2022-09-03T00:59:15.675+08:00",
              "agreement": "",
              "method": "DELETE",
              "url": "/delete",
              "params": {},
              "headers": {},
              "data": {},
              "params_json": [],
              "headers_json": [],
              "data_json": [],
              "json": {},
              "verify": false
            },
            "ThinkTimeID": 0,
            "TransactionID": 0,
            "RendezvousID": 0,
            "variables": {},
            "extract": {},
            "validate": [],
            "validate_number": 0,
            "validate_json": null,
            "extract_json": [],
            "variables_json": [],
            "hooks": "",
            "testCase": null,
            "sort": 7,
            "export_header": [],
            "export_parameter": []
          }
        ]
      }
    },
    {
      "ID": 2,
      "Name": "httpbin 步骤1",
      "testcase": {
        "ID": 2,
        "Name": "httpbin 步骤1",
        "config": {
          "ID": 1,
          "CreatedAt": "2022-08-22T19:40:06.294+08:00",
          "UpdatedAt": "2022-09-21T17:17:41.7+08:00",
          "name": "httpbin 配置1",
          "base_url": "http://httpbin.org",
          "variables": {},
          "headers": {},
          "parameters": {},
          "variables_json": [],
          "headers_json": [],
          "weight": 0,
          "default": true,
          "verify": false,
          "setup_case": null,
          "setup_case_id": 1
        },
        "teststeps": [
          {
            "ID": 15,
            "CreatedAt": "2022-08-22T19:44:50.753+08:00",
            "UpdatedAt": "2022-09-01T10:48:30.665+08:00",
            "name": "headers",
            "type": 2,
            "request": {
              "ID": 15,
              "CreatedAt": "2022-08-22T19:44:50.71+08:00",
              "UpdatedAt": "2022-08-22T19:44:50.71+08:00",
              "agreement": "",
              "method": "GET",
              "url": "/headers",
              "params": {},
              "headers": {},
              "data": {},
              "params_json": [],
              "headers_json": [],
              "data_json": [],
              "json": {},
              "verify": false
            },
            "ThinkTimeID": 0,
            "TransactionID": 0,
            "RendezvousID": 0,
            "variables": {},
            "extract": {},
            "validate": [
              {
                "assert": "equals",
                "check": "status_code",
                "expect": 200,
                "expectTemp": "200",
                "msg": "check status code",
                "type": 2
              }
            ],
            "validate_number": 1,
            "validate_json": null,
            "extract_json": [],
            "variables_json": [],
            "hooks": "",
            "testCase": null,
            "sort": 1,
            "export_header": [],
            "export_parameter": []
          }
        ]
      }
    }
  ]
}`
	//BuildHashicorpPyPlugin([]byte{}, "")
	//defer RemoveHashicorpPyPlugin("")
	testcase3 := &JsonToCase{
		JsonString:        jsonString,
		ID:                1,
		DebugTalkFilePath: "",
	}
	//NewRunner(t).SetHTTPStatOn().GenHTMLReport().SetFailfast(false).RunJsons(r.tcm.Case...)
	testCase, _ := testcase3.ToTestCase()
	r := NewRunner(t)
	r.GenHTMLReport()
	r.SetFailfast(false)
	r.RunJsons(testCase)
	//	if err != nil {
	//		t.Fatalf("run testcase error: %v", err)
	//	}
	//	reportStr, _ := json.Marshal(report)
	//	fmt.Println(string(reportStr))
}
