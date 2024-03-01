package time

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	t1 := time.Date(2024, time.February, 29, 12, 0, 0, 0, time.UTC)
	t1Str := t1.Format(time.RFC3339)
	longer, err := HasPassedGivenHours(t1Str, 5)
	if err != nil {
		panic(err)
	}
	fmt.Printf("longer: %t\n", longer)
}
