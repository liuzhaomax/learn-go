/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/1/29 18:11
 * @version     v1.0
 * @filename    RemainderSorting.go
 * @description
 ***************************************************************************/
package main

import (
	"sort"
)

func RemainderSorting(strArr []string) []string {
	m := make([]int, len(strArr))
	mapp := make(map[string]int)
	for i := 0; i < len(strArr); i++ {
		m[i] = len(strArr[i]) % 3
		mapp[strArr[i]] = m[i]
	}
	sort.Ints(m)
	var res []string
	var tmp []string
	for i := 0; i < len(m); i++ {
		tmp = []string{}
		if i < len(m)-1 && m[i] == m[i+1] {
			continue
		}
		for k, v := range mapp {
			if v == m[i] {
				tmp = append(tmp, k)
			}
		}
		if len(tmp) == 1 {
			res = append(res, tmp[0])
		}
		if len(tmp) > 1 {
			sort.Strings(tmp)
			res = append(res, tmp...)
		}
	}
	return res
}

//func main() {
//	a := []string{
//		"AnBmMIPbs",
//		"ANBmMIPbs",
//		"Abb",
//	} //BUG
//	fmt.Println(RemainderSorting(a))
//}
