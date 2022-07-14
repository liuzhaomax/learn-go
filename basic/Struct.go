/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/7/15 0:19
 * @version     v1.0
 * @filename    Struct.go
 * @description
 ***************************************************************************/
package main

import "fmt"

type Person struct {
	Request Request
}

type Request struct {
	Name string
}

func main() {
	aaa := Person{}
	fmt.Println(aaa.Request == Request{}) // true
}
