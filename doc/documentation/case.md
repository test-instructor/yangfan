

测试步骤由一个或多个接口组成，可以测试步骤作为前置步骤使用
测试用例为独立可执行单元，由一个或多个测试步骤

## 设计理念

将用例管理分为分解为测试用例、测试步骤和API三个模块的设计，可以提高测试用例的可维护性、复用性、可读性和可扩展性。

1. 提高测试用例的可维护性：通过将测试用例分解为测试步骤和API两个模块，可以将测试用例的复杂度降低，并且可以通过修改测试步骤和API来更新多个测试用例，从而提高了测试用例的可维护性。
2. 提高测试用例的复用性：通过将测试步骤和API单独设计，可以将其复用在多个测试用例中，减少测试用例的重复编写，从而提高了测试用例的复用性。
3. 提高测试用例的可读性：通过将测试用例分解为测试步骤和API两个模块，可以将测试用例的结构变得更加清晰，易于阅读和理解，从而提高了测试用例的可读性。
4. 方便测试用例的扩展：通过将测试步骤和API单独设计，可以方便地扩展测试用例，例如增加新的测试步骤或API，从而满足不同的测试需求。

### 后续优化

1. 用例关系图：展示接口、步骤、用例、任务的关联情况，方便管理
2. 导入用户：主要为导入用例json、yaml、swagger格式，json、yaml为httprunner格式的用例，通过swagger导入后，可以关联对应的服务、接口信息，运行失败后更方便定位
3. 用例录制：录制web中访问的记录

## 接口管理

1. Header: 请求头，复选框为将当前请求头字段导出为全局header，场景：token保持在后续的接口中，无需每个接口再添加token字段
2. Extract：提取变量，复选框为将当前请求头字段导出为全局变量，场景：用于步骤、用例的解耦，假设步骤执行顺序为步骤A、步骤B，那么可以在步骤B中引用步骤A导出的变量，`引用其他步骤的变量时需要确保执行顺序`
3. `Params`、`Form`、`Jsons`、`Validate`、`Variables`参考`httprunner`中的应用
   ![img.png](https://qiniu.yangfan.gd.cn/image/documents/apiadd.png)

## 测试步骤

复用于接口管理中的api，无法直接在步骤中进行新建

### 步骤使用

1. 新建步骤
2. 进入步骤详情，在详情中添加api，对已添加的api进行拖动排序
3. 在步骤详情对测试接口进行定制化修改、删除、拖动排序
   ![img.png](https://qiniu.yangfan.gd.cn/image/documents/stepDetail.png)
   ![img.png](https://qiniu.yangfan.gd.cn/image/documents/stepDetailAdd.png)


## 测试用例

> 各用例的数据应该独立，不能依赖于其他用例，如数据有依赖，需要在步骤中增加对应的依赖

1. 新建测试用例
2. 进入用例详情，在详情中添加步骤，一个用例可以同时添加多个步骤（一个步骤可以被多次添加）
3. 在用例详情中调整步骤执行顺序

![img.png](https://qiniu.yangfan.gd.cn/image/documents/caseDetail.png)
![img.png](https://qiniu.yangfan.gd.cn/image/documents/caseDetailAdd.png)