package main

import (
	"context"
	"github.com/sirupsen/logrus"
)

func main() {
	// 创建一个 Logrus 日志实例
	logger := logrus.New()

	// 示例 1: 添加字段
	logger.WithField("key", "value").Info("Logging with a field")

	// 示例 2: 添加多个字段
	logger.WithFields(logrus.Fields{
		"field1": "value1",
		"field2": "value2",
	}).Warn("Logging with multiple fields")

	// 示例 3: 使用上下文
	ctx := context.WithValue(context.Background(), "user_id", 123)
	logger.WithContext(ctx).Info("Logging with context")

	// 示例 4: 结合字段和上下文
	logger.WithField("key", "value").WithContext(ctx).Error("Logging with both field and context")
}
