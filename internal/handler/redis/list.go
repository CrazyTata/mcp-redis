package redis

import (
	"context"
	"mcp-redis/internal/svc"

	"github.com/zeromicro/go-zero/mcp"
)

func RedisLPush(svcCtx *svc.ServiceContext) mcp.Tool {
	redisLPushHandler := func(ctx context.Context, params map[string]any) (any, error) {
		var req struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}
		err := mcp.ParseArguments(params, &req)
		if err != nil {
			return nil, err
		}
		return svcCtx.Redis.LpushCtx(ctx, req.Key, req.Value)
	}

	return mcp.Tool{
		Name:        "redis_lpush",
		Description: "Push an element to the left of a list",
		InputSchema: mcp.InputSchema{
			Properties: map[string]any{
				"key": map[string]any{
					"type":        "string",
					"description": "The key of the list",
				},
				"value": map[string]any{
					"type":        "string",
					"description": "The value to push",
				},
			},
			Required: []string{"key", "value"},
		},
		Handler: redisLPushHandler,
	}
}

func RedisRPush(svcCtx *svc.ServiceContext) mcp.Tool {
	redisRPushHandler := func(ctx context.Context, params map[string]any) (any, error) {
		var req struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}
		err := mcp.ParseArguments(params, &req)
		if err != nil {
			return nil, err
		}
		return svcCtx.Redis.RpushCtx(ctx, req.Key, req.Value)
	}

	return mcp.Tool{
		Name:        "redis_rpush",
		Description: "Push an element to the right of a list",
		InputSchema: mcp.InputSchema{
			Properties: map[string]any{
				"key": map[string]any{
					"type":        "string",
					"description": "The key of the list",
				},
				"value": map[string]any{
					"type":        "string",
					"description": "The value to push",
				},
			},
			Required: []string{"key", "value"},
		},
		Handler: redisRPushHandler,
	}
}

func RedisLpop(svcCtx *svc.ServiceContext) mcp.Tool {
	redisLpopHandler := func(ctx context.Context, params map[string]any) (any, error) {
		var req struct {
			Key string `json:"key"`
		}
		err := mcp.ParseArguments(params, &req)
		if err != nil {
			return nil, err
		}
		return svcCtx.Redis.LpopCtx(ctx, req.Key)
	}

	return mcp.Tool{
		Name:        "redis_lpop",
		Description: "Pop an element from the left of a list",
		InputSchema: mcp.InputSchema{
			Properties: map[string]any{
				"key": map[string]any{
					"type":        "string",
					"description": "The key of the list",
				},
			},
			Required: []string{"key"},
		},
		Handler: redisLpopHandler,
	}
}

func RedisRpop(svcCtx *svc.ServiceContext) mcp.Tool {
	redisRpopHandler := func(ctx context.Context, params map[string]any) (any, error) {
		var req struct {
			Key string `json:"key"`
		}
		err := mcp.ParseArguments(params, &req)
		if err != nil {
			return nil, err
		}
		return svcCtx.Redis.RpopCtx(ctx, req.Key)
	}

	return mcp.Tool{
		Name:        "redis_rpop",
		Description: "Pop an element from the right of a list",
		InputSchema: mcp.InputSchema{
			Properties: map[string]any{
				"key": map[string]any{
					"type":        "string",
					"description": "The key of the list",
				},
			},
			Required: []string{"key"},
		},
		Handler: redisRpopHandler,
	}
}

func RedisLrange(svcCtx *svc.ServiceContext) mcp.Tool {
	redisLrangeHandler := func(ctx context.Context, params map[string]any) (any, error) {
		var req struct {
			Key   string  `json:"key"`
			Start float64 `json:"start"`
			End   float64 `json:"end"`
		}
		err := mcp.ParseArguments(params, &req)
		if err != nil {
			return nil, err
		}
		return svcCtx.Redis.LrangeCtx(ctx, req.Key, int(req.Start), int(req.End))
	}

	return mcp.Tool{
		Name:        "redis_lrange",
		Description: "Get a range of elements from a list",
		InputSchema: mcp.InputSchema{
			Properties: map[string]any{
				"key": map[string]any{
					"type":        "string",
					"description": "The key of the list",
				},
				"start": map[string]any{
					"type":        "number",
					"description": "The start index",
				},
				"end": map[string]any{
					"type":        "number",
					"description": "The end index",
				},
			},
			Required: []string{"key", "start", "end"},
		},
		Handler: redisLrangeHandler,
	}
}

func RedisLlen(svcCtx *svc.ServiceContext) mcp.Tool {
	redisLlenHandler := func(ctx context.Context, params map[string]any) (any, error) {
		var req struct {
			Key string `json:"key"`
		}
		err := mcp.ParseArguments(params, &req)
		if err != nil {
			return nil, err
		}
		return svcCtx.Redis.LlenCtx(ctx, req.Key)
	}

	return mcp.Tool{
		Name:        "redis_llen",
		Description: "Get the length of a list",
		InputSchema: mcp.InputSchema{
			Properties: map[string]any{
				"key": map[string]any{
					"type":        "string",
					"description": "The key of the list",
				},
			},
			Required: []string{"key"},
		},
		Handler: redisLlenHandler,
	}
}
