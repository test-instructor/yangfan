# 快速开始（本地跑通）

本章节用于在本地快速跑通扬帆测试平台的核心链路（推荐按“依赖服务 → server/web → 初始化 → 配置 → run/data”的顺序）。

## 前置条件

- Go：>= 1.23
- Node.js：>= 16.4
- Docker + Docker Compose：用于快速启动 MySQL/Redis/RabbitMQ

## 1. 启动依赖服务（MySQL/Redis/RabbitMQ）

在项目根目录执行：

```bash
docker-compose -f deploy/docker-compose/docker-compose.data.yml up -d
```

默认端口（宿主机）：

- MySQL：43306（root/123456，默认库 yf）
- RabbitMQ：45672（AMQP），45673（管理台）
- Redis：46378

## 2. 首次启动（只启动 server + web）

### 2.1 启动 server（后端 API）

```bash
cd server
go run .
```

### 2.2 启动 web（前端）

```bash
cd web
npm install
npm run serve
```

默认会在 `http://localhost:8080` 启动前端，并把 `/api` 代理到 `http://127.0.0.1:8888`（见 `web/.env.development`）。

## 3. 在前端完成初始化（首次必做）

打开 `http://localhost:8080`，如果系统检测到未初始化，会跳转到“初始化”页面。推荐使用本地 docker-compose 的 MySQL：

- DBType：mysql
- Host：127.0.0.1
- Port：43306
- UserName：root
- Password：123456
- DBName：yf
- AdminPassword：设置管理员账号 `admin` 的密码

提交初始化后，后端会建库/建表/写入初始数据，并把 DB 配置写入配置文件。

## 4. 重启 server，并配置 MQ / 平台配置

初始化完成后，需要重启 server 让配置在进程内生效（MQ 初始化、定时任务、数据仓库端口等都在启动时加载）。

server 重启完成后，在前端配置：

- 系统工具 → 系统配置 → MQ配置：填写 RabbitMQ 连接信息
  - Host：127.0.0.1
  - Port：45672
  - Username/Password：guest/guest（按你的实际环境）
- 系统工具 → 系统配置 → 平台配置：填写数据仓库（data）服务地址
  - Host：127.0.0.1
  - Port：9000（与 data 服务启动端口保持一致）

修改后会落盘到配置文件；若你已启动过 run/data，需要在修改配置后重启它们以重新加载。

## 5. 启动 run、data 服务

```bash
cd data
go run .
```

```bash
cd run
go run .
```

run 默认以 `RUN_SERVICE_MODE=runner` 启动（消费执行队列）。如需同时启用定时调度，可用：

```bash
RUN_SERVICE_MODE=all go run .
```

## 6. 继续配置并沉淀测试资产

1. 创建项目（pm / 项目配置）
2. 配置环境变量（platform / 环境变量管理）
3. 配置运行配置（platform / 运行配置）
4. 沉淀接口（APIAutomation / 接口管理）
5. 封装步骤（APIAutomation / 测试步骤）
6. 编排用例并调试（APIAutomation / 测试用例）
7. 创建定时任务并执行（APIAutomation / 定时任务）
8. 在自动报告查看结果（APIAutomation / 自动报告），并配置报告通知（pm / 报告通知）
