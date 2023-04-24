## docker 部署server方式

### 创建docker网络服务

无论是本地还是离线部署均需要创建网络服务因为使用的docker网络模式为host，所以需要创建一个docker网络服务，否则无法访问

```bash
docker network create yangfan
```

#### 构建镜像

1. 进入server目录
2. 执行

```bash 
docker build -t server:1.0 .
```

#### 本地启动容器部署

1. 修改配置文件
2. 复制配置文件到 /home/yangfan/config.yaml
3. 执行

```bash 
docker run -d -p 8888:8888 --network yangfan -v /home/yangfan/config.yaml:/yangfan/config.yaml --name yangfan-server
   server:1.0
```

### 离线部署方式

####           

1. 本地导出镜像

```bash
docker save -o server.tar server:1.0
```

1. 上传到服务器
2. 导入镜像
3. 上传配置文件到服务器
4. 执行

```bash
docker load -i server.tar
```

5. 执行

```bash
docker run -d -p 8888:8888 --network yangfan -v /home/yangfan/config.yaml:/yangfan/config.yaml --name yangfan-server
   server:1.0
```

### 问题

* 启动报错
  failed to initialize database, got error dial tcp 127.0.0.1:3306: connect: connection refusedfailed to initialize
  database, got error dial tcp 127.0.0.1:3306: connect: connection refused
  **_检查数据库配置是否正确，原因为未挂在成功本地配置文件_**

**注意：**

* 如果服务器不能访问外网必须在本地可以访问外网进行构建，或者在公司docker管理平台进行构建
* 这里默认是已经初始化过数据库的，如果是首次部署需要先初始化数据库，然后再进行部署
* 这里未使用docker-compose进行项目和中间件部署，需要独立搭建mysql、redis等中间件，或者使用已有的中间件