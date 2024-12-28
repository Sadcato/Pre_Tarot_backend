package bootstrap

import (
	"TAROT/pkg/arkapi"
	"TAROT/pkg/redis"
	"TAROT/service/chat"
)

type Application struct {
	ArkClient   *arkapi.Client
	RedisClient *redis.Client
	ChatService *chat.Service
}

func NewApplication() *Application {
	// 初始化客户端
	arkClient := arkapi.NewClient()
	redisClient := redis.NewClient()

	// 初始化服务
	chatService := chat.NewService(arkClient, redisClient)

	return &Application{
		ArkClient:   arkClient,
		RedisClient: redisClient,
		ChatService: chatService,
	}
}

func (app *Application) Close() error {
	return app.RedisClient.Close()
}
