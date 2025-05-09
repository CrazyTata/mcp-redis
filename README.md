# Redis MCP (Model Context Protocol)

Redis MCP 是一个基于 go-zero 框架开发的 Redis 模型上下文协议实现，支持多种交互方式，包括 SSE（Server-Sent Events）和命令行界面。它提供了一个统一的接口来管理和操作 Redis 数据，使开发者能够更方便地集成 Redis 功能到他们的应用中。

## 功能特点

- 支持多种交互方式：
  - SSE（Server-Sent Events）实时通信
  - 命令行界面（CLI）
- 完整的 Redis 数据类型支持：
  - 字符串（String）
  - 哈希表（Hash）
  - 列表（List）
  - 集合（Set）
  - 有序集合（Sorted Set）
- 统一的错误处理和响应格式
- 基于 go-zero 框架，提供高性能和可扩展性
- 支持 Cursor IDE 集成
- 符合 Model Context Protocol 规范

## 快速开始

### 环境要求

- Go 1.16 或更高版本
- Redis 服务器
- Cursor IDE（可选，用于 IDE 集成）

### 安装

1. 克隆仓库：
```bash
git clone https://github.com/yourusername/mcp-redis.git
cd mcp-redis
```

2. 安装依赖：
```bash
go mod tidy
```

### 配置

1. 修改配置文件 `etc/config-api.yaml`：
```yaml
Name: mcp-redis
Host: 0.0.0.0
Port: 8083

Redis:
  Host: localhost:6379
  Type: node
  Pass: "your-password"
  DB: 0
  PoolSize: 100
  MinIdleConns: 10
  MaxRetries: 3
  IdleTimeout: 240
  MaxConnAge: 3600
```

2. 配置 Cursor IDE（可选）：
在 `~/.cursor/mcp.json` 中添加：
```json
{
  "mcpServers": {
    "redis": {
      "name": "Redis MCP",
      "description": "Redis Model Context Protocol",
      "version": "1.0.0",
      "url": "http://localhost:8083/sse",
      "transport": "sse"
    },
    "redis-cli": {
      "name": "Redis MCP CLI",
      "description": "Redis Model Context Protocol CLI",
      "version": "1.0.0",
      "command": "go run main.go -m stdio",
      "transport": "stdio",
      "workingDir": "/path/to/mcp-redis"
    }
  }
}
```

### 运行

1. HTTP 模式（默认）：
```bash
go run main.go
```

2. 命令行模式：
```bash
go run main.go -m stdio
```

## API 文档

### SSE API

通过 `/sse` 端点提供实时通信：

```json
// 发送命令
{
  "type": "command.type",
  "payload": {
    "key": "test",
    "value": "hello"
  }
}

// 接收响应
{
  "type": "command.type.response",
  "payload": {
    "code": 200,
    "message": "success",
    "data": "hello"
  }
}
```

### 命令行界面

支持的命令：
```
set <key> <value>  - 设置键值
get <key>         - 获取键值
del <key>         - 删除键
exists <key>      - 检查键是否存在
help              - 显示帮助信息
exit              - 退出程序
```

## 项目结构

```
mcp-redis/
├── etc/                 # 配置文件
├── internal/           # 内部代码
│   ├── config/        # 配置结构
│   ├── handler/       # 处理器
│   └── svc/          # 服务上下文
├── main.go            # 主程序
└── README.md         # 项目文档
```

## 贡献指南

1. Fork 项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建 Pull Request

## 许可证

本项目采用 MIT 许可证
