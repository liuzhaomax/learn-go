package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func main() {
	// 创建消费者实例
	consumer1, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"106.15.94.179:9876"}),
		consumer.WithGroupName("testGroup"),
	)
	if err != nil {
		panic(err)
	}

	// 设置消息监听器
	err = consumer1.Subscribe("testTopic", consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range msgs {
			fmt.Printf("Received message: %s\n", msg.Body)
		}
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		panic(err)
	}

	// 启动消费者
	err = consumer1.Start()
	if err != nil {
		panic(err)
	}
	defer consumer1.Shutdown()

	// 保持主程序不退出
	select {}
}
