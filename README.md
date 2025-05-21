
# mcp-redis

mcp-redis 是一个基于 go-zero 框架的 Redis Model Context Protocol (MCP) 实现，支持 SSE（Server-Sent Events）。它为开发者提供了统一、便捷的 Redis 数据管理接口，适用于高性能、可扩展的服务场景，并支持与 Cursor IDE 集成。

---

## 目录

- [项目简介](#项目简介)
- [功能特性](#功能特性)
- [项目结构](#项目结构)
- [环境要求](#环境要求)
- [安装与配置](#安装与配置)
- [运行方式](#运行方式)
- [支持的 Redis 操作](#支持的-redis-操作)
- [API 交互说明](#api-交互说明)
- [Cursor IDE 集成](#cursor-ide-集成)
- [贡献指南](#贡献指南)
- [许可证](#许可证)
- [常见问题](#常见问题)

---

## 项目简介

mcp-redis 旨在为 Redis 提供标准化的模型上下文协议支持，简化 Redis 在多种应用场景下的集成与开发。通过 go-zero 框架实现高性能服务，支持实时通信和命令行操作，适合微服务、云原生等现代开发环境。

---

## 功能特性

- **完整的 Redis 数据类型支持**：字符串、列表等，还可以继续扩展其他数据类型
- **统一的错误处理与响应格式**：便于前后端协作
- **高性能与可扩展性**：基于 go-zero 框架
- **Cursor IDE 集成**：提升开发效率
- **符合 Model Context Protocol 规范**

---

## 项目结构

```
mcp-redis/
├── etc/                 # 配置文件
│   └── config-api.yaml  # 服务和 Redis 配置
├── internal/            # 内部代码
│   ├── config/          # 配置结构体
│   ├── handler/         # 路由与处理器
│   │   ├── toolshandler.go
│   │   └── redis/       # Redis 相关命令实现
│   │       ├── list.go
│   │       └── string.go
│   └── svc/             # 服务上下文
│       └── servicecontext.go
├── main.go              # 启动入口
├── go.mod               # Go 依赖管理
└── README.md            # 项目文档
```

---

## 环境要求

- Go 1.16 及以上
- Redis 服务器（建议 5.0+）
- Cursor IDE（可选）

---

## 安装与配置

### 1. 克隆项目

```bash
git clone https://github.com/yourusername/mcp-redis.git
cd mcp-redis
```

### 2. 安装依赖

```bash
go mod tidy
```

### 3. 配置服务

编辑 `etc/config-api.yaml`：

```yaml
Redis:
  Host: localhost:46379
  Type: node
  Pass: "123456"
  DB: 0
  PoolSize: 100
  MinIdleConns: 10
  MaxRetries: 3
  IdleTimeout: 240
  MaxConnAge: 3600

McpConf:
  Name: mcp-redis
  Host: 0.0.0.0
  Port: 8083
```

---

## 运行方式

### HTTP（SSE）模式

```bash
go run main.go
```

服务将监听 8083 端口，提供 SSE 实时通信接口。

---

## 支持的 Redis 操作

### 字符串相关

- `redis_set`：设置键值（支持过期时间）
- `redis_get`：获取键值
- `redis_del`：删除键

### 列表相关

- `redis_lpush`：左侧插入元素
- `redis_rpush`：右侧插入元素
- `redis_lpop`：左侧弹出元素
- `redis_rpop`：右侧弹出元素
- `redis_lrange`：获取区间元素
- `redis_llen`：获取列表长度

每个命令的参数和返回格式详见代码注释和 mcp.Tool 的 InputSchema。

---

## API 交互说明

### SSE API

- **请求格式**（发送命令）：

  ```json
  {
    "type": "redis_set",
    "payload": {
      "key": "test",
      "value": "hello",
      "expire": 60
    }
  }
  ```

- **响应格式**（接收响应）：

  ```json
  {
    "type": "redis_set.response",
    "payload": {
      "code": 200,
      "message": "success",
      "data": "写入成功"
    }
  }
  ```


---

## Cursor IDE 集成

可选步骤，提升开发体验。

1. 在 `~/.cursor/mcp.json` 添加：

   ```json
   {
     "mcpServers": {
        "redis": {
          "name": "Redis MCP",
          "description": "Redis Management Control Panel",
          "command": "npx",
          "args": ["mcp-remote", "http://localhost:8083/sse"]
        }
   }
   ```

---

