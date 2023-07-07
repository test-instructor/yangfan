yangfan 自动化测试平台[开源项目](https://testerhome.com/opensource_projects/yangfan)
了解如何接入httprunner
[测试平台接入 HttpRunner V4（一）基本功能接入](https://testerhome.com/topics/35126)
[测试平台接入 HttpRunner V4（二）使用 config 实现用例之间的参数传递](https://testerhome.com/topics/35125)
# 大致功能

> 1. 项目管理：项目创建后会初始化函数驱动，可根据实际需要对项目进行划分，各项目数据相互独立，无法查看、引用其他项目的数据
> 2. 配置管理：公共数据配置，可以配置域名、请求头、变量和前置步骤等
> 3. 树形菜单：接口管理、测试步骤、测试用例都包含了树形菜单，可以根据树形菜单对接口按功能模块、服务等进行划分，方便用例管理
> 4. 接口管理：接口测试最基础模块，测试用例、测试步骤、定时任务等都依赖与接口管理
> 5. 测试步骤：数据从接口管理的数据复制过来，数据相互独立，互不影响；运行配置只在调试时生效，测试用例、定时任务执行时无效
> 6. 测试用例：引用测试步骤，执行时以测试用例的配置为主；测试步骤的修改，会导致测试用例运行报错、无法运行等
> 7. 定时任务：引用多个定时任务，执行时各用例项目独立，没有依赖
> 8. 性能任务：引用测试步骤，增加性能测试相关特性（如：事务、集合点等）
> 9. 测试报告：展示除压测任务的报告外的所有接口调试、运行报告
> 10. 性能测试报告：展示性能测试报告
> 11. 环境变量(`开发中`)：自行设置`开发环境`、`测试环境`、`预发布环境`等多个环境，相对固定的变量进行设置，如：域名、账号等


# 功能使用介绍


## 配置管理

### 主要功能
> 1. 前置步骤：引用测试步骤，在接口、步骤、用例类型进行调试和运行时，运行前置步骤，一般用于初始化数据，如登录、创建用户等操作
> 2. verify：是否开启https安全验证
> 3. 默认配置：接口调试时会自动选择默认配置
> 4. 域名格式：http://httpbin.org
> 5. Header：默认请求头，在用例下的所有数据的默认请求头
> 6. Variables：可以引用的变量，在当前配置下的所有接口都可以引用
> 7. Parameters：参数化列表，用于对某写遍历场景

### Header(默认请求头)
#### 主要功能
> 1. 标签： 内置部分常用标签（User-Agent、Host等），可进行搜索或者增加自定义标签
> 2. 内容： 标签对应的值
> 3. 默认请求头，所有使用该配置的接口，请求头默认会带上

#### 使用配置

1. 使用变量：$version
2. 调用函数无传参：${get_user_agent()}
3. 调用函数：${sum_ints(1,$number)}

> 1. 运行前
> ![config_header](https://testerhome.com/uploads/photo/2022/408fdc81-9540-4879-8524-1d08e4572939.png)
> 2. 运行后
> ![config_header_run](https://testerhome.com/uploads/photo/2022/c4321dc6-3eb1-42a8-97d4-175789575c26.png)

### Variables(变量)
#### 主要功能
> 1. 变量名：调用时需要"$"符号，如设置的变量名为version，调用时为：$version
> 2. 类型：目前为String、Integer、Float、Boolean、List、Dict
> 3. 变量值：根据类型设置对应的值，如设置错误则无法使用对应的变量


> 1. 运行前
> ![config_variables](https://testerhome.com/uploads/photo/2022/4699151f-a94e-4bfa-8d8c-3deba3a20bb6.png)
> 2. 运行后
> ![config_header_run](https://testerhome.com/uploads/photo/2022/2b2fc1d6-293e-423d-8636-46c5000716ba.png)


## 接口管理 

### 主要功能
> 1. 请求方法：GET、POST、PUT、DELETE、HEAD、OPTIONS、PATCH
> 2. Header：请求头，会覆盖配置中的默认请求头
> 3. Params：url中携带的参数
> 4. Form：表单中的参数
> 5. Jsons：json格式参数
> 6. Extract：参数提取，提取后的变量在当前步骤中有效
> 7. Validate：断言
> 8. Variables：局部变量，设置的局部变量只能在当前接口中有效
> 9. Hooks：

### Extract（参数提取），提取后的变量在当前步骤中有效
> 1. 变量名：调用时需要"$"符号，如设置的变量名为version，调用时为：$version
> 2. 抽取表达式
>    * status_code：响应状态码
>    * body：response body，通过body.data获取到body下的data字段

> 1. 运行前
> ![api_extract](https://testerhome.com/uploads/photo/2022/806ca4f7-e29e-4e20-8e37-f41e81513f4f.png)
> 2. 运行后
> ![api_resp_extract](https://testerhome.com/uploads/photo/2022/78fc5b92-db0e-45a0-b551-e2a48c060665.png)

### Validate（断言）
> 1. 断言字段：同`Extract`中的`抽取表达式`
> 2. 断言类型:
>    * equals： 是否相等
>    * less_than： 小于
>    * less_than_or_equals： 小于等于
>    * greater_than： 大于
>    * greater_than_or_equals： 大于等于
>    * not_equals： 不等于
>    * string_equals： 字符串相等
>    * length_equals： 长度相等
>    * length_greater_than： 长度大于
>    * length_greater_than_or_equals： 长度大于等于
>    * length_less_than： 长度小于
>    * length_less_than_or_equals： 长度小于等于
>    * contains： 预期结果是否被包含在实际结果中
>    * contained_by： 实际结果是否被包含在预期结果中
>    * type_match： 类型是否匹配
>    * regex_match： 正则表达式是否匹配
>    * startswith： 字符串是否以什么开头
>    * endswith： 字符串是否以什么结尾
> 3. 期望类型、期望返回值：根据断言类型设置相对应的值

> 1. 运行前
> ![api_validate](https://testerhome.com/uploads/photo/2022/d4c2b4ed-137a-4f21-a927-12a5cef09fa3.png
)
> 2. 运行后
> ![api_resp_validate](https://testerhome.com/uploads/photo/2022/6eb1d2da-3972-4b5c-a2d1-2921e5450e2d.png)

### hooks

> hooks 分为 setup hooks 和 teardown hooks 
> 可以用来对数据加解密或者初始化和清理数据的操作

### 如何使用

> ${setup_hook_encryption($request)}
> ${setup_hook_decrypt($response)}

1. 传参：setup hooks 使用request，teardown hooks 使用response
2. setup hooks 函数返回为request对象，teardown hooks 函数返回为response对象

```python
def setup_hook_encryption(request):
    request["body"]["setup_hook_encryption_request"] = "setup_hook_encryption_request"
    return request


def setup_hook_decrypt(response):
    response["body"]["setup_hook_decrypt"] = "setup_hook_encryption_response"
    return response
```


## 测试步骤（测试步骤）

> * 测试步骤由`API`组成，复制接口管理中的`API`并做外键关联，方便后续的数据统计。
> * 测试步骤可以复制多个相同或者不同的`API`，与`API`为多对多关系
> * 测试步骤一般为一个操作，如：登录、查询订单等
> * 设置为前置步骤时，一般不能依赖于其他步骤的返回值

### 主要功能
> 1. 调试运行配置：调试时需要使用的配置，临时变量，`测试用例`、`定时任务`、`性能测试中`执行时不使用该临时变量
> 2. 前置步骤：是否可以被设置为前置步骤，默认为否
> 3. 测试步骤：进入步骤详情可以设置测试步骤，通过拖动的方式进行增加和调整执行顺序

> 1. 步骤详情
> ![step_detail](https://testerhome.com/uploads/photo/2022/2e0cb792-3864-4a26-8e32-1968423a2d44.png)
> 2. 步骤添加测试步骤
> ![step_add_case](https://testerhome.com/uploads/photo/2022/5c04af03-be89-4a1a-ade7-f9737a275691.png)


## 测试用例

> * 测试用例由`步骤`组成，引用`步骤`,`步骤`修改会导致测试用例的执行结果
> * 运行配置：在`测试用例`和`定时任务`中执行使用该配置
> * 测试用例一般为一个完整的操作，如：注册流程、创建订单并完成支付等

### 主要功能
> 1. 用例详情：对用例引用的`步骤`进行维护，通过拖动进行排序
> 2. 排序：添加后的用例，默认排序为`999`，需要进行拖动排序(以免执行时顺序错误)，拖动排序后会以最后的顺序执行
> 3. 添加步骤：添加时可以同时添加多个，一个`步骤`可以同时添加多次

> 1. 测试用例详情
> ![case_detail](https://testerhome.com/uploads/photo/2022/35d0c43c-e5f8-446b-9026-b971295dd2c5.png)
> 2. 测试用例添加步骤
> ![case_add_step](https://testerhome.com/uploads/photo/2022/b9b8c562-ffc1-4b21-8620-c2eef1614df3.png)


## 定时任务

> * 接口测试中最核心的执行部分
> * 运行配置：定时任务中没有运行配置，根据测试用例中的运行配置去执行，所以用例的数据互不影响
> * 定时任务引用测试用例
> * 并发执行(`开发中`)：考虑到执行时间问题，多个用例可以同时执行以节省等待时间

### 主要功能
> 1. 任务详情：对用例引用的`测试用例`进行维护，通过拖动进行排序，因为每个测试用例相对独立，所以执行顺序不会影响接测试结果
> 2. 添加用例：添加时可以同时添加多个，一个`用例`可以同时添加多次
> 3. 定时执行：定时执行为启用状态时，需要填写时间配置

> 1. 编辑任务
> ![case_detail](https://testerhome.com/uploads/photo/2022/b4af1543-00ca-428b-98fe-4b663ed98864.png)
> 2. 任务详情
> ![case_add_step](https://testerhome.com/uploads/photo/2022/54d2f8d8-2a06-4167-82f1-0f0b7b66cf58.png)
> 3. 任务添加测试用例
> ![case_add_step](https://testerhome.com/uploads/photo/2022/6ebbf708-eff7-41c2-9746-6b2c60681660.png)

## 测试报告

> * 用例类型：`api`、`步骤`、`用例`、`定时任务`、`性能测试`
> * 执行类型：`保存调试`、`调试运行`、`后台运行`、`定时运行`

### 主要功能
> 1. 除性能数据外的其他运行测试报告都在测试报告内
> 2. 测试报告详情只统计`用例状态`和`接口状态`，暂未统计`步骤状态`

> 1. 测试报告列表
> ![report_list](https://testerhome.com/uploads/photo/2022/2a689dbe-4140-4a70-b972-a8b15bac8b4e.png)
> 2. 测试报告详情
> ![report_detail](https://testerhome.com/uploads/photo/2022/5b9ccf62-5d39-4ee2-a3a2-a85a398c6702.png)
> 3. 测试报告接口详情
> ![report_api_detail](https://testerhome.com/uploads/photo/2022/74d2d544-c050-49b9-97b9-1d3e2039cc0a.png)

     