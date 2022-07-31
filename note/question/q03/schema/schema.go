/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/7/31 15:19
 * @version     v1.0
 * @filename    schema.go
 * @description
 ***************************************************************************/
package schema

type Book struct {
	ID   string `uri:"id" json:"id" binding:"required"`
	Name string `uri:"name" json:"name" binding:"required"`
}
