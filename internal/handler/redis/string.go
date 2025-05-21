package redis

import (
	"context"
	"fmt"
	"mcp-redis/internal/svc"

	"github.com/zeromicro/go-zero/mcp"
)

func RedisSet(svcCtx *svc.ServiceContext) mcp.Tool {
	redisSetHandler := func(ctx context.Context, params map[string]any) (any, error) {
		var req struct {
			Key    string  `json:"key"`
			Value  string  `json:"value"`
			Expire float64 `json:"expire"`
		}
		err := mcp.ParseArguments(params, &req)
		if err != nil {
			return nil, fmt.Errorf("参数解析错误: %v", err)
		}

		err = svcCtx.Redis.SetexCtx(ctx, req.Key, req.Value, int(req.Expire))
		if err != nil {
			return nil, fmt.Errorf("Redis写入失败: %v", err)
		}

		return "写入成功", nil
	}

	return mcp.Tool{
		Name:        "redis_set",
		Description: "Set a Redis key",
		InputSchema: mcp.InputSchema{
			Properties: map[string]any{
				"key": map[string]any{
					"type":        "string",
					"description": "Redis key",
				},
				"value": map[string]any{
					"type":        "string",
					"description": "Redis value",
				},
				"expire": map[string]any{
					"type":        "number",
					"description": "Expire time in seconds",
				},
			},
			Required: []string{"key", "value"},
		},
		Handler: redisSetHandler,
	}
}

func RedisGet(svcCtx *svc.ServiceContext) mcp.Tool {
	redisGetHandler := func(ctx context.Context, params map[string]any) (any, error) {
		var req struct {
			Key string `json:"key"`
		}
		err := mcp.ParseArguments(params, &req)
		if err != nil {
			return nil, err
		}
		return svcCtx.Redis.GetCtx(ctx, req.Key)
	}

	return mcp.Tool{
		Name:        "redis_get",
		Description: "Get a Redis key",
		InputSchema: mcp.InputSchema{
			Properties: map[string]any{
				"key": map[string]any{
					"type":        "string",
					"description": "Redis key",
				},
			},
			Required: []string{"key"},
		},
		Handler: redisGetHandler,
	}
}

func RedisDel(svcCtx *svc.ServiceContext) mcp.Tool {
	redisDelHandler := func(ctx context.Context, params map[string]any) (any, error) {
		var req struct {
			Key string `json:"key"`
		}
		err := mcp.ParseArguments(params, &req)
		if err != nil {
			return nil, err
		}
		return svcCtx.Redis.DelCtx(ctx, req.Key)
	}

	return mcp.Tool{
		Name:        "redis_del",
		Description: "Delete a Redis key",
		InputSchema: mcp.InputSchema{
			Properties: map[string]any{
				"key": map[string]any{
					"type":        "string",
					"description": "Redis key",
				},
			},
			Required: []string{"key"},
		},
		Handler: redisDelHandler,
	}
}
