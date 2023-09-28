

本阶段主要以接口自动化进行讲解并实现MQTT协议，所以对UI自动化这部分会将会一带过，后续阶段会讲解UI自动化相关的内容，并实现web ui自动化。对于命令行相关部分，这部分一般来说不会对其有太大的改动，所以这部分暂时不讲解。

## 

## 主要目录

* hrp：`httprunner` go版本代码目录，所有代码都在该目录下去实现，所以后面的内容都为hrp下的内容
* cmd：命令行工具，如果你用`hrp`运行，所有的命令都在`cmd`目录下
* internal：内部库，只允许hrp下的文件调用，在hrp外的其他目录层级无法调用，可以看成`hrp`下的私有方法
* tests：部分单元测试用例，除了该目录下的单元测试，其他文件名称为`_test.go`结尾的也是单元测试内容。对于测试框架来说，单元测试可以更好确保修改后的准确性

## 文件

### hrp 下的文件

```shell
├──` boomer.go`
├── `build.go`
├── `config.go`：配置相关内容
├── `convert.go`：内容为空
├── `loader.go`：读取测试用例，输入的用例格式为`ITestCase`接口，返回`TestCase`类型的用例列表
├── `parameters.go`：配置中`parameters`字段
├── `parser.go`：解析`url`、`body`、`params`等内容中字符串包含函数和变量的方法
├── `plugin.go`：函数插件，如果有使用函数，则从这边进入
├── `response.go`：测试报告中的`response`
├── `runner.go`：运行测试用例
├── `server.go`：性能测试`http`服务，使用多机负载时会启动，通过`start`、`rebalance`、`stop`、`quit`、`workers`、`master`接口来操作和获取节点信息
├── `step.go`：测试步骤、测试报告的结构体和测试步骤类型定义，`StepType`定义了不同步骤的类型
├── `step_api.go`：`api`测试步骤
├── `step_mobile_ui.go`：`ios`和`android`测试步骤中
├── `step_rendezvous.go`：性能测试中的集合
├── `step_request.go`：`http`中的`request`对象
├── `step_testcase.go`：测试用例（测试步骤中的一种）
├── `step_thinktime.go`：性能测试中的思考时间
├── `step_transaction.go`：性能测试中的事务
├── `step_websocket.go`：websocket测试步骤
├── `summary.go`：测试报告
└── `testcase.go`：测试用例定义，通过`json/yaml`由这里进行定义并解析
```



> 所有`xxx_test.go`一般为对应的`xxx.go`文件中各方法的单元测试用例，下面就不单独介绍单元测试的内容

1. `boomer.go`：性能测试入口，`cmd` 中运行也是从这个文件中开始调用的
    * `NewMasterBoomer`：创建多机负载测试的控制节点
    * `NewWorkerBoomer`：创建多机负载测试的工作节点
    * `NewStandaloneBoome`r：创建本地运行的boomer实例

### `hrp/cmd`

主要是命令行工具，这里暂不讲解

### `hrp/internal`

```shell
├── `builtin`：
│   ├── `assertion.go`：断言
│   ├── `function.go`：内置函数
│   └── `utils.go`：工具文件，有获取`json/yaml`、创建目录、创建文件、判断路径等方法
├── `code`：
│   └── `code.go`：自定义错误类型代码
├── `dial`：
│   ├── `dns.go`：DNS解析功能
│   └── `ping.go`：Ping功能
├── `env`：
│   └── `env.go`：获取环境变量
├── `json`：
│   └── `json.go`：定义json相关操作
├── `myexec`：python 的相关命令行操作，其中`_uixt` 和`_windows`在构建时会根据不同的系统选择到不同的文件
│   ├── `cmd.go`：
│   ├── `cmd_uixt.go`：
│   └── `cmd_windows.go`：
├── `pytest`：
│   └── `main.go`：执行`hrp pytest`命令时由这里进行转发
├── `scaffold`：函数插件和测试用例
│   ├── `examples_test.go`：
│   ├── `main.go`：创建插件
│   └── `templates`：
│       ├── `api`：api 用例文件
│       ├── `env`：
│       ├── `gitignore`：
│       ├── `plugin`：插件demo
│       ├── `pytest.ini`：pytest 配置文件
│       ├── `report`：测试报告模板
│       └── `testcases`：测试用例
├── `sdk`：环境变量相关
├── `version`：版本号信息
└── `wiki`：打开`httprunner`官网
```



### `hrp/pkg`

```shell
├── `boomer`：性能测试
│   ├── `boomer.go`：
│   ├── `client_grpc.go`：grpc 客户端，工作节点使用
│   ├── `data`：grpc服务证书
│   ├── `grpc`：proto 协议
│   ├── `message.go`：grpc 消息定义
│   ├── `output.go`：性能测试报告
│   ├── `ratelimiter.go`：并发用户数
│   ├── `runner.go`：运行性能测试，包括修改并发用户数、停止等操作
│   ├── `server_grpc.go`：grpc 服务端，master 节点使用
│   ├── `stats.go`：请求统计功能，用于记录请求的响应时间、成功与失败次数、错误信息等统计数据
│   ├── `task.go`：
│   ├── `ulimit.go`：
│   ├── `ulimit_windows.go`：
│   └── `utils.go`：性能监控
├── `convert`：用例转换
├── `gadb`：用于与 Android 设备进行交互
├── `gidevice`：iOS 设备进行通信
├── `httpstat`：获取`http`请求相关时间
│   ├── `demo`：
│   │   └── `main_test.go`：
│   └── `main.go`：
└── `uixt`：Android 和 iOS 操作上封装的方法
```

