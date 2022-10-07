package main

import (
	"go.uber.org/zap"
	"time"
)

func NewLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = append(config.OutputPaths, "./packages/zap/logger")
	return config.Build()
}

func main() {
	//logger, err := zap.NewProduction()
	logger, err := NewLogger() // 输出到disk
	//zap.ReplaceGlobals(logger) // 替换全局logger
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Info("error:", "abc", "time:", time.Now().Unix())
}
