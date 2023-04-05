## docker 部署web方式

### 创建docker网络服务

* 无论是本地还是离线部署均需要创建网络服务因为使用的docker网络模式为host，所以需要创建一个docker网络服务，否则无法访问,如果server之前已经创建过网络服务则不需要再次创建
* 如果是分开部署,则需要修改.docker-compose/nginx/conf.d/my.conf文件中/api下的proxy_pass地址为server的地址
* .env.production文件中的VITE_BASE_PATH为server的地址

#### 构建镜像

1. 进入web目录
2. 执行

```bash
docker build -t web:1.0 .
```

#### 本地启动容器部署

```bash
docker run -d -p 8080:8080 --network cheetah --name cheetah-web web:1.0
```

#### 离线部署方式

1. 本地导出镜像

```bash
docker save -o web.tar web:1.0
```

2. 上传到服务器
3. 导入镜像
4. 执行

```bash
docker load -i web.tar
```

5. 执行

```bash
docker run -d -p 8080:8080 --network cheetah --name cheetah-web web:1.0
```

**注意:** 前端构建较慢,请耐心等待

