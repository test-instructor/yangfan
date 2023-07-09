# 扬帆测试平台 - 部署文档

![image-20230614184040563](http://qiniu.yangfan.gd.cn/markdown/image-20230614184040563.png)




# 在线demo

首页：http://82.157.150.119:8080/
用户名：admin
密码： 123456

# 部署方式
> 部署前请先安装好数据看，以下所有部分方式都需要手动配置数据库
> 一键部署方式正在开发中，尽请期待
1. 新建数据库，并导入`docs/sql/yangfan.sql`文件
2. 管理员账号`yangfan`,密码`123456`

## shell 脚本一键安装
```shell
# 进入目录
mkdir -p /home/yangfan && cd /home/yangfan
# 执行安装
bash -c "$(curl -fsSLk http://qiniu.yangfan.gd.cn/install/install.sh)"
```
> 1. 脚本默认安装在/data/yangfan目录下
> 2. 如需修改配置请到安装目录下修改`./config/docker.config.yaml`文件
> 3. `./config/docker.config.yaml`中的`grafana-host`需要手动修改为`http://IP:3000`，否则无法访问grafana
> 4. `./config/docker.config.yaml`中的`fs`需要手动修改为飞书登录相关配置，否则无法登录
> 5. 修改配置后重新执行安装脚本即可

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


## docker 部署

* 前端：修改对应`docker-compose`文件中的`ENV_VITE_FS_APP_ID`、`ENV_VITE_FS_LOGIN`
* 后端：修改`deploy/docker-compose/config/docker.config.yaml`中的数据库`mysql`、飞书登录`fs`相关配置
* docker镜像源：目前使用阿里云镜像源(registry.cn-hangzhou.aliyuncs.com)，如需使用docker官方镜像源，请将阿里云镜像源(registry.cn-hangzhou.aliyuncs.com/)删除即可

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




