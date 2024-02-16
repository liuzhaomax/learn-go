package algorithm

import (
	"fmt"
	"testing"
)

func TestShuffle(t *testing.T) {
	strSlice := []string{"apple", "banana", "orange", "grape", "peach"}
	fmt.Println(Shuffle(strSlice))
}
