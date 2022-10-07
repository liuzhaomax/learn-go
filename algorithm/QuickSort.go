package algorithm

import (
	"math"
)

func QuickSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	pivotIndex := int(math.Floor(float64(length) / 2))
	pivot := arr[pivotIndex]
	arr = append(arr[:pivotIndex], arr[(pivotIndex+1):]...)
	left := make([]int, 0, length)
	right := make([]int, 0, length)
	for i := 0; i < length-1; i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	return append(append(QuickSort(left), pivot), QuickSort(right)...)
}
