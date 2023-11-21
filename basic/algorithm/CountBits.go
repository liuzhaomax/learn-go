package algorithm

import (
	"strconv"
)

func CountBits(num uint32) int32 {
	if num == 0 {
		return 0
	}
	var bin = ""
	for ; num > 0; num = num / 2 {
		bin = strconv.Itoa(int(num)%2) + bin
	}
	var res int32
	for i := 0; i < len(bin); i++ {
		if bin[i] == '1' {
			res++
		}
	}
	return res
}

//func main() {
//	a := CountBits(126)
//	fmt.Println(a)
//}
