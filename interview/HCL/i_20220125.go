/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/1/25 13:44
 * @version     v1.0
 * @filename    i_20220125.go
 * @description
 ***************************************************************************/
package HCL

import "fmt"

type Name struct {
	a int
	b int
	c int
}

func main() {
	name := &Name{
		a: 10,
		b: 20,
	}
	name.calc()
	fmt.Println(name.c)
	//var array [3]int
	//var slice = make([]int, 3, 6)
}

func (n *Name) calc() {
	n.c = n.a + n.b
}
