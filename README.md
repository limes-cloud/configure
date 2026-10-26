# configure —— 百度内部配置中心服务

> 本服务是 **百度内部** 统一应用管理平台（GoPlatform）的**配置中心**，为平台内各服务（manager / resource / ai-agent / cron 等）提供配置的集中管理、下发与热更新能力。

## 功能定位

- **配置管理**：配置（`configure`）的读取、修改、比对（compare）与**实时监听**（Watch，gRPC stream 下发）；服务启动时经 `CONF_HOST / CONF_TOKEN / APP_NAME` 环境变量从本中心拉取自身配置
- **业务配置**（`business`）：按业务维度组织与维护配置项（values）
- **环境管理**（`env`）：环境（dev / test / prod 等）与 Token 的申请、重置（供客户端拉取配置鉴权）
- **资源配置**（`resource`）：资源视图与值的管理
- **模板管理**（`template`）：配置模板与模板命名空间，支持配置的模板化生成

## 技术栈

| 维度 | 选型 |
|---|---|
| 语言 | Go 1.24.6 |
| 微服务框架 | go-kratos/kratos + limes-cloud/kratosx |
| 数据库 | MySQL（GORM） |
| 缓存 | Redis |
| 鉴权 | 经 `manager/api/authorize` 的 `AuthMiddleware` 统一校验（登录态 + X-App-Key 应用识别） |
| Proto | protoc + kratosx proto 插件，路由统一 `/configure/api/*` |

## 服务模块与路由

| 模块 | 说明 | 主要路由 |
|---|---|---|
| configure | 配置 CRUD / 比对 / 监听 | `GET\|PUT /configure/api/configure`、`POST /configure/api/configure/compare`、`WatchConfigure`（stream） |
| business | 业务配置与值 | `/configure/api/business*` |
| env | 环境与 Token | `/configure/api/env*` |
| resource | 资源与值 | `/configure/api/resource*` |
| template | 模板与命名空间 | `/configure/api/template*` |

## 目录结构

```
server/configure/
├── api/        # proto 定义与生成代码（独立 go 模块）
├── internal/   # core（上下文/配置）、app（handler 自注册）、domain（entity/repository/service）、
│               # infra（dbs 仓储 / 客户端）、initiator（启动自注册）、types（DTO）、pkg（工具）
├── conf/       # 服务配置（端口 / DB / Redis）
├── main.go     # 入口：kratosx.InitApp + 自注册
└── Dockerfile  # 多阶段构建
```

## 启动与构建

```bash
# 本地运行（默认读取 conf/ 配置）
go run main.go

# 构建
go build -o configure main.go

# 生成 proto（api 目录下执行）
kratosx proto client api/configure/proto/configure_service.proto
```

默认端口：HTTP **6091** / gRPC **6092**。

## 客户端接入

其他服务通过环境变量接入本配置中心：

```bash
export CONF_HOST=<https://configure 服务地址>
export CONF_TOKEN=<env 管理下发的 Token>
export APP_NAME=<服务名>
```

配置加载逻辑见 `internal/core/`，与 limes-cloud/manager 的配置加载约定一致。

## 相关服务

- [manager](../manager/README.md) —— 统一应用管理平台管理端
- [resource](../resource/README.md) —— 资源中心（对象存储 / 文件 / 目录 / 应用配置）