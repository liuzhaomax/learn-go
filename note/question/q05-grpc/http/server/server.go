package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	schema "learn-go/note/question/q05-grpc/http"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/hello", func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}
		defer c.Request.Body.Close()
		var result schema.ReqBody
		err = json.Unmarshal(body, &result)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "hello " + result.Data,
		})
	})
	r.Run()
}
