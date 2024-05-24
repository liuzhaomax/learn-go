package algorithm

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	target := 8

	result := check1(arr, target)
	fmt.Println(result)
}

func check(arr []int, target int) bool {
	for _, elem := range arr {
		for _, elem2 := range arr {
			if elem+elem2 == target {
				return true
			}
		}
	}
	return false
}

func check1(arr []int, target int) bool {
	i := 0
	j := len(arr) - 1

	for {
		if i >= j {
			break
		}
		if arr[i]+arr[j] == target {
			return true
		}
		if arr[i]+arr[j] < target {
			i++
		}
		if arr[i]+arr[j] > target {
			j--
		}
	}
	return false
}
