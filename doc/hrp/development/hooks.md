
在v4 go版本之前`setup_hooks`和`teardown_hooks`用于在请求前和返回后的一些操作操作，在`v4`中只能作为一个函数，无法修改任何的内容

## `V4` go版本与python版本的差异
### python
1. 调用函数进行数据处理
2. 由于函数驱动用的也是python，可以直接修改函对象的内容，所以不用返回数据也可以实现修改请求和返回的数据
3. 使用`request`作为请求内容的参数，使用`response`作为返回内容的参数

### go
1. 调用函数只能处理简单的逻辑，如：sleep、查库等
2. 驱动函数使用的是grpc，可以用python、Java、go等语言实现，结果依赖于grpc的返回，目前代码中没有实现读取返回内容
3. 使用`hrp_step_request`作为请求内容的参数，使用`hrp_step_response`作为返回内容的参数

在没看源码时，一直觉得应该是使用相同的参数，但是一直提示`request`、`response`变量不存在，

## 原有功能

1. `SetupHooks` 和 `TeardownHooks` 支持多个函数操作
2. 每个函数没有对返回值进行处理

## 实现功能
> 前提条件：尽量保存原有的逻辑

1. 由于接口需要加解密的操作，需要可以对请求内容进行加密、对返回内容进行解密
2. 使用`request`作为请求参数、`response`返回参数，如果用上前缀比较繁琐
3. 由于hooks是一个列表，如果有多个修改将以最后的修改为准

## 修改源码

> 如果你对源码还不够熟悉，可以先从[源码解析](https://yangfan.gd.cn/hrp/)开始
> 本次代码修改已在4.3.3中被采纳，如果低于当前版本可以自己修改
> 主要修改文件***hrp/step_request.go***

以下代码中，由于`SetupHooks` 和 `TeardownHooks`的实现方式不同，呈现出来修改的也不同，在尽量保持和源码写法一致时作出的修改，每次调用时都会更新最新内容

```golang
// runStepRequest 执行请求
func runStepRequest(r *SessionRunner, step *TStep) (stepResult *StepResult, err error) {
	// 省略非关键代码
	// add request object to step variables, could be used in setup hooks
	// 将请求对象添加到step变量中，hook中可以直接使用hrp_step_name、hrp_step_request、request 变量
	stepVariables["hrp_step_name"] = step.Name
	stepVariables["hrp_step_request"] = rb.requestMap
	// 增加request变量
	stepVariables["request"] = rb.requestMap  

	// deal with setup hooks
	// 遍历setp hook并执行
	for _, setupHook := range step.SetupHooks {
		// 调用hook
		req, err := parser.Parse(setupHook, stepVariables)
		if err != nil {
			return stepResult, errors.Wrap(err, "run setup hooks failed")
		}
		// 处理hook返回的结果
		reqMap, ok := req.(map[string]interface{})
		// 每次调用回写request信息
		if ok && reqMap != nil {
			rb.requestMap = reqMap
			stepVariables["request"] = reqMap
		}
	}
	// 将hook处理后的结果回写到请求对象中，主要更新body和header的内容
	// 由于每次调用信息已经回写，这里只需要判断是否有setup hooks
	if len(step.SetupHooks) > 0 {
		// 更新body的内容
		requestBody, ok := rb.requestMap["body"].(map[string]interface{})
		if ok {
			body, err := json.Marshal(requestBody)
			if err == nil {
				rb.req.Body = io.NopCloser(bytes.NewReader(body))
				rb.req.ContentLength = int64(len(body))
			}
		}
		// 更新header的内容
		headers, ok := rb.requestMap["headers"].(map[string]string)
		rb.req.Header = map[string][]string{}
		for key, value := range headers {
			rb.req.Header.Set(key, value)
		}
	}
    
    // 省略非关键代码
    
    // add response object to step variables, could be used in teardown hooks
	// 将响应对象添加到step变量中，hrp_step_response、response 变量
	stepVariables["hrp_step_response"] = respObj.respObjMeta
	// 添加response变量
	stepVariables["response"] = respObj.respObjMeta

	// add response object to step variables, could be used in teardown hooks
	// 将响应对象添加到step变量中，hrp_step_response、response 变量
	stepVariables["hrp_step_response"] = respObj.respObjMeta
	stepVariables["response"] = respObj.respObjMeta

	// deal with teardown hooks
	// 遍历teardown hook并执行
	// 于setup hooks不同的实现方式，直接更新respObj.respObjMeta
	for _, teardownHook := range step.TeardownHooks {
		// 调用hook
		res, err := parser.Parse(teardownHook, stepVariables)
		if err != nil {
			return stepResult, errors.Wrap(err, "run teardown hooks failed")
		}
		// 处理hook返回的结果
		resMpa, ok := res.(map[string]interface{})
		if ok {
			stepVariables["response"] = resMpa
			respObj.respObjMeta = resMpa
		}
	}
	// 将hook处理后的结果回写到响应对象中
	sessionData.ReqResps.Request = rb.requestMap
	// 使用更新后的respObj.respObjMeta
	sessionData.ReqResps.Response = builtin.FormatResponse(respObj.respObjMeta)

	// 省略非关键代码
}

```




