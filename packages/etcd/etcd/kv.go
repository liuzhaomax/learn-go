package etcd

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func KvDemo() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   GetEtcdEndpoint(),
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// 添加KV
	res, err := client.Put(context.Background(), "key1", "value", clientv3.WithPrevKV())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Put Prev Val is: %s \n", res.PrevKv)

	res1, err := client.Put(context.Background(), "key1", "value", clientv3.WithPrevKV())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Put Prev Val is: %s \n", res1.PrevKv)

	// 获取KV
	getRes, err := client.Get(context.Background(), "key", clientv3.WithPrefix())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Get Res is: %s \n", getRes.Kvs[0].Value)

	// 删除kv
	delRes, err := client.Delete(context.Background(), "key", clientv3.WithPrefix())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Delete amount is: %d \n", delRes.Deleted)
}
