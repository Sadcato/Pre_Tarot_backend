package chat

import (
	"TAROT/pkg/arkapi"
	"TAROT/pkg/redis"
	"context"
	"encoding/json"
	"os"
	"time"
)

type Service struct {
	arkClient   *arkapi.Client
	redisClient *redis.Client
}

type SystemPrompt struct {
	SystemPrompt string  `json:"system_prompt"`
	Temperature  float64 `json:"temperature"`
	MaxTokens    int     `json:"max_tokens"`
}

func NewService(arkClient *arkapi.Client, redisClient *redis.Client) *Service {
	return &Service{
		arkClient:   arkClient,
		redisClient: redisClient,
	}
}

func (s *Service) Chat(ctx context.Context, userMessage string) (string, error) {
	// 尝试从缓存获取响应
	cacheKey := "chat:" + userMessage
	if response, err := s.redisClient.Get(ctx, cacheKey); err == nil {
		return response, nil
	}

	// 读取系统提示词
	promptFile, err := os.ReadFile("prompt/prompt.json")
	if err != nil {
		return "", err
	}

	var systemPrompt SystemPrompt
	if err := json.Unmarshal(promptFile, &systemPrompt); err != nil {
		return "", err
	}

	// 调用AI服务
	response, err := s.arkClient.Chat(ctx, systemPrompt.SystemPrompt, userMessage)
	if err != nil {
		return "", err
	}

	// 缓存响应，设置20秒过期时间
	if err := s.redisClient.Set(ctx, cacheKey, response, 20*time.Second); err != nil {
		return "", err
	}

	return response, nil
}
