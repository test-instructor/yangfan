# 部署服务

平台支持 Docker Compose 与 Kubernetes 两种部署方式。部署前请先准备好依赖服务（MySQL / RabbitMQ / Redis 可选，按你的实际配置启用）。

## Docker Compose（推荐）

### 1. 配置文件

- 后端配置：`deploy/docker-compose/config/docker.config.yaml`（数据库、飞书登录等）
- 前端环境变量：`deploy/docker-compose/*.yml` 中的 `ENV_VITE_FS_APP_ID`、`ENV_VITE_FS_LOGIN`

### 2. 启动方式

`deploy/docker-compose` 目录下提供了不同模式的 compose 文件（本地构建/镜像模式/调试模式），按需选择其一启动即可。

## Kubernetes

部署文件位于 `deploy/kubernetes`，主要入口为 `k8s_yangfan.yaml`。通常需要：

1. 修改 ConfigMap 中的数据库与登录配置
2. 修改 Web Deployment 的前端环境变量
3. `kubectl apply -f k8s_yangfan.yaml`

## 首次初始化

首次启动后，建议通过前端“初始化”页面完成建库建表与管理员初始化，而不是手动导入 SQL。

## 更完整的部署说明

请参考项目中的部署文档：[deploy/README.md](file:///Users/taylor/Documents/yangfan/python/yangfan/deploy/README.md)

