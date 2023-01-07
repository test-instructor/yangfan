# cheetah 自动化测试平台性能测试

得益于[httprunner v4](https://github.com/httprunner/httprunner)优秀开源项目，基于go语言的v4，集成了性能测试。目前测试平台已经完成`接口自动化`、`性能测试`两部分，后续将增加`UI自动化`、`gRPC`等功能

## 实现功能
> * 在接口测试的基础上，使用请参考[测试平台接入 HttpRunner V4（三）详细功能介绍](https://testerhome.com/topics/35161)，引用接口测试的`测试套件`，并增加`事务`和`集合。
> * 目前只完成单机版的压测部分，k8s 部署分布式压测正在开发中
> * 压测报告：v4 压测需要`Prometheus` + `Grafana`，在部署过程会出现各种各样的问题，cheetah 将性能报告存储到数据库，并提供页面供用户查询，节省用户在环境上的使用问题

## 功能介绍

> * 性能任务：主要为`调试运行`和`启动压测`两个功能
> * 调试运行：运行后默认会打开测试报告，检验所有接口是否符合预期
> * 启动压测：启动时需要设置参数`并发用户数`和`初始每秒增加用户数`，压测过程中无法修改，动态修改参数功能正在开发中
> * 压测报告：所有压测报告都在此显示，`准备中`的压测报告无法进入详情

![performance_detail.png](https://testerhome.com/uploads/photo/2022/922f44d7-1591-46a0-8d50-1be1088484c0.png)
![performance_detail_all.png](https://testerhome.com/uploads/photo/2022/9b4a796b-7045-42a4-a33c-57c0a1f5696c.png)


# 往期文档
[cheetah 自动化测试平台](https://testerhome.com/opensource_projects/cheetah)

了解如何接入httprunner
* [测试平台接入 HttpRunner V4（一）基本功能接入](https://testerhome.com/topics/35126)
* [测试平台接入 HttpRunner V4（二）使用 config 实现用例之间的参数传递](https://testerhome.com/topics/35125)
* [测试平台接入 HttpRunner V4（三）详细功能介绍](https://testerhome.com/topics/35161)
