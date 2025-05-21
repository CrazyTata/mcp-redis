package main

import (
	"flag"
	"fmt"

	"mcp-redis/internal/config"
	"mcp-redis/internal/handler"
	"mcp-redis/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/mcp"
)

var (
	configFile = flag.String("f", "etc/config-api.yaml", "the config file")
)

func main() {
	// 解析命令行参数s
	flag.Parse()

	// 禁用日志统计
	logx.DisableStat()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	mcpSrv := mcp.NewMcpServer(c.McpConf)
	defer mcpSrv.Stop()

	svcCtx := svc.NewServiceContext(c)
	handler.NewToolsetHandler(mcpSrv, svcCtx)

	fmt.Printf("Starting MCP Server on %s:%d\n", c.McpConf.Host, c.McpConf.Port)
	mcpSrv.Start()
}
