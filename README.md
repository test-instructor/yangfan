# Cheetah
![logoLogin](docs/img/logoLogin.jpg)


- 前后端采用[gin-vue-admin v2.5.3](https://www.gin-vue-admin.com/), Gin + Vue全栈开发基础平台
- 测试引擎采用[HttpRunner V4](https://httprunner.com/),支持 HTTP(S)/HTTP2/WebSocket/RPC 等网络协议，涵盖接口测试、性能测试、数字体验监测等测试类型。简单易用，功能强大，具有丰富的插件化机制和高度的可扩展能力。

## 支持功能
- [x] 用户管理
- [x] 项目管理
- [x] 接口管理
- [x] 测试用例管理
- [x] 定时任务
- [x] 生成测试报告
- [ ] 并发执行多个定时任务/用例/接口
- [ ] 支持更多协议，`HTTP/2`、`WebSocket`、`TCP`、`RPC`等
- [ ] 数据驱动`parameterize`
- [ ] 支持用例导入，`json`、`postman`、`swagger`等
- [ ] 性能测试
- [ ] web UI 自动化测试

## 部署方式


> - [环境准备](https://www.gin-vue-admin.com/guide/start-quickly/env.html)
> - 1、新建数据库，并导入server/sql/cheetah.sql文件
> - 2、修改`server/config.yaml`文件中的数据库`mysql`、飞书登录`fs`相关配置

### [项目上线](https://www.gin-vue-admin.com/guide/deployment/)

## 项目概况

![login.png](docs/img/login.png)
![config](docs/img/config.png)
![api](docs/img/api.png)
![case](docs/img/case.png)
![timer](docs/img/timer.png)
![report](docs/img/report.png)
![reportDetail](docs/img/reportDetail.png)
