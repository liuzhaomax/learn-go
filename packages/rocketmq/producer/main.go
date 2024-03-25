package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func CreateTopic(topicName string) {
	endPoint := []string{"106.15.94.179:9876"}
	// 创建主题
	// 先连接远程的服务器，得到一个具柄testAdmin，然后利用该具柄创建CreateTopic()创建topic
	testAdmin, err := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver(endPoint)))
	// 检查是否连接成功
	if err != nil {
		fmt.Printf("connection error: %s\n", err.Error())
	}
	err = testAdmin.CreateTopic(context.Background(), admin.WithTopicCreate(topicName))
	// 检查是否创建topic失败
	if err != nil {
		fmt.Printf("createTopic error: %s\n", err.Error())
	}
}

func main() {
	CreateTopic("testTopic")
	// 创建生产者实例
	producer1, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{"106.15.94.179:9876"}),
		producer.WithRetry(2), // 尝试发送数据的次数
		producer.WithGroupName("testGroup"),
	)
	if err != nil {
		panic(err)
	}

	// 启动生产者
	err = producer1.Start()
	if err != nil {
		panic(err)
	}
	defer producer1.Shutdown()

	// 构造消息
	msg := primitive.NewMessage("testTopic", []byte("Hello RocketMQ Go!"))

	// 发送消息
	result, err := producer1.SendSync(context.Background(), msg)
	if err != nil {
		fmt.Printf("Send message error: %s\n", err)
	} else {
		fmt.Printf("Send message success: %s\n", result.String())
	}
}
