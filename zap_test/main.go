package main

import (
	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{
		"stdout",
		"./xxx.log",
	}
	return cfg.Build()
}

func main() {
	// logger, _ := zap.NewProduction()
	// logger, _ := zap.NewDevelopment()
	logger, _ := NewLogger()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		"url", "http://example.com",
		"attempt", 3,
	)
}
