/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/7/31 14:52
 * @version     v1.0
 * @filename    gin.go
 * @description
 ***************************************************************************/
package main

import (
	"github.com/gin-gonic/gin"
	"learn-go/note/question/q03-gin/handler"
)

func main() {
	r := gin.Default()

	r.GET("/", handler.GetIndex)
	r.GET("/:id/:name", handler.GetIDAndName)
	r.GET("/details", handler.GetQueryIDAndName)
	r.POST("/add", handler.PostAdd)

	r.Run()
}
