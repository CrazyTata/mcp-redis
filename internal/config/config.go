package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/mcp"
)

type Config struct {
	Redis   redis.RedisConf
	McpConf mcp.McpConf
}
