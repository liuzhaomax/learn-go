/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/7/31 15:49
 * @version     v1.0
 * @filename    post.go
 * @description
 ***************************************************************************/
package handler

import (
	"github.com/gin-gonic/gin"
	"learn-go/note/question/q03-gin/schema"
	"net/http"
)

func PostAdd(c *gin.Context) {
	var book schema.Book
	err := c.ShouldBindJSON(&book)
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
