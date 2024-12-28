package main

import (
	"TAROT/bootstrap"
	"TAROT/config"
	"TAROT/pkg/logger"
	"TAROT/routes"
	"TAROT/service/chat"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// App 应用程序上下文，用于优雅关闭
type App struct {
	server *http.Server
}

// 添加全局 logger
var appLogger *logger.Logger

// setupApplication 初始化应用程序所需的各种组件
func setupApplication() error {
	// 加载配置
	_, err := config.LoadConfig()
	if err != nil {
		return err
	}

	// 初始化日志
	appLogger, err = logger.NewLogger()
	if err != nil {
		return err
	}

	// 初始化应用
	app := bootstrap.NewApplication()
	defer app.Close()

	return nil
}

// setupServer 配置并返回 HTTP 服务器和处理器
func setupServer(chatService *chat.Service) *http.Server {
	// 注册路由
	handler := routes.RegisterRoutes(chatService)

	// 创建 HTTP 服务器
	server := &http.Server{
		Addr:    ":3000",
		Handler: handler,
	}

	return server
}

// start 启动服务器并处理优雅关闭
func (a *App) start() {
	// 创建系统信号监听器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		appLogger.Printf("服务器正在启动，监听端口%s\n", a.server.Addr)
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			appLogger.Fatalf("服务器启动失败: %v", err)
		}
	}()

	// 等待中断信号
	<-quit
	appLogger.Println("正在关闭服务器...")

	// 创建一个带超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 优雅关闭服务器
	if err := a.server.Shutdown(ctx); err != nil {
		appLogger.Fatalf("服务器关闭异常: %v", err)
	}

	appLogger.Println("服务器已成功关闭")
}

func main() {
	// 初始化应用程序
	if err := setupApplication(); err != nil {
		appLogger.Fatalf("初始化应用程序失败: %v", err)
	}

	// 初始化应用
	bootstrapApp := bootstrap.NewApplication()
	defer bootstrapApp.Close()

	// 创建并配置服务器
	server := setupServer(bootstrapApp.ChatService)

	// 创建应用实例
	app := &App{
		server: server,
	}

	// 启动服务器（包含优雅关闭）
	app.start()
}
