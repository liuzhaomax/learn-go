package algorithm

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type test struct {
		input []int
		want  []int
	}

	tests := map[string]test{
		"0": {input: []int{5, 4, 3, 2, 1}, want: []int{1, 2, 3, 4, 5}},
		"1": {input: []int{1}, want: []int{1}},
		"2": {input: []int{5, 3, 3, 2, 1}, want: []int{1, 2, 3, 3, 5}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := QuickSort(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name: %v failed \n want: %v \n got: %v", name, tc.want, got)
			}
		})
	}
}

// go test -bench=QuickSort -benchmem
func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort([]int{5, 4, 3, 2, 1})
	}
}
