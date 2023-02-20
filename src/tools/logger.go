package tools

import "go.uber.org/zap"

var (
	Logger *zap.Logger
)

func InitLogger() {
	Logger, _ := zap.NewProduction()
	Logger.Debug("Logger init success")
}

func ShutdownLogger() {
	_ = Logger.Sync()
}

func init() {
	InitLogger()
}
