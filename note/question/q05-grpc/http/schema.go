/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/8/4 21:37
 * @version     v1.0
 * @filename    schema.go
 * @description
 ***************************************************************************/
package schema

type Result struct {
	Msg string `json:"msg"`
}

type ReqBody struct {
	Data string `json:"data"`
}
