package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "43.14 567 agsdg 1.23 7. 8.9 1sdljgl 6.66 7.8   "
	//解释正则表达式
	reg := regexp.MustCompile(`\d+\.\d+`)
	if reg == nil {
		fmt.Println("MustCompile err")
		return
	}
	//提取关键信息
	result := reg.Match([]byte(buf))
	//result := reg.FindAllString(buf, -1)
	//result := reg.FindAllStringSubmatch(buf, -1)
	fmt.Println("result = ", result)

	buff := "992DC-2L3OD-SOKO2P390IPD9-SS1"
	regg := regexp.MustCompile(`^(\d|\w)[0-9A-Z-]+(\d|\w)$`)
	if regg == nil {
		fmt.Println("MustCompile err")
		return
	}
	res := regg.Match([]byte(buff))
	fmt.Println("res = ", res)
}
