package handler

import (
	"log"
	"mcp-redis/internal/handler/redis"
	"mcp-redis/internal/svc"

	"github.com/zeromicro/go-zero/mcp"
)

type ToolsetHandler struct {
	svcCtx *svc.ServiceContext
	mcpSrv mcp.McpServer
}

func NewToolsetHandler(mcpSrv mcp.McpServer, svcCtx *svc.ServiceContext) *ToolsetHandler {
	t := &ToolsetHandler{
		svcCtx: svcCtx,
		mcpSrv: mcpSrv,
	}
	t.registerTool()
	return t
}

func (t *ToolsetHandler) registerTool() {
	if err := t.mcpSrv.RegisterTool(redis.RedisSet(t.svcCtx)); err != nil {
		log.Fatalf("register tool error: %v", err)
	}
	if err := t.mcpSrv.RegisterTool(redis.RedisGet(t.svcCtx)); err != nil {
		log.Fatalf("register tool error: %v", err)
	}
	if err := t.mcpSrv.RegisterTool(redis.RedisDel(t.svcCtx)); err != nil {
		log.Fatalf("register tool error: %v", err)
	}
	if err := t.mcpSrv.RegisterTool(redis.RedisLPush(t.svcCtx)); err != nil {
		log.Fatalf("register tool error: %v", err)
	}
	if err := t.mcpSrv.RegisterTool(redis.RedisRPush(t.svcCtx)); err != nil {
		log.Fatalf("register tool error: %v", err)
	}
	if err := t.mcpSrv.RegisterTool(redis.RedisLpop(t.svcCtx)); err != nil {
		log.Fatalf("register tool error: %v", err)
	}
	if err := t.mcpSrv.RegisterTool(redis.RedisRpop(t.svcCtx)); err != nil {
		log.Fatalf("register tool error: %v", err)
	}
	if err := t.mcpSrv.RegisterTool(redis.RedisLrange(t.svcCtx)); err != nil {
		log.Fatalf("register tool error: %v", err)
	}
	if err := t.mcpSrv.RegisterTool(redis.RedisLlen(t.svcCtx)); err != nil {
		log.Fatalf("register tool error: %v", err)
	}
}
