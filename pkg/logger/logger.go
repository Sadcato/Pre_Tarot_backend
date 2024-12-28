package logger

import (
	"log"
	"os"
	"path/filepath"
)

type Logger struct {
	*log.Logger
}

func NewLogger() (*Logger, error) {
	// 确保日志目录存在
	if err := os.MkdirAll("storage/logs", 0755); err != nil {
		return nil, err
	}

	// 打开日志文件
	logFile, err := os.OpenFile(
		filepath.Join("storage/logs", "app.log"),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)
	if err != nil {
		return nil, err
	}

	return &Logger{
		Logger: log.New(logFile, "", log.LstdFlags),
	}, nil
}
