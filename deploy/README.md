# 扬帆测试平台 - 部署文档

![image-20230614184040563](http://qiniu.yangfan.gd.cn/markdown/image-20230614184040563.png)




# 在线demo

首页：http://82.157.150.119:8080/
用户名：admin
密码： 123456

# 部署方式
> 部署前请先准备好依赖服务（MySQL / RabbitMQ / Redis 可选）。
> 首次启动推荐通过前端“初始化”页面完成建库建表与管理员初始化，而不是手动导入 SQL。

使用文档：

- 本地调试：[本地调试.md](file:///Users/taylor/Documents/yangfan/yangfan/docs/使用文档/本地调试.md)
- 服务端部署：[服务端部署.md](file:///Users/taylor/Documents/yangfan/yangfan/docs/使用文档/服务端部署.md)
- 依赖服务清单：[项目前置依赖服务清单.md](file:///Users/taylor/Documents/yangfan/yangfan/docs/使用文档/项目前置依赖服务清单.md)


* 前端：修改对应`docker-compose`文件中的`ENV_VITE_FS_APP_ID`、`ENV_VITE_FS_LOGIN`
* docker镜像源：目前使用阿里云镜像源(registry.cn-hangzhou.aliyuncs.com)，如需使用docker官方镜像源，请将阿里云镜像源(registry.cn-hangzhou.aliyuncs.com/)删除即可

1. 本地构建模式文件：`deploy/docker-compose/docker-compose-build.yaml`
2. 远程镜像模式文件: `deploy/docker-compose/docker-compose-image.yaml`
3. 执行命令：
   ```shell
   cd deploy/docker-compose
   # 数据库相关配置，mysql、mq、redis
   docker-compose up -f docker-compose.data.yml
   docker-compose -f docker-compose.data.yml up -d
   docker-compose up -f docker-compose.yml
   docker-compose -f docker-compose.yml up -d
   ```
   
   
## 三、k8s 部署
文件目录`./deploy/kubernetes`
```shell
kubernetes
    ├── grafana-prometheus-pushgateway    # 性能测试报告监控
    ├── httpbin                           # http、grpc demo
    ├── k8s_yangfan.yaml                  # 部署文件
    ├── server                            # 后端部署文件
    ├── web                               # 前端部署文件
    └── yangfan-namespace.yaml            # 命名空间
```

1. 修改`ConfigMap/docker-config-yaml`中的数据库`mysql`、飞书登录`fs`相关配置
2. 修改`Deployment/yangfan-web`中的`ENV_VITE_FS_APP_ID`、`ENV_VITE_FS_LOGIN`
3. 执行命令：
   ```shell
   cd deploy/kubernetes
   kubectl apply -f k8s_yangfan.yaml
   ```



