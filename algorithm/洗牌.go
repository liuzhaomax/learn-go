package algorithm

import (
	"math/rand"
)

func Shuffle(strSlice []string) []string {
	// 使用 Fisher-Yates 算法随机排列切片中的元素
	for i := len(strSlice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		strSlice[i], strSlice[j] = strSlice[j], strSlice[i]
	}
	return strSlice
}
