# 扬帆测试平台 - 部署文档

![image-20230614184040563](http://qiniu.yangfan.gd.cn/markdown/image-20230614184040563.png)




# 在线demo

首页：http://82.157.150.119:8080/
用户名：admin
密码： 123456

# 部署方式
> 飞书登录的前端信息暂时不支持动态配置，飞书登录需要修改后重新构建镜像

## 


1. 新建数据库，并导入`docs/sql/yangfan.sql`文件
2. 修改`web/.env.development`、 `web/.env.production`中的`VITE_FS_LOGIN`、`VITE_FS_APP_ID`
3. 管理员账号`yangfan`,密码`123456`

## 本地调试

## docker 部署

* 前端：如果需要飞书登录则需要修改env:`web/.env.production`中的`VITE_FS_LOGIN`、`VITE_FS_APP_ID`
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

