/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/7/23 21:49
 * @version     v1.0
 * @filename    duck.go
 * @description
 ***************************************************************************/
package main

import "fmt"

type Duck struct {
}

func (d *Duck) run() {
	fmt.Println("duck running")
}

func (d *Duck) fly() {
	fmt.Println("duck flying")
}

type Plane struct {
}

func (p *Plane) fly() {
	fmt.Println("plane flying")
}

func Happy(v Go) {
	v.run()
	v.fly()
}

type Go interface {
	run()
	fly()
}

//func main() {
//	duck := Duck{}
//	//plane := Plane{}
//	Happy(&duck)
//	//Happy(&plane)
//}
