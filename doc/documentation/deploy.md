# 部署服务（服务端）

本文档描述扬帆测试平台在服务器上的推荐部署流程（Docker Compose 为主），并补充 run 服务多节点/异地部署要点。

## 1. 服务与依赖说明

### 1.1 依赖中间件

- MySQL：业务数据
- RabbitMQ：server ↔ run 的任务投递与定时任务控制
- Redis：可选（按配置启用）

依赖服务的默认镜像、端口与本地快速启动方式，见“项目前置依赖服务清单”文档内容（本页后面也有摘要）。

### 1.2 核心服务

- server：后端 API、平台初始化、MQ 初始化、定时任务管理等
- web：前端（Nginx/Vite 构建产物）
- run：用例运行服务（可多实例、多机部署）
- data：数据仓库服务（对外提供数据读取/写入能力）

## 2. 首次部署推荐流程（分阶段启动）

首次部署建议按“先 server/web → 初始化 → 配置 → 再 run/data”的顺序，避免 run/data 在基础配置未就绪时反复报错或连接失败。

### 2.1 启动依赖服务

```bash
cd deploy/docker-compose
docker-compose -f docker-compose.data.yml up -d
```

### 2.2 首次只启动 server、web

```bash
cd deploy/docker-compose
docker-compose -f docker-compose.yml up -d server web
```

server 默认端口：8888  
web 默认端口：8080

### 2.3 在前端完成初始化

访问 `http://<你的服务器IP>:8080`，如果系统检测到未初始化，会跳转到“初始化”页面，填写数据库连接信息并提交。

Docker Compose 同网络下的常见填法（以默认 compose 为例）：

- DBType：mysql
- Host：yangfan-mysql
- Port：3306
- UserName：root
- Password：123456
- DBName：yf
- AdminPassword：设置管理员账号 `admin` 的密码

初始化会：

- 建库/建表/写入初始数据
- 把 DB 配置写入 server 配置文件

在 compose 模式下，配置文件默认映射为主机侧的：

- `deploy/docker-compose/config/docker.config.yaml`

### 2.4 重启 server，并配置 MQ配置 / 平台配置

初始化完成后，需要重启 server 让配置在进程内生效（MQ 初始化、定时任务、数据仓库端口等都在启动时加载）。

```bash
docker restart yangfan-server
```

随后在前端配置：

- 系统工具 → 系统配置 → MQ配置：填写 MQ 连接信息（host/port/账号等）
- 系统工具 → 系统配置 → 平台配置：填写 data（数据仓库）服务地址（host/port）

配置保存后会落盘到配置文件；如果 run/data 已启动，需要重启它们以重新加载配置。

### 2.5 启动 run、data

```bash
cd deploy/docker-compose
docker-compose -f docker-compose.yml up -d run data
```

如果你只想启用定时调度（不消费执行队列），可启动 `run-timer`；若同时启用执行 + 定时，可启动 `run-all`（见 `docker-compose.yml` 内的示例 service）。

## 3. run 多节点/异地部署

run 服务天然支持多实例（同机多进程、同集群多 Pod、跨机房部署）。其核心工作方式是：

- 每个 run 节点都连接同一个 MySQL（读写业务数据、上报节点心跳/状态）
- 每个 run 节点都连接同一个 MQ（按节点名消费自己的队列）
- server 根据节点名把任务投递到对应队列，或按策略选择节点

### 3.1 必要条件

- 所有 run 节点必须能访问 MySQL 与 MQ
- 每个 run 节点必须使用唯一的节点名

### 3.2 关键环境变量

- `RUN_SERVICE_MODE`
  - `runner`：仅消费执行任务队列（默认）
  - `timer`：仅启用定时调度
  - `all`：同时启用执行 + 定时
- `NODE_NAME`：节点名称（必须唯一）
- `NODE_ALIAS`：节点别名（可选）

队列命名规则：

- 实际队列名为：`<queue-prefix><NODE_NAME>`（`queue-prefix` 来自 MQ配置）

### 3.3 多节点部署方式示例（Docker Compose 扩容思路）

通过复制 run service 并设置不同的 `NODE_NAME` 扩容多个执行节点，例如：

- runner 节点：`RUN_SERVICE_MODE=runner`，`NODE_NAME=runner-1/runner-2/...`
- timer 节点：`RUN_SERVICE_MODE=timer`，`NODE_NAME=timer-1/...`

## 4. 项目前置依赖服务清单（摘要）

本项目主要依赖 MySQL、Redis、RabbitMQ。推荐在本地开发或测试环境中，使用 `docker-compose.data.yml` 快速启动：

```bash
docker-compose -f deploy/docker-compose/docker-compose.data.yml up -d
```

默认端口（宿主机）：

- MySQL：43306（容器 3306）
- Redis：46378（容器 6378）
- RabbitMQ：45672（容器 5672），45673（容器 15672）

