package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn-go/packages/consul"
)

func init() {
	err := consul.Reg("test1", "test_1", "172.26.160.1", "9527", []string{"test1"})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("成功")
	}
}

func main() {
	addr := "0.0.0.0:9527"
	r := gin.Default()
	r.GET("/health", consul.HealthHandler)
	r.Run(addr)
}
