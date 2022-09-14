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
	"Config": {
		"ID": 1,
		"CreatedAt": "2022-08-22T19:40:06.294+08:00",
		"UpdatedAt": "2022-08-22T19:49:54.502+08:00",
		"name": "httpbin 配置",
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
		"setup_case_id": null
	},
	"TestSteps": [
		{
			"ID": 0,
			"CreatedAt": "0001-01-01T00:00:00Z",
			"UpdatedAt": "0001-01-01T00:00:00Z",
			"name": "headers",
			"type": 0,
			"request": {
				"ID": 0,
				"CreatedAt": "0001-01-01T00:00:00Z",
				"UpdatedAt": "0001-01-01T00:00:00Z",
				"agreement": "",
				"method": "",
				"url": "",
				"params": null,
				"headers": null,
				"data": null,
				"params_json": null,
				"headers_json": null,
				"data_json": null,
				"json": null,
				"verify": false
			},
			"variables": null,
			"extract": null,
			"validate": null,
			"validate_number": 0,
			"validate_json": null,
			"extract_json": null,
			"variables_json": null,
			"hooks": "",
			"testCase": null,
			"sort": 0,
			"export_header": null,
			"export_parameter": null,
			"TestCase": {
				"Config": {
					"ID": 1,
					"CreatedAt": "2022-08-22T19:40:06.294+08:00",
					"UpdatedAt": "2022-08-22T19:49:54.502+08:00",
					"name": "httpbin 配置",
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
					"setup_case_id": null
				},
				"TestSteps": [
					{
						"ID": 4,
						"CreatedAt": "2022-08-22T19:44:50.753+08:00",
						"UpdatedAt": "2022-08-22T19:44:50.753+08:00",
						"name": "headers",
						"type": 1,
						"request": {
							"ID": 4,
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
						"sort": 0,
						"export_header": [],
						"export_parameter": [],
						"TestCase": null
					}
				]
			}
		},{
			"ID": 0,
			"CreatedAt": "0001-01-01T00:00:00Z",
			"UpdatedAt": "0001-01-01T00:00:00Z",
			"name": "headers11111111111",
			"type": 0,
			"request": {
				"ID": 0,
				"CreatedAt": "0001-01-01T00:00:00Z",
				"UpdatedAt": "0001-01-01T00:00:00Z",
				"agreement": "",
				"method": "",
				"url": "",
				"params": null,
				"headers": null,
				"data": null,
				"params_json": null,
				"headers_json": null,
				"data_json": null,
				"json": null,
				"verify": false
			},
			"variables": null,
			"extract": null,
			"validate": null,
			"validate_number": 0,
			"validate_json": null,
			"extract_json": null,
			"variables_json": null,
			"hooks": "",
			"testCase": null,
			"sort": 0,
			"export_header": null,
			"export_parameter": null,
			"TestCase": {
				"Config": {
					"ID": 1,
					"CreatedAt": "2022-08-22T19:40:06.294+08:00",
					"UpdatedAt": "2022-08-22T19:49:54.502+08:00",
					"name": "httpbin 配置",
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
					"setup_case_id": null
				},
				"TestSteps": [
					{
						"ID": 4,
						"CreatedAt": "2022-08-22T19:44:50.753+08:00",
						"UpdatedAt": "2022-08-22T19:44:50.753+08:00",
						"name": "headers222222",
						"type": 1,
						"request": {
							"ID": 4,
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
						"sort": 0,
						"export_header": [],
						"export_parameter": [],
						"TestCase": null
					}
				]
			}
		}
	]
}`
	//BuildHashicorpPyPlugin([]byte{}, "")
	//defer RemoveHashicorpPyPlugin("")
	testcase3 := &TestCaseJson{
		JsonString:        jsonString,
		ID:                1,
		DebugTalkFilePath: "",
	}
	//NewRunner(t).SetHTTPStatOn().GenHTMLReport().SetFailfast(false).RunJsons(r.tcm.Case...)
	testCase, _ := testcase3.ToTestCase()
	r := NewRunner(t)
	r.GenHTMLReport()

	r.RunJsons(testCase)
	//	if err != nil {
	//		t.Fatalf("run testcase error: %v", err)
	//	}
	//	reportStr, _ := json.Marshal(report)
	//	fmt.Println(string(reportStr))
}
