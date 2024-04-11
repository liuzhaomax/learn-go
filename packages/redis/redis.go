package main

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"time"
)

const (
	redisPassword = "123456"
)

func InitRedis() (*redis.Client, func(), error) {
	client, clean, err := LoadRedis()
	if err != nil {
		return nil, clean, err
	}
	return client, clean, nil
}

func LoadRedis() (*redis.Client, func(), error) {
	ctx := context.Background()
	addr := fmt.Sprintf("%s:%s", "127.0.0.1", "6379")
	opts := &redis.Options{
		Network:               "",
		Addr:                  addr,
		ClientName:            "",
		Dialer:                nil,
		OnConnect:             nil,
		Protocol:              0,
		Username:              "",
		Password:              redisPassword,
		CredentialsProvider:   nil,
		DB:                    0,
		MaxRetries:            0,
		MinRetryBackoff:       0,
		MaxRetryBackoff:       0,
		DialTimeout:           0,
		ReadTimeout:           0,
		WriteTimeout:          0,
		ContextTimeoutEnabled: false,
		PoolFIFO:              false,
		PoolSize:              0,
		PoolTimeout:           0,
		MinIdleConns:          0,
		MaxIdleConns:          0,
		MaxActiveConns:        0,
		ConnMaxIdleTime:       0,
		ConnMaxLifetime:       0,
		TLSConfig:             nil,
		Limiter:               nil,
		DisableIndentity:      false,
	}
	client := redis.NewClient(opts)
	err := client.Ping(ctx).Err()
	if err != nil {
		return nil, nil, err
	}
	// 设置必备键的过期时间
	err = client.Expire(ctx, "signature", 24*time.Hour).Err()
	if err != nil {
		return nil, nil, err
	}
	clean := func() {
		err := client.Close()
		if err != nil {
		}
	}
	return client, clean, nil
}

func main() {
	client, _, _ := InitRedis()
	result, err := client.SAdd(context.Background(), "signature", "123").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
