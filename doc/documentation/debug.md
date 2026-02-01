# 开发调试

本章节用于在本地启动前后端进行开发与问题定位。

## 前端（web）

1. 安装 Node.js（建议 18+）
2. 安装依赖并启动

```shell
cd web
npm install
npm run serve
```

默认端口与后端联调参数见 `web/.env.development`。

## 后端（server）

1. 安装 Go（建议 1.20+）
2. 安装依赖并启动

```shell
cd server
go mod tidy
go run .
```

默认监听端口通常为 `8888`（以你的配置为准）。

## 依赖服务

后端运行依赖 MySQL 等基础服务，建议优先使用 `deploy/docker-compose` 启动依赖服务，再进行本地调试。

## 多服务说明（按需）

仓库中除了 `server` 外，还包含 `run`（用例运行）等服务。你可以分别在对应目录执行：

```shell
go mod tidy
go run .
```
