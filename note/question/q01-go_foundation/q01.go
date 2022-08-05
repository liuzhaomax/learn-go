/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/7/21 0:55
 * @version     v1.0
 * @filename    q01-go_foundation.go
 * @description
 ***************************************************************************/
package main

import (
	"fmt"
)

func main() {
	s := []string{"炭烤生蚝", "麻辣小龙虾", "干锅鸭"}
	s2 := make([]*string, len(s))
	for i, v := range s {
		s2[i] = &v
	}
	fmt.Println(s2)
	s2 = make([]*string, len(s))
	for i, v := range s {
		val := v
		s2[i] = &val
	}
	fmt.Println(*s2[0])
	fmt.Println(*s2[1])
	fmt.Println(*s2[2])
	s2 = make([]*string, len(s))
	for i, _ := range s {
		s2[i] = &s[i]
	}
	fmt.Println(*s2[0])
	fmt.Println(*s2[1])
	fmt.Println(*s2[2])

	fn := func() func() {
		aaa := 1
		return func() {
			fmt.Println(aaa)
		}
	}
	fnC := fn()
	fnC()

}
