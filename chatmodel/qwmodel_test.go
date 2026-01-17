package chatmodel

import (
	"context"
	"fmt"
	"github.com/cloudwego/eino-ext/components/model/qwen"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"testing"
)

func init() {
	_ = godotenv.Load("../.env")
}

var ctx = context.Background()

// 获取千问模型
func getQwModeClient() *qwen.ChatModel {
	//timeOut, _ := strconv.Atoi(os.Getenv("MODEL_TIME_OUT"))
	maxToken, _ := strconv.Atoi(os.Getenv("QW_RESP_MAX_TOKENS"))
	m, err := qwen.NewChatModel(context.Background(), &qwen.ChatModelConfig{
		APIKey:      os.Getenv("QW_KEY"),
		BaseURL:     "https://dashscope.aliyuncs.com/compatible-mode/v1",
		Model:       os.Getenv("QW_LLM_CHAT"),
		MaxTokens:   of(maxToken),
		Temperature: of(float32(0.5)),
		TopP:        of(float32(0.5)),
	})
	if err != nil {
		panic(err)
	}
	return m
}

// Description: 千问模型文本生成
// Author: LiuQHui
// Param t
// Date 2026-01-17 23:30:55
func TestQwModelGenerate(t *testing.T) {
	// 获取模型
	m := getQwModeClient()
	msgs := []*schema.Message{
		schema.SystemMessage("您现在扮演一名幽默家，给用户讲一些有趣且富有情感的故事"),
		schema.UserMessage("讲一个植物人相关的故事"),
	}
	outMsg, err := m.Generate(ctx, msgs)
	fmt.Println(err)
	fmt.Println(outMsg.Content)

}

func of[T any](n T) *T {
	return &n
}
