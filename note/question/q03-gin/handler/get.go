/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/7/31 15:18
 * @version     v1.0
 * @filename    get.go
 * @description
 ***************************************************************************/
package handler

import (
	"github.com/gin-gonic/gin"
	"learn-go/note/question/q03-gin/schema"
	"net/http"
)

func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func GetIDAndName(c *gin.Context) {
	var book schema.Book
	err := c.ShouldBindUri(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "missing params",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id":   book.ID,
			"name": book.Name,
		})
	}
}

func GetQueryIDAndName(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")
	if id == "" || name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "missing params",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id":   id,
			"name": name,
		})
	}
}
