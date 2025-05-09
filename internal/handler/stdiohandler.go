package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"mcp-redis/internal/svc"
)

type StdioHandler struct {
	svcCtx *svc.ServiceContext
}

func NewStdioHandler(svcCtx *svc.ServiceContext) *StdioHandler {
	return &StdioHandler{
		svcCtx: svcCtx,
	}
}

func (h *StdioHandler) Start() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Redis MCP CLI - 输入 'exit' 退出")
	fmt.Println("支持的命令格式: <command> <key> [value]")
	fmt.Println("示例: set test hello")

	for {
		fmt.Print("redis> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("读取输入错误: %v\n", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "exit" {
			break
		}

		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}

		command := strings.ToLower(parts[0])
		switch command {
		case "set":
			if len(parts) != 3 {
				fmt.Println("用法: set <key> <value>")
				continue
			}
			key, value := parts[1], parts[2]
			err := h.svcCtx.Redis.Set(key, value)
			if err != nil {
				fmt.Printf("设置键值失败: %v\n", err)
			} else {
				fmt.Println("OK")
			}

		case "get":
			if len(parts) != 2 {
				fmt.Println("用法: get <key>")
				continue
			}
			key := parts[1]
			val, err := h.svcCtx.Redis.Get(key)
			if err != nil {
				fmt.Printf("获取键值失败: %v\n", err)
			} else {
				fmt.Println(val)
			}

		case "del":
			if len(parts) != 2 {
				fmt.Println("用法: del <key>")
				continue
			}
			key := parts[1]
			_, err := h.svcCtx.Redis.Del(key)
			if err != nil {
				fmt.Printf("删除键失败: %v\n", err)
			} else {
				fmt.Println("OK")
			}

		case "exists":
			if len(parts) != 2 {
				fmt.Println("用法: exists <key>")
				continue
			}
			key := parts[1]
			exists, err := h.svcCtx.Redis.Exists(key)
			if err != nil {
				fmt.Printf("检查键存在失败: %v\n", err)
			} else {
				fmt.Println(exists)
			}

		case "help":
			fmt.Println("支持的命令:")
			fmt.Println("  set <key> <value>  - 设置键值")
			fmt.Println("  get <key>         - 获取键值")
			fmt.Println("  del <key>         - 删除键")
			fmt.Println("  exists <key>      - 检查键是否存在")
			fmt.Println("  help              - 显示帮助信息")
			fmt.Println("  exit              - 退出程序")

		default:
			fmt.Printf("未知命令: %s\n", command)
			fmt.Println("输入 'help' 查看支持的命令")
		}
	}
}
