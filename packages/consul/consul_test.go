package consul

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

func init() {
	err := Reg("test1", "test_1", "172.26.160.1", "9208", []string{"test1"})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("成功")
	}
}

func TestGin(t *testing.T) {
	addr := "0.0.0.0:9208"
	r := gin.Default()
	r.GET("/health", HealthHandler)
	r.Run(addr)
}
