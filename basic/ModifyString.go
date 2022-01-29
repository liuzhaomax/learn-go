/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/1/29 15:39
 * @version     v1.0
 * @filename    ModifyString.go
 * @description
 ***************************************************************************/
package main

import (
	"strings"
	"unicode"
)

func ModifyString(str string) string {
	arr := strings.Split(str, " ")
	trim := strings.Join(arr, "")
	noD := ""
	for i := 0; i < len(trim); i++ {
		if unicode.IsDigit(rune(trim[i])) == false {
			noD += string(trim[i])
		}
	}
	resArr := []rune(noD)
	for i, j := 0, len(resArr)-1; i < j; i, j = i+1, j-1 {
		resArr[i], resArr[j] = resArr[j], resArr[i]
	}
	return string(resArr)
}

//func main() {
//	a := ModifyString("oll123eH ")
//	fmt.Println(reflect.TypeOf(a[0]))
//}
