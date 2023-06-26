# 扬帆测试平台 - 部署文档

![image-20230614184040563](http://qiniu.yangfan.gd.cn/markdown/image-20230614184040563.png)




# 在线demo

首页：http://82.157.150.119:8080/
用户名：admin
密码： 123456

# 部署方式
> 飞书登录的前端信息暂时不支持动态配置，飞书登录需要修改后重新构建镜像

## 本地调试
### 后端
1. 下载golang安装 版本号需>=1.18
   * 国际: https://golang.org/dl/
   * 国内: https://golang.google.cn/dl/
2. goland 打开项目根目录
3. 修改`config.yaml`中的数据库`mysql`、飞书登录`fs`相关配置
4. 使用软件包进行运行，目前已有的软件包为
   ```shell
   github.com/test-instructor/yangfan/server  # 后端服务
   github.com/test-instructor/yangfan/run     # 用例运行服务
   github.com/test-instructor/yangfan/master  # 性能测试master服务
   github.com/test-instructor/yangfan/work    # 性能测试worker服务
   github.com/test-instructor/yangfan/timer   # 定时任务服务
   ```
### 前端
1. 前往https://nodejs.org/zh-cn/下载当前版本node 
2. 命令行运行 node -v 若控制台输出版本号则前端环境搭建成功 
3. node 版本需大于 16.4 
4. 开发工具推荐vscode https://code.visualstudio.com/

## 


1. 新建数据库，并导入`docs/sql/yangfan.sql`文件
2. 修改`web/.env.development`、 `web/.env.production`中的`VITE_FS_LOGIN`、`VITE_FS_APP_ID`
3. 管理员账号`yangfan`,密码`123456`

## 本地调试

## docker 部署

* 前端：修改对应`docker-compose`文件中的`ENV_VITE_FS_APP_ID`、`ENV_VITE_FS_LOGIN`
* 后端：修改`deploy/docker-compose/config/docker.config.yaml`中的数据库`mysql`、飞书登录`fs`相关配置

1. 本地构建模式文件：`deploy/docker-compose/docker-compose-build.yaml`
2. 远程镜像模式文件: `deploy/docker-compose/docker-compose-image.yaml`
3. 执行命令：
   ```shell
   cd deploy/docker-compose
   # 本地构建模式
   docker-compose up --build -f docker-compose-build.yaml --force-recreate -d
   # 远程镜像模式
   docker-compose up -f docker-compose-image.yaml
      
   ```
   
## k8s 部署
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




