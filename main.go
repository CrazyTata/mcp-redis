package main

import (
	"flag"
	"fmt"
	"net/http"

	"mcp-redis/internal/config"
	"mcp-redis/internal/handler"
	"mcp-redis/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var (
	configFile = flag.String("f", "etc/config-api.yaml", "the config file")
	mode       = flag.String("m", "http", "运行模式: http 或 stdio")
)

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)

	switch *mode {
	case "http":
		server := rest.MustNewServer(c.RestConf)
		defer server.Stop()

		// handler.RegisterHandlers(server, ctx)

		// 注册 SSE 处理器
		sseHandler := handler.NewSSEHandler(ctx)
		server.AddRoute(rest.Route{
			Method:  http.MethodPost,
			Path:    "/sse",
			Handler: http.HandlerFunc(sseHandler.ServeHTTP),
		})

		fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
		server.Start()

	case "stdio":
		stdioHandler := handler.NewStdioHandler(ctx)
		stdioHandler.Start()

	default:
		fmt.Printf("未知的运行模式: %s\n", *mode)
		fmt.Println("支持的模式: http 或 stdio")
	}
}
