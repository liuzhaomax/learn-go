/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/7/24 14:41
 * @version     v1.0
 * @filename    func_params.go
 * @description
 ***************************************************************************/
package main

type A struct {
	field string
}

func exec(a A) {
	a.field = "world"
}

func exec1(a *A) {
	a.field = "world"
}

//func main() {
//	a := A{
//		field: "hello",
//	}
//	exec(a)
//	fmt.Println(a.field)
//	exec1(&a)
//	fmt.Println(a.field)
//}
