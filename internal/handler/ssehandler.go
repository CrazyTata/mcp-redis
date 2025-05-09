package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"mcp-redis/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SSEHandler struct {
	svcCtx *svc.ServiceContext
}

func NewSSEHandler(svcCtx *svc.ServiceContext) *SSEHandler {
	return &SSEHandler{
		svcCtx: svcCtx,
	}
}

type Command struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type Response struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

func (h *SSEHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 设置 SSE 头部
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// 创建 SSE 编码器
	encoder := json.NewEncoder(w)

	// 处理命令
	decoder := json.NewDecoder(r.Body)
	var cmd Command
	if err := decoder.Decode(&cmd); err != nil {
		logx.Errorf("Failed to decode command: %v", err)
		return
	}

	// 根据命令类型处理
	switch cmd.Type {
	case "kv.get":
		var req struct {
			Key string `json:"key"`
		}
		if err := json.Unmarshal(cmd.Payload, &req); err != nil {
			logx.Errorf("Failed to unmarshal payload: %v", err)
			return
		}
		val, err := h.svcCtx.Redis.Get(req.Key)
		if err != nil {
			logx.Errorf("Failed to get key: %v", err)
			return
		}
		encoder.Encode(Response{
			Type:    "kv.get.response",
			Payload: val,
		})

	case "kv.set":
		var req struct {
			Key   string `json:"key"`
			Value string `json:"value"`
			TTL   int    `json:"ttl,optional"`
		}
		if err := json.Unmarshal(cmd.Payload, &req); err != nil {
			logx.Errorf("Failed to unmarshal payload: %v", err)
			return
		}
		err := h.svcCtx.Redis.Setex(req.Key, req.Value, req.TTL)
		if err != nil {
			logx.Errorf("Failed to set key: %v", err)
			return
		}
		encoder.Encode(Response{
			Type:    "kv.set.response",
			Payload: "success",
		})

	// 可以继续添加其他命令类型的处理...
	default:
		logx.Errorf("Unknown command type: %s", cmd.Type)
		return
	}

	// 保持连接
	flusher, ok := w.(http.Flusher)
	if !ok {
		logx.Error("Streaming unsupported!")
		return
	}

	// 定期发送心跳
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			encoder.Encode(Response{
				Type:    "heartbeat",
				Payload: time.Now().Unix(),
			})
			flusher.Flush()
		case <-r.Context().Done():
			return
		}
	}
}
