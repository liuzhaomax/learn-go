package main

import (
	"strings"
)

/*
 * Complete the 'getAnagram' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func getAnagram(s string) int32 {
	if len(s) == 0 || len(s) == 1 {
		return 0
	}
	var res = int32(len(s) / 2)
	first := s[0 : len(s)/2]
	second := s[len(s)/2:]
	var m = make(map[string]int32)
	for i := 0; i < len(second); i++ {
		m[string(second[i])] += 1
	}
	for i := 0; i < len(first); i++ {
		if strings.Contains(second, string(first[i])) && m[string(first[i])] > 0 {
			m[string(first[i])] -= 1
			res--
		}
	}
	return res
}

//func main() {
//	s := "123456"
//
//	fmt.Println(getAnagram(s))
//}
