# httprunner V4 版本通过 hooks 实现加解密

# 为何进行二次开发

1. 有hooks机制，但是调用时发现和v3版本不一样，识别不到request和response两个变量，查看源码发现加了前缀`hrp_step_`
2. 查看例子，发现hook对应的函数并没有返回，v3 中自定义函数只支持 python，在自定义函数中可以直接对 request 进行修改，由于 v4 用的是plugin，所以需要有返回值才可以

# 分析源码

通过源码得知SetupHooks和TeardownHooks并没有返回值，只要加上返回值，然后给rb.requestMap重新赋值，同时也要给rb.req赋值，rb.req是实际请求数据，在之前就处理过了，所以这里需要再进行处理
```go 
package hrp

func runStepRequest(r *SessionRunner, step *TStep) (stepResult *StepResult, err error) {
	// ... 省略

	// add request object to step variables, could be used in setup hooks
	stepVariables["hrp_step_name"] = step.Name
	stepVariables["hrp_step_request"] = rb.requestMap
	
	// deal with setup hooks
	for _, setupHook := range step.SetupHooks {
		_, err = parser.Parse(setupHook, stepVariables)
		if err != nil {
			return stepResult, errors.Wrap(err, "run setup hooks failed")
		}
	}

	// ... 省略
}
```

# 修改源码

> 以SetupHooks为例
1. 函数内修改数据，但是不能改变request的结构
```python
def setup_hook_encryption(request):
    request["body"]["setup_hook_encryption_request"] = "setup_hook_encryption_request"
    return request
```
2. stepVariables增加request，兼容v3的写法
3. 数据写入rb.requestMap
4. 数据写入rb.req.Body，修改后需要修改ContentLength，否则会因为长度不一致导致报错

```go 
package hrp

func runStepRequest(r *SessionRunner, step *TStep) (stepResult *StepResult, err error) {
	// ... 省略

	// add request object to step variables, could be used in setup hooks
	stepVariables["hrp_step_name"] = step.Name
	stepVariables["hrp_step_request"] = rb.requestMap
	stepVariables["request"] = rb.requestMap

	// deal with setup hooks
	for _, setupHook := range step.SetupHooks {
		req, err := parser.Parse(setupHook, stepVariables)
		if err != nil {
			return stepResult, errors.Wrap(err, "run setup hooks failed")
		}
		reqMap, ok := req.(map[string]interface{})
		if ok {
			rb.requestMap = reqMap
			stepVariables["request"] = reqMap
		}
	}
	if len(step.SetupHooks) > 0 {
		requestBody, ok := rb.requestMap["body"].(map[string]interface{})
		if ok {
			body, err := json.Marshal(requestBody)
			if err == nil {
				rb.req.Body = io.NopCloser(bytes.NewReader(body))
				rb.req.ContentLength = int64(len(body))
			}
		}
	}

	// ... 省略

}
```









